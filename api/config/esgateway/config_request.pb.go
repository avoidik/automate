// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/esgateway/config_request.proto

package esgateway

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
	V1                   *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0}
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
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0}
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
	Mlsa                 *shared.Mlsa                      `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Tls                  *shared.TLSCredentials            `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Service              *ConfigRequest_V1_System_Service  `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Log                  *ConfigRequest_V1_System_Log      `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	Ngx                  *ConfigRequest_V1_System_Nginx    `protobuf:"bytes,5,opt,name=ngx,proto3" json:"ngx,omitempty" toml:"ngx,omitempty" mapstructure:"ngx,omitempty"`
	External             *ConfigRequest_V1_System_External `protobuf:"bytes,6,opt,name=external,proto3" json:"external,omitempty" toml:"external,omitempty" mapstructure:"external,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0, 0}
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

func (m *ConfigRequest_V1_System) GetNgx() *ConfigRequest_V1_System_Nginx {
	if m != nil {
		return m.Ngx
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetExternal() *ConfigRequest_V1_System_External {
	if m != nil {
		return m.External
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
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0, 0, 0}
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
	Level                *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0, 0, 1}
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

type ConfigRequest_V1_System_Nginx struct {
	Main                 *ConfigRequest_V1_System_Nginx_Main   `protobuf:"bytes,1,opt,name=main,proto3" json:"main,omitempty" toml:"main,omitempty" mapstructure:"main,omitempty"`
	Events               *ConfigRequest_V1_System_Nginx_Events `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty" toml:"events,omitempty" mapstructure:"events,omitempty"`
	Http                 *ConfigRequest_V1_System_Nginx_Http   `protobuf:"bytes,3,opt,name=http,proto3" json:"http,omitempty" toml:"http,omitempty" mapstructure:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Nginx) Reset()         { *m = ConfigRequest_V1_System_Nginx{} }
func (m *ConfigRequest_V1_System_Nginx) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Nginx) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Nginx) Descriptor() ([]byte, []int) {
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0, 0, 2}
}

func (m *ConfigRequest_V1_System_Nginx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx.Size(m)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Nginx proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Nginx) GetMain() *ConfigRequest_V1_System_Nginx_Main {
	if m != nil {
		return m.Main
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetEvents() *ConfigRequest_V1_System_Nginx_Events {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetHttp() *ConfigRequest_V1_System_Nginx_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type ConfigRequest_V1_System_Nginx_Main struct {
	WorkerProcesses      *wrappers.Int32Value `protobuf:"bytes,1,opt,name=worker_processes,json=workerProcesses,proto3" json:"worker_processes,omitempty" toml:"worker_processes,omitempty" mapstructure:"worker_processes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Nginx_Main) Reset()         { *m = ConfigRequest_V1_System_Nginx_Main{} }
func (m *ConfigRequest_V1_System_Nginx_Main) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Nginx_Main) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Nginx_Main) Descriptor() ([]byte, []int) {
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0, 0, 2, 0}
}

func (m *ConfigRequest_V1_System_Nginx_Main) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Nginx_Main) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Nginx_Main) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Nginx_Main) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.Size(m)
}
func (m *ConfigRequest_V1_System_Nginx_Main) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Nginx_Main) GetWorkerProcesses() *wrappers.Int32Value {
	if m != nil {
		return m.WorkerProcesses
	}
	return nil
}

