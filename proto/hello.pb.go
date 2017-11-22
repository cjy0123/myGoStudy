// Code generated by protoc-gen-go.
// source: hello.proto
// DO NOT EDIT!

/*
Package greeter is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package greeter

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

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "greeter.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "greeter.HelloReply")
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d,
	0x52, 0x52, 0xe2, 0xe2, 0xf1, 0x00, 0x89, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09,
	0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a,
	0x6a, 0x5c, 0x5c, 0x50, 0x35, 0x05, 0x39, 0x95, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5,
	0x89, 0xe9, 0x30, 0x45, 0x30, 0xae, 0x91, 0x33, 0x17, 0xbb, 0x3b, 0xc4, 0x58, 0x21, 0x0b, 0x2e,
	0x8e, 0xe0, 0xc4, 0x4a, 0xb0, 0x2e, 0x21, 0x51, 0x3d, 0xa8, 0x65, 0x7a, 0xc8, 0x36, 0x49, 0x09,
	0xa3, 0x0b, 0x17, 0xe4, 0x54, 0x2a, 0x31, 0x38, 0x89, 0x70, 0x09, 0x25, 0xe7, 0xe7, 0xea, 0x25,
	0xe7, 0xe7, 0xe4, 0x27, 0x95, 0xea, 0x15, 0x15, 0x24, 0x97, 0xa4, 0x16, 0x97, 0x24, 0xb1, 0x81,
	0x9d, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x57, 0x34, 0x5f, 0xc5, 0x00, 0x00, 0x00,
}