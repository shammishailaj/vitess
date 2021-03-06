// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vschema.proto

package vschema // import "vitess.io/vitess/go/vt/proto/vschema"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import query "vitess.io/vitess/go/vt/proto/query"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Keyspace is the vschema for a keyspace.
type Keyspace struct {
	// If sharded is false, vindexes and tables are ignored.
	Sharded              bool               `protobuf:"varint,1,opt,name=sharded" json:"sharded,omitempty"`
	Vindexes             map[string]*Vindex `protobuf:"bytes,2,rep,name=vindexes" json:"vindexes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tables               map[string]*Table  `protobuf:"bytes,3,rep,name=tables" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Keyspace) Reset()         { *m = Keyspace{} }
func (m *Keyspace) String() string { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()    {}
func (*Keyspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_af169a2691d671e0, []int{0}
}
func (m *Keyspace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keyspace.Unmarshal(m, b)
}
func (m *Keyspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keyspace.Marshal(b, m, deterministic)
}
func (dst *Keyspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyspace.Merge(dst, src)
}
func (m *Keyspace) XXX_Size() int {
	return xxx_messageInfo_Keyspace.Size(m)
}
func (m *Keyspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyspace.DiscardUnknown(m)
}

var xxx_messageInfo_Keyspace proto.InternalMessageInfo

func (m *Keyspace) GetSharded() bool {
	if m != nil {
		return m.Sharded
	}
	return false
}

func (m *Keyspace) GetVindexes() map[string]*Vindex {
	if m != nil {
		return m.Vindexes
	}
	return nil
}

func (m *Keyspace) GetTables() map[string]*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

// Vindex is the vindex info for a Keyspace.
type Vindex struct {
	// The type must match one of the predefined
	// (or plugged in) vindex names.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// params is a map of attribute value pairs
	// that must be defined as required by the
	// vindex constructors. The values can only
	// be strings.
	Params map[string]string `protobuf:"bytes,2,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// A lookup vindex can have an owner table defined.
	// If so, rows in the lookup table are created or
	// deleted in sync with corresponding rows in the
	// owner table.
	Owner                string   `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vindex) Reset()         { *m = Vindex{} }
func (m *Vindex) String() string { return proto.CompactTextString(m) }
func (*Vindex) ProtoMessage()    {}
func (*Vindex) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_af169a2691d671e0, []int{1}
}
func (m *Vindex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vindex.Unmarshal(m, b)
}
func (m *Vindex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vindex.Marshal(b, m, deterministic)
}
func (dst *Vindex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vindex.Merge(dst, src)
}
func (m *Vindex) XXX_Size() int {
	return xxx_messageInfo_Vindex.Size(m)
}
func (m *Vindex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vindex.DiscardUnknown(m)
}

var xxx_messageInfo_Vindex proto.InternalMessageInfo

func (m *Vindex) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Vindex) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Vindex) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

// Table is the table info for a Keyspace.
type Table struct {
	// If the table is a sequence, type must be
	// "sequence". Otherwise, it should be empty.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// column_vindexes associates columns to vindexes.
	ColumnVindexes []*ColumnVindex `protobuf:"bytes,2,rep,name=column_vindexes,json=columnVindexes" json:"column_vindexes,omitempty"`
	// auto_increment is specified if a column needs
	// to be associated with a sequence.
	AutoIncrement *AutoIncrement `protobuf:"bytes,3,opt,name=auto_increment,json=autoIncrement" json:"auto_increment,omitempty"`
	// columns lists the columns for the table.
	Columns              []*Column `protobuf:"bytes,4,rep,name=columns" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_af169a2691d671e0, []int{2}
}
func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (dst *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(dst, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Table) GetColumnVindexes() []*ColumnVindex {
	if m != nil {
		return m.ColumnVindexes
	}
	return nil
}

func (m *Table) GetAutoIncrement() *AutoIncrement {
	if m != nil {
		return m.AutoIncrement
	}
	return nil
}

func (m *Table) GetColumns() []*Column {
	if m != nil {
		return m.Columns
	}
	return nil
}

// ColumnVindex is used to associate a column to a vindex.
type ColumnVindex struct {
	// Legacy implemenation, moving forward all vindexes should define a list of columns.
	Column string `protobuf:"bytes,1,opt,name=column" json:"column,omitempty"`
	// The name must match a vindex defined in Keyspace.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// List of columns that define this Vindex
	Columns              []string `protobuf:"bytes,3,rep,name=columns" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ColumnVindex) Reset()         { *m = ColumnVindex{} }
