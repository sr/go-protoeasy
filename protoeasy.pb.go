// Code generated by protoc-gen-go.
// source: protoeasy.proto
// DO NOT EDIT!

package protoeasy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/google-protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GoPluginType int32

const (
	GoPluginType_GO_PLUGIN_TYPE_NONE       GoPluginType = 0
	GoPluginType_GO_PLUGIN_TYPE_GO         GoPluginType = 1
	GoPluginType_GO_PLUGIN_TYPE_GOFAST     GoPluginType = 2
	GoPluginType_GO_PLUGIN_TYPE_GOGO       GoPluginType = 3
	GoPluginType_GO_PLUGIN_TYPE_GOGOFAST   GoPluginType = 4
	GoPluginType_GO_PLUGIN_TYPE_GOGOFASTER GoPluginType = 5
	GoPluginType_GO_PLUGIN_TYPE_GOGOSLICK  GoPluginType = 6
)

var GoPluginType_name = map[int32]string{
	0: "GO_PLUGIN_TYPE_NONE",
	1: "GO_PLUGIN_TYPE_GO",
	2: "GO_PLUGIN_TYPE_GOFAST",
	3: "GO_PLUGIN_TYPE_GOGO",
	4: "GO_PLUGIN_TYPE_GOGOFAST",
	5: "GO_PLUGIN_TYPE_GOGOFASTER",
	6: "GO_PLUGIN_TYPE_GOGOSLICK",
}
var GoPluginType_value = map[string]int32{
	"GO_PLUGIN_TYPE_NONE":       0,
	"GO_PLUGIN_TYPE_GO":         1,
	"GO_PLUGIN_TYPE_GOFAST":     2,
	"GO_PLUGIN_TYPE_GOGO":       3,
	"GO_PLUGIN_TYPE_GOGOFAST":   4,
	"GO_PLUGIN_TYPE_GOGOFASTER": 5,
	"GO_PLUGIN_TYPE_GOGOSLICK":  6,
}

func (x GoPluginType) String() string {
	return proto.EnumName(GoPluginType_name, int32(x))
}
func (GoPluginType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CompileOptions struct {
	Grpc                 bool              `protobuf:"varint,1,opt,name=grpc" json:"grpc,omitempty"`
	GrpcGateway          bool              `protobuf:"varint,2,opt,name=grpc_gateway" json:"grpc_gateway,omitempty"`
	ExcludePattern       []string          `protobuf:"bytes,3,rep,name=exclude_pattern" json:"exclude_pattern,omitempty"`
	Cpp                  bool              `protobuf:"varint,20,opt,name=cpp" json:"cpp,omitempty"`
	CppRelOut            string            `protobuf:"bytes,21,opt,name=cpp_rel_out" json:"cpp_rel_out,omitempty"`
	Csharp               bool              `protobuf:"varint,30,opt,name=csharp" json:"csharp,omitempty"`
	CsharpRelOut         string            `protobuf:"bytes,31,opt,name=csharp_rel_out" json:"csharp_rel_out,omitempty"`
	Go                   bool              `protobuf:"varint,40,opt,name=go" json:"go,omitempty"`
	GoRelOut             string            `protobuf:"bytes,41,opt,name=go_rel_out" json:"go_rel_out,omitempty"`
	GoImportPath         string            `protobuf:"bytes,42,opt,name=go_import_path" json:"go_import_path,omitempty"`
	GoNoDefaultModifiers bool              `protobuf:"varint,43,opt,name=go_no_default_modifiers" json:"go_no_default_modifiers,omitempty"`
	GoModifiers          map[string]string `protobuf:"bytes,44,rep,name=go_modifiers" json:"go_modifiers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	GoPluginType         GoPluginType      `protobuf:"varint,45,opt,name=go_plugin_type,enum=protoeasy.GoPluginType" json:"go_plugin_type,omitempty"`
	Objc                 bool              `protobuf:"varint,50,opt,name=objc" json:"objc,omitempty"`
	ObjcRelOut           string            `protobuf:"bytes,51,opt,name=objc_rel_out" json:"objc_rel_out,omitempty"`
	Python               bool              `protobuf:"varint,60,opt,name=python" json:"python,omitempty"`
	PythonRelOut         string            `protobuf:"bytes,61,opt,name=python_rel_out" json:"python_rel_out,omitempty"`
	Ruby                 bool              `protobuf:"varint,70,opt,name=ruby" json:"ruby,omitempty"`
	RubyRelOut           string            `protobuf:"bytes,71,opt,name=ruby_rel_out" json:"ruby_rel_out,omitempty"`
}

func (m *CompileOptions) Reset()                    { *m = CompileOptions{} }
func (m *CompileOptions) String() string            { return proto.CompactTextString(m) }
func (*CompileOptions) ProtoMessage()               {}
func (*CompileOptions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CompileOptions) GetGoModifiers() map[string]string {
	if m != nil {
		return m.GoModifiers
	}
	return nil
}

type Command struct {
	Arg []string `protobuf:"bytes,1,rep,name=arg" json:"arg,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CompileInfo struct {
	Command         []*Command                `protobuf:"bytes,1,rep,name=command" json:"command,omitempty"`
	InputSizeBytes  uint64                    `protobuf:"varint,2,opt,name=input_size_bytes" json:"input_size_bytes,omitempty"`
	OutputSizeBytes uint64                    `protobuf:"varint,3,opt,name=output_size_bytes" json:"output_size_bytes,omitempty"`
	Duration        *google_protobuf.Duration `protobuf:"bytes,4,opt,name=duration" json:"duration,omitempty"`
}

func (m *CompileInfo) Reset()                    { *m = CompileInfo{} }
func (m *CompileInfo) String() string            { return proto.CompactTextString(m) }
func (*CompileInfo) ProtoMessage()               {}
func (*CompileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CompileInfo) GetCommand() []*Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *CompileInfo) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type CompileRequest struct {
	Tar            []byte          `protobuf:"bytes,1,opt,name=tar,proto3" json:"tar,omitempty"`
	CompileOptions *CompileOptions `protobuf:"bytes,2,opt,name=compile_options" json:"compile_options,omitempty"`
}

