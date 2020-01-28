// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple.proto

package simple

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

type SimpleMessage struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsSimple             bool     `protobuf:"varint,2,opt,name=is_simple,json=isSimple,proto3" json:"is_simple,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	SimpleList           []int32  `protobuf:"varint,4,rep,packed,name=simple_list,json=simpleList,proto3" json:"simple_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleMessage) Reset()         { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()    {}
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ffd045dd4d042c1, []int{0}
}

func (m *SimpleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMessage.Unmarshal(m, b)
}
func (m *SimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMessage.Marshal(b, m, deterministic)
}
func (m *SimpleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMessage.Merge(m, src)
}
func (m *SimpleMessage) XXX_Size() int {
	return xxx_messageInfo_SimpleMessage.Size(m)
}
func (m *SimpleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMessage proto.InternalMessageInfo

func (m *SimpleMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SimpleMessage) GetIsSimple() bool {
	if m != nil {
		return m.IsSimple
	}
	return false
}

func (m *SimpleMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SimpleMessage) GetSimpleList() []int32 {
	if m != nil {
		return m.SimpleList
	}
	return nil
}

func init() {
	proto.RegisterType((*SimpleMessage)(nil), "simple.SimpleMessage")
}

func init() { proto.RegisterFile("simple.proto", fileDescriptor_5ffd045dd4d042c1) }

var fileDescriptor_5ffd045dd4d042c1 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x0a, 0xb9, 0x78,
	0x83, 0xc1, 0x2c, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x69, 0x2e, 0xce, 0xcc, 0xe2,
	0x78, 0x88, 0x6a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x8e, 0xcc, 0x62, 0x88, 0x1e, 0x21,
	0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b,
	0x48, 0x9e, 0x8b, 0x1b, 0xa2, 0x3a, 0x3e, 0x27, 0xb3, 0xb8, 0x44, 0x82, 0x45, 0x81, 0x59, 0x83,
	0x35, 0x88, 0x0b, 0x22, 0xe4, 0x93, 0x59, 0x5c, 0x92, 0xc4, 0x06, 0x76, 0x81, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x08, 0xee, 0xa1, 0x06, 0x91, 0x00, 0x00, 0x00,
}
