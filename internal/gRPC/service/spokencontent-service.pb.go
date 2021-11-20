// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/proto-files/service/spokencontent-service.proto

package service

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

// CRUD
type AddAudioRecordingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddAudioRecordingResponse) Reset() {
	*x = AddAudioRecordingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_files_service_spokencontent_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAudioRecordingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAudioRecordingResponse) ProtoMessage() {}

func (x *AddAudioRecordingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_files_service_spokencontent_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAudioRecordingResponse.ProtoReflect.Descriptor instead.
func (*AddAudioRecordingResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_files_service_spokencontent_service_proto_rawDescGZIP(), []int{0}
}

type UpdateRecordingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRecordingResponse) Reset() {
	*x = UpdateRecordingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_files_service_spokencontent_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRecordingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRecordingResponse) ProtoMessage() {}

func (x *UpdateRecordingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_files_service_spokencontent_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRecordingResponse.ProtoReflect.Descriptor instead.
func (*UpdateRecordingResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_files_service_spokencontent_service_proto_rawDescGZIP(), []int{1}
}

type DeleteAudioRecordingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAudioRecordingRequest) Reset() {
	*x = DeleteAudioRecordingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_files_service_spokencontent_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAudioRecordingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAudioRecordingRequest) ProtoMessage() {}

func (x *DeleteAudioRecordingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_files_service_spokencontent_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAudioRecordingRequest.ProtoReflect.Descriptor instead.
func (*DeleteAudioRecordingRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_files_service_spokencontent_service_proto_rawDescGZIP(), []int{2}
}

type DeleteAudioRecordingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAudioRecordingResponse) Reset() {
	*x = DeleteAudioRecordingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_files_service_spokencontent_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAudioRecordingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAudioRecordingResponse) ProtoMessage() {}

func (x *DeleteAudioRecordingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_files_service_spokencontent_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAudioRecordingResponse.ProtoReflect.Descriptor instead.
func (*DeleteAudioRecordingResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_files_service_spokencontent_service_proto_rawDescGZIP(), []int{3}
}

var File_internal_proto_files_service_spokencontent_service_proto protoreflect.FileDescriptor

var file_internal_proto_files_service_spokencontent_service_proto_rawDesc = []byte{
	0x0a, 0x38, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73,
	0x70, 0x6f, 0x6b, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x73, 0x70, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x19, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x0a, 0x1b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x1c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf0, 0x01, 0x0a, 0x10,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x41, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x16, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x1a,
	0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x75, 0x64,
	0x69, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x20, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x17,
	0x5a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_files_service_spokencontent_service_proto_rawDescOnce sync.Once
	file_internal_proto_files_service_spokencontent_service_proto_rawDescData = file_internal_proto_files_service_spokencontent_service_proto_rawDesc
)

func file_internal_proto_files_service_spokencontent_service_proto_rawDescGZIP() []byte {
	file_internal_proto_files_service_spokencontent_service_proto_rawDescOnce.Do(func() {
		file_internal_proto_files_service_spokencontent_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_files_service_spokencontent_service_proto_rawDescData)
	})
	return file_internal_proto_files_service_spokencontent_service_proto_rawDescData
}

var file_internal_proto_files_service_spokencontent_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_proto_files_service_spokencontent_service_proto_goTypes = []interface{}{
	(*AddAudioRecordingResponse)(nil),    // 0: service.AddAudioRecordingResponse
	(*UpdateRecordingResponse)(nil),      // 1: service.UpdateRecordingResponse
	(*DeleteAudioRecordingRequest)(nil),  // 2: service.DeleteAudioRecordingRequest
	(*DeleteAudioRecordingResponse)(nil), // 3: service.DeleteAudioRecordingResponse
	(*AudioRecording)(nil),               // 4: domain.AudioRecording
}
var file_internal_proto_files_service_spokencontent_service_proto_depIdxs = []int32{
	4, // 0: service.RecordingService.add:input_type -> domain.AudioRecording
	4, // 1: service.RecordingService.update:input_type -> domain.AudioRecording
	2, // 2: service.RecordingService.delete:input_type -> service.DeleteAudioRecordingRequest
	0, // 3: service.RecordingService.add:output_type -> service.AddAudioRecordingResponse
	1, // 4: service.RecordingService.update:output_type -> service.UpdateRecordingResponse
	3, // 5: service.RecordingService.delete:output_type -> service.DeleteAudioRecordingResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_files_service_spokencontent_service_proto_init() }
func file_internal_proto_files_service_spokencontent_service_proto_init() {
	if File_internal_proto_files_service_spokencontent_service_proto != nil {
		return
	}
	file_internal_proto_files_domain_spokenContent_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_files_service_spokencontent_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAudioRecordingResponse); i {
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
		file_internal_proto_files_service_spokencontent_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRecordingResponse); i {
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
		file_internal_proto_files_service_spokencontent_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAudioRecordingRequest); i {
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
		file_internal_proto_files_service_spokencontent_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAudioRecordingResponse); i {
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
			RawDescriptor: file_internal_proto_files_service_spokencontent_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_files_service_spokencontent_service_proto_goTypes,
		DependencyIndexes: file_internal_proto_files_service_spokencontent_service_proto_depIdxs,
		MessageInfos:      file_internal_proto_files_service_spokencontent_service_proto_msgTypes,
	}.Build()
	File_internal_proto_files_service_spokencontent_service_proto = out.File
	file_internal_proto_files_service_spokencontent_service_proto_rawDesc = nil
	file_internal_proto_files_service_spokencontent_service_proto_goTypes = nil
	file_internal_proto_files_service_spokencontent_service_proto_depIdxs = nil
}
