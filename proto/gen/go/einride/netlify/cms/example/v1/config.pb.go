// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: einride/netlify/cms/example/v1/config.proto

package examplev1

import (
	_ "go.einride.tech/protobuf-netlify-cms/proto/gen/go/einride/netlify/cms/v1"
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

var File_einride_netlify_cms_example_v1_config_proto protoreflect.FileDescriptor

var file_einride_netlify_cms_example_v1_config_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66,
	0x79, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x2e, 0x63,
	0x6d, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x65,
	0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x2f, 0x63,
	0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xbc, 0x04, 0xb2, 0xd3, 0xa1, 0xc7, 0x06, 0x8b,
	0x02, 0x0a, 0xc3, 0x01, 0x0a, 0x0b, 0x67, 0x69, 0x74, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x20, 0x01, 0x2a, 0xb1, 0x01, 0x0a, 0x25, 0x66, 0x65, 0x61, 0x74, 0x28, 0x7b, 0x7b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x7d, 0x29, 0x3a, 0x20, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x20, 0x7b, 0x7b, 0x73, 0x6c, 0x75, 0x67, 0x7d, 0x7d, 0x12, 0x25, 0x66,
	0x65, 0x61, 0x74, 0x28, 0x7b, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x7d, 0x29, 0x3a, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x7b, 0x7b, 0x73, 0x6c,
	0x75, 0x67, 0x7d, 0x7d, 0x1a, 0x25, 0x66, 0x65, 0x61, 0x74, 0x28, 0x7b, 0x7b, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x7d, 0x29, 0x3a, 0x20, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x20, 0x7b, 0x7b, 0x73, 0x6c, 0x75, 0x67, 0x7d, 0x7d, 0x22, 0x1c, 0x66, 0x65, 0x61,
	0x74, 0x28, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x29, 0x3a, 0x20, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x20, 0x7b, 0x7b, 0x70, 0x61, 0x74, 0x68, 0x7d, 0x7d, 0x2a, 0x1c, 0x66, 0x65, 0x61, 0x74, 0x28,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x29, 0x3a, 0x20, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x7b,
	0x7b, 0x70, 0x61, 0x74, 0x68, 0x7d, 0x7d, 0x12, 0x1e, 0x0a, 0x1c, 0x68, 0x74, 0x74, 0x70, 0x3a,
	0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x31,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x22, 0x0f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2a, 0x09, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x2e,
	0x73, 0x76, 0x67, 0x32, 0x07, 0x08, 0x02, 0x10, 0x01, 0x1a, 0x01, 0x2d, 0x0a, 0x22, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66,
	0x79, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x42, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x5a, 0x67, 0x6f, 0x2e, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x63, 0x68,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x6e, 0x65, 0x74, 0x6c, 0x69, 0x66,
	0x79, 0x2d, 0x63, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x65, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x6e, 0x65, 0x74, 0x6c, 0x69,
	0x66, 0x79, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x45, 0x4e,
	0x43, 0x45, 0xaa, 0x02, 0x1e, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x4e, 0x65, 0x74,
	0x6c, 0x69, 0x66, 0x79, 0x2e, 0x43, 0x6d, 0x73, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x4e, 0x65,
	0x74, 0x6c, 0x69, 0x66, 0x79, 0x5c, 0x43, 0x6d, 0x73, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2a, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x5c, 0x4e,
	0x65, 0x74, 0x6c, 0x69, 0x66, 0x79, 0x5c, 0x43, 0x6d, 0x73, 0x5c, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x22, 0x45, 0x69, 0x6e, 0x72, 0x69, 0x64, 0x65, 0x3a, 0x3a, 0x4e, 0x65, 0x74,
	0x6c, 0x69, 0x66, 0x79, 0x3a, 0x3a, 0x43, 0x6d, 0x73, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_einride_netlify_cms_example_v1_config_proto_goTypes = []interface{}{}
var file_einride_netlify_cms_example_v1_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_einride_netlify_cms_example_v1_config_proto_init() }
func file_einride_netlify_cms_example_v1_config_proto_init() {
	if File_einride_netlify_cms_example_v1_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_einride_netlify_cms_example_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_einride_netlify_cms_example_v1_config_proto_goTypes,
		DependencyIndexes: file_einride_netlify_cms_example_v1_config_proto_depIdxs,
	}.Build()
	File_einride_netlify_cms_example_v1_config_proto = out.File
	file_einride_netlify_cms_example_v1_config_proto_rawDesc = nil
	file_einride_netlify_cms_example_v1_config_proto_goTypes = nil
	file_einride_netlify_cms_example_v1_config_proto_depIdxs = nil
}
