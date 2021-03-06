// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/metrics/v4alpha/metrics_service.proto

package envoy_config_metrics_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type MetricsServiceConfig struct {
	GrpcService            *v4alpha.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	ReportCountersAsDeltas *wrappers.BoolValue  `protobuf:"bytes,2,opt,name=report_counters_as_deltas,json=reportCountersAsDeltas,proto3" json:"report_counters_as_deltas,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-"`
	XXX_unrecognized       []byte               `json:"-"`
	XXX_sizecache          int32                `json:"-"`
}

func (m *MetricsServiceConfig) Reset()         { *m = MetricsServiceConfig{} }
func (m *MetricsServiceConfig) String() string { return proto.CompactTextString(m) }
func (*MetricsServiceConfig) ProtoMessage()    {}
func (*MetricsServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_011de8cb9b8654e7, []int{0}
}

func (m *MetricsServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsServiceConfig.Unmarshal(m, b)
}
func (m *MetricsServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsServiceConfig.Marshal(b, m, deterministic)
}
func (m *MetricsServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsServiceConfig.Merge(m, src)
}
func (m *MetricsServiceConfig) XXX_Size() int {
	return xxx_messageInfo_MetricsServiceConfig.Size(m)
}
func (m *MetricsServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsServiceConfig proto.InternalMessageInfo

func (m *MetricsServiceConfig) GetGrpcService() *v4alpha.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *MetricsServiceConfig) GetReportCountersAsDeltas() *wrappers.BoolValue {
	if m != nil {
		return m.ReportCountersAsDeltas
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsServiceConfig)(nil), "envoy.config.metrics.v4alpha.MetricsServiceConfig")
}

func init() {
	proto.RegisterFile("envoy/config/metrics/v4alpha/metrics_service.proto", fileDescriptor_011de8cb9b8654e7)
}

var fileDescriptor_011de8cb9b8654e7 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x65, 0x2b, 0x16, 0xd9, 0x7a, 0x28, 0x55, 0xb4, 0x16, 0x2d, 0xea, 0x41, 0x44, 0x4a, 0x02,
	0xad, 0x27, 0x6f, 0x6e, 0x05, 0x0f, 0x22, 0x94, 0x16, 0xbd, 0x2e, 0x69, 0x9a, 0xae, 0x81, 0x35,
	0x13, 0x26, 0xd9, 0xd5, 0xde, 0x3c, 0x8a, 0x9f, 0xe0, 0xa7, 0x78, 0x17, 0xbc, 0xfa, 0x35, 0x82,
	0x27, 0xe9, 0x66, 0x57, 0x5d, 0x2c, 0xde, 0x32, 0x79, 0xf3, 0xde, 0xbc, 0x37, 0xe3, 0x77, 0x85,
	0x4a, 0x61, 0x46, 0x39, 0xa8, 0xa9, 0x8c, 0xe8, 0xad, 0xb0, 0x28, 0xb9, 0xa1, 0xe9, 0x31, 0x8b,
	0xf5, 0x0d, 0x2b, 0xea, 0xd0, 0x08, 0x4c, 0x25, 0x17, 0x44, 0x23, 0x58, 0x68, 0x6c, 0x67, 0x1c,
	0xe2, 0x38, 0x24, 0xef, 0x21, 0x39, 0xa7, 0xd5, 0x29, 0x29, 0x72, 0x40, 0xf1, 0x2d, 0x17, 0xa1,
	0xe6, 0x65, 0xad, 0x56, 0x3b, 0x02, 0x88, 0x62, 0x41, 0xb3, 0x6a, 0x9c, 0x4c, 0xe9, 0x1d, 0x32,
	0xad, 0x05, 0x9a, 0x1c, 0xdf, 0x49, 0x26, 0x9a, 0x51, 0xa6, 0x14, 0x58, 0x66, 0x25, 0x28, 0x43,
	0x8d, 0x65, 0x36, 0x29, 0xe0, 0xbd, 0x3f, 0x70, 0x2a, 0xd0, 0x48, 0x50, 0x52, 0x45, 0x79, 0xcb,
	0x66, 0xca, 0x62, 0x39, 0x61, 0x56, 0xd0, 0xe2, 0xe1, 0x80, 0xfd, 0x0f, 0xcf, 0x5f, 0xbf, 0x74,
	0xe6, 0x47, 0xce, 0x53, 0x3f, 0xf3, 0xdc, 0x18, 0xf9, 0xab, 0xbf, 0x9d, 0x36, 0xbd, 0x5d, 0xef,
	0xb0, 0xd6, 0x3d, 0x20, 0xa5, 0xd8, 0xf3, 0x60, 0x45, 0x66, 0x72, 0x8e, 0x9a, 0xe7, 0x1a, 0xc1,
	0xca, 0x67, 0xb0, 0xfc, 0xe4, 0x55, 0xea, 0xde, 0xb0, 0x16, 0xfd, 0x7c, 0x37, 0xae, 0xfc, 0x2d,
	0x14, 0x1a, 0xd0, 0x86, 0x1c, 0x12, 0x65, 0x05, 0x9a, 0x90, 0x99, 0x70, 0x22, 0x62, 0xcb, 0x4c,
	0xb3, 0x92, 0x4d, 0x68, 0x11, 0xb7, 0x0c, 0x52, 0x2c, 0x83, 0x04, 0x00, 0xf1, 0x35, 0x8b, 0x13,
	0x31, 0xdc, 0x70, 0xe4, 0x7e, 0xce, 0x3d, 0x35, 0x67, 0x19, 0xf3, 0xa4, 0xf7, 0xfc, 0xfa, 0xd8,
	0x26, 0x7e, 0x67, 0xf1, 0x49, 0x7a, 0x64, 0x51, 0xc0, 0xe0, 0xe2, 0xe5, 0xe1, 0xed, 0xbd, 0x5a,
	0xa9, 0x2f, 0xf9, 0x47, 0x12, 0x5c, 0x2c, 0x8d, 0x70, 0x3f, 0x23, 0xff, 0x1d, 0x36, 0x58, 0x2b,
	0x6b, 0x0d, 0xe6, 0x26, 0x07, 0xde, 0xb8, 0x9a, 0xb9, 0xed, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x8a, 0x86, 0xf5, 0x89, 0x4a, 0x02, 0x00, 0x00,
}
