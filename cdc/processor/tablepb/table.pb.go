// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: processor/tablepb/table.proto

package tablepb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_pingcap_tiflow_cdc_model "github.com/pingcap/tiflow/cdc/model"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// TableState is the state of table replication in processor.
//
//	┌────────┐   ┌───────────┐   ┌──────────┐
//	│ Absent ├─> │ Preparing ├─> │ Prepared │
//	└────────┘   └───────────┘   └─────┬────┘
//	                                   v
//	┌─────────┐   ┌──────────┐   ┌─────────────┐
//	│ Stopped │ <─┤ Stopping │ <─┤ Replicating │
//	└─────────┘   └──────────┘   └─────────────┘
type TableState int32

const (
	TableStateUnknown     TableState = 0
	TableStateAbsent      TableState = 1
	TableStatePreparing   TableState = 2
	TableStatePrepared    TableState = 3
	TableStateReplicating TableState = 4
	TableStateStopping    TableState = 5
	TableStateStopped     TableState = 6
)

var TableState_name = map[int32]string{
	0: "Unknown",
	1: "Absent",
	2: "Preparing",
	3: "Prepared",
	4: "Replicating",
	5: "Stopping",
	6: "Stopped",
}

var TableState_value = map[string]int32{
	"Unknown":     0,
	"Absent":      1,
	"Preparing":   2,
	"Prepared":    3,
	"Replicating": 4,
	"Stopping":    5,
	"Stopped":     6,
}

func (x TableState) String() string {
	return proto.EnumName(TableState_name, int32(x))
}

func (TableState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae83c9c6cf5ef75c, []int{0}
}

