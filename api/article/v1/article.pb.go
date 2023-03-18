// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: article/v1/article.proto

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

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_v1_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_v1_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_article_v1_article_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*ArticleListItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Count int64              `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Page  int32              `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Size  int32              `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_v1_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_v1_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_article_v1_article_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetList() []*ArticleListItem {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListResponse) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ArticleListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title           string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author          string                 `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Avatar          string                 `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Summary         string                 `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	Cover           string                 `protobuf:"bytes,6,opt,name=cover,proto3" json:"cover,omitempty"`
	Time            string                 `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	Category        []*ArticleCategoryItem `protobuf:"bytes,8,rep,name=category,proto3" json:"category,omitempty"`
	ViewCount       int32                  `protobuf:"varint,9,opt,name=viewCount,proto3" json:"viewCount,omitempty"`
	CollectionCount int32                  `protobuf:"varint,10,opt,name=collectionCount,proto3" json:"collectionCount,omitempty"`
	ShareCount      int32                  `protobuf:"varint,11,opt,name=shareCount,proto3" json:"shareCount,omitempty"`
	ZanCount        int32                  `protobuf:"varint,12,opt,name=zanCount,proto3" json:"zanCount,omitempty"`
	CommentCount    int32                  `protobuf:"varint,13,opt,name=commentCount,proto3" json:"commentCount,omitempty"`
}

func (x *ArticleListItem) Reset() {
	*x = ArticleListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_v1_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleListItem) ProtoMessage() {}

func (x *ArticleListItem) ProtoReflect() protoreflect.Message {
	mi := &file_article_v1_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleListItem.ProtoReflect.Descriptor instead.
func (*ArticleListItem) Descriptor() ([]byte, []int) {
	return file_article_v1_article_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleListItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArticleListItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleListItem) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *ArticleListItem) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ArticleListItem) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *ArticleListItem) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *ArticleListItem) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *ArticleListItem) GetCategory() []*ArticleCategoryItem {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ArticleListItem) GetViewCount() int32 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

func (x *ArticleListItem) GetCollectionCount() int32 {
	if x != nil {
		return x.CollectionCount
	}
	return 0
}

func (x *ArticleListItem) GetShareCount() int32 {
	if x != nil {
		return x.ShareCount
	}
	return 0
}

func (x *ArticleListItem) GetZanCount() int32 {
	if x != nil {
		return x.ZanCount
	}
	return 0
}

func (x *ArticleListItem) GetCommentCount() int32 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_v1_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_article_v1_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_article_v1_article_proto_rawDescGZIP(), []int{3}
}

func (x *InfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title              string                     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author             string                     `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Avatar             string                     `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Summary            string                     `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	Cover              string                     `protobuf:"bytes,6,opt,name=cover,proto3" json:"cover,omitempty"`
	Time               string                     `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	Date               string                     `protobuf:"bytes,8,opt,name=date,proto3" json:"date,omitempty"`
	Category           []*ArticleCategoryItem     `protobuf:"bytes,9,rep,name=category,proto3" json:"category,omitempty"`
	ViewCount          int32                      `protobuf:"varint,10,opt,name=viewCount,proto3" json:"viewCount,omitempty"`
	CollectionCount    int32                      `protobuf:"varint,11,opt,name=collectionCount,proto3" json:"collectionCount,omitempty"`
	ShareCount         int32                      `protobuf:"varint,12,opt,name=shareCount,proto3" json:"shareCount,omitempty"`
	ZanCount           int32                      `protobuf:"varint,13,opt,name=zanCount,proto3" json:"zanCount,omitempty"`
	CommentCount       int32                      `protobuf:"varint,14,opt,name=commentCount,proto3" json:"commentCount,omitempty"`
	Source             string                     `protobuf:"bytes,15,opt,name=source,proto3" json:"source,omitempty"`
	SourceUrl          string                     `protobuf:"bytes,16,opt,name=sourceUrl,proto3" json:"sourceUrl,omitempty"`
	Content            string                     `protobuf:"bytes,17,opt,name=content,proto3" json:"content,omitempty"`
	PrevArticle        *BasicIntroduceArticleItem `protobuf:"bytes,18,opt,name=prevArticle,proto3" json:"prevArticle,omitempty"`
	NextArticle        *BasicIntroduceArticleItem `protobuf:"bytes,19,opt,name=nextArticle,proto3" json:"nextArticle,omitempty"`
	CopyrightAuthor    string                     `protobuf:"bytes,20,opt,name=copyrightAuthor,proto3" json:"copyrightAuthor,omitempty"`
	CopyrightArticleId string                     `protobuf:"bytes,21,opt,name=copyrightArticleId,proto3" json:"copyrightArticleId,omitempty"`
	CopyrightLink      string                     `protobuf:"bytes,22,opt,name=copyrightLink,proto3" json:"copyrightLink,omitempty"`
	CopyrightStatement string                     `protobuf:"bytes,23,opt,name=copyrightStatement,proto3" json:"copyrightStatement,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_v1_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_article_v1_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_article_v1_article_proto_rawDescGZIP(), []int{4}
}

