// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spec.proto

package redisCache

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Method               string            `protobuf:"bytes,1,opt,name=Method,proto3" json:"Method,omitempty"`
	URL                  map[string]string `protobuf:"bytes,2,rep,name=URL,proto3" json:"URL,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Header               map[string]string `protobuf:"bytes,3,rep,name=Header,proto3" json:"Header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Host                 string            `protobuf:"bytes,4,opt,name=Host,proto3" json:"Host,omitempty"`
	Form                 map[string]string `protobuf:"bytes,5,rep,name=Form,proto3" json:"Form,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PostForm             map[string]string `protobuf:"bytes,6,rep,name=PostForm,proto3" json:"PostForm,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_spec_d31611f0279a4ded, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Request) GetURL() map[string]string {
	if m != nil {
		return m.URL
	}
	return nil
}

func (m *Request) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Request) GetForm() map[string]string {
	if m != nil {
		return m.Form
	}
	return nil
}

func (m *Request) GetPostForm() map[string]string {
	if m != nil {
		return m.PostForm
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "redisCache.Request")
	proto.RegisterMapType((map[string]string)(nil), "redisCache.Request.FormEntry")
	proto.RegisterMapType((map[string]string)(nil), "redisCache.Request.HeaderEntry")
	proto.RegisterMapType((map[string]string)(nil), "redisCache.Request.PostFormEntry")
	proto.RegisterMapType((map[string]string)(nil), "redisCache.Request.URLEntry")
}

func init() { proto.RegisterFile("spec.proto", fileDescriptor_spec_d31611f0279a4ded) }

var fileDescriptor_spec_d31611f0279a4ded = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2e, 0x48, 0x4d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0x4a, 0x4d, 0xc9, 0x2c, 0x76, 0x4e, 0x4c,
	0xce, 0x48, 0x55, 0xfa, 0xc9, 0xcc, 0xc5, 0x1e, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24,
	0xc6, 0xc5, 0xe6, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe5, 0x09, 0xe9, 0x71, 0x31, 0x87, 0x06, 0xf9, 0x48, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x1b, 0xc9,
	0xe8, 0x21, 0x74, 0xeb, 0x41, 0x75, 0xea, 0x85, 0x06, 0xf9, 0xb8, 0xe6, 0x95, 0x14, 0x55, 0x06,
	0x81, 0x14, 0x0a, 0x99, 0x73, 0xb1, 0x79, 0xa4, 0x26, 0xa6, 0xa4, 0x16, 0x49, 0x30, 0x83, 0xb5,
	0xc8, 0x63, 0xd3, 0x02, 0x51, 0x01, 0xd1, 0x05, 0x55, 0x2e, 0x24, 0xc4, 0xc5, 0xe2, 0x91, 0x5f,
	0x5c, 0x22, 0xc1, 0x02, 0xb6, 0x1e, 0xcc, 0x16, 0x32, 0xe4, 0x62, 0x71, 0xcb, 0x2f, 0xca, 0x95,
	0x60, 0x05, 0x1b, 0x25, 0x8b, 0xcd, 0x28, 0x90, 0x3c, 0xc4, 0x20, 0xb0, 0x52, 0x21, 0x5b, 0x2e,
	0x8e, 0x80, 0xfc, 0xe2, 0x12, 0xb0, 0x36, 0x36, 0xb0, 0x36, 0x45, 0x6c, 0xda, 0x60, 0x6a, 0x20,
	0x5a, 0xe1, 0x5a, 0xa4, 0xcc, 0xb8, 0x38, 0x60, 0xfe, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad,
	0x84, 0x86, 0x07, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04,
	0x16, 0x83, 0x70, 0xac, 0x98, 0x2c, 0x18, 0xa5, 0x2c, 0xb9, 0xb8, 0x91, 0x3c, 0x45, 0x92, 0x56,
	0x73, 0x2e, 0x4e, 0xb8, 0x4b, 0x48, 0xd2, 0x68, 0xcd, 0xc5, 0x8b, 0xe2, 0x0d, 0x52, 0x34, 0x27,
	0xb1, 0x81, 0x93, 0x83, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xed, 0x05, 0x21, 0x3a, 0x1c, 0x02,
	0x00, 0x00,
}