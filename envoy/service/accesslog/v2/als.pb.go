// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/accesslog/v2/als.proto

package envoy_service_accesslog_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type StreamAccessLogsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamAccessLogsResponse) Reset()         { *m = StreamAccessLogsResponse{} }
func (m *StreamAccessLogsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsResponse) ProtoMessage()    {}
func (*StreamAccessLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{0}
}

func (m *StreamAccessLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsResponse.Unmarshal(m, b)
}
func (m *StreamAccessLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsResponse.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsResponse.Merge(m, src)
}
func (m *StreamAccessLogsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsResponse.Size(m)
}
func (m *StreamAccessLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsResponse proto.InternalMessageInfo

type StreamAccessLogsMessage struct {
	Identifier *StreamAccessLogsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Types that are valid to be assigned to LogEntries:
	//	*StreamAccessLogsMessage_HttpLogs
	//	*StreamAccessLogsMessage_TcpLogs
	LogEntries           isStreamAccessLogsMessage_LogEntries `protobuf_oneof:"log_entries"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *StreamAccessLogsMessage) Reset()         { *m = StreamAccessLogsMessage{} }
func (m *StreamAccessLogsMessage) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage) ProtoMessage()    {}
func (*StreamAccessLogsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1}
}

func (m *StreamAccessLogsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage.Merge(m, src)
}
func (m *StreamAccessLogsMessage) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage.Size(m)
}
func (m *StreamAccessLogsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage proto.InternalMessageInfo

func (m *StreamAccessLogsMessage) GetIdentifier() *StreamAccessLogsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

type isStreamAccessLogsMessage_LogEntries interface {
	isStreamAccessLogsMessage_LogEntries()
}

type StreamAccessLogsMessage_HttpLogs struct {
	HttpLogs *StreamAccessLogsMessage_HTTPAccessLogEntries `protobuf:"bytes,2,opt,name=http_logs,json=httpLogs,proto3,oneof"`
}

type StreamAccessLogsMessage_TcpLogs struct {
	TcpLogs *StreamAccessLogsMessage_TCPAccessLogEntries `protobuf:"bytes,3,opt,name=tcp_logs,json=tcpLogs,proto3,oneof"`
}

func (*StreamAccessLogsMessage_HttpLogs) isStreamAccessLogsMessage_LogEntries() {}

func (*StreamAccessLogsMessage_TcpLogs) isStreamAccessLogsMessage_LogEntries() {}

func (m *StreamAccessLogsMessage) GetLogEntries() isStreamAccessLogsMessage_LogEntries {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetHttpLogs() *StreamAccessLogsMessage_HTTPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_HttpLogs); ok {
		return x.HttpLogs
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetTcpLogs() *StreamAccessLogsMessage_TCPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_TcpLogs); ok {
		return x.TcpLogs
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StreamAccessLogsMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StreamAccessLogsMessage_HttpLogs)(nil),
		(*StreamAccessLogsMessage_TcpLogs)(nil),
	}
}

type StreamAccessLogsMessage_Identifier struct {
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	LogName              string     `protobuf:"bytes,2,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamAccessLogsMessage_Identifier) Reset()         { *m = StreamAccessLogsMessage_Identifier{} }
func (m *StreamAccessLogsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage_Identifier) ProtoMessage()    {}
func (*StreamAccessLogsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1, 0}
}

func (m *StreamAccessLogsMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_Identifier.Merge(m, src)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Size(m)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_Identifier proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *StreamAccessLogsMessage_Identifier) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

type StreamAccessLogsMessage_HTTPAccessLogEntries struct {
	LogEntry             []*v2.HTTPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_HTTPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1, 1}
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Merge(m, src)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Size(m)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) GetLogEntry() []*v2.HTTPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

