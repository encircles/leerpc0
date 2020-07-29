// Code generated by protoc-gen-go. DO NOT EDIT.
// source: x1.proto

package foo_bar

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

type SearchRequest_Corpus int32

const (
	SearchRequest_UNIVERSAL SearchRequest_Corpus = 0
	SearchRequest_WEB       SearchRequest_Corpus = 1
	SearchRequest_IMAGES    SearchRequest_Corpus = 2
	SearchRequest_LOCAL     SearchRequest_Corpus = 3
	SearchRequest_NEWS      SearchRequest_Corpus = 4
	SearchRequest_PRODUCTS  SearchRequest_Corpus = 5
	SearchRequest_VIDEO     SearchRequest_Corpus = 6
)

var SearchRequest_Corpus_name = map[int32]string{
	0: "UNIVERSAL",
	1: "WEB",
	2: "IMAGES",
	3: "LOCAL",
	4: "NEWS",
	5: "PRODUCTS",
	6: "VIDEO",
}

var SearchRequest_Corpus_value = map[string]int32{
	"UNIVERSAL": 0,
	"WEB":       1,
	"IMAGES":    2,
	"LOCAL":     3,
	"NEWS":      4,
	"PRODUCTS":  5,
	"VIDEO":     6,
}

func (x SearchRequest_Corpus) String() string {
	return proto.EnumName(SearchRequest_Corpus_name, int32(x))
}

func (SearchRequest_Corpus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a83375e8069f221e, []int{0, 0}
}

type SearchRequest struct {
	Query                string               `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int32                `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32                `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	Corpus               SearchRequest_Corpus `protobuf:"varint,4,opt,name=corpus,proto3,enum=foo.bar.SearchRequest_Corpus" json:"corpus,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a83375e8069f221e, []int{0}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchRequest) GetPageNumber() int32 {
	if m != nil {
		return m.PageNumber
	}
	return 0
}

func (m *SearchRequest) GetResultPerPage() int32 {
	if m != nil {
		return m.ResultPerPage
	}
	return 0
}

func (m *SearchRequest) GetCorpus() SearchRequest_Corpus {
	if m != nil {
		return m.Corpus
	}
	return SearchRequest_UNIVERSAL
}

type SearchResponse struct {
	Results              []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a83375e8069f221e, []int{1}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type Result struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Snippets             []string `protobuf:"bytes,3,rep,name=snippets,proto3" json:"snippets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_a83375e8069f221e, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Result) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Result) GetSnippets() []string {
	if m != nil {
		return m.Snippets
	}
	return nil
}

type Foo struct {
	Open                 *Open    `protobuf:"bytes,1,opt,name=open,proto3" json:"open,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a83375e8069f221e, []int{3}
}

func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetOpen() *Open {
	if m != nil {
		return m.Open
	}
	return nil
}

type Open struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Open) Reset()         { *m = Open{} }
func (m *Open) String() string { return proto.CompactTextString(m) }
func (*Open) ProtoMessage()    {}
func (*Open) Descriptor() ([]byte, []int) {
	return fileDescriptor_a83375e8069f221e, []int{4}
}

func (m *Open) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Open.Unmarshal(m, b)
}
func (m *Open) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Open.Marshal(b, m, deterministic)
}
func (m *Open) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Open.Merge(m, src)
}
func (m *Open) XXX_Size() int {
	return xxx_messageInfo_Open.Size(m)
}
func (m *Open) XXX_DiscardUnknown() {
	xxx_messageInfo_Open.DiscardUnknown(m)
}

var xxx_messageInfo_Open proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("foo.bar.SearchRequest_Corpus", SearchRequest_Corpus_name, SearchRequest_Corpus_value)
	proto.RegisterType((*SearchRequest)(nil), "foo.bar.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "foo.bar.SearchResponse")
	proto.RegisterType((*Result)(nil), "foo.bar.Result")
	proto.RegisterType((*Foo)(nil), "foo.bar.Foo")
	proto.RegisterType((*Open)(nil), "foo.bar.Open")
}

func init() { proto.RegisterFile("x1.proto", fileDescriptor_a83375e8069f221e) }

