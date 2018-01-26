// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/api/scheduler/v1/scheduler.proto

/*
Package scheduler is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/scheduler/api/scheduler/v1/scheduler.proto
	go.chromium.org/luci/scheduler/api/scheduler/v1/triggers.proto

It has these top-level messages:
	JobsRequest
	JobsReply
	InvocationsRequest
	InvocationsReply
	EmitTriggersRequest
	JobRef
	InvocationRef
	Job
	JobState
	Invocation
	Trigger
	NoopTrigger
	GitilesTrigger
	BuildbucketTrigger
*/
package scheduler

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JobsRequest struct {
	// If not specified or "", all projects' jobs are returned.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	Cursor  string `protobuf:"bytes,2,opt,name=cursor" json:"cursor,omitempty"`
	// page_size is currently not implemented and is ignored.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *JobsRequest) Reset()                    { *m = JobsRequest{} }
func (m *JobsRequest) String() string            { return proto.CompactTextString(m) }
func (*JobsRequest) ProtoMessage()               {}
func (*JobsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JobsRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *JobsRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *JobsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type JobsReply struct {
	Jobs       []*Job `protobuf:"bytes,1,rep,name=jobs" json:"jobs,omitempty"`
	NextCursor string `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
}

func (m *JobsReply) Reset()                    { *m = JobsReply{} }
func (m *JobsReply) String() string            { return proto.CompactTextString(m) }
func (*JobsReply) ProtoMessage()               {}
func (*JobsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobsReply) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *JobsReply) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

type InvocationsRequest struct {
	JobRef *JobRef `protobuf:"bytes,1,opt,name=job_ref,json=jobRef" json:"job_ref,omitempty"`
	Cursor string  `protobuf:"bytes,2,opt,name=cursor" json:"cursor,omitempty"`
	// page_size defaults to 50 which is maximum.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *InvocationsRequest) Reset()                    { *m = InvocationsRequest{} }
func (m *InvocationsRequest) String() string            { return proto.CompactTextString(m) }
func (*InvocationsRequest) ProtoMessage()               {}
func (*InvocationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InvocationsRequest) GetJobRef() *JobRef {
	if m != nil {
		return m.JobRef
	}
	return nil
}

func (m *InvocationsRequest) GetCursor() string {
	if m != nil {
		return m.Cursor
	}
	return ""
}

func (m *InvocationsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type InvocationsReply struct {
	Invocations []*Invocation `protobuf:"bytes,1,rep,name=invocations" json:"invocations,omitempty"`
	NextCursor  string        `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
}

func (m *InvocationsReply) Reset()                    { *m = InvocationsReply{} }
func (m *InvocationsReply) String() string            { return proto.CompactTextString(m) }
func (*InvocationsReply) ProtoMessage()               {}
func (*InvocationsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InvocationsReply) GetInvocations() []*Invocation {
	if m != nil {
		return m.Invocations
	}
	return nil
}

func (m *InvocationsReply) GetNextCursor() string {
	if m != nil {
		return m.NextCursor
	}
	return ""
}

type EmitTriggersRequest struct {
	// A trigger and jobs it should be delivered to.
	//
	// Order is important. Triggers that are listed earlier are considered older.
	Batches []*EmitTriggersRequest_Batch `protobuf:"bytes,1,rep,name=batches" json:"batches,omitempty"`
	// An optional timestamp to use as trigger creation time, as unix timestamp in
	// microseconds. Assigned by the server by default. If given, must be within
	// +-15 min of the current time.
	//
	// Under some conditions triggers are ordered by timestamp of when they are
	// created. By allowing the client to specify this timestamp, we make
	// EmitTrigger RPC idempotent: if EmitTrigger call fails midway, the caller
	// can retry it providing exact same timestamp to get the correct final order
	// of the triggers.
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *EmitTriggersRequest) Reset()                    { *m = EmitTriggersRequest{} }
func (m *EmitTriggersRequest) String() string            { return proto.CompactTextString(m) }
func (*EmitTriggersRequest) ProtoMessage()               {}
func (*EmitTriggersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EmitTriggersRequest) GetBatches() []*EmitTriggersRequest_Batch {
	if m != nil {
		return m.Batches
	}
	return nil
}

func (m *EmitTriggersRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type EmitTriggersRequest_Batch struct {
	Trigger *Trigger  `protobuf:"bytes,1,opt,name=trigger" json:"trigger,omitempty"`
	Jobs    []*JobRef `protobuf:"bytes,2,rep,name=jobs" json:"jobs,omitempty"`
}

func (m *EmitTriggersRequest_Batch) Reset()                    { *m = EmitTriggersRequest_Batch{} }
func (m *EmitTriggersRequest_Batch) String() string            { return proto.CompactTextString(m) }
func (*EmitTriggersRequest_Batch) ProtoMessage()               {}
func (*EmitTriggersRequest_Batch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *EmitTriggersRequest_Batch) GetTrigger() *Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (m *EmitTriggersRequest_Batch) GetJobs() []*JobRef {
	if m != nil {
		return m.Jobs
	}
	return nil
}

// JobRef uniquely identifies a job.
type JobRef struct {
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	Job     string `protobuf:"bytes,2,opt,name=job" json:"job,omitempty"`
}

func (m *JobRef) Reset()                    { *m = JobRef{} }
func (m *JobRef) String() string            { return proto.CompactTextString(m) }
func (*JobRef) ProtoMessage()               {}
func (*JobRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *JobRef) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *JobRef) GetJob() string {
	if m != nil {
		return m.Job
	}
	return ""
}

// InvocationRef uniquely identifies an invocation of a job.
type InvocationRef struct {
	JobRef *JobRef `protobuf:"bytes,1,opt,name=job_ref,json=jobRef" json:"job_ref,omitempty"`
	// invocation_id is a unique integer among all invocations for a given job.
	// However, there could be invocations with the same invocation_id but
	// belonging to different jobs.
	InvocationId int64 `protobuf:"varint,2,opt,name=invocation_id,json=invocationId" json:"invocation_id,omitempty"`
}

func (m *InvocationRef) Reset()                    { *m = InvocationRef{} }
func (m *InvocationRef) String() string            { return proto.CompactTextString(m) }
func (*InvocationRef) ProtoMessage()               {}
func (*InvocationRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *InvocationRef) GetJobRef() *JobRef {
	if m != nil {
		return m.JobRef
	}
	return nil
}

func (m *InvocationRef) GetInvocationId() int64 {
	if m != nil {
		return m.InvocationId
	}
	return 0
}

// Job descibes currently configured job.
type Job struct {
	JobRef   *JobRef   `protobuf:"bytes,1,opt,name=job_ref,json=jobRef" json:"job_ref,omitempty"`
	Schedule string    `protobuf:"bytes,2,opt,name=schedule" json:"schedule,omitempty"`
	State    *JobState `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	Paused   bool      `protobuf:"varint,4,opt,name=paused" json:"paused,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Job) GetJobRef() *JobRef {
	if m != nil {
		return m.JobRef
	}
	return nil
}

func (m *Job) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *Job) GetState() *JobState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Job) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

