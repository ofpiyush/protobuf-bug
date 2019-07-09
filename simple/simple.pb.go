// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple/simple.proto

package simple

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

type BytesTest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesTest) Reset()         { *m = BytesTest{} }
func (m *BytesTest) String() string { return proto.CompactTextString(m) }
func (*BytesTest) ProtoMessage()    {}
func (*BytesTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_simple_68e58cbecfd78430, []int{0}
}
func (m *BytesTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesTest.Unmarshal(m, b)
}
func (m *BytesTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesTest.Marshal(b, m, deterministic)
}
func (dst *BytesTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesTest.Merge(dst, src)
}
func (m *BytesTest) XXX_Size() int {
	return xxx_messageInfo_BytesTest.Size(m)
}
func (m *BytesTest) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesTest.DiscardUnknown(m)
}

var xxx_messageInfo_BytesTest proto.InternalMessageInfo

func (m *BytesTest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BytesTest)(nil), "simple.BytesTest")
}

func init() { proto.RegisterFile("simple/simple.proto", fileDescriptor_simple_68e58cbecfd78430) }

var fileDescriptor_simple_68e58cbecfd78430 = []byte{
	// 79 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92,
	0x3c, 0x17, 0xa7, 0x53, 0x65, 0x49, 0x6a, 0x71, 0x48, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x8b,
	0x4b, 0x62, 0x49, 0xa2, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x98, 0x9d, 0xc4, 0x06, 0x56,
	0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x59, 0x76, 0x72, 0x46, 0x00, 0x00, 0x00,
}
