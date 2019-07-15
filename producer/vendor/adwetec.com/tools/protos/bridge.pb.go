// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bridge.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRealUlrRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRealUlrRequest) Reset()         { *m = GetRealUlrRequest{} }
func (m *GetRealUlrRequest) String() string { return proto.CompactTextString(m) }
func (*GetRealUlrRequest) ProtoMessage()    {}
func (*GetRealUlrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bridge_c426662c903519eb, []int{0}
}
func (m *GetRealUlrRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRealUlrRequest.Unmarshal(m, b)
}
func (m *GetRealUlrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRealUlrRequest.Marshal(b, m, deterministic)
}
func (dst *GetRealUlrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRealUlrRequest.Merge(dst, src)
}
func (m *GetRealUlrRequest) XXX_Size() int {
	return xxx_messageInfo_GetRealUlrRequest.Size(m)
}
func (m *GetRealUlrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRealUlrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRealUlrRequest proto.InternalMessageInfo

func (m *GetRealUlrRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type GetRealUlrResponse struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRealUlrResponse) Reset()         { *m = GetRealUlrResponse{} }
func (m *GetRealUlrResponse) String() string { return proto.CompactTextString(m) }
func (*GetRealUlrResponse) ProtoMessage()    {}
func (*GetRealUlrResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bridge_c426662c903519eb, []int{1}
}
func (m *GetRealUlrResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRealUlrResponse.Unmarshal(m, b)
}
func (m *GetRealUlrResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRealUlrResponse.Marshal(b, m, deterministic)
}
func (dst *GetRealUlrResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRealUlrResponse.Merge(dst, src)
}
func (m *GetRealUlrResponse) XXX_Size() int {
	return xxx_messageInfo_GetRealUlrResponse.Size(m)
}
func (m *GetRealUlrResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRealUlrResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRealUlrResponse proto.InternalMessageInfo

func (m *GetRealUlrResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRealUlrRequest)(nil), "protos.GetRealUlrRequest")
	proto.RegisterType((*GetRealUlrResponse)(nil), "protos.GetRealUlrResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BridgeService service

type BridgeServiceClient interface {
	GetRealUrl(ctx context.Context, in *GetRealUlrRequest, opts ...grpc.CallOption) (*GetRealUlrResponse, error)
}

type bridgeServiceClient struct {
	cc *grpc.ClientConn
}

func NewBridgeServiceClient(cc *grpc.ClientConn) BridgeServiceClient {
	return &bridgeServiceClient{cc}
}

func (c *bridgeServiceClient) GetRealUrl(ctx context.Context, in *GetRealUlrRequest, opts ...grpc.CallOption) (*GetRealUlrResponse, error) {
	out := new(GetRealUlrResponse)
	err := grpc.Invoke(ctx, "/protos.BridgeService/GetRealUrl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BridgeService service

type BridgeServiceServer interface {
	GetRealUrl(context.Context, *GetRealUlrRequest) (*GetRealUlrResponse, error)
}

func RegisterBridgeServiceServer(s *grpc.Server, srv BridgeServiceServer) {
	s.RegisterService(&_BridgeService_serviceDesc, srv)
}

func _BridgeService_GetRealUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRealUlrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServiceServer).GetRealUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.BridgeService/GetRealUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServiceServer).GetRealUrl(ctx, req.(*GetRealUlrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BridgeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.BridgeService",
	HandlerType: (*BridgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRealUrl",
			Handler:    _BridgeService_GetRealUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bridge.proto",
}

func init() { proto.RegisterFile("bridge.proto", fileDescriptor_bridge_c426662c903519eb) }

var fileDescriptor_bridge_c426662c903519eb = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2a, 0xca, 0x4c,
	0x49, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xaa, 0x5c,
	0x82, 0xee, 0xa9, 0x25, 0x41, 0xa9, 0x89, 0x39, 0xa1, 0x39, 0x45, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc, 0xa5, 0x45, 0x39, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x20, 0xa6, 0x92, 0x1a, 0x97, 0x10, 0xb2, 0xb2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x4c, 0x75,
	0x46, 0x61, 0x5c, 0xbc, 0x4e, 0x60, 0x6b, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0x5c,
	0xb9, 0xb8, 0x60, 0x1a, 0x8b, 0x72, 0x84, 0x24, 0x21, 0xb6, 0x17, 0xeb, 0x61, 0xd8, 0x29, 0x25,
	0x85, 0x4d, 0x0a, 0x62, 0x8f, 0x12, 0x43, 0x12, 0xc4, 0xb9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xee, 0xda, 0x5a, 0x93, 0xc5, 0x00, 0x00, 0x00,
}