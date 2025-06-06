// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: steam/client/steammessages_secrets.steamclient.proto

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

type EKeyEscrowUsage int32

const (
	EKeyEscrowUsage_k_EKeyEscrowUsageStreamingDevice EKeyEscrowUsage = 0
)

// Enum value maps for EKeyEscrowUsage.
var (
	EKeyEscrowUsage_name = map[int32]string{
		0: "k_EKeyEscrowUsageStreamingDevice",
	}
	EKeyEscrowUsage_value = map[string]int32{
		"k_EKeyEscrowUsageStreamingDevice": 0,
	}
)

func (x EKeyEscrowUsage) Enum() *EKeyEscrowUsage {
	p := new(EKeyEscrowUsage)
	*p = x
	return p
}

func (x EKeyEscrowUsage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EKeyEscrowUsage) Descriptor() protoreflect.EnumDescriptor {
	return file_steam_client_steammessages_secrets_steamclient_proto_enumTypes[0].Descriptor()
}

func (EKeyEscrowUsage) Type() protoreflect.EnumType {
	return &file_steam_client_steammessages_secrets_steamclient_proto_enumTypes[0]
}

func (x EKeyEscrowUsage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EKeyEscrowUsage) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EKeyEscrowUsage(num)
	return nil
}

// Deprecated: Use EKeyEscrowUsage.Descriptor instead.
func (EKeyEscrowUsage) EnumDescriptor() ([]byte, []int) {
	return file_steam_client_steammessages_secrets_steamclient_proto_rawDescGZIP(), []int{0}
}

type CKeyEscrow_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RsaOaepShaTicket []byte           `protobuf:"bytes,1,opt,name=rsa_oaep_sha_ticket,json=rsaOaepShaTicket" json:"rsa_oaep_sha_ticket,omitempty"`
	Password         []byte           `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Usage            *EKeyEscrowUsage `protobuf:"varint,3,opt,name=usage,enum=steam.client.EKeyEscrowUsage,def=0" json:"usage,omitempty"`
	DeviceName       *string          `protobuf:"bytes,4,opt,name=device_name,json=deviceName" json:"device_name,omitempty"`
}

// Default values for CKeyEscrow_Request fields.
const (
	Default_CKeyEscrow_Request_Usage = EKeyEscrowUsage_k_EKeyEscrowUsageStreamingDevice
)

func (x *CKeyEscrow_Request) Reset() {
	*x = CKeyEscrow_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_secrets_steamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CKeyEscrow_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CKeyEscrow_Request) ProtoMessage() {}

func (x *CKeyEscrow_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_secrets_steamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CKeyEscrow_Request.ProtoReflect.Descriptor instead.
func (*CKeyEscrow_Request) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_secrets_steamclient_proto_rawDescGZIP(), []int{0}
}

func (x *CKeyEscrow_Request) GetRsaOaepShaTicket() []byte {
	if x != nil {
		return x.RsaOaepShaTicket
	}
	return nil
}

func (x *CKeyEscrow_Request) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *CKeyEscrow_Request) GetUsage() EKeyEscrowUsage {
	if x != nil && x.Usage != nil {
		return *x.Usage
	}
	return Default_CKeyEscrow_Request_Usage
}

func (x *CKeyEscrow_Request) GetDeviceName() string {
	if x != nil && x.DeviceName != nil {
		return *x.DeviceName
	}
	return ""
}

type CKeyEscrow_Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password             []byte           `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	Identifier           *uint64          `protobuf:"varint,2,opt,name=identifier" json:"identifier,omitempty"`
	Payload              []byte           `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	Timestamp            *uint32          `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Usage                *EKeyEscrowUsage `protobuf:"varint,5,opt,name=usage,enum=steam.client.EKeyEscrowUsage,def=0" json:"usage,omitempty"`
	DeviceName           *string          `protobuf:"bytes,6,opt,name=device_name,json=deviceName" json:"device_name,omitempty"`
	DeviceModel          *string          `protobuf:"bytes,7,opt,name=device_model,json=deviceModel" json:"device_model,omitempty"`
	DeviceSerial         *string          `protobuf:"bytes,8,opt,name=device_serial,json=deviceSerial" json:"device_serial,omitempty"`
	DeviceProvisioningId *uint32          `protobuf:"varint,9,opt,name=device_provisioning_id,json=deviceProvisioningId" json:"device_provisioning_id,omitempty"`
}

