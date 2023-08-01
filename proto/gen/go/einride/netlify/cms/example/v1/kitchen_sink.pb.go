// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: einride/netlify/cms/example/v1/kitchen_sink.proto

package examplev1

import (
	_ "go.einride.tech/protobuf-netlify-cms/proto/gen/go/einride/netlify/cms/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Example enum.
type KitchenSink_ExampleEnum int32

const (
	// Default value. This value is unused.
	KitchenSink_EXAMPLE_ENUM_UNSPECIFIED KitchenSink_ExampleEnum = 0
	// One.
	KitchenSink_ONE KitchenSink_ExampleEnum = 1
	// Two.
	KitchenSink_TWO KitchenSink_ExampleEnum = 2
)

// Enum value maps for KitchenSink_ExampleEnum.
var (
	KitchenSink_ExampleEnum_name = map[int32]string{
		0: "EXAMPLE_ENUM_UNSPECIFIED",
		1: "ONE",
		2: "TWO",
	}
	KitchenSink_ExampleEnum_value = map[string]int32{
		"EXAMPLE_ENUM_UNSPECIFIED": 0,
		"ONE":                      1,
		"TWO":                      2,
	}
)

func (x KitchenSink_ExampleEnum) Enum() *KitchenSink_ExampleEnum {
	p := new(KitchenSink_ExampleEnum)
	*p = x
	return p
}

func (x KitchenSink_ExampleEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KitchenSink_ExampleEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_einride_netlify_cms_example_v1_kitchen_sink_proto_enumTypes[0].Descriptor()
}

func (KitchenSink_ExampleEnum) Type() protoreflect.EnumType {
	return &file_einride_netlify_cms_example_v1_kitchen_sink_proto_enumTypes[0]
}

func (x KitchenSink_ExampleEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KitchenSink_ExampleEnum.Descriptor instead.
func (KitchenSink_ExampleEnum) EnumDescriptor() ([]byte, []int) {
	return file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescGZIP(), []int{0, 0}
}

// A kitchen sink example message.
type KitchenSink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the kitchen sink.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The timestamp the kitchen sink was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The revision ID of the kitchen sink.
	RevisionId string `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	// The timestamp when the kitchen sink revision was created.
	RevisionCreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=revision_create_time,json=revisionCreateTime,proto3" json:"revision_create_time,omitempty"`
	// Display name of the kitchen sink.
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// An example enum.
	ExampleEnum KitchenSink_ExampleEnum `protobuf:"varint,6,opt,name=example_enum,json=exampleEnum,proto3,enum=einride.netlify.cms.example.v1.KitchenSink_ExampleEnum" json:"example_enum,omitempty"`
	// A double value.
	DoubleValue float64 `protobuf:"fixed64,7,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	// A float value.
	FloatValue float32 `protobuf:"fixed32,8,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	// An int64 value.
	Int64Value int64 `protobuf:"varint,9,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	// A value with a custom widget.
	CustomValue string `protobuf:"bytes,10,opt,name=custom_value,json=customValue,proto3" json:"custom_value,omitempty"`
	// A value with relation to another entity
	Book string `protobuf:"bytes,11,opt,name=book,proto3" json:"book,omitempty"`
	// A nested list of some specs to show usage of a list widget.
	Specs []*KitchenSink_SomeSpec `protobuf:"bytes,12,rep,name=specs,proto3" json:"specs,omitempty"`
}

func (x *KitchenSink) Reset() {
	*x = KitchenSink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_netlify_cms_example_v1_kitchen_sink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KitchenSink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KitchenSink) ProtoMessage() {}

func (x *KitchenSink) ProtoReflect() protoreflect.Message {
	mi := &file_einride_netlify_cms_example_v1_kitchen_sink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KitchenSink.ProtoReflect.Descriptor instead.
func (*KitchenSink) Descriptor() ([]byte, []int) {
	return file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescGZIP(), []int{0}
}

func (x *KitchenSink) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KitchenSink) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *KitchenSink) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

func (x *KitchenSink) GetRevisionCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.RevisionCreateTime
	}
	return nil
}

func (x *KitchenSink) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *KitchenSink) GetExampleEnum() KitchenSink_ExampleEnum {
	if x != nil {
		return x.ExampleEnum
	}
	return KitchenSink_EXAMPLE_ENUM_UNSPECIFIED
}

func (x *KitchenSink) GetDoubleValue() float64 {
	if x != nil {
		return x.DoubleValue
	}
	return 0
}

func (x *KitchenSink) GetFloatValue() float32 {
	if x != nil {
		return x.FloatValue
	}
	return 0
}

func (x *KitchenSink) GetInt64Value() int64 {
	if x != nil {
		return x.Int64Value
	}
	return 0
}

func (x *KitchenSink) GetCustomValue() string {
	if x != nil {
		return x.CustomValue
	}
	return ""
}

func (x *KitchenSink) GetBook() string {
	if x != nil {
		return x.Book
	}
	return ""
}

func (x *KitchenSink) GetSpecs() []*KitchenSink_SomeSpec {
	if x != nil {
		return x.Specs
	}
	return nil
}

// SomeSpec is a dummy message struct holds some dummy fields.
type KitchenSink_SomeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *KitchenSink_SomeSpec) Reset() {
	*x = KitchenSink_SomeSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_einride_netlify_cms_example_v1_kitchen_sink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KitchenSink_SomeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KitchenSink_SomeSpec) ProtoMessage() {}

func (x *KitchenSink_SomeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_einride_netlify_cms_example_v1_kitchen_sink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KitchenSink_SomeSpec.ProtoReflect.Descriptor instead.
func (*KitchenSink_SomeSpec) Descriptor() ([]byte, []int) {
	return file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescGZIP(), []int{0, 0}
}

func (x *KitchenSink_SomeSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KitchenSink_SomeSpec) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_einride_netlify_cms_example_v1_kitchen_sink_proto protoreflect.FileDescriptor

var file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDesc = []byte{
	0x0a, 0x31, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66,
	0x79, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x6e, 0x65, 0x74,
	0x6c, 0x69, 0x66, 0x79, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x28, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x6e, 0x65, 0x74,
	0x6c, 0x69, 0x66, 0x79, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x09, 0x0a, 0x0b, 0x4b,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x56, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x42, 0xaa, 0xf6, 0xa1, 0xf3, 0x07, 0x3c,
	0x22, 0x3a, 0xaa, 0x01, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x20, 0x27, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e,
	0x53, 0x69, 0x6e, 0x6b, 0x73, 0x2f, 0x27, 0x12, 0x06, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x3a, 0x12,
	0x0b, 0x20, 0x20, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x20, 0x34, 0x32, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x05, 0xe0, 0x41,
	0x03, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x51, 0x0a,
	0x14, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x12, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37,
	0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79,
	0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x49,
	0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x26, 0xaa, 0xf6, 0xa1, 0xf3, 0x07, 0x20, 0x22, 0x1e, 0xaa, 0x01,
	0x1b, 0x0a, 0x04, 0x74, 0x65, 0x73, 0x74, 0x12, 0x06, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x3a, 0x12,
	0x0b, 0x20, 0x20, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x3a, 0x20, 0x34, 0x32, 0x52, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x6d, 0x0a, 0x04, 0x62, 0x6f, 0x6f,
	0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x59, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x27, 0x0a,
	0x25, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x2d, 0x63, 0x6d, 0x73, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63,
	0x68, 0x2f, 0x42, 0x6f, 0x6f, 0x6b, 0xaa, 0xf6, 0xa1, 0xf3, 0x07, 0x26, 0x22, 0x24, 0x8a, 0x01,
	0x21, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x81, 0x01, 0x0a, 0x05, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x6e, 0x53, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x63, 0x42, 0x35,
	0xaa, 0xf6, 0xa1, 0xf3, 0x07, 0x2f, 0x22, 0x2d, 0x62, 0x2b, 0x1a, 0x29, 0x7b, 0x7b, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x20, 0x2d, 0x20, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x3a, 0x20, 0x7b, 0x7b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x7d, 0x7d, 0x52, 0x05, 0x73, 0x70, 0x65, 0x63, 0x73, 0x1a, 0x34, 0x0a, 0x08,
	0x53, 0x6f, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x3d, 0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e, 0x75,
	0x6d, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x41, 0x4d, 0x50, 0x4c, 0x45, 0x5f, 0x45, 0x4e, 0x55,
	0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x07, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10,
	0x02, 0x3a, 0xd8, 0x01, 0xea, 0x41, 0x4b, 0x0a, 0x2c, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79,
	0x2d, 0x63, 0x6d, 0x73, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x6e, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x1b, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53, 0x69,
	0x6e, 0x6b, 0x73, 0x2f, 0x7b, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x5f, 0x73, 0x69, 0x6e,
	0x6b, 0x7d, 0xda, 0xf6, 0xf1, 0x97, 0x02, 0x83, 0x01, 0x0a, 0x0d, 0x6b, 0x69, 0x74, 0x63, 0x68,
	0x65, 0x6e, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x0d,
	0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x20, 0x53, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x0c, 0x4b,
	0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x20, 0x53, 0x69, 0x6e, 0x6b, 0x2a, 0x1d, 0x4b, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x6e, 0x20, 0x73, 0x69, 0x6e, 0x6b, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x14, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53, 0x69, 0x6e, 0x6b, 0x73,
	0x38, 0x01, 0x42, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x4a, 0x10, 0x7b, 0x7b, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x7d, 0x52, 0x00, 0x42, 0xaf, 0x02, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x6e, 0x65, 0x74,
	0x6c, 0x69, 0x66, 0x79, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x10, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53, 0x69, 0x6e, 0x6b,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5a, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2d, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x2d, 0x63, 0x6d, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69,
	0x64, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x45, 0x4e, 0x43, 0x45, 0xaa, 0x02, 0x1e, 0x45, 0x69, 0x6e,
	0x72, 0x69, 0x64, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x2e, 0x43, 0x6d, 0x73,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x45, 0x69,
	0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x4e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x5c, 0x43, 0x6d,
	0x73, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2a, 0x45,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x4e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x5c, 0x43,
	0x6d, 0x73, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x22, 0x45, 0x69, 0x6e, 0x72,
	0x69, 0x64, 0x65, 0x3a, 0x3a, 0x4e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x3a, 0x3a, 0x43, 0x6d,
	0x73, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescOnce sync.Once
	file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescData = file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDesc
)

func file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescGZIP() []byte {
	file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescOnce.Do(func() {
		file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescData = protoimpl.X.CompressGZIP(file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescData)
	})
	return file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDescData
}

var file_einride_netlify_cms_example_v1_kitchen_sink_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_einride_netlify_cms_example_v1_kitchen_sink_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_einride_netlify_cms_example_v1_kitchen_sink_proto_goTypes = []interface{}{
	(KitchenSink_ExampleEnum)(0),  // 0: einride.netlify.cms.example.v1.KitchenSink.ExampleEnum
	(*KitchenSink)(nil),           // 1: einride.netlify.cms.example.v1.KitchenSink
	(*KitchenSink_SomeSpec)(nil),  // 2: einride.netlify.cms.example.v1.KitchenSink.SomeSpec
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_einride_netlify_cms_example_v1_kitchen_sink_proto_depIdxs = []int32{
	3, // 0: einride.netlify.cms.example.v1.KitchenSink.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: einride.netlify.cms.example.v1.KitchenSink.revision_create_time:type_name -> google.protobuf.Timestamp
	0, // 2: einride.netlify.cms.example.v1.KitchenSink.example_enum:type_name -> einride.netlify.cms.example.v1.KitchenSink.ExampleEnum
	2, // 3: einride.netlify.cms.example.v1.KitchenSink.specs:type_name -> einride.netlify.cms.example.v1.KitchenSink.SomeSpec
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_einride_netlify_cms_example_v1_kitchen_sink_proto_init() }
func file_einride_netlify_cms_example_v1_kitchen_sink_proto_init() {
	if File_einride_netlify_cms_example_v1_kitchen_sink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_einride_netlify_cms_example_v1_kitchen_sink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KitchenSink); i {
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
		file_einride_netlify_cms_example_v1_kitchen_sink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KitchenSink_SomeSpec); i {
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
			RawDescriptor: file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_netlify_cms_example_v1_kitchen_sink_proto_goTypes,
		DependencyIndexes: file_einride_netlify_cms_example_v1_kitchen_sink_proto_depIdxs,
		EnumInfos:         file_einride_netlify_cms_example_v1_kitchen_sink_proto_enumTypes,
		MessageInfos:      file_einride_netlify_cms_example_v1_kitchen_sink_proto_msgTypes,
	}.Build()
	File_einride_netlify_cms_example_v1_kitchen_sink_proto = out.File
	file_einride_netlify_cms_example_v1_kitchen_sink_proto_rawDesc = nil
	file_einride_netlify_cms_example_v1_kitchen_sink_proto_goTypes = nil
	file_einride_netlify_cms_example_v1_kitchen_sink_proto_depIdxs = nil
}
