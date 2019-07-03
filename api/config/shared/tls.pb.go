// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/shared/tls.proto

package shared

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

type TLSCredentials struct {
	RootCertContents     string   `protobuf:"bytes,1,opt,name=root_cert_contents,json=rootCertContents,proto3" json:"root_cert_contents,omitempty" toml:"root_cert_contents,omitempty" mapstructure:"root_cert_contents,omitempty"`
	KeyContents          string   `protobuf:"bytes,2,opt,name=key_contents,json=keyContents,proto3" json:"key_contents,omitempty" toml:"key_contents,omitempty" mapstructure:"key_contents,omitempty"`
	CertContents         string   `protobuf:"bytes,3,opt,name=cert_contents,json=certContents,proto3" json:"cert_contents,omitempty" toml:"cert_contents,omitempty" mapstructure:"cert_contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *TLSCredentials) Reset()         { *m = TLSCredentials{} }
func (m *TLSCredentials) String() string { return proto.CompactTextString(m) }
func (*TLSCredentials) ProtoMessage()    {}
func (*TLSCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbee6a658298948, []int{0}
}

func (m *TLSCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TLSCredentials.Unmarshal(m, b)
}
func (m *TLSCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TLSCredentials.Marshal(b, m, deterministic)
}
func (m *TLSCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TLSCredentials.Merge(m, src)
}
func (m *TLSCredentials) XXX_Size() int {
	return xxx_messageInfo_TLSCredentials.Size(m)
}
func (m *TLSCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_TLSCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_TLSCredentials proto.InternalMessageInfo

func (m *TLSCredentials) GetRootCertContents() string {
	if m != nil {
		return m.RootCertContents
	}
	return ""
}

func (m *TLSCredentials) GetKeyContents() string {
	if m != nil {
		return m.KeyContents
	}
	return ""
}

func (m *TLSCredentials) GetCertContents() string {
	if m != nil {
		return m.CertContents
	}
	return ""
}

type FrontendTLSCredential struct {
	ServerName           string   `protobuf:"bytes,1,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty" toml:"server_name,omitempty" mapstructure:"server_name,omitempty"`
	Cert                 string   `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty" toml:"cert,omitempty" mapstructure:"cert,omitempty"`
	CertPath             string   `protobuf:"bytes,3,opt,name=cert_path,json=certPath,proto3" json:"cert_path,omitempty" toml:"cert_path,omitempty" mapstructure:"cert_path,omitempty"`
	Key                  string   `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty" toml:"key,omitempty" mapstructure:"key,omitempty"`
	KeyPath              string   `protobuf:"bytes,5,opt,name=key_path,json=keyPath,proto3" json:"key_path,omitempty" toml:"key_path,omitempty" mapstructure:"key_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FrontendTLSCredential) Reset()         { *m = FrontendTLSCredential{} }
func (m *FrontendTLSCredential) String() string { return proto.CompactTextString(m) }
func (*FrontendTLSCredential) ProtoMessage()    {}
func (*FrontendTLSCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbee6a658298948, []int{1}
}

func (m *FrontendTLSCredential) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrontendTLSCredential.Unmarshal(m, b)
}
func (m *FrontendTLSCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrontendTLSCredential.Marshal(b, m, deterministic)
}
func (m *FrontendTLSCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontendTLSCredential.Merge(m, src)
}
func (m *FrontendTLSCredential) XXX_Size() int {
	return xxx_messageInfo_FrontendTLSCredential.Size(m)
}
func (m *FrontendTLSCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontendTLSCredential.DiscardUnknown(m)
}

var xxx_messageInfo_FrontendTLSCredential proto.InternalMessageInfo

func (m *FrontendTLSCredential) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *FrontendTLSCredential) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func (m *FrontendTLSCredential) GetCertPath() string {
	if m != nil {
		return m.CertPath
	}
	return ""
}

func (m *FrontendTLSCredential) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FrontendTLSCredential) GetKeyPath() string {
	if m != nil {
		return m.KeyPath
	}
	return ""
}

func init() {
	proto.RegisterType((*TLSCredentials)(nil), "chef.automate.infra.config.TLSCredentials")
	proto.RegisterType((*FrontendTLSCredential)(nil), "chef.automate.infra.config.FrontendTLSCredential")
}

func init() { proto.RegisterFile("api/config/shared/tls.proto", fileDescriptor_2fbee6a658298948) }

var fileDescriptor_2fbee6a658298948 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x89, 0xad, 0xda, 0x4e, 0xab, 0x94, 0x05, 0xa1, 0xda, 0x83, 0x5a, 0x2f, 0x22, 0x25,
	0x7b, 0xf0, 0x0d, 0x0c, 0x78, 0x12, 0x11, 0xf5, 0xe4, 0x25, 0x6c, 0x37, 0x93, 0x26, 0xa4, 0xd9,
	0x0d, 0xb3, 0x53, 0x21, 0x2f, 0xe0, 0x23, 0xf8, 0xbc, 0xb2, 0xdb, 0xf8, 0xa7, 0x78, 0x9b, 0xfd,
	0xf6, 0xc7, 0xf7, 0xcd, 0x7c, 0x30, 0x53, 0x4d, 0x29, 0xb5, 0x35, 0x79, 0xb9, 0x92, 0xae, 0x50,
	0x84, 0x99, 0xe4, 0xb5, 0x8b, 0x1b, 0xb2, 0x6c, 0xc5, 0x99, 0x2e, 0x30, 0x8f, 0xd5, 0x86, 0x6d,
	0xad, 0x18, 0xe3, 0xd2, 0xe4, 0xa4, 0xe2, 0x2d, 0x3c, 0xff, 0x88, 0xe0, 0xf8, 0xf5, 0xe1, 0x25,
	0x21, 0xcc, 0xd0, 0x70, 0xa9, 0xd6, 0x4e, 0x2c, 0x40, 0x90, 0xb5, 0x9c, 0x6a, 0x24, 0x4e, 0xb5,
	0x35, 0x8c, 0x86, 0xdd, 0x34, 0xba, 0x88, 0xae, 0x87, 0xcf, 0x13, 0xff, 0x93, 0x20, 0x71, 0xd2,
	0xe9, 0xe2, 0x12, 0xc6, 0x15, 0xb6, 0xbf, 0xdc, 0x5e, 0xe0, 0x46, 0x15, 0xb6, 0x3f, 0xc8, 0x15,
	0x1c, 0xed, 0x7a, 0xf5, 0x02, 0x33, 0xd6, 0x7f, 0x7c, 0xe6, 0x9f, 0x11, 0x9c, 0xdc, 0x53, 0x78,
	0x65, 0x3b, 0x0b, 0x89, 0x73, 0x18, 0x39, 0xa4, 0x77, 0xa4, 0xd4, 0xa8, 0x1a, 0xbb, 0x45, 0x60,
	0x2b, 0x3d, 0xaa, 0x1a, 0x85, 0x80, 0xbe, 0xb7, 0xea, 0xa2, 0xc3, 0x2c, 0x66, 0x30, 0x0c, 0x99,
	0x8d, 0xe2, 0xa2, 0xcb, 0x1b, 0x78, 0xe1, 0x49, 0x71, 0x21, 0x26, 0xd0, 0xab, 0xb0, 0x9d, 0xf6,
	0x83, 0xec, 0x47, 0x71, 0x0a, 0x03, 0x7f, 0x45, 0xa0, 0xf7, 0x83, 0x7c, 0x58, 0x61, 0xeb, 0xe1,
	0xbb, 0xc5, 0xdb, 0xcd, 0xaa, 0xe4, 0x62, 0xb3, 0x8c, 0xb5, 0xad, 0xa5, 0xaf, 0x52, 0x7e, 0x57,
	0x29, 0xff, 0xb5, 0xbe, 0x3c, 0x08, 0x95, 0xdf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x5d,
	0x56, 0xae, 0x91, 0x01, 0x00, 0x00,
}
