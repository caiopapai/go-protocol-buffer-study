// Code generated by protoc-gen-go. DO NOT EDIT.
// source: day_of_week.proto

package enumspb

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

type DayOfWeek_Day int32

const (
	DayOfWeek_UNDEFINED_DAY DayOfWeek_Day = 0
	DayOfWeek_SEGUNDA       DayOfWeek_Day = 1
	DayOfWeek_TERCA         DayOfWeek_Day = 2
	DayOfWeek_QUARTA        DayOfWeek_Day = 3
	DayOfWeek_QUINTA        DayOfWeek_Day = 4
	DayOfWeek_SEXTA         DayOfWeek_Day = 5
	DayOfWeek_SABADO        DayOfWeek_Day = 6
	DayOfWeek_DOMINGO       DayOfWeek_Day = 7
)

var DayOfWeek_Day_name = map[int32]string{
	0: "UNDEFINED_DAY",
	1: "SEGUNDA",
	2: "TERCA",
	3: "QUARTA",
	4: "QUINTA",
	5: "SEXTA",
	6: "SABADO",
	7: "DOMINGO",
}

var DayOfWeek_Day_value = map[string]int32{
	"UNDEFINED_DAY": 0,
	"SEGUNDA":       1,
	"TERCA":         2,
	"QUARTA":        3,
	"QUINTA":        4,
	"SEXTA":         5,
	"SABADO":        6,
	"DOMINGO":       7,
}

func (x DayOfWeek_Day) String() string {
	return proto.EnumName(DayOfWeek_Day_name, int32(x))
}

func (DayOfWeek_Day) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_466685131aab1de7, []int{0, 0}
}

type DayOfWeek struct {
	DayOfWeek            DayOfWeek_Day `protobuf:"varint,1,opt,name=day_of_week,json=dayOfWeek,proto3,enum=enumeration.DayOfWeek_Day" json:"day_of_week,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DayOfWeek) Reset()         { *m = DayOfWeek{} }
func (m *DayOfWeek) String() string { return proto.CompactTextString(m) }
func (*DayOfWeek) ProtoMessage()    {}
func (*DayOfWeek) Descriptor() ([]byte, []int) {
	return fileDescriptor_466685131aab1de7, []int{0}
}

func (m *DayOfWeek) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DayOfWeek.Unmarshal(m, b)
}
func (m *DayOfWeek) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DayOfWeek.Marshal(b, m, deterministic)
}
func (m *DayOfWeek) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DayOfWeek.Merge(m, src)
}
func (m *DayOfWeek) XXX_Size() int {
	return xxx_messageInfo_DayOfWeek.Size(m)
}
func (m *DayOfWeek) XXX_DiscardUnknown() {
	xxx_messageInfo_DayOfWeek.DiscardUnknown(m)
}

var xxx_messageInfo_DayOfWeek proto.InternalMessageInfo

func (m *DayOfWeek) GetDayOfWeek() DayOfWeek_Day {
	if m != nil {
		return m.DayOfWeek
	}
	return DayOfWeek_UNDEFINED_DAY
}

func init() {
	proto.RegisterEnum("enumeration.DayOfWeek_Day", DayOfWeek_Day_name, DayOfWeek_Day_value)
	proto.RegisterType((*DayOfWeek)(nil), "enumeration.DayOfWeek")
}

func init() { proto.RegisterFile("day_of_week.proto", fileDescriptor_466685131aab1de7) }

var fileDescriptor_466685131aab1de7 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x49, 0xac, 0x8c,
	0xcf, 0x4f, 0x8b, 0x2f, 0x4f, 0x4d, 0xcd, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e,
	0xcd, 0x2b, 0xcd, 0x4d, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x53, 0xda, 0xca, 0xc8, 0xc5, 0xe9,
	0x92, 0x58, 0xe9, 0x9f, 0x16, 0x9e, 0x9a, 0x9a, 0x2d, 0x64, 0xc5, 0xc5, 0x8d, 0xa4, 0x5e, 0x82,
	0x51, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x4a, 0x0f, 0x49, 0x83, 0x1e, 0x5c, 0x31, 0x88, 0x15, 0xc4,
	0x99, 0x02, 0xe3, 0x2a, 0xe5, 0x70, 0x31, 0xbb, 0x24, 0x56, 0x0a, 0x09, 0x72, 0xf1, 0x86, 0xfa,
	0xb9, 0xb8, 0xba, 0x79, 0xfa, 0xb9, 0xba, 0xc4, 0xbb, 0x38, 0x46, 0x0a, 0x30, 0x08, 0x71, 0x73,
	0xb1, 0x07, 0xbb, 0xba, 0x87, 0xfa, 0xb9, 0x38, 0x0a, 0x30, 0x0a, 0x71, 0x72, 0xb1, 0x86, 0xb8,
	0x06, 0x39, 0x3b, 0x0a, 0x30, 0x09, 0x71, 0x71, 0xb1, 0x05, 0x86, 0x3a, 0x06, 0x85, 0x38, 0x0a,
	0x30, 0x43, 0xd8, 0x9e, 0x7e, 0x21, 0x8e, 0x02, 0x2c, 0x20, 0x25, 0xc1, 0xae, 0x11, 0x21, 0x8e,
	0x02, 0xac, 0x20, 0xe1, 0x60, 0x47, 0x27, 0x47, 0x17, 0x7f, 0x01, 0x36, 0x90, 0x31, 0x2e, 0xfe,
	0xbe, 0x9e, 0x7e, 0xee, 0xfe, 0x02, 0xec, 0x4e, 0x9c, 0x51, 0xec, 0x20, 0x57, 0x15, 0x17, 0x24,
	0x25, 0xb1, 0x81, 0xbd, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x51, 0x94, 0x22, 0x05, 0xeb,
	0x00, 0x00, 0x00,
}
