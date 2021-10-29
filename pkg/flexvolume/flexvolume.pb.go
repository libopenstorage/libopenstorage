// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/flexvolume/flexvolume.proto

package flexvolume

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import emptypb "google.golang.org/protobuf/types/known/emptypb"

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

type AttachRequest struct {
	JsonOptions          map[string]string `protobuf:"bytes,1,rep,name=json_options,json=jsonOptions" json:"json_options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AttachRequest) Reset()         { *m = AttachRequest{} }
func (m *AttachRequest) String() string { return proto.CompactTextString(m) }
func (*AttachRequest) ProtoMessage()    {}
func (*AttachRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flexvolume_f3351e70cdde2d38, []int{0}
}
func (m *AttachRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachRequest.Unmarshal(m, b)
}
func (m *AttachRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachRequest.Marshal(b, m, deterministic)
}
func (dst *AttachRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachRequest.Merge(dst, src)
}
func (m *AttachRequest) XXX_Size() int {
	return xxx_messageInfo_AttachRequest.Size(m)
}
func (m *AttachRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttachRequest proto.InternalMessageInfo

func (m *AttachRequest) GetJsonOptions() map[string]string {
	if m != nil {
		return m.JsonOptions
	}
	return nil
}

type DetachRequest struct {
	MountDevice          string   `protobuf:"bytes,1,opt,name=mount_device,json=mountDevice" json:"mount_device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetachRequest) Reset()         { *m = DetachRequest{} }
func (m *DetachRequest) String() string { return proto.CompactTextString(m) }
func (*DetachRequest) ProtoMessage()    {}
func (*DetachRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flexvolume_f3351e70cdde2d38, []int{1}
}
func (m *DetachRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetachRequest.Unmarshal(m, b)
}
func (m *DetachRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetachRequest.Marshal(b, m, deterministic)
}
func (dst *DetachRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetachRequest.Merge(dst, src)
}
func (m *DetachRequest) XXX_Size() int {
	return xxx_messageInfo_DetachRequest.Size(m)
}
func (m *DetachRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetachRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetachRequest proto.InternalMessageInfo

func (m *DetachRequest) GetMountDevice() string {
	if m != nil {
		return m.MountDevice
	}
	return ""
}

type MountRequest struct {
	TargetMountDir       string            `protobuf:"bytes,1,opt,name=target_mount_dir,json=targetMountDir" json:"target_mount_dir,omitempty"`
	MountDevice          string            `protobuf:"bytes,2,opt,name=mount_device,json=mountDevice" json:"mount_device,omitempty"`
	JsonOptions          map[string]string `protobuf:"bytes,3,rep,name=json_options,json=jsonOptions" json:"json_options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MountRequest) Reset()         { *m = MountRequest{} }
func (m *MountRequest) String() string { return proto.CompactTextString(m) }
func (*MountRequest) ProtoMessage()    {}
func (*MountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flexvolume_f3351e70cdde2d38, []int{2}
}
func (m *MountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountRequest.Unmarshal(m, b)
}
func (m *MountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountRequest.Marshal(b, m, deterministic)
}
func (dst *MountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountRequest.Merge(dst, src)
}
func (m *MountRequest) XXX_Size() int {
	return xxx_messageInfo_MountRequest.Size(m)
}
func (m *MountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MountRequest proto.InternalMessageInfo

func (m *MountRequest) GetTargetMountDir() string {
	if m != nil {
		return m.TargetMountDir
	}
	return ""
}

func (m *MountRequest) GetMountDevice() string {
	if m != nil {
		return m.MountDevice
	}
	return ""
}

func (m *MountRequest) GetJsonOptions() map[string]string {
	if m != nil {
		return m.JsonOptions
	}
	return nil
}

type UnmountRequest struct {
	MountDir             string   `protobuf:"bytes,1,opt,name=mount_dir,json=mountDir" json:"mount_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnmountRequest) Reset()         { *m = UnmountRequest{} }
func (m *UnmountRequest) String() string { return proto.CompactTextString(m) }
func (*UnmountRequest) ProtoMessage()    {}
func (*UnmountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_flexvolume_f3351e70cdde2d38, []int{3}
}
func (m *UnmountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnmountRequest.Unmarshal(m, b)
}
func (m *UnmountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnmountRequest.Marshal(b, m, deterministic)
}
func (dst *UnmountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnmountRequest.Merge(dst, src)
}
func (m *UnmountRequest) XXX_Size() int {
	return xxx_messageInfo_UnmountRequest.Size(m)
}
func (m *UnmountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnmountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnmountRequest proto.InternalMessageInfo

func (m *UnmountRequest) GetMountDir() string {
	if m != nil {
		return m.MountDir
	}
	return ""
}

func init() {
	proto.RegisterType((*AttachRequest)(nil), "flexvolume.AttachRequest")
	proto.RegisterMapType((map[string]string)(nil), "flexvolume.AttachRequest.JsonOptionsEntry")
	proto.RegisterType((*DetachRequest)(nil), "flexvolume.DetachRequest")
	proto.RegisterType((*MountRequest)(nil), "flexvolume.MountRequest")
	proto.RegisterMapType((map[string]string)(nil), "flexvolume.MountRequest.JsonOptionsEntry")
	proto.RegisterType((*UnmountRequest)(nil), "flexvolume.UnmountRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	Init(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Attach(ctx context.Context, in *AttachRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Detach(ctx context.Context, in *DetachRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Mount(ctx context.Context, in *MountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Unmount(ctx context.Context, in *UnmountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Init(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := grpc.Invoke(ctx, "/flexvolume.API/Init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Attach(ctx context.Context, in *AttachRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := grpc.Invoke(ctx, "/flexvolume.API/Attach", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Detach(ctx context.Context, in *DetachRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := grpc.Invoke(ctx, "/flexvolume.API/Detach", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Mount(ctx context.Context, in *MountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := grpc.Invoke(ctx, "/flexvolume.API/Mount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Unmount(ctx context.Context, in *UnmountRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := grpc.Invoke(ctx, "/flexvolume.API/Unmount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	Init(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Attach(context.Context, *AttachRequest) (*emptypb.Empty, error)
	Detach(context.Context, *DetachRequest) (*emptypb.Empty, error)
	Mount(context.Context, *MountRequest) (*emptypb.Empty, error)
	Unmount(context.Context, *UnmountRequest) (*emptypb.Empty, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flexvolume.API/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Init(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Attach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Attach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flexvolume.API/Attach",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Attach(ctx, req.(*AttachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Detach_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Detach(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flexvolume.API/Detach",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Detach(ctx, req.(*DetachRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Mount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Mount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flexvolume.API/Mount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Mount(ctx, req.(*MountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Unmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Unmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flexvolume.API/Unmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Unmount(ctx, req.(*UnmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flexvolume.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _API_Init_Handler,
		},
		{
			MethodName: "Attach",
			Handler:    _API_Attach_Handler,
		},
		{
			MethodName: "Detach",
			Handler:    _API_Detach_Handler,
		},
		{
			MethodName: "Mount",
			Handler:    _API_Mount_Handler,
		},
		{
			MethodName: "Unmount",
			Handler:    _API_Unmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/flexvolume/flexvolume.proto",
}

func init() {
	proto.RegisterFile("pkg/flexvolume/flexvolume.proto", fileDescriptor_flexvolume_f3351e70cdde2d38)
}

var fileDescriptor_flexvolume_f3351e70cdde2d38 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xb1, 0xbd, 0xa4, 0xed, 0xb3, 0x53, 0x3c, 0x6d, 0x0c, 0xcf, 0x1d, 0xac, 0xd3, 0x29,
	0x2b, 0xcc, 0x82, 0xec, 0x32, 0x7a, 0x18, 0x14, 0x52, 0x58, 0x47, 0x4b, 0x87, 0x61, 0xe7, 0xe0,
	0xb6, 0x8a, 0xa7, 0xc4, 0x96, 0x3c, 0x5b, 0x0e, 0xcb, 0x75, 0x5f, 0x21, 0xf7, 0x7d, 0xa9, 0x7d,
	0x85, 0x5d, 0xf6, 0x2d, 0x86, 0x25, 0x9b, 0xd8, 0x09, 0x86, 0x5d, 0x7a, 0xb3, 0xfe, 0x7a, 0xef,
	0x67, 0xbd, 0xff, 0xfb, 0xc3, 0xeb, 0x6c, 0x19, 0x93, 0x79, 0x42, 0x7f, 0xac, 0x44, 0x52, 0xa6,
	0xb4, 0xf5, 0x19, 0x64, 0xb9, 0x90, 0x02, 0xc1, 0x56, 0xf1, 0x5f, 0xc5, 0x42, 0xc4, 0x09, 0x25,
	0x51, 0xc6, 0x48, 0xc4, 0xb9, 0x90, 0x91, 0x64, 0x82, 0x17, 0xba, 0xd2, 0x3f, 0xa9, 0x6f, 0xd5,
	0xe9, 0xae, 0x9c, 0x13, 0x9a, 0x66, 0x72, 0xad, 0x2f, 0xf1, 0x2f, 0x03, 0x46, 0x17, 0x52, 0x46,
	0xf7, 0xdf, 0x42, 0xfa, 0xbd, 0xa4, 0x85, 0x44, 0x37, 0xe0, 0x2c, 0x0a, 0xc1, 0x67, 0x22, 0x53,
	0x10, 0xcf, 0x38, 0xb5, 0xc6, 0xf6, 0xe4, 0x2c, 0x68, 0xbd, 0xa0, 0xd3, 0x10, 0x7c, 0x2e, 0x04,
	0xbf, 0xd5, 0xc5, 0x97, 0x5c, 0xe6, 0xeb, 0xd0, 0x5e, 0x6c, 0x15, 0xff, 0x23, 0xb8, 0xbb, 0x05,
	0xc8, 0x05, 0x6b, 0x49, 0xd7, 0x9e, 0x71, 0x6a, 0x8c, 0x8f, 0xc2, 0xea, 0x13, 0x3d, 0x87, 0xc1,
	0x2a, 0x4a, 0x4a, 0xea, 0x99, 0x4a, 0xd3, 0x87, 0x73, 0xf3, 0x83, 0x81, 0x27, 0x30, 0x9a, 0xd2,
	0xf6, 0xfb, 0xde, 0x80, 0x93, 0x8a, 0x92, 0xcb, 0xd9, 0x03, 0x5d, 0xb1, 0x7b, 0x5a, 0x53, 0x6c,
	0xa5, 0x4d, 0x95, 0x84, 0xff, 0x1a, 0xe0, 0xdc, 0x54, 0xe7, 0xa6, 0x67, 0x0c, 0xae, 0x8c, 0xf2,
	0x98, 0xca, 0x59, 0xdd, 0xca, 0xf2, 0xba, 0xef, 0x58, 0xeb, 0xaa, 0x7a, 0xca, 0xf2, 0x3d, 0xba,
	0xb9, 0x47, 0x47, 0xd7, 0x3b, 0x06, 0x59, 0xca, 0xa0, 0xb7, 0x6d, 0x83, 0xda, 0x3f, 0x7f, 0x64,
	0x7f, 0xde, 0xc1, 0xf1, 0x57, 0x9e, 0xb6, 0x87, 0x3d, 0x81, 0xa3, 0xdd, 0x29, 0x0f, 0xd3, 0x7a,
	0xbe, 0xc9, 0xc6, 0x02, 0xeb, 0xe2, 0xcb, 0x15, 0xfa, 0x04, 0x4f, 0xae, 0x38, 0x93, 0xe8, 0x45,
	0xa0, 0xd3, 0x11, 0x34, 0xe9, 0x08, 0x2e, 0xab, 0x74, 0xf8, 0x3d, 0x3a, 0x76, 0x7f, 0xfe, 0xfe,
	0xb3, 0x31, 0x01, 0x0f, 0x08, 0xe3, 0x4c, 0x9e, 0x1b, 0x67, 0xe8, 0x16, 0x86, 0x3a, 0x0f, 0xe8,
	0x65, 0x6f, 0x46, 0x7a, 0x71, 0x48, 0xe1, 0x1c, 0x7c, 0x40, 0x22, 0x55, 0x5f, 0x03, 0xf5, 0xc6,
	0xbb, 0xc0, 0x4e, 0x0a, 0xfe, 0x03, 0xf8, 0x40, 0x1b, 0xe0, 0x35, 0x0c, 0xd4, 0x42, 0x90, 0xd7,
	0xb7, 0xa3, 0x5e, 0xdc, 0x53, 0x85, 0xb3, 0xf1, 0x90, 0x28, 0x07, 0x2b, 0x5a, 0x08, 0x07, 0xb5,
	0xe1, 0xc8, 0x6f, 0xf3, 0xba, 0x5b, 0xe8, 0x25, 0x3e, 0x53, 0xc4, 0x11, 0x3e, 0x24, 0x25, 0x6f,
	0x98, 0x77, 0x43, 0x55, 0xf4, 0xfe, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x30, 0xd3, 0x1a,
	0xf6, 0x03, 0x00, 0x00,
}