func (m *ColumnVindex) String() string { return proto.CompactTextString(m) }
func (*ColumnVindex) ProtoMessage()    {}
func (*ColumnVindex) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_af169a2691d671e0, []int{3}
}
func (m *ColumnVindex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ColumnVindex.Unmarshal(m, b)
}
func (m *ColumnVindex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ColumnVindex.Marshal(b, m, deterministic)
}
func (dst *ColumnVindex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ColumnVindex.Merge(dst, src)
}
func (m *ColumnVindex) XXX_Size() int {
	return xxx_messageInfo_ColumnVindex.Size(m)
}
func (m *ColumnVindex) XXX_DiscardUnknown() {
	xxx_messageInfo_ColumnVindex.DiscardUnknown(m)
}

var xxx_messageInfo_ColumnVindex proto.InternalMessageInfo

func (m *ColumnVindex) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *ColumnVindex) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ColumnVindex) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

// Autoincrement is used to designate a column as auto-inc.
type AutoIncrement struct {
	Column string `protobuf:"bytes,1,opt,name=column" json:"column,omitempty"`
	// The sequence must match a table of type SEQUENCE.
	Sequence             string   `protobuf:"bytes,2,opt,name=sequence" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoIncrement) Reset()         { *m = AutoIncrement{} }
func (m *AutoIncrement) String() string { return proto.CompactTextString(m) }
func (*AutoIncrement) ProtoMessage()    {}
func (*AutoIncrement) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_af169a2691d671e0, []int{4}
}
func (m *AutoIncrement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoIncrement.Unmarshal(m, b)
}
func (m *AutoIncrement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoIncrement.Marshal(b, m, deterministic)
}
func (dst *AutoIncrement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoIncrement.Merge(dst, src)
}
func (m *AutoIncrement) XXX_Size() int {
	return xxx_messageInfo_AutoIncrement.Size(m)
}
func (m *AutoIncrement) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoIncrement.DiscardUnknown(m)
}

var xxx_messageInfo_AutoIncrement proto.InternalMessageInfo

func (m *AutoIncrement) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *AutoIncrement) GetSequence() string {
	if m != nil {
		return m.Sequence
	}
	return ""
}

// Column describes a column.
type Column struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type                 query.Type `protobuf:"varint,2,opt,name=type,enum=query.Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Column) Reset()         { *m = Column{} }
func (m *Column) String() string { return proto.CompactTextString(m) }
func (*Column) ProtoMessage()    {}
func (*Column) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_af169a2691d671e0, []int{5}
}
func (m *Column) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Column.Unmarshal(m, b)
}
func (m *Column) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Column.Marshal(b, m, deterministic)
}
func (dst *Column) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Column.Merge(dst, src)
}
func (m *Column) XXX_Size() int {
	return xxx_messageInfo_Column.Size(m)
}
func (m *Column) XXX_DiscardUnknown() {
	xxx_messageInfo_Column.DiscardUnknown(m)
}

var xxx_messageInfo_Column proto.InternalMessageInfo

func (m *Column) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Column) GetType() query.Type {
	if m != nil {
		return m.Type
	}
	return query.Type_NULL_TYPE
}

