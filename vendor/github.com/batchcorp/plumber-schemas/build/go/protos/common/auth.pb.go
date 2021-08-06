// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package common

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

type Auth struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Auth)(nil), "protos.common.Auth")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x05, 0x53, 0xc5, 0x7a, 0xc9, 0xf9, 0xb9, 0xb9,
	0xf9, 0x79, 0x4a, 0x32, 0x5c, 0x2c, 0x8e, 0xa5, 0x25, 0x19, 0x42, 0x22, 0x5c, 0xac, 0x25, 0xf9,
	0xd9, 0xa9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x93, 0x6d, 0x94, 0x75,
	0x7a, 0x66, 0x49, 0x46, 0x69, 0x12, 0x48, 0xb9, 0x7e, 0x52, 0x62, 0x49, 0x72, 0x46, 0x72, 0x7e,
	0x51, 0x81, 0x7e, 0x41, 0x4e, 0x69, 0x6e, 0x52, 0x6a, 0x91, 0x6e, 0x71, 0x72, 0x46, 0x6a, 0x6e,
	0x62, 0xb1, 0x7e, 0x52, 0x69, 0x66, 0x4e, 0x8a, 0x7e, 0x7a, 0xbe, 0x3e, 0xc4, 0x70, 0x7d, 0x88,
	0xe1, 0x49, 0x6c, 0x60, 0xae, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x81, 0xae, 0xa1, 0xd2, 0x80,
	0x00, 0x00, 0x00,
}