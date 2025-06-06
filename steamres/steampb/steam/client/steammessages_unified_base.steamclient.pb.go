// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: steam/client/steammessages_unified_base.steamclient.proto

package client

import (
	proto "google.golang.org/protobuf/proto"
	descriptor "google.golang.org/protobuf/protoc-gen-go/descriptor"
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

type EProtoExecutionSite int32

const (
	EProtoExecutionSite_k_EProtoExecutionSiteUnknown     EProtoExecutionSite = 0
	EProtoExecutionSite_k_EProtoExecutionSiteSteamClient EProtoExecutionSite = 2
)

// Enum value maps for EProtoExecutionSite.
var (
	EProtoExecutionSite_name = map[int32]string{
		0: "k_EProtoExecutionSiteUnknown",
		2: "k_EProtoExecutionSiteSteamClient",
	}
	EProtoExecutionSite_value = map[string]int32{
		"k_EProtoExecutionSiteUnknown":     0,
		"k_EProtoExecutionSiteSteamClient": 2,
	}
)

func (x EProtoExecutionSite) Enum() *EProtoExecutionSite {
	p := new(EProtoExecutionSite)
	*p = x
	return p
}

func (x EProtoExecutionSite) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EProtoExecutionSite) Descriptor() protoreflect.EnumDescriptor {
	return file_steam_client_steammessages_unified_base_steamclient_proto_enumTypes[0].Descriptor()
}

func (EProtoExecutionSite) Type() protoreflect.EnumType {
	return &file_steam_client_steammessages_unified_base_steamclient_proto_enumTypes[0]
}

func (x EProtoExecutionSite) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EProtoExecutionSite) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EProtoExecutionSite(num)
	return nil
}

// Deprecated: Use EProtoExecutionSite.Descriptor instead.
func (EProtoExecutionSite) EnumDescriptor() ([]byte, []int) {
	return file_steam_client_steammessages_unified_base_steamclient_proto_rawDescGZIP(), []int{0}
}

type NoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoResponse) Reset() {
	*x = NoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steam_client_steammessages_unified_base_steamclient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoResponse) ProtoMessage() {}

