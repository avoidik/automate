// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/authz/v2/authz.proto

package v2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
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

// TODO (sr): consider dropping this validation:
// a) it takes time
// b) bad input will with certainty lead to an "unauthorized" response
type IsAuthorizedReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" toml:"resource,omitempty" mapstructure:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty" toml:"action,omitempty" mapstructure:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *IsAuthorizedReq) Reset()         { *m = IsAuthorizedReq{} }
func (m *IsAuthorizedReq) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedReq) ProtoMessage()    {}
func (*IsAuthorizedReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{0}
}

func (m *IsAuthorizedReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedReq.Unmarshal(m, b)
}
func (m *IsAuthorizedReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedReq.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedReq.Merge(m, src)
}
func (m *IsAuthorizedReq) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedReq.Size(m)
}
func (m *IsAuthorizedReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedReq.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedReq proto.InternalMessageInfo

func (m *IsAuthorizedReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *IsAuthorizedReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *IsAuthorizedReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type IsAuthorizedResp struct {
	Authorized           bool     `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty" toml:"authorized,omitempty" mapstructure:"authorized,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *IsAuthorizedResp) Reset()         { *m = IsAuthorizedResp{} }
func (m *IsAuthorizedResp) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedResp) ProtoMessage()    {}
func (*IsAuthorizedResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{1}
}

func (m *IsAuthorizedResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedResp.Unmarshal(m, b)
}
func (m *IsAuthorizedResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedResp.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedResp.Merge(m, src)
}
func (m *IsAuthorizedResp) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedResp.Size(m)
}
func (m *IsAuthorizedResp) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedResp.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedResp proto.InternalMessageInfo

func (m *IsAuthorizedResp) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

type ProjectsAuthorizedReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" toml:"resource,omitempty" mapstructure:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty" toml:"action,omitempty" mapstructure:"action,omitempty"`
	ProjectsFilter       []string `protobuf:"bytes,4,rep,name=projects_filter,json=projectsFilter,proto3" json:"projects_filter,omitempty" toml:"projects_filter,omitempty" mapstructure:"projects_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ProjectsAuthorizedReq) Reset()         { *m = ProjectsAuthorizedReq{} }
func (m *ProjectsAuthorizedReq) String() string { return proto.CompactTextString(m) }
func (*ProjectsAuthorizedReq) ProtoMessage()    {}
func (*ProjectsAuthorizedReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{2}
}

func (m *ProjectsAuthorizedReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectsAuthorizedReq.Unmarshal(m, b)
}
func (m *ProjectsAuthorizedReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectsAuthorizedReq.Marshal(b, m, deterministic)
}
func (m *ProjectsAuthorizedReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectsAuthorizedReq.Merge(m, src)
}
func (m *ProjectsAuthorizedReq) XXX_Size() int {
	return xxx_messageInfo_ProjectsAuthorizedReq.Size(m)
}
func (m *ProjectsAuthorizedReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectsAuthorizedReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectsAuthorizedReq proto.InternalMessageInfo

func (m *ProjectsAuthorizedReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *ProjectsAuthorizedReq) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *ProjectsAuthorizedReq) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ProjectsAuthorizedReq) GetProjectsFilter() []string {
	if m != nil {
		return m.ProjectsFilter
	}
	return nil
}

type ProjectsAuthorizedResp struct {
	Projects             []string `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ProjectsAuthorizedResp) Reset()         { *m = ProjectsAuthorizedResp{} }
func (m *ProjectsAuthorizedResp) String() string { return proto.CompactTextString(m) }
func (*ProjectsAuthorizedResp) ProtoMessage()    {}
func (*ProjectsAuthorizedResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{3}
}

func (m *ProjectsAuthorizedResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectsAuthorizedResp.Unmarshal(m, b)
}
func (m *ProjectsAuthorizedResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectsAuthorizedResp.Marshal(b, m, deterministic)
}
func (m *ProjectsAuthorizedResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectsAuthorizedResp.Merge(m, src)
}
func (m *ProjectsAuthorizedResp) XXX_Size() int {
	return xxx_messageInfo_ProjectsAuthorizedResp.Size(m)
}
func (m *ProjectsAuthorizedResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectsAuthorizedResp.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectsAuthorizedResp proto.InternalMessageInfo

func (m *ProjectsAuthorizedResp) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type FilterAuthorizedPairsReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	Pairs                []*Pair  `protobuf:"bytes,2,rep,name=pairs,proto3" json:"pairs,omitempty" toml:"pairs,omitempty" mapstructure:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedPairsReq) Reset()         { *m = FilterAuthorizedPairsReq{} }
func (m *FilterAuthorizedPairsReq) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedPairsReq) ProtoMessage()    {}
func (*FilterAuthorizedPairsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{4}
}

func (m *FilterAuthorizedPairsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedPairsReq.Unmarshal(m, b)
}
func (m *FilterAuthorizedPairsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedPairsReq.Marshal(b, m, deterministic)
}
func (m *FilterAuthorizedPairsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedPairsReq.Merge(m, src)
}
func (m *FilterAuthorizedPairsReq) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedPairsReq.Size(m)
}
func (m *FilterAuthorizedPairsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedPairsReq.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedPairsReq proto.InternalMessageInfo

func (m *FilterAuthorizedPairsReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

func (m *FilterAuthorizedPairsReq) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type FilterAuthorizedPairsResp struct {
	Pairs                []*Pair  `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty" toml:"pairs,omitempty" mapstructure:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedPairsResp) Reset()         { *m = FilterAuthorizedPairsResp{} }
func (m *FilterAuthorizedPairsResp) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedPairsResp) ProtoMessage()    {}
func (*FilterAuthorizedPairsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{5}
}

func (m *FilterAuthorizedPairsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedPairsResp.Unmarshal(m, b)
}
func (m *FilterAuthorizedPairsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedPairsResp.Marshal(b, m, deterministic)
}
func (m *FilterAuthorizedPairsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedPairsResp.Merge(m, src)
}
func (m *FilterAuthorizedPairsResp) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedPairsResp.Size(m)
}
func (m *FilterAuthorizedPairsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedPairsResp.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedPairsResp proto.InternalMessageInfo

func (m *FilterAuthorizedPairsResp) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type Pair struct {
	Resource             string   `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty" toml:"resource,omitempty" mapstructure:"resource,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty" toml:"action,omitempty" mapstructure:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{6}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *Pair) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type FilterAuthorizedProjectsReq struct {
	Subjects             []string `protobuf:"bytes,1,rep,name=subjects,proto3" json:"subjects,omitempty" toml:"subjects,omitempty" mapstructure:"subjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedProjectsReq) Reset()         { *m = FilterAuthorizedProjectsReq{} }
func (m *FilterAuthorizedProjectsReq) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedProjectsReq) ProtoMessage()    {}
func (*FilterAuthorizedProjectsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{7}
}

func (m *FilterAuthorizedProjectsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedProjectsReq.Unmarshal(m, b)
}
func (m *FilterAuthorizedProjectsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedProjectsReq.Marshal(b, m, deterministic)
}
func (m *FilterAuthorizedProjectsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedProjectsReq.Merge(m, src)
}
func (m *FilterAuthorizedProjectsReq) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedProjectsReq.Size(m)
}
func (m *FilterAuthorizedProjectsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedProjectsReq.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedProjectsReq proto.InternalMessageInfo

func (m *FilterAuthorizedProjectsReq) GetSubjects() []string {
	if m != nil {
		return m.Subjects
	}
	return nil
}

type FilterAuthorizedProjectsResp struct {
	Projects             []string `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty" toml:"projects,omitempty" mapstructure:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *FilterAuthorizedProjectsResp) Reset()         { *m = FilterAuthorizedProjectsResp{} }
func (m *FilterAuthorizedProjectsResp) String() string { return proto.CompactTextString(m) }
func (*FilterAuthorizedProjectsResp) ProtoMessage()    {}
func (*FilterAuthorizedProjectsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_744c9d667935614c, []int{8}
}

func (m *FilterAuthorizedProjectsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAuthorizedProjectsResp.Unmarshal(m, b)
}
func (m *FilterAuthorizedProjectsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAuthorizedProjectsResp.Marshal(b, m, deterministic)
}
func (m *FilterAuthorizedProjectsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAuthorizedProjectsResp.Merge(m, src)
}
func (m *FilterAuthorizedProjectsResp) XXX_Size() int {
	return xxx_messageInfo_FilterAuthorizedProjectsResp.Size(m)
}
func (m *FilterAuthorizedProjectsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAuthorizedProjectsResp.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAuthorizedProjectsResp proto.InternalMessageInfo

func (m *FilterAuthorizedProjectsResp) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*IsAuthorizedReq)(nil), "chef.automate.domain.authz.v2.IsAuthorizedReq")
	proto.RegisterType((*IsAuthorizedResp)(nil), "chef.automate.domain.authz.v2.IsAuthorizedResp")
	proto.RegisterType((*ProjectsAuthorizedReq)(nil), "chef.automate.domain.authz.v2.ProjectsAuthorizedReq")
	proto.RegisterType((*ProjectsAuthorizedResp)(nil), "chef.automate.domain.authz.v2.ProjectsAuthorizedResp")
	proto.RegisterType((*FilterAuthorizedPairsReq)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedPairsReq")
	proto.RegisterType((*FilterAuthorizedPairsResp)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedPairsResp")
	proto.RegisterType((*Pair)(nil), "chef.automate.domain.authz.v2.Pair")
	proto.RegisterType((*FilterAuthorizedProjectsReq)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedProjectsReq")
	proto.RegisterType((*FilterAuthorizedProjectsResp)(nil), "chef.automate.domain.authz.v2.FilterAuthorizedProjectsResp")
}

func init() {
	proto.RegisterFile("api/interservice/authz/v2/authz.proto", fileDescriptor_744c9d667935614c)
}

var fileDescriptor_744c9d667935614c = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x96, 0x4d, 0x6f, 0xd3, 0x3e,
	0x1c, 0xc7, 0xff, 0x6e, 0xf7, 0x9f, 0x3a, 0xff, 0x1f, 0x06, 0x9e, 0xca, 0xb2, 0xee, 0x41, 0xc5,
	0x2b, 0x28, 0x2d, 0x34, 0x01, 0x33, 0x1e, 0x16, 0x04, 0xd3, 0x76, 0x60, 0x9a, 0x40, 0x62, 0x8a,
	0x04, 0x13, 0x9b, 0x5a, 0xf0, 0x52, 0x8f, 0x06, 0xd2, 0x26, 0xb3, 0x9d, 0x1e, 0xba, 0xee, 0x80,
	0xb8, 0x22, 0x21, 0x55, 0xe2, 0xca, 0x5b, 0xe0, 0xce, 0x89, 0x33, 0xef, 0x84, 0x2b, 0xaf, 0x00,
	0x25, 0x4d, 0xda, 0x6d, 0x6d, 0xa9, 0xb6, 0x0b, 0x1c, 0xb8, 0x25, 0xae, 0x3f, 0x5f, 0xff, 0xfc,
	0xfd, 0xfa, 0xe7, 0x06, 0x5e, 0xa2, 0x9e, 0xad, 0xdb, 0x75, 0xc9, 0xb8, 0x60, 0xbc, 0x61, 0x5b,
	0x4c, 0xa7, 0xbe, 0xac, 0x36, 0xf5, 0x06, 0xe9, 0x3c, 0x68, 0x1e, 0x77, 0xa5, 0x8b, 0xe6, 0xad,
	0x2a, 0xdb, 0xd3, 0xa8, 0x2f, 0xdd, 0x1a, 0x95, 0x4c, 0xab, 0xb8, 0x35, 0x6a, 0xd7, 0xb5, 0xce,
	0x8c, 0x06, 0xc9, 0x4c, 0x37, 0xa8, 0x63, 0x57, 0xa8, 0x64, 0x7a, 0xfc, 0xd0, 0xe1, 0xf0, 0xc7,
	0x04, 0x9c, 0xdc, 0x10, 0xab, 0xbe, 0xac, 0xba, 0xdc, 0x6e, 0xb2, 0x8a, 0xc9, 0xf6, 0xd1, 0x1b,
	0x00, 0x53, 0xc2, 0xdf, 0x7d, 0xc5, 0x2c, 0x29, 0x14, 0x90, 0x4d, 0xaa, 0x13, 0x6b, 0xec, 0xf3,
	0xb7, 0x2f, 0xc9, 0x17, 0x6d, 0x50, 0x4a, 0x01, 0xfc, 0x8c, 0x6f, 0x91, 0x27, 0x65, 0x75, 0xc5,
	0x90, 0x8c, 0xd6, 0x5a, 0xbe, 0x60, 0x3c, 0x6f, 0xa8, 0x2b, 0x86, 0xe3, 0x5a, 0xd4, 0x69, 0x39,
	0x15, 0xea, 0xb5, 0x04, 0xad, 0x39, 0x79, 0x63, 0xa7, 0x6c, 0x14, 0x4a, 0x57, 0x72, 0xad, 0xb2,
	0x74, 0x5f, 0xb3, 0xfa, 0x91, 0x57, 0x47, 0x18, 0xd1, 0x5e, 0xa2, 0xc1, 0xf8, 0x37, 0xb3, 0xbb,
	0x2c, 0xba, 0x0f, 0x53, 0x9c, 0x09, 0xd7, 0xe7, 0x16, 0x53, 0x12, 0x59, 0xa0, 0x4e, 0xac, 0xe1,
	0xa0, 0x84, 0x79, 0x3e, 0x4b, 0x66, 0xca, 0x3b, 0xb4, 0xd8, 0x2c, 0x85, 0x4c, 0x41, 0x5d, 0x31,
	0x22, 0x3a, 0x5f, 0xc8, 0x99, 0x5d, 0x06, 0xad, 0xc3, 0x71, 0x6a, 0x49, 0xdb, 0xad, 0x2b, 0xc9,
	0x90, 0xd6, 0x03, 0xba, 0xc0, 0x55, 0x72, 0x39, 0xa2, 0x69, 0xb1, 0xb9, 0x5a, 0xdc, 0x8e, 0x04,
	0x8e, 0x8d, 0xe4, 0x0f, 0xc8, 0x61, 0xce, 0x8c, 0x70, 0x4c, 0xe0, 0xb9, 0xe3, 0xfe, 0x08, 0x0f,
	0x2d, 0x40, 0x48, 0xbb, 0x23, 0x0a, 0xc8, 0x02, 0x35, 0x65, 0x1e, 0x19, 0xc1, 0xdf, 0x13, 0x30,
	0xbd, 0xc9, 0xdd, 0x70, 0x27, 0x7f, 0xac, 0x1d, 0x66, 0x2d, 0x7a, 0x04, 0x27, 0xbd, 0xc8, 0xa5,
	0xe7, 0x7b, 0xb6, 0x23, 0x19, 0x57, 0xc6, 0x42, 0x4b, 0x16, 0x03, 0xc5, 0x85, 0x36, 0x98, 0x55,
	0x00, 0x9e, 0xe6, 0x69, 0x32, 0x15, 0x0a, 0x5f, 0x2b, 0x2e, 0xab, 0xf9, 0x62, 0xe9, 0xe0, 0xfa,
	0xd5, 0x5b, 0x4b, 0x87, 0x39, 0xf3, 0xff, 0x98, 0x7d, 0x10, 0xa2, 0x78, 0x0b, 0x5e, 0x18, 0xe4,
	0xb9, 0xf0, 0xd0, 0x3d, 0x98, 0x8a, 0xe7, 0x46, 0x9e, 0x5f, 0x0c, 0x16, 0x98, 0x6b, 0x83, 0x19,
	0x05, 0xe0, 0x34, 0x9f, 0x22, 0xe7, 0xe3, 0x05, 0x7a, 0xf2, 0x5d, 0x04, 0x7f, 0x05, 0x50, 0xe9,
	0xac, 0xd1, 0xd3, 0xdd, 0xa4, 0x36, 0x17, 0x41, 0xa0, 0xa2, 0x2f, 0xcf, 0xad, 0x40, 0xdb, 0x6c,
	0x83, 0xc7, 0x29, 0x80, 0x1f, 0xf2, 0x0d, 0xb2, 0x1e, 0xe4, 0x39, 0x32, 0xd2, 0x56, 0x98, 0x64,
	0xab, 0x3f, 0xc0, 0xfc, 0x80, 0x04, 0x97, 0xe1, 0xdf, 0x5e, 0x50, 0x80, 0x92, 0xc8, 0x26, 0xd5,
	0x7f, 0xc8, 0xa2, 0xf6, 0xd3, 0xe6, 0xd7, 0x82, 0x62, 0xcd, 0x0e, 0x81, 0x9f, 0xc2, 0x99, 0x21,
	0x7b, 0x11, 0x5e, 0x4f, 0x17, 0x9c, 0x5a, 0xf7, 0x3d, 0x80, 0x63, 0xc1, 0xfb, 0xef, 0xd3, 0xb8,
	0x6d, 0x00, 0x67, 0xfb, 0xb6, 0x1a, 0x65, 0xfa, 0xab, 0x92, 0xc3, 0x06, 0x9c, 0x1b, 0x5e, 0x93,
	0xf0, 0x50, 0xe6, 0xe4, 0x51, 0xed, 0x9d, 0x43, 0xf2, 0x69, 0x0c, 0xfe, 0x17, 0x63, 0x34, 0x6c,
	0xa0, 0x7d, 0xf8, 0xef, 0xd1, 0xbb, 0x09, 0x69, 0x23, 0x02, 0x3b, 0x71, 0xd1, 0x67, 0xf4, 0x53,
	0xcd, 0x17, 0x1e, 0xfe, 0x0b, 0xbd, 0x03, 0x30, 0x3d, 0xf0, 0x00, 0xa1, 0xdb, 0x23, 0xc4, 0x86,
	0xb5, 0x50, 0xe6, 0xce, 0xd9, 0xc0, 0xb0, 0x9c, 0x0f, 0x83, 0x7a, 0x33, 0x32, 0x0c, 0x19, 0xa7,
	0x15, 0xee, 0x9d, 0x8e, 0xcc, 0xdd, 0x33, 0xb3, 0x61, 0x5d, 0x6f, 0x01, 0x44, 0xfd, 0xb7, 0x11,
	0x5a, 0x1a, 0xd5, 0x51, 0x83, 0xfe, 0x34, 0x32, 0x37, 0xcf, 0x40, 0x05, 0x55, 0xac, 0x2d, 0x6d,
	0x93, 0x97, 0xb6, 0xac, 0xfa, 0xbb, 0x9a, 0xe5, 0xd6, 0xf4, 0x40, 0x44, 0x8f, 0x45, 0xf4, 0xa1,
	0x9f, 0x15, 0xbb, 0xe3, 0xe1, 0x97, 0xc1, 0x8d, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x34,
	0x2f, 0x5f, 0x7a, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	IsAuthorized(ctx context.Context, in *IsAuthorizedReq, opts ...grpc.CallOption) (*IsAuthorizedResp, error)
	FilterAuthorizedPairs(ctx context.Context, in *FilterAuthorizedPairsReq, opts ...grpc.CallOption) (*FilterAuthorizedPairsResp, error)
	FilterAuthorizedProjects(ctx context.Context, in *FilterAuthorizedProjectsReq, opts ...grpc.CallOption) (*FilterAuthorizedProjectsResp, error)
	ProjectsAuthorized(ctx context.Context, in *ProjectsAuthorizedReq, opts ...grpc.CallOption) (*ProjectsAuthorizedResp, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) IsAuthorized(ctx context.Context, in *IsAuthorizedReq, opts ...grpc.CallOption) (*IsAuthorizedResp, error) {
	out := new(IsAuthorizedResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/IsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) FilterAuthorizedPairs(ctx context.Context, in *FilterAuthorizedPairsReq, opts ...grpc.CallOption) (*FilterAuthorizedPairsResp, error) {
	out := new(FilterAuthorizedPairsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedPairs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) FilterAuthorizedProjects(ctx context.Context, in *FilterAuthorizedProjectsReq, opts ...grpc.CallOption) (*FilterAuthorizedProjectsResp, error) {
	out := new(FilterAuthorizedProjectsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) ProjectsAuthorized(ctx context.Context, in *ProjectsAuthorizedReq, opts ...grpc.CallOption) (*ProjectsAuthorizedResp, error) {
	out := new(ProjectsAuthorizedResp)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.authz.v2.Authorization/ProjectsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	IsAuthorized(context.Context, *IsAuthorizedReq) (*IsAuthorizedResp, error)
	FilterAuthorizedPairs(context.Context, *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error)
	FilterAuthorizedProjects(context.Context, *FilterAuthorizedProjectsReq) (*FilterAuthorizedProjectsResp, error)
	ProjectsAuthorized(context.Context, *ProjectsAuthorizedReq) (*ProjectsAuthorizedResp, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthorizedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IsAuthorized(ctx, req.(*IsAuthorizedReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_FilterAuthorizedPairs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterAuthorizedPairsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).FilterAuthorizedPairs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedPairs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).FilterAuthorizedPairs(ctx, req.(*FilterAuthorizedPairsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_FilterAuthorizedProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterAuthorizedProjectsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).FilterAuthorizedProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/FilterAuthorizedProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).FilterAuthorizedProjects(ctx, req.(*FilterAuthorizedProjectsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_ProjectsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectsAuthorizedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).ProjectsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.authz.v2.Authorization/ProjectsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).ProjectsAuthorized(ctx, req.(*ProjectsAuthorizedReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.authz.v2.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuthorized",
			Handler:    _Authorization_IsAuthorized_Handler,
		},
		{
			MethodName: "FilterAuthorizedPairs",
			Handler:    _Authorization_FilterAuthorizedPairs_Handler,
		},
		{
			MethodName: "FilterAuthorizedProjects",
			Handler:    _Authorization_FilterAuthorizedProjects_Handler,
		},
		{
			MethodName: "ProjectsAuthorized",
			Handler:    _Authorization_ProjectsAuthorized_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interservice/authz/v2/authz.proto",
}
