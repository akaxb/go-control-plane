// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/tracers/opencensus/v4alpha/opencensus.proto

package envoy_extensions_tracers_opencensus_v4alpha

import (
	fmt "fmt"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type OpenCensusConfig_TraceContext int32

const (
	OpenCensusConfig_NONE                OpenCensusConfig_TraceContext = 0
	OpenCensusConfig_TRACE_CONTEXT       OpenCensusConfig_TraceContext = 1
	OpenCensusConfig_GRPC_TRACE_BIN      OpenCensusConfig_TraceContext = 2
	OpenCensusConfig_CLOUD_TRACE_CONTEXT OpenCensusConfig_TraceContext = 3
	OpenCensusConfig_B3                  OpenCensusConfig_TraceContext = 4
)

var OpenCensusConfig_TraceContext_name = map[int32]string{
	0: "NONE",
	1: "TRACE_CONTEXT",
	2: "GRPC_TRACE_BIN",
	3: "CLOUD_TRACE_CONTEXT",
	4: "B3",
}

var OpenCensusConfig_TraceContext_value = map[string]int32{
	"NONE":                0,
	"TRACE_CONTEXT":       1,
	"GRPC_TRACE_BIN":      2,
	"CLOUD_TRACE_CONTEXT": 3,
	"B3":                  4,
}

func (x OpenCensusConfig_TraceContext) String() string {
	return proto.EnumName(OpenCensusConfig_TraceContext_name, int32(x))
}

func (OpenCensusConfig_TraceContext) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ab1caa00b30f88d, []int{0, 0}
}

