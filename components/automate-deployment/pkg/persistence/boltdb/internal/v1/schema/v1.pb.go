// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-deployment/pkg/persistence/boltdb/internal/v1/schema/v1.proto

package schema

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Service_DeploymentState int32

const (
	Service_SKIP      Service_DeploymentState = 0
	Service_INSTALLED Service_DeploymentState = 1
	Service_RUNNING   Service_DeploymentState = 2
	Service_REMOVED   Service_DeploymentState = 3
)

var Service_DeploymentState_name = map[int32]string{
	0: "SKIP",
	1: "INSTALLED",
	2: "RUNNING",
	3: "REMOVED",
}

var Service_DeploymentState_value = map[string]int32{
	"SKIP":      0,
	"INSTALLED": 1,
	"RUNNING":   2,
	"REMOVED":   3,
}

func (x Service_DeploymentState) String() string {
	return proto.EnumName(Service_DeploymentState_name, int32(x))
}

func (Service_DeploymentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4545dc933ad0947d, []int{3, 0}
}

type Service_InstallType int32

const (
	Service_DEPOT Service_InstallType = 0
	Service_HART  Service_InstallType = 1
)

var Service_InstallType_name = map[int32]string{
	0: "DEPOT",
	1: "HART",
}

var Service_InstallType_value = map[string]int32{
	"DEPOT": 0,
	"HART":  1,
}

func (x Service_InstallType) String() string {
	return proto.EnumName(Service_InstallType_name, int32(x))
}

func (Service_InstallType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4545dc933ad0947d, []int{3, 1}
}

