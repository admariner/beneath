// Code generated by protoc-gen-go. DO NOT EDIT.
// source: engine/driver/bigtable/proto/cursor.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Cursor_Type int32

const (
	Cursor_LOG   Cursor_Type = 0
	Cursor_INDEX Cursor_Type = 1
	Cursor_PEEK  Cursor_Type = 2
)

var Cursor_Type_name = map[int32]string{
	0: "LOG",
	1: "INDEX",
	2: "PEEK",
}

var Cursor_Type_value = map[string]int32{
	"LOG":   0,
	"INDEX": 1,
	"PEEK":  2,
}

func (x Cursor_Type) String() string {
	return proto.EnumName(Cursor_Type_name, int32(x))
}

func (Cursor_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_09dade45cc45a2f4, []int{1, 0}
}

type CursorSet struct {
	Cursors              []*Cursor `protobuf:"bytes,1,rep,name=cursors,proto3" json:"cursors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CursorSet) Reset()         { *m = CursorSet{} }
func (m *CursorSet) String() string { return proto.CompactTextString(m) }
func (*CursorSet) ProtoMessage()    {}
func (*CursorSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_09dade45cc45a2f4, []int{0}
}

func (m *CursorSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CursorSet.Unmarshal(m, b)
}
func (m *CursorSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CursorSet.Marshal(b, m, deterministic)
}
func (m *CursorSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CursorSet.Merge(m, src)
}
func (m *CursorSet) XXX_Size() int {
	return xxx_messageInfo_CursorSet.Size(m)
}
func (m *CursorSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CursorSet.DiscardUnknown(m)
}

var xxx_messageInfo_CursorSet proto.InternalMessageInfo

func (m *CursorSet) GetCursors() []*Cursor {
	if m != nil {
		return m.Cursors
	}
	return nil
}

type Cursor struct {
	Type                 Cursor_Type `protobuf:"varint,1,opt,name=type,proto3,enum=bigtable.Cursor_Type" json:"type,omitempty"`
	LogStart             int64       `protobuf:"varint,10,opt,name=log_start,json=logStart,proto3" json:"log_start,omitempty"`
	LogEnd               int64       `protobuf:"varint,11,opt,name=log_end,json=logEnd,proto3" json:"log_end,omitempty"`
	IndexId              int32       `protobuf:"varint,20,opt,name=index_id,json=indexId,proto3" json:"index_id,omitempty"`
	IndexStart           []byte      `protobuf:"bytes,21,opt,name=index_start,json=indexStart,proto3" json:"index_start,omitempty"`
	IndexEnd             []byte      `protobuf:"bytes,22,opt,name=index_end,json=indexEnd,proto3" json:"index_end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Cursor) Reset()         { *m = Cursor{} }
func (m *Cursor) String() string { return proto.CompactTextString(m) }
func (*Cursor) ProtoMessage()    {}
func (*Cursor) Descriptor() ([]byte, []int) {
	return fileDescriptor_09dade45cc45a2f4, []int{1}
}

func (m *Cursor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cursor.Unmarshal(m, b)
}
func (m *Cursor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cursor.Marshal(b, m, deterministic)
}
func (m *Cursor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cursor.Merge(m, src)
}
func (m *Cursor) XXX_Size() int {
	return xxx_messageInfo_Cursor.Size(m)
}
func (m *Cursor) XXX_DiscardUnknown() {
	xxx_messageInfo_Cursor.DiscardUnknown(m)
}

var xxx_messageInfo_Cursor proto.InternalMessageInfo

func (m *Cursor) GetType() Cursor_Type {
	if m != nil {
		return m.Type
	}
	return Cursor_LOG
}

func (m *Cursor) GetLogStart() int64 {
	if m != nil {
		return m.LogStart
	}
	return 0
}

func (m *Cursor) GetLogEnd() int64 {
	if m != nil {
		return m.LogEnd
	}
	return 0
}

func (m *Cursor) GetIndexId() int32 {
	if m != nil {
		return m.IndexId
	}
	return 0
}

func (m *Cursor) GetIndexStart() []byte {
	if m != nil {
		return m.IndexStart
	}
	return nil
}

func (m *Cursor) GetIndexEnd() []byte {
	if m != nil {
		return m.IndexEnd
	}
	return nil
}

func init() {
	proto.RegisterEnum("bigtable.Cursor_Type", Cursor_Type_name, Cursor_Type_value)
	proto.RegisterType((*CursorSet)(nil), "bigtable.CursorSet")
	proto.RegisterType((*Cursor)(nil), "bigtable.Cursor")
}

func init() {
	proto.RegisterFile("engine/driver/bigtable/proto/cursor.proto", fileDescriptor_09dade45cc45a2f4)
}

var fileDescriptor_09dade45cc45a2f4 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x8d, 0xeb, 0xdf, 0xb7, 0x22, 0xe5, 0xc5, 0x69, 0x65, 0x07, 0x4b, 0xf1, 0xd0, 0x79,
	0x68, 0x61, 0x1e, 0xbc, 0xab, 0x45, 0x86, 0xa2, 0xd2, 0x79, 0x10, 0x2f, 0x63, 0x35, 0xa1, 0x04,
	0x4a, 0x52, 0xd2, 0x28, 0xee, 0x1b, 0xfb, 0x31, 0xa4, 0x09, 0xbd, 0x78, 0x0a, 0xcf, 0xf3, 0xfb,
	0xe5, 0x21, 0x04, 0x96, 0x4c, 0xb4, 0x5c, 0xb0, 0x92, 0x2a, 0xfe, 0xcd, 0x54, 0xd9, 0xf0, 0x56,
	0xef, 0x9a, 0x8e, 0x95, 0xbd, 0x92, 0x5a, 0x96, 0x9f, 0x5f, 0x6a, 0x90, 0xaa, 0x30, 0x01, 0x83,
	0x09, 0x66, 0x37, 0x10, 0xde, 0x19, 0xb2, 0x61, 0x1a, 0xaf, 0xc0, 0xb7, 0xda, 0x90, 0x90, 0x74,
	0x96, 0x47, 0xab, 0xb8, 0x98, 0xc4, 0xc2, 0x5a, 0xf5, 0x24, 0x64, 0xbf, 0x04, 0x3c, 0xdb, 0xe1,
	0x12, 0x1c, 0xbd, 0xef, 0x59, 0x42, 0x52, 0x92, 0x1f, 0xaf, 0xe6, 0xff, 0xef, 0x14, 0x6f, 0xfb,
	0x9e, 0xd5, 0x46, 0xc1, 0x05, 0x84, 0x9d, 0x6c, 0xb7, 0x83, 0xde, 0x29, 0x9d, 0x40, 0x4a, 0xf2,
	0x59, 0x1d, 0x74, 0xb2, 0xdd, 0x8c, 0x19, 0xcf, 0xc0, 0x1f, 0x21, 0x13, 0x34, 0x89, 0x0c, 0xf2,
	0x3a, 0xd9, 0x56, 0x82, 0xe2, 0x39, 0x04, 0x5c, 0x50, 0xf6, 0xb3, 0xe5, 0x34, 0x39, 0x49, 0x49,
	0xee, 0xd6, 0xbe, 0xc9, 0x6b, 0x8a, 0x17, 0x10, 0x59, 0x64, 0x27, 0xe7, 0x29, 0xc9, 0x8f, 0x6a,
	0x30, 0x95, 0x1d, 0x5d, 0x40, 0x68, 0x85, 0x71, 0xf6, 0xd4, 0x60, 0x3b, 0x56, 0x09, 0x9a, 0x5d,
	0x82, 0x33, 0x3e, 0x0e, 0x7d, 0x98, 0x3d, 0xbd, 0x3c, 0xc4, 0x07, 0x18, 0x82, 0xbb, 0x7e, 0xbe,
	0xaf, 0xde, 0x63, 0x82, 0x01, 0x38, 0xaf, 0x55, 0xf5, 0x18, 0x1f, 0xde, 0xfa, 0x1f, 0xae, 0xf9,
	0xb6, 0xc6, 0x33, 0xc7, 0xf5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x0c, 0x9d, 0x84, 0x6a,
	0x01, 0x00, 0x00,
}
