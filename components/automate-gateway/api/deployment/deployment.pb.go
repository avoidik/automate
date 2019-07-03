// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/deployment/deployment.proto

package deployment

import (
	context "context"
	fmt "fmt"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Version message
//
// The manifest version constructed with:
// * build_timestamp
type Version struct {
	BuildTimestamp       string   `protobuf:"bytes,1,opt,name=build_timestamp,json=buildTimestamp,proto3" json:"build_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4362896e7bedc82, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetBuildTimestamp() string {
	if m != nil {
		return m.BuildTimestamp
	}
	return ""
}

type ServiceVersionsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceVersionsRequest) Reset()         { *m = ServiceVersionsRequest{} }
func (m *ServiceVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceVersionsRequest) ProtoMessage()    {}
func (*ServiceVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4362896e7bedc82, []int{1}
}

func (m *ServiceVersionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceVersionsRequest.Unmarshal(m, b)
}
func (m *ServiceVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceVersionsRequest.Marshal(b, m, deterministic)
}
func (m *ServiceVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceVersionsRequest.Merge(m, src)
}
func (m *ServiceVersionsRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceVersionsRequest.Size(m)
}
func (m *ServiceVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceVersionsRequest proto.InternalMessageInfo

type ServiceVersionsResponse struct {
	Services             []*ServiceVersion `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceVersionsResponse) Reset()         { *m = ServiceVersionsResponse{} }
func (m *ServiceVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceVersionsResponse) ProtoMessage()    {}
func (*ServiceVersionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4362896e7bedc82, []int{2}
}

func (m *ServiceVersionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceVersionsResponse.Unmarshal(m, b)
}
func (m *ServiceVersionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceVersionsResponse.Marshal(b, m, deterministic)
}
func (m *ServiceVersionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceVersionsResponse.Merge(m, src)
}
func (m *ServiceVersionsResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceVersionsResponse.Size(m)
}
func (m *ServiceVersionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceVersionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceVersionsResponse proto.InternalMessageInfo

func (m *ServiceVersionsResponse) GetServices() []*ServiceVersion {
	if m != nil {
		return m.Services
	}
	return nil
}

type ServiceVersion struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Origin               string   `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Release              string   `protobuf:"bytes,4,opt,name=release,proto3" json:"release,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceVersion) Reset()         { *m = ServiceVersion{} }
func (m *ServiceVersion) String() string { return proto.CompactTextString(m) }
func (*ServiceVersion) ProtoMessage()    {}
func (*ServiceVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4362896e7bedc82, []int{3}
}

func (m *ServiceVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceVersion.Unmarshal(m, b)
}
func (m *ServiceVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceVersion.Marshal(b, m, deterministic)
}
func (m *ServiceVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceVersion.Merge(m, src)
}
func (m *ServiceVersion) XXX_Size() int {
	return xxx_messageInfo_ServiceVersion.Size(m)
}
func (m *ServiceVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceVersion proto.InternalMessageInfo

func (m *ServiceVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceVersion) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *ServiceVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceVersion) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func init() {
	proto.RegisterType((*Version)(nil), "chef.automate.api.deployment.Version")
	proto.RegisterType((*ServiceVersionsRequest)(nil), "chef.automate.api.deployment.ServiceVersionsRequest")
	proto.RegisterType((*ServiceVersionsResponse)(nil), "chef.automate.api.deployment.ServiceVersionsResponse")
	proto.RegisterType((*ServiceVersion)(nil), "chef.automate.api.deployment.ServiceVersion")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/deployment/deployment.proto", fileDescriptor_e4362896e7bedc82)
}

var fileDescriptor_e4362896e7bedc82 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x28, 0x2d, 0x83, 0xd4, 0x4a, 0x2b, 0x64, 0x16, 0x37, 0x42, 0x95, 0x25, 0x44,
	0x0f, 0x64, 0x2d, 0x05, 0xb8, 0x84, 0x0b, 0x02, 0x2a, 0x38, 0x07, 0xc4, 0x81, 0x4b, 0xb5, 0x71,
	0x26, 0xee, 0x4a, 0xde, 0x0f, 0xbc, 0xeb, 0x22, 0x5f, 0x23, 0xb8, 0xf4, 0xca, 0x6f, 0xc9, 0x8d,
	0x7f, 0xd1, 0xbf, 0xc0, 0x0f, 0x41, 0xb1, 0xd7, 0xa1, 0xa1, 0x25, 0x22, 0x12, 0xb7, 0x9d, 0x79,
	0xf3, 0xde, 0xce, 0xcb, 0xcb, 0x1a, 0x5e, 0xa4, 0x5a, 0x1a, 0xad, 0x50, 0x39, 0x9b, 0xf0, 0xd2,
	0x69, 0xc9, 0x1d, 0x0e, 0x33, 0xee, 0xf0, 0x0b, 0xaf, 0x12, 0x6e, 0x44, 0x32, 0x43, 0x93, 0xeb,
	0x4a, 0xa2, 0x72, 0xd7, 0x8e, 0xcc, 0x14, 0xda, 0x69, 0x32, 0x48, 0xcf, 0x71, 0xce, 0x5a, 0x1a,
	0xe3, 0x46, 0xb0, 0xdf, 0x33, 0xd1, 0x20, 0xd3, 0x3a, 0xcb, 0xb1, 0x56, 0xe1, 0x4a, 0x69, 0xc7,
	0x9d, 0xd0, 0xca, 0x36, 0xdc, 0xe8, 0xc8, 0xa3, 0x75, 0x35, 0x2d, 0xe7, 0x09, 0x4a, 0xe3, 0x2a,
	0x0f, 0xbe, 0xbc, 0x75, 0xab, 0xc2, 0xa4, 0xcd, 0x78, 0x3a, 0xcc, 0x50, 0x0d, 0x8d, 0xce, 0x45,
	0x5a, 0xfd, 0x45, 0x7e, 0x17, 0x05, 0xc1, 0xe5, 0x4d, 0x85, 0x78, 0x04, 0x7b, 0x1f, 0xb1, 0xb0,
	0x42, 0x2b, 0xf2, 0x18, 0x0e, 0xa7, 0xa5, 0xc8, 0x67, 0x67, 0x4e, 0x48, 0xb4, 0x8e, 0x4b, 0x43,
	0x83, 0xe3, 0xe0, 0xe4, 0xce, 0xe4, 0xa0, 0x6e, 0x7f, 0x68, 0xbb, 0x31, 0x85, 0xf0, 0x3d, 0x16,
	0x17, 0x22, 0x45, 0x4f, 0xb5, 0x13, 0xfc, 0x5c, 0xa2, 0x75, 0x71, 0x0a, 0xf7, 0x6f, 0x20, 0xd6,
	0x68, 0x65, 0x91, 0xbc, 0x83, 0x7d, 0xdb, 0x40, 0x96, 0x06, 0xc7, 0xdd, 0x93, 0xbb, 0xa3, 0x27,
	0x6c, 0xdb, 0x0f, 0xcb, 0x36, 0x85, 0x26, 0x6b, 0x76, 0x6c, 0xe0, 0x60, 0x13, 0x23, 0x04, 0x7a,
	0x8a, 0x4b, 0xf4, 0xeb, 0xd6, 0x67, 0x12, 0x42, 0x5f, 0x17, 0x22, 0x13, 0x8a, 0x76, 0xea, 0xae,
	0xaf, 0x08, 0x85, 0xbd, 0x8b, 0x86, 0x46, 0xbb, 0x35, 0xd0, 0x96, 0x2b, 0xa4, 0xc0, 0x1c, 0xb9,
	0x45, 0xda, 0x6b, 0x10, 0x5f, 0x8e, 0xae, 0xba, 0x00, 0x6f, 0xd6, 0x9b, 0x91, 0x1f, 0x01, 0xc0,
	0x5b, 0x74, 0xed, 0xed, 0x21, 0x6b, 0x42, 0x66, 0x6d, 0xc8, 0xec, 0x74, 0x15, 0x72, 0xf4, 0x68,
	0xbb, 0x3f, 0x4f, 0x8f, 0xf5, 0x62, 0x49, 0x43, 0xb8, 0xe7, 0xed, 0x9d, 0x09, 0x35, 0xd7, 0x63,
	0xbf, 0xd2, 0x62, 0x49, 0xfb, 0xa4, 0x57, 0x20, 0x9f, 0x5d, 0x2e, 0x29, 0x85, 0xd0, 0x56, 0xd6,
	0xa1, 0x1c, 0xfb, 0xd1, 0x76, 0xea, 0x72, 0x49, 0x8f, 0xc8, 0x83, 0x4d, 0xcc, 0x8b, 0x8f, 0x33,
	0x74, 0x8b, 0xab, 0x9f, 0xdf, 0x3b, 0x40, 0xf6, 0x93, 0xd6, 0xe7, 0xb7, 0x0e, 0x1c, 0xfe, 0x91,
	0x12, 0x79, 0xb6, 0x4b, 0x16, 0x6d, 0xdc, 0xd1, 0xf3, 0x1d, 0x59, 0xcd, 0x5f, 0x21, 0xfe, 0x1a,
	0xfc, 0x07, 0xcb, 0x03, 0x12, 0xdd, 0x6e, 0x39, 0x17, 0xb6, 0xf1, 0xfc, 0x90, 0x0c, 0xae, 0xbf,
	0xeb, 0xf6, 0x36, 0x2f, 0x61, 0x5f, 0x9d, 0x7e, 0x7a, 0x9d, 0x09, 0x77, 0x5e, 0x4e, 0x59, 0xaa,
	0x65, 0xb2, 0x72, 0xb2, 0x7e, 0x43, 0xc9, 0xbf, 0x7f, 0x2f, 0xa6, 0xfd, 0x3a, 0xf6, 0xa7, 0xbf,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x38, 0x9d, 0x71, 0x64, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeploymentClient is the client API for Deployment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeploymentClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error)
	ServiceVersions(ctx context.Context, in *ServiceVersionsRequest, opts ...grpc.CallOption) (*ServiceVersionsResponse, error)
}