type StreamAccessLogsMessage_TCPAccessLogEntries struct {
	LogEntry             []*v2.TCPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_TCPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1, 2}
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Merge(m, src)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Size(m)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) GetLogEntry() []*v2.TCPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamAccessLogsResponse)(nil), "envoy.service.accesslog.v2.StreamAccessLogsResponse")
	proto.RegisterType((*StreamAccessLogsMessage)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage")
	proto.RegisterType((*StreamAccessLogsMessage_Identifier)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.Identifier")
	proto.RegisterType((*StreamAccessLogsMessage_HTTPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.HTTPAccessLogEntries")
	proto.RegisterType((*StreamAccessLogsMessage_TCPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.TCPAccessLogEntries")
}

func init() {
	proto.RegisterFile("envoy/service/accesslog/v2/als.proto", fileDescriptor_e4f3a3a69261b513)
}

var fileDescriptor_e4f3a3a69261b513 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x77, 0xda, 0xdd, 0x6d, 0xfa, 0x7a, 0x29, 0xa3, 0xd0, 0x12, 0x14, 0x4a, 0x11, 0x2c,
	0x0a, 0x09, 0x64, 0xf5, 0x2a, 0x34, 0x22, 0x56, 0xd0, 0xa5, 0x66, 0x7b, 0x76, 0x99, 0x4d, 0x9e,
	0xd9, 0x91, 0x74, 0x26, 0x64, 0x66, 0x83, 0xbd, 0xe9, 0x4d, 0x3c, 0x7a, 0xf0, 0xe2, 0x37, 0xf1,
	0x13, 0x78, 0xf5, 0xdb, 0xc8, 0x9e, 0x64, 0x26, 0xd9, 0xa8, 0xdd, 0x56, 0xd8, 0xde, 0x42, 0xe6,
	0xff, 0x7e, 0xbf, 0x79, 0x33, 0xf3, 0xe0, 0x1e, 0x8a, 0x52, 0xae, 0x7c, 0x85, 0x45, 0xc9, 0x63,
	0xf4, 0x59, 0x1c, 0xa3, 0x52, 0x99, 0x4c, 0xfd, 0x32, 0xf0, 0x59, 0xa6, 0xbc, 0xbc, 0x90, 0x5a,
	0x52, 0xd7, 0xa6, 0xbc, 0x3a, 0xe5, 0x35, 0x29, 0xaf, 0x0c, 0xdc, 0x3b, 0x15, 0x81, 0xe5, 0xdc,
	0xd4, 0xc4, 0xb2, 0x40, 0xff, 0x8c, 0x29, 0xac, 0x2a, 0xdd, 0xfb, 0xd5, 0x6a, 0xc2, 0x34, 0x5b,
	0x83, 0x37, 0x8c, 0x2a, 0x78, 0xf7, 0x22, 0xc9, 0x99, 0xcf, 0x84, 0x90, 0x9a, 0x69, 0x2e, 0x85,
	0xf2, 0x95, 0x66, 0xfa, 0xa2, 0xde, 0x81, 0x3b, 0x28, 0x59, 0xc6, 0x13, 0xa6, 0xd1, 0xbf, 0xfa,
	0xa8, 0x16, 0xc6, 0x2e, 0x0c, 0x4f, 0x74, 0x81, 0x6c, 0x39, 0xb5, 0xc0, 0x97, 0x32, 0x55, 0x11,
	0xaa, 0x5c, 0x0a, 0x85, 0xe3, 0x6f, 0x07, 0x30, 0x58, 0x5f, 0x7c, 0x85, 0x4a, 0xb1, 0x14, 0xe9,
	0x1b, 0x00, 0x9e, 0xa0, 0xd0, 0xfc, 0x2d, 0xc7, 0x62, 0x48, 0x46, 0x64, 0xd2, 0x0b, 0x9e, 0x78,
	0xdb, 0xfb, 0xf4, 0xb6, 0x80, 0xbc, 0x17, 0x0d, 0x25, 0xfa, 0x8b, 0x48, 0x53, 0xe8, 0x9e, 0x6b,
	0x9d, 0x9f, 0x66, 0x32, 0x55, 0xc3, 0x96, 0xc5, 0xcf, 0x76, 0xc1, 0xcf, 0x16, 0x8b, 0x79, 0xf3,
	0xf7, 0x99, 0xd0, 0x05, 0x47, 0x35, 0xdb, 0x8b, 0x1c, 0x03, 0x37, 0x39, 0x9a, 0x80, 0xa3, 0xe3,
	0xda, 0xd3, 0xb6, 0x9e, 0xe7, 0xbb, 0x78, 0x16, 0x4f, 0x37, 0x69, 0x3a, 0x3a, 0xb6, 0x16, 0x37,
	0x05, 0xf8, 0xd3, 0x28, 0x7d, 0x0c, 0xfb, 0x42, 0x26, 0x58, 0x1f, 0xdb, 0xa0, 0xf6, 0xb1, 0x9c,
	0x1b, 0x83, 0x79, 0x02, 0xde, 0xb1, 0x4c, 0x30, 0x74, 0x2e, 0xc3, 0x83, 0xcf, 0xa4, 0xd5, 0x27,
	0x91, 0x8d, 0xd3, 0x31, 0x38, 0x99, 0x4c, 0x4f, 0x05, 0x5b, 0xa2, 0x3d, 0x92, 0x6e, 0xd8, 0xb9,
	0x0c, 0xf7, 0x8b, 0xd6, 0x88, 0x44, 0x9d, 0x4c, 0xa6, 0xc7, 0x6c, 0x89, 0xee, 0x3b, 0xb8, 0xbd,
	0xa9, 0x65, 0x1a, 0x41, 0xd7, 0xd4, 0xa2, 0xd0, 0xc5, 0x6a, 0x48, 0x46, 0xed, 0x49, 0x2f, 0x78,
	0x58, 0x7b, 0xcd, 0xe3, 0xfa, 0xb7, 0xc9, 0x6b, 0x84, 0x95, 0xdd, 0xcb, 0x17, 0xd2, 0x72, 0x48,
	0x64, 0xf6, 0x60, 0xff, 0xb9, 0xe7, 0x70, 0x6b, 0x43, 0xdb, 0xf4, 0xf5, 0x75, 0xd5, 0x83, 0xad,
	0xaa, 0x75, 0xc0, 0x46, 0x53, 0x48, 0xa1, 0x77, 0x85, 0x34, 0x86, 0xf6, 0xaf, 0x90, 0x04, 0x5f,
	0x09, 0xf4, 0x9b, 0xd2, 0x93, 0xea, 0xae, 0xe8, 0x47, 0x02, 0xfd, 0xf5, 0x2b, 0xa2, 0x47, 0x3b,
	0x5c, 0xa8, 0xfb, 0xe8, 0x26, 0x45, 0xcd, 0xc8, 0xec, 0x4d, 0x48, 0x38, 0xfd, 0xfe, 0xe1, 0xc7,
	0xcf, 0xc3, 0x56, 0x9f, 0xc0, 0x84, 0xcb, 0x8a, 0x92, 0x17, 0xf2, 0xfd, 0xea, 0x3f, 0xc0, 0xd0,
	0x99, 0x66, 0x6a, 0x6e, 0x06, 0x72, 0x4e, 0x3e, 0x11, 0x72, 0x76, 0x68, 0x87, 0xf3, 0xe8, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x3b, 0x59, 0x76, 0x5f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessLogServiceClient is the client API for AccessLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessLogServiceClient interface {
	StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error)
}

type accessLogServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessLogServiceClient(cc *grpc.ClientConn) AccessLogServiceClient {
	return &accessLogServiceClient{cc}
}

func (c *accessLogServiceClient) StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccessLogService_serviceDesc.Streams[0], "/envoy.service.accesslog.v2.AccessLogService/StreamAccessLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &accessLogServiceStreamAccessLogsClient{stream}
	return x, nil
}

type AccessLogService_StreamAccessLogsClient interface {
	Send(*StreamAccessLogsMessage) error
	CloseAndRecv() (*StreamAccessLogsResponse, error)
	grpc.ClientStream
}

type accessLogServiceStreamAccessLogsClient struct {
	grpc.ClientStream
}

func (x *accessLogServiceStreamAccessLogsClient) Send(m *StreamAccessLogsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsClient) CloseAndRecv() (*StreamAccessLogsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamAccessLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccessLogServiceServer is the server API for AccessLogService service.
type AccessLogServiceServer interface {
	StreamAccessLogs(AccessLogService_StreamAccessLogsServer) error
}

// UnimplementedAccessLogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccessLogServiceServer struct {
}

func (*UnimplementedAccessLogServiceServer) StreamAccessLogs(srv AccessLogService_StreamAccessLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAccessLogs not implemented")
}

func RegisterAccessLogServiceServer(s *grpc.Server, srv AccessLogServiceServer) {
	s.RegisterService(&_AccessLogService_serviceDesc, srv)
}

func _AccessLogService_StreamAccessLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessLogServiceServer).StreamAccessLogs(&accessLogServiceStreamAccessLogsServer{stream})
}

type AccessLogService_StreamAccessLogsServer interface {
	SendAndClose(*StreamAccessLogsResponse) error
	Recv() (*StreamAccessLogsMessage, error)
	grpc.ServerStream
}

type accessLogServiceStreamAccessLogsServer struct {
	grpc.ServerStream
}

func (x *accessLogServiceStreamAccessLogsServer) SendAndClose(m *StreamAccessLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsServer) Recv() (*StreamAccessLogsMessage, error) {
	m := new(StreamAccessLogsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AccessLogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.accesslog.v2.AccessLogService",
	HandlerType: (*AccessLogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAccessLogs",
			Handler:       _AccessLogService_StreamAccessLogs_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/accesslog/v2/als.proto",
}
