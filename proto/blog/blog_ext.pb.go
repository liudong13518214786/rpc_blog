// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go/src/rpc_blog/proto/blog/blog_ext.proto

package blog_ext

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//LIST
type BlogListRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogListRequest) Reset()         { *m = BlogListRequest{} }
func (m *BlogListRequest) String() string { return proto.CompactTextString(m) }
func (*BlogListRequest) ProtoMessage()    {}
func (*BlogListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{0}
}

func (m *BlogListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogListRequest.Unmarshal(m, b)
}
func (m *BlogListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogListRequest.Marshal(b, m, deterministic)
}
func (m *BlogListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogListRequest.Merge(m, src)
}
func (m *BlogListRequest) XXX_Size() int {
	return xxx_messageInfo_BlogListRequest.Size(m)
}
func (m *BlogListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlogListRequest proto.InternalMessageInfo

func (m *BlogListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type BlogListResponse struct {
	BlogList             []*BlogInfo `protobuf:"bytes,1,rep,name=blogList,proto3" json:"blogList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BlogListResponse) Reset()         { *m = BlogListResponse{} }
func (m *BlogListResponse) String() string { return proto.CompactTextString(m) }
func (*BlogListResponse) ProtoMessage()    {}
func (*BlogListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{1}
}

func (m *BlogListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogListResponse.Unmarshal(m, b)
}
func (m *BlogListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogListResponse.Marshal(b, m, deterministic)
}
func (m *BlogListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogListResponse.Merge(m, src)
}
func (m *BlogListResponse) XXX_Size() int {
	return xxx_messageInfo_BlogListResponse.Size(m)
}
func (m *BlogListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlogListResponse proto.InternalMessageInfo

func (m *BlogListResponse) GetBlogList() []*BlogInfo {
	if m != nil {
		return m.BlogList
	}
	return nil
}

type BlogInfo struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	BuildTime            string   `protobuf:"bytes,2,opt,name=buildTime,proto3" json:"buildTime,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogInfo) Reset()         { *m = BlogInfo{} }
func (m *BlogInfo) String() string { return proto.CompactTextString(m) }
func (*BlogInfo) ProtoMessage()    {}
func (*BlogInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{2}
}

func (m *BlogInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogInfo.Unmarshal(m, b)
}
func (m *BlogInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogInfo.Marshal(b, m, deterministic)
}
func (m *BlogInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogInfo.Merge(m, src)
}
func (m *BlogInfo) XXX_Size() int {
	return xxx_messageInfo_BlogInfo.Size(m)
}
func (m *BlogInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlogInfo proto.InternalMessageInfo

func (m *BlogInfo) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *BlogInfo) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

func (m *BlogInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

//WRITE
type WriteBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteBlogRequest) Reset()         { *m = WriteBlogRequest{} }
func (m *WriteBlogRequest) String() string { return proto.CompactTextString(m) }
func (*WriteBlogRequest) ProtoMessage()    {}
func (*WriteBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{3}
}

func (m *WriteBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteBlogRequest.Unmarshal(m, b)
}
func (m *WriteBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteBlogRequest.Marshal(b, m, deterministic)
}
func (m *WriteBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteBlogRequest.Merge(m, src)
}
func (m *WriteBlogRequest) XXX_Size() int {
	return xxx_messageInfo_WriteBlogRequest.Size(m)
}
func (m *WriteBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteBlogRequest proto.InternalMessageInfo

func (m *WriteBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

func (m *WriteBlogRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type WriteBlogResponse struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteBlogResponse) Reset()         { *m = WriteBlogResponse{} }
func (m *WriteBlogResponse) String() string { return proto.CompactTextString(m) }
func (*WriteBlogResponse) ProtoMessage()    {}
func (*WriteBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{4}
}

func (m *WriteBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteBlogResponse.Unmarshal(m, b)
}
func (m *WriteBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteBlogResponse.Marshal(b, m, deterministic)
}
func (m *WriteBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteBlogResponse.Merge(m, src)
}
func (m *WriteBlogResponse) XXX_Size() int {
	return xxx_messageInfo_WriteBlogResponse.Size(m)
}
func (m *WriteBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteBlogResponse proto.InternalMessageInfo

func (m *WriteBlogResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *WriteBlogResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//delete
type DeleteBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogRequest) Reset()         { *m = DeleteBlogRequest{} }
func (m *DeleteBlogRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogRequest) ProtoMessage()    {}
func (*DeleteBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{5}
}

func (m *DeleteBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogRequest.Unmarshal(m, b)
}
func (m *DeleteBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogRequest.Merge(m, src)
}
func (m *DeleteBlogRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogRequest.Size(m)
}
func (m *DeleteBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogRequest proto.InternalMessageInfo

func (m *DeleteBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

func (m *DeleteBlogRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteBlogResponse struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogResponse) Reset()         { *m = DeleteBlogResponse{} }
func (m *DeleteBlogResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogResponse) ProtoMessage()    {}
func (*DeleteBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{6}
}

func (m *DeleteBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogResponse.Unmarshal(m, b)
}
func (m *DeleteBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogResponse.Marshal(b, m, deterministic)
}
func (m *DeleteBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogResponse.Merge(m, src)
}
func (m *DeleteBlogResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogResponse.Size(m)
}
func (m *DeleteBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogResponse proto.InternalMessageInfo

func (m *DeleteBlogResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteBlogResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//modify
type ModifyBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Detail               string   `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyBlogRequest) Reset()         { *m = ModifyBlogRequest{} }
func (m *ModifyBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyBlogRequest) ProtoMessage()    {}
func (*ModifyBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{7}
}

func (m *ModifyBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyBlogRequest.Unmarshal(m, b)
}
func (m *ModifyBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyBlogRequest.Marshal(b, m, deterministic)
}
func (m *ModifyBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyBlogRequest.Merge(m, src)
}
func (m *ModifyBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ModifyBlogRequest.Size(m)
}
func (m *ModifyBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyBlogRequest proto.InternalMessageInfo

func (m *ModifyBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

func (m *ModifyBlogRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ModifyBlogRequest) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *ModifyBlogRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ModifyBlogResponse struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyBlogResponse) Reset()         { *m = ModifyBlogResponse{} }
func (m *ModifyBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyBlogResponse) ProtoMessage()    {}
func (*ModifyBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{8}
}

func (m *ModifyBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyBlogResponse.Unmarshal(m, b)
}
func (m *ModifyBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyBlogResponse.Marshal(b, m, deterministic)
}
func (m *ModifyBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyBlogResponse.Merge(m, src)
}
func (m *ModifyBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ModifyBlogResponse.Size(m)
}
func (m *ModifyBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyBlogResponse proto.InternalMessageInfo

func (m *ModifyBlogResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ModifyBlogResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//detail
type BlogDetailRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blogId,proto3" json:"blogId,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogDetailRequest) Reset()         { *m = BlogDetailRequest{} }
func (m *BlogDetailRequest) String() string { return proto.CompactTextString(m) }
func (*BlogDetailRequest) ProtoMessage()    {}
func (*BlogDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{9}
}

func (m *BlogDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogDetailRequest.Unmarshal(m, b)
}
func (m *BlogDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogDetailRequest.Marshal(b, m, deterministic)
}
func (m *BlogDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogDetailRequest.Merge(m, src)
}
func (m *BlogDetailRequest) XXX_Size() int {
	return xxx_messageInfo_BlogDetailRequest.Size(m)
}
func (m *BlogDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlogDetailRequest proto.InternalMessageInfo

func (m *BlogDetailRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

func (m *BlogDetailRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type BlogDetailResponse struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	BuildTime            string   `protobuf:"bytes,2,opt,name=buildTime,proto3" json:"buildTime,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Detail               string   `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlogDetailResponse) Reset()         { *m = BlogDetailResponse{} }
func (m *BlogDetailResponse) String() string { return proto.CompactTextString(m) }
func (*BlogDetailResponse) ProtoMessage()    {}
func (*BlogDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0c6db8a9d783a32, []int{10}
}

func (m *BlogDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlogDetailResponse.Unmarshal(m, b)
}
func (m *BlogDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlogDetailResponse.Marshal(b, m, deterministic)
}
func (m *BlogDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlogDetailResponse.Merge(m, src)
}
func (m *BlogDetailResponse) XXX_Size() int {
	return xxx_messageInfo_BlogDetailResponse.Size(m)
}
func (m *BlogDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlogDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlogDetailResponse proto.InternalMessageInfo

func (m *BlogDetailResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *BlogDetailResponse) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

func (m *BlogDetailResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BlogDetailResponse) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterType((*BlogListRequest)(nil), "BlogListRequest")
	proto.RegisterType((*BlogListResponse)(nil), "BlogListResponse")
	proto.RegisterType((*BlogInfo)(nil), "BlogInfo")
	proto.RegisterType((*WriteBlogRequest)(nil), "WriteBlogRequest")
	proto.RegisterType((*WriteBlogResponse)(nil), "WriteBlogResponse")
	proto.RegisterType((*DeleteBlogRequest)(nil), "DeleteBlogRequest")
	proto.RegisterType((*DeleteBlogResponse)(nil), "DeleteBlogResponse")
	proto.RegisterType((*ModifyBlogRequest)(nil), "ModifyBlogRequest")
	proto.RegisterType((*ModifyBlogResponse)(nil), "ModifyBlogResponse")
	proto.RegisterType((*BlogDetailRequest)(nil), "BlogDetailRequest")
	proto.RegisterType((*BlogDetailResponse)(nil), "BlogDetailResponse")
}

func init() {
	proto.RegisterFile("go/src/rpc_blog/proto/blog/blog_ext.proto", fileDescriptor_a0c6db8a9d783a32)
}

var fileDescriptor_a0c6db8a9d783a32 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0xcb, 0xd3, 0x40,
	0x10, 0x6d, 0xd3, 0xf8, 0xf9, 0x65, 0x7a, 0xb0, 0x99, 0x8a, 0x84, 0xe0, 0xe1, 0x23, 0x20, 0xd6,
	0xcb, 0x06, 0x6a, 0x41, 0xbc, 0x69, 0x29, 0x48, 0x41, 0x2f, 0x51, 0xf0, 0x58, 0x9a, 0x64, 0x1a,
	0x16, 0xd3, 0x6e, 0x4d, 0x36, 0xa2, 0xff, 0xc7, 0x1f, 0x2a, 0xbb, 0x4d, 0xba, 0x4b, 0x72, 0x91,
	0xea, 0x25, 0xec, 0xcc, 0xee, 0xbe, 0x79, 0xfb, 0xe6, 0x4d, 0xe0, 0x55, 0x21, 0xe2, 0xba, 0xca,
	0xe2, 0xea, 0x9c, 0xed, 0xd2, 0x52, 0x14, 0xf1, 0xb9, 0x12, 0x52, 0xc4, 0x7a, 0xa9, 0x3e, 0x3b,
	0xfa, 0x29, 0x99, 0xce, 0x45, 0x2f, 0xe1, 0xc9, 0xba, 0x14, 0xc5, 0x47, 0x5e, 0xcb, 0x84, 0xbe,
	0x37, 0x54, 0x4b, 0x7c, 0x0a, 0x8f, 0xa4, 0xf8, 0x46, 0xa7, 0x60, 0xfc, 0x30, 0x5e, 0x78, 0xc9,
	0x25, 0x88, 0xde, 0xc2, 0xcc, 0x1c, 0xac, 0xcf, 0xe2, 0x54, 0x13, 0xbe, 0x80, 0xfb, 0xb4, 0xcd,
	0x05, 0xe3, 0x87, 0xc9, 0x62, 0xba, 0xf4, 0x98, 0x3a, 0xb4, 0x3d, 0x1d, 0x44, 0x72, 0xdd, 0x8a,
	0x12, 0xb8, 0xef, 0xb2, 0x88, 0xe0, 0x36, 0x0d, 0xcf, 0x5b, 0x6c, 0xbd, 0xc6, 0xe7, 0xe0, 0xa5,
	0x0d, 0x2f, 0xf3, 0x2f, 0xfc, 0x48, 0x81, 0xa3, 0x37, 0x4c, 0x42, 0xd3, 0xe1, 0xb2, 0xa4, 0x60,
	0xd2, 0xd2, 0x51, 0x41, 0xf4, 0x0e, 0x66, 0x5f, 0x2b, 0x2e, 0x49, 0x01, 0x77, 0xc4, 0x9f, 0xc1,
	0x9d, 0xaa, 0xb9, 0xed, 0xd0, 0xdb, 0xc8, 0x3c, 0xc8, 0xb1, 0x1f, 0xf4, 0x1e, 0x7c, 0x0b, 0xa1,
	0x7d, 0x11, 0x82, 0x9b, 0x89, 0x9c, 0x34, 0xc0, 0x24, 0xd1, 0x6b, 0x0c, 0xe0, 0xf1, 0x91, 0xea,
	0x7a, 0x5f, 0x74, 0xe4, 0xba, 0x50, 0x41, 0x6c, 0xa8, 0xa4, 0x7f, 0x61, 0xb1, 0x06, 0xb4, 0x21,
	0x6e, 0xa2, 0x21, 0xc0, 0xff, 0x24, 0x72, 0x7e, 0xf8, 0xf5, 0xb7, 0x34, 0xb4, 0x9c, 0x8e, 0x25,
	0xa7, 0x3a, 0x9d, 0x93, 0xdc, 0xf3, 0xb2, 0x55, 0xb9, 0x8d, 0x0c, 0x69, 0xb7, 0x47, 0xda, 0x2e,
	0x78, 0xab, 0x76, 0xea, 0xf6, 0x46, 0xd7, 0xb9, 0x4d, 0x3b, 0x09, 0x68, 0x43, 0x18, 0x1a, 0xff,
	0xc3, 0x61, 0x96, 0x24, 0xae, 0x2d, 0xc9, 0xf2, 0xb7, 0x03, 0x53, 0x55, 0xf6, 0x33, 0x55, 0x3f,
	0x78, 0x46, 0xb8, 0x82, 0xe9, 0x07, 0x92, 0xdd, 0x6c, 0xe0, 0x8c, 0xf5, 0xe6, 0x29, 0xf4, 0x59,
	0x7f, 0x70, 0xa2, 0x11, 0xae, 0xc0, 0xbb, 0xba, 0x0f, 0x7d, 0xd6, 0xf7, 0x72, 0x88, 0x6c, 0x60,
	0xce, 0x68, 0x84, 0x6f, 0x00, 0x8c, 0x5b, 0x10, 0xd9, 0xc0, 0x7d, 0xe1, 0x9c, 0x0d, 0xed, 0x74,
	0xb9, 0x68, 0x3a, 0x86, 0xc8, 0x06, 0x7e, 0x09, 0xe7, 0x6c, 0xd8, 0xd2, 0xcb, 0x45, 0xa3, 0x31,
	0x22, 0x1b, 0xf4, 0x2c, 0x9c, 0xb3, 0x61, 0x13, 0xa2, 0x51, 0x7a, 0xa7, 0xff, 0x2f, 0xaf, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x90, 0x00, 0x0b, 0x8c, 0x04, 0x00, 0x00,
}
