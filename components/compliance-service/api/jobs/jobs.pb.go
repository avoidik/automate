// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/compliance-service/api/jobs/jobs.proto

package jobs

import (
	context "context"
	fmt "fmt"
	common "github.com/chef/automate/components/compliance-service/api/common"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Query_OrderType int32

const (
	Query_ASC  Query_OrderType = 0
	Query_DESC Query_OrderType = 1
)

var Query_OrderType_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}

var Query_OrderType_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x Query_OrderType) String() string {
	return proto.EnumName(Query_OrderType_name, int32(x))
}

func (Query_OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{6, 0}
}

type GetJobResultByNodeIdRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	NodeId               string   `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetJobResultByNodeIdRequest) Reset()         { *m = GetJobResultByNodeIdRequest{} }
func (m *GetJobResultByNodeIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetJobResultByNodeIdRequest) ProtoMessage()    {}
func (*GetJobResultByNodeIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{0}
}

func (m *GetJobResultByNodeIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetJobResultByNodeIdRequest.Unmarshal(m, b)
}
func (m *GetJobResultByNodeIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetJobResultByNodeIdRequest.Marshal(b, m, deterministic)
}
func (m *GetJobResultByNodeIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetJobResultByNodeIdRequest.Merge(m, src)
}
func (m *GetJobResultByNodeIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetJobResultByNodeIdRequest.Size(m)
}
func (m *GetJobResultByNodeIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetJobResultByNodeIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetJobResultByNodeIdRequest proto.InternalMessageInfo

func (m *GetJobResultByNodeIdRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *GetJobResultByNodeIdRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type RerunResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RerunResponse) Reset()         { *m = RerunResponse{} }
func (m *RerunResponse) String() string { return proto.CompactTextString(m) }
func (*RerunResponse) ProtoMessage()    {}
func (*RerunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{1}
}

func (m *RerunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RerunResponse.Unmarshal(m, b)
}
func (m *RerunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RerunResponse.Marshal(b, m, deterministic)
}
func (m *RerunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RerunResponse.Merge(m, src)
}
func (m *RerunResponse) XXX_Size() int {
	return xxx_messageInfo_RerunResponse.Size(m)
}
func (m *RerunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RerunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RerunResponse proto.InternalMessageInfo

type Id struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{2}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Ids struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ids) Reset()         { *m = Ids{} }
func (m *Ids) String() string { return proto.CompactTextString(m) }
func (*Ids) ProtoMessage()    {}
func (*Ids) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{3}
}

func (m *Ids) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ids.Unmarshal(m, b)
}
func (m *Ids) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ids.Marshal(b, m, deterministic)
}
func (m *Ids) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ids.Merge(m, src)
}
func (m *Ids) XXX_Size() int {
	return xxx_messageInfo_Ids.Size(m)
}
func (m *Ids) XXX_DiscardUnknown() {
	xxx_messageInfo_Ids.DiscardUnknown(m)
}

var xxx_messageInfo_Ids proto.InternalMessageInfo

func (m *Ids) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type TimeQuery struct {
	StartTime            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeQuery) Reset()         { *m = TimeQuery{} }
func (m *TimeQuery) String() string { return proto.CompactTextString(m) }
func (*TimeQuery) ProtoMessage()    {}
func (*TimeQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{4}
}

func (m *TimeQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeQuery.Unmarshal(m, b)
}
func (m *TimeQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeQuery.Marshal(b, m, deterministic)
}
func (m *TimeQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeQuery.Merge(m, src)
}
func (m *TimeQuery) XXX_Size() int {
	return xxx_messageInfo_TimeQuery.Size(m)
}
func (m *TimeQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeQuery.DiscardUnknown(m)
}

var xxx_messageInfo_TimeQuery proto.InternalMessageInfo

func (m *TimeQuery) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

type ManagerFilter struct {
	ManagerId            string           `protobuf:"bytes,1,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	Filters              []*common.Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ManagerFilter) Reset()         { *m = ManagerFilter{} }
func (m *ManagerFilter) String() string { return proto.CompactTextString(m) }
func (*ManagerFilter) ProtoMessage()    {}
func (*ManagerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{5}
}

