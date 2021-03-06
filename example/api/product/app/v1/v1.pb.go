// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product/app/v1/v1.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetArticlesReq struct {
	// @inject_tag: form:"title"
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" form:"title"`
	// @inject_tag: form:"page"
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty" form:"page"`
	// @inject_tag: form:"page_size"
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" form:"page_size"`
	// @inject_tag: form:"author_id" uri:"author_id"
	AuthorId             int32    `protobuf:"varint,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty" form:"author_id" uri:"author_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetArticlesReq) Reset()         { *m = GetArticlesReq{} }
func (m *GetArticlesReq) String() string { return proto.CompactTextString(m) }
func (*GetArticlesReq) ProtoMessage()    {}
func (*GetArticlesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6090e3745c002b5, []int{0}
}

func (m *GetArticlesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArticlesReq.Unmarshal(m, b)
}
func (m *GetArticlesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArticlesReq.Marshal(b, m, deterministic)
}
func (m *GetArticlesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArticlesReq.Merge(m, src)
}
func (m *GetArticlesReq) XXX_Size() int {
	return xxx_messageInfo_GetArticlesReq.Size(m)
}
func (m *GetArticlesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArticlesReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetArticlesReq proto.InternalMessageInfo

func (m *GetArticlesReq) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *GetArticlesReq) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetArticlesReq) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetArticlesReq) GetAuthorId() int32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

type GetArticlesResp struct {
	Total                int64      `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Articles             []*Article `protobuf:"bytes,2,rep,name=articles,proto3" json:"articles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetArticlesResp) Reset()         { *m = GetArticlesResp{} }
func (m *GetArticlesResp) String() string { return proto.CompactTextString(m) }
func (*GetArticlesResp) ProtoMessage()    {}
func (*GetArticlesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6090e3745c002b5, []int{1}
}

func (m *GetArticlesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArticlesResp.Unmarshal(m, b)
}
func (m *GetArticlesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArticlesResp.Marshal(b, m, deterministic)
}
func (m *GetArticlesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArticlesResp.Merge(m, src)
}
func (m *GetArticlesResp) XXX_Size() int {
	return xxx_messageInfo_GetArticlesResp.Size(m)
}
func (m *GetArticlesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArticlesResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetArticlesResp proto.InternalMessageInfo

func (m *GetArticlesResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GetArticlesResp) GetArticles() []*Article {
	if m != nil {
		return m.Articles
	}
	return nil
}

type Article struct {
	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// @inject_tag: form:"author_id" uri:"author_id"
	AuthorId             int32    `protobuf:"varint,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty" form:"author_id" uri:"author_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6090e3745c002b5, []int{2}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Article) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Article) GetAuthorId() int32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

func init() {
	proto.RegisterType((*GetArticlesReq)(nil), "product.app.v1.GetArticlesReq")
	proto.RegisterType((*GetArticlesResp)(nil), "product.app.v1.GetArticlesResp")
	proto.RegisterType((*Article)(nil), "product.app.v1.Article")
}

func init() { proto.RegisterFile("product/app/v1/v1.proto", fileDescriptor_c6090e3745c002b5) }

