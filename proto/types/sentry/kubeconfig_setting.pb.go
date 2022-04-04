// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/types/sentry/kubeconfig_setting.proto

package sentry

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KubeconfigRevocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationID string                 `protobuf:"bytes,2,opt,name=organizationID,proto3" json:"organizationID,omitempty"`
	PartnerID      string                 `protobuf:"bytes,3,opt,name=partnerID,proto3" json:"partnerID,omitempty"`
	AccountID      string                 `protobuf:"bytes,4,opt,name=accountID,proto3" json:"accountID,omitempty"`
	IsSSOUser      bool                   `protobuf:"varint,5,opt,name=isSSOUser,proto3" json:"isSSOUser,omitempty"`
	RevokedAt      *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=revokedAt,proto3" json:"revokedAt,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *KubeconfigRevocation) Reset() {
	*x = KubeconfigRevocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_sentry_kubeconfig_setting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeconfigRevocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeconfigRevocation) ProtoMessage() {}

func (x *KubeconfigRevocation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_sentry_kubeconfig_setting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeconfigRevocation.ProtoReflect.Descriptor instead.
func (*KubeconfigRevocation) Descriptor() ([]byte, []int) {
	return file_proto_types_sentry_kubeconfig_setting_proto_rawDescGZIP(), []int{0}
}

func (x *KubeconfigRevocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KubeconfigRevocation) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

func (x *KubeconfigRevocation) GetPartnerID() string {
	if x != nil {
		return x.PartnerID
	}
	return ""
}

func (x *KubeconfigRevocation) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *KubeconfigRevocation) GetIsSSOUser() bool {
	if x != nil {
		return x.IsSSOUser
	}
	return false
}

func (x *KubeconfigRevocation) GetRevokedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RevokedAt
	}
	return nil
}

func (x *KubeconfigRevocation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type KubeconfigSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationID              string                 `protobuf:"bytes,2,opt,name=organizationID,proto3" json:"organizationID,omitempty"`
	PartnerID                   string                 `protobuf:"bytes,3,opt,name=partnerID,proto3" json:"partnerID,omitempty"`
	AccountID                   string                 `protobuf:"bytes,4,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Scope                       string                 `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
	ValiditySeconds             int64                  `protobuf:"zigzag64,6,opt,name=validitySeconds,proto3" json:"validitySeconds,omitempty"`
	CreatedAt                   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ModifiedAt                  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=modifiedAt,proto3" json:"modifiedAt,omitempty"`
	DeletedAt                   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	EnableSessionCheck          bool                   `protobuf:"varint,10,opt,name=enableSessionCheck,proto3" json:"enableSessionCheck,omitempty"`
	IsSSOUser                   bool                   `protobuf:"varint,11,opt,name=isSSOUser,proto3" json:"isSSOUser,omitempty"`
	EnablePrivateRelay          bool                   `protobuf:"varint,12,opt,name=enablePrivateRelay,proto3" json:"enablePrivateRelay,omitempty"`
	EnforceOrgAdminSecretAccess bool                   `protobuf:"varint,13,opt,name=enforceOrgAdminSecretAccess,proto3" json:"enforceOrgAdminSecretAccess,omitempty"`
	DisableWebKubectl           bool                   `protobuf:"varint,14,opt,name=disableWebKubectl,proto3" json:"disableWebKubectl,omitempty"`
	DisableCLIKubectl           bool                   `protobuf:"varint,15,opt,name=disableCLIKubectl,proto3" json:"disableCLIKubectl,omitempty"`
}

func (x *KubeconfigSetting) Reset() {
	*x = KubeconfigSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_types_sentry_kubeconfig_setting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubeconfigSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeconfigSetting) ProtoMessage() {}

func (x *KubeconfigSetting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_types_sentry_kubeconfig_setting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeconfigSetting.ProtoReflect.Descriptor instead.
func (*KubeconfigSetting) Descriptor() ([]byte, []int) {
	return file_proto_types_sentry_kubeconfig_setting_proto_rawDescGZIP(), []int{1}
}

func (x *KubeconfigSetting) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KubeconfigSetting) GetOrganizationID() string {
	if x != nil {
		return x.OrganizationID
	}
	return ""
}

func (x *KubeconfigSetting) GetPartnerID() string {
	if x != nil {
		return x.PartnerID
	}
	return ""
}

func (x *KubeconfigSetting) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *KubeconfigSetting) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *KubeconfigSetting) GetValiditySeconds() int64 {
	if x != nil {
		return x.ValiditySeconds
	}
	return 0
}

func (x *KubeconfigSetting) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *KubeconfigSetting) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

func (x *KubeconfigSetting) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *KubeconfigSetting) GetEnableSessionCheck() bool {
	if x != nil {
		return x.EnableSessionCheck
	}
	return false
}

func (x *KubeconfigSetting) GetIsSSOUser() bool {
	if x != nil {
		return x.IsSSOUser
	}
	return false
}

func (x *KubeconfigSetting) GetEnablePrivateRelay() bool {
	if x != nil {
		return x.EnablePrivateRelay
	}
	return false
}

