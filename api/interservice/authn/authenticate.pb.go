// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/authn/authenticate.proto

package authn

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AuthenticateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *AuthenticateRequest) Reset()         { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()    {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57390d1a555f5d05, []int{0}
}

func (m *AuthenticateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateRequest.Unmarshal(m, b)
}
func (m *AuthenticateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateRequest.Marshal(b, m, deterministic)
}
func (m *AuthenticateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateRequest.Merge(m, src)
}
func (m *AuthenticateRequest) XXX_Size() int {
	return xxx_messageInfo_AuthenticateRequest.Size(m)
}
func (m *AuthenticateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateRequest proto.InternalMessageInfo

type AuthenticateResponse struct {
	// This could be either "user:{local,ldap,saml}:<some-id>",
	// "tls:service:<some-id> or "token:<some-id>",
	// depending on the authentication method that was successful.
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty" toml:"subject,omitempty" mapstructure:"subject,omitempty"`
	// Only human users have teams. The teams are provided either by the external
	// IdP (in which case they're extracted from the id_token; TODO), or, for local
	// users, by teams-service (TODO).
	Teams                []string `protobuf:"bytes,2,rep,name=teams,proto3" json:"teams,omitempty" toml:"teams,omitempty" mapstructure:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *AuthenticateResponse) Reset()         { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string { return proto.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()    {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57390d1a555f5d05, []int{1}
}

func (m *AuthenticateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticateResponse.Unmarshal(m, b)
}
func (m *AuthenticateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticateResponse.Marshal(b, m, deterministic)
}
func (m *AuthenticateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticateResponse.Merge(m, src)
}
func (m *AuthenticateResponse) XXX_Size() int {
	return xxx_messageInfo_AuthenticateResponse.Size(m)
}
func (m *AuthenticateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticateResponse proto.InternalMessageInfo

func (m *AuthenticateResponse) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *AuthenticateResponse) GetTeams() []string {
	if m != nil {
		return m.Teams
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthenticateRequest)(nil), "chef.automate.domain.authn.AuthenticateRequest")
	proto.RegisterType((*AuthenticateResponse)(nil), "chef.automate.domain.authn.AuthenticateResponse")
}

func init() {
	proto.RegisterFile("api/interservice/authn/authenticate.proto", fileDescriptor_57390d1a555f5d05)
}

var fileDescriptor_57390d1a555f5d05 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4a, 0xc4, 0x40,
	0x10, 0x85, 0xc9, 0x88, 0xca, 0x34, 0x22, 0xd8, 0x33, 0x42, 0x0c, 0x2e, 0x86, 0xac, 0xc6, 0x4d,
	0xb7, 0x3f, 0x27, 0xd0, 0x85, 0x07, 0xc8, 0xd2, 0x5d, 0xa7, 0x2d, 0x93, 0x12, 0x53, 0x15, 0xd3,
	0xd5, 0x1e, 0xc0, 0xb5, 0x3b, 0x0f, 0xe0, 0xa1, 0xbc, 0x82, 0x07, 0x91, 0x24, 0x0c, 0x44, 0x18,
	0x61, 0x36, 0x0d, 0xf5, 0x78, 0xef, 0xe3, 0xf1, 0x5a, 0x5d, 0xb8, 0x16, 0x2d, 0x92, 0x40, 0x17,
	0xa0, 0x7b, 0x43, 0x0f, 0xd6, 0x45, 0xa9, 0x69, 0x78, 0x81, 0x04, 0xbd, 0x13, 0x30, 0x6d, 0xc7,
	0xc2, 0x3a, 0xf3, 0x35, 0x3c, 0x19, 0x17, 0x85, 0x9b, 0x5e, 0x7c, 0xe4, 0xc6, 0x21, 0x99, 0xc1,
	0x9e, 0x9d, 0x57, 0xcc, 0xd5, 0x0b, 0xd8, 0x9e, 0xe6, 0x88, 0x58, 0x9c, 0x20, 0x53, 0x18, 0x93,
	0xf9, 0xa9, 0x5a, 0xdc, 0x4e, 0x78, 0x05, 0xbc, 0x46, 0x08, 0x92, 0xdf, 0xab, 0xe5, 0x5f, 0x39,
	0xb4, 0x4c, 0x01, 0x74, 0xaa, 0x0e, 0x43, 0x2c, 0x9f, 0xc1, 0x4b, 0x9a, 0xac, 0x92, 0xf5, 0xbc,
	0xd8, 0x9c, 0x7a, 0xa9, 0xf6, 0x05, 0x5c, 0x13, 0xd2, 0xd9, 0x6a, 0x6f, 0x3d, 0x2f, 0xc6, 0xe3,
	0xfa, 0x2b, 0x51, 0xc7, 0x13, 0x10, 0x32, 0xe9, 0x8f, 0x44, 0x1d, 0x4d, 0xd9, 0xda, 0x9a, 0xff,
	0xdb, 0x9b, 0x2d, 0xe5, 0xb2, 0xcb, 0xdd, 0x03, 0x63, 0xed, 0xfc, 0xec, 0xfd, 0xfb, 0xe7, 0x73,
	0xb6, 0xd0, 0x27, 0xe3, 0x0a, 0x13, 0xcb, 0xdd, 0xd5, 0x83, 0xad, 0x50, 0xea, 0x58, 0x1a, 0xcf,
	0x8d, 0xed, 0xc1, 0x76, 0x03, 0xb6, 0xdb, 0x3f, 0xa0, 0x3c, 0x18, 0xa6, 0xbb, 0xf9, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x5d, 0xe5, 0xc8, 0xc5, 0xa1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	// Authenticate inspects the request's metadata -- for this, an empty argument
	// is just enough. Getting a response means it was authenticated successfully.
	// If the metadata does not contain what is needed to authenticate the
	// request, or the tokens are wrong, the AuthenticationService will return the
	// corresponding error code, with details in the message.
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authn.Authentication/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	// Authenticate inspects the request's metadata -- for this, an empty argument
	// is just enough. Getting a response means it was authenticated successfully.
	// If the metadata does not contain what is needed to authenticate the
	// request, or the tokens are wrong, the AuthenticationService will return the
	// corresponding error code, with details in the message.
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authn.Authentication/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.authn.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Authentication_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/authn/authenticate.proto",
}
