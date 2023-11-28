package csv

import (
	"context"
	"testing"
	"github.com/pingcap/tiflow/pkg/sink/cloudstorage"
	"encoding/json"
	"io/ioutil"
	"github.com/pingcap/tiflow/pkg/sink/codec/common"
	"os"
)

// TestCSVReaderSearch tests the search rows match all filters
func TestCSVReaderSearch(t *testing.T) {
	// Create a temporary file in /tmp directory with schema content
	// and remove it after the test
	schemaFile, err := ioutil.TempFile("/tmp", "csv_reader_test_search_*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(schemaFile.Name())
	defer schemaFile.Close()

	schemaFile.WriteString(`{
		"Table": "t1",
		"Schema": "db1",
		"Version": 1,
		"TableVersion": 445778110496374786,
		"Query": "",
		"Type": 0,
		"TableColumns": [
			{
				"ColumnName": "id",
				"ColumnType": "INT",
				"ColumnPrecision": "10",
				"ColumnNullable": "false",
				"ColumnIsPk": "true"
			},
			{
				"ColumnName": "name",
				"ColumnType": "VARCHAR",
				"ColumnPrecision": "255"
			}
		],
		"TableColumnsTotal": 2
	}`)
	schemaFile.Sync()

	csvFile, err := ioutil.TempFile("/tmp", "csv_reader_test_search_*.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(csvFile.Name())
	defer csvFile.Close()

	// OP, Schema, Table, CommitTs, id, name
	csvFile.WriteString(`
		I, db1, t1, 445778110496374790, 1, abc
		I, db1, t1, 445778110496374790, 2, def
		U, db1, t1, 445778110496374790, 3, ghi
		D, db1, t1, 445778110496374790, 4, jkl
		D, db1, t1, 445778110496374792, 5, jkl
		D, db1, t1, 445778110496374792, 6, jkl
		D, db1, t1, 445778110496374792, 7, jkl
		D, db1, t1, 445778110496374792, 8, jkl
	`)
	csvFile.Sync()

	// Create a CSVReader with the schema and data files
	cfg := &common.Config{
		Delimiter: ",",
		Quote:    "",
		NullString: "NULL",
		IncludeCommitTs: true,
	}
	csvReader := NewCSVReader(cfg)

	// Without any filters, all rows should be returned
	res, err := csvReader.Search(schemaFile.Name(), csvFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 8 {	
		t.Fatalf("Expected 8 rows, got %d", len(res))
	}

	// 
}
