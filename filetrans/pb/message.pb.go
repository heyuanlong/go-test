// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	File
*/
package pb

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

type File struct {
	Type             *int32  `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Remain           *int32  `protobuf:"varint,3,req,name=remain" json:"remain,omitempty"`
	Content          *string `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *File) Reset()                    { *m = File{} }
func (m *File) String() string            { return proto.CompactTextString(m) }
func (*File) ProtoMessage()               {}
func (*File) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *File) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *File) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *File) GetRemain() int32 {
	if m != nil && m.Remain != nil {
		return *m.Remain
	}
	return 0
}

func (m *File) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*File)(nil), "pb.file")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72, 0xe6,
	0x62, 0x49, 0xcb, 0xcc, 0x49, 0x15, 0xe2, 0xe1, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54,
	0x60, 0xd2, 0x60, 0x05, 0xf1, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x98, 0x34, 0x38, 0x85,
	0xf8, 0xb8, 0xd8, 0x8a, 0x52, 0x73, 0x13, 0x33, 0xf3, 0x24, 0x98, 0xc1, 0xb2, 0xfc, 0x5c, 0xec,
	0xc9, 0xf9, 0x79, 0x25, 0xa9, 0x79, 0x25, 0x12, 0x2c, 0x20, 0x05, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x40, 0xce, 0x15, 0xa9, 0x58, 0x00, 0x00, 0x00,
}
