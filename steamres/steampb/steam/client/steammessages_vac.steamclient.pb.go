// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: steam/client/steammessages_vac.steamclient.proto

package client

import (
	proto "google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CFileVerification_SignatureCheck_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steamid              *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	Appid                *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	FileSize             *uint64 `protobuf:"varint,3,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	FileTimestamp        *uint32 `protobuf:"varint,4,opt,name=file_timestamp,json=fileTimestamp" json:"file_timestamp,omitempty"`
	FileTimestamp2       *uint32 `protobuf:"varint,5,opt,name=file_timestamp2,json=fileTimestamp2" json:"file_timestamp2,omitempty"`
	SignatureResult      *uint32 `protobuf:"varint,6,opt,name=signature_result,json=signatureResult" json:"signature_result,omitempty"`
	Filename             *string `protobuf:"bytes,7,opt,name=filename" json:"filename,omitempty"`
	ClientPackageVersion *uint32 `protobuf:"varint,8,opt,name=client_package_version,json=clientPackageVersion" json:"client_package_version,omitempty"`
	Sha1Hash             []byte  `protobuf:"bytes,9,opt,name=sha1hash" json:"sha1hash,omitempty"`
}

func (x *CFileVerification_SignatureCheck_Request) Reset() {
	*x = CFileVerification_SignatureCheck_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_vac_steamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CFileVerification_SignatureCheck_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CFileVerification_SignatureCheck_Request) ProtoMessage() {}

func (x *CFileVerification_SignatureCheck_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_vac_steamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CFileVerification_SignatureCheck_Request.ProtoReflect.Descriptor instead.
func (*CFileVerification_SignatureCheck_Request) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_vac_steamclient_proto_rawDescGZIP(), []int{0}
}

func (x *CFileVerification_SignatureCheck_Request) GetSteamid() uint64 {
	if x != nil && x.Steamid != nil {
		return *x.Steamid
	}
	return 0
}

func (x *CFileVerification_SignatureCheck_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CFileVerification_SignatureCheck_Request) GetFileSize() uint64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *CFileVerification_SignatureCheck_Request) GetFileTimestamp() uint32 {
	if x != nil && x.FileTimestamp != nil {
		return *x.FileTimestamp
	}
	return 0
}

func (x *CFileVerification_SignatureCheck_Request) GetFileTimestamp2() uint32 {
	if x != nil && x.FileTimestamp2 != nil {
		return *x.FileTimestamp2
	}
	return 0
}

func (x *CFileVerification_SignatureCheck_Request) GetSignatureResult() uint32 {
	if x != nil && x.SignatureResult != nil {
		return *x.SignatureResult
	}
	return 0
}

func (x *CFileVerification_SignatureCheck_Request) GetFilename() string {
	if x != nil && x.Filename != nil {
		return *x.Filename
	}
	return ""
}

func (x *CFileVerification_SignatureCheck_Request) GetClientPackageVersion() uint32 {
	if x != nil && x.ClientPackageVersion != nil {
		return *x.ClientPackageVersion
	}
	return 0
}

func (x *CFileVerification_SignatureCheck_Request) GetSha1Hash() []byte {
	if x != nil {
		return x.Sha1Hash
	}
	return nil
}

type CFileVerification_SignatureCheck_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DenyOperation *bool `protobuf:"varint,1,opt,name=deny_operation,json=denyOperation" json:"deny_operation,omitempty"`
}

func (x *CFileVerification_SignatureCheck_Response) Reset() {
	*x = CFileVerification_SignatureCheck_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_vac_steamclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CFileVerification_SignatureCheck_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CFileVerification_SignatureCheck_Response) ProtoMessage() {}

func (x *CFileVerification_SignatureCheck_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_vac_steamclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CFileVerification_SignatureCheck_Response.ProtoReflect.Descriptor instead.
func (*CFileVerification_SignatureCheck_Response) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_vac_steamclient_proto_rawDescGZIP(), []int{1}
}

func (x *CFileVerification_SignatureCheck_Response) GetDenyOperation() bool {
	if x != nil && x.DenyOperation != nil {
		return *x.DenyOperation
	}
	return false
}

