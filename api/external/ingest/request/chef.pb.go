// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/ingest/request/chef.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Run struct {
	// 1 through 15 are for frequently occuring fields
	// Reserving for shared fields between run_start and run_converge mesages.
	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RunId            string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	EntityUuid       string `protobuf:"bytes,3,opt,name=entity_uuid,json=entityUuid,proto3" json:"entity_uuid,omitempty"`
	MessageVersion   string `protobuf:"bytes,4,opt,name=message_version,json=messageVersion,proto3" json:"message_version,omitempty"`
	MessageType      string `protobuf:"bytes,5,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	NodeName         string `protobuf:"bytes,6,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	OrganizationName string `protobuf:"bytes,7,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	StartTime        string `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	ChefServerFqdn   string `protobuf:"bytes,9,opt,name=chef_server_fqdn,json=chefServerFqdn,proto3" json:"chef_server_fqdn,omitempty"`
	// This new field called 'content' is being used to send the entire raw JSON
	// message in bytes, this field is heavily used by the gateway for the DataCollector
	// Func Handler that will send the Run message to the (receiver) ingest-service
	// that will manually unmarshal the message from this field if it is provided.
	// The main purpose of this field it to improve the performance of ingestion when
	// the requests comes in REST/HTTP format.
	Content              []byte           `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty"`
	EndTime              string           `protobuf:"bytes,16,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Status               string           `protobuf:"bytes,17,opt,name=status,proto3" json:"status,omitempty"`
	TotalResourceCount   int32            `protobuf:"varint,18,opt,name=total_resource_count,json=totalResourceCount,proto3" json:"total_resource_count,omitempty"`
	UpdatedResourceCount int32            `protobuf:"varint,19,opt,name=updated_resource_count,json=updatedResourceCount,proto3" json:"updated_resource_count,omitempty"`
	Source               string           `protobuf:"bytes,20,opt,name=source,proto3" json:"source,omitempty"`
	ExpandedRunList      *ExpandedRunList `protobuf:"bytes,21,opt,name=expanded_run_list,json=expandedRunList,proto3" json:"expanded_run_list,omitempty"`
	Resources            []*Resource      `protobuf:"bytes,22,rep,name=resources,proto3" json:"resources,omitempty"`
	RunList              []string         `protobuf:"bytes,23,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	Node                 *_struct.Struct  `protobuf:"bytes,24,opt,name=node,proto3" json:"node,omitempty"`
	Error                *Error           `protobuf:"bytes,25,opt,name=error,proto3" json:"error,omitempty"`
	PolicyName           string           `protobuf:"bytes,26,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyGroup          string           `protobuf:"bytes,27,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	Deprecations         []*Deprecation   `protobuf:"bytes,28,rep,name=deprecations,proto3" json:"deprecations,omitempty"`
	Tags                 []string         `protobuf:"bytes,29,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Run) Reset()         { *m = Run{} }
func (m *Run) String() string { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()    {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_3065b16c78965565, []int{0}
}

func (m *Run) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Run.Unmarshal(m, b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Run.Marshal(b, m, deterministic)
}
func (m *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(m, src)
}
func (m *Run) XXX_Size() int {
	return xxx_messageInfo_Run.Size(m)
}
func (m *Run) XXX_DiscardUnknown() {
	xxx_messageInfo_Run.DiscardUnknown(m)
}

var xxx_messageInfo_Run proto.InternalMessageInfo

func (m *Run) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Run) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *Run) GetEntityUuid() string {
	if m != nil {
		return m.EntityUuid
	}
	return ""
}

func (m *Run) GetMessageVersion() string {
	if m != nil {
		return m.MessageVersion
	}
	return ""
}

func (m *Run) GetMessageType() string {
	if m != nil {
		return m.MessageType
	}
	return ""
}

func (m *Run) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Run) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *Run) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Run) GetChefServerFqdn() string {
	if m != nil {
		return m.ChefServerFqdn
	}
	return ""
}

func (m *Run) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Run) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Run) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Run) GetTotalResourceCount() int32 {
	if m != nil {
		return m.TotalResourceCount
	}
	return 0
}

func (m *Run) GetUpdatedResourceCount() int32 {
	if m != nil {
		return m.UpdatedResourceCount
	}
	return 0
}

func (m *Run) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Run) GetExpandedRunList() *ExpandedRunList {
	if m != nil {
		return m.ExpandedRunList
	}
	return nil
}

func (m *Run) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Run) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func (m *Run) GetNode() *_struct.Struct {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Run) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Run) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *Run) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *Run) GetDeprecations() []*Deprecation {
	if m != nil {
		return m.Deprecations
	}
	return nil
}

func (m *Run) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Deprecation struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Location             string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deprecation) Reset()         { *m = Deprecation{} }
func (m *Deprecation) String() string { return proto.CompactTextString(m) }
func (*Deprecation) ProtoMessage()    {}
func (*Deprecation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3065b16c78965565, []int{1}
}

func (m *Deprecation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deprecation.Unmarshal(m, b)
}
func (m *Deprecation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deprecation.Marshal(b, m, deterministic)
}
func (m *Deprecation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deprecation.Merge(m, src)
}
func (m *Deprecation) XXX_Size() int {
	return xxx_messageInfo_Deprecation.Size(m)
}
func (m *Deprecation) XXX_DiscardUnknown() {
	xxx_messageInfo_Deprecation.DiscardUnknown(m)
}

var xxx_messageInfo_Deprecation proto.InternalMessageInfo

func (m *Deprecation) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Deprecation) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Deprecation) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type ExpandedRunList struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RunList              []*RunList `protobuf:"bytes,2,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ExpandedRunList) Reset()         { *m = ExpandedRunList{} }
func (m *ExpandedRunList) String() string { return proto.CompactTextString(m) }
func (*ExpandedRunList) ProtoMessage()    {}
func (*ExpandedRunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3065b16c78965565, []int{2}
}