func (x *NoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_steam_client_steammessages_unified_base_steamclient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoResponse.ProtoReflect.Descriptor instead.
func (*NoResponse) Descriptor() ([]byte, []int) {
	return file_steam_client_steammessages_unified_base_steamclient_proto_rawDescGZIP(), []int{0}
}

var file_steam_client_steammessages_unified_base_steamclient_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "steam.client.description",
		Tag:           "bytes,50000,opt,name=description",
		Filename:      "steam/client/steammessages_unified_base.steamclient.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "steam.client.service_description",
		Tag:           "bytes,50000,opt,name=service_description",
		Filename:      "steam/client/steammessages_unified_base.steamclient.proto",
	},
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*EProtoExecutionSite)(nil),
		Field:         50008,
		Name:          "steam.client.service_execution_site",
		Tag:           "varint,50008,opt,name=service_execution_site,enum=steam.client.EProtoExecutionSite,def=0",
		Filename:      "steam/client/steammessages_unified_base.steamclient.proto",
	},
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "steam.client.method_description",
		Tag:           "bytes,50000,opt,name=method_description",
		Filename:      "steam/client/steammessages_unified_base.steamclient.proto",
	},
	{
		ExtendedType:  (*descriptor.EnumOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "steam.client.enum_description",
		Tag:           "bytes,50000,opt,name=enum_description",
		Filename:      "steam/client/steammessages_unified_base.steamclient.proto",
	},
	{
		ExtendedType:  (*descriptor.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         50000,
		Name:          "steam.client.enum_value_description",
		Tag:           "bytes,50000,opt,name=enum_value_description",
		Filename:      "steam/client/steammessages_unified_base.steamclient.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional string description = 50000;
	E_Description = &file_steam_client_steammessages_unified_base_steamclient_proto_extTypes[0]
)

// Extension fields to descriptor.ServiceOptions.
var (
	// optional string service_description = 50000;
	E_ServiceDescription = &file_steam_client_steammessages_unified_base_steamclient_proto_extTypes[1]
	// optional steam.client.EProtoExecutionSite service_execution_site = 50008;
	E_ServiceExecutionSite = &file_steam_client_steammessages_unified_base_steamclient_proto_extTypes[2]
)

// Extension fields to descriptor.MethodOptions.
var (
	// optional string method_description = 50000;
	E_MethodDescription = &file_steam_client_steammessages_unified_base_steamclient_proto_extTypes[3]
)

// Extension fields to descriptor.EnumOptions.
var (
	// optional string enum_description = 50000;
	E_EnumDescription = &file_steam_client_steammessages_unified_base_steamclient_proto_extTypes[4]
)

// Extension fields to descriptor.EnumValueOptions.
var (
	// optional string enum_value_description = 50000;
	E_EnumValueDescription = &file_steam_client_steammessages_unified_base_steamclient_proto_extTypes[5]
)

var File_steam_client_steammessages_unified_base_steamclient_proto protoreflect.FileDescriptor

var file_steam_client_steammessages_unified_base_steamclient_proto_rawDesc = []byte{
	0x0a, 0x39, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x74, 0x65,
	0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x4e,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x5d, 0x0a, 0x13, 0x45, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x1c, 0x6b, 0x5f, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x74, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x6b, 0x5f, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x74, 0x65, 0x53, 0x74, 0x65, 0x61, 0x6d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x3a, 0x41, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x52, 0x0a, 0x13, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x98, 0x01, 0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd8, 0x86, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x69, 0x74, 0x65, 0x3a, 0x1c, 0x6b, 0x5f, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x74, 0x65, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x74, 0x65, 0x3a, 0x4f, 0x0a, 0x12, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x49, 0x0a, 0x10, 0x65,
	0x6e, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x75, 0x6d, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x59, 0x0a, 0x16, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x6e, 0x75,
	0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x40, 0x48, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x31, 0x33, 0x6b, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x80, 0x01, 0x00,
}

var (
	file_steam_client_steammessages_unified_base_steamclient_proto_rawDescOnce sync.Once
	file_steam_client_steammessages_unified_base_steamclient_proto_rawDescData = file_steam_client_steammessages_unified_base_steamclient_proto_rawDesc
)

func file_steam_client_steammessages_unified_base_steamclient_proto_rawDescGZIP() []byte {
	file_steam_client_steammessages_unified_base_steamclient_proto_rawDescOnce.Do(func() {
		file_steam_client_steammessages_unified_base_steamclient_proto_rawDescData = protoimpl.X.CompressGZIP(file_steam_client_steammessages_unified_base_steamclient_proto_rawDescData)
	})
	return file_steam_client_steammessages_unified_base_steamclient_proto_rawDescData
}

var file_steam_client_steammessages_unified_base_steamclient_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_steam_client_steammessages_unified_base_steamclient_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_steam_client_steammessages_unified_base_steamclient_proto_goTypes = []interface{}{
	(EProtoExecutionSite)(0),            // 0: steam.client.EProtoExecutionSite
	(*NoResponse)(nil),                  // 1: steam.client.NoResponse
	(*descriptor.FieldOptions)(nil),     // 2: google.protobuf.FieldOptions
	(*descriptor.ServiceOptions)(nil),   // 3: google.protobuf.ServiceOptions
	(*descriptor.MethodOptions)(nil),    // 4: google.protobuf.MethodOptions
	(*descriptor.EnumOptions)(nil),      // 5: google.protobuf.EnumOptions
	(*descriptor.EnumValueOptions)(nil), // 6: google.protobuf.EnumValueOptions
}
var file_steam_client_steammessages_unified_base_steamclient_proto_depIdxs = []int32{
	2, // 0: steam.client.description:extendee -> google.protobuf.FieldOptions
	3, // 1: steam.client.service_description:extendee -> google.protobuf.ServiceOptions
	3, // 2: steam.client.service_execution_site:extendee -> google.protobuf.ServiceOptions
	4, // 3: steam.client.method_description:extendee -> google.protobuf.MethodOptions
	5, // 4: steam.client.enum_description:extendee -> google.protobuf.EnumOptions
	6, // 5: steam.client.enum_value_description:extendee -> google.protobuf.EnumValueOptions
	0, // 6: steam.client.service_execution_site:type_name -> steam.client.EProtoExecutionSite
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	6, // [6:7] is the sub-list for extension type_name
	0, // [0:6] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_steam_client_steammessages_unified_base_steamclient_proto_init() }
func file_steam_client_steammessages_unified_base_steamclient_proto_init() {
	if File_steam_client_steammessages_unified_base_steamclient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_steam_client_steammessages_unified_base_steamclient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoResponse); i {
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
			RawDescriptor: file_steam_client_steammessages_unified_base_steamclient_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_steam_client_steammessages_unified_base_steamclient_proto_goTypes,
		DependencyIndexes: file_steam_client_steammessages_unified_base_steamclient_proto_depIdxs,
		EnumInfos:         file_steam_client_steammessages_unified_base_steamclient_proto_enumTypes,
		MessageInfos:      file_steam_client_steammessages_unified_base_steamclient_proto_msgTypes,
		ExtensionInfos:    file_steam_client_steammessages_unified_base_steamclient_proto_extTypes,
	}.Build()
	File_steam_client_steammessages_unified_base_steamclient_proto = out.File
	file_steam_client_steammessages_unified_base_steamclient_proto_rawDesc = nil
	file_steam_client_steammessages_unified_base_steamclient_proto_goTypes = nil
	file_steam_client_steammessages_unified_base_steamclient_proto_depIdxs = nil
}
