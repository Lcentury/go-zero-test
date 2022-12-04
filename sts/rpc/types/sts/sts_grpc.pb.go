// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: sts.proto

package sts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StsClient is the client API for Sts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StsClient interface {
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
}

type stsClient struct {
	cc grpc.ClientConnInterface
}

func NewStsClient(cc grpc.ClientConnInterface) StsClient {
	return &stsClient{cc}
}

func (c *stsClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*InsertResponse, error) {
	out := new(InsertResponse)
	err := c.cc.Invoke(ctx, "/user.Sts/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stsClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/user.Sts/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stsClient) Find(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/user.Sts/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StsServer is the server API for Sts service.
// All implementations must embed UnimplementedStsServer
// for forward compatibility
type StsServer interface {
	Insert(context.Context, *InsertRequest) (*InsertResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Find(context.Context, *FindRequest) (*FindResponse, error)
	mustEmbedUnimplementedStsServer()
}

// UnimplementedStsServer must be embedded to have forward compatible implementations.
type UnimplementedStsServer struct {
}

func (UnimplementedStsServer) Insert(context.Context, *InsertRequest) (*InsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedStsServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStsServer) Find(context.Context, *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedStsServer) mustEmbedUnimplementedStsServer() {}

// UnsafeStsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StsServer will
// result in compilation errors.
type UnsafeStsServer interface {
	mustEmbedUnimplementedStsServer()
}

func RegisterStsServer(s grpc.ServiceRegistrar, srv StsServer) {
	s.RegisterService(&Sts_ServiceDesc, srv)
}

func _Sts_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StsServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Sts/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StsServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Sts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StsServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sts_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StsServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Sts/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StsServer).Find(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sts_ServiceDesc is the grpc.ServiceDesc for Sts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Sts",
	HandlerType: (*StsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _Sts_Insert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Sts_Delete_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _Sts_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sts.proto",
}
