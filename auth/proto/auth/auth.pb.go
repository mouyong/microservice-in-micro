// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package mu_micro_book_srv_auth

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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "mu.micro.book.srv.auth.Error")
	proto.RegisterType((*Request)(nil), "mu.micro.book.srv.auth.Request")
	proto.RegisterType((*Response)(nil), "mu.micro.book.srv.auth.Response")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x8d, 0x35, 0x4d, 0x1c, 0x0f, 0xc2, 0xa0, 0x25, 0x14, 0xc4, 0x90, 0x53, 0x4f, 0x2b,
	0x34, 0x9f, 0x40, 0xd0, 0x83, 0x07, 0x3d, 0x6c, 0x55, 0xf0, 0x98, 0x26, 0x03, 0x86, 0x34, 0x9d,
	0xba, 0x7f, 0xfa, 0xe9, 0xfc, 0x70, 0xb2, 0x93, 0xe8, 0xa9, 0xbd, 0x79, 0x09, 0xf3, 0xe3, 0xbd,
	0x79, 0x93, 0xc7, 0xc2, 0xf5, 0xce, 0xb0, 0xe3, 0xbb, 0xca, 0xbb, 0x4f, 0xf9, 0x28, 0x61, 0x9c,
	0xf5, 0x5e, 0xf5, 0x6d, 0x6d, 0x58, 0xad, 0x99, 0x3b, 0x65, 0xcd, 0x5e, 0x05, 0xb5, 0x28, 0x21,
	0x7e, 0x34, 0x86, 0x0d, 0x22, 0x9c, 0xd5, 0xdc, 0x50, 0x16, 0xe5, 0xd1, 0x22, 0xd6, 0x32, 0xe3,
	0x0c, 0xa6, 0x0d, 0xb9, 0xaa, 0xdd, 0x64, 0xa7, 0x79, 0xb4, 0x38, 0xd7, 0x23, 0x15, 0x2b, 0x48,
	0x34, 0x7d, 0x79, 0xb2, 0x2e, 0x58, 0xbc, 0x25, 0xf3, 0xd4, 0xc8, 0xe2, 0x44, 0x8f, 0x84, 0x73,
	0x48, 0xc3, 0xf4, 0x52, 0xf5, 0x34, 0x2e, 0xff, 0x31, 0x5e, 0x41, 0xec, 0xb8, 0xa3, 0x6d, 0x36,
	0x11, 0x61, 0x80, 0x82, 0x21, 0xd5, 0x64, 0x77, 0xbc, 0xb5, 0x84, 0x19, 0x24, 0xd6, 0xd7, 0x35,
	0x59, 0x2b, 0xb1, 0xa9, 0xfe, 0x45, 0x2c, 0x21, 0xa6, 0xf0, 0xbf, 0x12, 0x7a, 0xb1, 0xbc, 0x51,
	0x87, 0x7b, 0x29, 0x29, 0xa5, 0x07, 0xef, 0xe1, 0x83, 0xcb, 0xef, 0x08, 0x92, 0x15, 0x99, 0x7d,
	0x5b, 0x13, 0xbe, 0xc3, 0xe5, 0x73, 0xd5, 0xd1, 0xbd, 0x1c, 0x79, 0x0d, 0x32, 0xde, 0x1e, 0x8b,
	0x1e, 0xab, 0xcf, 0xf3, 0xe3, 0x86, 0xa1, 0x46, 0x71, 0x82, 0x1f, 0x80, 0x0f, 0xb4, 0x79, 0xb3,
	0x64, 0xfe, 0x3b, 0x7a, 0x3d, 0x95, 0x87, 0x2d, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xa0,
	0x15, 0xfe, 0xf1, 0x01, 0x00, 0x00,
}
