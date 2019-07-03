// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/compliance-service/ingest/events/compliance/compliance.proto

package compliance

import (
	fmt "fmt"
	common "github.com/chef/automate/components/compliance-service/api/common"
	inspec "github.com/chef/automate/components/compliance-service/ingest/events/inspec"
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

type Report struct {
	// inspec full json report fields
	Version     string             `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Platform    *inspec.Platform   `protobuf:"bytes,16,opt,name=platform,proto3" json:"platform,omitempty"`
	Statistics  *inspec.Statistics `protobuf:"bytes,17,opt,name=statistics,proto3" json:"statistics,omitempty"`
	Profiles    []*inspec.Profile  `protobuf:"bytes,18,rep,name=profiles,proto3" json:"profiles,omitempty"`
	OtherChecks []string           `protobuf:"bytes,19,rep,name=other_checks,json=otherChecks,proto3" json:"other_checks,omitempty"`
	// extra report fields added by the audit cookbook
	ReportUuid           string       `protobuf:"bytes,20,opt,name=report_uuid,json=reportUuid,proto3" json:"report_uuid,omitempty"`
	NodeUuid             string       `protobuf:"bytes,21,opt,name=node_uuid,json=nodeUuid,proto3" json:"node_uuid,omitempty"`
	JobUuid              string       `protobuf:"bytes,22,opt,name=job_uuid,json=jobUuid,proto3" json:"job_uuid,omitempty"`
	NodeName             string       `protobuf:"bytes,23,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Environment          string       `protobuf:"bytes,24,opt,name=environment,proto3" json:"environment,omitempty"`
	Roles                []string     `protobuf:"bytes,25,rep,name=roles,proto3" json:"roles,omitempty"`
	Recipes              []string     `protobuf:"bytes,26,rep,name=recipes,proto3" json:"recipes,omitempty"`
	EndTime              string       `protobuf:"bytes,27,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Type                 string       `protobuf:"bytes,28,opt,name=type,proto3" json:"type,omitempty"`
	SourceId             string       `protobuf:"bytes,29,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	SourceRegion         string       `protobuf:"bytes,30,opt,name=source_region,json=sourceRegion,proto3" json:"source_region,omitempty"`
	SourceAccountId      string       `protobuf:"bytes,31,opt,name=source_account_id,json=sourceAccountId,proto3" json:"source_account_id,omitempty"`
	PolicyName           string       `protobuf:"bytes,32,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyGroup          string       `protobuf:"bytes,33,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	OrganizationName     string       `protobuf:"bytes,34,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	SourceFqdn           string       `protobuf:"bytes,35,opt,name=source_fqdn,json=sourceFqdn,proto3" json:"source_fqdn,omitempty"`
	ChefTags             []string     `protobuf:"bytes,36,rep,name=chef_tags,json=chefTags,proto3" json:"chef_tags,omitempty"`
	Ipaddress            string       `protobuf:"bytes,37,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	Fqdn                 string       `protobuf:"bytes,38,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Tags                 []*common.Kv `protobuf:"bytes,39,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e96c54e075a0df3, []int{0}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Report) GetPlatform() *inspec.Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

func (m *Report) GetStatistics() *inspec.Statistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *Report) GetProfiles() []*inspec.Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Report) GetOtherChecks() []string {
	if m != nil {
		return m.OtherChecks
	}
	return nil
}

func (m *Report) GetReportUuid() string {
	if m != nil {
		return m.ReportUuid
	}
	return ""
}

func (m *Report) GetNodeUuid() string {
	if m != nil {
		return m.NodeUuid
	}
	return ""
}

func (m *Report) GetJobUuid() string {
	if m != nil {
		return m.JobUuid
	}
	return ""
}

func (m *Report) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Report) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *Report) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Report) GetRecipes() []string {
	if m != nil {
		return m.Recipes
	}
	return nil
}

func (m *Report) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Report) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Report) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Report) GetSourceRegion() string {
	if m != nil {
		return m.SourceRegion
	}
	return ""
}

func (m *Report) GetSourceAccountId() string {
	if m != nil {
		return m.SourceAccountId
	}
	return ""
}

func (m *Report) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *Report) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *Report) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *Report) GetSourceFqdn() string {
	if m != nil {
		return m.SourceFqdn
	}
	return ""
}

func (m *Report) GetChefTags() []string {
	if m != nil {
		return m.ChefTags
	}
	return nil
}

