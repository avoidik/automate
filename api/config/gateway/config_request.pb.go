// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/gateway/config_request.proto

package gateway

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
	return fileDescriptor_9be4cf630cc5f57c, []int{0}
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
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0}
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
	Tls                  *shared.TLSCredentials           `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Service              *ConfigRequest_V1_System_Service `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Log                  *ConfigRequest_V1_System_Log     `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0, 0}
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

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
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
	Host                 *wrappers.StringValue        `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value         `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	ExternalFqdn         *wrappers.StringValue        `protobuf:"bytes,3,opt,name=external_fqdn,json=externalFqdn,proto3" json:"external_fqdn,omitempty" toml:"external_fqdn,omitempty" mapstructure:"external_fqdn,omitempty"`
	GrpcPort             *wrappers.Int32Value         `protobuf:"bytes,4,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty" toml:"grpc_port,omitempty" mapstructure:"grpc_port,omitempty"`
	TrialLicenseUrl      *wrappers.StringValue        `protobuf:"bytes,5,opt,name=trial_license_url,json=trialLicenseUrl,proto3" json:"trial_license_url,omitempty" toml:"trial_license_url,omitempty" mapstructure:"trial_license_url,omitempty"`
	Log                  *ConfigRequest_V1_System_Log `protobuf:"bytes,6,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	AuthMiddleware       *wrappers.StringValue        `protobuf:"bytes,7,opt,name=auth_middleware,json=authMiddleware,proto3" json:"auth_middleware,omitempty" toml:"auth_middleware,omitempty" mapstructure:"auth_middleware,omitempty"` // Deprecated: Do not use.
	EnableAppsFeature    *wrappers.BoolValue          `protobuf:"bytes,9,opt,name=enable_apps_feature,json=enableAppsFeature,proto3" json:"enable_apps_feature,omitempty" toml:"enable_apps_feature,omitempty" mapstructure:"enable_apps_feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                       `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                        `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0, 0, 0}
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

func (m *ConfigRequest_V1_System_Service) GetExternalFqdn() *wrappers.StringValue {
	if m != nil {
		return m.ExternalFqdn
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetGrpcPort() *wrappers.Int32Value {
	if m != nil {
		return m.GrpcPort
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetTrialLicenseUrl() *wrappers.StringValue {
	if m != nil {
		return m.TrialLicenseUrl
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

// Deprecated: Do not use.
func (m *ConfigRequest_V1_System_Service) GetAuthMiddleware() *wrappers.StringValue {
	if m != nil {
		return m.AuthMiddleware
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetEnableAppsFeature() *wrappers.BoolValue {
	if m != nil {
		return m.EnableAppsFeature
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Level                *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	Format               *wrappers.StringValue `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0, 0, 1}
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

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
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
	return fileDescriptor_9be4cf630cc5f57c, []int{0, 0, 1}
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
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.api.config.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.api.config.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.api.config.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.api.config.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.api.config.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.api.config.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("api/config/gateway/config_request.proto", fileDescriptor_9be4cf630cc5f57c)
}