// Default values for CKeyEscrow_Ticket fields.
const (
	Default_CKeyEscrow_Ticket_Usage = EKeyEscrowUsage_k_EKeyEscrowUsageStreamingDevice
)

func (x *CKeyEscrow_Ticket) Reset() {
	*x = CKeyEscrow_Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_secrets_steamclient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CKeyEscrow_Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CKeyEscrow_Ticket) ProtoMessage() {}

func (x *CKeyEscrow_Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_secrets_steamclient_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CKeyEscrow_Ticket.ProtoReflect.Descriptor instead.
func (*CKeyEscrow_Ticket) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_secrets_steamclient_proto_rawDescGZIP(), []int{1}
}

func (x *CKeyEscrow_Ticket) GetPassword() []byte {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *CKeyEscrow_Ticket) GetIdentifier() uint64 {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return 0
}

func (x *CKeyEscrow_Ticket) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *CKeyEscrow_Ticket) GetTimestamp() uint32 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *CKeyEscrow_Ticket) GetUsage() EKeyEscrowUsage {
	if x != nil && x.Usage != nil {
		return *x.Usage
	}
	return Default_CKeyEscrow_Ticket_Usage
}

func (x *CKeyEscrow_Ticket) GetDeviceName() string {
	if x != nil && x.DeviceName != nil {
		return *x.DeviceName
	}
	return ""
}

func (x *CKeyEscrow_Ticket) GetDeviceModel() string {
	if x != nil && x.DeviceModel != nil {
		return *x.DeviceModel
	}
	return ""
}

func (x *CKeyEscrow_Ticket) GetDeviceSerial() string {
	if x != nil && x.DeviceSerial != nil {
		return *x.DeviceSerial
	}
	return ""
}

func (x *CKeyEscrow_Ticket) GetDeviceProvisioningId() uint32 {
	if x != nil && x.DeviceProvisioningId != nil {
		return *x.DeviceProvisioningId
	}
	return 0
}