func (m *Report) GetIpaddress() string {
	if m != nil {
		return m.Ipaddress
	}
	return ""
}

func (m *Report) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Report) GetTags() []*common.Kv {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Report)(nil), "chef.automate.domain.compliance.ingest.events.compliance.Report")
}

func init() {
	proto.RegisterFile("components/compliance-service/ingest/events/compliance/compliance.proto", fileDescriptor_8e96c54e075a0df3)
}

var fileDescriptor_8e96c54e075a0df3 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x5d, 0x4f, 0x14, 0x3f,
	0x14, 0xc6, 0x43, 0xe0, 0x0f, 0xbb, 0x5d, 0xfe, 0x11, 0x2a, 0x6a, 0x79, 0x51, 0x16, 0xf0, 0x65,
	0xa3, 0x71, 0x36, 0xc1, 0x98, 0x78, 0xe3, 0x1b, 0x26, 0x12, 0x62, 0xe2, 0xcb, 0x88, 0x17, 0x72,
	0xb3, 0x76, 0xdb, 0xb3, 0xb3, 0xc5, 0x9d, 0xb6, 0xb4, 0x9d, 0x4d, 0xf0, 0x13, 0xfb, 0x31, 0x4c,
	0x7b, 0x06, 0x18, 0xbd, 0xd0, 0xc0, 0xd5, 0x74, 0x9e, 0xe7, 0x9c, 0xdf, 0x9c, 0x79, 0xf6, 0xec,
	0x90, 0x7d, 0x61, 0x4a, 0x6b, 0x34, 0xe8, 0xe0, 0xfb, 0xf1, 0x38, 0x51, 0x5c, 0x0b, 0x78, 0xec,
	0xc1, 0x4d, 0x95, 0x80, 0xbe, 0xd2, 0x05, 0xf8, 0xd0, 0x87, 0xe9, 0x1f, 0x05, 0x8d, 0x63, 0x66,
	0x9d, 0x09, 0x86, 0x3e, 0x13, 0x63, 0x18, 0x65, 0xbc, 0x0a, 0xa6, 0xe4, 0x01, 0x32, 0x69, 0x4a,
	0xae, 0x74, 0xd6, 0x28, 0x43, 0x54, 0x86, 0xa8, 0x86, 0xb1, 0xf6, 0xf2, 0x32, 0x23, 0x28, 0xed,
	0x2d, 0x88, 0xfa, 0x82, 0x8f, 0x5e, 0x7b, 0xfa, 0x77, 0x00, 0xb7, 0x2a, 0xca, 0xa5, 0xd1, 0xf5,
	0x05, 0xdb, 0xb6, 0x7f, 0x2e, 0x90, 0xf9, 0x1c, 0xac, 0x71, 0x81, 0x32, 0xb2, 0x30, 0x05, 0xe7,
	0x95, 0xd1, 0x6c, 0xa6, 0x3b, 0xd3, 0x6b, 0xe7, 0x67, 0xb7, 0xf4, 0x88, 0xb4, 0xec, 0x84, 0x87,
	0x91, 0x71, 0x25, 0x5b, 0xea, 0xce, 0xf4, 0x3a, 0xbb, 0x2f, 0xb2, 0xcb, 0xbd, 0x69, 0x3d, 0xea,
	0xc7, 0x9a, 0x92, 0x9f, 0xf3, 0xe8, 0x37, 0x42, 0x7c, 0xe0, 0x41, 0xf9, 0xa0, 0x84, 0x67, 0xcb,
	0x89, 0xfe, 0xea, 0x6a, 0xf4, 0xcf, 0xe7, 0x9c, 0xbc, 0xc1, 0xa4, 0x5f, 0x49, 0xcb, 0x3a, 0x33,
	0x52, 0x13, 0xf0, 0x8c, 0x76, 0x67, 0x7b, 0x9d, 0xdd, 0xe7, 0x57, 0x9c, 0x1e, 0x29, 0xf9, 0x39,
	0x8e, 0x6e, 0x91, 0x45, 0x13, 0xc6, 0xe0, 0x06, 0x62, 0x0c, 0xe2, 0xbb, 0x67, 0xd7, 0xbb, 0xb3,
	0xbd, 0x76, 0xde, 0x49, 0xda, 0x9b, 0x24, 0xd1, 0x4d, 0xd2, 0x71, 0x29, 0xdf, 0x41, 0x55, 0x29,
	0xc9, 0x56, 0x52, 0xb2, 0x04, 0xa5, 0x2f, 0x95, 0x92, 0x74, 0x9d, 0xb4, 0xb5, 0x91, 0x80, 0xf6,
	0x8d, 0x64, 0xb7, 0xa2, 0x90, 0xcc, 0x55, 0xd2, 0x3a, 0x36, 0x43, 0xf4, 0x6e, 0xe2, 0x8f, 0x72,
	0x6c, 0x86, 0xbf, 0xf5, 0x69, 0x5e, 0x02, 0xbb, 0x75, 0xd1, 0xf7, 0x9e, 0x97, 0x40, 0xbb, 0xa4,
	0x03, 0x7a, 0xaa, 0x9c, 0xd1, 0x25, 0xe8, 0xc0, 0x58, 0xb2, 0x9b, 0x12, 0x5d, 0x21, 0xff, 0x39,
	0x13, 0x23, 0x59, 0x4d, 0x33, 0xe3, 0x4d, 0xdc, 0x01, 0x07, 0x42, 0x59, 0xf0, 0x6c, 0x2d, 0xe9,
	0x67, 0xb7, 0x71, 0x12, 0xd0, 0x72, 0x10, 0x54, 0x09, 0x6c, 0x1d, 0x27, 0x01, 0x2d, 0x0f, 0x55,
	0x09, 0x94, 0x92, 0xb9, 0x70, 0x6a, 0x81, 0x6d, 0x24, 0x39, 0x9d, 0xe3, 0x74, 0xde, 0x54, 0x4e,
	0xc0, 0x40, 0x49, 0x76, 0x1b, 0xa7, 0x43, 0xe1, 0x40, 0xd2, 0x1d, 0xf2, 0x7f, 0x6d, 0x3a, 0x28,
	0xe2, 0xbe, 0xdd, 0x49, 0x05, 0x8b, 0x28, 0xe6, 0x49, 0xa3, 0x0f, 0xc9, 0x72, 0x5d, 0xc4, 0x85,
	0x30, 0x95, 0x0e, 0x91, 0xb4, 0x99, 0x0a, 0xaf, 0xa1, 0xf1, 0x1a, 0xf5, 0x03, 0x19, 0x43, 0xb6,
	0x66, 0xa2, 0xc4, 0x29, 0xa6, 0xd1, 0xc5, 0x90, 0x51, 0x4a, 0x79, 0x6c, 0x91, 0xc5, 0xba, 0xa0,
	0x70, 0xa6, 0xb2, 0x6c, 0x0b, 0x03, 0x41, 0x6d, 0x3f, 0x4a, 0xf4, 0x11, 0x59, 0x36, 0xae, 0xe0,
	0x5a, 0xfd, 0xe0, 0x41, 0x19, 0x8d, 0xa4, 0xed, 0x54, 0xb7, 0xd4, 0x34, 0x12, 0x6f, 0x93, 0x74,
	0xea, 0xe1, 0x46, 0x27, 0x52, 0xb3, 0x1d, 0x7c, 0x20, 0x4a, 0x6f, 0x4f, 0xa4, 0x8e, 0xef, 0x1f,
	0x77, 0x6c, 0x10, 0x78, 0xe1, 0xd9, 0xdd, 0x14, 0x65, 0x2b, 0x0a, 0x87, 0xbc, 0xf0, 0x74, 0x83,
	0xb4, 0x95, 0xe5, 0x52, 0x3a, 0xf0, 0x9e, 0xdd, 0x4b, 0xbd, 0x17, 0x42, 0x8c, 0x33, 0x41, 0xef,
	0x63, 0x9c, 0xf1, 0x4c, 0xf7, 0xc8, 0x5c, 0x22, 0x3d, 0x48, 0xfb, 0x9b, 0xfd, 0x73, 0x7f, 0xb9,
	0x55, 0x59, 0xfd, 0x3f, 0x7f, 0x37, 0xcd, 0x53, 0xef, 0xde, 0xa7, 0xa3, 0x0f, 0x85, 0x0a, 0xe3,
	0x6a, 0x18, 0x9d, 0x7e, 0x24, 0xf4, 0xcf, 0x08, 0xfd, 0xab, 0x7d, 0x00, 0x87, 0xf3, 0xe9, 0x23,
	0xf2, 0xe4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x48, 0x50, 0x63, 0x41, 0x05, 0x00, 0x00,
}