type ConfigRequest_V1_System_Nginx_Events struct {
	WorkerConnections    *wrappers.Int32Value `protobuf:"bytes,1,opt,name=worker_connections,json=workerConnections,proto3" json:"worker_connections,omitempty" toml:"worker_connections,omitempty" mapstructure:"worker_connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Nginx_Events) Reset()         { *m = ConfigRequest_V1_System_Nginx_Events{} }
func (m *ConfigRequest_V1_System_Nginx_Events) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Nginx_Events) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Nginx_Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0, 0, 2, 1}
}

func (m *ConfigRequest_V1_System_Nginx_Events) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Nginx_Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Nginx_Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Nginx_Events) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.Size(m)
}
func (m *ConfigRequest_V1_System_Nginx_Events) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Nginx_Events) GetWorkerConnections() *wrappers.Int32Value {
	if m != nil {
		return m.WorkerConnections
	}
	return nil
}

type ConfigRequest_V1_System_Nginx_Http struct {
	ClientMaxBodySize   *wrappers.StringValue `protobuf:"bytes,1,opt,name=client_max_body_size,json=clientMaxBodySize,proto3" json:"client_max_body_size,omitempty" toml:"client_max_body_size,omitempty" mapstructure:"client_max_body_size,omitempty"`
	ProxyConnectTimeout *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=proxy_connect_timeout,json=proxyConnectTimeout,proto3" json:"proxy_connect_timeout,omitempty" toml:"proxy_connect_timeout,omitempty" mapstructure:"proxy_connect_timeout,omitempty"`
	KeepaliveTimeout    *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=keepalive_timeout,json=keepaliveTimeout,proto3" json:"keepalive_timeout,omitempty" toml:"keepalive_timeout,omitempty" mapstructure:"keepalive_timeout,omitempty"`
	Gzip                *wrappers.StringValue `protobuf:"bytes,4,opt,name=gzip,proto3" json:"gzip,omitempty" toml:"gzip,omitempty" mapstructure:"gzip,omitempty"`
	GzipHttpVersion     *wrappers.StringValue `protobuf:"bytes,5,opt,name=gzip_http_version,json=gzipHttpVersion,proto3" json:"gzip_http_version,omitempty" toml:"gzip_http_version,omitempty" mapstructure:"gzip_http_version,omitempty"`
	// StringValue for consitency with other nginx configs
	GzipCompLevel             *wrappers.StringValue `protobuf:"bytes,6,opt,name=gzip_comp_level,json=gzipCompLevel,proto3" json:"gzip_comp_level,omitempty" toml:"gzip_comp_level,omitempty" mapstructure:"gzip_comp_level,omitempty"`
	GzipProxied               *wrappers.StringValue `protobuf:"bytes,7,opt,name=gzip_proxied,json=gzipProxied,proto3" json:"gzip_proxied,omitempty" toml:"gzip_proxied,omitempty" mapstructure:"gzip_proxied,omitempty"`
	GzipTypes                 *wrappers.StringValue `protobuf:"bytes,8,opt,name=gzip_types,json=gzipTypes,proto3" json:"gzip_types,omitempty" toml:"gzip_types,omitempty" mapstructure:"gzip_types,omitempty"`
	Sendfile                  *wrappers.StringValue `protobuf:"bytes,9,opt,name=sendfile,proto3" json:"sendfile,omitempty" toml:"sendfile,omitempty" mapstructure:"sendfile,omitempty"`
	TcpNodelay                *wrappers.StringValue `protobuf:"bytes,10,opt,name=tcp_nodelay,json=tcpNodelay,proto3" json:"tcp_nodelay,omitempty" toml:"tcp_nodelay,omitempty" mapstructure:"tcp_nodelay,omitempty"`
	TcpNopush                 *wrappers.StringValue `protobuf:"bytes,11,opt,name=tcp_nopush,json=tcpNopush,proto3" json:"tcp_nopush,omitempty" toml:"tcp_nopush,omitempty" mapstructure:"tcp_nopush,omitempty"`
	SslCiphers                *wrappers.StringValue `protobuf:"bytes,12,opt,name=ssl_ciphers,json=sslCiphers,proto3" json:"ssl_ciphers,omitempty" toml:"ssl_ciphers,omitempty" mapstructure:"ssl_ciphers,omitempty"`
	SslProtocols              *wrappers.StringValue `protobuf:"bytes,13,opt,name=ssl_protocols,json=sslProtocols,proto3" json:"ssl_protocols,omitempty" toml:"ssl_protocols,omitempty" mapstructure:"ssl_protocols,omitempty"`
	SslVerifyDepth            *wrappers.Int32Value  `protobuf:"bytes,15,opt,name=ssl_verify_depth,json=sslVerifyDepth,proto3" json:"ssl_verify_depth,omitempty" toml:"ssl_verify_depth,omitempty" mapstructure:"ssl_verify_depth,omitempty"`
	ServerNamesHashBucketSize *wrappers.Int32Value  `protobuf:"bytes,16,opt,name=server_names_hash_bucket_size,json=serverNamesHashBucketSize,proto3" json:"server_names_hash_bucket_size,omitempty" toml:"server_names_hash_bucket_size,omitempty" mapstructure:"server_names_hash_bucket_size,omitempty"`
	ClientBodyBufferSize      *wrappers.StringValue `protobuf:"bytes,17,opt,name=client_body_buffer_size,json=clientBodyBufferSize,proto3" json:"client_body_buffer_size,omitempty" toml:"client_body_buffer_size,omitempty" mapstructure:"client_body_buffer_size,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized          []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache             int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Nginx_Http) Reset()         { *m = ConfigRequest_V1_System_Nginx_Http{} }
func (m *ConfigRequest_V1_System_Nginx_Http) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Nginx_Http) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Nginx_Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0, 0, 2, 2}
}

func (m *ConfigRequest_V1_System_Nginx_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Nginx_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Nginx_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Nginx_Http) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.Size(m)
}
func (m *ConfigRequest_V1_System_Nginx_Http) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Nginx_Http) GetClientMaxBodySize() *wrappers.StringValue {
	if m != nil {
		return m.ClientMaxBodySize
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetProxyConnectTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.ProxyConnectTimeout
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetKeepaliveTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.KeepaliveTimeout
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzip() *wrappers.StringValue {
	if m != nil {
		return m.Gzip
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzipHttpVersion() *wrappers.StringValue {
	if m != nil {
		return m.GzipHttpVersion
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzipCompLevel() *wrappers.StringValue {
	if m != nil {
		return m.GzipCompLevel
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzipProxied() *wrappers.StringValue {
	if m != nil {
		return m.GzipProxied
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzipTypes() *wrappers.StringValue {
	if m != nil {
		return m.GzipTypes
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetSendfile() *wrappers.StringValue {
	if m != nil {
		return m.Sendfile
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetTcpNodelay() *wrappers.StringValue {
	if m != nil {
		return m.TcpNodelay
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetTcpNopush() *wrappers.StringValue {
	if m != nil {
		return m.TcpNopush
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetSslCiphers() *wrappers.StringValue {
	if m != nil {
		return m.SslCiphers
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetSslProtocols() *wrappers.StringValue {
	if m != nil {
		return m.SslProtocols
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetSslVerifyDepth() *wrappers.Int32Value {
	if m != nil {
		return m.SslVerifyDepth
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetServerNamesHashBucketSize() *wrappers.Int32Value {
	if m != nil {
		return m.ServerNamesHashBucketSize
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetClientBodyBufferSize() *wrappers.StringValue {
	if m != nil {
		return m.ClientBodyBufferSize
	}
	return nil
}

type ConfigRequest_V1_System_External struct {
	Enable               *wrappers.BoolValue     `protobuf:"bytes,1,opt,name=enable,proto3" json:"enable,omitempty" toml:"enable,omitempty" mapstructure:"enable,omitempty"`
	SslUpstream          *wrappers.BoolValue     `protobuf:"bytes,2,opt,name=ssl_upstream,json=sslUpstream,proto3" json:"ssl_upstream,omitempty" toml:"ssl_upstream,omitempty" mapstructure:"ssl_upstream,omitempty"`
	Endpoints            []*wrappers.StringValue `protobuf:"bytes,3,rep,name=endpoints,proto3" json:"endpoints,omitempty" toml:"endpoints,omitempty" mapstructure:"endpoints,omitempty"`
	BasicAuthCredentials *wrappers.StringValue   `protobuf:"bytes,4,opt,name=basic_auth_credentials,json=basicAuthCredentials,proto3" json:"basic_auth_credentials,omitempty" toml:"basic_auth_credentials,omitempty" mapstructure:"basic_auth_credentials,omitempty"`
	RootCert             *wrappers.StringValue   `protobuf:"bytes,5,opt,name=root_cert,json=rootCert,proto3" json:"root_cert,omitempty" toml:"root_cert,omitempty" mapstructure:"root_cert,omitempty"`
	ServerName           *wrappers.StringValue   `protobuf:"bytes,6,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty" toml:"server_name,omitempty" mapstructure:"server_name,omitempty"`
	RootCertFile         *wrappers.StringValue   `protobuf:"bytes,7,opt,name=root_cert_file,json=rootCertFile,proto3" json:"root_cert_file,omitempty" toml:"root_cert_file,omitempty" mapstructure:"root_cert_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                   `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_External) Reset()         { *m = ConfigRequest_V1_System_External{} }
func (m *ConfigRequest_V1_System_External) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_External) ProtoMessage()    {}
func (*ConfigRequest_V1_System_External) Descriptor() ([]byte, []int) {
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0, 0, 3}
}