func (m *ManagerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManagerFilter.Unmarshal(m, b)
}
func (m *ManagerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManagerFilter.Marshal(b, m, deterministic)
}
func (m *ManagerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManagerFilter.Merge(m, src)
}
func (m *ManagerFilter) XXX_Size() int {
	return xxx_messageInfo_ManagerFilter.Size(m)
}
func (m *ManagerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ManagerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ManagerFilter proto.InternalMessageInfo

func (m *ManagerFilter) GetManagerId() string {
	if m != nil {
		return m.ManagerId
	}
	return ""
}

func (m *ManagerFilter) GetFilters() []*common.Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

type Query struct {
	Filters              []*common.Filter `protobuf:"bytes,20,rep,name=filters,proto3" json:"filters,omitempty"`
	Order                Query_OrderType  `protobuf:"varint,21,opt,name=order,proto3,enum=chef.automate.domain.compliance.api.jobs.Query_OrderType" json:"order,omitempty"`
	Sort                 string           `protobuf:"bytes,22,opt,name=sort,proto3" json:"sort,omitempty"`
	Page                 int32            `protobuf:"varint,23,opt,name=page,proto3" json:"page,omitempty"`
	PerPage              int32            `protobuf:"varint,24,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{6}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetFilters() []*common.Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *Query) GetOrder() Query_OrderType {
	if m != nil {
		return m.Order
	}
	return Query_ASC
}

func (m *Query) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *Query) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Query) GetPerPage() int32 {
	if m != nil {
		return m.PerPage
	}
	return 0
}

type Job struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string               `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Timeout              int32                `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Tags                 []*common.Kv         `protobuf:"bytes,20,rep,name=tags,proto3" json:"tags,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,21,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,22,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Status               string               `protobuf:"bytes,23,opt,name=status,proto3" json:"status,omitempty"`
	Retries              int32                `protobuf:"varint,26,opt,name=retries,proto3" json:"retries,omitempty"`
	RetriesLeft          int32                `protobuf:"varint,27,opt,name=retries_left,json=retriesLeft,proto3" json:"retries_left,omitempty"`
	Results              []*ResultsRow        `protobuf:"bytes,28,rep,name=results,proto3" json:"results,omitempty"`
	Nodes                []string             `protobuf:"bytes,100,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Profiles             []string             `protobuf:"bytes,101,rep,name=profiles,proto3" json:"profiles,omitempty"`
	NodeCount            int32                `protobuf:"varint,102,opt,name=node_count,json=nodeCount,proto3" json:"node_count,omitempty"`
	ProfileCount         int32                `protobuf:"varint,103,opt,name=profile_count,json=profileCount,proto3" json:"profile_count,omitempty"`
	NodeSelectors        []*ManagerFilter     `protobuf:"bytes,104,rep,name=node_selectors,json=nodeSelectors,proto3" json:"node_selectors,omitempty"`
	ScheduledTime        *timestamp.Timestamp `protobuf:"bytes,105,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
	Recurrence           string               `protobuf:"bytes,106,opt,name=recurrence,proto3" json:"recurrence,omitempty"`
	ParentId             string               `protobuf:"bytes,107,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	JobCount             int32                `protobuf:"varint,108,opt,name=job_count,json=jobCount,proto3" json:"job_count,omitempty"`
	Deleted              bool                 `protobuf:"varint,109,opt,name=deleted,proto3" json:"deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{7}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Job) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Job) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Job) GetTags() []*common.Kv {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Job) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Job) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Job) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Job) GetRetries() int32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

func (m *Job) GetRetriesLeft() int32 {
	if m != nil {
		return m.RetriesLeft
	}
	return 0
}

func (m *Job) GetResults() []*ResultsRow {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *Job) GetNodes() []string {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *Job) GetProfiles() []string {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Job) GetNodeCount() int32 {
	if m != nil {
		return m.NodeCount
	}
	return 0
}

func (m *Job) GetProfileCount() int32 {
	if m != nil {
		return m.ProfileCount
	}
	return 0
}

func (m *Job) GetNodeSelectors() []*ManagerFilter {
	if m != nil {
		return m.NodeSelectors
	}
	return nil
}

func (m *Job) GetScheduledTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduledTime
	}
	return nil
}

func (m *Job) GetRecurrence() string {
	if m != nil {
		return m.Recurrence
	}
	return ""
}

func (m *Job) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *Job) GetJobCount() int32 {
	if m != nil {
		return m.JobCount
	}
	return 0
}

