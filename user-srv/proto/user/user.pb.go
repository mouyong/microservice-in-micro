// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package mu_micro_book_srv_user

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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=pwd,proto3" json:"pwd,omitempty"`
	CreatedTime          uint64   `protobuf:"varint,4,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	UpdatedTime          uint64   `protobuf:"varint,5,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func (m *User) GetCreatedTime() uint64 {
	if m != nil {
		return m.CreatedTime
	}
	return 0
}

func (m *User) GetUpdatedTime() uint64 {
	if m != nil {
		return m.UpdatedTime
	}
	return 0
}

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
	return fileDescriptor_9b283a848145d6b7, []int{1}
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
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserPwd              string   `protobuf:"bytes,3,opt,name=userPwd,proto3" json:"userPwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
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

func (m *Request) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetUserPwd() string {
	if m != nil {
		return m.UserPwd
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
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

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "mu.micro.book.srv.user.user")
	proto.RegisterType((*Error)(nil), "mu.micro.book.srv.user.Error")
	proto.RegisterType((*Request)(nil), "mu.micro.book.srv.user.Request")
	proto.RegisterType((*Response)(nil), "mu.micro.book.srv.user.Response")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3d, 0x4f, 0xfb, 0x30,
	0x10, 0xc6, 0xff, 0x6e, 0x93, 0xbe, 0x5c, 0xa5, 0x3f, 0xc8, 0x12, 0x55, 0x54, 0x81, 0x88, 0x32,
	0x75, 0x32, 0xa8, 0xfd, 0x06, 0x48, 0x0c, 0x2c, 0x08, 0x2c, 0x5e, 0x36, 0xa4, 0x36, 0xbe, 0x21,
	0x82, 0xd4, 0xc1, 0x8e, 0x41, 0xdd, 0xd8, 0xf9, 0xd2, 0xe8, 0x2e, 0x69, 0xe9, 0x40, 0x16, 0xeb,
	0x79, 0x2e, 0x3f, 0xdb, 0xf7, 0x5c, 0x0c, 0x27, 0x95, 0xb3, 0xb5, 0xbd, 0x08, 0x1e, 0x1d, 0x2f,
	0x8a, 0xbd, 0x9c, 0x96, 0x41, 0x95, 0x45, 0xee, 0xac, 0x5a, 0x5b, 0xfb, 0xaa, 0xbc, 0xfb, 0x50,
	0xf4, 0x35, 0xfb, 0x12, 0x10, 0x91, 0x90, 0xff, 0xa1, 0x57, 0x98, 0x44, 0xa4, 0x62, 0xde, 0xd7,
	0xbd, 0xc2, 0x48, 0x09, 0xd1, 0x66, 0x55, 0x62, 0xd2, 0x4b, 0xc5, 0x7c, 0xac, 0x59, 0xcb, 0x63,
	0xe8, 0x57, 0x9f, 0x26, 0xe9, 0x73, 0x89, 0xa4, 0x4c, 0x61, 0x92, 0x3b, 0x5c, 0xd5, 0x68, 0x1e,
	0x8a, 0x12, 0x93, 0x28, 0x15, 0xf3, 0x48, 0x1f, 0x96, 0x88, 0x08, 0x95, 0xd9, 0x13, 0x71, 0x43,
	0x1c, 0x94, 0xb2, 0x25, 0xc4, 0xd7, 0xce, 0x59, 0x47, 0x57, 0xe6, 0xd6, 0x20, 0x37, 0x11, 0x6b,
	0xd6, 0x72, 0x0a, 0x03, 0x83, 0xf5, 0xaa, 0x78, 0x6b, 0x1b, 0x69, 0x5d, 0xf6, 0x0c, 0x43, 0x8d,
	0xef, 0x01, 0x7d, 0x4d, 0x08, 0x25, 0xb8, 0x69, 0xba, 0x1f, 0xeb, 0xd6, 0xc9, 0x19, 0x8c, 0x48,
	0xdd, 0xfe, 0xa6, 0xd8, 0x7b, 0x99, 0xc0, 0x90, 0xf4, 0xdd, 0x3e, 0xcd, 0xce, 0x66, 0xdf, 0x02,
	0x46, 0x1a, 0x7d, 0x65, 0x37, 0x9e, 0x31, 0x1f, 0xf2, 0x1c, 0xbd, 0xe7, 0xb3, 0x47, 0x7a, 0x67,
	0xe5, 0x12, 0x62, 0xa4, 0xa6, 0xf9, 0xe4, 0xc9, 0xe2, 0x4c, 0xfd, 0x3d, 0x5f, 0xc5, 0xc9, 0x74,
	0xc3, 0xca, 0xcb, 0x66, 0xd6, 0x7c, 0xe5, 0x64, 0x71, 0xda, 0xb5, 0x87, 0x16, 0xcd, 0xe4, 0xe2,
	0x05, 0xa2, 0x47, 0xfa, 0x3b, 0x4f, 0x70, 0x74, 0x1f, 0xd0, 0x6d, 0xc9, 0x5c, 0x6d, 0x39, 0xc2,
	0x79, 0xd7, 0xf6, 0x76, 0x2e, 0xb3, 0xb4, 0x1b, 0x68, 0xe2, 0x65, 0xff, 0xd6, 0x03, 0x7e, 0x1d,
	0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x58, 0xbc, 0x2e, 0x36, 0x02, 0x00, 0x00,
}