type OpenCensusConfig struct {
	TraceConfig                *v1.TraceConfig                 `protobuf:"bytes,1,opt,name=trace_config,json=traceConfig,proto3" json:"trace_config,omitempty"`
	StdoutExporterEnabled      bool                            `protobuf:"varint,2,opt,name=stdout_exporter_enabled,json=stdoutExporterEnabled,proto3" json:"stdout_exporter_enabled,omitempty"`
	StackdriverExporterEnabled bool                            `protobuf:"varint,3,opt,name=stackdriver_exporter_enabled,json=stackdriverExporterEnabled,proto3" json:"stackdriver_exporter_enabled,omitempty"`
	StackdriverProjectId       string                          `protobuf:"bytes,4,opt,name=stackdriver_project_id,json=stackdriverProjectId,proto3" json:"stackdriver_project_id,omitempty"`
	StackdriverAddress         string                          `protobuf:"bytes,10,opt,name=stackdriver_address,json=stackdriverAddress,proto3" json:"stackdriver_address,omitempty"`
	StackdriverGrpcService     *v4alpha.GrpcService            `protobuf:"bytes,13,opt,name=stackdriver_grpc_service,json=stackdriverGrpcService,proto3" json:"stackdriver_grpc_service,omitempty"`
	ZipkinExporterEnabled      bool                            `protobuf:"varint,5,opt,name=zipkin_exporter_enabled,json=zipkinExporterEnabled,proto3" json:"zipkin_exporter_enabled,omitempty"`
	ZipkinUrl                  string                          `protobuf:"bytes,6,opt,name=zipkin_url,json=zipkinUrl,proto3" json:"zipkin_url,omitempty"`
	OcagentExporterEnabled     bool                            `protobuf:"varint,11,opt,name=ocagent_exporter_enabled,json=ocagentExporterEnabled,proto3" json:"ocagent_exporter_enabled,omitempty"`
	OcagentAddress             string                          `protobuf:"bytes,12,opt,name=ocagent_address,json=ocagentAddress,proto3" json:"ocagent_address,omitempty"`
	OcagentGrpcService         *v4alpha.GrpcService            `protobuf:"bytes,14,opt,name=ocagent_grpc_service,json=ocagentGrpcService,proto3" json:"ocagent_grpc_service,omitempty"`
	IncomingTraceContext       []OpenCensusConfig_TraceContext `protobuf:"varint,8,rep,packed,name=incoming_trace_context,json=incomingTraceContext,proto3,enum=envoy.extensions.tracers.opencensus.v4alpha.OpenCensusConfig_TraceContext" json:"incoming_trace_context,omitempty"`
	OutgoingTraceContext       []OpenCensusConfig_TraceContext `protobuf:"varint,9,rep,packed,name=outgoing_trace_context,json=outgoingTraceContext,proto3,enum=envoy.extensions.tracers.opencensus.v4alpha.OpenCensusConfig_TraceContext" json:"outgoing_trace_context,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                        `json:"-"`
	XXX_unrecognized           []byte                          `json:"-"`
	XXX_sizecache              int32                           `json:"-"`
}

func (m *OpenCensusConfig) Reset()         { *m = OpenCensusConfig{} }
func (m *OpenCensusConfig) String() string { return proto.CompactTextString(m) }
func (*OpenCensusConfig) ProtoMessage()    {}
func (*OpenCensusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ab1caa00b30f88d, []int{0}
}

func (m *OpenCensusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenCensusConfig.Unmarshal(m, b)
}
func (m *OpenCensusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenCensusConfig.Marshal(b, m, deterministic)
}
func (m *OpenCensusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenCensusConfig.Merge(m, src)
}
func (m *OpenCensusConfig) XXX_Size() int {
	return xxx_messageInfo_OpenCensusConfig.Size(m)
}
func (m *OpenCensusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenCensusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OpenCensusConfig proto.InternalMessageInfo

func (m *OpenCensusConfig) GetTraceConfig() *v1.TraceConfig {
	if m != nil {
		return m.TraceConfig
	}
	return nil
}

func (m *OpenCensusConfig) GetStdoutExporterEnabled() bool {
	if m != nil {
		return m.StdoutExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetStackdriverExporterEnabled() bool {
	if m != nil {
		return m.StackdriverExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetStackdriverProjectId() string {
	if m != nil {
		return m.StackdriverProjectId
	}
	return ""
}

func (m *OpenCensusConfig) GetStackdriverAddress() string {
	if m != nil {
		return m.StackdriverAddress
	}
	return ""
}

func (m *OpenCensusConfig) GetStackdriverGrpcService() *v4alpha.GrpcService {
	if m != nil {
		return m.StackdriverGrpcService
	}
	return nil
}

func (m *OpenCensusConfig) GetZipkinExporterEnabled() bool {
	if m != nil {
		return m.ZipkinExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetZipkinUrl() string {
	if m != nil {
		return m.ZipkinUrl
	}
	return ""
}

func (m *OpenCensusConfig) GetOcagentExporterEnabled() bool {
	if m != nil {
		return m.OcagentExporterEnabled
	}
	return false
}

func (m *OpenCensusConfig) GetOcagentAddress() string {
	if m != nil {
		return m.OcagentAddress
	}
	return ""
}

func (m *OpenCensusConfig) GetOcagentGrpcService() *v4alpha.GrpcService {
	if m != nil {
		return m.OcagentGrpcService
	}
	return nil
}

func (m *OpenCensusConfig) GetIncomingTraceContext() []OpenCensusConfig_TraceContext {
	if m != nil {
		return m.IncomingTraceContext
	}
	return nil
}

func (m *OpenCensusConfig) GetOutgoingTraceContext() []OpenCensusConfig_TraceContext {
	if m != nil {
		return m.OutgoingTraceContext
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.extensions.tracers.opencensus.v4alpha.OpenCensusConfig_TraceContext", OpenCensusConfig_TraceContext_name, OpenCensusConfig_TraceContext_value)
	proto.RegisterType((*OpenCensusConfig)(nil), "envoy.extensions.tracers.opencensus.v4alpha.OpenCensusConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/tracers/opencensus/v4alpha/opencensus.proto", fileDescriptor_8ab1caa00b30f88d)
}

var fileDescriptor_8ab1caa00b30f88d = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x18, 0x25, 0x6d, 0x29, 0x9d, 0xb7, 0x75, 0xc1, 0x1b, 0x5b, 0x54, 0x31, 0x54, 0x76, 0x18, 0x95,
	0x18, 0x89, 0xf6, 0x43, 0x08, 0x10, 0x07, 0xd6, 0x50, 0x4d, 0x9b, 0x50, 0x5b, 0x85, 0x0e, 0xed,
	0x96, 0x79, 0x89, 0x29, 0x66, 0xc5, 0x8e, 0x1c, 0x37, 0xea, 0x38, 0xed, 0xc8, 0xdf, 0xc0, 0x9f,
	0xc2, 0x1d, 0x89, 0xeb, 0xfe, 0x23, 0x14, 0x3b, 0xe9, 0xdc, 0x94, 0xcb, 0x24, 0x6e, 0xee, 0xf7,
	0xbe, 0xf7, 0xbe, 0xd7, 0xe7, 0xcf, 0x01, 0x6f, 0x31, 0x4d, 0xd8, 0x95, 0x83, 0x27, 0x02, 0xd3,
	0x98, 0x30, 0x1a, 0x3b, 0x82, 0xa3, 0x00, 0xf3, 0xd8, 0x61, 0x11, 0xa6, 0x01, 0xa6, 0xf1, 0x38,
	0x76, 0x92, 0x03, 0x34, 0x8a, 0xbe, 0x20, 0xad, 0x64, 0x47, 0x9c, 0x09, 0x06, 0x9f, 0x4b, 0xb6,
	0x7d, 0xcb, 0xb6, 0x33, 0xb6, 0xad, 0xb5, 0x66, 0xec, 0xc6, 0x8e, 0x1a, 0x15, 0x30, 0xfa, 0x99,
	0x0c, 0x9d, 0x80, 0x71, 0x3c, 0x15, 0x1e, 0xf2, 0x28, 0xf0, 0x63, 0xcc, 0x13, 0x12, 0x60, 0x25,
	0xdd, 0xd8, 0xd1, 0xe6, 0xcb, 0x8a, 0x32, 0xe6, 0x24, 0xbb, 0xea, 0xe0, 0x2b, 0x9d, 0xac, 0x7b,
	0x73, 0x1c, 0x46, 0xc8, 0x41, 0x94, 0x32, 0x81, 0x84, 0xfc, 0x1b, 0xb1, 0x40, 0x22, 0xf7, 0xd9,
	0x78, 0x3a, 0x07, 0x27, 0x98, 0xa7, 0x86, 0x09, 0xcd, 0x15, 0x36, 0x12, 0x34, 0x22, 0x21, 0x12,
	0xd8, 0xc9, 0x0f, 0x0a, 0xd8, 0xba, 0xa9, 0x01, 0xb3, 0x17, 0x61, 0xea, 0x4a, 0x2f, 0xae, 0x9c,
	0x0a, 0x8f, 0xc1, 0x92, 0xee, 0xc2, 0x32, 0x9a, 0x46, 0x6b, 0x71, 0x6f, 0xdb, 0x2e, 0x26, 0xa4,
	0xf2, 0xb0, 0x93, 0x5d, 0x7b, 0x90, 0x1e, 0x14, 0xdb, 0x5b, 0x14, 0xb7, 0x3f, 0xe0, 0x4b, 0xb0,
	0x11, 0x8b, 0x90, 0x8d, 0x85, 0x8f, 0x27, 0x11, 0xe3, 0x02, 0x73, 0x1f, 0x53, 0x74, 0x31, 0xc2,
	0xa1, 0x55, 0x6a, 0x1a, 0xad, 0x9a, 0xf7, 0x48, 0xc1, 0x9d, 0x0c, 0xed, 0x28, 0x10, 0xbe, 0x03,
	0x8f, 0x63, 0x81, 0x82, 0xcb, 0x90, 0x93, 0x24, 0xe5, 0x14, 0xc9, 0x65, 0x49, 0x6e, 0x68, 0x3d,
	0x45, 0x85, 0x03, 0xb0, 0xae, 0x2b, 0x44, 0x9c, 0x7d, 0xc5, 0x81, 0xf0, 0x49, 0x68, 0x55, 0x9a,
	0x46, 0x6b, 0xc1, 0x5b, 0xd3, 0xd0, 0xbe, 0x02, 0x8f, 0x43, 0xe8, 0x80, 0x55, 0x9d, 0x85, 0xc2,
	0x90, 0xe3, 0x38, 0xb6, 0x80, 0xa4, 0x40, 0x0d, 0x3a, 0x54, 0x08, 0x3c, 0x07, 0x96, 0x4e, 0xd0,
	0xef, 0xda, 0x5a, 0xce, 0x72, 0x53, 0x7b, 0x94, 0x5d, 0x69, 0xba, 0x1a, 0xf9, 0xd6, 0xd8, 0x47,
	0x3c, 0x0a, 0x3e, 0xaa, 0x6e, 0x4f, 0xb7, 0xab, 0xd5, 0xd3, 0x08, 0xbf, 0x93, 0xe8, 0x92, 0xd0,
	0xf9, 0x14, 0xee, 0xab, 0x08, 0x15, 0x5c, 0x0c, 0x60, 0x13, 0x80, 0x8c, 0x37, 0xe6, 0x23, 0xab,
	0x2a, 0xff, 0xc1, 0x82, 0xaa, 0x9c, 0xf2, 0x11, 0x7c, 0x05, 0x2c, 0x16, 0xa0, 0x21, 0xa6, 0xff,
	0xb8, 0x9a, 0x45, 0xa9, 0xbb, 0x9e, 0xe1, 0x45, 0xe1, 0x67, 0x60, 0x25, 0x67, 0xe6, 0xf9, 0x2c,
	0x49, 0xf5, 0x7a, 0x56, 0xce, 0xb3, 0x39, 0x03, 0x6b, 0x79, 0xe3, 0x4c, 0x2e, 0xf5, 0x3b, 0xe5,
	0x02, 0x33, 0x0d, 0x3d, 0x93, 0x6b, 0x03, 0xac, 0x13, 0x1a, 0xb0, 0x6f, 0x84, 0x0e, 0xfd, 0xe9,
	0xae, 0x0a, 0x3c, 0x11, 0x56, 0xad, 0x59, 0x6e, 0xd5, 0xf7, 0x4e, 0xec, 0x3b, 0x3c, 0x5e, 0xbb,
	0xf8, 0x02, 0xa6, 0xfb, 0x9c, 0x2a, 0x7a, 0x6b, 0xf9, 0x24, 0xbd, 0x2a, 0x2d, 0xb0, 0xb1, 0x18,
	0xb2, 0x79, 0x0b, 0x0b, 0xff, 0xdf, 0x42, 0x3e, 0x49, 0xaf, 0x6e, 0x9d, 0x83, 0xa5, 0x19, 0x4b,
	0x35, 0x50, 0xe9, 0xf6, 0xba, 0x1d, 0xf3, 0x1e, 0x7c, 0x08, 0x96, 0x07, 0xde, 0xa1, 0xdb, 0xf1,
	0xdd, 0x5e, 0x77, 0xd0, 0x39, 0x1b, 0x98, 0x06, 0x84, 0xa0, 0x7e, 0xe4, 0xf5, 0x5d, 0x5f, 0xd5,
	0xdb, 0xc7, 0x5d, 0xb3, 0x04, 0x37, 0xc0, 0xaa, 0xfb, 0xa1, 0x77, 0xfa, 0xde, 0x9f, 0x6d, 0x2e,
	0xc3, 0x2a, 0x28, 0xb5, 0xf7, 0xcd, 0xca, 0x9b, 0x17, 0x3f, 0x7f, 0xff, 0x78, 0xd2, 0x02, 0xdb,
	0x33, 0x37, 0x95, 0xbd, 0xfa, 0xfd, 0x39, 0xcf, 0x27, 0x95, 0xda, 0x03, 0xb3, 0xd6, 0xfe, 0xf4,
	0xeb, 0xfa, 0xcf, 0x4d, 0xb5, 0x64, 0x96, 0xc1, 0x6b, 0xc2, 0x54, 0x08, 0x11, 0x67, 0x93, 0xab,
	0xbb, 0xe4, 0xd1, 0x5e, 0xe9, 0x4d, 0x6b, 0xfd, 0xf4, 0x5b, 0xd3, 0x37, 0x2e, 0xaa, 0xf2, 0xa3,
	0xb3, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x67, 0x7c, 0xc0, 0x08, 0xd6, 0x05, 0x00, 0x00,
}
