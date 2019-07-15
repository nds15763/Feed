// Code generated by protoc-gen-go. DO NOT EDIT.
// source: material.proto

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

type GetUlrRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUlrRequest) Reset()         { *m = GetUlrRequest{} }
func (m *GetUlrRequest) String() string { return proto.CompactTextString(m) }
func (*GetUlrRequest) ProtoMessage()    {}
func (*GetUlrRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_material_26e685616562b79d, []int{0}
}
func (m *GetUlrRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUlrRequest.Unmarshal(m, b)
}
func (m *GetUlrRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUlrRequest.Marshal(b, m, deterministic)
}
func (dst *GetUlrRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUlrRequest.Merge(dst, src)
}
func (m *GetUlrRequest) XXX_Size() int {
	return xxx_messageInfo_GetUlrRequest.Size(m)
}
func (m *GetUlrRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUlrRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUlrRequest proto.InternalMessageInfo

func (m *GetUlrRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type GetUlrResponse struct {
	Md5                  string   `protobuf:"bytes,1,opt,name=md5" json:"md5,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Mine                 string   `protobuf:"bytes,4,opt,name=mine" json:"mine,omitempty"`
	Phash                uint64   `protobuf:"varint,5,opt,name=phash" json:"phash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUlrResponse) Reset()         { *m = GetUlrResponse{} }
func (m *GetUlrResponse) String() string { return proto.CompactTextString(m) }
func (*GetUlrResponse) ProtoMessage()    {}
func (*GetUlrResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_material_26e685616562b79d, []int{1}
}
func (m *GetUlrResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUlrResponse.Unmarshal(m, b)
}
func (m *GetUlrResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUlrResponse.Marshal(b, m, deterministic)
}
func (dst *GetUlrResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUlrResponse.Merge(dst, src)
}
func (m *GetUlrResponse) XXX_Size() int {
	return xxx_messageInfo_GetUlrResponse.Size(m)
}
func (m *GetUlrResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUlrResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUlrResponse proto.InternalMessageInfo

func (m *GetUlrResponse) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *GetUlrResponse) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *GetUlrResponse) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetUlrResponse) GetMine() string {
	if m != nil {
		return m.Mine
	}
	return ""
}

func (m *GetUlrResponse) GetPhash() uint64 {
	if m != nil {
		return m.Phash
	}
	return 0
}

func init() {
	proto.RegisterType((*GetUlrRequest)(nil), "protos.GetUlrRequest")
	proto.RegisterType((*GetUlrResponse)(nil), "protos.GetUlrResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MaterialService service

type MaterialServiceClient interface {
	GetRealUrl(ctx context.Context, in *GetUlrRequest, opts ...grpc.CallOption) (*GetUlrResponse, error)
}

type materialServiceClient struct {
	cc *grpc.ClientConn
}

func NewMaterialServiceClient(cc *grpc.ClientConn) MaterialServiceClient {
	return &materialServiceClient{cc}
}

func (c *materialServiceClient) GetRealUrl(ctx context.Context, in *GetUlrRequest, opts ...grpc.CallOption) (*GetUlrResponse, error) {
	out := new(GetUlrResponse)
	err := grpc.Invoke(ctx, "/protos.MaterialService/GetRealUrl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MaterialService service

type MaterialServiceServer interface {
	GetRealUrl(context.Context, *GetUlrRequest) (*GetUlrResponse, error)
}

func RegisterMaterialServiceServer(s *grpc.Server, srv MaterialServiceServer) {
	s.RegisterService(&_MaterialService_serviceDesc, srv)
}

func _MaterialService_GetRealUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUlrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).GetRealUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.MaterialService/GetRealUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).GetRealUrl(ctx, req.(*GetUlrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MaterialService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.MaterialService",
	HandlerType: (*MaterialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRealUrl",
			Handler:    _MaterialService_GetRealUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "material.proto",
}

func init() { proto.RegisterFile("material.proto", fileDescriptor_material_26e685616562b79d) }

var fileDescriptor_material_26e685616562b79d = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x4b, 0xc5, 0x30,
	0x10, 0x84, 0x8d, 0xaf, 0x7d, 0xe0, 0x82, 0x4f, 0x59, 0xb4, 0x04, 0x4f, 0xb5, 0xa7, 0x9e, 0x7a,
	0x50, 0x7a, 0xf4, 0xdc, 0x93, 0x20, 0x91, 0xfe, 0x80, 0x68, 0x17, 0x13, 0x48, 0xda, 0x9a, 0xa4,
	0x0a, 0xfe, 0x7a, 0x69, 0xd2, 0x1e, 0xf4, 0xb4, 0x3b, 0xc3, 0x30, 0xcc, 0x07, 0x27, 0x2b, 0x03,
	0x39, 0x2d, 0x4d, 0x33, 0xbb, 0x29, 0x4c, 0x78, 0x8c, 0xc7, 0x57, 0xf7, 0x70, 0xd9, 0x51, 0xe8,
	0x8d, 0x13, 0xf4, 0xb9, 0x90, 0x0f, 0x78, 0x0d, 0x87, 0xc5, 0x19, 0xce, 0x4a, 0x56, 0x5f, 0x88,
	0xf5, 0xad, 0x7e, 0xe0, 0xb4, 0x47, 0xfc, 0x3c, 0x8d, 0x9e, 0xd6, 0x8c, 0x1d, 0xda, 0x3d, 0x63,
	0x87, 0x16, 0x6f, 0x20, 0xff, 0xd6, 0x43, 0x50, 0xfc, 0xbc, 0x64, 0x75, 0x2e, 0x92, 0xc0, 0x02,
	0x8e, 0x8a, 0xf4, 0x87, 0x0a, 0xfc, 0x10, 0xed, 0x4d, 0x21, 0x42, 0x66, 0xf5, 0x48, 0x3c, 0x8b,
	0x05, 0xf1, 0x5f, 0x1b, 0x66, 0x25, 0xbd, 0xe2, 0x79, 0xc9, 0xea, 0x4c, 0x24, 0xf1, 0xf0, 0x02,
	0x57, 0xcf, 0xdb, 0xf0, 0x57, 0x72, 0x5f, 0xfa, 0x9d, 0xf0, 0x09, 0xa0, 0xa3, 0x20, 0x48, 0x9a,
	0xde, 0x19, 0xbc, 0x4d, 0x3c, 0xbe, 0xf9, 0x43, 0x71, 0x57, 0xfc, 0xb7, 0xd3, 0xf2, 0xea, 0xec,
	0x2d, 0x81, 0x3f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x70, 0x3c, 0x51, 0x11, 0x01, 0x00,
	0x00,
}
