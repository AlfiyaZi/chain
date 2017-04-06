// Code generated by protoc-gen-go.
// source: hash.proto
// DO NOT EDIT!

/*
Package bc is a generated protocol buffer package.

It is generated from these files:
	hash.proto

It has these top-level messages:
	Hash
*/
package bc

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

type Hash struct {
	V0 uint64 `protobuf:"fixed64,1,opt,name=v0" json:"v0,omitempty"`
	V1 uint64 `protobuf:"fixed64,2,opt,name=v1" json:"v1,omitempty"`
	V2 uint64 `protobuf:"fixed64,3,opt,name=v2" json:"v2,omitempty"`
	V3 uint64 `protobuf:"fixed64,4,opt,name=v3" json:"v3,omitempty"`
}

func (m *Hash) Reset()                    { *m = Hash{} }
func (m *Hash) String() string            { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()               {}
func (*Hash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Hash) GetV0() uint64 {
	if m != nil {
		return m.V0
	}
	return 0
}

func (m *Hash) GetV1() uint64 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *Hash) GetV2() uint64 {
	if m != nil {
		return m.V2
	}
	return 0
}

func (m *Hash) GetV3() uint64 {
	if m != nil {
		return m.V3
	}
	return 0
}

func init() {
	proto.RegisterType((*Hash)(nil), "bc.Hash")
}

func init() { proto.RegisterFile("hash.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 92 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x48, 0x2c, 0xce,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4a, 0x4a, 0x56, 0x72, 0xe3, 0x62, 0xf1, 0x48,
	0x2c, 0xce, 0x10, 0xe2, 0xe3, 0x62, 0x2a, 0x33, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0b, 0x62,
	0x2a, 0x33, 0x00, 0xf3, 0x0d, 0x25, 0x98, 0xa0, 0x7c, 0x43, 0x30, 0xdf, 0x48, 0x82, 0x19, 0xca,
	0x37, 0x02, 0xf3, 0x8d, 0x25, 0x58, 0xa0, 0x7c, 0xe3, 0x24, 0x36, 0xb0, 0x91, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x78, 0xc6, 0xcf, 0xd1, 0x60, 0x00, 0x00, 0x00,
}
