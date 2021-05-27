// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: proto/twit/twit.proto

package twit

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Twit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *TwitUUID              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Date     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Text     string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Nickname string                 `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *Twit) Reset() {
	*x = Twit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twit_twit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Twit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Twit) ProtoMessage() {}

func (x *Twit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twit_twit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Twit.ProtoReflect.Descriptor instead.
func (*Twit) Descriptor() ([]byte, []int) {
	return file_proto_twit_twit_proto_rawDescGZIP(), []int{0}
}

func (x *Twit) GetId() *TwitUUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Twit) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Twit) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Twit) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type ResponseTwit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ResponseTwit) Reset() {
	*x = ResponseTwit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twit_twit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseTwit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseTwit) ProtoMessage() {}

func (x *ResponseTwit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twit_twit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseTwit.ProtoReflect.Descriptor instead.
func (*ResponseTwit) Descriptor() ([]byte, []int) {
	return file_proto_twit_twit_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseTwit) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TwitUUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TwitUUID) Reset() {
	*x = TwitUUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_twit_twit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwitUUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwitUUID) ProtoMessage() {}

func (x *TwitUUID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_twit_twit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwitUUID.ProtoReflect.Descriptor instead.
func (*TwitUUID) Descriptor() ([]byte, []int) {
	return file_proto_twit_twit_proto_rawDescGZIP(), []int{2}
}

func (x *TwitUUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_twit_twit_proto protoreflect.FileDescriptor

var file_proto_twit_twit_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x77, 0x69, 0x74, 0x2f, 0x74, 0x77, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x04,
	0x54, 0x77, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x55, 0x55, 0x49, 0x44,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x54, 0x77, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x74, 0x77,
	0x69, 0x74, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xcd, 0x01, 0x0a,
	0x0b, 0x54, 0x77, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08,
	0x67, 0x65, 0x74, 0x54, 0x77, 0x69, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x77, 0x69, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x2d, 0x0a, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x77, 0x69, 0x74, 0x12, 0x0a, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x77, 0x69, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x77, 0x69, 0x74, 0x22, 0x00, 0x12,
	0x27, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x54, 0x77, 0x69, 0x74, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x54, 0x77, 0x69, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x77, 0x69, 0x74, 0x12, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x74, 0x77,
	0x69, 0x74, 0x55, 0x55, 0x49, 0x44, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x77, 0x69, 0x74, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_twit_twit_proto_rawDescOnce sync.Once
	file_proto_twit_twit_proto_rawDescData = file_proto_twit_twit_proto_rawDesc
)

func file_proto_twit_twit_proto_rawDescGZIP() []byte {
	file_proto_twit_twit_proto_rawDescOnce.Do(func() {
		file_proto_twit_twit_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_twit_twit_proto_rawDescData)
	})
	return file_proto_twit_twit_proto_rawDescData
}

var file_proto_twit_twit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_twit_twit_proto_goTypes = []interface{}{
	(*Twit)(nil),                  // 0: main.Twit
	(*ResponseTwit)(nil),          // 1: main.ResponseTwit
	(*TwitUUID)(nil),              // 2: main.twitUUID
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_proto_twit_twit_proto_depIdxs = []int32{
	2, // 0: main.Twit.id:type_name -> main.twitUUID
	3, // 1: main.Twit.date:type_name -> google.protobuf.Timestamp
	4, // 2: main.TwitService.getTwits:input_type -> google.protobuf.Empty
	0, // 3: main.TwitService.writeTwit:input_type -> main.Twit
	2, // 4: main.TwitService.getTwit:input_type -> main.twitUUID
	2, // 5: main.TwitService.deleteTwit:input_type -> main.twitUUID
	0, // 6: main.TwitService.getTwits:output_type -> main.Twit
	1, // 7: main.TwitService.writeTwit:output_type -> main.ResponseTwit
	0, // 8: main.TwitService.getTwit:output_type -> main.Twit
	1, // 9: main.TwitService.deleteTwit:output_type -> main.ResponseTwit
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_twit_twit_proto_init() }
func file_proto_twit_twit_proto_init() {
	if File_proto_twit_twit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_twit_twit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Twit); i {
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
		file_proto_twit_twit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseTwit); i {
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
		file_proto_twit_twit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwitUUID); i {
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
			RawDescriptor: file_proto_twit_twit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_twit_twit_proto_goTypes,
		DependencyIndexes: file_proto_twit_twit_proto_depIdxs,
		MessageInfos:      file_proto_twit_twit_proto_msgTypes,
	}.Build()
	File_proto_twit_twit_proto = out.File
	file_proto_twit_twit_proto_rawDesc = nil
	file_proto_twit_twit_proto_goTypes = nil
	file_proto_twit_twit_proto_depIdxs = nil
}
