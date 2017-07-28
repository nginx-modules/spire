// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/node_attestor/proto/node_attestor.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	plugins/node_attestor/proto/node_attestor.proto

It has these top-level messages:
	FetchAttestationDataRequest
	FetchAttestationDataResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto2 "github.com/spiffe/node-agent/plugins/common/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// ConfigureRequest from public import github.com/spiffe/node-agent/plugins/common/proto/common.proto
type ConfigureRequest proto2.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*proto2.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*proto2.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*proto2.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/node-agent/plugins/common/proto/common.proto
type ConfigureResponse proto2.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*proto2.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*proto2.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}

// GetPluginInfoRequest from public import github.com/spiffe/node-agent/plugins/common/proto/common.proto
type GetPluginInfoRequest proto2.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*proto2.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*proto2.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/node-agent/plugins/common/proto/common.proto
type GetPluginInfoResponse proto2.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset()         { (*proto2.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string { return (*proto2.GetPluginInfoResponse)(m).String() }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (m *GetPluginInfoResponse) GetPluginName() string {
	return (*proto2.GetPluginInfoResponse)(m).GetPluginName()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*proto2.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*proto2.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*proto2.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*proto2.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*proto2.GetPluginInfoResponse)(m).GetCompany()
}

type FetchAttestationDataRequest struct {
}

func (m *FetchAttestationDataRequest) Reset()                    { *m = FetchAttestationDataRequest{} }
func (m *FetchAttestationDataRequest) String() string            { return proto1.CompactTextString(m) }
func (*FetchAttestationDataRequest) ProtoMessage()               {}
func (*FetchAttestationDataRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FetchAttestationDataResponse struct {
	AttestationData []byte `protobuf:"bytes,1,opt,name=attestationData,proto3" json:"attestationData,omitempty"`
}

func (m *FetchAttestationDataResponse) Reset()                    { *m = FetchAttestationDataResponse{} }
func (m *FetchAttestationDataResponse) String() string            { return proto1.CompactTextString(m) }
func (*FetchAttestationDataResponse) ProtoMessage()               {}
func (*FetchAttestationDataResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FetchAttestationDataResponse) GetAttestationData() []byte {
	if m != nil {
		return m.AttestationData
	}
	return nil
}

func init() {
	proto1.RegisterType((*FetchAttestationDataRequest)(nil), "proto.FetchAttestationDataRequest")
	proto1.RegisterType((*FetchAttestationDataResponse)(nil), "proto.FetchAttestationDataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeAttestor service

type NodeAttestorClient interface {
	FetchAttestationData(ctx context.Context, in *FetchAttestationDataRequest, opts ...grpc.CallOption) (*FetchAttestationDataResponse, error)
	Configure(ctx context.Context, in *proto2.ConfigureRequest, opts ...grpc.CallOption) (*proto2.ConfigureResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) FetchAttestationData(ctx context.Context, in *FetchAttestationDataRequest, opts ...grpc.CallOption) (*FetchAttestationDataResponse, error) {
	out := new(FetchAttestationDataResponse)
	err := grpc.Invoke(ctx, "/proto.NodeAttestor/FetchAttestationData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *proto2.ConfigureRequest, opts ...grpc.CallOption) (*proto2.ConfigureResponse, error) {
	out := new(proto2.ConfigureResponse)
	err := grpc.Invoke(ctx, "/proto.NodeAttestor/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeAttestor service

type NodeAttestorServer interface {
	FetchAttestationData(context.Context, *FetchAttestationDataRequest) (*FetchAttestationDataResponse, error)
	Configure(context.Context, *proto2.ConfigureRequest) (*proto2.ConfigureResponse, error)
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_FetchAttestationData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAttestationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).FetchAttestationData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeAttestor/FetchAttestationData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).FetchAttestationData(ctx, req.(*FetchAttestationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto2.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*proto2.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAttestationData",
			Handler:    _NodeAttestor_FetchAttestationData_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugins/node_attestor/proto/node_attestor.proto",
}

func init() { proto1.RegisterFile("plugins/node_attestor/proto/node_attestor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xed, 0x41, 0xc1, 0xb0, 0x20, 0x04, 0xc1, 0x65, 0x55, 0x90, 0x7a, 0xd9, 0x8b, 0x0d,
	0xe8, 0x7d, 0x61, 0x51, 0xc4, 0x93, 0x48, 0x5f, 0xa0, 0xa4, 0xed, 0x34, 0x0d, 0xd8, 0x4c, 0x6c,
	0x26, 0xcf, 0xe5, 0x2b, 0x0a, 0x49, 0x2c, 0x54, 0x4b, 0x4f, 0x21, 0xff, 0x37, 0xf3, 0xf1, 0x0f,
	0x13, 0xf6, 0xd3, 0x2b, 0x6d, 0x9c, 0x30, 0xd8, 0x42, 0x25, 0x89, 0xc0, 0x11, 0x8e, 0xc2, 0x8e,
	0x48, 0x38, 0xcf, 0x8a, 0x90, 0xf1, 0xd3, 0xf0, 0xec, 0x0e, 0x4a, 0x53, 0xef, 0xeb, 0xa2, 0xc1,
	0x41, 0x38, 0xab, 0xbb, 0x0e, 0xc2, 0xf4, 0x83, 0x54, 0x60, 0x68, 0x92, 0x36, 0x38, 0x0c, 0x68,
	0x92, 0x2d, 0x7e, 0xa2, 0x26, 0xbf, 0x65, 0xd7, 0xaf, 0x40, 0x4d, 0x7f, 0x0c, 0x76, 0x49, 0x1a,
	0xcd, 0x8b, 0x24, 0x59, 0xc2, 0x97, 0x07, 0x47, 0xf9, 0x1b, 0xbb, 0x59, 0xc6, 0xce, 0xa2, 0x71,
	0xc0, 0xf7, 0xec, 0x42, 0xce, 0xd1, 0x36, 0xbb, 0xcb, 0xf6, 0x9b, 0xf2, 0x6f, 0xfc, 0xf8, 0x9d,
	0xb1, 0xcd, 0x3b, 0xb6, 0x70, 0x4c, 0x67, 0xf0, 0x8a, 0x5d, 0x2e, 0xa9, 0x79, 0x1e, 0x9b, 0x15,
	0x2b, 0xb5, 0x76, 0xf7, 0xab, 0x33, 0xa9, 0xdb, 0x81, 0x9d, 0x3f, 0xa3, 0xe9, 0xb4, 0xf2, 0x23,
	0xf0, 0xab, 0xb4, 0x31, 0x25, 0xbf, 0xaa, 0xed, 0x7f, 0x10, 0xf7, 0x3f, 0x4e, 0xea, 0xb3, 0x80,
	0x9e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xe5, 0x0e, 0x26, 0x9d, 0x01, 0x00, 0x00,
}
