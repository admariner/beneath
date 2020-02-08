// Code generated by protoc-gen-go. DO NOT EDIT.
// source: engine.proto

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

type Record struct {
	AvroData             []byte   `protobuf:"bytes,1,opt,name=avro_data,json=avroData,proto3" json:"avro_data,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_770b178c3aab763f, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetAvroData() []byte {
	if m != nil {
		return m.AvroData
	}
	return nil
}

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type WriteReport struct {
	InstanceId           []byte   `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	RecordsCount         int32    `protobuf:"varint,2,opt,name=records_count,json=recordsCount,proto3" json:"records_count,omitempty"`
	BytesTotal           int32    `protobuf:"varint,3,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteReport) Reset()         { *m = WriteReport{} }
func (m *WriteReport) String() string { return proto.CompactTextString(m) }
func (*WriteReport) ProtoMessage()    {}
func (*WriteReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_770b178c3aab763f, []int{1}
}

func (m *WriteReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteReport.Unmarshal(m, b)
}
func (m *WriteReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteReport.Marshal(b, m, deterministic)
}
func (m *WriteReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteReport.Merge(m, src)
}
func (m *WriteReport) XXX_Size() int {
	return xxx_messageInfo_WriteReport.Size(m)
}
func (m *WriteReport) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteReport.DiscardUnknown(m)
}

var xxx_messageInfo_WriteReport proto.InternalMessageInfo

func (m *WriteReport) GetInstanceId() []byte {
	if m != nil {
		return m.InstanceId
	}
	return nil
}

func (m *WriteReport) GetRecordsCount() int32 {
	if m != nil {
		return m.RecordsCount
	}
	return 0
}

func (m *WriteReport) GetBytesTotal() int32 {
	if m != nil {
		return m.BytesTotal
	}
	return 0
}

func init() {
	proto.RegisterType((*Record)(nil), "proto.Record")
	proto.RegisterType((*WriteReport)(nil), "proto.WriteReport")
}

func init() { proto.RegisterFile("engine.proto", fileDescriptor_770b178c3aab763f) }

var fileDescriptor_770b178c3aab763f = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xcf, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0x62, 0x69, 0xb1, 0xd3, 0x78, 0x09, 0x08, 0x0b, 0x0a, 0x96, 0x7a, 0xb0, 0xa7,
	0x5e, 0x7c, 0x83, 0xad, 0x17, 0x6f, 0x25, 0x08, 0x1e, 0xc3, 0xec, 0x66, 0xd0, 0xa0, 0x9b, 0x59,
	0x92, 0x51, 0xf1, 0xed, 0x25, 0xd9, 0x95, 0x9e, 0x86, 0xff, 0x63, 0xf8, 0x87, 0x01, 0x4d, 0xf1,
	0x2d, 0x44, 0x3a, 0x8c, 0x89, 0x85, 0xcd, 0xb2, 0x8e, 0xdd, 0x11, 0x56, 0x96, 0x7a, 0x4e, 0xde,
	0xdc, 0xc0, 0x1a, 0xbf, 0x13, 0x3b, 0x8f, 0x82, 0x8d, 0xda, 0xaa, 0xbd, 0xb6, 0x97, 0x05, 0x9e,
	0x50, 0xd0, 0xdc, 0xc2, 0x5a, 0xc2, 0x40, 0x59, 0x70, 0x18, 0x9b, 0x8b, 0xad, 0xda, 0x2f, 0xec,
	0x19, 0x76, 0x02, 0x9b, 0xd7, 0x14, 0x84, 0x2c, 0x8d, 0x9c, 0xc4, 0xdc, 0xc1, 0x26, 0xc4, 0x2c,
	0x18, 0x7b, 0x72, 0xc1, 0xcf, 0x5d, 0xf0, 0x4f, 0xcf, 0xde, 0xdc, 0xc3, 0x55, 0xaa, 0x47, 0xb3,
	0xeb, 0xf9, 0x2b, 0x4a, 0x6d, 0x5c, 0x5a, 0x3d, 0xe3, 0xb1, 0x58, 0x69, 0xe9, 0x7e, 0x85, 0xb2,
	0x13, 0x16, 0xfc, 0x6c, 0x16, 0x75, 0x05, 0x2a, 0xbd, 0x14, 0x69, 0x1f, 0xe0, 0x3a, 0x92, 0xfc,
	0x70, 0xfa, 0x38, 0x74, 0x14, 0x09, 0xe5, 0x7d, 0x7a, 0xad, 0xd5, 0xed, 0x14, 0x4f, 0x25, 0x9d,
	0x54, 0xb7, 0xaa, 0xfc, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x17, 0x59, 0xd0, 0x01, 0x01,
	0x00, 0x00,
}
