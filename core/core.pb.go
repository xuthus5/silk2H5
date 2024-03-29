// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: core.proto

package core

import (
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

type TransType int32

const (
	TransType_TransTypeNil TransType = 0
	TransType_TransTypeMP3 TransType = 1
	TransType_TransTypeWAV TransType = 2
)

// Enum value maps for TransType.
var (
	TransType_name = map[int32]string{
		0: "TransTypeNil",
		1: "TransTypeMP3",
		2: "TransTypeWAV",
	}
	TransType_value = map[string]int32{
		"TransTypeNil": 0,
		"TransTypeMP3": 1,
		"TransTypeWAV": 2,
	}
)

func (x TransType) Enum() *TransType {
	p := new(TransType)
	*p = x
	return p
}

func (x TransType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransType) Descriptor() protoreflect.EnumDescriptor {
	return file_core_proto_enumTypes[0].Descriptor()
}

func (TransType) Type() protoreflect.EnumType {
	return &file_core_proto_enumTypes[0]
}

func (x TransType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransType.Descriptor instead.
func (TransType) EnumDescriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x69,
	0x6c, 0x6b, 0x32, 0x48, 0x35, 0x2a, 0x41, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x4e,
	0x69, 0x6c, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x4d, 0x50, 0x33, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x57, 0x41, 0x56, 0x10, 0x02, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x73,
	0x69, 0x6c, 0x6b, 0x32, 0x48, 0x35, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_proto_goTypes = []interface{}{
	(TransType)(0), // 0: silk2H5.TransType
}
var file_core_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		EnumInfos:         file_core_proto_enumTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}
