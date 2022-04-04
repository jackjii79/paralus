// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/rpc/sentry/audit_info.proto

package sentry

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

// AuditInformationClient is the client API for AuditInformation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditInformationClient interface {
	LookupUser(ctx context.Context, in *LookupUserRequest, opts ...grpc.CallOption) (*LookupUserResponse, error)
	LookupCluster(ctx context.Context, in *LookupClusterRequest, opts ...grpc.CallOption) (*LookupClusterResponse, error)
}

type auditInformationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditInformationClient(cc grpc.ClientConnInterface) AuditInformationClient {
	return &auditInformationClient{cc}
}

func (c *auditInformationClient) LookupUser(ctx context.Context, in *LookupUserRequest, opts ...grpc.CallOption) (*LookupUserResponse, error) {
	out := new(LookupUserResponse)
	err := c.cc.Invoke(ctx, "/rafay.dev.sentry.rpc.AuditInformation/LookupUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditInformationClient) LookupCluster(ctx context.Context, in *LookupClusterRequest, opts ...grpc.CallOption) (*LookupClusterResponse, error) {
	out := new(LookupClusterResponse)
	err := c.cc.Invoke(ctx, "/rafay.dev.sentry.rpc.AuditInformation/LookupCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditInformationServer is the server API for AuditInformation service.
// All implementations should embed UnimplementedAuditInformationServer
// for forward compatibility
type AuditInformationServer interface {
	LookupUser(context.Context, *LookupUserRequest) (*LookupUserResponse, error)
	LookupCluster(context.Context, *LookupClusterRequest) (*LookupClusterResponse, error)
}

// UnimplementedAuditInformationServer should be embedded to have forward compatible implementations.
type UnimplementedAuditInformationServer struct {
}

func (UnimplementedAuditInformationServer) LookupUser(context.Context, *LookupUserRequest) (*LookupUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupUser not implemented")
}
func (UnimplementedAuditInformationServer) LookupCluster(context.Context, *LookupClusterRequest) (*LookupClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupCluster not implemented")
}

// UnsafeAuditInformationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditInformationServer will
// result in compilation errors.
type UnsafeAuditInformationServer interface {
	mustEmbedUnimplementedAuditInformationServer()
}

func RegisterAuditInformationServer(s grpc.ServiceRegistrar, srv AuditInformationServer) {
	s.RegisterService(&AuditInformation_ServiceDesc, srv)
}

func _AuditInformation_LookupUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditInformationServer).LookupUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.sentry.rpc.AuditInformation/LookupUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditInformationServer).LookupUser(ctx, req.(*LookupUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuditInformation_LookupCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditInformationServer).LookupCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rafay.dev.sentry.rpc.AuditInformation/LookupCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditInformationServer).LookupCluster(ctx, req.(*LookupClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuditInformation_ServiceDesc is the grpc.ServiceDesc for AuditInformation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuditInformation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rafay.dev.sentry.rpc.AuditInformation",
	HandlerType: (*AuditInformationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookupUser",
			Handler:    _AuditInformation_LookupUser_Handler,
		},
		{
			MethodName: "LookupCluster",
			Handler:    _AuditInformation_LookupCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rpc/sentry/audit_info.proto",
}