var fileDescriptor_c6090e3745c002b5 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x6b, 0xe2, 0x40,
	0x18, 0x25, 0x89, 0xae, 0x3a, 0xee, 0xba, 0x30, 0x2c, 0x18, 0xdc, 0x65, 0xd5, 0x5c, 0xd6, 0x8b,
	0x19, 0xa2, 0x87, 0x3d, 0xab, 0x87, 0x65, 0xaf, 0x11, 0x7a, 0x90, 0x82, 0x8c, 0x71, 0x98, 0x4c,
	0x89, 0x99, 0x69, 0xf2, 0x25, 0x14, 0x4b, 0x2f, 0xed, 0xa1, 0x3f, 0xa0, 0x3f, 0xad, 0x7f, 0xa1,
	0x3f, 0xa4, 0x38, 0x89, 0xd2, 0x14, 0x6c, 0x4f, 0xc9, 0xfb, 0xde, 0x23, 0xef, 0xbd, 0x2f, 0x1f,
	0xea, 0xaa, 0x44, 0x6e, 0xb3, 0x00, 0x08, 0x55, 0x8a, 0xe4, 0x1e, 0xc9, 0x3d, 0x57, 0x25, 0x12,
	0x24, 0xee, 0x94, 0x84, 0x4b, 0x95, 0x72, 0x73, 0xaf, 0xf7, 0x8b, 0x4b, 0xc9, 0x23, 0x46, 0xa8,
	0x12, 0x84, 0xc6, 0xb1, 0x04, 0x0a, 0x42, 0xc6, 0x69, 0xa1, 0x76, 0x00, 0x75, 0xfe, 0x31, 0x98,
	0x25, 0x20, 0x82, 0x88, 0xa5, 0x3e, 0xbb, 0xc6, 0x3f, 0x50, 0x1d, 0x04, 0x44, 0xcc, 0x36, 0x06,
	0xc6, 0xa8, 0xe5, 0x17, 0x00, 0x63, 0x54, 0x53, 0x94, 0x33, 0xdb, 0x1c, 0x18, 0xa3, 0xba, 0xaf,
	0xdf, 0xf1, 0x4f, 0xd4, 0x3a, 0x3c, 0xd7, 0xa9, 0xd8, 0x33, 0xdb, 0xd2, 0x44, 0xf3, 0x30, 0x58,
	0x8a, 0xbd, 0x26, 0x69, 0x06, 0xa1, 0x4c, 0xd6, 0x62, 0x6b, 0xd7, 0x0a, 0xb2, 0x18, 0xfc, 0xdf,
	0x3a, 0x97, 0xe8, 0x7b, 0xc5, 0x35, 0x55, 0xda, 0x56, 0x02, 0x8d, 0xb4, 0xad, 0xe5, 0x17, 0x00,
	0x4f, 0x51, 0x93, 0x96, 0x2a, 0xdb, 0x1c, 0x58, 0xa3, 0xf6, 0xa4, 0xeb, 0x56, 0xfb, 0xb9, 0xe5,
	0x57, 0xfc, 0x93, 0xd0, 0xb9, 0x40, 0x8d, 0x72, 0x78, 0xa6, 0x8c, 0x8d, 0x1a, 0x81, 0x8c, 0x81,
	0xc5, 0xa0, 0xfb, 0xb4, 0xfc, 0x23, 0xac, 0xa6, 0xb6, 0xaa, 0xa9, 0x27, 0x0f, 0x26, 0x6a, 0xcf,
	0x23, 0xc9, 0x97, 0x2c, 0xc9, 0x45, 0xc0, 0xf0, 0xa3, 0x81, 0xda, 0x6f, 0x6a, 0xe0, 0xdf, 0xef,
	0xa3, 0x55, 0x37, 0xdb, 0xeb, 0x7f, 0xc8, 0xa7, 0xca, 0xf9, 0x7b, 0xff, 0xfc, 0xf2, 0x64, 0x7a,
	0xf8, 0xeb, 0xe1, 0x87, 0x1e, 0xeb, 0xac, 0x86, 0xb8, 0xaf, 0xb1, 0x8e, 0x41, 0x6e, 0x4f, 0xf9,
	0xee, 0x4e, 0x12, 0x7c, 0x85, 0xbe, 0x2d, 0x12, 0x46, 0x81, 0x1d, 0x7b, 0x9f, 0xdb, 0x52, 0xef,
	0x1c, 0xe1, 0xfc, 0xd1, 0xde, 0x43, 0xe7, 0x33, 0xaf, 0xf9, 0x62, 0x35, 0xe3, 0x02, 0xc2, 0x6c,
	0xe3, 0x06, 0x72, 0x47, 0x76, 0x32, 0xcc, 0x44, 0x1a, 0xca, 0x8c, 0xe8, 0x7b, 0x0a, 0xc6, 0x9c,
	0xc5, 0x63, 0x2e, 0xc7, 0x5c, 0xc4, 0x84, 0xdd, 0xd0, 0x9d, 0x2a, 0x2f, 0xaf, 0x7a, 0xad, 0x9b,
	0x2f, 0x5a, 0x3d, 0x7d, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x25, 0xba, 0x8f, 0x91, 0xc6, 0x02, 0x00,
	0x00,
}
