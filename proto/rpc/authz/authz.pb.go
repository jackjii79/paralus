// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/rpc/authz/authz.proto

package authzv1

import (
	authz "github.com/RafayLabs/rcloud-base/proto/types/authz"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_rpc_authz_authz_proto protoreflect.FileDescriptor

var file_proto_rpc_authz_authz_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9e, 0x08, 0x0a, 0x05, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x12, 0x5a,
	0x0a, 0x07, 0x45, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x72, 0x61, 0x66, 0x61,
	0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x61, 0x66,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x1a, 0x22, 0x2e, 0x72,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x1a, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x59, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x20, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x1a, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x23, 0x2e, 0x72,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x1a, 0x24, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x24, 0x2e,
	0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x1a, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x23,
	0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x1a, 0x23, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x1a, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x37, 0x2e, 0x72, 0x61, 0x66, 0x61,
	0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x52, 0x6f, 0x6c,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x1a, 0x33, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x7a, 0x0a, 0x1c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x33, 0x2e, 0x72, 0x61, 0x66, 0x61,
	0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x7e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x37, 0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x23,
	0x2e, 0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0xdf, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61,
	0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52,
	0x61, 0x66, 0x61, 0x79, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x52,
	0x44, 0x52, 0x41, 0xaa, 0x02, 0x16, 0x52, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x2e,
	0x52, 0x70, 0x63, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x16, 0x52,
	0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x41, 0x75, 0x74,
	0x68, 0x7a, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65,
	0x76, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x52, 0x61, 0x66,
	0x61, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x76, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x41, 0x75,
	0x74, 0x68, 0x7a, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_rpc_authz_authz_proto_goTypes = []interface{}{
	(*authz.EnforceRequest)(nil),                // 0: rafay.dev.types.authz.v1.EnforceRequest
	(*authz.Policy)(nil),                        // 1: rafay.dev.types.authz.v1.Policy
	(*authz.Policies)(nil),                      // 2: rafay.dev.types.authz.v1.Policies
	(*authz.UserGroup)(nil),                     // 3: rafay.dev.types.authz.v1.UserGroup
	(*authz.UserGroups)(nil),                    // 4: rafay.dev.types.authz.v1.UserGroups
	(*authz.FilteredRolePermissionMapping)(nil), // 5: rafay.dev.types.authz.v1.FilteredRolePermissionMapping
	(*authz.RolePermissionMappingList)(nil),     // 6: rafay.dev.types.authz.v1.RolePermissionMappingList
	(*authz.BoolReply)(nil),                     // 7: rafay.dev.types.authz.v1.BoolReply
}
var file_proto_rpc_authz_authz_proto_depIdxs = []int32{
	0,  // 0: rafay.dev.rpc.authz.v1.Authz.Enforce:input_type -> rafay.dev.types.authz.v1.EnforceRequest
	1,  // 1: rafay.dev.rpc.authz.v1.Authz.ListPolicies:input_type -> rafay.dev.types.authz.v1.Policy
	2,  // 2: rafay.dev.rpc.authz.v1.Authz.CreatePolicies:input_type -> rafay.dev.types.authz.v1.Policies
	1,  // 3: rafay.dev.rpc.authz.v1.Authz.DeletePolicies:input_type -> rafay.dev.types.authz.v1.Policy
	3,  // 4: rafay.dev.rpc.authz.v1.Authz.ListUserGroups:input_type -> rafay.dev.types.authz.v1.UserGroup
	4,  // 5: rafay.dev.rpc.authz.v1.Authz.CreateUserGroups:input_type -> rafay.dev.types.authz.v1.UserGroups
	3,  // 6: rafay.dev.rpc.authz.v1.Authz.DeleteUserGroups:input_type -> rafay.dev.types.authz.v1.UserGroup
	5,  // 7: rafay.dev.rpc.authz.v1.Authz.ListRolePermissionMappings:input_type -> rafay.dev.types.authz.v1.FilteredRolePermissionMapping
	6,  // 8: rafay.dev.rpc.authz.v1.Authz.CreateRolePermissionMappings:input_type -> rafay.dev.types.authz.v1.RolePermissionMappingList
	5,  // 9: rafay.dev.rpc.authz.v1.Authz.DeleteRolePermissionMappings:input_type -> rafay.dev.types.authz.v1.FilteredRolePermissionMapping
	7,  // 10: rafay.dev.rpc.authz.v1.Authz.Enforce:output_type -> rafay.dev.types.authz.v1.BoolReply
	2,  // 11: rafay.dev.rpc.authz.v1.Authz.ListPolicies:output_type -> rafay.dev.types.authz.v1.Policies
	7,  // 12: rafay.dev.rpc.authz.v1.Authz.CreatePolicies:output_type -> rafay.dev.types.authz.v1.BoolReply
	7,  // 13: rafay.dev.rpc.authz.v1.Authz.DeletePolicies:output_type -> rafay.dev.types.authz.v1.BoolReply
	4,  // 14: rafay.dev.rpc.authz.v1.Authz.ListUserGroups:output_type -> rafay.dev.types.authz.v1.UserGroups
	7,  // 15: rafay.dev.rpc.authz.v1.Authz.CreateUserGroups:output_type -> rafay.dev.types.authz.v1.BoolReply
	7,  // 16: rafay.dev.rpc.authz.v1.Authz.DeleteUserGroups:output_type -> rafay.dev.types.authz.v1.BoolReply
	6,  // 17: rafay.dev.rpc.authz.v1.Authz.ListRolePermissionMappings:output_type -> rafay.dev.types.authz.v1.RolePermissionMappingList
	7,  // 18: rafay.dev.rpc.authz.v1.Authz.CreateRolePermissionMappings:output_type -> rafay.dev.types.authz.v1.BoolReply
	7,  // 19: rafay.dev.rpc.authz.v1.Authz.DeleteRolePermissionMappings:output_type -> rafay.dev.types.authz.v1.BoolReply
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rpc_authz_authz_proto_init() }
func file_proto_rpc_authz_authz_proto_init() {
	if File_proto_rpc_authz_authz_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_rpc_authz_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_authz_authz_proto_goTypes,
		DependencyIndexes: file_proto_rpc_authz_authz_proto_depIdxs,
	}.Build()
	File_proto_rpc_authz_authz_proto = out.File
	file_proto_rpc_authz_authz_proto_rawDesc = nil
	file_proto_rpc_authz_authz_proto_goTypes = nil
	file_proto_rpc_authz_authz_proto_depIdxs = nil
}
