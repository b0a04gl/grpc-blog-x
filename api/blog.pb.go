// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: api/blog.proto

package api

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

type Blog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Blog) Reset() {
	*x = Blog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blog) ProtoMessage() {}

func (x *Blog) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blog.ProtoReflect.Descriptor instead.
func (*Blog) Descriptor() ([]byte, []int) {
	return file_api_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Blog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Blog) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type BlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *BlogRequest) Reset() {
	*x = BlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogRequest) ProtoMessage() {}

func (x *BlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogRequest.ProtoReflect.Descriptor instead.
func (*BlogRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_proto_rawDescGZIP(), []int{1}
}

func (x *BlogRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type BlogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BlogResponse) Reset() {
	*x = BlogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogResponse) ProtoMessage() {}

func (x *BlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogResponse.ProtoReflect.Descriptor instead.
func (*BlogResponse) Descriptor() ([]byte, []int) {
	return file_api_blog_proto_rawDescGZIP(), []int{2}
}

func (x *BlogResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBlogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBlogRequest) Reset() {
	*x = GetBlogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlogRequest) ProtoMessage() {}

func (x *GetBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlogRequest.ProtoReflect.Descriptor instead.
func (*GetBlogRequest) Descriptor() ([]byte, []int) {
	return file_api_blog_proto_rawDescGZIP(), []int{3}
}

func (x *GetBlogRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_blog_proto protoreflect.FileDescriptor

var file_api_blog_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x22, 0x46, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3d,
	0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x1e, 0x0a,
	0x0c, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0x6f, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x11, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x14,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x6c, 0x6f, 0x67,
	0x42, 0x1a, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x52, 0x50, 0x43, 0x42, 0x6c, 0x6f, 0x67, 0x58, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_blog_proto_rawDescOnce sync.Once
	file_api_blog_proto_rawDescData = file_api_blog_proto_rawDesc
)

func file_api_blog_proto_rawDescGZIP() []byte {
	file_api_blog_proto_rawDescOnce.Do(func() {
		file_api_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_blog_proto_rawDescData)
	})
	return file_api_blog_proto_rawDescData
}

var file_api_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_blog_proto_goTypes = []interface{}{
	(*Blog)(nil),           // 0: blog.Blog
	(*BlogRequest)(nil),    // 1: blog.BlogRequest
	(*BlogResponse)(nil),   // 2: blog.BlogResponse
	(*GetBlogRequest)(nil), // 3: blog.GetBlogRequest
}
var file_api_blog_proto_depIdxs = []int32{
	1, // 0: blog.BlogService.CreateBlog:input_type -> blog.BlogRequest
	3, // 1: blog.BlogService.GetBlog:input_type -> blog.GetBlogRequest
	2, // 2: blog.BlogService.CreateBlog:output_type -> blog.BlogResponse
	0, // 3: blog.BlogService.GetBlog:output_type -> blog.Blog
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_blog_proto_init() }
func file_api_blog_proto_init() {
	if File_api_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blog); i {
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
		file_api_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogRequest); i {
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
		file_api_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogResponse); i {
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
		file_api_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlogRequest); i {
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
			RawDescriptor: file_api_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_blog_proto_goTypes,
		DependencyIndexes: file_api_blog_proto_depIdxs,
		MessageInfos:      file_api_blog_proto_msgTypes,
	}.Build()
	File_api_blog_proto = out.File
	file_api_blog_proto_rawDesc = nil
	file_api_blog_proto_goTypes = nil
	file_api_blog_proto_depIdxs = nil
}
