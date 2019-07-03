// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/secrets/config_request.proto

package secrets

import (
	fmt "fmt"
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,3,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_835457c260c4333a, []int{0}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_835457c260c4333a, []int{0, 0}
}

func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(m, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Service              *ConfigRequest_V1_System_Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Tls                  *shared.TLSCredentials           `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	SecretsKey           *wrappers.StringValue            `protobuf:"bytes,4,opt,name=secrets_key,json=secretsKey,proto3" json:"secrets_key,omitempty" toml:"secrets_key,omitempty" mapstructure:"secrets_key,omitempty"`
	Storage              *ConfigRequest_V1_System_Storage `protobuf:"bytes,5,opt,name=storage,proto3" json:"storage,omitempty" toml:"storage,omitempty" mapstructure:"storage,omitempty"`
	Log                  *ConfigRequest_V1_System_Log     `protobuf:"bytes,6,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_835457c260c4333a, []int{0, 0, 0}
}

func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(m, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetSecretsKey() *wrappers.StringValue {
	if m != nil {
		return m.SecretsKey
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetStorage() *ConfigRequest_V1_System_Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_835457c260c4333a, []int{0, 0, 0, 0}
}

func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Format               *wrappers.StringValue `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	Level                *wrappers.StringValue `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_835457c260c4333a, []int{0, 0, 0, 1}
}

func (m *ConfigRequest_V1_System_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Log.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Log) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Size(m)
}
func (m *ConfigRequest_V1_System_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Log.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Log proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Log) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

type ConfigRequest_V1_System_Storage struct {
	Database             *wrappers.StringValue `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty" toml:"database,omitempty" mapstructure:"database,omitempty"`
	User                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" toml:"user,omitempty" mapstructure:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Storage) Reset()         { *m = ConfigRequest_V1_System_Storage{} }
func (m *ConfigRequest_V1_System_Storage) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Storage) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_835457c260c4333a, []int{0, 0, 0, 2}
}

func (m *ConfigRequest_V1_System_Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Storage.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Storage.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Storage.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Storage) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Storage.Size(m)
}
func (m *ConfigRequest_V1_System_Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Storage proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Storage) GetDatabase() *wrappers.StringValue {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *ConfigRequest_V1_System_Storage) GetUser() *wrappers.StringValue {
	if m != nil {
		return m.User
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_835457c260c4333a, []int{0, 0, 1}
}

func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.secrets.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.secrets.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.secrets.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.domain.secrets.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.domain.secrets.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_System_Storage)(nil), "chef.automate.domain.secrets.ConfigRequest.V1.System.Storage")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.secrets.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("api/config/secrets/config_request.proto", fileDescriptor_835457c260c4333a)
}