func (m *ExpandedRunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpandedRunList.Unmarshal(m, b)
}
func (m *ExpandedRunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpandedRunList.Marshal(b, m, deterministic)
}
func (m *ExpandedRunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpandedRunList.Merge(m, src)
}
func (m *ExpandedRunList) XXX_Size() int {
	return xxx_messageInfo_ExpandedRunList.Size(m)
}
func (m *ExpandedRunList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpandedRunList.DiscardUnknown(m)
}

var xxx_messageInfo_ExpandedRunList proto.InternalMessageInfo

func (m *ExpandedRunList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExpandedRunList) GetRunList() []*RunList {
	if m != nil {
		return m.RunList
	}
	return nil
}

type RunList struct {
	Type                 string     `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string     `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Skipped              bool       `protobuf:"varint,4,opt,name=skipped,proto3" json:"skipped,omitempty"`
	Children             []*RunList `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RunList) Reset()         { *m = RunList{} }
func (m *RunList) String() string { return proto.CompactTextString(m) }
func (*RunList) ProtoMessage()    {}
func (*RunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3065b16c78965565, []int{3}
}

func (m *RunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunList.Unmarshal(m, b)
}
func (m *RunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunList.Marshal(b, m, deterministic)
}
func (m *RunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunList.Merge(m, src)
}
func (m *RunList) XXX_Size() int {
	return xxx_messageInfo_RunList.Size(m)
}
func (m *RunList) XXX_DiscardUnknown() {
	xxx_messageInfo_RunList.DiscardUnknown(m)
}

var xxx_messageInfo_RunList proto.InternalMessageInfo

func (m *RunList) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RunList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RunList) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RunList) GetSkipped() bool {
	if m != nil {
		return m.Skipped
	}
	return false
}

func (m *RunList) GetChildren() []*RunList {
	if m != nil {
		return m.Children
	}
	return nil
}

type Resource struct {
	Type                 string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string          `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	After                *_struct.Struct `protobuf:"bytes,4,opt,name=after,proto3" json:"after,omitempty"`
	Before               *_struct.Struct `protobuf:"bytes,5,opt,name=before,proto3" json:"before,omitempty"`
	Duration             string          `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Delta                string          `protobuf:"bytes,7,opt,name=delta,proto3" json:"delta,omitempty"`
	CookbookName         string          `protobuf:"bytes,8,opt,name=cookbook_name,json=cookbookName,proto3" json:"cookbook_name,omitempty"`
	CookbookVersion      string          `protobuf:"bytes,9,opt,name=cookbook_version,json=cookbookVersion,proto3" json:"cookbook_version,omitempty"`
	Status               string          `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	RecipeName           string          `protobuf:"bytes,12,opt,name=recipe_name,json=recipeName,proto3" json:"recipe_name,omitempty"`
	IgnoreFailure        *_struct.Value  `protobuf:"bytes,13,opt,name=ignore_failure,json=ignoreFailure,proto3" json:"ignore_failure,omitempty"`
	Conditional          string          `protobuf:"bytes,16,opt,name=conditional,proto3" json:"conditional,omitempty"`
	Result               string          `protobuf:"bytes,17,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_3065b16c78965565, []int{4}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Resource) GetAfter() *_struct.Struct {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *Resource) GetBefore() *_struct.Struct {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Resource) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Resource) GetDelta() string {
	if m != nil {
		return m.Delta
	}
	return ""
}

func (m *Resource) GetCookbookName() string {
	if m != nil {
		return m.CookbookName
	}
	return ""
}

func (m *Resource) GetCookbookVersion() string {
	if m != nil {
		return m.CookbookVersion
	}
	return ""
}

func (m *Resource) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Resource) GetRecipeName() string {
	if m != nil {
		return m.RecipeName
	}
	return ""
}

func (m *Resource) GetIgnoreFailure() *_struct.Value {
	if m != nil {
		return m.IgnoreFailure
	}
	return nil
}

func (m *Resource) GetConditional() string {
	if m != nil {
		return m.Conditional
	}
	return ""
}

func (m *Resource) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type Error struct {
	Class                string       `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	Message              string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Backtrace            []string     `protobuf:"bytes,16,rep,name=backtrace,proto3" json:"backtrace,omitempty"`
	Description          *Description `protobuf:"bytes,17,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_3065b16c78965565, []int{5}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetBacktrace() []string {
	if m != nil {
		return m.Backtrace
	}
	return nil
}

func (m *Error) GetDescription() *Description {
	if m != nil {
		return m.Description
	}
	return nil
}

type Description struct {
	Title                string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Sections             []*_struct.Struct `protobuf:"bytes,2,rep,name=sections,proto3" json:"sections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Description) Reset()         { *m = Description{} }
func (m *Description) String() string { return proto.CompactTextString(m) }
func (*Description) ProtoMessage()    {}
func (*Description) Descriptor() ([]byte, []int) {
	return fileDescriptor_3065b16c78965565, []int{6}
}

func (m *Description) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Description.Unmarshal(m, b)
}
func (m *Description) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Description.Marshal(b, m, deterministic)
}
func (m *Description) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Description.Merge(m, src)
}
func (m *Description) XXX_Size() int {
	return xxx_messageInfo_Description.Size(m)
}
func (m *Description) XXX_DiscardUnknown() {
	xxx_messageInfo_Description.DiscardUnknown(m)
}

