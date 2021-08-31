// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: mutuals.proto

package mutual

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetMutualRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target  int32 `protobuf:"varint,1,opt,name=target,proto3" json:"target,omitempty"`
	Current int32 `protobuf:"varint,2,opt,name=current,proto3" json:"current,omitempty"`
}

func (x *GetMutualRequest) Reset() {
	*x = GetMutualRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutuals_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMutualRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMutualRequest) ProtoMessage() {}

func (x *GetMutualRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mutuals_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMutualRequest.ProtoReflect.Descriptor instead.
func (*GetMutualRequest) Descriptor() ([]byte, []int) {
	return file_mutuals_proto_rawDescGZIP(), []int{0}
}

func (x *GetMutualRequest) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

func (x *GetMutualRequest) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutuals_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_mutuals_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_mutuals_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Group) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetMutualResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetMutualResponse) Reset() {
	*x = GetMutualResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mutuals_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMutualResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMutualResponse) ProtoMessage() {}

func (x *GetMutualResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mutuals_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMutualResponse.ProtoReflect.Descriptor instead.
func (*GetMutualResponse) Descriptor() ([]byte, []int) {
	return file_mutuals_proto_rawDescGZIP(), []int{2}
}

func (x *GetMutualResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_mutuals_proto protoreflect.FileDescriptor

var file_mutuals_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x74, 0x75,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x05, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d,
	0x75, 0x74, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x32, 0x65, 0x0a, 0x0a, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x57, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x12,
	0x18, 0x2e, 0x6d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x74, 0x75,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x75, 0x74, 0x75,
	0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x67,
	0x65, 0x74, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x3a, 0x01, 0x2a, 0x42, 0x17, 0x5a, 0x15, 0x6b,
	0x61, 0x72, 0x61, 0x73, 0x65, 0x6e, 0x6b, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x75,
	0x74, 0x75, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mutuals_proto_rawDescOnce sync.Once
	file_mutuals_proto_rawDescData = file_mutuals_proto_rawDesc
)

func file_mutuals_proto_rawDescGZIP() []byte {
	file_mutuals_proto_rawDescOnce.Do(func() {
		file_mutuals_proto_rawDescData = protoimpl.X.CompressGZIP(file_mutuals_proto_rawDescData)
	})
	return file_mutuals_proto_rawDescData
}

var file_mutuals_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mutuals_proto_goTypes = []interface{}{
	(*GetMutualRequest)(nil),  // 0: mutual.GetMutualRequest
	(*Group)(nil),             // 1: mutual.Group
	(*GetMutualResponse)(nil), // 2: mutual.GetMutualResponse
}
var file_mutuals_proto_depIdxs = []int32{
	1, // 0: mutual.GetMutualResponse.groups:type_name -> mutual.Group
	0, // 1: mutual.rusprofile.GetMutual:input_type -> mutual.GetMutualRequest
	2, // 2: mutual.rusprofile.GetMutual:output_type -> mutual.GetMutualResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mutuals_proto_init() }
func file_mutuals_proto_init() {
	if File_mutuals_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mutuals_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMutualRequest); i {
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
		file_mutuals_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_mutuals_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMutualResponse); i {
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
			RawDescriptor: file_mutuals_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mutuals_proto_goTypes,
		DependencyIndexes: file_mutuals_proto_depIdxs,
		MessageInfos:      file_mutuals_proto_msgTypes,
	}.Build()
	File_mutuals_proto = out.File
	file_mutuals_proto_rawDesc = nil
	file_mutuals_proto_goTypes = nil
	file_mutuals_proto_depIdxs = nil
}