type Deployment struct {
	Id                     string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	CreatedAt              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" toml:"created_at,omitempty" mapstructure:"created_at,omitempty"`
	LastAction             string               `protobuf:"bytes,3,opt,name=last_action,json=lastAction,proto3" json:"last_action,omitempty" toml:"last_action,omitempty" mapstructure:"last_action,omitempty"`
	Deployed               bool                 `protobuf:"varint,4,opt,name=deployed,proto3" json:"deployed,omitempty" toml:"deployed,omitempty" mapstructure:"deployed,omitempty"`
	Services               []*Service           `protobuf:"bytes,5,rep,name=services,proto3" json:"services,omitempty" toml:"services,omitempty" mapstructure:"services,omitempty"`
	CurrentReleaseManifest *ReleaseManifest     `protobuf:"bytes,6,opt,name=current_release_manifest,json=currentReleaseManifest,proto3" json:"current_release_manifest,omitempty" toml:"current_release_manifest,omitempty" mapstructure:"current_release_manifest,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized       []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache          int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_4545dc933ad0947d, []int{0}
}

func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (m *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(m, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Deployment) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Deployment) GetLastAction() string {
	if m != nil {
		return m.LastAction
	}
	return ""
}

func (m *Deployment) GetDeployed() bool {
	if m != nil {
		return m.Deployed
	}
	return false
}

func (m *Deployment) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Deployment) GetCurrentReleaseManifest() *ReleaseManifest {
	if m != nil {
		return m.CurrentReleaseManifest
	}
	return nil
}

type ReleaseManifest struct {
	Build                string   `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty" toml:"build,omitempty" mapstructure:"build,omitempty"`
	BuildSHA             string   `protobuf:"bytes,2,opt,name=buildSHA,proto3" json:"buildSHA,omitempty" toml:"buildSHA,omitempty" mapstructure:"buildSHA,omitempty"`
	Packages             []string `protobuf:"bytes,3,rep,name=packages,proto3" json:"packages,omitempty" toml:"packages,omitempty" mapstructure:"packages,omitempty"`
	Harts                []*Hart  `protobuf:"bytes,4,rep,name=harts,proto3" json:"harts,omitempty" toml:"harts,omitempty" mapstructure:"harts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ReleaseManifest) Reset()         { *m = ReleaseManifest{} }
func (m *ReleaseManifest) String() string { return proto.CompactTextString(m) }
func (*ReleaseManifest) ProtoMessage()    {}
func (*ReleaseManifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4545dc933ad0947d, []int{1}
}

func (m *ReleaseManifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseManifest.Unmarshal(m, b)
}
func (m *ReleaseManifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseManifest.Marshal(b, m, deterministic)
}
func (m *ReleaseManifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseManifest.Merge(m, src)
}
func (m *ReleaseManifest) XXX_Size() int {
	return xxx_messageInfo_ReleaseManifest.Size(m)
}
func (m *ReleaseManifest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseManifest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseManifest proto.InternalMessageInfo

func (m *ReleaseManifest) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *ReleaseManifest) GetBuildSHA() string {
	if m != nil {
		return m.BuildSHA
	}
	return ""
}

func (m *ReleaseManifest) GetPackages() []string {
	if m != nil {
		return m.Packages
	}
	return nil
}

func (m *ReleaseManifest) GetHarts() []*Hart {
	if m != nil {
		return m.Harts
	}
	return nil
}

type Hart struct {
	Origin               string   `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty" toml:"origin,omitempty" mapstructure:"origin,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty" toml:"path,omitempty" mapstructure:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Hart) Reset()         { *m = Hart{} }
func (m *Hart) String() string { return proto.CompactTextString(m) }
func (*Hart) ProtoMessage()    {}
func (*Hart) Descriptor() ([]byte, []int) {
	return fileDescriptor_4545dc933ad0947d, []int{2}
}

func (m *Hart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hart.Unmarshal(m, b)
}
func (m *Hart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hart.Marshal(b, m, deterministic)
}
func (m *Hart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hart.Merge(m, src)
}
func (m *Hart) XXX_Size() int {
	return xxx_messageInfo_Hart.Size(m)
}
func (m *Hart) XXX_DiscardUnknown() {
	xxx_messageInfo_Hart.DiscardUnknown(m)
}

var xxx_messageInfo_Hart proto.InternalMessageInfo

func (m *Hart) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Hart) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hart) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type Service struct {
	Origin               string                  `protobuf:"bytes,1,opt,name=origin,proto3" json:"origin,omitempty" toml:"origin,omitempty" mapstructure:"origin,omitempty"`
	OriginalOrigin       string                  `protobuf:"bytes,2,opt,name=original_origin,json=originalOrigin,proto3" json:"original_origin,omitempty" toml:"original_origin,omitempty" mapstructure:"original_origin,omitempty"`
	Name                 string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	InstallType          Service_InstallType     `protobuf:"varint,4,opt,name=install_type,json=installType,proto3,enum=Service_InstallType" json:"install_type,omitempty" toml:"install_type,omitempty" mapstructure:"install_type,omitempty"`
	HartPath             string                  `protobuf:"bytes,5,opt,name=hart_path,json=hartPath,proto3" json:"hart_path,omitempty" toml:"hart_path,omitempty" mapstructure:"hart_path,omitempty"`
	DeploymentState      Service_DeploymentState `protobuf:"varint,8,opt,name=deployment_state,json=deploymentState,proto3,enum=Service_DeploymentState" json:"deployment_state,omitempty" toml:"deployment_state,omitempty" mapstructure:"deployment_state,omitempty"`
	SslKey               string                  `protobuf:"bytes,9,opt,name=ssl_key,json=sslKey,proto3" json:"ssl_key,omitempty" toml:"ssl_key,omitempty" mapstructure:"ssl_key,omitempty"`
	SslCert              string                  `protobuf:"bytes,10,opt,name=ssl_cert,json=sslCert,proto3" json:"ssl_cert,omitempty" toml:"ssl_cert,omitempty" mapstructure:"ssl_cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                   `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_4545dc933ad0947d, []int{3}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Service) GetOriginalOrigin() string {
	if m != nil {
		return m.OriginalOrigin
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetInstallType() Service_InstallType {
	if m != nil {
		return m.InstallType
	}
	return Service_DEPOT
}

func (m *Service) GetHartPath() string {
	if m != nil {
		return m.HartPath
	}
	return ""
}

func (m *Service) GetDeploymentState() Service_DeploymentState {
	if m != nil {
		return m.DeploymentState
	}
	return Service_SKIP
}

func (m *Service) GetSslKey() string {
	if m != nil {
		return m.SslKey
	}
	return ""
}

func (m *Service) GetSslCert() string {
	if m != nil {
		return m.SslCert
	}
	return ""
}

func init() {
	proto.RegisterEnum("Service_DeploymentState", Service_DeploymentState_name, Service_DeploymentState_value)
	proto.RegisterEnum("Service_InstallType", Service_InstallType_name, Service_InstallType_value)
	proto.RegisterType((*Deployment)(nil), "Deployment")
	proto.RegisterType((*ReleaseManifest)(nil), "ReleaseManifest")
	proto.RegisterType((*Hart)(nil), "Hart")
	proto.RegisterType((*Service)(nil), "Service")
}

func init() {
	proto.RegisterFile("components/automate-deployment/pkg/persistence/boltdb/internal/v1/schema/v1.proto", fileDescriptor_4545dc933ad0947d)
}

var fileDescriptor_4545dc933ad0947d = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xda, 0x4a,
	0x14, 0x0d, 0x60, 0xc0, 0x5c, 0xe7, 0x05, 0x6b, 0x14, 0xe5, 0xf9, 0x25, 0x8b, 0x20, 0xeb, 0x49,
	0x65, 0x53, 0x5b, 0x49, 0x17, 0x55, 0x97, 0x34, 0xd0, 0x06, 0x92, 0x90, 0xd4, 0xd0, 0x2e, 0xba,
	0xb1, 0x06, 0xfb, 0x06, 0x46, 0xf1, 0x97, 0x3c, 0x97, 0x48, 0x48, 0x5d, 0xf7, 0xbf, 0xf4, 0x5f,
	0x56, 0x1e, 0x1b, 0xf2, 0x21, 0x75, 0xd7, 0xdd, 0x3d, 0xe7, 0x5c, 0x9f, 0xfb, 0x31, 0x33, 0x86,
	0x2f, 0x41, 0x1a, 0x67, 0x69, 0x82, 0x09, 0x49, 0x97, 0xaf, 0x29, 0x8d, 0x39, 0xe1, 0xdb, 0x10,
	0xb3, 0x28, 0xdd, 0xc4, 0x98, 0x90, 0x9b, 0x3d, 0x2c, 0xdd, 0x0c, 0x73, 0x29, 0x24, 0x61, 0x12,
	0xa0, 0xbb, 0x48, 0x23, 0x0a, 0x17, 0xae, 0x48, 0x08, 0xf3, 0x84, 0x47, 0xee, 0xe3, 0x99, 0x2b,
	0x83, 0x15, 0xc6, 0xdc, 0x7d, 0x3c, 0x73, 0xb2, 0x3c, 0xa5, 0xf4, 0xf8, 0x74, 0x99, 0xa6, 0xcb,
	0x08, 0x5d, 0x85, 0x16, 0xeb, 0x7b, 0x97, 0x44, 0x8c, 0x92, 0x78, 0x9c, 0x95, 0x09, 0xf6, 0xcf,
	0x3a, 0xc0, 0x70, 0x57, 0x82, 0x1d, 0x40, 0x5d, 0x84, 0x56, 0xad, 0x57, 0xeb, 0x77, 0xbc, 0xba,
	0x08, 0xd9, 0x07, 0x80, 0x20, 0x47, 0x4e, 0x18, 0xfa, 0x9c, 0xac, 0x7a, 0xaf, 0xd6, 0x37, 0xce,
	0x8f, 0x9d, 0xd2, 0xd4, 0xd9, 0x9a, 0x3a, 0xf3, 0xad, 0xa9, 0xd7, 0xa9, 0xb2, 0x07, 0xc4, 0x4e,
	0xc1, 0x88, 0xb8, 0x24, 0x9f, 0x07, 0x24, 0xd2, 0xc4, 0x6a, 0x28, 0x4f, 0x28, 0xa8, 0x81, 0x62,
	0xd8, 0x31, 0xe8, 0xe5, 0x70, 0x18, 0x5a, 0x5a, 0xaf, 0xd6, 0xd7, 0xbd, 0x1d, 0x66, 0xff, 0x83,
	0x2e, 0x31, 0x7f, 0x14, 0x01, 0x4a, 0xab, 0xd9, 0x6b, 0xf4, 0x8d, 0x73, 0xdd, 0x99, 0x95, 0x84,
	0xb7, 0x53, 0xd8, 0x04, 0xac, 0x60, 0x9d, 0xe7, 0x98, 0x90, 0x9f, 0x63, 0x84, 0x5c, 0xa2, 0x1f,
	0xf3, 0x44, 0xdc, 0xa3, 0x24, 0xab, 0xa5, 0x7a, 0x35, 0x1d, 0xaf, 0x14, 0x6e, 0x2a, 0xde, 0x3b,
	0xaa, 0xbe, 0x78, 0xc5, 0xdb, 0x3f, 0xa0, 0xfb, 0x8a, 0x62, 0x87, 0xd0, 0x5c, 0xac, 0x45, 0xb4,
	0xdd, 0x47, 0x09, 0x8a, 0xb6, 0x55, 0x30, 0xbb, 0x1c, 0xa8, 0x85, 0x74, 0xbc, 0x1d, 0x2e, 0xb4,
	0x8c, 0x07, 0x0f, 0x7c, 0x89, 0xd2, 0x6a, 0xf4, 0x1a, 0x85, 0xb6, 0xc5, 0xec, 0x04, 0x9a, 0x2b,
	0x9e, 0x93, 0xb4, 0x34, 0x35, 0x4f, 0xd3, 0xb9, 0xe4, 0x39, 0x79, 0x25, 0x67, 0x7f, 0x02, 0xad,
	0x80, 0xec, 0x08, 0x5a, 0x69, 0x2e, 0x96, 0x22, 0xa9, 0x6a, 0x56, 0x88, 0x31, 0xd0, 0x12, 0x1e,
	0x63, 0x55, 0x50, 0xc5, 0x05, 0x97, 0x71, 0x5a, 0x55, 0x9b, 0x55, 0xb1, 0xfd, 0xab, 0x01, 0xed,
	0x6a, 0x4f, 0x7f, 0xf4, 0x7a, 0x03, 0xdd, 0x32, 0xe2, 0x91, 0x5f, 0x25, 0x94, 0xb6, 0x07, 0x5b,
	0xfa, 0xf6, 0x65, 0xd1, 0xc6, 0xb3, 0xa2, 0xef, 0x61, 0x5f, 0x24, 0x92, 0x78, 0x14, 0xf9, 0xb4,
	0xc9, 0x50, 0x1d, 0xdc, 0xc1, 0xf9, 0xe1, 0xf6, 0x70, 0x9c, 0x71, 0x29, 0xce, 0x37, 0x19, 0x7a,
	0x86, 0x78, 0x02, 0xec, 0x04, 0x3a, 0xc5, 0xa8, 0xbe, 0x6a, 0xb9, 0x59, 0xee, 0xad, 0x20, 0xee,
	0x38, 0xad, 0xd8, 0x05, 0x98, 0x4f, 0xf7, 0xdc, 0x97, 0xc4, 0x09, 0x2d, 0x5d, 0x39, 0x5b, 0x3b,
	0xe7, 0xa7, 0x5b, 0x3a, 0x2b, 0x74, 0xaf, 0x1b, 0xbe, 0x24, 0xd8, 0xbf, 0xd0, 0x96, 0x32, 0xf2,
	0x1f, 0x70, 0x63, 0x75, 0xca, 0x81, 0xa5, 0x8c, 0xae, 0x70, 0xc3, 0xfe, 0x03, 0xbd, 0x10, 0x02,
	0xcc, 0xc9, 0x02, 0xa5, 0x14, 0x89, 0x17, 0x98, 0x93, 0x3d, 0x84, 0xee, 0x2b, 0x5f, 0xa6, 0x83,
	0x36, 0xbb, 0x1a, 0xdf, 0x99, 0x7b, 0xec, 0x1f, 0xe8, 0x8c, 0xa7, 0xb3, 0xf9, 0xe0, 0xfa, 0x7a,
	0x34, 0x34, 0x6b, 0xcc, 0x80, 0xb6, 0xf7, 0x75, 0x3a, 0x1d, 0x4f, 0x3f, 0x9b, 0x75, 0x05, 0x46,
	0x37, 0xb7, 0xdf, 0x46, 0x43, 0xb3, 0x61, 0xdb, 0x60, 0x3c, 0x9b, 0x9b, 0x75, 0xa0, 0x39, 0x1c,
	0xdd, 0xdd, 0xce, 0xcd, 0xbd, 0xc2, 0xec, 0x72, 0xe0, 0xcd, 0xcd, 0xda, 0x44, 0xd3, 0x5b, 0x66,
	0x7b, 0xa2, 0xe9, 0x6d, 0x53, 0x9f, 0x68, 0xba, 0x61, 0xee, 0x7f, 0x0c, 0xbe, 0xf3, 0xa5, 0xa0,
	0xd5, 0x7a, 0xe1, 0x04, 0x69, 0xec, 0x06, 0x2b, 0xbc, 0xdf, 0xbd, 0x7a, 0xf7, 0x6f, 0xfd, 0x09,
	0x16, 0x2d, 0xf5, 0x48, 0xdf, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x93, 0x78, 0x8d, 0x61, 0x5c,
	0x04, 0x00, 0x00,
}