func (x *InfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InfoResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InfoResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *InfoResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *InfoResponse) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *InfoResponse) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *InfoResponse) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *InfoResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *InfoResponse) GetCategory() []*ArticleCategoryItem {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *InfoResponse) GetViewCount() int32 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

func (x *InfoResponse) GetCollectionCount() int32 {
	if x != nil {
		return x.CollectionCount
	}
	return 0
}

func (x *InfoResponse) GetShareCount() int32 {
	if x != nil {
		return x.ShareCount
	}
	return 0
}

func (x *InfoResponse) GetZanCount() int32 {
	if x != nil {
		return x.ZanCount
	}
	return 0
}

func (x *InfoResponse) GetCommentCount() int32 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *InfoResponse) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *InfoResponse) GetSourceUrl() string {
	if x != nil {
		return x.SourceUrl
	}
	return ""
}

func (x *InfoResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *InfoResponse) GetPrevArticle() *BasicIntroduceArticleItem {
	if x != nil {
		return x.PrevArticle
	}
	return nil
}

func (x *InfoResponse) GetNextArticle() *BasicIntroduceArticleItem {
	if x != nil {
		return x.NextArticle
	}
	return nil
}

func (x *InfoResponse) GetCopyrightAuthor() string {
	if x != nil {
		return x.CopyrightAuthor
	}
	return ""
}

func (x *InfoResponse) GetCopyrightArticleId() string {
	if x != nil {
		return x.CopyrightArticleId
	}
	return ""
}

func (x *InfoResponse) GetCopyrightLink() string {
	if x != nil {
		return x.CopyrightLink
	}
	return ""
}

func (x *InfoResponse) GetCopyrightStatement() string {
	if x != nil {
		return x.CopyrightStatement
	}
	return ""
}

type ArticleCategoryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Color string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *ArticleCategoryItem) Reset() {
	*x = ArticleCategoryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_v1_article_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleCategoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCategoryItem) ProtoMessage() {}

func (x *ArticleCategoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_article_v1_article_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCategoryItem.ProtoReflect.Descriptor instead.
func (*ArticleCategoryItem) Descriptor() ([]byte, []int) {
	return file_article_v1_article_proto_rawDescGZIP(), []int{5}
}

func (x *ArticleCategoryItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ArticleCategoryItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ArticleCategoryItem) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type BasicIntroduceArticleItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *BasicIntroduceArticleItem) Reset() {
	*x = BasicIntroduceArticleItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_article_v1_article_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicIntroduceArticleItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicIntroduceArticleItem) ProtoMessage() {}