// SrvVSchema is the roll-up of all the Keyspace schema for a cell.
type SrvVSchema struct {
	// keyspaces is a map of keyspace name -> Keyspace object.
	Keyspaces            map[string]*Keyspace `protobuf:"bytes,1,rep,name=keyspaces" json:"keyspaces,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SrvVSchema) Reset()         { *m = SrvVSchema{} }
func (m *SrvVSchema) String() string { return proto.CompactTextString(m) }
func (*SrvVSchema) ProtoMessage()    {}
func (*SrvVSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_vschema_af169a2691d671e0, []int{6}
}
func (m *SrvVSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrvVSchema.Unmarshal(m, b)
}
func (m *SrvVSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrvVSchema.Marshal(b, m, deterministic)
}
func (dst *SrvVSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrvVSchema.Merge(dst, src)
}
func (m *SrvVSchema) XXX_Size() int {
	return xxx_messageInfo_SrvVSchema.Size(m)
}
func (m *SrvVSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_SrvVSchema.DiscardUnknown(m)
}

var xxx_messageInfo_SrvVSchema proto.InternalMessageInfo

func (m *SrvVSchema) GetKeyspaces() map[string]*Keyspace {
	if m != nil {
		return m.Keyspaces
	}
	return nil
}

func init() {
	proto.RegisterType((*Keyspace)(nil), "vschema.Keyspace")
	proto.RegisterMapType((map[string]*Table)(nil), "vschema.Keyspace.TablesEntry")
	proto.RegisterMapType((map[string]*Vindex)(nil), "vschema.Keyspace.VindexesEntry")
	proto.RegisterType((*Vindex)(nil), "vschema.Vindex")
	proto.RegisterMapType((map[string]string)(nil), "vschema.Vindex.ParamsEntry")
	proto.RegisterType((*Table)(nil), "vschema.Table")
	proto.RegisterType((*ColumnVindex)(nil), "vschema.ColumnVindex")
	proto.RegisterType((*AutoIncrement)(nil), "vschema.AutoIncrement")
	proto.RegisterType((*Column)(nil), "vschema.Column")
	proto.RegisterType((*SrvVSchema)(nil), "vschema.SrvVSchema")
	proto.RegisterMapType((map[string]*Keyspace)(nil), "vschema.SrvVSchema.KeyspacesEntry")
}

func init() { proto.RegisterFile("vschema.proto", fileDescriptor_vschema_af169a2691d671e0) }

var fileDescriptor_vschema_af169a2691d671e0 = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0xed, 0x36, 0x6d, 0x6f, 0x6c, 0xaa, 0xc3, 0xba, 0x84, 0x88, 0x6c, 0x09, 0xab, 0xd6,
	0x97, 0x04, 0xba, 0x08, 0x7e, 0xb0, 0xa2, 0x16, 0x1f, 0x16, 0x05, 0x25, 0x5b, 0xf6, 0xc1, 0x97,
	0x65, 0x36, 0xbd, 0xb8, 0x65, 0x9b, 0x8f, 0xcd, 0x24, 0xd1, 0xfc, 0x1a, 0xc1, 0x7f, 0xe0, 0x8f,
	0xf0, 0x7f, 0x49, 0x66, 0x26, 0xe9, 0xa4, 0xad, 0x6f, 0x73, 0x7a, 0xee, 0x39, 0xf7, 0xcc, 0xcd,
	0x9d, 0xc2, 0xa8, 0x60, 0xc1, 0x0d, 0x86, 0xd4, 0x4d, 0xd2, 0x38, 0x8b, 0x49, 0x5f, 0x42, 0xdb,
	0xb8, 0xcb, 0x31, 0x2d, 0xc5, 0xaf, 0xce, 0x9f, 0x0e, 0x0c, 0x3e, 0x61, 0xc9, 0x12, 0x1a, 0x20,
	0xb1, 0xa0, 0xcf, 0x6e, 0x68, 0xba, 0xc4, 0xa5, 0xa5, 0x4d, 0xb4, 0xe9, 0xc0, 0xaf, 0x21, 0x79,
	0x03, 0x83, 0x62, 0x15, 0x2d, 0xf1, 0x27, 0x32, 0xab, 0x33, 0xe9, 0x4e, 0x8d, 0xd9, 0xb1, 0x5b,
	0xdb, 0xd7, 0x72, 0xf7, 0x52, 0x56, 0x7c, 0x8c, 0xb2, 0xb4, 0xf4, 0x1b, 0x01, 0x79, 0x01, 0x7a,
	0x46, 0xaf, 0xd7, 0xc8, 0xac, 0x2e, 0x97, 0x3e, 0xde, 0x95, 0x2e, 0x38, 0x2f, 0x84, 0xb2, 0xd8,
	0xfe, 0x0c, 0xa3, 0x96, 0x23, 0xb9, 0x0f, 0xdd, 0x5b, 0x2c, 0x79, 0xb4, 0xa1, 0x5f, 0x1d, 0xc9,
	0x13, 0xe8, 0x15, 0x74, 0x9d, 0xa3, 0xd5, 0x99, 0x68, 0x53, 0x63, 0x36, 0x6e, 0x8c, 0x85, 0xd0,
	0x17, 0xec, 0xeb, 0xce, 0x4b, 0xcd, 0x3e, 0x07, 0x43, 0x69, 0xb2, 0xc7, 0xeb, 0xa4, 0xed, 0x65,
	0x36, 0x5e, 0x5c, 0xa6, 0x58, 0x39, 0xbf, 0x35, 0xd0, 0x45, 0x03, 0x42, 0xe0, 0x20, 0x2b, 0x13,
	0x94, 0x3e, 0xfc, 0x4c, 0x4e, 0x41, 0x4f, 0x68, 0x4a, 0xc3, 0x7a, 0x52, 0x8f, 0xb6, 0x52, 0xb9,
	0x5f, 0x39, 0x2b, 0x2f, 0x2b, 0x4a, 0xc9, 0x21, 0xf4, 0xe2, 0x1f, 0x11, 0xa6, 0x56, 0x97, 0x3b,
	0x09, 0x60, 0xbf, 0x02, 0x43, 0x29, 0xde, 0x13, 0xfa, 0x50, 0x0d, 0x3d, 0x54, 0x43, 0xfe, 0xd5,
	0xa0, 0xc7, 0x93, 0xef, 0xcd, 0xf8, 0x16, 0xc6, 0x41, 0xbc, 0xce, 0xc3, 0xe8, 0x6a, 0xeb, 0xb3,
	0x3e, 0x6c, 0xc2, 0xce, 0x39, 0x2f, 0x07, 0x69, 0x06, 0x0a, 0x42, 0x46, 0xce, 0xc0, 0xa4, 0x79,
	0x16, 0x5f, 0xad, 0xa2, 0x20, 0xc5, 0x10, 0xa3, 0x8c, 0xe7, 0x36, 0x66, 0x47, 0x8d, 0xfc, 0x7d,
	0x9e, 0xc5, 0xe7, 0x35, 0xeb, 0x8f, 0xa8, 0x0a, 0xc9, 0x73, 0xe8, 0x0b, 0x43, 0x66, 0x1d, 0xf0,
	0xb6, 0xe3, 0xad, 0xb6, 0x7e, 0xcd, 0x3b, 0x0b, 0xb8, 0xa7, 0x26, 0x21, 0x47, 0xa0, 0x0b, 0x4a,
	0xde, 0x47, 0xa2, 0xea, 0x96, 0x11, 0x0d, 0xeb, 0x41, 0xf0, 0x73, 0xb5, 0xcf, 0x75, 0x9b, 0x6a,
	0xf3, 0x86, 0x1b, 0xd7, 0x39, 0x8c, 0x5a, 0x01, 0xff, 0x6b, 0x6b, 0xc3, 0x80, 0xe1, 0x5d, 0x8e,
	0x51, 0x50, 0x5b, 0x37, 0xd8, 0x39, 0x03, 0x7d, 0xde, 0x6e, 0xae, 0x29, 0xcd, 0x8f, 0xe5, 0xd8,
	0x2b, 0x95, 0x39, 0x33, 0x5c, 0xf1, 0xea, 0x16, 0x65, 0x82, 0xe2, 0x1b, 0x38, 0xbf, 0x34, 0x80,
	0x8b, 0xb4, 0xb8, 0xbc, 0xe0, 0x17, 0x27, 0xef, 0x60, 0x78, 0x2b, 0x9f, 0x03, 0xb3, 0x34, 0x3e,
	0x15, 0xa7, 0x99, 0xca, 0xa6, 0xae, 0x79, 0x33, 0x72, 0x81, 0x36, 0x22, 0xfb, 0x0b, 0x98, 0x6d,
	0x72, 0xcf, 0xc2, 0x3c, 0x6b, 0x6f, 0xf9, 0x83, 0x9d, 0xa7, 0xa8, 0xec, 0xd0, 0x87, 0xa7, 0xdf,
	0x4e, 0x8a, 0x55, 0x86, 0x8c, 0xb9, 0xab, 0xd8, 0x13, 0x27, 0xef, 0x7b, 0xec, 0x15, 0x99, 0xc7,
	0xff, 0x3c, 0x3c, 0xa9, 0xbd, 0xd6, 0x39, 0x3c, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x29,
	0xa2, 0x18, 0x72, 0x04, 0x00, 0x00,
}