type CFileVerification_SteamServiceCheck_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceStatus        *uint32 `protobuf:"varint,2,opt,name=service_status,json=serviceStatus" json:"service_status,omitempty"`
	ClientPackageVersion *uint32 `protobuf:"varint,3,opt,name=client_package_version,json=clientPackageVersion" json:"client_package_version,omitempty"`
	LauncherType         *uint32 `protobuf:"varint,4,opt,name=launcher_type,json=launcherType" json:"launcher_type,omitempty"`
	OsType               *uint32 `protobuf:"varint,5,opt,name=os_type,json=osType" json:"os_type,omitempty"`
	ServiceRepair        *uint32 `protobuf:"varint,6,opt,name=service_repair,json=serviceRepair" json:"service_repair,omitempty"`
}

func (x *CFileVerification_SteamServiceCheck_Request) Reset() {
	*x = CFileVerification_SteamServiceCheck_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_vac_steamclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CFileVerification_SteamServiceCheck_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CFileVerification_SteamServiceCheck_Request) ProtoMessage() {}

func (x *CFileVerification_SteamServiceCheck_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_vac_steamclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CFileVerification_SteamServiceCheck_Request.ProtoReflect.Descriptor instead.
func (*CFileVerification_SteamServiceCheck_Request) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_vac_steamclient_proto_rawDescGZIP(), []int{2}
}

func (x *CFileVerification_SteamServiceCheck_Request) GetServiceStatus() uint32 {
	if x != nil && x.ServiceStatus != nil {
		return *x.ServiceStatus
	}
	return 0
}

func (x *CFileVerification_SteamServiceCheck_Request) GetClientPackageVersion() uint32 {
	if x != nil && x.ClientPackageVersion != nil {
		return *x.ClientPackageVersion
	}
	return 0
}

func (x *CFileVerification_SteamServiceCheck_Request) GetLauncherType() uint32 {
	if x != nil && x.LauncherType != nil {
		return *x.LauncherType
	}
	return 0
}

func (x *CFileVerification_SteamServiceCheck_Request) GetOsType() uint32 {
	if x != nil && x.OsType != nil {
		return *x.OsType
	}
	return 0
}

func (x *CFileVerification_SteamServiceCheck_Request) GetServiceRepair() uint32 {
	if x != nil && x.ServiceRepair != nil {
		return *x.ServiceRepair
	}
	return 0
}

type CFileVerification_SteamServiceCheck_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttemptRepair *bool `protobuf:"varint,1,opt,name=attempt_repair,json=attemptRepair" json:"attempt_repair,omitempty"`
}

func (x *CFileVerification_SteamServiceCheck_Response) Reset() {
	*x = CFileVerification_SteamServiceCheck_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_vac_steamclient_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CFileVerification_SteamServiceCheck_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CFileVerification_SteamServiceCheck_Response) ProtoMessage() {}

func (x *CFileVerification_SteamServiceCheck_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_vac_steamclient_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CFileVerification_SteamServiceCheck_Response.ProtoReflect.Descriptor instead.
func (*CFileVerification_SteamServiceCheck_Response) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_vac_steamclient_proto_rawDescGZIP(), []int{3}
}

func (x *CFileVerification_SteamServiceCheck_Response) GetAttemptRepair() bool {
	if x != nil && x.AttemptRepair != nil {
		return *x.AttemptRepair
	}
	return false
}

var File_steam_client_steammessages_vac_steamclient_proto protoreflect.FileDescriptor