func (x *KubeconfigSetting) GetEnforceOrgAdminSecretAccess() bool {
	if x != nil {
		return x.EnforceOrgAdminSecretAccess
	}
	return false
}

func (x *KubeconfigSetting) GetDisableWebKubectl() bool {
	if x != nil {
		return x.DisableWebKubectl
	}
	return false
}

func (x *KubeconfigSetting) GetDisableCLIKubectl() bool {
	if x != nil {
		return x.DisableCLIKubectl
	}
	return false
}

var File_proto_types_sentry_kubeconfig_setting_proto protoreflect.FileDescriptor

var file_proto_types_sentry_kubeconfig_setting_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x14, 0x4b, 0x75, 0x62, 0x65, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x53, 0x4f, 0x55, 0x73, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x53, 0x4f, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x38, 0x0a, 0x09, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x92, 0x41, 0x02, 0x40,
	0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa8, 0x05, 0x0a,
	0x11, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x92, 0x41, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x92, 0x41, 0x02, 0x40, 0x01, 0x52, 0x0a,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x05, 0x92, 0x41, 0x02, 0x40, 0x01,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x73, 0x53, 0x53, 0x4f, 0x55, 0x73, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x53, 0x53, 0x4f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x40, 0x0a, 0x1b, 0x65, 0x6e, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x4f, 0x72, 0x67, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b,
	0x65, 0x6e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x4f, 0x72, 0x67, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x57, 0x65, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x74, 0x6c,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x57,
	0x65, 0x62, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x74, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x43, 0x4c, 0x49, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x74, 0x6c, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x4c, 0x49,
	0x4b, 0x75, 0x62, 0x65, 0x63, 0x74, 0x6c, 0x42, 0xe5, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x61, 0x66, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x73, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x16, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x66,
	0x61, 0x79, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x72, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0xa2, 0x02, 0x04, 0x52, 0x44, 0x54, 0x53, 0xaa, 0x02, 0x16, 0x52,
	0x61, 0x66, 0x61, 0x79, 0x2e, 0x44, 0x65, 0x76, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0xca, 0x02, 0x16, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65,
	0x76, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0xe2, 0x02,
	0x22, 0x52, 0x61, 0x66, 0x61, 0x79, 0x5c, 0x44, 0x65, 0x76, 0x5c, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x5c, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x52, 0x61, 0x66, 0x61, 0x79, 0x3a, 0x3a, 0x44, 0x65, 0x76,
	0x3a, 0x3a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x53, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_types_sentry_kubeconfig_setting_proto_rawDescOnce sync.Once
	file_proto_types_sentry_kubeconfig_setting_proto_rawDescData = file_proto_types_sentry_kubeconfig_setting_proto_rawDesc
)

func file_proto_types_sentry_kubeconfig_setting_proto_rawDescGZIP() []byte {
	file_proto_types_sentry_kubeconfig_setting_proto_rawDescOnce.Do(func() {
		file_proto_types_sentry_kubeconfig_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_types_sentry_kubeconfig_setting_proto_rawDescData)
	})
	return file_proto_types_sentry_kubeconfig_setting_proto_rawDescData
}

var file_proto_types_sentry_kubeconfig_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_types_sentry_kubeconfig_setting_proto_goTypes = []interface{}{
	(*KubeconfigRevocation)(nil),  // 0: rafay.dev.types.sentry.KubeconfigRevocation
	(*KubeconfigSetting)(nil),     // 1: rafay.dev.types.sentry.KubeconfigSetting
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_types_sentry_kubeconfig_setting_proto_depIdxs = []int32{
	2, // 0: rafay.dev.types.sentry.KubeconfigRevocation.revokedAt:type_name -> google.protobuf.Timestamp
	2, // 1: rafay.dev.types.sentry.KubeconfigRevocation.createdAt:type_name -> google.protobuf.Timestamp
	2, // 2: rafay.dev.types.sentry.KubeconfigSetting.createdAt:type_name -> google.protobuf.Timestamp
	2, // 3: rafay.dev.types.sentry.KubeconfigSetting.modifiedAt:type_name -> google.protobuf.Timestamp
	2, // 4: rafay.dev.types.sentry.KubeconfigSetting.deletedAt:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_types_sentry_kubeconfig_setting_proto_init() }
func file_proto_types_sentry_kubeconfig_setting_proto_init() {
	if File_proto_types_sentry_kubeconfig_setting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_types_sentry_kubeconfig_setting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeconfigRevocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_types_sentry_kubeconfig_setting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubeconfigSetting); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_types_sentry_kubeconfig_setting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_types_sentry_kubeconfig_setting_proto_goTypes,
		DependencyIndexes: file_proto_types_sentry_kubeconfig_setting_proto_depIdxs,
		MessageInfos:      file_proto_types_sentry_kubeconfig_setting_proto_msgTypes,
	}.Build()
	File_proto_types_sentry_kubeconfig_setting_proto = out.File
	file_proto_types_sentry_kubeconfig_setting_proto_rawDesc = nil
	file_proto_types_sentry_kubeconfig_setting_proto_goTypes = nil
	file_proto_types_sentry_kubeconfig_setting_proto_depIdxs = nil
}