type Checkpoint struct {
	CheckpointTs github_com_pingcap_tiflow_cdc_model.Ts `protobuf:"varint,1,opt,name=checkpoint_ts,json=checkpointTs,proto3,casttype=github.com/pingcap/tiflow/cdc/model.Ts" json:"checkpoint_ts,omitempty"`
	ResolvedTs   github_com_pingcap_tiflow_cdc_model.Ts `protobuf:"varint,2,opt,name=resolved_ts,json=resolvedTs,proto3,casttype=github.com/pingcap/tiflow/cdc/model.Ts" json:"resolved_ts,omitempty"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae83c9c6cf5ef75c, []int{0}
}
func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return m.Size()
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetCheckpointTs() github_com_pingcap_tiflow_cdc_model.Ts {
	if m != nil {
		return m.CheckpointTs
	}
	return 0
}

func (m *Checkpoint) GetResolvedTs() github_com_pingcap_tiflow_cdc_model.Ts {
	if m != nil {
		return m.ResolvedTs
	}
	return 0
}

// Stats holds a statistic for a table.
type Stats struct {
	// Number of captured regions.
	RegionCount uint64 `protobuf:"varint,1,opt,name=region_count,json=regionCount,proto3" json:"region_count,omitempty"`
	// The current timestamp from the table's point of view.
	CurrentTs github_com_pingcap_tiflow_cdc_model.Ts `protobuf:"varint,2,opt,name=current_ts,json=currentTs,proto3,casttype=github.com/pingcap/tiflow/cdc/model.Ts" json:"current_ts,omitempty"`
	// Checkponits at each stage.
	StageCheckpoints map[string]Checkpoint `protobuf:"bytes,3,rep,name=stage_checkpoints,json=stageCheckpoints,proto3" json:"stage_checkpoints" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae83c9c6cf5ef75c, []int{1}
}
func (m *Stats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return m.Size()
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetRegionCount() uint64 {
	if m != nil {
		return m.RegionCount
	}
	return 0
}

func (m *Stats) GetCurrentTs() github_com_pingcap_tiflow_cdc_model.Ts {
	if m != nil {
		return m.CurrentTs
	}
	return 0
}

func (m *Stats) GetStageCheckpoints() map[string]Checkpoint {
	if m != nil {
		return m.StageCheckpoints
	}
	return nil
}

// TableStatus is the running status of a table.
type TableStatus struct {
	TableID    github_com_pingcap_tiflow_cdc_model.TableID `protobuf:"varint,1,opt,name=table_id,json=tableId,proto3,casttype=github.com/pingcap/tiflow/cdc/model.TableID" json:"table_id,omitempty"`
	State      TableState                                  `protobuf:"varint,2,opt,name=state,proto3,enum=pingcap.tiflow.cdc.processor.tablepb.TableState" json:"state,omitempty"`
	Checkpoint Checkpoint                                  `protobuf:"bytes,3,opt,name=checkpoint,proto3" json:"checkpoint"`
	Stats      Stats                                       `protobuf:"bytes,4,opt,name=stats,proto3" json:"stats"`
}

func (m *TableStatus) Reset()         { *m = TableStatus{} }
func (m *TableStatus) String() string { return proto.CompactTextString(m) }
func (*TableStatus) ProtoMessage()    {}
func (*TableStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae83c9c6cf5ef75c, []int{2}
}
func (m *TableStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TableStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TableStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TableStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableStatus.Merge(m, src)
}
func (m *TableStatus) XXX_Size() int {
	return m.Size()
}
func (m *TableStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TableStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TableStatus proto.InternalMessageInfo

func (m *TableStatus) GetTableID() github_com_pingcap_tiflow_cdc_model.TableID {
	if m != nil {
		return m.TableID
	}
	return 0
}

func (m *TableStatus) GetState() TableState {
	if m != nil {
		return m.State
	}
	return TableStateUnknown
}

func (m *TableStatus) GetCheckpoint() Checkpoint {
	if m != nil {
		return m.Checkpoint
	}
	return Checkpoint{}
}

func (m *TableStatus) GetStats() Stats {
	if m != nil {
		return m.Stats
	}
	return Stats{}
}

func init() {
	proto.RegisterEnum("pingcap.tiflow.cdc.processor.tablepb.TableState", TableState_name, TableState_value)
	proto.RegisterType((*Checkpoint)(nil), "pingcap.tiflow.cdc.processor.tablepb.Checkpoint")
	proto.RegisterType((*Stats)(nil), "pingcap.tiflow.cdc.processor.tablepb.Stats")
	proto.RegisterMapType((map[string]Checkpoint)(nil), "pingcap.tiflow.cdc.processor.tablepb.Stats.StageCheckpointsEntry")
	proto.RegisterType((*TableStatus)(nil), "pingcap.tiflow.cdc.processor.tablepb.TableStatus")
}

func init() { proto.RegisterFile("processor/tablepb/table.proto", fileDescriptor_ae83c9c6cf5ef75c) }

var fileDescriptor_ae83c9c6cf5ef75c = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0xed, 0xf4, 0x75, 0x5d, 0x90, 0x3b, 0xb4, 0x10, 0x2c, 0xe1, 0x98, 0xa8, 0xaa, 0xaa,
	0x56, 0xb2, 0x51, 0xd9, 0xa0, 0xee, 0x9a, 0xf2, 0x50, 0x85, 0x10, 0xc8, 0x4d, 0x59, 0xb0, 0x89,
	0x9c, 0xf1, 0xe0, 0x5a, 0x75, 0x3d, 0x96, 0x67, 0xdc, 0xaa, 0xbf, 0x90, 0x15, 0x2b, 0x76, 0xf9,
	0x01, 0xbe, 0xa4, 0x1b, 0xa4, 0x2e, 0x59, 0x45, 0x90, 0xfe, 0x45, 0x57, 0x68, 0x3c, 0x6e, 0x5c,
	0x05, 0x84, 0x12, 0x36, 0xc9, 0xcc, 0x9c, 0x7b, 0xce, 0x3d, 0xe7, 0x8e, 0x35, 0xf0, 0x24, 0xcd,
	0x28, 0x26, 0x8c, 0xd1, 0xcc, 0xe5, 0x7e, 0x2f, 0x26, 0x69, 0x4f, 0xfe, 0x3b, 0x69, 0x46, 0x39,
	0x45, 0xeb, 0x69, 0x94, 0x84, 0xd8, 0x4f, 0x1d, 0x1e, 0x7d, 0x8e, 0xe9, 0xb9, 0x83, 0x03, 0xec,
	0x8c, 0x19, 0x4e, 0xc9, 0x30, 0x57, 0x43, 0x1a, 0xd2, 0x82, 0xe0, 0x8a, 0x95, 0xe4, 0xb6, 0xbe,
	0x29, 0x00, 0xfb, 0xc7, 0x04, 0x9f, 0xa4, 0x34, 0x4a, 0x38, 0x7a, 0x0f, 0xf7, 0xf0, 0x78, 0xd7,
	0xe5, 0xac, 0xa1, 0xd8, 0xca, 0x66, 0xbd, 0xbd, 0x75, 0x33, 0x6c, 0x6e, 0x84, 0x11, 0x3f, 0xce,
	0x7b, 0x0e, 0xa6, 0xa7, 0x6e, 0xd9, 0xd0, 0x95, 0x0d, 0x5d, 0x1c, 0x60, 0xf7, 0x94, 0x06, 0x24,
	0x76, 0x3a, 0xcc, 0x5b, 0xae, 0x04, 0x3a, 0x0c, 0xbd, 0x05, 0x3d, 0x23, 0x8c, 0xc6, 0x67, 0x24,
	0x10, 0x72, 0xea, 0xcc, 0x72, 0x70, 0x4b, 0xef, 0xb0, 0xd6, 0x48, 0x85, 0xb9, 0x43, 0xee, 0x73,
	0x86, 0x9e, 0xc2, 0x72, 0x46, 0xc2, 0x88, 0x26, 0x5d, 0x4c, 0xf3, 0x84, 0x4b, 0x9b, 0x9e, 0x2e,
	0xcf, 0xf6, 0xc5, 0x11, 0x3a, 0x00, 0xc0, 0x79, 0x96, 0x11, 0x99, 0x63, 0xf6, 0xc6, 0x4b, 0x25,
	0xbb, 0xc3, 0x10, 0x87, 0x15, 0xc6, 0xfd, 0x90, 0x74, 0xab, 0x68, 0xac, 0xa1, 0xd9, 0xda, 0xa6,
	0xbe, 0xb3, 0xe7, 0x4c, 0x33, 0x7c, 0xa7, 0x70, 0x2d, 0x7e, 0x43, 0x52, 0x4d, 0x9b, 0xbd, 0x4a,
	0x78, 0x76, 0xd1, 0xae, 0x5f, 0x0e, 0x9b, 0x35, 0xcf, 0x60, 0x13, 0xa0, 0x99, 0xc3, 0xda, 0x5f,
	0x09, 0xc8, 0x00, 0xed, 0x84, 0x5c, 0x14, 0x99, 0x97, 0x3c, 0xb1, 0x44, 0xaf, 0x61, 0xee, 0xcc,
	0x8f, 0x73, 0x52, 0xc4, 0xd4, 0x77, 0x9e, 0x4d, 0x67, 0xaa, 0x12, 0xf6, 0x24, 0x7d, 0x57, 0x7d,
	0xa1, 0xb4, 0xbe, 0xab, 0xa0, 0x77, 0x44, 0x85, 0xf0, 0x9c, 0x33, 0x74, 0x04, 0x8b, 0x05, 0xa1,
	0x1b, 0x05, 0x45, 0x4b, 0xad, 0xbd, 0x3b, 0x1a, 0x36, 0x17, 0x8a, 0x92, 0x83, 0x97, 0x37, 0xc3,
	0xe6, 0xf6, 0x54, 0x03, 0x95, 0xe5, 0xde, 0x42, 0xa1, 0x75, 0x10, 0x08, 0xcb, 0x8c, 0xfb, 0x5c,
	0x5a, 0xbe, 0x3f, 0xad, 0xe5, 0xb1, 0x31, 0xe2, 0x49, 0x3a, 0xfa, 0x08, 0x50, 0xdd, 0x4a, 0x43,
	0xfb, 0xbf, 0xfc, 0xe5, 0x1d, 0xdc, 0x51, 0x42, 0x6f, 0xa4, 0x3f, 0xd6, 0xa8, 0x17, 0x92, 0xdb,
	0x33, 0xdc, 0x73, 0xa9, 0x26, 0xf9, 0x5b, 0x5f, 0x55, 0x80, 0xca, 0x36, 0x6a, 0xc1, 0xc2, 0x51,
	0x72, 0x92, 0xd0, 0xf3, 0xc4, 0xa8, 0x99, 0x6b, 0xfd, 0x81, 0xbd, 0x52, 0x81, 0x25, 0x80, 0x6c,
	0x98, 0xdf, 0xeb, 0x31, 0x92, 0x70, 0x43, 0x31, 0x57, 0xfb, 0x03, 0xdb, 0xa8, 0x4a, 0xe4, 0x39,
	0xda, 0x80, 0xa5, 0x0f, 0x19, 0x49, 0xfd, 0x2c, 0x4a, 0x42, 0x43, 0x35, 0x1f, 0xf5, 0x07, 0xf6,
	0x83, 0xaa, 0x68, 0x0c, 0xa1, 0x75, 0x58, 0x94, 0x1b, 0x12, 0x18, 0x9a, 0xf9, 0xb0, 0x3f, 0xb0,
	0xd1, 0x64, 0x19, 0x09, 0xd0, 0x16, 0xe8, 0x1e, 0x49, 0xe3, 0x08, 0xfb, 0x5c, 0xe8, 0xd5, 0xcd,
	0xc7, 0xfd, 0x81, 0xbd, 0x76, 0x67, 0xd6, 0x15, 0x28, 0x14, 0x0f, 0x39, 0x4d, 0xc5, 0x34, 0x8c,
	0xb9, 0x49, 0xc5, 0x5b, 0x44, 0xa4, 0x2c, 0xd6, 0x24, 0x30, 0xe6, 0x27, 0x53, 0x96, 0x40, 0xfb,
	0xdd, 0xd5, 0x2f, 0xab, 0x76, 0x39, 0xb2, 0x94, 0xab, 0x91, 0xa5, 0xfc, 0x1c, 0x59, 0xca, 0x97,
	0x6b, 0xab, 0x76, 0x75, 0x6d, 0xd5, 0x7e, 0x5c, 0x5b, 0xb5, 0x4f, 0xee, 0xbf, 0xbf, 0xaa, 0x3f,
	0x5e, 0xc4, 0xde, 0x7c, 0xf1, 0xa0, 0x3d, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x50, 0x46, 0x7c,
	0xc9, 0x2d, 0x05, 0x00, 0x00,
}

func (m *Checkpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Checkpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Checkpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ResolvedTs != 0 {
		i = encodeVarintTable(dAtA, i, uint64(m.ResolvedTs))
		i--
		dAtA[i] = 0x10
	}
	if m.CheckpointTs != 0 {
		i = encodeVarintTable(dAtA, i, uint64(m.CheckpointTs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Stats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StageCheckpoints) > 0 {
		for k := range m.StageCheckpoints {
			v := m.StageCheckpoints[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTable(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTable(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTable(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.CurrentTs != 0 {
		i = encodeVarintTable(dAtA, i, uint64(m.CurrentTs))
		i--
		dAtA[i] = 0x10
	}
	if m.RegionCount != 0 {
		i = encodeVarintTable(dAtA, i, uint64(m.RegionCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TableStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TableStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTable(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Checkpoint.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTable(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.State != 0 {
		i = encodeVarintTable(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	if m.TableID != 0 {
		i = encodeVarintTable(dAtA, i, uint64(m.TableID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTable(dAtA []byte, offset int, v uint64) int {
	offset -= sovTable(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Checkpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CheckpointTs != 0 {
		n += 1 + sovTable(uint64(m.CheckpointTs))
	}
	if m.ResolvedTs != 0 {
		n += 1 + sovTable(uint64(m.ResolvedTs))
	}
	return n
}

func (m *Stats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RegionCount != 0 {
		n += 1 + sovTable(uint64(m.RegionCount))
	}
	if m.CurrentTs != 0 {
		n += 1 + sovTable(uint64(m.CurrentTs))
	}
	if len(m.StageCheckpoints) > 0 {
		for k, v := range m.StageCheckpoints {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovTable(uint64(len(k))) + 1 + l + sovTable(uint64(l))
			n += mapEntrySize + 1 + sovTable(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *TableStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TableID != 0 {
		n += 1 + sovTable(uint64(m.TableID))
	}
	if m.State != 0 {
		n += 1 + sovTable(uint64(m.State))
	}
	l = m.Checkpoint.Size()
	n += 1 + l + sovTable(uint64(l))
	l = m.Stats.Size()
	n += 1 + l + sovTable(uint64(l))
	return n
}

func sovTable(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTable(x uint64) (n int) {
	return sovTable(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Checkpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTable
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Checkpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Checkpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckpointTs", wireType)
			}
			m.CheckpointTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CheckpointTs |= github_com_pingcap_tiflow_cdc_model.Ts(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolvedTs", wireType)
			}
			m.ResolvedTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolvedTs |= github_com_pingcap_tiflow_cdc_model.Ts(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTable(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTable
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Stats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTable
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Stats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionCount", wireType)
			}
			m.RegionCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegionCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentTs", wireType)
			}
			m.CurrentTs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentTs |= github_com_pingcap_tiflow_cdc_model.Ts(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StageCheckpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTable
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StageCheckpoints == nil {
				m.StageCheckpoints = make(map[string]Checkpoint)
			}
			var mapkey string
			mapvalue := &Checkpoint{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTable
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTable
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTable
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTable
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTable
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTable
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTable
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Checkpoint{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTable(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTable
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.StageCheckpoints[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTable(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTable
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TableStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTable
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TableStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TableID", wireType)
			}
			m.TableID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TableID |= github_com_pingcap_tiflow_cdc_model.TableID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= TableState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Checkpoint", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTable
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Checkpoint.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTable
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTable
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTable
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTable(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTable
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTable(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTable
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTable
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTable
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTable
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTable
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTable
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTable        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTable          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTable = fmt.Errorf("proto: unexpected end of group")
)