var fileDescriptor_835457c260c4333a = []byte{
	// 577 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xd1, 0x6a, 0x13, 0x4f,
	0x14, 0xc6, 0x69, 0x76, 0x9b, 0xe4, 0x3f, 0xe5, 0x8f, 0x75, 0x40, 0x59, 0xb6, 0xa5, 0x14, 0x6f,
	0x14, 0x65, 0x67, 0x4d, 0x5a, 0x45, 0xc5, 0x7a, 0xd1, 0x5e, 0x54, 0xdb, 0x88, 0xb0, 0x91, 0x08,
	0xde, 0x94, 0xc9, 0xe6, 0x64, 0xb2, 0x38, 0x3b, 0xb3, 0xce, 0xcc, 0x46, 0xf2, 0x0a, 0xde, 0xfa,
	0x20, 0x3e, 0x81, 0x37, 0x7d, 0x06, 0xdf, 0xa4, 0x2f, 0x20, 0x3b, 0x3b, 0x09, 0x6a, 0x20, 0xa4,
	0xf5, 0x72, 0x39, 0xe7, 0xfb, 0x9d, 0x73, 0xbe, 0x8f, 0x1d, 0x74, 0x9f, 0x16, 0x59, 0x9c, 0x4a,
	0x31, 0xce, 0x58, 0xac, 0x21, 0x55, 0x60, 0xb4, 0xfb, 0xbc, 0x50, 0xf0, 0xb9, 0x04, 0x6d, 0x48,
	0xa1, 0xa4, 0x91, 0x78, 0x37, 0x9d, 0xc0, 0x98, 0xd0, 0xd2, 0xc8, 0x9c, 0x1a, 0x20, 0x23, 0x99,
	0xd3, 0x4c, 0x10, 0x27, 0x09, 0xf7, 0x7e, 0xc7, 0x4c, 0xa8, 0x82, 0x51, 0xcc, 0xb8, 0x1c, 0x52,
	0x5e, 0xab, 0xc3, 0x9d, 0xe5, 0xba, 0xe1, 0xda, 0x15, 0xcf, 0x52, 0x99, 0x17, 0x52, 0x80, 0x30,
	0x3a, 0x9e, 0x0f, 0x88, 0x98, 0x2a, 0xd2, 0xd8, 0xd6, 0xd3, 0x88, 0x81, 0x88, 0x68, 0x37, 0x72,
	0xfa, 0x0a, 0x45, 0xbb, 0xd5, 0x47, 0x4c, 0x85, 0x90, 0x86, 0x9a, 0x4c, 0x8a, 0x39, 0x6b, 0x8f,
	0x49, 0xc9, 0x38, 0xd4, 0xca, 0x61, 0x39, 0x8e, 0xbf, 0x28, 0x5a, 0x14, 0xa0, 0x5c, 0xfd, 0xde,
	0x8f, 0x36, 0xfa, 0xff, 0xc4, 0x72, 0x92, 0xfa, 0x3c, 0xfc, 0x0a, 0x35, 0xa6, 0x9d, 0xc0, 0xdb,
	0xdf, 0x78, 0xb0, 0xd5, 0x25, 0x64, 0xd5, 0x95, 0xe4, 0x0f, 0x21, 0x19, 0x74, 0x92, 0xc6, 0xb4,
	0x13, 0xfe, 0x6c, 0xa1, 0xc6, 0xa0, 0x83, 0x4f, 0x91, 0xa7, 0x67, 0x3a, 0xd8, 0xb0, 0x9c, 0x27,
	0xd7, 0xe3, 0x90, 0xfe, 0x4c, 0x1b, 0xc8, 0x93, 0x8a, 0x80, 0x5f, 0x23, 0x4f, 0x4f, 0xd3, 0xa0,
	0x61, 0x41, 0x4f, 0xaf, 0x0b, 0x02, 0x35, 0xcd, 0x52, 0x48, 0x2a, 0x44, 0xf8, 0xbd, 0x89, 0x9a,
	0x35, 0x19, 0x1f, 0x22, 0x3f, 0xe7, 0x9a, 0xba, 0xf5, 0xf6, 0xff, 0xa2, 0x66, 0x62, 0xac, 0x28,
	0xa9, 0xed, 0x25, 0x6f, 0xb9, 0xa6, 0x89, 0xed, 0xc6, 0x1f, 0x50, 0x4b, 0xd7, 0x40, 0xb7, 0xce,
	0xd1, 0x8d, 0xee, 0x5a, 0x6c, 0x35, 0xa7, 0xe1, 0x97, 0xc8, 0x33, 0x5c, 0x3b, 0xd3, 0x1f, 0xae,
	0xda, 0xe6, 0x7d, 0xaf, 0x7f, 0xa2, 0x60, 0x04, 0xc2, 0x64, 0x94, 0xeb, 0xa4, 0x92, 0xe1, 0x23,
	0xb4, 0xe5, 0x26, 0x5e, 0x7c, 0x82, 0x59, 0xe0, 0x5b, 0xca, 0x2e, 0xa9, 0x93, 0x27, 0xf3, 0xe4,
	0x49, 0xdf, 0xa8, 0x4c, 0xb0, 0x01, 0xe5, 0x25, 0x24, 0xc8, 0x09, 0xce, 0x61, 0x66, 0xaf, 0x32,
	0x52, 0x51, 0x06, 0xc1, 0xe6, 0x3f, 0x5d, 0x55, 0x43, 0x92, 0x39, 0x0d, 0x9f, 0x23, 0x8f, 0x4b,
	0x16, 0x34, 0x2d, 0xf4, 0xf9, 0xcd, 0xa0, 0x3d, 0xc9, 0x92, 0x8a, 0x12, 0x7e, 0xdd, 0x40, 0x2d,
	0xe7, 0x1b, 0x7e, 0x8c, 0xfc, 0x89, 0xd4, 0xc6, 0xa5, 0xb7, 0xfa, 0x52, 0xdb, 0x89, 0x4f, 0x91,
	0x5f, 0x48, 0x65, 0x5c, 0x6c, 0x3b, 0x4b, 0x8a, 0x37, 0xc2, 0x1c, 0x74, 0xad, 0xe0, 0xf8, 0xee,
	0xe5, 0x55, 0x80, 0x17, 0x41, 0x6f, 0x7f, 0x7b, 0x17, 0xfa, 0xd5, 0xbf, 0x97, 0x58, 0xc0, 0x99,
	0xdf, 0xf6, 0xb6, 0xfd, 0x50, 0x22, 0xaf, 0x27, 0x19, 0x3e, 0x44, 0xcd, 0xb1, 0x54, 0x39, 0x5d,
	0x6f, 0x13, 0xd7, 0x8b, 0xbb, 0x68, 0x93, 0xc3, 0x14, 0xb8, 0x5b, 0x66, 0xb5, 0xa8, 0x6e, 0x0d,
	0x4b, 0xd4, 0x72, 0xf6, 0xe2, 0x67, 0xa8, 0x3d, 0xa2, 0x86, 0x0e, 0xa9, 0x86, 0xb5, 0xc6, 0x2e,
	0xba, 0x2b, 0xdb, 0x4a, 0x0d, 0x6a, 0xad, 0xb9, 0xb6, 0x33, 0xfc, 0x6f, 0xe1, 0xf9, 0x8b, 0x3b,
	0x97, 0x57, 0xc1, 0x6d, 0x74, 0xcb, 0xe5, 0x15, 0x39, 0x7b, 0x8e, 0xa3, 0x8f, 0x8f, 0x58, 0x66,
	0x26, 0xe5, 0x90, 0xa4, 0x32, 0x8f, 0xab, 0x88, 0x17, 0x4f, 0x56, 0xbc, 0xfc, 0x94, 0x0e, 0x9b,
	0x76, 0xd8, 0xc1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xc1, 0xeb, 0x20, 0x67, 0x05, 0x00,
	0x00,
}
