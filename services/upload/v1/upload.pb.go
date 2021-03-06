// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: services/upload/v1/upload.proto

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
		mi := &file_services_upload_v1_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamUploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamUploadFileRequest) ProtoMessage() {}

func (x *StreamUploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_upload_v1_upload_proto_msgTypes[0]
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
	return file_services_upload_v1_upload_proto_rawDescGZIP(), []int{0}
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

type StreamUploadFileReply struct {
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

func (x *StreamUploadFileReply) Reset() {
	*x = StreamUploadFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_upload_v1_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamUploadFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamUploadFileReply) ProtoMessage() {}

func (x *StreamUploadFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_upload_v1_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamUploadFileReply.ProtoReflect.Descriptor instead.
func (*StreamUploadFileReply) Descriptor() ([]byte, []int) {
	return file_services_upload_v1_upload_proto_rawDescGZIP(), []int{1}
}

func (x *StreamUploadFileReply) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *StreamUploadFileReply) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *StreamUploadFileReply) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StreamUploadFileReply) GetFsize() int64 {
	if x != nil {
		return x.Fsize
	}
	return 0
}

func (x *StreamUploadFileReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *StreamUploadFileReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamUploadFileReply) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *StreamUploadFileReply) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *StreamUploadFileReply) GetExt() string {
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
		mi := &file_services_upload_v1_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlUploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlUploadFileRequest) ProtoMessage() {}

func (x *UrlUploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_upload_v1_upload_proto_msgTypes[2]
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
	return file_services_upload_v1_upload_proto_rawDescGZIP(), []int{2}
}

func (x *UrlUploadFileRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type UrlUploadFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash     string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Key      string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Fsize    int64  `protobuf:"varint,4,opt,name=fsize,proto3" json:"fsize,omitempty"`
	Url      string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	MimeType string `protobuf:"bytes,8,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *UrlUploadFileReply) Reset() {
	*x = UrlUploadFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_upload_v1_upload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlUploadFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlUploadFileReply) ProtoMessage() {}

func (x *UrlUploadFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_services_upload_v1_upload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlUploadFileReply.ProtoReflect.Descriptor instead.
func (*UrlUploadFileReply) Descriptor() ([]byte, []int) {
	return file_services_upload_v1_upload_proto_rawDescGZIP(), []int{3}
}

func (x *UrlUploadFileReply) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *UrlUploadFileReply) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UrlUploadFileReply) GetFsize() int64 {
	if x != nil {
		return x.Fsize
	}
	return 0
}

func (x *UrlUploadFileReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UrlUploadFileReply) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

var File_services_upload_v1_upload_proto protoreflect.FileDescriptor

var file_services_upload_v1_upload_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x17,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02, 0x10, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x24, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd4,
	0x01, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x66, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x78, 0x74, 0x22, 0x32, 0x0a, 0x14, 0x55, 0x72, 0x6c, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0x88, 0x01, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x7f, 0x0a, 0x12, 0x55, 0x72, 0x6c,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x66, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x32, 0x87, 0x02, 0x0a, 0x06, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x83, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x1a, 0x0c, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a, 0x0d, 0x55,
	0x72, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x72, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x72, 0x6c, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x72,
	0x6c, 0x3a, 0x01, 0x2a, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x79, 0x6c, 0x69, 0x6e, 0x36, 0x36, 0x36, 0x2f, 0x67, 0x6f, 0x2d,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_upload_v1_upload_proto_rawDescOnce sync.Once
	file_services_upload_v1_upload_proto_rawDescData = file_services_upload_v1_upload_proto_rawDesc
)

func file_services_upload_v1_upload_proto_rawDescGZIP() []byte {
	file_services_upload_v1_upload_proto_rawDescOnce.Do(func() {
		file_services_upload_v1_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_upload_v1_upload_proto_rawDescData)
	})
	return file_services_upload_v1_upload_proto_rawDescData
}

var file_services_upload_v1_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_services_upload_v1_upload_proto_goTypes = []interface{}{
	(*StreamUploadFileRequest)(nil), // 0: services.upload.v1.StreamUploadFileRequest
	(*StreamUploadFileReply)(nil),   // 1: services.upload.v1.StreamUploadFileReply
	(*UrlUploadFileRequest)(nil),    // 2: services.upload.v1.UrlUploadFileRequest
	(*UrlUploadFileReply)(nil),      // 3: services.upload.v1.UrlUploadFileReply
}
var file_services_upload_v1_upload_proto_depIdxs = []int32{
	0, // 0: services.upload.v1.Upload.StreamUploadFile:input_type -> services.upload.v1.StreamUploadFileRequest
	2, // 1: services.upload.v1.Upload.UrlUploadFile:input_type -> services.upload.v1.UrlUploadFileRequest
	1, // 2: services.upload.v1.Upload.StreamUploadFile:output_type -> services.upload.v1.StreamUploadFileReply
	3, // 3: services.upload.v1.Upload.UrlUploadFile:output_type -> services.upload.v1.UrlUploadFileReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_upload_v1_upload_proto_init() }
func file_services_upload_v1_upload_proto_init() {
	if File_services_upload_v1_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_upload_v1_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_upload_v1_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamUploadFileReply); i {
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
		file_services_upload_v1_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_upload_v1_upload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlUploadFileReply); i {
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
			RawDescriptor: file_services_upload_v1_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_upload_v1_upload_proto_goTypes,
		DependencyIndexes: file_services_upload_v1_upload_proto_depIdxs,
		MessageInfos:      file_services_upload_v1_upload_proto_msgTypes,
	}.Build()
	File_services_upload_v1_upload_proto = out.File
	file_services_upload_v1_upload_proto_rawDesc = nil
	file_services_upload_v1_upload_proto_goTypes = nil
	file_services_upload_v1_upload_proto_depIdxs = nil
}
