// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/cfgmgmt/response/stats.proto

package response

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

type RunsCounts struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty" toml:"total,omitempty" mapstructure:"total,omitempty"`
	Success              int32    `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty" toml:"success,omitempty" mapstructure:"success,omitempty"`
	Failure              int32    `protobuf:"varint,3,opt,name=failure,proto3" json:"failure,omitempty" toml:"failure,omitempty" mapstructure:"failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *RunsCounts) Reset()         { *m = RunsCounts{} }
func (m *RunsCounts) String() string { return proto.CompactTextString(m) }
func (*RunsCounts) ProtoMessage()    {}
func (*RunsCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_97bc85b7bb3c746f, []int{0}
}

func (m *RunsCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunsCounts.Unmarshal(m, b)
}
func (m *RunsCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunsCounts.Marshal(b, m, deterministic)
}
func (m *RunsCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunsCounts.Merge(m, src)
}
func (m *RunsCounts) XXX_Size() int {
	return xxx_messageInfo_RunsCounts.Size(m)
}
func (m *RunsCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_RunsCounts.DiscardUnknown(m)
}

var xxx_messageInfo_RunsCounts proto.InternalMessageInfo

func (m *RunsCounts) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RunsCounts) GetSuccess() int32 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *RunsCounts) GetFailure() int32 {
	if m != nil {
		return m.Failure
	}
	return 0
}

type NodesCounts struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty" toml:"total,omitempty" mapstructure:"total,omitempty"`
	Success              int32    `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty" toml:"success,omitempty" mapstructure:"success,omitempty"`
	Failure              int32    `protobuf:"varint,3,opt,name=failure,proto3" json:"failure,omitempty" toml:"failure,omitempty" mapstructure:"failure,omitempty"`
	Missing              int32    `protobuf:"varint,4,opt,name=missing,proto3" json:"missing,omitempty" toml:"missing,omitempty" mapstructure:"missing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *NodesCounts) Reset()         { *m = NodesCounts{} }
func (m *NodesCounts) String() string { return proto.CompactTextString(m) }
func (*NodesCounts) ProtoMessage()    {}
func (*NodesCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_97bc85b7bb3c746f, []int{1}
}

func (m *NodesCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesCounts.Unmarshal(m, b)
}
func (m *NodesCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesCounts.Marshal(b, m, deterministic)
}
func (m *NodesCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesCounts.Merge(m, src)
}
func (m *NodesCounts) XXX_Size() int {
	return xxx_messageInfo_NodesCounts.Size(m)
}
func (m *NodesCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesCounts.DiscardUnknown(m)
}

var xxx_messageInfo_NodesCounts proto.InternalMessageInfo

func (m *NodesCounts) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *NodesCounts) GetSuccess() int32 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *NodesCounts) GetFailure() int32 {
	if m != nil {
		return m.Failure
	}
	return 0
}

func (m *NodesCounts) GetMissing() int32 {
	if m != nil {
		return m.Missing
	}
	return 0
}

func init() {
	proto.RegisterType((*RunsCounts)(nil), "chef.automate.domain.cfgmgmt.response.RunsCounts")
	proto.RegisterType((*NodesCounts)(nil), "chef.automate.domain.cfgmgmt.response.NodesCounts")
}

func init() {
	proto.RegisterFile("api/interservice/cfgmgmt/response/stats.proto", fileDescriptor_97bc85b7bb3c746f)
}

var fileDescriptor_97bc85b7bb3c746f = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xbf, 0x4b, 0xc0, 0x30,
	0x10, 0x85, 0xa9, 0x5a, 0x85, 0xb8, 0x15, 0x87, 0x8c, 0x52, 0x10, 0x5c, 0x4c, 0x06, 0x37, 0x71,
	0xd2, 0xdd, 0xa1, 0x83, 0x83, 0x5b, 0x9a, 0x5e, 0xd3, 0x83, 0x26, 0x57, 0x73, 0x17, 0xff, 0x7e,
	0xe9, 0xaf, 0xd9, 0xc5, 0xf1, 0xbb, 0x8f, 0x77, 0xf0, 0x9e, 0x7a, 0x72, 0x0b, 0x5a, 0x4c, 0x02,
	0x99, 0x21, 0xff, 0xa0, 0x07, 0xeb, 0xc7, 0x10, 0x43, 0x14, 0x9b, 0x81, 0x17, 0x4a, 0x0c, 0x96,
	0xc5, 0x09, 0x9b, 0x25, 0x93, 0x50, 0xf3, 0xe0, 0x27, 0x18, 0x8d, 0x2b, 0x42, 0xd1, 0x09, 0x98,
	0x81, 0xa2, 0xc3, 0x64, 0x8e, 0x88, 0x39, 0x23, 0xed, 0xa7, 0x52, 0x5d, 0x49, 0xfc, 0x4e, 0x25,
	0x09, 0x37, 0x77, 0xaa, 0x16, 0x12, 0x37, 0xeb, 0xea, 0xbe, 0x7a, 0xac, 0xbb, 0x1d, 0x1a, 0xad,
	0x6e, 0xb8, 0x78, 0x0f, 0xcc, 0xfa, 0x62, 0xbb, 0x9f, 0xb8, 0x9a, 0xd1, 0xe1, 0x5c, 0x32, 0xe8,
	0xcb, 0xdd, 0x1c, 0xd8, 0x7e, 0xab, 0xdb, 0x0f, 0x1a, 0xe0, 0xdf, 0x1f, 0xaf, 0x26, 0x22, 0x33,
	0xa6, 0xa0, 0xaf, 0x76, 0x73, 0xe0, 0xdb, 0xeb, 0xd7, 0x4b, 0x40, 0x99, 0x4a, 0x6f, 0x3c, 0x45,
	0xbb, 0xd6, 0xb7, 0x67, 0x7d, 0xfb, 0xe7, 0x76, 0xfd, 0xf5, 0x36, 0xdb, 0xf3, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x65, 0xcc, 0x95, 0xc9, 0x67, 0x01, 0x00, 0x00,
}