func (m *CompileRequest) Reset()                    { *m = CompileRequest{} }
func (m *CompileRequest) String() string            { return proto.CompactTextString(m) }
func (*CompileRequest) ProtoMessage()               {}
func (*CompileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CompileRequest) GetCompileOptions() *CompileOptions {
	if m != nil {
		return m.CompileOptions
	}
	return nil
}

type CompileResponse struct {
	Tar         []byte       `protobuf:"bytes,1,opt,name=tar,proto3" json:"tar,omitempty"`
	CompileInfo *CompileInfo `protobuf:"bytes,2,opt,name=compile_info" json:"compile_info,omitempty"`
}

func (m *CompileResponse) Reset()                    { *m = CompileResponse{} }
func (m *CompileResponse) String() string            { return proto.CompactTextString(m) }
func (*CompileResponse) ProtoMessage()               {}
func (*CompileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CompileResponse) GetCompileInfo() *CompileInfo {
	if m != nil {
		return m.CompileInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*CompileOptions)(nil), "protoeasy.CompileOptions")
	proto.RegisterType((*Command)(nil), "protoeasy.Command")
	proto.RegisterType((*CompileInfo)(nil), "protoeasy.CompileInfo")
	proto.RegisterType((*CompileRequest)(nil), "protoeasy.CompileRequest")
	proto.RegisterType((*CompileResponse)(nil), "protoeasy.CompileResponse")
	proto.RegisterEnum("protoeasy.GoPluginType", GoPluginType_name, GoPluginType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error) {
	out := new(CompileResponse)
	err := grpc.Invoke(ctx, "/protoeasy.API/Compile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	Compile(context.Context, *CompileRequest) (*CompileResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Compile(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoeasy.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compile",
			Handler:    _API_Compile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x52, 0x5f, 0x53, 0xd3, 0x4e,
	0x14, 0xa5, 0x7f, 0xf8, 0xd3, 0xdb, 0xfe, 0xda, 0xb0, 0xfc, 0xe9, 0xd2, 0x9f, 0x02, 0x53, 0x5f,
	0x10, 0xb0, 0xcc, 0x84, 0x17, 0xc7, 0xd1, 0x19, 0x11, 0x4b, 0xa6, 0x23, 0xd2, 0x0a, 0xf8, 0xe0,
	0xd3, 0x4e, 0x9a, 0x6e, 0x43, 0x34, 0xcd, 0xae, 0xc9, 0x46, 0x8d, 0x5f, 0xc0, 0x2f, 0xe0, 0xc7,
	0xf2, 0x43, 0xb9, 0xbb, 0x69, 0x22, 0x85, 0xf2, 0x96, 0x7b, 0xce, 0xc9, 0xd9, 0x7b, 0xcf, 0xbd,
	0xd0, 0xe0, 0x21, 0x13, 0x8c, 0xda, 0x51, 0xd2, 0xd1, 0x5f, 0xa8, 0x92, 0x03, 0xad, 0x6d, 0x97,
	0x31, 0xd7, 0xa7, 0x47, 0x1a, 0x19, 0xc6, 0xe3, 0xa3, 0x51, 0x1c, 0xda, 0xc2, 0x63, 0x41, 0x2a,
	0x6d, 0xff, 0x2a, 0x43, 0xfd, 0x94, 0x4d, 0xb8, 0xe7, 0xd3, 0x3e, 0x57, 0x78, 0x84, 0x6a, 0x50,
	0x76, 0x43, 0xee, 0xe0, 0xc2, 0x6e, 0x61, 0x6f, 0x05, 0xad, 0x43, 0x4d, 0x55, 0xc4, 0xb5, 0x05,
	0xfd, 0x6e, 0x27, 0xb8, 0xa8, 0xd1, 0x26, 0x34, 0xe8, 0x0f, 0xc7, 0x8f, 0x47, 0x94, 0x70, 0x5b,
	0x08, 0x1a, 0x06, 0xb8, 0xb4, 0x5b, 0xda, 0xab, 0xa0, 0x2a, 0x94, 0x1c, 0xce, 0xf1, 0xba, 0x56,
	0xad, 0x41, 0x55, 0x16, 0x24, 0xa4, 0x3e, 0x61, 0xb1, 0xc0, 0x1b, 0x12, 0xac, 0xa0, 0x3a, 0x2c,
	0x39, 0xd1, 0x8d, 0x1d, 0x72, 0xbc, 0xad, 0x45, 0x9b, 0x50, 0x4f, 0xeb, 0x5c, 0xb7, 0xa3, 0x75,
	0x00, 0x45, 0x97, 0xe1, 0x3d, 0xad, 0x91, 0x85, 0xcb, 0x72, 0xfe, 0xa9, 0xe6, 0xe5, 0x7f, 0x12,
	0xf3, 0x26, 0x9c, 0x85, 0x42, 0x35, 0x71, 0x83, 0xf7, 0x35, 0xbe, 0x03, 0x4d, 0x89, 0x07, 0x8c,
	0x8c, 0xe8, 0xd8, 0x8e, 0x7d, 0x41, 0x26, 0x6c, 0xe4, 0x8d, 0x3d, 0x1a, 0x46, 0xf8, 0x40, 0x9b,
	0xbd, 0x96, 0x13, 0xb1, 0x5b, 0xe8, 0xa1, 0x6c, 0xbc, 0x6a, 0xee, 0x77, 0xfe, 0xa5, 0x38, 0x1b,
	0x48, 0xc7, 0x62, 0xef, 0x33, 0x71, 0x37, 0x10, 0x61, 0x82, 0x8e, 0xf4, 0xd3, 0xdc, 0x8f, 0x5d,
	0x2f, 0x20, 0x22, 0xe1, 0x14, 0x3f, 0x93, 0xce, 0x75, 0xb3, 0x79, 0xcb, 0xc3, 0x62, 0x03, 0xcd,
	0x5f, 0x4b, 0x5a, 0x45, 0xca, 0x86, 0x9f, 0x1d, 0x6c, 0x66, 0x91, 0xaa, 0x2a, 0x9f, 0xe7, 0x38,
	0xcb, 0x85, 0x27, 0xe2, 0x86, 0x05, 0xf8, 0x65, 0x96, 0x4b, 0x5a, 0xe7, 0xba, 0x57, 0x5a, 0x27,
	0xbd, 0xc2, 0x78, 0x98, 0xe0, 0xb3, 0xcc, 0x4b, 0x55, 0xb9, 0xc6, 0x52, 0x9a, 0x96, 0x09, 0xc6,
	0xbd, 0xa6, 0xe5, 0x66, 0xbe, 0xd0, 0x44, 0x6f, 0xb5, 0x82, 0xfe, 0x83, 0xc5, 0x6f, 0xb6, 0x1f,
	0x53, 0xbd, 0xce, 0xca, 0x8b, 0xe2, 0xf3, 0x42, 0x7b, 0x13, 0x96, 0xe5, 0xdc, 0x13, 0x3b, 0x18,
	0x29, 0xa9, 0x1d, 0xba, 0x52, 0x2a, 0x37, 0xda, 0xfe, 0x5d, 0x80, 0xea, 0x34, 0x90, 0x5e, 0x30,
	0x66, 0xe8, 0x09, 0x2c, 0x3b, 0xa9, 0x4e, 0x0b, 0xaa, 0x26, 0x9a, 0x4d, 0x4e, 0x3b, 0x60, 0x30,
	0xbc, 0x80, 0xc7, 0x82, 0x44, 0xde, 0x4f, 0x4a, 0x86, 0x89, 0xa0, 0x91, 0x7e, 0xaa, 0x8c, 0xb6,
	0x60, 0x55, 0xf6, 0x79, 0x87, 0x2a, 0x69, 0xea, 0x00, 0x56, 0xb2, 0xeb, 0xc4, 0x65, 0x89, 0x54,
	0xcd, 0xad, 0x4e, 0x7a, 0xbe, 0x9d, 0xec, 0x7c, 0x3b, 0x6f, 0xa7, 0x82, 0xf6, 0x87, 0xfc, 0x6e,
	0x2f, 0xe9, 0xd7, 0x98, 0x46, 0x42, 0x75, 0x2d, 0xec, 0x50, 0x0f, 0x58, 0x43, 0x26, 0x34, 0x9c,
	0x94, 0x26, 0x2c, 0x5d, 0xa3, 0x7e, 0x5f, 0x59, 0x3e, 0xb4, 0xe7, 0xf6, 0x39, 0x34, 0x72, 0xcb,
	0x88, 0x4b, 0x84, 0xce, 0x7a, 0x1e, 0x42, 0x2d, 0xf3, 0xf4, 0x64, 0x12, 0x53, 0xc3, 0xcd, 0xfb,
	0x86, 0x2a, 0xa7, 0xfd, 0x3f, 0x05, 0xa8, 0xcd, 0x1c, 0x41, 0x13, 0xd6, 0xac, 0x3e, 0x19, 0x9c,
	0x7f, 0xb4, 0x7a, 0x17, 0xe4, 0xfa, 0xd3, 0xa0, 0x4b, 0x2e, 0xfa, 0x17, 0x5d, 0x63, 0x01, 0x6d,
	0xc0, 0xea, 0x1d, 0xc2, 0xea, 0x1b, 0x05, 0x99, 0xd4, 0xc6, 0x3d, 0xf8, 0xec, 0xe4, 0xea, 0xda,
	0x28, 0xce, 0xb1, 0xb2, 0xfa, 0xf2, 0x9f, 0x12, 0xfa, 0x1f, 0x9a, 0x73, 0x08, 0xfd, 0x57, 0x19,
	0x3d, 0x86, 0xad, 0x07, 0xc8, 0xee, 0xa5, 0xb1, 0x88, 0x1e, 0x01, 0x9e, 0x43, 0x5f, 0x9d, 0xf7,
	0x4e, 0xdf, 0x19, 0x4b, 0x66, 0x0f, 0x4a, 0x27, 0x83, 0x1e, 0x7a, 0xa3, 0xaf, 0x44, 0x0d, 0x89,
	0xe6, 0x24, 0x39, 0x5d, 0x45, 0xab, 0x35, 0x8f, 0x4a, 0x23, 0x6d, 0x2f, 0x0c, 0x97, 0x34, 0x79,
	0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xb4, 0x0d, 0xc8, 0xb8, 0x04, 0x00, 0x00,
}
