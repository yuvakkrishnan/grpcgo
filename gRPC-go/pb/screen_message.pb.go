// Code generated by protoc-gen-go. DO NOT EDIT.
// source: screen_message.proto

package pb

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

type Screen_Panel int32

const (
	Screen_UNKNOWN Screen_Panel = 0
	Screen_IPS     Screen_Panel = 1
	Screen_OLED    Screen_Panel = 2
)

var Screen_Panel_name = map[int32]string{
	0: "UNKNOWN",
	1: "IPS",
	2: "OLED",
}

var Screen_Panel_value = map[string]int32{
	"UNKNOWN": 0,
	"IPS":     1,
	"OLED":    2,
}

func (x Screen_Panel) String() string {
	return proto.EnumName(Screen_Panel_name, int32(x))
}

func (Screen_Panel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

type Screen struct {
	SizeInch             float32            `protobuf:"fixed32,1,opt,name=size_inch,json=sizeInch,proto3" json:"size_inch,omitempty"`
	Resolution           *Screen_Resolution `protobuf:"bytes,2,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Panel                Screen_Panel       `protobuf:"varint,3,opt,name=panel,proto3,enum=spacecode.gogrpc.Screen_Panel" json:"panel,omitempty"`
	Multitouch           bool               `protobuf:"varint,4,opt,name=multitouch,proto3" json:"multitouch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Screen) Reset()         { *m = Screen{} }
func (m *Screen) String() string { return proto.CompactTextString(m) }
func (*Screen) ProtoMessage()    {}
func (*Screen) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0}
}

func (m *Screen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen.Unmarshal(m, b)
}
func (m *Screen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen.Marshal(b, m, deterministic)
}
func (m *Screen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen.Merge(m, src)
}
func (m *Screen) XXX_Size() int {
	return xxx_messageInfo_Screen.Size(m)
}
func (m *Screen) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen.DiscardUnknown(m)
}

var xxx_messageInfo_Screen proto.InternalMessageInfo

func (m *Screen) GetSizeInch() float32 {
	if m != nil {
		return m.SizeInch
	}
	return 0
}

func (m *Screen) GetResolution() *Screen_Resolution {
	if m != nil {
		return m.Resolution
	}
	return nil
}

func (m *Screen) GetPanel() Screen_Panel {
	if m != nil {
		return m.Panel
	}
	return Screen_UNKNOWN
}

func (m *Screen) GetMultitouch() bool {
	if m != nil {
		return m.Multitouch
	}
	return false
}

type Screen_Resolution struct {
	Width                uint32   `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Screen_Resolution) Reset()         { *m = Screen_Resolution{} }
func (m *Screen_Resolution) String() string { return proto.CompactTextString(m) }
func (*Screen_Resolution) ProtoMessage()    {}
func (*Screen_Resolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2814cd02b8aba7, []int{0, 0}
}

func (m *Screen_Resolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Screen_Resolution.Unmarshal(m, b)
}
func (m *Screen_Resolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Screen_Resolution.Marshal(b, m, deterministic)
}
func (m *Screen_Resolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Screen_Resolution.Merge(m, src)
}
func (m *Screen_Resolution) XXX_Size() int {
	return xxx_messageInfo_Screen_Resolution.Size(m)
}
func (m *Screen_Resolution) XXX_DiscardUnknown() {
	xxx_messageInfo_Screen_Resolution.DiscardUnknown(m)
}

var xxx_messageInfo_Screen_Resolution proto.InternalMessageInfo

func (m *Screen_Resolution) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Screen_Resolution) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterEnum("spacecode.gogrpc.Screen_Panel", Screen_Panel_name, Screen_Panel_value)
	proto.RegisterType((*Screen)(nil), "spacecode.gogrpc.Screen")
	proto.RegisterType((*Screen_Resolution)(nil), "spacecode.gogrpc.Screen.Resolution")
}

func init() {
	proto.RegisterFile("screen_message.proto", fileDescriptor_8a2814cd02b8aba7)
}

var fileDescriptor_8a2814cd02b8aba7 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0x49, 0x93, 0x34, 0xdf, 0x2d, 0x95, 0x70, 0x29, 0x12, 0x14, 0x4a, 0xa8, 0x0b,
	0xb3, 0xca, 0xa2, 0xba, 0x72, 0xe9, 0x9f, 0x45, 0x51, 0xd2, 0x32, 0x45, 0x04, 0x37, 0x25, 0x9d,
	0x5e, 0x92, 0x81, 0x34, 0x33, 0x64, 0x26, 0x08, 0xbe, 0x8f, 0xef, 0x29, 0x4d, 0x44, 0x8b, 0xe0,
	0xf2, 0x1c, 0x7e, 0xe7, 0x9c, 0x99, 0x0b, 0x13, 0x23, 0x1a, 0xa2, 0x7a, 0xb3, 0x27, 0x63, 0xf2,
	0x82, 0x52, 0xdd, 0x28, 0xab, 0x30, 0x34, 0x3a, 0x17, 0x24, 0xd4, 0x8e, 0xd2, 0x42, 0x15, 0x8d,
	0x16, 0xb3, 0x0f, 0x07, 0xfc, 0x75, 0x87, 0xe2, 0x39, 0xfc, 0x37, 0xf2, 0x9d, 0x36, 0xb2, 0x16,
	0x65, 0xc4, 0x62, 0x96, 0x38, 0x3c, 0x38, 0x18, 0x8b, 0x5a, 0x94, 0x78, 0x07, 0xd0, 0x90, 0x51,
	0x55, 0x6b, 0xa5, 0xaa, 0x23, 0x27, 0x66, 0xc9, 0x68, 0x7e, 0x91, 0xfe, 0xae, 0x4b, 0xfb, 0xaa,
	0x94, 0x7f, 0xa3, 0xfc, 0x28, 0x86, 0xd7, 0xe0, 0xe9, 0xbc, 0xa6, 0x2a, 0x1a, 0xc4, 0x2c, 0x39,
	0x99, 0x4f, 0xff, 0xcc, 0xaf, 0x0e, 0x14, 0xef, 0x61, 0x9c, 0x02, 0xec, 0xdb, 0xca, 0x4a, 0xab,
	0x5a, 0x51, 0x46, 0x6e, 0xcc, 0x92, 0x80, 0x1f, 0x39, 0x67, 0x37, 0x00, 0x3f, 0x7b, 0x38, 0x01,
	0xef, 0x4d, 0xee, 0x6c, 0xff, 0x83, 0x31, 0xef, 0x05, 0x9e, 0x82, 0x5f, 0x92, 0x2c, 0x4a, 0xdb,
	0x3d, 0x7d, 0xcc, 0xbf, 0xd4, 0xec, 0x12, 0xbc, 0x6e, 0x0b, 0x47, 0x30, 0x7c, 0xce, 0x1e, 0xb3,
	0xe5, 0x4b, 0x16, 0xfe, 0xc3, 0x21, 0x0c, 0x16, 0xab, 0x75, 0xc8, 0x30, 0x00, 0x77, 0xf9, 0xf4,
	0x70, 0x1f, 0x3a, 0xb7, 0xee, 0xab, 0xa3, 0xb7, 0x5b, 0xbf, 0x3b, 0xe3, 0xd5, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbf, 0x7e, 0x5f, 0x91, 0x5e, 0x01, 0x00, 0x00,
}