type deploymentClient struct {
	cc *grpc.ClientConn
}

func NewDeploymentClient(cc *grpc.ClientConn) DeploymentClient {
	return &deploymentClient{cc}
}

func (c *deploymentClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/chef.automate.api.deployment.Deployment/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentClient) ServiceVersions(ctx context.Context, in *ServiceVersionsRequest, opts ...grpc.CallOption) (*ServiceVersionsResponse, error) {
	out := new(ServiceVersionsResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.deployment.Deployment/ServiceVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentServer is the server API for Deployment service.
type DeploymentServer interface {
	GetVersion(context.Context, *empty.Empty) (*Version, error)
	ServiceVersions(context.Context, *ServiceVersionsRequest) (*ServiceVersionsResponse, error)
}

func RegisterDeploymentServer(s *grpc.Server, srv DeploymentServer) {
	s.RegisterService(&_Deployment_serviceDesc, srv)
}

func _Deployment_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.deployment.Deployment/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployment_ServiceVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).ServiceVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.deployment.Deployment/ServiceVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).ServiceVersions(ctx, req.(*ServiceVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deployment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.deployment.Deployment",
	HandlerType: (*DeploymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Deployment_GetVersion_Handler,
		},
		{
			MethodName: "ServiceVersions",
			Handler:    _Deployment_ServiceVersions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/deployment/deployment.proto",
}