var fileDescriptor_a83375e8069f221e = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x6b, 0xe2, 0x40,
	0x18, 0xc5, 0x37, 0x4e, 0x12, 0x93, 0xcf, 0x8d, 0x0e, 0xc3, 0x1e, 0xc2, 0xc2, 0xb2, 0xd9, 0x1c,
	0x96, 0xec, 0x25, 0xb0, 0x96, 0x9e, 0x7a, 0xb2, 0x9a, 0x16, 0x21, 0x35, 0x32, 0xa9, 0x0a, 0xbd,
	0x48, 0x94, 0xa9, 0x15, 0x6c, 0x66, 0x9c, 0x49, 0xa0, 0x3d, 0xf6, 0x3f, 0x2f, 0x99, 0xa8, 0xd0,
	0x5b, 0xde, 0xef, 0x3d, 0xf2, 0x3d, 0xde, 0x80, 0xf3, 0xf6, 0x3f, 0x16, 0x92, 0x57, 0x9c, 0x74,
	0x9f, 0x39, 0x8f, 0x37, 0x85, 0x0c, 0x3f, 0x3a, 0xe0, 0xe5, 0xac, 0x90, 0xdb, 0x17, 0xca, 0x8e,
	0x35, 0x53, 0x15, 0xf9, 0x01, 0xd6, 0xb1, 0x66, 0xf2, 0xdd, 0x37, 0x02, 0x23, 0x72, 0x69, 0x2b,
	0xc8, 0x6f, 0xe8, 0x89, 0x62, 0xc7, 0xd6, 0x65, 0xfd, 0xba, 0x61, 0xd2, 0xef, 0x04, 0x46, 0x64,
	0x51, 0x68, 0xd0, 0x4c, 0x13, 0xf2, 0x17, 0x06, 0x92, 0xa9, 0xfa, 0x50, 0xad, 0x05, 0x93, 0xeb,
	0xc6, 0xf0, 0x91, 0x0e, 0x79, 0x2d, 0x9e, 0x33, 0x39, 0x2f, 0x76, 0x8c, 0x5c, 0x83, 0xbd, 0xe5,
	0x52, 0xd4, 0xca, 0x37, 0x03, 0x23, 0xea, 0x0f, 0x7f, 0xc5, 0xa7, 0x2a, 0xf1, 0x97, 0x1a, 0xf1,
	0x58, 0x87, 0xe8, 0x29, 0x1c, 0x3e, 0x81, 0xdd, 0x12, 0xe2, 0x81, 0xbb, 0x98, 0x4d, 0x97, 0x09,
	0xcd, 0x47, 0x29, 0xfe, 0x46, 0xba, 0x80, 0x56, 0xc9, 0x2d, 0x36, 0x08, 0x80, 0x3d, 0x7d, 0x18,
	0xdd, 0x27, 0x39, 0xee, 0x10, 0x17, 0xac, 0x34, 0x1b, 0x8f, 0x52, 0x8c, 0x88, 0x03, 0xe6, 0x2c,
	0x59, 0xe5, 0xd8, 0x24, 0xdf, 0xc1, 0x99, 0xd3, 0x6c, 0xb2, 0x18, 0x3f, 0xe6, 0xd8, 0x6a, 0x22,
	0xcb, 0xe9, 0x24, 0xc9, 0xb0, 0x1d, 0xde, 0x40, 0xff, 0x7c, 0x5b, 0x09, 0x5e, 0x2a, 0x46, 0xfe,
	0x41, 0xb7, 0x6d, 0xad, 0x7c, 0x23, 0x40, 0x51, 0x6f, 0x38, 0xb8, 0xb4, 0xa4, 0x9a, 0xd3, 0xb3,
	0x1f, 0xa6, 0x60, 0xb7, 0x88, 0x60, 0x40, 0xb5, 0x3c, 0x9c, 0x66, 0x6b, 0x3e, 0x9b, 0x29, 0xab,
	0x7d, 0x75, 0x60, 0x7a, 0x2e, 0x97, 0xb6, 0x82, 0xfc, 0x04, 0x47, 0x95, 0x7b, 0x21, 0x58, 0xa5,
	0x7c, 0x14, 0xa0, 0xc8, 0xa5, 0x17, 0x1d, 0x46, 0x80, 0xee, 0x38, 0x27, 0x7f, 0xc0, 0xe4, 0x82,
	0x95, 0xfa, 0x5f, 0xbd, 0xa1, 0x77, 0x39, 0x9e, 0x09, 0x56, 0x52, 0x6d, 0x85, 0x36, 0x98, 0x8d,
	0xda, 0xd8, 0xfa, 0x41, 0xaf, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x69, 0xde, 0x74, 0xdc,
	0x01, 0x00, 0x00,
}
