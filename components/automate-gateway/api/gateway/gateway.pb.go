// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/gateway/gateway.proto

package gateway

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
// The service version constructed with:
// * Service name
// * Built time
// * Semantic version
// * Git SHA
type Version struct {
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Built                string   `protobuf:"bytes,1,opt,name=built,proto3" json:"built,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Sha                  string   `protobuf:"bytes,4,opt,name=sha,proto3" json:"sha,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_33efad17cb9f3272, []int{0}
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

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetBuilt() string {
	if m != nil {
		return m.Built
	}
	return ""
}

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

// Health message
//
// The automate-gateway service health is constructed with:
// * Status:
//            => ok:             Everything is alright
//            => initialization: The service is in its initialization process
//            => warning:        Something might be wrong?
//            => critical:       Something is wrong!
//
// @afiune: Here we can add more health information to the response
type Health struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Health) Reset()         { *m = Health{} }
func (m *Health) String() string { return proto.CompactTextString(m) }
func (*Health) ProtoMessage()    {}
func (*Health) Descriptor() ([]byte, []int) {
	return fileDescriptor_33efad17cb9f3272, []int{1}
}

func (m *Health) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Health.Unmarshal(m, b)
}
func (m *Health) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Health.Marshal(b, m, deterministic)
}
func (m *Health) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Health.Merge(m, src)
}
func (m *Health) XXX_Size() int {
	return xxx_messageInfo_Health.Size(m)
}
func (m *Health) XXX_DiscardUnknown() {
	xxx_messageInfo_Health.DiscardUnknown(m)
}

var xxx_messageInfo_Health proto.InternalMessageInfo

func (m *Health) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Version)(nil), "chef.automate.api.Version")
	proto.RegisterType((*Health)(nil), "chef.automate.api.Health")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/gateway/gateway.proto", fileDescriptor_33efad17cb9f3272)
}