var fileDescriptor_9be4cf630cc5f57c = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0xd5, 0x26, 0x6b, 0x57, 0xc3, 0x68, 0x67, 0x10, 0x8a, 0x32, 0x34, 0x4d, 0x5c, 0x40,
	0x43, 0x4d, 0x68, 0x37, 0x0e, 0x4c, 0x5c, 0xd6, 0xc1, 0x60, 0x53, 0x27, 0x4d, 0x29, 0xec, 0xc0,
	0x25, 0x72, 0xd3, 0x27, 0x69, 0x24, 0xc7, 0xce, 0x6c, 0xa7, 0x63, 0x1f, 0x8f, 0x7d, 0x01, 0xe0,
	0xc6, 0x87, 0xe0, 0xb6, 0x2f, 0x80, 0x12, 0x27, 0xe5, 0xa5, 0x62, 0x2f, 0xe2, 0x52, 0xd5, 0xf2,
	0xf3, 0xfb, 0xf5, 0xef, 0xc7, 0x4f, 0x8d, 0x9e, 0x90, 0x34, 0x76, 0x03, 0xce, 0xc2, 0x38, 0x72,
	0x23, 0xa2, 0xe0, 0x8c, 0x9c, 0x97, 0x4b, 0x5f, 0xc0, 0x69, 0x06, 0x52, 0x39, 0xa9, 0xe0, 0x8a,
	0x63, 0x2b, 0x98, 0x42, 0xe8, 0x90, 0x4c, 0xf1, 0x84, 0x28, 0x70, 0x48, 0x1a, 0x3b, 0xba, 0xce,
	0x5e, 0xff, 0x4d, 0x21, 0xa7, 0x44, 0xc0, 0xc4, 0x8d, 0x28, 0x1f, 0x13, 0xaa, 0x49, 0x7b, 0x6d,
	0x71, 0x5f, 0x51, 0x59, 0x6e, 0x1e, 0x06, 0x3c, 0x49, 0x39, 0x03, 0xa6, 0xa4, 0x5b, 0xc9, 0xbb,
	0x91, 0x48, 0x03, 0xb7, 0xd8, 0x0f, 0xba, 0x11, 0xb0, 0x2e, 0xe9, 0x77, 0x4b, 0x3e, 0x57, 0x91,
	0x7e, 0xbe, 0x70, 0x09, 0x63, 0x5c, 0x11, 0x15, 0x73, 0x56, 0xb9, 0xd6, 0x23, 0xce, 0x23, 0x0a,
	0x9a, 0x1c, 0x67, 0xa1, 0x7b, 0x26, 0x48, 0x9a, 0x82, 0x28, 0xf7, 0x1f, 0x7f, 0x6f, 0xa1, 0x95,
	0xbd, 0xc2, 0xe3, 0xe9, 0xa3, 0xe1, 0x1d, 0x54, 0x9f, 0xf5, 0x2c, 0x63, 0xa3, 0xf6, 0xf4, 0x4e,
	0x7f, 0xd3, 0xf9, 0xd7, 0x09, 0x9d, 0x3f, 0x20, 0xe7, 0xa4, 0xe7, 0xd5, 0x67, 0x3d, 0xfb, 0xc7,
	0x32, 0xaa, 0x9f, 0xf4, 0xf0, 0x1e, 0x32, 0xe4, 0xb9, 0xb4, 0x6a, 0x85, 0xa3, 0x77, 0x73, 0x87,
	0x33, 0x3a, 0x97, 0x0a, 0x12, 0x2f, 0xa7, 0xf1, 0x6b, 0x64, 0xc8, 0x59, 0x60, 0xd5, 0x0b, 0x49,
	0xff, 0x36, 0x12, 0x10, 0xb3, 0x38, 0x00, 0x2f, 0xc7, 0xed, 0x2f, 0x4d, 0xd4, 0xd0, 0x56, 0xbc,
	0x8d, 0xcc, 0x84, 0x4a, 0x52, 0xc6, 0xda, 0xf8, 0xcb, 0x18, 0xb3, 0x50, 0x90, 0xca, 0x79, 0x44,
	0x25, 0xf1, 0x8a, 0x6a, 0xfc, 0x0a, 0x19, 0x8a, 0xca, 0x32, 0xc6, 0xe6, 0x55, 0xd0, 0xfb, 0xe1,
	0x68, 0x4f, 0xc0, 0x04, 0x98, 0x8a, 0x09, 0x95, 0x5e, 0x8e, 0xe1, 0x11, 0x6a, 0x4a, 0x1d, 0xa7,
	0xec, 0xe8, 0xcb, 0x5b, 0x77, 0x63, 0x7e, 0x9e, 0xca, 0x84, 0xdf, 0x22, 0x83, 0xf2, 0xc8, 0x32,
	0x0b, 0xe1, 0x8b, 0xdb, 0x0b, 0x87, 0x3c, 0xf2, 0x72, 0x83, 0xfd, 0xd9, 0x44, 0xcd, 0xd2, 0x8e,
	0x9f, 0x23, 0x73, 0xca, 0xa5, 0x2a, 0xbb, 0xf3, 0xc8, 0xd1, 0x73, 0xe3, 0x54, 0x73, 0xe3, 0x8c,
	0x94, 0x88, 0x59, 0x74, 0x42, 0x68, 0x06, 0x5e, 0x51, 0x89, 0xdf, 0x20, 0x33, 0xe5, 0x42, 0x95,
	0xad, 0x59, 0x5b, 0x20, 0x0e, 0x98, 0xda, 0xea, 0x17, 0xc0, 0xe0, 0xc1, 0xc5, 0xa5, 0xd5, 0x41,
	0xe6, 0x54, 0xa9, 0xb4, 0xf3, 0xb5, 0x6d, 0x2f, 0xe5, 0x5f, 0xa4, 0x57, 0xe0, 0x78, 0x17, 0xad,
	0xc0, 0x27, 0x05, 0x82, 0x11, 0xea, 0x87, 0xa7, 0x13, 0x56, 0x36, 0xea, 0xea, 0x04, 0x77, 0x2b,
	0x64, 0xff, 0x74, 0xc2, 0xf0, 0x31, 0x6a, 0xe5, 0x7f, 0x0f, 0xbf, 0x88, 0x63, 0x5e, 0x1f, 0xe7,
	0xe1, 0xc5, 0xa5, 0x85, 0xe7, 0x37, 0xd3, 0xf9, 0xd6, 0xb6, 0xcd, 0x9c, 0xf7, 0x96, 0xf3, 0xcf,
	0xe3, 0x3c, 0xd4, 0x3b, 0xb4, 0xaa, 0x44, 0x4c, 0xa8, 0x4f, 0xe3, 0x00, 0x98, 0x04, 0x3f, 0x13,
	0xd4, 0x5a, 0xba, 0x41, 0xb0, 0x76, 0x81, 0x0d, 0x35, 0xf5, 0x41, 0xd0, 0xea, 0xb2, 0x1a, 0xff,
	0x7b, 0x59, 0xf8, 0x00, 0xb5, 0x49, 0xa6, 0xa6, 0x7e, 0x12, 0x4f, 0x26, 0x14, 0xce, 0x88, 0x00,
	0xab, 0x79, 0x7d, 0xa0, 0x41, 0xdd, 0xaa, 0x79, 0xf7, 0x72, 0xf0, 0x68, 0xce, 0xe1, 0x43, 0x74,
	0x1f, 0x18, 0x19, 0x53, 0xf0, 0x49, 0x9a, 0x4a, 0x3f, 0x04, 0xa2, 0x32, 0x01, 0x56, 0xab, 0xd0,
	0xd9, 0x0b, 0xba, 0x01, 0xe7, 0x54, 0x9f, 0x6e, 0x55, 0x63, 0xbb, 0x69, 0x2a, 0xf7, 0x35, 0x64,
	0x73, 0x64, 0x0c, 0x79, 0x84, 0xfb, 0x68, 0x89, 0xc2, 0x0c, 0xe8, 0x8d, 0xe6, 0x47, 0x97, 0xe2,
	0x6d, 0xd4, 0x08, 0xb9, 0x48, 0x48, 0x35, 0x42, 0x57, 0x43, 0x65, 0xad, 0xdd, 0x9a, 0xcf, 0xec,
	0x8e, 0xbe, 0xc7, 0xce, 0xaf, 0x37, 0x52, 0x3f, 0xd6, 0x87, 0xe6, 0x72, 0xad, 0x63, 0x0c, 0xba,
	0x1f, 0x9f, 0x45, 0xb1, 0x9a, 0x66, 0x63, 0x27, 0xe0, 0x89, 0x9b, 0x37, 0x7e, 0xfe, 0x9a, 0xba,
	0x8b, 0x2f, 0xfc, 0xb8, 0x51, 0xfc, 0xea, 0xd6, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x4e,
	0x68, 0xcf, 0xfe, 0x05, 0x00, 0x00,
}