var xxx_messageInfo_Description proto.InternalMessageInfo

func (m *Description) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Description) GetSections() []*_struct.Struct {
	if m != nil {
		return m.Sections
	}
	return nil
}

func init() {
	proto.RegisterType((*Run)(nil), "chef.automate.api.ingest.request.Run")
	proto.RegisterType((*Deprecation)(nil), "chef.automate.api.ingest.request.Deprecation")
	proto.RegisterType((*ExpandedRunList)(nil), "chef.automate.api.ingest.request.ExpandedRunList")
	proto.RegisterType((*RunList)(nil), "chef.automate.api.ingest.request.RunList")
	proto.RegisterType((*Resource)(nil), "chef.automate.api.ingest.request.Resource")
	proto.RegisterType((*Error)(nil), "chef.automate.api.ingest.request.Error")
	proto.RegisterType((*Description)(nil), "chef.automate.api.ingest.request.Description")
}

func init() {
	proto.RegisterFile("api/external/ingest/request/chef.proto", fileDescriptor_3065b16c78965565)
}

var fileDescriptor_3065b16c78965565 = []byte{
	// 984 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x4d, 0x6f, 0x1b, 0x37,
	0x10, 0x85, 0x2c, 0xcb, 0x96, 0x46, 0xb2, 0x2d, 0xb3, 0x8e, 0xc3, 0x38, 0x0e, 0xa2, 0xa8, 0x40,
	0xa3, 0x34, 0x88, 0xd4, 0x26, 0x41, 0x6f, 0xb9, 0xb4, 0x49, 0xfa, 0x81, 0xa2, 0x45, 0x37, 0x1f,
	0x28, 0x0a, 0x14, 0x0b, 0x6a, 0x77, 0x24, 0x13, 0x5e, 0x91, 0x6b, 0x2e, 0x19, 0xc4, 0xfd, 0x41,
	0x3d, 0xf5, 0xe7, 0xf4, 0xd8, 0x1f, 0x53, 0x70, 0xc8, 0x95, 0xa5, 0x04, 0xb5, 0xdd, 0x93, 0x77,
	0x1e, 0xdf, 0x90, 0x43, 0xce, 0x9b, 0x67, 0xc1, 0x67, 0xa2, 0x94, 0x13, 0x7c, 0x6f, 0xd1, 0x28,
	0x51, 0x4c, 0xa4, 0x9a, 0x63, 0x65, 0x27, 0x06, 0xcf, 0x9c, 0xff, 0x9b, 0x9d, 0xe0, 0x6c, 0x5c,
	0x1a, 0x6d, 0x35, 0x1b, 0xd0, 0xb7, 0x70, 0x56, 0x2f, 0x84, 0xc5, 0xb1, 0x28, 0xe5, 0x38, 0x90,
	0xc7, 0x91, 0x7c, 0x74, 0x3c, 0xd7, 0x7a, 0x5e, 0xe0, 0x84, 0xf8, 0x53, 0x37, 0x9b, 0x54, 0xd6,
	0xb8, 0xcc, 0x86, 0xfc, 0xe1, 0xdf, 0xdb, 0xd0, 0x4c, 0x9c, 0x62, 0xbb, 0xb0, 0x21, 0x73, 0xde,
	0x18, 0x34, 0x46, 0x9d, 0x64, 0x43, 0xe6, 0xec, 0x06, 0x6c, 0x19, 0xa7, 0x52, 0x99, 0xf3, 0x0d,
	0xc2, 0x5a, 0xc6, 0xa9, 0xef, 0x73, 0x76, 0x17, 0xba, 0xa8, 0xac, 0xb4, 0xe7, 0xa9, 0x73, 0x32,
	0xe7, 0x4d, 0x5a, 0x83, 0x00, 0xbd, 0x71, 0x32, 0x67, 0xf7, 0x61, 0x6f, 0x81, 0x55, 0x25, 0xe6,
	0x98, 0xbe, 0x43, 0x53, 0x49, 0xad, 0xf8, 0x26, 0x91, 0x76, 0x23, 0xfc, 0x36, 0xa0, 0xec, 0x1e,
	0xf4, 0x6a, 0xa2, 0x3d, 0x2f, 0x91, 0xb7, 0x88, 0xd5, 0x8d, 0xd8, 0xeb, 0xf3, 0x12, 0xd9, 0x6d,
	0xe8, 0x28, 0x9d, 0x63, 0xaa, 0xc4, 0x02, 0xf9, 0x16, 0xad, 0xb7, 0x3d, 0xf0, 0x93, 0x58, 0x20,
	0x7b, 0x08, 0xfb, 0xda, 0xcc, 0x85, 0x92, 0x7f, 0x08, 0x2b, 0xb5, 0x0a, 0xa4, 0x6d, 0x22, 0xf5,
	0x57, 0x17, 0x88, 0x7c, 0x07, 0xa0, 0xb2, 0xc2, 0xd8, 0xd4, 0xca, 0x05, 0xf2, 0x36, 0xb1, 0x3a,
	0x84, 0xbc, 0x96, 0x0b, 0x64, 0x23, 0xe8, 0xfb, 0x67, 0x4c, 0x2b, 0x34, 0xef, 0xd0, 0xa4, 0xb3,
	0xb3, 0x5c, 0xf1, 0x4e, 0xa8, 0xda, 0xe3, 0xaf, 0x08, 0x7e, 0x79, 0x96, 0x2b, 0xc6, 0x61, 0x3b,
	0xd3, 0xca, 0xa2, 0xb2, 0x1c, 0x06, 0x8d, 0x51, 0x2f, 0xa9, 0x43, 0x76, 0x0b, 0xda, 0xa8, 0xf2,
	0x70, 0x40, 0x9f, 0x72, 0xb7, 0x51, 0xe5, 0xb4, 0xfd, 0x21, 0x6c, 0x55, 0x56, 0x58, 0x57, 0xf1,
	0x7d, 0x5a, 0x88, 0x11, 0xfb, 0x02, 0x0e, 0xac, 0xb6, 0xa2, 0x48, 0x0d, 0x56, 0xda, 0x99, 0x0c,
	0xd3, 0x4c, 0x3b, 0x65, 0x39, 0x1b, 0x34, 0x46, 0xad, 0x84, 0xd1, 0x5a, 0x12, 0x97, 0xbe, 0xf1,
	0x2b, 0xec, 0x29, 0x1c, 0xba, 0x32, 0x17, 0x16, 0xf3, 0x0f, 0x73, 0x3e, 0xa1, 0x9c, 0x83, 0xb8,
	0xba, 0x9e, 0xe5, 0xcf, 0xa7, 0x90, 0x1f, 0xc4, 0xf3, 0x29, 0x62, 0xbf, 0xc3, 0x3e, 0xbe, 0x2f,
	0x85, 0xca, 0xfd, 0x76, 0x4e, 0xa5, 0x85, 0xac, 0x2c, 0xbf, 0x31, 0x68, 0x8c, 0xba, 0x8f, 0xbf,
	0x1c, 0x5f, 0xa5, 0xab, 0xf1, 0x8b, 0x98, 0x9a, 0x38, 0xf5, 0xa3, 0xac, 0x6c, 0xb2, 0x87, 0xeb,
	0x00, 0xfb, 0x0e, 0x3a, 0x75, 0x91, 0x15, 0x3f, 0x1c, 0x34, 0x47, 0xdd, 0xc7, 0x9f, 0x5f, 0xbd,
	0x6d, 0x5d, 0x7a, 0x72, 0x91, 0xec, 0xdf, 0x76, 0x59, 0xdf, 0xcd, 0x41, 0xd3, 0xbf, 0xad, 0x89,
	0x87, 0x3c, 0x84, 0x4d, 0x2f, 0x09, 0xce, 0xa9, 0xec, 0x9b, 0xe3, 0x20, 0xf6, 0x71, 0x2d, 0xf6,
	0xf1, 0x2b, 0x12, 0x7b, 0x42, 0x24, 0xf6, 0x0c, 0x5a, 0x68, 0x8c, 0x36, 0xfc, 0x16, 0xb1, 0xef,
	0x5f, 0xe3, 0x92, 0x9e, 0x9e, 0x84, 0x2c, 0x2f, 0xfe, 0x52, 0x17, 0x32, 0x3b, 0x0f, 0x62, 0x3b,
	0x0a, 0xe2, 0x0f, 0x10, 0xc9, 0xec, 0x1e, 0xf4, 0x22, 0x61, 0x6e, 0xb4, 0x2b, 0xf9, 0xed, 0xa0,
	0xe9, 0x80, 0x7d, 0xeb, 0x21, 0xf6, 0x0b, 0xf4, 0x72, 0x2c, 0x0d, 0x66, 0x24, 0xce, 0x8a, 0x1f,
	0xd3, 0xbb, 0x3c, 0xba, 0xba, 0x92, 0xe7, 0x17, 0x59, 0xc9, 0xda, 0x16, 0x8c, 0xc1, 0xa6, 0x15,
	0xf3, 0x8a, 0xdf, 0xa1, 0x97, 0xa1, 0xef, 0xe1, 0x1b, 0xe8, 0xae, 0x24, 0x78, 0xd9, 0xc6, 0xc1,
	0x8a, 0x23, 0x5e, 0x87, 0xac, 0x0f, 0x4d, 0x67, 0x8a, 0x38, 0xe4, 0xfe, 0x93, 0x1d, 0x41, 0xbb,
	0xd0, 0x21, 0x2f, 0xce, 0xf7, 0x32, 0x1e, 0xce, 0x61, 0xef, 0x83, 0xb6, 0x7f, 0x64, 0x1c, 0xcf,
	0x57, 0x7a, 0xb5, 0x41, 0x97, 0x7b, 0x70, 0x8d, 0xa6, 0x47, 0x0d, 0xd5, 0x6d, 0x1d, 0xfe, 0xd5,
	0x80, 0xed, 0xfa, 0x04, 0x7f, 0x3f, 0xef, 0x10, 0xe1, 0x0c, 0xfa, 0xf6, 0x18, 0xf5, 0x20, 0xd4,
	0x4d, 0xdf, 0xfe, 0x92, 0xb5, 0xe5, 0x84, 0xba, 0xeb, 0xd0, 0xaf, 0x54, 0xa7, 0xb2, 0x2c, 0x31,
	0x27, 0x33, 0x6a, 0x27, 0x75, 0xc8, 0x5e, 0x40, 0x3b, 0x3b, 0x91, 0x45, 0x6e, 0x50, 0xf1, 0xd6,
	0xff, 0xad, 0x76, 0x99, 0x3a, 0xfc, 0xa7, 0x09, 0xed, 0x5a, 0xb8, 0xd7, 0xae, 0x37, 0xbc, 0x5c,
	0x73, 0xf9, 0x72, 0x8f, 0xa0, 0x25, 0x66, 0x16, 0x0d, 0xd5, 0x78, 0x89, 0x96, 0x03, 0x8b, 0x4d,
	0x60, 0x6b, 0x8a, 0x33, 0x6d, 0x82, 0x75, 0x5e, 0xc2, 0x8f, 0x34, 0xdf, 0xd8, 0xdc, 0x99, 0xd0,
	0xd8, 0xe8, 0xa6, 0x75, 0xcc, 0x0e, 0xa0, 0x95, 0x63, 0x61, 0x45, 0x74, 0xd0, 0x10, 0xb0, 0x4f,
	0x61, 0x27, 0xd3, 0xfa, 0x74, 0xaa, 0xf5, 0x69, 0x90, 0x7c, 0x70, 0xce, 0x5e, 0x0d, 0x92, 0xe8,
	0x1f, 0x40, 0x7f, 0x49, 0xaa, 0xdf, 0x3f, 0x98, 0xe7, 0x5e, 0x8d, 0xd7, 0x9e, 0x7f, 0x61, 0x84,
	0xdd, 0x35, 0x23, 0xbc, 0x0b, 0x5d, 0x83, 0x99, 0x2c, 0xa3, 0xd5, 0xf7, 0xc2, 0x60, 0x05, 0x88,
	0xce, 0x78, 0x06, 0xbb, 0x72, 0xae, 0xb4, 0xc1, 0x74, 0x26, 0x64, 0xe1, 0x0c, 0xf2, 0x1d, 0xba,
	0xf3, 0xe1, 0x47, 0x77, 0x7e, 0x2b, 0x0a, 0x87, 0xc9, 0x4e, 0x60, 0xbf, 0x0c, 0x64, 0x36, 0x80,
	0x6e, 0xa6, 0x55, 0x2e, 0xfd, 0x55, 0x45, 0x11, 0xed, 0x79, 0x15, 0xf2, 0x95, 0x19, 0xac, 0x5c,
	0x61, 0x6b, 0x8b, 0x0e, 0xd1, 0x0f, 0x9b, 0x6d, 0xe8, 0x77, 0x87, 0x7f, 0x36, 0xa0, 0x45, 0x4e,
	0xe0, 0xdf, 0x29, 0x2b, 0x44, 0x55, 0xc5, 0xe6, 0x86, 0x60, 0x75, 0xbc, 0x36, 0xd6, 0xc7, 0xeb,
	0x18, 0x3a, 0x53, 0x91, 0x9d, 0x5a, 0x23, 0x32, 0xff, 0x6f, 0xc1, 0x0f, 0xe8, 0x05, 0xc0, 0x7e,
	0x86, 0x6e, 0x8e, 0x55, 0x66, 0x64, 0x49, 0x4d, 0xd9, 0xa7, 0x3b, 0x5d, 0xcb, 0x0b, 0x96, 0x49,
	0xc9, 0xea, 0x0e, 0xc3, 0x5f, 0xfd, 0xd8, 0x2f, 0x43, 0x5f, 0xad, 0x95, 0xb6, 0xa8, 0xa5, 0x18,
	0x02, 0xf6, 0x04, 0xda, 0x15, 0x66, 0xc1, 0x7e, 0xc2, 0x84, 0xfe, 0xa7, 0x74, 0x96, 0xc4, 0xaf,
	0xbf, 0xfa, 0xed, 0xe9, 0x5c, 0xda, 0x13, 0x37, 0x1d, 0x67, 0x7a, 0x41, 0x3f, 0x40, 0x26, 0x75,
	0x85, 0x93, 0x4b, 0x7e, 0xaa, 0x4c, 0xb7, 0x68, 0xcb, 0x27, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x8c, 0x23, 0x87, 0xfd, 0xd0, 0x08, 0x00, 0x00,
}