var file_steam_client_steammessages_vac_steamclient_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x63,
	0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x1a, 0x39, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x02, 0x0a, 0x28,
	0x43, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66,
	0x69, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x0f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x32, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x31, 0x68, 0x61, 0x73, 0x68, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x73, 0x68, 0x61, 0x31, 0x68, 0x61, 0x73, 0x68, 0x22, 0x52,
	0x0a, 0x29, 0x43, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64,
	0x65, 0x6e, 0x79, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x64, 0x65, 0x6e, 0x79, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xef, 0x01, 0x0a, 0x2b, 0x43, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x22, 0x55, 0x0a, 0x2c, 0x43, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f,
	0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x32, 0x92, 0x03, 0x0a, 0x10,
	0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0xa9, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x36, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x46, 0x69, 0x6c, 0x65,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xb5, 0x18, 0x22, 0x46, 0x69, 0x6c, 0x65, 0x20, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x77,
	0x61, 0x73, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x12, 0xb1, 0x01, 0x0a,
	0x11, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x39, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x46, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e,
	0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x46, 0x69,
	0x6c, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x53,
	0x74, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xb5, 0x18, 0x21, 0x53,
	0x74, 0x65, 0x61, 0x6d, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x20, 0x77, 0x61, 0x73, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64,
	0x1a, 0x1e, 0x82, 0xb5, 0x18, 0x1a, 0x46, 0x69, 0x6c, 0x65, 0x20, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x42, 0x3e, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x31,
	0x33, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x70, 0x62,
	0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x80, 0x01, 0x01,
}

var (
	file_steam_client_steammessages_vac_steamclient_proto_rawDescOnce sync.Once
	file_steam_client_steammessages_vac_steamclient_proto_rawDescData = file_steam_client_steammessages_vac_steamclient_proto_rawDesc
)

func file_steam_client_steammessages_vac_steamclient_proto_rawDescGZIP() []byte {
	file_steam_client_steammessages_vac_steamclient_proto_rawDescOnce.Do(func() {
		file_steam_client_steammessages_vac_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steam_client_steammessages_vac_steamclient_proto_rawDescData)
	})
	return file_steam_client_steammessages_vac_steamclient_proto_rawDescData
}

var file_steam_client_steammessages_vac_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_steam_client_steammessages_vac_steamclient_proto_goTypes = []interface{}{
	(*CFileVerification_SignatureCheck_Request)(nil),     // 0: steam.client.CFileVerification_SignatureCheck_Request
	(*CFileVerification_SignatureCheck_Response)(nil),    // 1: steam.client.CFileVerification_SignatureCheck_Response
	(*CFileVerification_SteamServiceCheck_Request)(nil),  // 2: steam.client.CFileVerification_SteamServiceCheck_Request
	(*CFileVerification_SteamServiceCheck_Response)(nil), // 3: steam.client.CFileVerification_SteamServiceCheck_Response
}
var file_steam_client_steammessages_vac_steamclient_proto_depIdxs = []int32{
	0, // 0: steam.client.FileVerification.SignatureCheck:input_type -> steam.client.CFileVerification_SignatureCheck_Request
	2, // 1: steam.client.FileVerification.SteamServiceCheck:input_type -> steam.client.CFileVerification_SteamServiceCheck_Request
	1, // 2: steam.client.FileVerification.SignatureCheck:output_type -> steam.client.CFileVerification_SignatureCheck_Response
	3, // 3: steam.client.FileVerification.SteamServiceCheck:output_type -> steam.client.CFileVerification_SteamServiceCheck_Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_steam_client_steammessages_vac_steamclient_proto_init() }
func file_steam_client_steammessages_vac_steamclient_proto_init() {
	if File_steam_client_steammessages_vac_steamclient_proto != nil {
		return
	}
	file_steam_client_steammessages_unified_base_steamclient_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steam_client_steammessages_vac_steamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CFileVerification_SignatureCheck_Request); i {
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
		file_steam_client_steammessages_vac_steamclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CFileVerification_SignatureCheck_Response); i {
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
		file_steam_client_steammessages_vac_steamclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CFileVerification_SteamServiceCheck_Request); i {
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
		file_steam_client_steammessages_vac_steamclient_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CFileVerification_SteamServiceCheck_Response); i {
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
			RawDescriptor: file_steam_client_steammessages_vac_steamclient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steam_client_steammessages_vac_steamclient_proto_goTypes,
		DependencyIndexes: file_steam_client_steammessages_vac_steamclient_proto_depIdxs,
		MessageInfos:      file_steam_client_steammessages_vac_steamclient_proto_msgTypes,
	}.Build()
	File_steam_client_steammessages_vac_steamclient_proto = out.File
	file_steam_client_steammessages_vac_steamclient_proto_rawDesc = nil
	file_steam_client_steammessages_vac_steamclient_proto_goTypes = nil
	file_steam_client_steammessages_vac_steamclient_proto_depIdxs = nil
}