// JobState describes current Job state as one of these strings:
//   "DISABLED"
//   "OVERRUN"
//   "PAUSED"
//   "RETRYING"
//   "RUNNING"
//   "SCHEDULED"
//   "STARTING"
//   "SUSPENDED"
//   "WAITING"
type JobState struct {
	UiStatus string `protobuf:"bytes,1,opt,name=ui_status,json=uiStatus" json:"ui_status,omitempty"`
}

func (m *JobState) Reset()                    { *m = JobState{} }
func (m *JobState) String() string            { return proto.CompactTextString(m) }
func (*JobState) ProtoMessage()               {}
func (*JobState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *JobState) GetUiStatus() string {
	if m != nil {
		return m.UiStatus
	}
	return ""
}

// Invocation describes properties of one job execution.
type Invocation struct {
	InvocationRef *InvocationRef `protobuf:"bytes,1,opt,name=invocation_ref,json=invocationRef" json:"invocation_ref,omitempty"`
	// start_ts is unix timestamp in microseconds.
	StartedTs int64 `protobuf:"varint,2,opt,name=started_ts,json=startedTs" json:"started_ts,omitempty"`
	// finished_ts is unix timestamp in microseconds. Set only if final is true.
	FinishedTs int64 `protobuf:"varint,3,opt,name=finished_ts,json=finishedTs" json:"finished_ts,omitempty"`
	// triggered_by is an identity ("kind:value") which is specified only if
	// invocation was triggered by not the scheduler service itself.
	TriggeredBy string `protobuf:"bytes,4,opt,name=triggered_by,json=triggeredBy" json:"triggered_by,omitempty"`
	// Latest status of a job.
	Status string `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	// If true, this invocation properties are final and won't be changed.
	Final bool `protobuf:"varint,6,opt,name=final" json:"final,omitempty"`
	// config_revision pins project/job config version according to which this
	// invocation was created.
	ConfigRevision string `protobuf:"bytes,7,opt,name=config_revision,json=configRevision" json:"config_revision,omitempty"`
	// view_url points to human readable page for a given invocation if available.
	ViewUrl string `protobuf:"bytes,8,opt,name=view_url,json=viewUrl" json:"view_url,omitempty"`
}

func (m *Invocation) Reset()                    { *m = Invocation{} }
func (m *Invocation) String() string            { return proto.CompactTextString(m) }
func (*Invocation) ProtoMessage()               {}
func (*Invocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Invocation) GetInvocationRef() *InvocationRef {
	if m != nil {
		return m.InvocationRef
	}
	return nil
}

func (m *Invocation) GetStartedTs() int64 {
	if m != nil {
		return m.StartedTs
	}
	return 0
}

func (m *Invocation) GetFinishedTs() int64 {
	if m != nil {
		return m.FinishedTs
	}
	return 0
}

func (m *Invocation) GetTriggeredBy() string {
	if m != nil {
		return m.TriggeredBy
	}
	return ""
}

func (m *Invocation) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Invocation) GetFinal() bool {
	if m != nil {
		return m.Final
	}
	return false
}

func (m *Invocation) GetConfigRevision() string {
	if m != nil {
		return m.ConfigRevision
	}
	return ""
}

func (m *Invocation) GetViewUrl() string {
	if m != nil {
		return m.ViewUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*JobsRequest)(nil), "scheduler.JobsRequest")
	proto.RegisterType((*JobsReply)(nil), "scheduler.JobsReply")
	proto.RegisterType((*InvocationsRequest)(nil), "scheduler.InvocationsRequest")
	proto.RegisterType((*InvocationsReply)(nil), "scheduler.InvocationsReply")
	proto.RegisterType((*EmitTriggersRequest)(nil), "scheduler.EmitTriggersRequest")
	proto.RegisterType((*EmitTriggersRequest_Batch)(nil), "scheduler.EmitTriggersRequest.Batch")
	proto.RegisterType((*JobRef)(nil), "scheduler.JobRef")
	proto.RegisterType((*InvocationRef)(nil), "scheduler.InvocationRef")
	proto.RegisterType((*Job)(nil), "scheduler.Job")
	proto.RegisterType((*JobState)(nil), "scheduler.JobState")
	proto.RegisterType((*Invocation)(nil), "scheduler.Invocation")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Scheduler service

type SchedulerClient interface {
	// GetJobs fetches all jobs satisfying JobsRequest and visibility ACLs.
	// If JobsRequest.project is specified but the project doesn't exist, empty
	// list of Jobs is returned.
	GetJobs(ctx context.Context, in *JobsRequest, opts ...grpc.CallOption) (*JobsReply, error)
	// GetInvocations fetches invocations of a given job, most recent first.
	GetInvocations(ctx context.Context, in *InvocationsRequest, opts ...grpc.CallOption) (*InvocationsReply, error)
	// PauseJob will prevent automatic triggering of a job. Manual triggering such
	// as through this API is still allowed. Any pending or running invocations
	// are still executed. PauseJob does nothing if job is already paused.
	//
	// Requires OWNER Job permission.
	PauseJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// ResumeJob resumes paused job. ResumeJob does nothing if job is not paused.
	//
	// Requires OWNER Job permission.
	ResumeJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// AbortJob resets the job to scheduled state, aborting a currently pending or
	// running invocation if any.
	//
	// Note, that this is similar to AbortInvocation except that AbortInvocation
	// requires invocation ID and doesn't ensure that the invocation aborted is
	// actually latest triggered for the job.
	//
	// Requires OWNER Job permission.
	AbortJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// AbortInvocation aborts a given job invocation.
	// If an invocation is final, AbortInvocation does nothing.
	//
	// If you want to abort a specific hung invocation, use this request instead
	// of AbortJob.
	//
	// Requires OWNER Job permission.
	AbortInvocation(ctx context.Context, in *InvocationRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// EmitTriggers puts one or more triggers into pending trigger queues of the
	// specified jobs.
	//
	// This eventually causes jobs to start executing. The scheduler may merge
	// multiple triggers into one job execution, based on how the job is
	// configured.
	//
	// If at least one job doesn't exist or the caller has no permission to
	// trigger it, the entire request is aborted. Otherwise, the request is NOT
	// transactional: if it fails midway (e.g by returning internal server error),
	// some triggers may have been submitted and some may not. It is safe to retry
	// the call, supplying the same trigger IDs. Triggers with the same IDs will
	// be deduplicated. See Trigger message for more details.
	//
	// Requires OWNER Job permission.
	//
	// TODO(vadimsh): add new TRIGGERER role.
	// TODO(vadimsh): deduplication doesn't work in v1.
	EmitTriggers(ctx context.Context, in *EmitTriggersRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}
type schedulerPRPCClient struct {
	client *prpc.Client
}

func NewSchedulerPRPCClient(client *prpc.Client) SchedulerClient {
	return &schedulerPRPCClient{client}
}

func (c *schedulerPRPCClient) GetJobs(ctx context.Context, in *JobsRequest, opts ...grpc.CallOption) (*JobsReply, error) {
	out := new(JobsReply)
	err := c.client.Call(ctx, "scheduler.Scheduler", "GetJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerPRPCClient) GetInvocations(ctx context.Context, in *InvocationsRequest, opts ...grpc.CallOption) (*InvocationsReply, error) {
	out := new(InvocationsReply)
	err := c.client.Call(ctx, "scheduler.Scheduler", "GetInvocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerPRPCClient) PauseJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := c.client.Call(ctx, "scheduler.Scheduler", "PauseJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerPRPCClient) ResumeJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := c.client.Call(ctx, "scheduler.Scheduler", "ResumeJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerPRPCClient) AbortJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := c.client.Call(ctx, "scheduler.Scheduler", "AbortJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerPRPCClient) AbortInvocation(ctx context.Context, in *InvocationRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := c.client.Call(ctx, "scheduler.Scheduler", "AbortInvocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerPRPCClient) EmitTriggers(ctx context.Context, in *EmitTriggersRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := c.client.Call(ctx, "scheduler.Scheduler", "EmitTriggers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type schedulerClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerClient(cc *grpc.ClientConn) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) GetJobs(ctx context.Context, in *JobsRequest, opts ...grpc.CallOption) (*JobsReply, error) {
	out := new(JobsReply)
	err := grpc.Invoke(ctx, "/scheduler.Scheduler/GetJobs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetInvocations(ctx context.Context, in *InvocationsRequest, opts ...grpc.CallOption) (*InvocationsReply, error) {
	out := new(InvocationsReply)
	err := grpc.Invoke(ctx, "/scheduler.Scheduler/GetInvocations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) PauseJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/scheduler.Scheduler/PauseJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) ResumeJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/scheduler.Scheduler/ResumeJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AbortJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/scheduler.Scheduler/AbortJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AbortInvocation(ctx context.Context, in *InvocationRef, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/scheduler.Scheduler/AbortInvocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) EmitTriggers(ctx context.Context, in *EmitTriggersRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/scheduler.Scheduler/EmitTriggers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Scheduler service

type SchedulerServer interface {
	// GetJobs fetches all jobs satisfying JobsRequest and visibility ACLs.
	// If JobsRequest.project is specified but the project doesn't exist, empty
	// list of Jobs is returned.
	GetJobs(context.Context, *JobsRequest) (*JobsReply, error)
	// GetInvocations fetches invocations of a given job, most recent first.
	GetInvocations(context.Context, *InvocationsRequest) (*InvocationsReply, error)
	// PauseJob will prevent automatic triggering of a job. Manual triggering such
	// as through this API is still allowed. Any pending or running invocations
	// are still executed. PauseJob does nothing if job is already paused.
	//
	// Requires OWNER Job permission.
	PauseJob(context.Context, *JobRef) (*google_protobuf.Empty, error)
	// ResumeJob resumes paused job. ResumeJob does nothing if job is not paused.
	//
	// Requires OWNER Job permission.
	ResumeJob(context.Context, *JobRef) (*google_protobuf.Empty, error)
	// AbortJob resets the job to scheduled state, aborting a currently pending or
	// running invocation if any.
	//
	// Note, that this is similar to AbortInvocation except that AbortInvocation
	// requires invocation ID and doesn't ensure that the invocation aborted is
	// actually latest triggered for the job.
	//
	// Requires OWNER Job permission.
	AbortJob(context.Context, *JobRef) (*google_protobuf.Empty, error)
	// AbortInvocation aborts a given job invocation.
	// If an invocation is final, AbortInvocation does nothing.
	//
	// If you want to abort a specific hung invocation, use this request instead
	// of AbortJob.
	//
	// Requires OWNER Job permission.
	AbortInvocation(context.Context, *InvocationRef) (*google_protobuf.Empty, error)
	// EmitTriggers puts one or more triggers into pending trigger queues of the
	// specified jobs.
	//
	// This eventually causes jobs to start executing. The scheduler may merge
	// multiple triggers into one job execution, based on how the job is
	// configured.
	//
	// If at least one job doesn't exist or the caller has no permission to
	// trigger it, the entire request is aborted. Otherwise, the request is NOT
	// transactional: if it fails midway (e.g by returning internal server error),
	// some triggers may have been submitted and some may not. It is safe to retry
	// the call, supplying the same trigger IDs. Triggers with the same IDs will
	// be deduplicated. See Trigger message for more details.
	//
	// Requires OWNER Job permission.
	//
	// TODO(vadimsh): add new TRIGGERER role.
	// TODO(vadimsh): deduplication doesn't work in v1.
	EmitTriggers(context.Context, *EmitTriggersRequest) (*google_protobuf.Empty, error)
}

func RegisterSchedulerServer(s prpc.Registrar, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_GetJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/GetJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetJobs(ctx, req.(*JobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetInvocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetInvocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/GetInvocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetInvocations(ctx, req.(*InvocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_PauseJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).PauseJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/PauseJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).PauseJob(ctx, req.(*JobRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_ResumeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).ResumeJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/ResumeJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).ResumeJob(ctx, req.(*JobRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AbortJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AbortJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/AbortJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AbortJob(ctx, req.(*JobRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AbortInvocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvocationRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AbortInvocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/AbortInvocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AbortInvocation(ctx, req.(*InvocationRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_EmitTriggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitTriggersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).EmitTriggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/EmitTriggers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).EmitTriggers(ctx, req.(*EmitTriggersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJobs",
			Handler:    _Scheduler_GetJobs_Handler,
		},
		{
			MethodName: "GetInvocations",
			Handler:    _Scheduler_GetInvocations_Handler,
		},
		{
			MethodName: "PauseJob",
			Handler:    _Scheduler_PauseJob_Handler,
		},
		{
			MethodName: "ResumeJob",
			Handler:    _Scheduler_ResumeJob_Handler,
		},
		{
			MethodName: "AbortJob",
			Handler:    _Scheduler_AbortJob_Handler,
		},
		{
			MethodName: "AbortInvocation",
			Handler:    _Scheduler_AbortInvocation_Handler,
		},
		{
			MethodName: "EmitTriggers",
			Handler:    _Scheduler_EmitTriggers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/scheduler/api/scheduler/v1/scheduler.proto",
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/api/scheduler/v1/scheduler.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x4e, 0xe3, 0x46,
	0x18, 0x55, 0x08, 0x49, 0xec, 0x2f, 0x10, 0xe8, 0x40, 0x91, 0x1b, 0x4a, 0x9b, 0xba, 0xad, 0x48,
	0xab, 0xca, 0x51, 0xd3, 0x52, 0xee, 0x40, 0x05, 0x51, 0x04, 0xea, 0x05, 0x72, 0xe8, 0x1d, 0x92,
	0x6b, 0x3b, 0x63, 0x67, 0x22, 0x27, 0xe3, 0x7a, 0xc6, 0x69, 0xc3, 0x53, 0xf4, 0x19, 0xf6, 0x5d,
	0xf6, 0x01, 0xf6, 0x8d, 0x56, 0x33, 0x1e, 0xc7, 0xce, 0x92, 0xb0, 0xca, 0x5e, 0x25, 0xe7, 0x7c,
	0x7f, 0x73, 0xbe, 0x39, 0x1e, 0xb8, 0x0c, 0xa9, 0xe5, 0x8f, 0x12, 0x3a, 0x21, 0xe9, 0xc4, 0xa2,
	0x49, 0xd8, 0x8b, 0x52, 0x9f, 0xf4, 0x98, 0x3f, 0xc2, 0xc3, 0x34, 0xc2, 0x49, 0xcf, 0x8d, 0xcb,
	0x68, 0xf6, 0x73, 0x01, 0xac, 0x38, 0xa1, 0x9c, 0x22, 0x7d, 0x41, 0xb4, 0x8f, 0x43, 0x4a, 0xc3,
	0x08, 0xf7, 0x64, 0xc0, 0x4b, 0x83, 0x1e, 0x9e, 0xc4, 0x7c, 0x9e, 0xe5, 0xb5, 0x2f, 0x36, 0x1d,
	0xc4, 0x13, 0x12, 0x86, 0x38, 0x61, 0x59, 0xbd, 0xf9, 0x04, 0xcd, 0x7b, 0xea, 0x31, 0x1b, 0xff,
	0x93, 0x62, 0xc6, 0x91, 0x01, 0x8d, 0x38, 0xa1, 0x63, 0xec, 0x73, 0xa3, 0xd2, 0xa9, 0x74, 0x75,
	0x3b, 0x87, 0xe8, 0x08, 0xea, 0x7e, 0x9a, 0x30, 0x9a, 0x18, 0x5b, 0x32, 0xa0, 0x10, 0x3a, 0x06,
	0x3d, 0x76, 0x43, 0xec, 0x30, 0xf2, 0x8c, 0x8d, 0x6a, 0xa7, 0xd2, 0xad, 0xd9, 0x9a, 0x20, 0x06,
	0xe4, 0x19, 0x9b, 0x0f, 0xa0, 0x67, 0xdd, 0xe3, 0x68, 0x8e, 0x4c, 0xd8, 0x1e, 0x53, 0x8f, 0x19,
	0x95, 0x4e, 0xb5, 0xdb, 0xec, 0xb7, 0xac, 0x42, 0xf2, 0x3d, 0xf5, 0x6c, 0x19, 0x43, 0x5f, 0x43,
	0x73, 0x8a, 0xff, 0xe3, 0xce, 0xd2, 0x28, 0x10, 0xd4, 0xb5, 0x64, 0xcc, 0x14, 0xd0, 0xdd, 0x74,
	0x46, 0x7d, 0x97, 0x13, 0x3a, 0x5d, 0x1c, 0xfb, 0x47, 0x68, 0x8c, 0xa9, 0xe7, 0x24, 0x38, 0x90,
	0xc7, 0x6e, 0xf6, 0x3f, 0xfb, 0xa0, 0x3b, 0x0e, 0xec, 0xfa, 0x58, 0xfe, 0x7e, 0x9a, 0x90, 0x08,
	0xf6, 0x97, 0xc6, 0x0a, 0x3d, 0xe7, 0xd0, 0x24, 0x05, 0xa7, 0x64, 0x7d, 0x5e, 0x1a, 0x5c, 0x54,
	0xd8, 0xe5, 0xcc, 0x8f, 0x8b, 0x7c, 0x57, 0x81, 0x83, 0x9b, 0x09, 0xe1, 0x8f, 0xea, 0xae, 0x72,
	0x99, 0x17, 0xd0, 0xf0, 0x5c, 0xee, 0x8f, 0x70, 0x3e, 0xed, 0xbb, 0xd2, 0xb4, 0x15, 0x05, 0xd6,
	0x95, 0xc8, 0xb6, 0xf3, 0x22, 0xf4, 0x25, 0xe8, 0x9c, 0x4c, 0x30, 0xe3, 0xee, 0x24, 0x96, 0x63,
	0xab, 0x76, 0x41, 0xb4, 0x9f, 0xa0, 0x26, 0xf3, 0xd1, 0x4f, 0xd0, 0x50, 0x2e, 0x51, 0xdb, 0x44,
	0xa5, 0x31, 0x6a, 0x84, 0x9d, 0xa7, 0xa0, 0xef, 0xd5, 0xb5, 0x6e, 0xc9, 0x13, 0xad, 0x58, 0xbc,
	0x0c, 0x9b, 0xbf, 0x42, 0x3d, 0xc3, 0xaf, 0x78, 0x6c, 0x1f, 0xaa, 0x63, 0xea, 0xa9, 0x85, 0x88,
	0xbf, 0xe6, 0xdf, 0xb0, 0x5b, 0xda, 0x22, 0x0e, 0x36, 0xba, 0xe9, 0x6f, 0x61, 0xb7, 0x58, 0xbb,
	0x43, 0x86, 0x4a, 0xf2, 0x4e, 0x41, 0xde, 0x0d, 0xcd, 0xff, 0x2b, 0x50, 0xbd, 0xa7, 0xde, 0x46,
	0x8d, 0xdb, 0xa0, 0xe5, 0x31, 0x75, 0xd8, 0x05, 0x46, 0x3f, 0x40, 0x8d, 0x71, 0x97, 0x67, 0x16,
	0x6a, 0xf6, 0x0f, 0x96, 0xbb, 0x0c, 0x44, 0xc8, 0xce, 0x32, 0x84, 0x13, 0x63, 0x37, 0x65, 0x78,
	0x68, 0x6c, 0x77, 0x2a, 0x5d, 0xcd, 0x56, 0xc8, 0x3c, 0x05, 0x2d, 0x4f, 0x15, 0xae, 0x4c, 0x89,
	0x23, 0xf2, 0x53, 0xa6, 0xd6, 0xa5, 0xa5, 0x64, 0x20, 0xb1, 0xf9, 0x66, 0x0b, 0xa0, 0x58, 0x0f,
	0xba, 0x84, 0x56, 0x49, 0x6f, 0xa1, 0xc4, 0x58, 0xed, 0x49, 0x1c, 0xd8, 0xa5, 0xfd, 0x08, 0x5d,
	0x27, 0x00, 0x8c, 0xbb, 0x09, 0xc7, 0x43, 0x87, 0xb3, 0xdc, 0x20, 0x8a, 0x79, 0x94, 0xbe, 0x0d,
	0xc8, 0x94, 0xb0, 0x51, 0x16, 0xaf, 0xca, 0x38, 0xe4, 0xd4, 0x23, 0x43, 0xdf, 0xc0, 0x8e, 0x72,
	0x05, 0x1e, 0x3a, 0xde, 0x5c, 0xca, 0xd2, 0xed, 0xe6, 0x82, 0xbb, 0x9a, 0x0b, 0xcd, 0x4a, 0x4c,
	0x2d, 0xfb, 0xfa, 0x32, 0x84, 0x0e, 0xa1, 0x16, 0x90, 0xa9, 0x1b, 0x19, 0x75, 0xb9, 0x8a, 0x0c,
	0xa0, 0x53, 0xd8, 0xf3, 0xe9, 0x34, 0x20, 0xa1, 0x93, 0xe0, 0x19, 0x61, 0x84, 0x4e, 0x8d, 0x86,
	0x2c, 0x6b, 0x65, 0xb4, 0xad, 0x58, 0xf4, 0x05, 0x68, 0x33, 0x82, 0xff, 0x75, 0xd2, 0x24, 0x32,
	0xb4, 0xcc, 0x54, 0x02, 0xff, 0x95, 0x44, 0xfd, 0xb7, 0x55, 0xd0, 0x07, 0xb9, 0x7e, 0x74, 0x0e,
	0x8d, 0x5b, 0xcc, 0xc5, 0xa3, 0x84, 0x8e, 0x96, 0xaf, 0x26, 0xff, 0x68, 0xda, 0x87, 0x2f, 0x78,
	0xf1, 0xb5, 0xff, 0x09, 0xad, 0x5b, 0xcc, 0x4b, 0x8f, 0x00, 0x3a, 0x59, 0xb9, 0xd6, 0x45, 0x9b,
	0xe3, 0x75, 0x61, 0xd1, 0xed, 0x0c, 0xb4, 0x07, 0x71, 0xd9, 0xc2, 0x79, 0x2f, 0x8d, 0xd6, 0x3e,
	0xb2, 0xb2, 0x37, 0xdf, 0xca, 0xdf, 0x7c, 0xeb, 0x46, 0xbc, 0xf9, 0xe8, 0x37, 0xd0, 0x6d, 0xcc,
	0xd2, 0xc9, 0xa6, 0x75, 0x67, 0xa0, 0xfd, 0xee, 0xd1, 0x84, 0x6f, 0x58, 0x76, 0x0d, 0x7b, 0xb2,
	0xac, 0xe4, 0xb1, 0xb5, 0x5e, 0x5a, 0xdb, 0xe4, 0x0f, 0xd8, 0x29, 0x3f, 0x4d, 0xe8, 0xab, 0xd7,
	0xdf, 0xac, 0x75, 0x7d, 0xbc, 0xba, 0xc4, 0xbf, 0xbc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x30, 0x3f,
	0x34, 0xcd, 0x5b, 0x07, 0x00, 0x00,
}
