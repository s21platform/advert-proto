// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.26.1
// source: advert.proto

package advert_proto

import (
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

type GetAdvertIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAdvertIn) Reset() {
	*x = GetAdvertIn{}
	mi := &file_advert_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAdvertIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdvertIn) ProtoMessage() {}

func (x *GetAdvertIn) ProtoReflect() protoreflect.Message {
	mi := &file_advert_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdvertIn.ProtoReflect.Descriptor instead.
func (*GetAdvertIn) Descriptor() ([]byte, []int) {
	return file_advert_proto_rawDescGZIP(), []int{0}
}

type AdvertText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TextContent string                 `protobuf:"bytes,1,opt,name=text_content,json=textContent,proto3" json:"text_content,omitempty"`
	ExpiredAt   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
}

func (x *AdvertText) Reset() {
	*x = AdvertText{}
	mi := &file_advert_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdvertText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvertText) ProtoMessage() {}

func (x *AdvertText) ProtoReflect() protoreflect.Message {
	mi := &file_advert_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvertText.ProtoReflect.Descriptor instead.
func (*AdvertText) Descriptor() ([]byte, []int) {
	return file_advert_proto_rawDescGZIP(), []int{1}
}

func (x *AdvertText) GetTextContent() string {
	if x != nil {
		return x.TextContent
	}
	return ""
}

func (x *AdvertText) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

type GetAdvertOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adverts []*AdvertText `protobuf:"bytes,1,rep,name=adverts,proto3" json:"adverts,omitempty"`
}

func (x *GetAdvertOut) Reset() {
	*x = GetAdvertOut{}
	mi := &file_advert_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAdvertOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdvertOut) ProtoMessage() {}

func (x *GetAdvertOut) ProtoReflect() protoreflect.Message {
	mi := &file_advert_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdvertOut.ProtoReflect.Descriptor instead.
func (*GetAdvertOut) Descriptor() ([]byte, []int) {
	return file_advert_proto_rawDescGZIP(), []int{2}
}

func (x *GetAdvertOut) GetAdverts() []*AdvertText {
	if x != nil {
		return x.Adverts
	}
	return nil
}

type UserFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os string `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
}

func (x *UserFilter) Reset() {
	*x = UserFilter{}
	mi := &file_advert_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFilter) ProtoMessage() {}

func (x *UserFilter) ProtoReflect() protoreflect.Message {
	mi := &file_advert_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFilter.ProtoReflect.Descriptor instead.
func (*UserFilter) Descriptor() ([]byte, []int) {
	return file_advert_proto_rawDescGZIP(), []int{3}
}

func (x *UserFilter) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

type CreateAdvertIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text      string                 `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	User      *UserFilter            `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	ExpiredAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at,omitempty"`
}

func (x *CreateAdvertIn) Reset() {
	*x = CreateAdvertIn{}
	mi := &file_advert_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAdvertIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdvertIn) ProtoMessage() {}

func (x *CreateAdvertIn) ProtoReflect() protoreflect.Message {
	mi := &file_advert_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdvertIn.ProtoReflect.Descriptor instead.
func (*CreateAdvertIn) Descriptor() ([]byte, []int) {
	return file_advert_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAdvertIn) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreateAdvertIn) GetUser() *UserFilter {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateAdvertIn) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

type CreateAdvertOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAdvertOut) Reset() {
	*x = CreateAdvertOut{}
	mi := &file_advert_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAdvertOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdvertOut) ProtoMessage() {}

func (x *CreateAdvertOut) ProtoReflect() protoreflect.Message {
	mi := &file_advert_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdvertOut.ProtoReflect.Descriptor instead.
func (*CreateAdvertOut) Descriptor() ([]byte, []int) {
	return file_advert_proto_rawDescGZIP(), []int{5}
}

var File_advert_proto protoreflect.FileDescriptor

var file_advert_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x22, 0x6a,
	0x0a, 0x0a, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0x35, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x07, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x73, 0x22, 0x1c, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x22,
	0x80, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x49, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x4f, 0x75, 0x74, 0x32, 0x70, 0x0a, 0x0d, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x12, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x49,
	0x6e, 0x1a, 0x0d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x4f, 0x75, 0x74,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x12, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x49, 0x6e, 0x1a, 0x10, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x61, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_advert_proto_rawDescOnce sync.Once
	file_advert_proto_rawDescData = file_advert_proto_rawDesc
)

func file_advert_proto_rawDescGZIP() []byte {
	file_advert_proto_rawDescOnce.Do(func() {
		file_advert_proto_rawDescData = protoimpl.X.CompressGZIP(file_advert_proto_rawDescData)
	})
	return file_advert_proto_rawDescData
}

var file_advert_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_advert_proto_goTypes = []any{
	(*GetAdvertIn)(nil),           // 0: GetAdvertIn
	(*AdvertText)(nil),            // 1: AdvertText
	(*GetAdvertOut)(nil),          // 2: GetAdvertOut
	(*UserFilter)(nil),            // 3: UserFilter
	(*CreateAdvertIn)(nil),        // 4: CreateAdvertIn
	(*CreateAdvertOut)(nil),       // 5: CreateAdvertOut
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_advert_proto_depIdxs = []int32{
	6, // 0: AdvertText.expired_at:type_name -> google.protobuf.Timestamp
	1, // 1: GetAdvertOut.adverts:type_name -> AdvertText
	3, // 2: CreateAdvertIn.user:type_name -> UserFilter
	6, // 3: CreateAdvertIn.expired_at:type_name -> google.protobuf.Timestamp
	0, // 4: AdvertService.GetAdvert:input_type -> GetAdvertIn
	4, // 5: AdvertService.CreateAdvert:input_type -> CreateAdvertIn
	2, // 6: AdvertService.GetAdvert:output_type -> GetAdvertOut
	5, // 7: AdvertService.CreateAdvert:output_type -> CreateAdvertOut
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_advert_proto_init() }
func file_advert_proto_init() {
	if File_advert_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_advert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_advert_proto_goTypes,
		DependencyIndexes: file_advert_proto_depIdxs,
		MessageInfos:      file_advert_proto_msgTypes,
	}.Build()
	File_advert_proto = out.File
	file_advert_proto_rawDesc = nil
	file_advert_proto_goTypes = nil
	file_advert_proto_depIdxs = nil
}
