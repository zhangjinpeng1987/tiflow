// Copyright 2023 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package csv

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/pingcap/errors"
	parserconfig "github.com/pingcap/tidb/br/pkg/lightning/config"
	"github.com/pingcap/tidb/br/pkg/lightning/mydump"
	"github.com/pingcap/tidb/br/pkg/lightning/worker"
	"github.com/pingcap/tiflow/pkg/sink/cloudstorage"
	"github.com/pingcap/tiflow/pkg/sink/codec/common"
	"github.com/pingcap/tidb/types"
)

// The csv files are stored in the following directory structure:
// $ tree
// .
// ├── db1
// │   ├── meta
// │   │   └── schema_445777854499389449_1495827732.json
// │   └── table1
// │       ├── 445777864526397449
// │       │   └── 2023-11-20
// │       │       ├── CDC00000000000000000001.csv
// │       │       ├── CDC00000000000000000002.csv
// │       │       └── meta
// │       │           └── CDC.index
// │       ├── 445777908658339860
// │       │   └── 2023-11-20
// │       │       ├── CDC00000000000000000001.csv
// │       │       └── meta
// │       │           └── CDC.index
// │       └── meta
// │           ├── schema_445777864526397449_2572608321.json
// │           ├── schema_445777908658339860_2572608321.json
// │           └── schema_445777975348035588_2572608321.json
// └── metadata
// ...
//
// metadata file
// The 'metadata' file contains the checkpoint-ts of the whole task,
// it looks like this:
// {"checkpoint-ts":445777802503389188}
//
// db1.meta.schema_445777854499389449_1495827732.json
// The schema file under db directory contains the schema of the db, it looks like this:
// {
//     "Table": "",
//     "Schema": "cdcdb",
//     "Version": 1,
//     "TableVersion": 445777854499389449,
//     "Query": "CREATE DATABASE `cdcdb`",
//     "Type": 1,
//     "TableColumns": null,
//     "TableColumnsTotal": 0
// }
//
// db1.table1.meta.schema_445777864526397449_2572608321.json
// The schema file under table directory contains DDL information of table, it looks like this:
// {
//     "Table": "t1",
//     "Schema": "cdcdb",
//     "Version": 1,
//     "TableVersion": 445777864526397449,
//     "Query": "CREATE TABLE `t1` (`id` INT(10) PRIMARY KEY NOT NULL,`name` VARCHAR(255))",
//     "Type": 3,
//     "TableColumns": [
//         {
//             "ColumnName": "id",
//             "ColumnType": "INT",
//             "ColumnPrecision": "10",
//             "ColumnNullable": "false",
//             "ColumnIsPk": "true"
//         },
//         {
//             "ColumnName": "name",
//             "ColumnType": "VARCHAR",
//             "ColumnPrecision": "255"
//         }
//     ],
//     "TableColumnsTotal": 2
// }
//
// db1.table1.445777864526397449.2023-11-20.meta.CDC.index
// The index file under table directory contains the latest index of the csv files, it looks like this:
// "CDC00000000000000000002.csv"
//

const (
	// OpType, Schema, Table, [CommitTS], Columns...
	CSVMetaColCnt = 3
)

// CsvSchemaReader is the schema reader
type CsvSchemaReader struct {
	// The file path of the schema file
	filePath string
}

func NewCsvSchemaReader(filePath string) (*CsvSchemaReader, error) {
	// Verify the file path is a valid schema file.
	// If the file path is not a valid schema file, return error.
	if !isValidSchemaFile(filePath) {
		return nil, errors.Errorf("invalid schema file path %s", filePath)
	}

	return &CsvSchemaReader{
		filePath: filePath,
	}, nil
}

func (reader *CsvSchemaReader) TableSchema() (*cloudstorage.TableDefinition, error) {
	// Open the schema file.
	handler, err := os.Open(reader.filePath)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer handler.Close()

	tableDef := &cloudstorage.TableDefinition{}
	err = json.NewDecoder(handler).Decode(tableDef)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return tableDef, nil
}

func isValidCsvFile(filePath string) bool {
	return strings.HasSuffix(filePath, ".csv")
}

func isValidSchemaFile(filePath string) bool {
	return strings.HasSuffix(filePath, ".json")
}

// CsvFileReader is the csv reader, each reader is responsible for reading a csv file
// and sending the data to the downstream.
// The csv file is read line by line, and each line is sent to the downstream
// as a message.
type CsvFileReader struct {
	// The file path of the csv file
	filePath string

	// The paser of the csv file.
	cvsParser *mydump.CSVParser

	codecConfig *common.Config
}

// NewCsvReader creates a new csv reader.
func NewCsvReader(ctx context.Context, filePath string, codecConfig *common.Config) (*CsvFileReader, error) {
	// Verify the file path is a valid csv file.
	// If the file path is not a valid csv file, return error.
	if !isValidCsvFile(filePath) {
		return nil, errors.Errorf("invalid csv file path %s", filePath)
	}

	// Read all content from the CSV file to the buf.
	// The max size of cloud storage data file is limited , the default max file size is 64MB.
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var backslashEscape bool

	// if quote is not set in config, we should unespace backslash
	// when parsing csv columns.
	if len(codecConfig.Quote) == 0 {
		backslashEscape = true
	}
	cfg := &parserconfig.CSVConfig{
		Separator:       codecConfig.Delimiter,
		Delimiter:       codecConfig.Quote,
		Terminator:      codecConfig.Terminator,
		Null:            []string{codecConfig.NullString},
		BackslashEscape: backslashEscape,
		// Todo: use csvMessage related escape function
	}
	csvParser, err := mydump.NewCSVParser(ctx, cfg, mydump.NewStringReader(string(b)),
		512*1024, /*block size*/
		worker.NewPool(ctx, 1, "io"), false, nil)
	if err != nil {
		return nil, err
	}

	return &CsvFileReader{
		filePath:  filePath,
		cvsParser: csvParser,
	}, nil
}