var fileDescriptor_33efad17cb9f3272 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3d, 0x6a, 0x1b, 0x41,
	0x14, 0xc7, 0x91, 0xac, 0x48, 0xf8, 0x41, 0xb0, 0x35, 0x76, 0xc4, 0x78, 0x9d, 0xc2, 0xa8, 0x4a,
	0xa3, 0x59, 0x48, 0x48, 0xa3, 0xca, 0x18, 0x82, 0x52, 0xa7, 0x48, 0x91, 0xc6, 0x8c, 0xd6, 0x4f,
	0xbb, 0x03, 0xda, 0x99, 0x65, 0xe7, 0xad, 0x8d, 0x5a, 0x95, 0x6a, 0x73, 0x85, 0x5c, 0x61, 0x8b,
	0x9c, 0x23, 0x57, 0x48, 0x99, 0x43, 0x04, 0xcd, 0x87, 0x60, 0xb1, 0x05, 0x49, 0xa5, 0xf7, 0xf9,
	0xe7, 0xf7, 0xfe, 0x9a, 0x85, 0x8f, 0x99, 0x29, 0x2b, 0xa3, 0x51, 0x93, 0x4d, 0x65, 0x43, 0xa6,
	0x94, 0x84, 0xb3, 0x5c, 0x12, 0x3e, 0xc9, 0x4d, 0x2a, 0x2b, 0x95, 0xc6, 0x38, 0xfc, 0x8a, 0xaa,
	0x36, 0x64, 0xd8, 0x38, 0x2b, 0x70, 0x25, 0xe2, 0x82, 0x90, 0x95, 0x4a, 0xde, 0xe6, 0xc6, 0xe4,
	0x6b, 0x74, 0x4b, 0x52, 0x6b, 0x43, 0x92, 0x94, 0xd1, 0xd6, 0x2f, 0x24, 0xd7, 0xa1, 0xeb, 0xb2,
	0x65, 0xb3, 0x4a, 0xb1, 0xac, 0x28, 0xa8, 0x25, 0xb7, 0x2f, 0x42, 0xd4, 0x55, 0xe6, 0xc7, 0xb3,
	0x59, 0x8e, 0x7a, 0x56, 0x99, 0xb5, 0xca, 0x36, 0x47, 0xe4, 0xff, 0x47, 0x41, 0xc9, 0xf2, 0xb9,
	0xc2, 0xf4, 0x1e, 0x46, 0x5f, 0xb1, 0xb6, 0xca, 0x68, 0xc6, 0x61, 0xf4, 0xe8, 0x43, 0xde, 0xbf,
	0xe9, 0xbd, 0x3b, 0xfd, 0x12, 0x53, 0x76, 0x09, 0xaf, 0x96, 0x8d, 0x5a, 0x13, 0xef, 0xb9, 0xba,
	0x4f, 0x18, 0x83, 0x81, 0x96, 0x25, 0xf2, 0x13, 0x57, 0x74, 0x31, 0x3b, 0x87, 0x13, 0x5b, 0x48,
	0x3e, 0x70, 0xa5, 0x7d, 0x38, 0xbd, 0x81, 0xe1, 0x67, 0x94, 0x6b, 0x2a, 0xd8, 0x04, 0x86, 0x96,
	0x24, 0x35, 0x36, 0xc8, 0x84, 0xec, 0xfd, 0x9f, 0x3e, 0x8c, 0x16, 0xde, 0x66, 0xf6, 0xb3, 0x07,
	0xb0, 0x40, 0x8a, 0x48, 0x13, 0xe1, 0xfd, 0x13, 0xd1, 0x3f, 0xf1, 0x69, 0xef, 0x5f, 0x92, 0x88,
	0x67, 0x7f, 0x84, 0x08, 0x3b, 0xd3, 0xa7, 0x6d, 0xcb, 0x27, 0x70, 0x69, 0xb1, 0x7e, 0x54, 0x19,
	0xde, 0x2b, 0xbd, 0x32, 0xf3, 0x70, 0xc8, 0xb6, 0xe5, 0x43, 0x36, 0xa8, 0x51, 0x3e, 0xec, 0x5a,
	0xce, 0x61, 0x62, 0x37, 0x96, 0xb0, 0x9c, 0x87, 0xd1, 0x38, 0xb5, 0x6b, 0xf9, 0x35, 0xbb, 0xea,
	0xf6, 0x82, 0xf8, 0x3c, 0x47, 0xda, 0xfe, 0xfa, 0xfd, 0xbd, 0xcf, 0xd8, 0xf9, 0xe1, 0x89, 0x44,
	0x97, 0x7e, 0xf4, 0xe0, 0x74, 0x81, 0x14, 0xaf, 0x3d, 0x82, 0x7e, 0xf5, 0x02, 0xba, 0x5f, 0x99,
	0x3e, 0x6c, 0x5b, 0xfe, 0x06, 0x2e, 0x3a, 0xe4, 0x85, 0x6b, 0x75, 0xc0, 0xcf, 0xe0, 0x75, 0x80,
	0xf3, 0xcd, 0x5d, 0xcb, 0x2f, 0xd8, 0xb8, 0x53, 0x3a, 0x70, 0x8e, 0xd9, 0xd9, 0x81, 0xd3, 0xb7,
	0xee, 0xee, 0xbe, 0xdd, 0xe6, 0x8a, 0x8a, 0x66, 0x29, 0x32, 0x53, 0xa6, 0x7b, 0x98, 0xc3, 0xd3,
	0x49, 0xff, 0xf1, 0xab, 0x58, 0x0e, 0xdd, 0x51, 0x1f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf1,
	0x19, 0xab, 0x8a, 0x47, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error)
	GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/chef.automate.api.Gateway/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Health, error) {
	out := new(Health)
	err := c.cc.Invoke(ctx, "/chef.automate.api.Gateway/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	GetVersion(context.Context, *empty.Empty) (*Version, error)
	GetHealth(context.Context, *empty.Empty) (*Health, error)
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.Gateway/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.Gateway/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetHealth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Gateway_GetVersion_Handler,
		},
		{
			MethodName: "GetHealth",
			Handler:    _Gateway_GetHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/gateway/gateway.proto",
}
