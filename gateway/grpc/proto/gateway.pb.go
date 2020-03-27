// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway/grpc/proto/gateway.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	return fileDescriptor_1fb7a2419213bb16, []int{0}
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

type PingRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientVersion        string   `protobuf:"bytes,2,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{1}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *PingRequest) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

type PingResponse struct {
	Authenticated        bool     `protobuf:"varint,1,opt,name=authenticated,proto3" json:"authenticated,omitempty"`
	VersionStatus        string   `protobuf:"bytes,2,opt,name=version_status,json=versionStatus,proto3" json:"version_status,omitempty"`
	RecommendedVersion   string   `protobuf:"bytes,3,opt,name=recommended_version,json=recommendedVersion,proto3" json:"recommended_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{2}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetAuthenticated() bool {
	if m != nil {
		return m.Authenticated
	}
	return false
}

func (m *PingResponse) GetVersionStatus() string {
	if m != nil {
		return m.VersionStatus
	}
	return ""
}

func (m *PingResponse) GetRecommendedVersion() string {
	if m != nil {
		return m.RecommendedVersion
	}
	return ""
}

type WriteRequest struct {
	InstanceId           []byte    `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Records              []*Record `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{3}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetInstanceId() []byte {
	if m != nil {
		return m.InstanceId
	}
	return nil
}

func (m *WriteRequest) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type WriteResponse struct {
	WriteId              []byte   `protobuf:"bytes,1,opt,name=write_id,json=writeId,proto3" json:"write_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{4}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

func (m *WriteResponse) GetWriteId() []byte {
	if m != nil {
		return m.WriteId
	}
	return nil
}

type QueryLogRequest struct {
	InstanceId           []byte   `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Partitions           int32    `protobuf:"varint,2,opt,name=partitions,proto3" json:"partitions,omitempty"`
	Peek                 bool     `protobuf:"varint,3,opt,name=peek,proto3" json:"peek,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryLogRequest) Reset()         { *m = QueryLogRequest{} }
func (m *QueryLogRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLogRequest) ProtoMessage()    {}
func (*QueryLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{5}
}

func (m *QueryLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryLogRequest.Unmarshal(m, b)
}
func (m *QueryLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryLogRequest.Marshal(b, m, deterministic)
}
func (m *QueryLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLogRequest.Merge(m, src)
}
func (m *QueryLogRequest) XXX_Size() int {
	return xxx_messageInfo_QueryLogRequest.Size(m)
}
func (m *QueryLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLogRequest proto.InternalMessageInfo

func (m *QueryLogRequest) GetInstanceId() []byte {
	if m != nil {
		return m.InstanceId
	}
	return nil
}

func (m *QueryLogRequest) GetPartitions() int32 {
	if m != nil {
		return m.Partitions
	}
	return 0
}

func (m *QueryLogRequest) GetPeek() bool {
	if m != nil {
		return m.Peek
	}
	return false
}

type QueryLogResponse struct {
	ReplayCursors        [][]byte `protobuf:"bytes,1,rep,name=replay_cursors,json=replayCursors,proto3" json:"replay_cursors,omitempty"`
	ChangeCursors        [][]byte `protobuf:"bytes,2,rep,name=change_cursors,json=changeCursors,proto3" json:"change_cursors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryLogResponse) Reset()         { *m = QueryLogResponse{} }
func (m *QueryLogResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLogResponse) ProtoMessage()    {}
func (*QueryLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{6}
}

func (m *QueryLogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryLogResponse.Unmarshal(m, b)
}
func (m *QueryLogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryLogResponse.Marshal(b, m, deterministic)
}
func (m *QueryLogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLogResponse.Merge(m, src)
}
func (m *QueryLogResponse) XXX_Size() int {
	return xxx_messageInfo_QueryLogResponse.Size(m)
}
func (m *QueryLogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLogResponse proto.InternalMessageInfo

func (m *QueryLogResponse) GetReplayCursors() [][]byte {
	if m != nil {
		return m.ReplayCursors
	}
	return nil
}

func (m *QueryLogResponse) GetChangeCursors() [][]byte {
	if m != nil {
		return m.ChangeCursors
	}
	return nil
}

type QueryIndexRequest struct {
	InstanceId           []byte   `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Partitions           int32    `protobuf:"varint,2,opt,name=partitions,proto3" json:"partitions,omitempty"`
	Filter               string   `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryIndexRequest) Reset()         { *m = QueryIndexRequest{} }
func (m *QueryIndexRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIndexRequest) ProtoMessage()    {}
func (*QueryIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{7}
}

func (m *QueryIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryIndexRequest.Unmarshal(m, b)
}
func (m *QueryIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryIndexRequest.Marshal(b, m, deterministic)
}
func (m *QueryIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIndexRequest.Merge(m, src)
}
func (m *QueryIndexRequest) XXX_Size() int {
	return xxx_messageInfo_QueryIndexRequest.Size(m)
}
func (m *QueryIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIndexRequest proto.InternalMessageInfo

func (m *QueryIndexRequest) GetInstanceId() []byte {
	if m != nil {
		return m.InstanceId
	}
	return nil
}

func (m *QueryIndexRequest) GetPartitions() int32 {
	if m != nil {
		return m.Partitions
	}
	return 0
}

func (m *QueryIndexRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type QueryIndexResponse struct {
	ReplayCursors        [][]byte `protobuf:"bytes,1,rep,name=replay_cursors,json=replayCursors,proto3" json:"replay_cursors,omitempty"`
	ChangeCursors        [][]byte `protobuf:"bytes,2,rep,name=change_cursors,json=changeCursors,proto3" json:"change_cursors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryIndexResponse) Reset()         { *m = QueryIndexResponse{} }
func (m *QueryIndexResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIndexResponse) ProtoMessage()    {}
func (*QueryIndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{8}
}

func (m *QueryIndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryIndexResponse.Unmarshal(m, b)
}
func (m *QueryIndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryIndexResponse.Marshal(b, m, deterministic)
}
func (m *QueryIndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIndexResponse.Merge(m, src)
}
func (m *QueryIndexResponse) XXX_Size() int {
	return xxx_messageInfo_QueryIndexResponse.Size(m)
}
func (m *QueryIndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIndexResponse proto.InternalMessageInfo

func (m *QueryIndexResponse) GetReplayCursors() [][]byte {
	if m != nil {
		return m.ReplayCursors
	}
	return nil
}

func (m *QueryIndexResponse) GetChangeCursors() [][]byte {
	if m != nil {
		return m.ChangeCursors
	}
	return nil
}

type ReadRequest struct {
	InstanceId           []byte   `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Cursor               []byte   `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{9}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetInstanceId() []byte {
	if m != nil {
		return m.InstanceId
	}
	return nil
}

func (m *ReadRequest) GetCursor() []byte {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *ReadRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ReadResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	NextCursor           []byte    `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{10}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *ReadResponse) GetNextCursor() []byte {
	if m != nil {
		return m.NextCursor
	}
	return nil
}

type SubscribeRequest struct {
	InstanceId           []byte   `protobuf:"bytes,1,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Cursor               []byte   `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{11}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetInstanceId() []byte {
	if m != nil {
		return m.InstanceId
	}
	return nil
}

func (m *SubscribeRequest) GetCursor() []byte {
	if m != nil {
		return m.Cursor
	}
	return nil
}

type SubscribeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeResponse) Reset()         { *m = SubscribeResponse{} }
func (m *SubscribeResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeResponse) ProtoMessage()    {}
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fb7a2419213bb16, []int{12}
}

func (m *SubscribeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeResponse.Unmarshal(m, b)
}
func (m *SubscribeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeResponse.Merge(m, src)
}
func (m *SubscribeResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeResponse.Size(m)
}
func (m *SubscribeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Record)(nil), "gateway.v1.Record")
	proto.RegisterType((*PingRequest)(nil), "gateway.v1.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "gateway.v1.PingResponse")
	proto.RegisterType((*WriteRequest)(nil), "gateway.v1.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "gateway.v1.WriteResponse")
	proto.RegisterType((*QueryLogRequest)(nil), "gateway.v1.QueryLogRequest")
	proto.RegisterType((*QueryLogResponse)(nil), "gateway.v1.QueryLogResponse")
	proto.RegisterType((*QueryIndexRequest)(nil), "gateway.v1.QueryIndexRequest")
	proto.RegisterType((*QueryIndexResponse)(nil), "gateway.v1.QueryIndexResponse")
	proto.RegisterType((*ReadRequest)(nil), "gateway.v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "gateway.v1.ReadResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "gateway.v1.SubscribeRequest")
	proto.RegisterType((*SubscribeResponse)(nil), "gateway.v1.SubscribeResponse")
}

func init() { proto.RegisterFile("gateway/grpc/proto/gateway.proto", fileDescriptor_1fb7a2419213bb16) }

var fileDescriptor_1fb7a2419213bb16 = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x6d, 0x4f, 0x13, 0x41,
	0x10, 0xe6, 0x28, 0x2d, 0xed, 0xb4, 0x45, 0x58, 0x0c, 0x96, 0xf2, 0x62, 0x73, 0x91, 0xa4, 0x21,
	0xd8, 0x2a, 0x7e, 0x24, 0xf1, 0x83, 0x98, 0x90, 0x46, 0x4c, 0x64, 0x49, 0x34, 0x31, 0x9a, 0xba,
	0xbd, 0x1b, 0xda, 0x8d, 0xed, 0xde, 0xb9, 0xb7, 0xc7, 0xcb, 0x7f, 0xf0, 0x8f, 0xf8, 0x2f, 0xcd,
	0xed, 0xde, 0xb5, 0x7b, 0x14, 0x22, 0x89, 0xfa, 0xa9, 0x9d, 0x67, 0xe6, 0x9e, 0x7d, 0x9e, 0xd9,
	0x99, 0x3b, 0x68, 0x0d, 0x99, 0xc2, 0x2b, 0x76, 0xd3, 0x1d, 0xca, 0xd0, 0xeb, 0x86, 0x32, 0x50,
	0x41, 0x37, 0x85, 0x3a, 0x3a, 0x22, 0x90, 0x85, 0x97, 0x2f, 0xdd, 0x63, 0x28, 0x51, 0xf4, 0x02,
	0xe9, 0x93, 0x2d, 0xa8, 0xb0, 0x4b, 0x19, 0xf4, 0x7d, 0xa6, 0x58, 0xc3, 0x69, 0x39, 0xed, 0x1a,
	0x2d, 0x27, 0xc0, 0x5b, 0xa6, 0x18, 0xd9, 0x86, 0x8a, 0xe2, 0x13, 0x8c, 0x14, 0x9b, 0x84, 0x8d,
	0xc5, 0x96, 0xd3, 0x2e, 0xd0, 0x19, 0xe0, 0x9e, 0x41, 0xf5, 0x03, 0x17, 0x43, 0x8a, 0x3f, 0x62,
	0x8c, 0x54, 0xc2, 0xe4, 0x8d, 0x39, 0x0a, 0xd5, 0xe7, 0xbe, 0x66, 0xaa, 0xd0, 0xb2, 0x01, 0x7a,
	0x3e, 0xd9, 0x83, 0x95, 0x34, 0x79, 0x89, 0x32, 0xe2, 0x81, 0xd0, 0x74, 0x15, 0x5a, 0x37, 0xe8,
	0x47, 0x03, 0xba, 0x3f, 0x1d, 0xa8, 0x19, 0xce, 0x28, 0x0c, 0x44, 0x84, 0xe4, 0x19, 0xd4, 0x59,
	0xac, 0x46, 0x28, 0x14, 0xf7, 0x98, 0x42, 0x43, 0x5c, 0xa6, 0x79, 0x30, 0x61, 0x4f, 0x69, 0xfb,
	0x91, 0x62, 0x2a, 0x8e, 0x32, 0xf6, 0x14, 0x3d, 0xd7, 0x20, 0xe9, 0xc2, 0xba, 0x44, 0x2f, 0x98,
	0x4c, 0x50, 0xf8, 0xe8, 0x4f, 0x95, 0x14, 0x74, 0x2d, 0xb1, 0x52, 0x99, 0x9c, 0xaf, 0x50, 0xfb,
	0x24, 0xb9, 0xc2, 0xcc, 0xe2, 0x53, 0xa8, 0x72, 0x11, 0x29, 0x26, 0x3c, 0xcc, 0x4c, 0xd6, 0x28,
	0x64, 0x50, 0xcf, 0x27, 0x07, 0xb0, 0x2c, 0x75, 0x5f, 0x13, 0x05, 0x85, 0x76, 0xf5, 0x90, 0x74,
	0x66, 0x5d, 0xef, 0x98, 0x96, 0xd3, 0xac, 0xc4, 0xdd, 0x87, 0x7a, 0x4a, 0x9f, 0xba, 0xdd, 0x84,
	0xf2, 0x55, 0x02, 0xcc, 0xc8, 0x97, 0x75, 0xdc, 0xf3, 0xdd, 0x0b, 0x78, 0x74, 0x16, 0xa3, 0xbc,
	0x39, 0x0d, 0x86, 0x0f, 0x56, 0xb3, 0x0b, 0x10, 0x32, 0xa9, 0xb8, 0xe2, 0x81, 0x30, 0x2d, 0x29,
	0x52, 0x0b, 0x21, 0x04, 0x96, 0x42, 0xc4, 0xef, 0xba, 0x01, 0x65, 0xaa, 0xff, 0xbb, 0xdf, 0x60,
	0x75, 0x76, 0x4e, 0x2a, 0x6b, 0x0f, 0x56, 0x24, 0x86, 0x63, 0x76, 0xd3, 0xf7, 0x62, 0x19, 0x05,
	0x32, 0x6a, 0x38, 0xad, 0x42, 0xbb, 0x46, 0xeb, 0x06, 0x3d, 0x36, 0xa0, 0xbe, 0xe3, 0x11, 0x13,
	0x43, 0x9c, 0x96, 0x2d, 0x9a, 0x32, 0x83, 0xa6, 0x65, 0xee, 0x18, 0xd6, 0xf4, 0x09, 0x3d, 0xe1,
	0xe3, 0xf5, 0x3f, 0xf3, 0xb2, 0x01, 0xa5, 0x0b, 0x3e, 0x56, 0x28, 0xd3, 0xeb, 0x4c, 0x23, 0x77,
	0x00, 0xc4, 0x3e, 0xed, 0xbf, 0x38, 0xfa, 0x02, 0x55, 0x8a, 0xcc, 0x7f, 0xb0, 0x97, 0x0d, 0x28,
	0x19, 0x3e, 0xed, 0xa3, 0x46, 0xd3, 0x88, 0x3c, 0x86, 0xe2, 0x98, 0x4f, 0xb8, 0xd2, 0x16, 0x8a,
	0xd4, 0x04, 0xc9, 0x10, 0x1a, 0xf6, 0x54, 0xbb, 0x35, 0x63, 0xce, 0x1f, 0x67, 0x2c, 0x11, 0x23,
	0xf0, 0x5a, 0xf5, 0x73, 0x07, 0x42, 0x02, 0x19, 0xf5, 0xee, 0x3b, 0x58, 0x3d, 0x8f, 0x07, 0x91,
	0x27, 0xf9, 0x00, 0xff, 0xd6, 0x81, 0xbb, 0x0e, 0x6b, 0x16, 0x99, 0x11, 0x7c, 0xf8, 0xab, 0x00,
	0xcb, 0x27, 0x46, 0x21, 0x39, 0x82, 0xa5, 0x64, 0xbf, 0xc9, 0x13, 0x5b, 0xb3, 0xf5, 0x16, 0x69,
	0x36, 0xe6, 0x13, 0x86, 0xc6, 0x5d, 0x20, 0xaf, 0xa1, 0xa8, 0xf7, 0x85, 0xe4, 0x8a, 0xec, 0x0d,
	0x6d, 0x6e, 0xde, 0x91, 0x99, 0x3e, 0x7f, 0x02, 0xe5, 0x6c, 0xb6, 0xc9, 0x96, 0x5d, 0x78, 0x6b,
	0xb3, 0x9a, 0xdb, 0x77, 0x27, 0xa7, 0x44, 0xef, 0x01, 0x66, 0x43, 0x45, 0x76, 0xe6, 0xaa, 0xed,
	0xd1, 0x6e, 0xee, 0xde, 0x97, 0x9e, 0xd2, 0x1d, 0xc1, 0x52, 0x72, 0xc3, 0xf9, 0xa6, 0x58, 0x13,
	0x95, 0x6f, 0x8a, 0x3d, 0x0c, 0xee, 0x02, 0x39, 0x85, 0xca, 0xb4, 0xe5, 0x24, 0x27, 0xfc, 0xf6,
	0xb5, 0x36, 0x77, 0xee, 0xc9, 0x66, 0x5c, 0x2f, 0x9c, 0x37, 0x07, 0x9f, 0xf7, 0x87, 0x5c, 0x8d,
	0xe2, 0x41, 0xc7, 0x0b, 0x26, 0xdd, 0x01, 0x0a, 0x64, 0x6a, 0xf4, 0xdc, 0x0b, 0x24, 0x76, 0xe7,
	0x3f, 0x30, 0x83, 0x92, 0xfe, 0x79, 0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x84, 0xc8, 0xd3, 0xf6,
	0x7d, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	QueryLog(ctx context.Context, in *QueryLogRequest, opts ...grpc.CallOption) (*QueryLogResponse, error)
	QueryIndex(ctx context.Context, in *QueryIndexRequest, opts ...grpc.CallOption) (*QueryIndexResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Gateway_SubscribeClient, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/gateway.v1.Gateway/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/gateway.v1.Gateway/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) QueryLog(ctx context.Context, in *QueryLogRequest, opts ...grpc.CallOption) (*QueryLogResponse, error) {
	out := new(QueryLogResponse)
	err := c.cc.Invoke(ctx, "/gateway.v1.Gateway/QueryLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) QueryIndex(ctx context.Context, in *QueryIndexRequest, opts ...grpc.CallOption) (*QueryIndexResponse, error) {
	out := new(QueryIndexResponse)
	err := c.cc.Invoke(ctx, "/gateway.v1.Gateway/QueryIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/gateway.v1.Gateway/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Gateway_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Gateway_serviceDesc.Streams[0], "/gateway.v1.Gateway/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewaySubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gateway_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type gatewaySubscribeClient struct {
	grpc.ClientStream
}

func (x *gatewaySubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	QueryLog(context.Context, *QueryLogRequest) (*QueryLogResponse, error)
	QueryIndex(context.Context, *QueryIndexRequest) (*QueryIndexResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Subscribe(*SubscribeRequest, Gateway_SubscribeServer) error
}

// UnimplementedGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (*UnimplementedGatewayServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedGatewayServer) Write(ctx context.Context, req *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedGatewayServer) QueryLog(ctx context.Context, req *QueryLogRequest) (*QueryLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryLog not implemented")
}
func (*UnimplementedGatewayServer) QueryIndex(ctx context.Context, req *QueryIndexRequest) (*QueryIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryIndex not implemented")
}
func (*UnimplementedGatewayServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedGatewayServer) Subscribe(req *SubscribeRequest, srv Gateway_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.v1.Gateway/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.v1.Gateway/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_QueryLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).QueryLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.v1.Gateway/QueryLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).QueryLog(ctx, req.(*QueryLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_QueryIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).QueryIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.v1.Gateway/QueryIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).QueryIndex(ctx, req.(*QueryIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.v1.Gateway/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GatewayServer).Subscribe(m, &gatewaySubscribeServer{stream})
}

type Gateway_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type gatewaySubscribeServer struct {
	grpc.ServerStream
}

func (x *gatewaySubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.v1.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Gateway_Ping_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Gateway_Write_Handler,
		},
		{
			MethodName: "QueryLog",
			Handler:    _Gateway_QueryLog_Handler,
		},
		{
			MethodName: "QueryIndex",
			Handler:    _Gateway_QueryIndex_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Gateway_Read_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Gateway_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gateway/grpc/proto/gateway.proto",
}