func (m *ConfigRequest_V1_System_External) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_External.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_External) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_External.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_External) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_External.Merge(m, src)
}
func (m *ConfigRequest_V1_System_External) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_External.Size(m)
}
func (m *ConfigRequest_V1_System_External) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_External.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_External proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_External) GetEnable() *wrappers.BoolValue {
	if m != nil {
		return m.Enable
	}
	return nil
}

func (m *ConfigRequest_V1_System_External) GetSslUpstream() *wrappers.BoolValue {
	if m != nil {
		return m.SslUpstream
	}
	return nil
}

func (m *ConfigRequest_V1_System_External) GetEndpoints() []*wrappers.StringValue {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ConfigRequest_V1_System_External) GetBasicAuthCredentials() *wrappers.StringValue {
	if m != nil {
		return m.BasicAuthCredentials
	}
	return nil
}

func (m *ConfigRequest_V1_System_External) GetRootCert() *wrappers.StringValue {
	if m != nil {
		return m.RootCert
	}
	return nil
}

func (m *ConfigRequest_V1_System_External) GetServerName() *wrappers.StringValue {
	if m != nil {
		return m.ServerName
	}
	return nil
}

func (m *ConfigRequest_V1_System_External) GetRootCertFile() *wrappers.StringValue {
	if m != nil {
		return m.RootCertFile
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
	return fileDescriptor_70f0b1ef0a70f9f5, []int{0, 0, 1}
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
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.infra.esgateway.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_System_Nginx)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1.System.Nginx")
	proto.RegisterType((*ConfigRequest_V1_System_Nginx_Main)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1.System.Nginx.Main")
	proto.RegisterType((*ConfigRequest_V1_System_Nginx_Events)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1.System.Nginx.Events")
	proto.RegisterType((*ConfigRequest_V1_System_Nginx_Http)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1.System.Nginx.Http")
	proto.RegisterType((*ConfigRequest_V1_System_External)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1.System.External")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.infra.esgateway.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("api/config/esgateway/config_request.proto", fileDescriptor_70f0b1ef0a70f9f5)
}

var fileDescriptor_70f0b1ef0a70f9f5 = []byte{
	// 1152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x97, 0x4f, 0x6f, 0x14, 0x37,
	0x14, 0xc0, 0x45, 0x76, 0x59, 0x36, 0x2f, 0x09, 0x49, 0xcc, 0xbf, 0x61, 0x28, 0x08, 0xf5, 0xd4,
	0x56, 0xca, 0x2c, 0x09, 0x55, 0x5b, 0x28, 0x14, 0xb1, 0x0b, 0x34, 0xa0, 0x24, 0xa0, 0x49, 0x88,
	0xaa, 0xa2, 0x6a, 0xe4, 0x9d, 0x7d, 0x3b, 0x63, 0xe1, 0xb5, 0xa7, 0xb6, 0x77, 0xc9, 0x72, 0xee,
	0xa1, 0x5f, 0xa2, 0x52, 0x0f, 0xbd, 0xf7, 0x3b, 0x70, 0xed, 0x27, 0xe9, 0x99, 0x2f, 0x50, 0xd9,
	0xe3, 0x9d, 0x50, 0x81, 0x92, 0x21, 0x9c, 0xa2, 0xec, 0xbc, 0xdf, 0xcf, 0xcf, 0xcf, 0x9e, 0x67,
	0x0f, 0x7c, 0x49, 0x0b, 0xd6, 0x49, 0xa5, 0x18, 0xb2, 0xac, 0x83, 0x3a, 0xa3, 0x06, 0x5f, 0xd1,
	0xa9, 0xff, 0x21, 0x51, 0xf8, 0xeb, 0x18, 0xb5, 0x89, 0x0a, 0x25, 0x8d, 0x24, 0x57, 0xd3, 0x1c,
	0x87, 0x11, 0x1d, 0x1b, 0x39, 0xa2, 0x06, 0x23, 0x26, 0x86, 0x8a, 0x46, 0x15, 0x13, 0x5e, 0x7b,
	0xc7, 0xa4, 0x73, 0xaa, 0x70, 0xd0, 0xc9, 0xb8, 0xec, 0x53, 0x5e, 0xe2, 0xe1, 0x95, 0xf7, 0x9f,
	0x1b, 0xae, 0xfd, 0xc3, 0x27, 0xa9, 0x1c, 0x15, 0x52, 0xa0, 0x30, 0xba, 0x33, 0x1b, 0x61, 0x2d,
	0x53, 0x45, 0xda, 0x71, 0xcf, 0xd3, 0xb5, 0x0c, 0xc5, 0x1a, 0xdd, 0x58, 0xf3, 0xbc, 0x55, 0xd1,
	0x0d, 0xfb, 0x4f, 0x87, 0x0a, 0x21, 0x0d, 0x35, 0x4c, 0x8a, 0x99, 0xeb, 0x5a, 0x26, 0x65, 0xc6,
	0xb1, 0x24, 0xfb, 0xe3, 0x61, 0xe7, 0x95, 0xa2, 0x45, 0x81, 0xca, 0x3f, 0xff, 0xfc, 0x8f, 0x4b,
	0xb0, 0xd4, 0x73, 0x9e, 0xb8, 0x9c, 0x1f, 0xb9, 0x07, 0x73, 0x93, 0xf5, 0xe0, 0xd4, 0xf5, 0x53,
	0x5f, 0x2c, 0x6c, 0x74, 0xa2, 0x23, 0xa7, 0x19, 0xfd, 0x8f, 0x8c, 0xf6, 0xd7, 0xe3, 0xb9, 0xc9,
	0x7a, 0xf8, 0xf7, 0x45, 0x98, 0xdb, 0x5f, 0x27, 0x9b, 0xd0, 0xd0, 0x53, 0xed, 0x45, 0xdf, 0x7c,
	0xa4, 0x28, 0xda, 0x9d, 0x6a, 0x83, 0xa3, 0xd8, 0x2a, 0xc8, 0x63, 0x68, 0xe8, 0x49, 0x1a, 0xcc,
	0x39, 0xd3, 0xb7, 0x1f, 0x6d, 0x42, 0x35, 0x61, 0x29, 0xc6, 0xd6, 0x11, 0xfe, 0x7e, 0x01, 0x5a,
	0xa5, 0x9a, 0x7c, 0x0d, 0xcd, 0x11, 0xd7, 0xd4, 0x27, 0x78, 0xfd, 0x83, 0xda, 0xb2, 0xc2, 0xd1,
	0x36, 0xd7, 0x34, 0x76, 0xd1, 0xe4, 0x0e, 0x34, 0x0c, 0xd7, 0x3e, 0x97, 0xaf, 0x8e, 0x82, 0xf6,
	0xb6, 0x76, 0x7b, 0x0a, 0x07, 0x28, 0x0c, 0xa3, 0x5c, 0xc7, 0x16, 0x23, 0x3f, 0xc1, 0x19, 0x5d,
	0xa6, 0x13, 0x34, 0x9c, 0xe1, 0x87, 0x93, 0xd5, 0xa5, 0x9a, 0xd4, 0x4c, 0x47, 0xb6, 0xa0, 0xc1,
	0x65, 0x16, 0x34, 0x9d, 0xf5, 0xf6, 0x09, 0xad, 0x5b, 0x32, 0x8b, 0xad, 0x86, 0xec, 0x40, 0x43,
	0x64, 0x07, 0xc1, 0x69, 0x67, 0xbb, 0x73, 0x42, 0xdb, 0x4e, 0xc6, 0xc4, 0x41, 0x6c, 0x45, 0xe4,
	0x05, 0xb4, 0xf1, 0xc0, 0xa0, 0x12, 0x94, 0x07, 0x2d, 0x27, 0xbd, 0x77, 0x42, 0xe9, 0x43, 0xaf,
	0x89, 0x2b, 0x61, 0xf8, 0xdb, 0x29, 0x38, 0xe3, 0xeb, 0x41, 0x6e, 0x40, 0x33, 0x97, 0xda, 0xf8,
	0x45, 0xfd, 0x2c, 0x2a, 0x77, 0x7f, 0x34, 0xdb, 0xfd, 0xd1, 0xae, 0x51, 0x4c, 0x64, 0xfb, 0x94,
	0x8f, 0x31, 0x76, 0x91, 0xe4, 0x47, 0x68, 0x16, 0x52, 0x19, 0xbf, 0xa2, 0x57, 0xde, 0x23, 0x1e,
	0x0b, 0x73, 0x73, 0xc3, 0x01, 0xdd, 0x8b, 0x6f, 0xde, 0x06, 0xa4, 0x5a, 0xc1, 0x95, 0x3f, 0x9f,
	0x86, 0xcd, 0xdc, 0x98, 0x22, 0x76, 0x82, 0xf0, 0x16, 0x34, 0xb6, 0x64, 0x46, 0x36, 0xe0, 0x34,
	0xc7, 0x09, 0xf2, 0x5a, 0x29, 0x94, 0xa1, 0xe1, 0x5f, 0x8b, 0x70, 0xda, 0x55, 0x8b, 0x3c, 0x87,
	0xe6, 0x88, 0x32, 0xe1, 0xe1, 0xfb, 0x9f, 0x52, 0xf9, 0x68, 0x9b, 0x32, 0x11, 0x3b, 0x1d, 0x79,
	0x01, 0x2d, 0x9c, 0xd8, 0x7e, 0xe2, 0xa7, 0xd9, 0xfb, 0x24, 0xf1, 0x43, 0xa7, 0x8a, 0xbd, 0xd2,
	0xe6, 0x6c, 0xcb, 0xe0, 0x77, 0xf4, 0xa7, 0xe5, 0xbc, 0xe9, 0xea, 0x69, 0x75, 0xe1, 0x0e, 0x34,
	0xed, 0x0c, 0xc8, 0x23, 0x58, 0x79, 0x25, 0xd5, 0x4b, 0x54, 0x49, 0xa1, 0x64, 0x8a, 0x5a, 0xe3,
	0xac, 0xa9, 0x1c, 0xb5, 0x58, 0xf1, 0x72, 0x09, 0x3d, 0x9b, 0x31, 0xe1, 0x1e, 0xb4, 0xca, 0xc4,
	0xc9, 0x13, 0x20, 0xde, 0x98, 0x4a, 0x21, 0x30, 0x75, 0xfd, 0xb2, 0x8e, 0x73, 0xb5, 0xc4, 0x7a,
	0x87, 0x54, 0xf8, 0x6f, 0x1b, 0x9a, 0x36, 0x69, 0xb2, 0x0d, 0xe7, 0x53, 0xce, 0x50, 0x98, 0x64,
	0x44, 0x0f, 0x92, 0xbe, 0x1c, 0x4c, 0x13, 0xcd, 0x5e, 0x63, 0xad, 0x6d, 0xb0, 0x5a, 0x92, 0xdb,
	0xf4, 0xa0, 0x2b, 0x07, 0xd3, 0x5d, 0xf6, 0x1a, 0xc9, 0x53, 0xb8, 0x50, 0x28, 0x79, 0x30, 0x9d,
	0xa5, 0x98, 0x18, 0x36, 0x42, 0x39, 0xae, 0xb3, 0x4f, 0xe3, 0x73, 0x8e, 0xf4, 0x59, 0xee, 0x95,
	0x1c, 0xd9, 0x84, 0xd5, 0x97, 0x88, 0x05, 0xe5, 0x6c, 0x82, 0x95, 0xac, 0x71, 0xbc, 0x6c, 0xa5,
	0xa2, 0x66, 0xa6, 0x1b, 0xd0, 0xcc, 0x5e, 0xb3, 0xc2, 0xf7, 0x9a, 0x63, 0xde, 0x31, 0x1b, 0x69,
	0xc7, 0xb6, 0x7f, 0x13, 0xbb, 0xae, 0xc9, 0x04, 0x95, 0x66, 0x52, 0xf8, 0xe6, 0x72, 0x34, 0xbe,
	0x6c, 0x31, 0x5b, 0xde, 0xfd, 0x12, 0x22, 0x0f, 0xc0, 0xfd, 0x94, 0xd8, 0x13, 0x32, 0x29, 0xdf,
	0xb3, 0x56, 0x0d, 0xcf, 0x92, 0x85, 0x7a, 0x72, 0x54, 0x6c, 0x59, 0x84, 0xdc, 0x83, 0x45, 0x67,
	0xb1, 0x75, 0x62, 0x38, 0x08, 0xce, 0xd4, 0x50, 0x2c, 0x58, 0xe2, 0x59, 0x09, 0x90, 0xef, 0x01,
	0x9c, 0xc0, 0x4c, 0x0b, 0xd4, 0x41, 0xbb, 0x06, 0x3e, 0x6f, 0xe3, 0xf7, 0x6c, 0x38, 0xf9, 0x0e,
	0xda, 0x1a, 0xc5, 0x60, 0xc8, 0x38, 0x06, 0xf3, 0x35, 0xd0, 0x2a, 0x9a, 0xdc, 0x85, 0x05, 0x93,
	0x16, 0x89, 0x90, 0x03, 0xe4, 0x74, 0x1a, 0x40, 0x0d, 0x18, 0x4c, 0x5a, 0xec, 0x94, 0xf1, 0x36,
	0xeb, 0x12, 0x2f, 0xc6, 0x3a, 0x0f, 0x16, 0xea, 0x64, 0xed, 0x68, 0x1b, 0x6e, 0xc7, 0xd6, 0x9a,
	0x27, 0x29, 0x2b, 0x72, 0x54, 0x3a, 0x58, 0xac, 0x33, 0xb6, 0xd6, 0xbc, 0x57, 0xc6, 0x93, 0xfb,
	0xb0, 0x64, 0xf1, 0xf2, 0x02, 0x23, 0xb9, 0x0e, 0x96, 0x6a, 0x08, 0x16, 0xb5, 0xe6, 0xcf, 0x66,
	0x04, 0x79, 0x08, 0x2b, 0x56, 0x31, 0x41, 0xc5, 0x86, 0xd3, 0x64, 0x80, 0x85, 0xc9, 0x83, 0xe5,
	0xe3, 0x37, 0xf0, 0x59, 0xad, 0xf9, 0xbe, 0x63, 0x1e, 0x58, 0x84, 0xfc, 0x02, 0x57, 0x6d, 0x07,
	0x47, 0x95, 0x08, 0x3a, 0x42, 0x9d, 0xe4, 0x54, 0xe7, 0x49, 0x7f, 0x9c, 0xbe, 0x44, 0x53, 0xbe,
	0xb1, 0x2b, 0xc7, 0x3b, 0x2f, 0x97, 0x86, 0x1d, 0x2b, 0xd8, 0xa4, 0x3a, 0xef, 0x3a, 0xdc, 0xbd,
	0xb8, 0xbb, 0x70, 0xc9, 0xf7, 0x01, 0xd7, 0x03, 0xfa, 0xe3, 0xe1, 0x10, 0x55, 0x29, 0x5e, 0xad,
	0x31, 0x65, 0xdf, 0x44, 0x6c, 0x1f, 0xe8, 0x3a, 0xd4, 0x4a, 0x9f, 0x34, 0xdb, 0x67, 0x57, 0x96,
	0xc3, 0x7f, 0x1a, 0xd0, 0x9e, 0x9d, 0x7f, 0x64, 0x03, 0x5a, 0x28, 0x68, 0x9f, 0xcf, 0x3a, 0x4c,
	0xf8, 0x9e, 0xb6, 0x2b, 0x25, 0x2f, 0xa5, 0x3e, 0x92, 0xdc, 0x05, 0x5b, 0xd1, 0x64, 0x5c, 0x68,
	0xa3, 0x90, 0x8e, 0x7c, 0x2f, 0x39, 0x8a, 0xb4, 0x6b, 0xfe, 0xdc, 0x87, 0x93, 0xdb, 0x30, 0x8f,
	0x62, 0x50, 0x48, 0x66, 0x0f, 0x92, 0xc6, 0xf5, 0xc6, 0xf1, 0xdb, 0xa7, 0x0a, 0x27, 0x31, 0x5c,
	0xec, 0x53, 0xcd, 0xd2, 0x84, 0x8e, 0x4d, 0x9e, 0xa4, 0x87, 0x17, 0xa3, 0x5a, 0x6d, 0xe4, 0xbc,
	0x63, 0xef, 0x8f, 0x4d, 0xfe, 0xce, 0x95, 0x8a, 0xdc, 0x82, 0x79, 0x25, 0xa5, 0x49, 0x52, 0x54,
	0xa6, 0x56, 0x3b, 0x69, 0xdb, 0xf0, 0x1e, 0x2a, 0xe3, 0x76, 0xf3, 0xe1, 0x26, 0xa8, 0xd5, 0x43,
	0xe0, 0x70, 0xcd, 0x49, 0x17, 0xce, 0x56, 0x23, 0x27, 0xee, 0x45, 0xae, 0xd3, 0x42, 0x16, 0x67,
	0xc3, 0x3f, 0x62, 0x1c, 0xc3, 0xf9, 0xea, 0xd6, 0x72, 0xfb, 0xf2, 0x9b, 0xb7, 0xc1, 0x05, 0x38,
	0x57, 0xdd, 0xf5, 0x51, 0xaf, 0xf9, 0xe3, 0xb2, 0xdb, 0xf9, 0x79, 0x2d, 0x63, 0x26, 0x1f, 0xf7,
	0xa3, 0x54, 0x8e, 0x3a, 0xf6, 0x68, 0xad, 0x3e, 0x09, 0x3a, 0x1f, 0xfa, 0x5a, 0xe9, 0xb7, 0xdc,
	0xd0, 0x37, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x44, 0x86, 0xd3, 0x61, 0xcc, 0x0c, 0x00, 0x00,
}