func (m *Job) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type Jobs struct {
	Jobs                 []*Job   `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	Total                int32    `protobuf:"varint,20,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Jobs) Reset()         { *m = Jobs{} }
func (m *Jobs) String() string { return proto.CompactTextString(m) }
func (*Jobs) ProtoMessage()    {}
func (*Jobs) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{8}
}

func (m *Jobs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Jobs.Unmarshal(m, b)
}
func (m *Jobs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Jobs.Marshal(b, m, deterministic)
}
func (m *Jobs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Jobs.Merge(m, src)
}
func (m *Jobs) XXX_Size() int {
	return xxx_messageInfo_Jobs.Size(m)
}
func (m *Jobs) XXX_DiscardUnknown() {
	xxx_messageInfo_Jobs.DiscardUnknown(m)
}

var xxx_messageInfo_Jobs proto.InternalMessageInfo

func (m *Jobs) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *Jobs) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ResultsRow struct {
	NodeId               string               `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ReportId             string               `protobuf:"bytes,2,opt,name=report_id,json=reportId,proto3" json:"report_id,omitempty"`
	Status               string               `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Result               string               `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	JobId                string               `protobuf:"bytes,5,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,20,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,21,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ResultsRow) Reset()         { *m = ResultsRow{} }
func (m *ResultsRow) String() string { return proto.CompactTextString(m) }
func (*ResultsRow) ProtoMessage()    {}
func (*ResultsRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_2efaf54b9c1cd8a6, []int{9}
}

func (m *ResultsRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultsRow.Unmarshal(m, b)
}
func (m *ResultsRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultsRow.Marshal(b, m, deterministic)
}
func (m *ResultsRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultsRow.Merge(m, src)
}
func (m *ResultsRow) XXX_Size() int {
	return xxx_messageInfo_ResultsRow.Size(m)
}
func (m *ResultsRow) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultsRow.DiscardUnknown(m)
}

var xxx_messageInfo_ResultsRow proto.InternalMessageInfo

func (m *ResultsRow) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ResultsRow) GetReportId() string {
	if m != nil {
		return m.ReportId
	}
	return ""
}

func (m *ResultsRow) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ResultsRow) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *ResultsRow) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *ResultsRow) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ResultsRow) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("chef.automate.domain.compliance.api.jobs.Query_OrderType", Query_OrderType_name, Query_OrderType_value)
	proto.RegisterType((*GetJobResultByNodeIdRequest)(nil), "chef.automate.domain.compliance.api.jobs.GetJobResultByNodeIdRequest")
	proto.RegisterType((*RerunResponse)(nil), "chef.automate.domain.compliance.api.jobs.RerunResponse")
	proto.RegisterType((*Id)(nil), "chef.automate.domain.compliance.api.jobs.Id")
	proto.RegisterType((*Ids)(nil), "chef.automate.domain.compliance.api.jobs.Ids")
	proto.RegisterType((*TimeQuery)(nil), "chef.automate.domain.compliance.api.jobs.TimeQuery")
	proto.RegisterType((*ManagerFilter)(nil), "chef.automate.domain.compliance.api.jobs.ManagerFilter")
	proto.RegisterType((*Query)(nil), "chef.automate.domain.compliance.api.jobs.Query")
	proto.RegisterType((*Job)(nil), "chef.automate.domain.compliance.api.jobs.Job")
	proto.RegisterType((*Jobs)(nil), "chef.automate.domain.compliance.api.jobs.Jobs")
	proto.RegisterType((*ResultsRow)(nil), "chef.automate.domain.compliance.api.jobs.ResultsRow")
}

func init() {
	proto.RegisterFile("components/compliance-service/api/jobs/jobs.proto", fileDescriptor_2efaf54b9c1cd8a6)
}