func (reader *CsvFileReader) ReadNextRow() (bool, error) {
	// Read the next line from the csv file.
	// If the line is needed, return the line.
	// If the line is not needed, return the next line.
	err := reader.cvsParser.ReadRow()
	if err != nil {
		if errors.Cause(err) == io.EOF {
			return false, nil
		}
		return false, errors.Trace(err)
	}

	return true, nil
}

// Row returns the last row read by ReadNextRow.
func (reader *CsvFileReader) Row() (mydump.Row, error) {
	return reader.cvsParser.LastRow(), nil
}

type CSVReader struct {
	codecConfig *common.Config
}

func NewCSVReader(cfg *common.Config) *CSVReader {
	return &CSVReader{
		codecConfig: cfg,
	}
}

type RowFilter interface {
	// Filter returns true if the row should be filtered out.
	Filter(row mydump.Row, schema *cloudstorage.TableDefinition) bool
}

// Search searches rows/operations that satisfiy all filters/conditions.
func (reader *CSVReader) Search(schemaFile, csvFile string, filters ...RowFilter) ([]mydump.Row, error) {
	schemaReader, err := NewCsvSchemaReader(schemaFile)
	if err != nil {
		return nil, errors.Trace(err)
	}
	schema, err := schemaReader.TableSchema()
	if err != nil {
		return nil, errors.Trace(err)
	}
	err = schema.BuildColNameToIdx()
	if err != nil {
		return nil, errors.Trace(err)
	}

	expectedColCnt := len(schema.Columns) + CSVMetaColCnt
	if reader.codecConfig.IncludeCommitTs {
		expectedColCnt++
	}

	ctx := context.Background()
	fileReader, err := NewCsvReader(ctx, csvFile, reader.codecConfig)
	if err != nil {
		return nil, errors.Trace(err)
	}

	// Read the csv file row by row, match the row with all filters.
	result := make([]mydump.Row, 1024)
	valid, err := fileReader.ReadNextRow()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for valid {
		row, err := fileReader.Row()
		if err != nil {
			return nil, errors.Trace(err)
		}

		if len(row.Row) != expectedColCnt {
			return nil, errors.Errorf("row column count not match schema's column count, row %v, schema %v", row, *schema)
		}

		// Filter the row.
		want := true
		for _, filter := range filters {
			if !filter.Filter(row, schema) {
				want = false
				break
			}
		}
		if want {
			result = append(result, row)
		}

		valid, err = fileReader.ReadNextRow()
		if err != nil {
			return nil, errors.Trace(err)
		}
	}

	return result, nil
}

type MatchOP int

const (
	Equal MatchOP = iota
	NotEqual
	LessThan
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual
	Cotains
)

type ColumnMatchFilter struct {
	ColumnName string
	Value     types.Datum
	Match      MatchOP
}

func (f *ColumnMatchFilter) Filter(row mydump.Row, schema *cloudstorage.TableDefinition) bool {
	colIdx, ok := schema.ColNameToIdx[f.ColumnName]
	if !ok || colIdx < 0 {
		return false
	}
	colType := schema.Columns[colIdx].Tp
	colVal := row.Row[colIdx]

	switch colType {
	case "INT":
		return f.matchInt(colVal)
	case "VARCHAR", "TEXT", "CHAR", "BLOB", "JSON":
		return f.matchString(colVal)
	default:
		return false
	}
}

func (f *ColumnMatchFilter) matchInt(colVal types.Datum) bool {
	val := colVal.GetInt64()
	switch f.Match {
	case Equal:
		return val == f.Value.GetInt64()
	case NotEqual:
		return val != f.Value.GetInt64()
	case LessThan:
		return val < f.Value.GetInt64()
	case LessThanOrEqual:
		return val <= f.Value.GetInt64()
	case GreaterThan:
		return val > f.Value.GetInt64()
	case GreaterThanOrEqual:
		return val >= f.Value.GetInt64()
	default:
		return false
	}
}

func (f *ColumnMatchFilter) matchString(colVal types.Datum) bool {
	val := colVal.GetString()
	switch f.Match {
	case Equal:
		return val == f.Value.GetString()
	case NotEqual:
		return val != f.Value.GetString()
	case Cotains:
		return strings.Contains(val, f.Value.GetString())
	default:
		return false
	}
}

type CommitTsFilter struct {
	CommitTs uint64
	Match    MatchOP
}

func (f *CommitTsFilter) Filter(row mydump.Row, schema *cloudstorage.TableDefinition) bool {
	if len(row.Row) != len(schema.Columns) + CSVMetaColCnt + 1 {
		return false
	}

	colVal := row.Row[CSVMetaColCnt]
	val := colVal.GetUint64()
	switch f.Match {
	case Equal:
		return val == f.CommitTs
	case NotEqual:
		return val != f.CommitTs
	case LessThan:
		return val < f.CommitTs
	case LessThanOrEqual:
		return val <= f.CommitTs
	case GreaterThan:
		return val > f.CommitTs
	case GreaterThanOrEqual:
		return val >= f.CommitTs
	default:
		return false
	}
}

type OptypeFilter struct {
	OpType string
	Match MatchOP
}

func (f *OptypeFilter) Filter(row mydump.Row, schema *cloudstorage.TableDefinition) bool {
	colVal := row.Row[0]
	val := colVal.GetString()
	switch f.Match {
	case Equal:
		return val == f.OpType
	case NotEqual:
		return val != f.OpType
	default:
		return false
	}
}