func (x *BasicIntroduceArticleItem) ProtoReflect() protoreflect.Message {
	mi := &file_article_v1_article_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicIntroduceArticleItem.ProtoReflect.Descriptor instead.
func (*BasicIntroduceArticleItem) Descriptor() ([]byte, []int) {
	return file_article_v1_article_proto_rawDescGZIP(), []int{6}
}

func (x *BasicIntroduceArticleItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BasicIntroduceArticleItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_article_v1_article_proto protoreflect.FileDescriptor

var file_article_v1_article_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a,
	0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x64,
	0x20, 0x00, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x7d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x90, 0x03, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x7a, 0x61,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x7a, 0x61,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1d, 0x0a, 0x0b, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb3, 0x06, 0x0a, 0x0c, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x7a, 0x61, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x7a, 0x61, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55,
	0x72, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x47, 0x0a,
	0x0b, 0x70, 0x72, 0x65, 0x76, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x76, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e,
	0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x70,
	0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x70,
	0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x70,
	0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x4f, 0x0a, 0x13, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x22, 0x41, 0x0a, 0x19, 0x42, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x32, 0xa2, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x48, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x07, 0x12, 0x05, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x17, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x69,
	0x6e, 0x66, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x13, 0x5a, 0x11, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_article_v1_article_proto_rawDescOnce sync.Once
	file_article_v1_article_proto_rawDescData = file_article_v1_article_proto_rawDesc
)

func file_article_v1_article_proto_rawDescGZIP() []byte {
	file_article_v1_article_proto_rawDescOnce.Do(func() {
		file_article_v1_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_article_v1_article_proto_rawDescData)
	})
	return file_article_v1_article_proto_rawDescData
}

var file_article_v1_article_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_article_v1_article_proto_goTypes = []interface{}{
	(*ListRequest)(nil),               // 0: article.v1.ListRequest
	(*ListResponse)(nil),              // 1: article.v1.ListResponse
	(*ArticleListItem)(nil),           // 2: article.v1.ArticleListItem
	(*InfoRequest)(nil),               // 3: article.v1.InfoRequest
	(*InfoResponse)(nil),              // 4: article.v1.InfoResponse
	(*ArticleCategoryItem)(nil),       // 5: article.v1.ArticleCategoryItem
	(*BasicIntroduceArticleItem)(nil), // 6: article.v1.BasicIntroduceArticleItem
}
var file_article_v1_article_proto_depIdxs = []int32{
	2, // 0: article.v1.ListResponse.list:type_name -> article.v1.ArticleListItem
	5, // 1: article.v1.ArticleListItem.category:type_name -> article.v1.ArticleCategoryItem
	5, // 2: article.v1.InfoResponse.category:type_name -> article.v1.ArticleCategoryItem
	6, // 3: article.v1.InfoResponse.prevArticle:type_name -> article.v1.BasicIntroduceArticleItem
	6, // 4: article.v1.InfoResponse.nextArticle:type_name -> article.v1.BasicIntroduceArticleItem
	0, // 5: article.v1.Article.List:input_type -> article.v1.ListRequest
	3, // 6: article.v1.Article.Info:input_type -> article.v1.InfoRequest
	1, // 7: article.v1.Article.List:output_type -> article.v1.ListResponse
	4, // 8: article.v1.Article.Info:output_type -> article.v1.InfoResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_article_v1_article_proto_init() }
func file_article_v1_article_proto_init() {
	if File_article_v1_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_article_v1_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_article_v1_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_article_v1_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleListItem); i {
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
		file_article_v1_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_article_v1_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
		file_article_v1_article_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleCategoryItem); i {
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
		file_article_v1_article_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicIntroduceArticleItem); i {
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
			RawDescriptor: file_article_v1_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_article_v1_article_proto_goTypes,
		DependencyIndexes: file_article_v1_article_proto_depIdxs,
		MessageInfos:      file_article_v1_article_proto_msgTypes,
	}.Build()
	File_article_v1_article_proto = out.File
	file_article_v1_article_proto_rawDesc = nil
	file_article_v1_article_proto_goTypes = nil
	file_article_v1_article_proto_depIdxs = nil
}
