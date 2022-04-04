// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/role/rolepermission.proto

package rpcv3

import (
	context "context"
	v3 "github.com/RafayLabs/rcloud-base/proto/types/commonpb/v3"
	v31 "github.com/RafayLabs/rcloud-base/proto/types/rolepb/v3"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RolepermissionClient is the client API for Rolepermission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolepermissionClient interface {
	GetRolepermissions(ctx context.Context, in *v3.QueryOptions, opts ...grpc.CallOption) (*v31.RolePermissionList, error)
}

type rolepermissionClient struct {
	cc grpc.ClientConnInterface
}

func NewRolepermissionClient(cc grpc.ClientConnInterface) RolepermissionClient {
	return &rolepermissionClient{cc}
}

func (c *rolepermissionClient) GetRolepermissions(ctx context.Context, in *v3.QueryOptions, opts ...grpc.CallOption) (*v31.RolePermissionList, error) {
	out := new(v31.RolePermissionList)
	err := c.cc.Invoke(ctx, "/rafay.dev.rpc.v3.Rolepermission/GetRolepermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolepermissionServer is the server API for Rolepermission service.
// All implementations should embed UnimplementedRolepermissionServer
// for forward compatibility
type RolepermissionServer interface {
	GetRolepermissions(context.Context, *v3.QueryOptions) (*v31.RolePermissionList, error)
}

// UnimplementedRolepermissionServer should be embedded to have forward compatible implementations.
type UnimplementedRolepermissionServer struct {
}

func (UnimplementedRolepermissionServer) GetRolepermissions(context.Context, *v3.QueryOptions) (*v31.RolePermissionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRolepermissions not implemented")
}

// UnsafeRolepermissionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolepermissionServer will
// result in compilation errors.
type UnsafeRolepermissionServer interface {
	mustEmbedUnimplementedRolepermissionServer()
}

func RegisterRolepermissionServer(s grpc.ServiceRegistrar, srv RolepermissionServer) {
	s.RegisterService(&Rolepermission_ServiceDesc, srv)
}

func _Rolepermission_GetRolepermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3.QueryOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolepermissionServer).GetRolepermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.rpc.v3.Rolepermission/GetRolepermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolepermissionServer).GetRolepermissions(ctx, req.(*v3.QueryOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// Rolepermission_ServiceDesc is the grpc.ServiceDesc for Rolepermission service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rolepermission_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rafay.dev.rpc.v3.Rolepermission",
	HandlerType: (*RolepermissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRolepermissions",
			Handler:    _Rolepermission_GetRolepermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/role/rolepermission.proto",
}