type CKeyEscrow_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket *CKeyEscrow_Ticket `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
}

func (x *CKeyEscrow_Response) Reset() {
	*x = CKeyEscrow_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_secrets_steamclient_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CKeyEscrow_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CKeyEscrow_Response) ProtoMessage() {}

func (x *CKeyEscrow_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_secrets_steamclient_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CKeyEscrow_Response.ProtoReflect.Descriptor instead.
func (*CKeyEscrow_Response) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_secrets_steamclient_proto_rawDescGZIP(), []int{2}
}

func (x *CKeyEscrow_Response) GetTicket() *CKeyEscrow_Ticket {
	if x != nil {
		return x.Ticket
	}
	return nil
}

var File_steam_client_steammessages_secrets_steamclient_proto protoreflect.FileDescriptor

var file_steam_client_steammessages_secrets_steamclient_proto_rawDesc = []byte{
	0x0a, 0x34, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x1a, 0x39, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x5f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd7, 0x01, 0x0a, 0x12, 0x43, 0x4b, 0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x5f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x72, 0x73, 0x61, 0x5f, 0x6f, 0x61,
	0x65, 0x70, 0x5f, 0x73, 0x68, 0x61, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x10, 0x72, 0x73, 0x61, 0x4f, 0x61, 0x65, 0x70, 0x53, 0x68, 0x61, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x55, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x45, 0x4b, 0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x55, 0x73, 0x61, 0x67, 0x65, 0x3a,
	0x20, 0x6b, 0x5f, 0x45, 0x4b, 0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xfd, 0x02, 0x0a, 0x11, 0x43, 0x4b,
	0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x5f, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x55, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x4b, 0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x3a, 0x20, 0x6b, 0x5f, 0x45, 0x4b, 0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x14, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x13, 0x43, 0x4b, 0x65,
	0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x4b, 0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x5f, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2a, 0x37, 0x0a, 0x0f, 0x45, 0x4b, 0x65,
	0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x20,
	0x6b, 0x5f, 0x45, 0x4b, 0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x10, 0x00, 0x32, 0xe2, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12, 0x9a,
	0x01, 0x0a, 0x09, 0x4b, 0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x12, 0x20, 0x2e, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x4b, 0x65, 0x79,
	0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x4b,
	0x65, 0x79, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x48, 0x82, 0xb5, 0x18, 0x44, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x74,
	0x6f, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x20, 0x6b, 0x65, 0x79, 0x2d, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x20, 0x69, 0x6e, 0x76, 0x6f, 0x6c, 0x76, 0x69, 0x6e, 0x67, 0x20, 0x53,
	0x74, 0x65, 0x61, 0x6d, 0x20, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x3a, 0x82, 0xb5, 0x18,
	0x36, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x20,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x42, 0x3e, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x65,
	0x61, 0x6d, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x2f,
	0x73, 0x74, 0x65, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x80, 0x01, 0x01,
}

var (
	file_steam_client_steammessages_secrets_steamclient_proto_rawDescOnce sync.Once
	file_steam_client_steammessages_secrets_steamclient_proto_rawDescData = file_steam_client_steammessages_secrets_steamclient_proto_rawDesc
)

func file_steam_client_steammessages_secrets_steamclient_proto_rawDescGZIP() []byte {
	file_steam_client_steammessages_secrets_steamclient_proto_rawDescOnce.Do(func() {
		file_steam_client_steammessages_secrets_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steam_client_steammessages_secrets_steamclient_proto_rawDescData)
	})
	return file_steam_client_steammessages_secrets_steamclient_proto_rawDescData
}

var file_steam_client_steammessages_secrets_steamclient_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_steam_client_steammessages_secrets_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_steam_client_steammessages_secrets_steamclient_proto_goTypes = []interface{}{
	(EKeyEscrowUsage)(0),        // 0: steam.client.EKeyEscrowUsage
	(*CKeyEscrow_Request)(nil),  // 1: steam.client.CKeyEscrow_Request
	(*CKeyEscrow_Ticket)(nil),   // 2: steam.client.CKeyEscrow_Ticket
	(*CKeyEscrow_Response)(nil), // 3: steam.client.CKeyEscrow_Response
}
var file_steam_client_steammessages_secrets_steamclient_proto_depIdxs = []int32{
	0, // 0: steam.client.CKeyEscrow_Request.usage:type_name -> steam.client.EKeyEscrowUsage
	0, // 1: steam.client.CKeyEscrow_Ticket.usage:type_name -> steam.client.EKeyEscrowUsage
	2, // 2: steam.client.CKeyEscrow_Response.ticket:type_name -> steam.client.CKeyEscrow_Ticket
	1, // 3: steam.client.Secrets.KeyEscrow:input_type -> steam.client.CKeyEscrow_Request
	3, // 4: steam.client.Secrets.KeyEscrow:output_type -> steam.client.CKeyEscrow_Response
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_steam_client_steammessages_secrets_steamclient_proto_init() }
func file_steam_client_steammessages_secrets_steamclient_proto_init() {
	if File_steam_client_steammessages_secrets_steamclient_proto != nil {
		return
	}
	file_steam_client_steammessages_unified_base_steamclient_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steam_client_steammessages_secrets_steamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CKeyEscrow_Request); i {
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
		file_steam_client_steammessages_secrets_steamclient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CKeyEscrow_Ticket); i {
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
		file_steam_client_steammessages_secrets_steamclient_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CKeyEscrow_Response); i {
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
			RawDescriptor: file_steam_client_steammessages_secrets_steamclient_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steam_client_steammessages_secrets_steamclient_proto_goTypes,
		DependencyIndexes: file_steam_client_steammessages_secrets_steamclient_proto_depIdxs,
		EnumInfos:         file_steam_client_steammessages_secrets_steamclient_proto_enumTypes,
		MessageInfos:      file_steam_client_steammessages_secrets_steamclient_proto_msgTypes,
	}.Build()
	File_steam_client_steammessages_secrets_steamclient_proto = out.File
	file_steam_client_steammessages_secrets_steamclient_proto_rawDesc = nil
	file_steam_client_steammessages_secrets_steamclient_proto_goTypes = nil
	file_steam_client_steammessages_secrets_steamclient_proto_depIdxs = nil
}
