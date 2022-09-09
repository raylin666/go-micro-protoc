// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: basic/v1/upload.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type StreamUploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream   []byte `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *StreamUploadFileRequest) Reset() {
	*x = StreamUploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_v1_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamUploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamUploadFileRequest) ProtoMessage() {}

func (x *StreamUploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basic_v1_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamUploadFileRequest.ProtoReflect.Descriptor instead.
func (*StreamUploadFileRequest) Descriptor() ([]byte, []int) {
	return file_basic_v1_upload_proto_rawDescGZIP(), []int{0}
}

func (x *StreamUploadFileRequest) GetStream() []byte {
	if x != nil {
		return x.Stream
	}
	return nil
}

func (x *StreamUploadFileRequest) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

type StreamUploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash     string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Uuid     string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Key      string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Fsize    int64  `protobuf:"varint,4,opt,name=fsize,proto3" json:"fsize,omitempty"`
	Url      string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Name     string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Bucket   string `protobuf:"bytes,7,opt,name=bucket,proto3" json:"bucket,omitempty"`
	MimeType string `protobuf:"bytes,8,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Ext      string `protobuf:"bytes,9,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *StreamUploadFileResponse) Reset() {
	*x = StreamUploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_v1_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamUploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamUploadFileResponse) ProtoMessage() {}

func (x *StreamUploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basic_v1_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamUploadFileResponse.ProtoReflect.Descriptor instead.
func (*StreamUploadFileResponse) Descriptor() ([]byte, []int) {
	return file_basic_v1_upload_proto_rawDescGZIP(), []int{1}
}

func (x *StreamUploadFileResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *StreamUploadFileResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *StreamUploadFileResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StreamUploadFileResponse) GetFsize() int64 {
	if x != nil {
		return x.Fsize
	}
	return 0
}

func (x *StreamUploadFileResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StreamUploadFileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamUploadFileResponse) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *StreamUploadFileResponse) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *StreamUploadFileResponse) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type UrlUploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *UrlUploadFileRequest) Reset() {
	*x = UrlUploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_v1_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlUploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlUploadFileRequest) ProtoMessage() {}

func (x *UrlUploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basic_v1_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlUploadFileRequest.ProtoReflect.Descriptor instead.
func (*UrlUploadFileRequest) Descriptor() ([]byte, []int) {
	return file_basic_v1_upload_proto_rawDescGZIP(), []int{2}
}

func (x *UrlUploadFileRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UrlUploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash     string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Key      string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Fsize    int64  `protobuf:"varint,4,opt,name=fsize,proto3" json:"fsize,omitempty"`
	Url      string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	MimeType string `protobuf:"bytes,8,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *UrlUploadFileResponse) Reset() {
	*x = UrlUploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_v1_upload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlUploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlUploadFileResponse) ProtoMessage() {}

func (x *UrlUploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basic_v1_upload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlUploadFileResponse.ProtoReflect.Descriptor instead.
func (*UrlUploadFileResponse) Descriptor() ([]byte, []int) {
	return file_basic_v1_upload_proto_rawDescGZIP(), []int{3}
}

func (x *UrlUploadFileResponse) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *UrlUploadFileResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UrlUploadFileResponse) GetFsize() int64 {
	if x != nil {
		return x.Fsize
	}
	return 0
}

func (x *UrlUploadFileResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UrlUploadFileResponse) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

var File_basic_v1_upload_proto protoreflect.FileDescriptor

var file_basic_v1_upload_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x17, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02, 0x10, 0x0a, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x24, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03,
	0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x18, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x66, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x78, 0x74, 0x22, 0x32, 0x0a, 0x14, 0x55, 0x72, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0x88, 0x01, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x55, 0x72, 0x6c,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x32, 0xf2, 0x01,
	0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x79, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x1a, 0x13, 0x2f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x3a, 0x01, 0x2a, 0x12, 0x6d, 0x0a, 0x0d, 0x55, 0x72, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x72, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x72, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x72, 0x6c, 0x3a,
	0x01, 0x2a, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basic_v1_upload_proto_rawDescOnce sync.Once
	file_basic_v1_upload_proto_rawDescData = file_basic_v1_upload_proto_rawDesc
)

func file_basic_v1_upload_proto_rawDescGZIP() []byte {
	file_basic_v1_upload_proto_rawDescOnce.Do(func() {
		file_basic_v1_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_basic_v1_upload_proto_rawDescData)
	})
	return file_basic_v1_upload_proto_rawDescData
}

var file_basic_v1_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_basic_v1_upload_proto_goTypes = []interface{}{
	(*StreamUploadFileRequest)(nil),  // 0: basic.v1.StreamUploadFileRequest
	(*StreamUploadFileResponse)(nil), // 1: basic.v1.StreamUploadFileResponse
	(*UrlUploadFileRequest)(nil),     // 2: basic.v1.UrlUploadFileRequest
	(*UrlUploadFileResponse)(nil),    // 3: basic.v1.UrlUploadFileResponse
}
var file_basic_v1_upload_proto_depIdxs = []int32{
	0, // 0: basic.v1.Upload.StreamUploadFile:input_type -> basic.v1.StreamUploadFileRequest
	2, // 1: basic.v1.Upload.UrlUploadFile:input_type -> basic.v1.UrlUploadFileRequest
	1, // 2: basic.v1.Upload.StreamUploadFile:output_type -> basic.v1.StreamUploadFileResponse
	3, // 3: basic.v1.Upload.UrlUploadFile:output_type -> basic.v1.UrlUploadFileResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_basic_v1_upload_proto_init() }
func file_basic_v1_upload_proto_init() {
	if File_basic_v1_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basic_v1_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamUploadFileRequest); i {
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
		file_basic_v1_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamUploadFileResponse); i {
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
		file_basic_v1_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlUploadFileRequest); i {
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
		file_basic_v1_upload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlUploadFileResponse); i {
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
			RawDescriptor: file_basic_v1_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basic_v1_upload_proto_goTypes,
		DependencyIndexes: file_basic_v1_upload_proto_depIdxs,
		MessageInfos:      file_basic_v1_upload_proto_msgTypes,
	}.Build()
	File_basic_v1_upload_proto = out.File
	file_basic_v1_upload_proto_rawDesc = nil
	file_basic_v1_upload_proto_goTypes = nil
	file_basic_v1_upload_proto_depIdxs = nil
}