var fileDescriptor_2efaf54b9c1cd8a6 = []byte{
	// 1046 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4b, 0x73, 0x23, 0x35,
	0x10, 0xf6, 0xfb, 0xd1, 0xd9, 0x84, 0x94, 0xca, 0x49, 0x06, 0x67, 0x59, 0xcc, 0x70, 0xf1, 0x81,
	0x1d, 0x17, 0x5e, 0xb6, 0xb6, 0xf6, 0x44, 0x25, 0xd9, 0x2c, 0xe5, 0x90, 0x7d, 0x4d, 0xc2, 0x85,
	0x03, 0x29, 0x79, 0xd4, 0x76, 0x26, 0xcc, 0x48, 0x83, 0xa4, 0x59, 0x2a, 0x14, 0x7f, 0x80, 0x1f,
	0xc0, 0xef, 0xe1, 0x47, 0x71, 0xe4, 0x42, 0x49, 0x1a, 0x27, 0x4e, 0x78, 0xc4, 0x93, 0xbd, 0xd8,
	0xea, 0x9e, 0xee, 0x4f, 0xea, 0xaf, 0x1f, 0x12, 0x7c, 0x19, 0x89, 0x34, 0x13, 0x1c, 0xb9, 0x56,
	0x23, 0xb3, 0x4c, 0x62, 0xca, 0x23, 0x7c, 0xac, 0x50, 0xbe, 0x8f, 0x23, 0x1c, 0xd1, 0x2c, 0x1e,
	0x5d, 0x88, 0xa9, 0xb2, 0x3f, 0x41, 0x26, 0x85, 0x16, 0x64, 0x18, 0x9d, 0xe3, 0x2c, 0xa0, 0xb9,
	0x16, 0x29, 0xd5, 0x18, 0x30, 0x91, 0xd2, 0x98, 0x07, 0xd7, 0xce, 0x01, 0xcd, 0xe2, 0xc0, 0xd8,
	0xf7, 0x3f, 0x9d, 0x0b, 0x31, 0x4f, 0x70, 0x64, 0xfd, 0xa6, 0xf9, 0x6c, 0xa4, 0xe3, 0x14, 0x95,
	0xa6, 0x69, 0xe6, 0xa0, 0xfa, 0x0f, 0x0b, 0x03, 0xb3, 0x0d, 0xe5, 0x5c, 0x68, 0xaa, 0x63, 0xc1,
	0x8b, 0x8d, 0xfa, 0xbb, 0xb7, 0xdd, 0x31, 0xcd, 0xf4, 0x65, 0xf1, 0xf1, 0xe9, 0xdd, 0x07, 0x8f,
	0x44, 0x9a, 0x0a, 0x5e, 0xfc, 0x39, 0x37, 0xff, 0x15, 0xec, 0x7e, 0x83, 0xfa, 0x48, 0x4c, 0x43,
	0x54, 0x79, 0xa2, 0xf7, 0x2f, 0x5f, 0x0b, 0x86, 0x13, 0x16, 0xe2, 0x4f, 0x39, 0x2a, 0x4d, 0xb6,
	0xa0, 0x75, 0x21, 0xa6, 0x67, 0x31, 0xf3, 0xaa, 0x83, 0xea, 0xb0, 0x1b, 0x36, 0x2f, 0xc4, 0x74,
	0xc2, 0xc8, 0x0e, 0xb4, 0xb9, 0x60, 0x68, 0xf4, 0x35, 0xab, 0x6f, 0x71, 0xeb, 0xe6, 0x7f, 0x04,
	0xeb, 0x21, 0xca, 0x9c, 0x87, 0xa8, 0x32, 0xc1, 0x15, 0xfa, 0x3d, 0xa8, 0x4d, 0x18, 0xd9, 0x80,
	0xda, 0x15, 0x44, 0x2d, 0x66, 0xfe, 0x0e, 0xd4, 0x27, 0x4c, 0x91, 0x4d, 0xa8, 0xc7, 0x4c, 0x79,
	0xd5, 0x41, 0x7d, 0xd8, 0x0d, 0xcd, 0xd2, 0x7f, 0x09, 0xdd, 0xd3, 0x38, 0xc5, 0x77, 0x39, 0xca,
	0x4b, 0xf2, 0x1c, 0x40, 0x69, 0x2a, 0xf5, 0x99, 0xa1, 0xc9, 0x7a, 0xaf, 0x8d, 0xfb, 0x81, 0x23,
	0x21, 0x58, 0x90, 0x10, 0x9c, 0x2e, 0x38, 0x0c, 0xbb, 0xd6, 0xda, 0xc8, 0xfe, 0xaf, 0xb0, 0xfe,
	0x8a, 0x72, 0x3a, 0x47, 0xf9, 0x32, 0x4e, 0x34, 0x4a, 0xf2, 0x09, 0x40, 0xea, 0x14, 0xd7, 0xc1,
	0x74, 0x0b, 0xcd, 0x84, 0x91, 0x63, 0x68, 0xcf, 0xac, 0xa1, 0xf2, 0x6a, 0x83, 0xfa, 0x70, 0x6d,
	0x3c, 0x0e, 0x56, 0xc9, 0x6a, 0x41, 0xa5, 0xdb, 0x23, 0x5c, 0x40, 0xf8, 0xbf, 0xd5, 0xa0, 0xe9,
	0x42, 0x58, 0xc2, 0xed, 0x7d, 0x30, 0x2e, 0x79, 0x03, 0x4d, 0x21, 0x19, 0x4a, 0x6f, 0x6b, 0x50,
	0x1d, 0x6e, 0x8c, 0x9f, 0x07, 0xab, 0x56, 0x5e, 0x60, 0x4f, 0x13, 0xbc, 0x31, 0xce, 0xa7, 0x97,
	0x19, 0x86, 0x0e, 0x87, 0x10, 0x68, 0x28, 0x21, 0xb5, 0xb7, 0x6d, 0xf9, 0xb0, 0x6b, 0xa3, 0xcb,
	0xe8, 0x1c, 0xbd, 0x9d, 0x41, 0x75, 0xd8, 0x0c, 0xed, 0x9a, 0x7c, 0x0c, 0x9d, 0x0c, 0xe5, 0x99,
	0xd5, 0x7b, 0x56, 0xdf, 0xce, 0x50, 0xbe, 0xa5, 0x73, 0xf4, 0x1f, 0x41, 0xf7, 0x0a, 0x96, 0xb4,
	0xa1, 0xbe, 0x77, 0x72, 0xb0, 0x59, 0x21, 0x1d, 0x68, 0xbc, 0x38, 0x3c, 0x39, 0xd8, 0xac, 0xfa,
	0x7f, 0xb4, 0xa0, 0x7e, 0x24, 0xa6, 0xb7, 0x4b, 0xc0, 0x6c, 0xc3, 0x69, 0x8a, 0x45, 0xfd, 0xd8,
	0xb5, 0xd1, 0xe9, 0xcb, 0x0c, 0xbd, 0xba, 0xd3, 0x99, 0x35, 0xf1, 0xa0, 0x6d, 0xd2, 0x2f, 0x72,
	0xed, 0x35, 0xdc, 0xce, 0x85, 0x48, 0xf6, 0xa1, 0xa1, 0xe9, 0x7c, 0x41, 0x6c, 0x50, 0x86, 0xd8,
	0x6f, 0xdf, 0x87, 0xd6, 0xf7, 0x56, 0x89, 0x6d, 0x95, 0x28, 0x31, 0xf2, 0x14, 0x3a, 0xc8, 0x99,
	0x73, 0xdc, 0xbe, 0xd3, 0xb1, 0x8d, 0x9c, 0x59, 0xb7, 0x6d, 0x68, 0x29, 0x4d, 0x75, 0xae, 0x2c,
	0xc1, 0xdd, 0xb0, 0x90, 0x4c, 0x9c, 0x12, 0xb5, 0x8c, 0x51, 0x79, 0x7d, 0x17, 0x67, 0x21, 0x92,
	0xcf, 0xe0, 0x41, 0xb1, 0x3c, 0x4b, 0x70, 0xa6, 0xbd, 0x5d, 0xfb, 0x79, 0xad, 0xd0, 0x1d, 0xe3,
	0x4c, 0x93, 0xd7, 0xc6, 0xd9, 0xf4, 0xaf, 0xf2, 0x1e, 0x5a, 0x36, 0xbe, 0x5a, 0xbd, 0x34, 0x5c,
	0xe3, 0xab, 0x50, 0xfc, 0x1c, 0x2e, 0x40, 0x48, 0x0f, 0x9a, 0xa6, 0xa1, 0x95, 0xc7, 0x6c, 0x6b,
	0x3a, 0x81, 0xf4, 0xa1, 0x93, 0x49, 0x31, 0x8b, 0x13, 0x54, 0x1e, 0xda, 0x0f, 0x57, 0xb2, 0xe9,
	0x2f, 0x3b, 0x11, 0x22, 0x91, 0x73, 0xed, 0xcd, 0xec, 0x11, 0xbb, 0x46, 0x73, 0x60, 0x14, 0xe4,
	0x73, 0x58, 0x2f, 0x4c, 0x0b, 0x8b, 0xb9, 0xb5, 0x78, 0x50, 0x28, 0x9d, 0xd1, 0x0f, 0xb0, 0x61,
	0x31, 0x14, 0x26, 0x18, 0x69, 0x21, 0x95, 0x77, 0x6e, 0x83, 0x79, 0xb6, 0x7a, 0x30, 0x37, 0x9a,
	0x3e, 0x5c, 0x37, 0x70, 0x27, 0x0b, 0x34, 0xb2, 0x07, 0x1b, 0x2a, 0x3a, 0x47, 0x96, 0x27, 0x58,
	0xe4, 0x2d, 0xbe, 0x33, 0x6f, 0xeb, 0x57, 0x1e, 0x36, 0x7b, 0x8f, 0x00, 0x24, 0x46, 0xb9, 0x94,
	0xc8, 0x23, 0xf4, 0x2e, 0x6c, 0x06, 0x97, 0x34, 0x64, 0x17, 0xba, 0x19, 0x95, 0xc8, 0xb5, 0x99,
	0x32, 0x3f, 0xda, 0xcf, 0x1d, 0xa7, 0x98, 0x30, 0xf3, 0xd1, 0x0c, 0x53, 0x47, 0x40, 0x62, 0x09,
	0xe8, 0x5c, 0x88, 0xa9, 0x0b, 0xde, 0x83, 0x36, 0xc3, 0x04, 0x35, 0x32, 0x2f, 0x1d, 0x54, 0x87,
	0x9d, 0x70, 0x21, 0xfa, 0x67, 0xd0, 0x38, 0x12, 0x53, 0x73, 0xfc, 0x86, 0x89, 0xd1, 0x8e, 0xcb,
	0xb5, 0xf1, 0xe3, 0xd5, 0x49, 0x31, 0xd3, 0xdd, 0xba, 0x9a, 0xbc, 0x6a, 0xa1, 0x69, 0xe2, 0xf5,
	0xec, 0xee, 0x4e, 0xf0, 0xff, 0xaa, 0x02, 0x5c, 0x57, 0xc1, 0xf2, 0x70, 0xaf, 0x2e, 0x0f, 0x77,
	0x73, 0x7e, 0x89, 0x99, 0x90, 0xfa, 0x7a, 0xee, 0x77, 0x9c, 0x62, 0xc2, 0x96, 0xea, 0xba, 0x7e,
	0xa3, 0xae, 0xb7, 0xa1, 0xe5, 0xaa, 0xca, 0xb6, 0x6f, 0x37, 0x2c, 0xa4, 0xa5, 0x9b, 0xa5, 0xb9,
	0x7c, 0xb3, 0xdc, 0x6c, 0xc8, 0xde, 0x7d, 0x1b, 0x72, 0x6b, 0xe5, 0x86, 0x1c, 0xff, 0xd9, 0x82,
	0x35, 0xc3, 0xef, 0x89, 0xbb, 0x29, 0xc9, 0x1c, 0x5a, 0x07, 0x12, 0xa9, 0x46, 0x52, 0x8e, 0xe2,
	0xfe, 0x17, 0xab, 0x9b, 0x4f, 0x98, 0x5f, 0x21, 0x08, 0x8d, 0x10, 0x29, 0x23, 0xa5, 0xfc, 0xfa,
	0xe5, 0x0e, 0xe5, 0x57, 0xc8, 0x3b, 0x68, 0x7d, 0x97, 0xb1, 0x7b, 0xc4, 0xb3, 0xfd, 0x0f, 0xf6,
	0x0e, 0xcd, 0x7b, 0xc3, 0xaf, 0x90, 0xb7, 0xd0, 0x7a, 0x61, 0x8b, 0xb3, 0xe4, 0xd9, 0xff, 0x1b,
	0x31, 0x86, 0xc6, 0x71, 0xac, 0x34, 0x19, 0x95, 0xbc, 0xd2, 0xfa, 0x41, 0xa9, 0x98, 0x94, 0x5f,
	0x21, 0x19, 0x34, 0xed, 0x13, 0xa5, 0xe4, 0xd9, 0x9f, 0x95, 0x99, 0xa8, 0xcb, 0x2f, 0xa0, 0x0a,
	0xf9, 0x05, 0x88, 0x09, 0x6e, 0xc2, 0x63, 0x1d, 0x53, 0x8d, 0xec, 0x24, 0xa2, 0x5c, 0x91, 0x27,
	0xab, 0x03, 0x5e, 0x3d, 0x89, 0xca, 0x64, 0x7f, 0xc2, 0x4c, 0xb4, 0xbf, 0x57, 0xa1, 0xf7, 0x6f,
	0x0f, 0x3c, 0x72, 0xb8, 0x3a, 0xd2, 0xff, 0x3c, 0x10, 0xfb, 0xf7, 0xba, 0x68, 0xfc, 0xca, 0xfe,
	0xde, 0xf7, 0x5f, 0xcf, 0x63, 0x7d, 0x9e, 0x4f, 0x8d, 0xe9, 0xc8, 0x60, 0x8c, 0x16, 0x18, 0xa3,
	0xd5, 0x9e, 0xe0, 0xd3, 0x96, 0xad, 0xa2, 0x27, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x3c,
	0x84, 0x76, 0xb3, 0x0b, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobsServiceClient is the client API for JobsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobsServiceClient interface {
	Create(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Id, error)
	Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Job, error)
	Update(ctx context.Context, in *Job, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Jobs, error)
	Rerun(ctx context.Context, in *Id, opts ...grpc.CallOption) (*RerunResponse, error)
	ListInitiatedScans(ctx context.Context, in *TimeQuery, opts ...grpc.CallOption) (*Ids, error)
	GetJobResultByNodeId(ctx context.Context, in *GetJobResultByNodeIdRequest, opts ...grpc.CallOption) (*ResultsRow, error)
}

type jobsServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobsServiceClient(cc *grpc.ClientConn) JobsServiceClient {
	return &jobsServiceClient{cc}
}

func (c *jobsServiceClient) Create(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.api.jobs.JobsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.api.jobs.JobsService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) Update(ctx context.Context, in *Job, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.api.jobs.JobsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.api.jobs.JobsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) List(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Jobs, error) {
	out := new(Jobs)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.api.jobs.JobsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) Rerun(ctx context.Context, in *Id, opts ...grpc.CallOption) (*RerunResponse, error) {
	out := new(RerunResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.api.jobs.JobsService/Rerun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) ListInitiatedScans(ctx context.Context, in *TimeQuery, opts ...grpc.CallOption) (*Ids, error) {
	out := new(Ids)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.api.jobs.JobsService/ListInitiatedScans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) GetJobResultByNodeId(ctx context.Context, in *GetJobResultByNodeIdRequest, opts ...grpc.CallOption) (*ResultsRow, error) {
	out := new(ResultsRow)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.compliance.api.jobs.JobsService/GetJobResultByNodeId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobsServiceServer is the server API for JobsService service.
type JobsServiceServer interface {
	Create(context.Context, *Job) (*Id, error)
	Read(context.Context, *Id) (*Job, error)
	Update(context.Context, *Job) (*empty.Empty, error)
	Delete(context.Context, *Id) (*empty.Empty, error)
	List(context.Context, *Query) (*Jobs, error)
	Rerun(context.Context, *Id) (*RerunResponse, error)
	ListInitiatedScans(context.Context, *TimeQuery) (*Ids, error)
	GetJobResultByNodeId(context.Context, *GetJobResultByNodeIdRequest) (*ResultsRow, error)
}

func RegisterJobsServiceServer(s *grpc.Server, srv JobsServiceServer) {
	s.RegisterService(&_JobsService_serviceDesc, srv)
}

func _JobsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.api.jobs.JobsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).Create(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.api.jobs.JobsService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).Read(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.api.jobs.JobsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).Update(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.api.jobs.JobsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.api.jobs.JobsService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).List(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_Rerun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).Rerun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.api.jobs.JobsService/Rerun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).Rerun(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_ListInitiatedScans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).ListInitiatedScans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.api.jobs.JobsService/ListInitiatedScans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).ListInitiatedScans(ctx, req.(*TimeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_GetJobResultByNodeId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobResultByNodeIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).GetJobResultByNodeId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.compliance.api.jobs.JobsService/GetJobResultByNodeId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).GetJobResultByNodeId(ctx, req.(*GetJobResultByNodeIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.compliance.api.jobs.JobsService",
	HandlerType: (*JobsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _JobsService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _JobsService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _JobsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _JobsService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _JobsService_List_Handler,
		},
		{
			MethodName: "Rerun",
			Handler:    _JobsService_Rerun_Handler,
		},
		{
			MethodName: "ListInitiatedScans",
			Handler:    _JobsService_ListInitiatedScans_Handler,
		},
		{
			MethodName: "GetJobResultByNodeId",
			Handler:    _JobsService_GetJobResultByNodeId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/compliance-service/api/jobs/jobs.proto",
}
