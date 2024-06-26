// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: s3.proto

package proto

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

type NoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParam) Reset() {
	*x = NoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParam) ProtoMessage() {}

func (x *NoParam) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParam.ProtoReflect.Descriptor instead.
func (*NoParam) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{0}
}

type UploadImageFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *ObjectParams `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Data []byte        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadImageFileRequest) Reset() {
	*x = UploadImageFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageFileRequest) ProtoMessage() {}

func (x *UploadImageFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageFileRequest.ProtoReflect.Descriptor instead.
func (*UploadImageFileRequest) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{1}
}

func (x *UploadImageFileRequest) GetMeta() *ObjectParams {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *UploadImageFileRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadImageURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *ObjectParams `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Url  string        `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UploadImageURLRequest) Reset() {
	*x = UploadImageURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageURLRequest) ProtoMessage() {}

func (x *UploadImageURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageURLRequest.ProtoReflect.Descriptor instead.
func (*UploadImageURLRequest) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{2}
}

func (x *UploadImageURLRequest) GetMeta() *ObjectParams {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *UploadImageURLRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ObjectParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectID   int32  `protobuf:"varint,1,opt,name=objectID,proto3" json:"objectID,omitempty"`
	BucketName string `protobuf:"bytes,2,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	Format     string `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"` // png jpeg etc..
}

func (x *ObjectParams) Reset() {
	*x = ObjectParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectParams) ProtoMessage() {}

func (x *ObjectParams) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectParams.ProtoReflect.Descriptor instead.
func (*ObjectParams) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectParams) GetObjectID() int32 {
	if x != nil {
		return x.ObjectID
	}
	return 0
}

func (x *ObjectParams) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *ObjectParams) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageURL string `protobuf:"bytes,1,opt,name=imageURL,proto3" json:"imageURL,omitempty"`
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{4}
}

func (x *UploadImageResponse) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

var File_s3_proto protoreflect.FileDescriptor

var file_s3_proto_rawDesc = []byte{
	0x0a, 0x08, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x33, 0x5f, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x22, 0x59, 0x0a, 0x16, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x33, 0x5f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x15, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x33, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x22, 0x62, 0x0a, 0x0c, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x31, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x32, 0xbd, 0x01, 0x0a, 0x09, 0x53,
	0x33, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x21, 0x2e, 0x73, 0x33, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x33, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x56, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x2e, 0x73, 0x33, 0x5f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x33, 0x5f,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s3_proto_rawDescOnce sync.Once
	file_s3_proto_rawDescData = file_s3_proto_rawDesc
)

func file_s3_proto_rawDescGZIP() []byte {
	file_s3_proto_rawDescOnce.Do(func() {
		file_s3_proto_rawDescData = protoimpl.X.CompressGZIP(file_s3_proto_rawDescData)
	})
	return file_s3_proto_rawDescData
}

var file_s3_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_s3_proto_goTypes = []interface{}{
	(*NoParam)(nil),                // 0: s3_bucket.NoParam
	(*UploadImageFileRequest)(nil), // 1: s3_bucket.UploadImageFileRequest
	(*UploadImageURLRequest)(nil),  // 2: s3_bucket.UploadImageURLRequest
	(*ObjectParams)(nil),           // 3: s3_bucket.ObjectParams
	(*UploadImageResponse)(nil),    // 4: s3_bucket.UploadImageResponse
}
var file_s3_proto_depIdxs = []int32{
	3, // 0: s3_bucket.UploadImageFileRequest.meta:type_name -> s3_bucket.ObjectParams
	3, // 1: s3_bucket.UploadImageURLRequest.meta:type_name -> s3_bucket.ObjectParams
	1, // 2: s3_bucket.S3Service.UploadImageFromFile:input_type -> s3_bucket.UploadImageFileRequest
	2, // 3: s3_bucket.S3Service.UploadImageFromURL:input_type -> s3_bucket.UploadImageURLRequest
	4, // 4: s3_bucket.S3Service.UploadImageFromFile:output_type -> s3_bucket.UploadImageResponse
	4, // 5: s3_bucket.S3Service.UploadImageFromURL:output_type -> s3_bucket.UploadImageResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_s3_proto_init() }
func file_s3_proto_init() {
	if File_s3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoParam); i {
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
		file_s3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageFileRequest); i {
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
		file_s3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageURLRequest); i {
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
		file_s3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectParams); i {
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
		file_s3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageResponse); i {
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
			RawDescriptor: file_s3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_s3_proto_goTypes,
		DependencyIndexes: file_s3_proto_depIdxs,
		MessageInfos:      file_s3_proto_msgTypes,
	}.Build()
	File_s3_proto = out.File
	file_s3_proto_rawDesc = nil
	file_s3_proto_goTypes = nil
	file_s3_proto_depIdxs = nil
}
