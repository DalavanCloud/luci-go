// Code generated by protoc-gen-go.
// source: walk_graph.proto
// DO NOT EDIT!

package dm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/luci/luci-go/common/proto/google"
import _ "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Direction indicates that direction of dependencies that the request should
// walk.
type WalkGraphReq_Mode_Direction int32

const (
	WalkGraphReq_Mode_FORWARDS  WalkGraphReq_Mode_Direction = 0
	WalkGraphReq_Mode_BACKWARDS WalkGraphReq_Mode_Direction = 1
	WalkGraphReq_Mode_BOTH      WalkGraphReq_Mode_Direction = 2
)

var WalkGraphReq_Mode_Direction_name = map[int32]string{
	0: "FORWARDS",
	1: "BACKWARDS",
	2: "BOTH",
}
var WalkGraphReq_Mode_Direction_value = map[string]int32{
	"FORWARDS":  0,
	"BACKWARDS": 1,
	"BOTH":      2,
}

func (x WalkGraphReq_Mode_Direction) String() string {
	return proto.EnumName(WalkGraphReq_Mode_Direction_name, int32(x))
}
func (WalkGraphReq_Mode_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0, 0, 0}
}

// WalkGraphReq allows you to walk from one or more Quests through their
// Attempt's forward dependencies.
//
//
// The handler will evaluate all of the queries, executing them in parallel.
// For each attempt or quest produced by the query, it will queue a walk
// operation for that node, respecting the options set (max_depth, etc.).
type WalkGraphReq struct {
	// optional. See Include.AttemptResult for restrictions.
	Auth *Execution_Auth `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	// Query specifies a list of queries to start the graph traversal on. The
	// traversal will occur as a union of the query results. Redundant
	// specification will not cause additional heavy work; every graph node will
	// be processed exactly once, regardless of how many times it appears in the
	// query results. However, redundancy in the queries will cause the server to
	// retrieve and discard more information.
	Query *GraphQuery         `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	Mode  *WalkGraphReq_Mode  `protobuf:"bytes,3,opt,name=mode" json:"mode,omitempty"`
	Limit *WalkGraphReq_Limit `protobuf:"bytes,4,opt,name=limit" json:"limit,omitempty"`
	// Include allows you to add additional information to the returned
	// GraphData which is typically medium-to-large sized.
	Include *WalkGraphReq_Include `protobuf:"bytes,5,opt,name=include" json:"include,omitempty"`
	Exclude *WalkGraphReq_Exclude `protobuf:"bytes,6,opt,name=exclude" json:"exclude,omitempty"`
}

func (m *WalkGraphReq) Reset()                    { *m = WalkGraphReq{} }
func (m *WalkGraphReq) String() string            { return proto.CompactTextString(m) }
func (*WalkGraphReq) ProtoMessage()               {}
func (*WalkGraphReq) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *WalkGraphReq) GetAuth() *Execution_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *WalkGraphReq) GetQuery() *GraphQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *WalkGraphReq) GetMode() *WalkGraphReq_Mode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *WalkGraphReq) GetLimit() *WalkGraphReq_Limit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *WalkGraphReq) GetInclude() *WalkGraphReq_Include {
	if m != nil {
		return m.Include
	}
	return nil
}

func (m *WalkGraphReq) GetExclude() *WalkGraphReq_Exclude {
	if m != nil {
		return m.Exclude
	}
	return nil
}

type WalkGraphReq_Mode struct {
	// DFS sets whether this is a Depth-first (ish) or a Breadth-first (ish) load.
	// Since the load operation is multi-threaded, the search order is best
	// effort, but will actually be some hybrid between DFS and BFS. This setting
	// controls the bias direction of the hybrid loading algorithm.
	Dfs       bool                        `protobuf:"varint,1,opt,name=dfs" json:"dfs,omitempty"`
	Direction WalkGraphReq_Mode_Direction `protobuf:"varint,2,opt,name=direction,enum=dm.WalkGraphReq_Mode_Direction" json:"direction,omitempty"`
}

func (m *WalkGraphReq_Mode) Reset()                    { *m = WalkGraphReq_Mode{} }
func (m *WalkGraphReq_Mode) String() string            { return proto.CompactTextString(m) }
func (*WalkGraphReq_Mode) ProtoMessage()               {}
func (*WalkGraphReq_Mode) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 0} }

type WalkGraphReq_Limit struct {
	// MaxDepth sets the number of attempts to traverse; 0 means 'immediate'
	// (no dependencies), -1 means 'no limit', and >0 is a limit.
	//
	// Any negative value besides -1 is an error.
	MaxDepth int64 `protobuf:"varint,1,opt,name=max_depth,json=maxDepth" json:"max_depth,omitempty"`
	// MaxTime sets the maximum amount of time that the query processor should
	// take. Application of this deadline is 'best effort', which means the query
	// may take a bit longer than this timeout and still attempt to return data.
	//
	// This is different than the grpc timeout header, which will set a hard
	// deadline for the request.
	MaxTime *google_protobuf2.Duration `protobuf:"bytes,2,opt,name=max_time,json=maxTime" json:"max_time,omitempty"`
	// MaxDataSize sets the maximum amount of 'Data' (in bytes) that can be
	// returned, if include.quest_data, include.attempt_data, and/or
	// include.attempt_result are set. If this limit is hit, then the
	// appropriate 'partial' value will be set for that object, but the base
	// object would still be included in the result.
	//
	// If this limit is 0, a default limit of 16MB will be used. If this limit
	// exceeds 30MB, it will be reduced to 30MB.
	MaxDataSize uint32 `protobuf:"varint,3,opt,name=max_data_size,json=maxDataSize" json:"max_data_size,omitempty"`
}

func (m *WalkGraphReq_Limit) Reset()                    { *m = WalkGraphReq_Limit{} }
func (m *WalkGraphReq_Limit) String() string            { return proto.CompactTextString(m) }
func (*WalkGraphReq_Limit) ProtoMessage()               {}
func (*WalkGraphReq_Limit) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 1} }

func (m *WalkGraphReq_Limit) GetMaxTime() *google_protobuf2.Duration {
	if m != nil {
		return m.MaxTime
	}
	return nil
}

type WalkGraphReq_Include struct {
	Quest     *WalkGraphReq_Include_Options `protobuf:"bytes,1,opt,name=quest" json:"quest,omitempty"`
	Attempt   *WalkGraphReq_Include_Options `protobuf:"bytes,2,opt,name=attempt" json:"attempt,omitempty"`
	Execution *WalkGraphReq_Include_Options `protobuf:"bytes,3,opt,name=execution" json:"execution,omitempty"`
	// Executions is the number of Executions to include per Attempt. If this
	// is 0, then the execution data will be omitted completely.
	//
	// Executions included are from high ids to low ids. So setting this to `1`
	// would return the LAST execution made for this Attempt.
	NumExecutions uint32 `protobuf:"varint,4,opt,name=num_executions,json=numExecutions" json:"num_executions,omitempty"`
	// FwdDeps instructs WalkGraph to include forward dependency information
	// from the result. This only changes the presence of information in the
	// result; if the query is walking forward attempt dependencies, that will
	// still occur even if this is false.
	FwdDeps bool `protobuf:"varint,5,opt,name=fwd_deps,json=fwdDeps" json:"fwd_deps,omitempty"`
	// BackDeps instructs WalkGraph to include the backwards dependency
	// information. This only changes the presence of information in the result;
	// if the query is walking backward attempt dependencies, that will still
	// occur even if this is false.
	BackDeps bool `protobuf:"varint,6,opt,name=back_deps,json=backDeps" json:"back_deps,omitempty"`
}

func (m *WalkGraphReq_Include) Reset()                    { *m = WalkGraphReq_Include{} }
func (m *WalkGraphReq_Include) String() string            { return proto.CompactTextString(m) }
func (*WalkGraphReq_Include) ProtoMessage()               {}
func (*WalkGraphReq_Include) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 2} }

func (m *WalkGraphReq_Include) GetQuest() *WalkGraphReq_Include_Options {
	if m != nil {
		return m.Quest
	}
	return nil
}

func (m *WalkGraphReq_Include) GetAttempt() *WalkGraphReq_Include_Options {
	if m != nil {
		return m.Attempt
	}
	return nil
}

func (m *WalkGraphReq_Include) GetExecution() *WalkGraphReq_Include_Options {
	if m != nil {
		return m.Execution
	}
	return nil
}

type WalkGraphReq_Include_Options struct {
	// fills the 'id' field.
	//
	// If this is false, it will be omitted.
	//
	// Note that there's enough information contextually to derive these ids
	// on the client side, though it can be handy to have the server produce
	// them for you.
	Ids bool `protobuf:"varint,1,opt,name=ids" json:"ids,omitempty"`
	// instructs the request to include the Data field
	Data bool `protobuf:"varint,2,opt,name=data" json:"data,omitempty"`
	// instructs finished objects to include the Result field.
	//
	// If the requestor is an execution, the query logic will only include the
	// result if the execution's Attempt depends on it, otherwise it will be
	// blank.
	//
	// If the request's cumulative result data would be more than
	// limit.max_data_size of data, the remaining results will have their
	// Partial.Result set to DATA_SIZE_LIMIT.
	//
	// Has no effect for Quests.
	Result bool `protobuf:"varint,3,opt,name=result" json:"result,omitempty"`
	// If set to true, objects with an abnormal termination will be included.
	Abnormal bool `protobuf:"varint,4,opt,name=abnormal" json:"abnormal,omitempty"`
	// If set to true, expired objects will be included.
	Expired bool `protobuf:"varint,5,opt,name=expired" json:"expired,omitempty"`
}

func (m *WalkGraphReq_Include_Options) Reset()         { *m = WalkGraphReq_Include_Options{} }
func (m *WalkGraphReq_Include_Options) String() string { return proto.CompactTextString(m) }
func (*WalkGraphReq_Include_Options) ProtoMessage()    {}
func (*WalkGraphReq_Include_Options) Descriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0, 2, 0}
}

type WalkGraphReq_Exclude struct {
	// do not include data from the following quests in the response.
	Quests []string `protobuf:"bytes,1,rep,name=quests" json:"quests,omitempty"`
	// do not include data from the following attempts in the response.
	Attempts *AttemptList `protobuf:"bytes,2,opt,name=attempts" json:"attempts,omitempty"`
}

func (m *WalkGraphReq_Exclude) Reset()                    { *m = WalkGraphReq_Exclude{} }
func (m *WalkGraphReq_Exclude) String() string            { return proto.CompactTextString(m) }
func (*WalkGraphReq_Exclude) ProtoMessage()               {}
func (*WalkGraphReq_Exclude) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0, 3} }

func (m *WalkGraphReq_Exclude) GetAttempts() *AttemptList {
	if m != nil {
		return m.Attempts
	}
	return nil
}

func init() {
	proto.RegisterType((*WalkGraphReq)(nil), "dm.WalkGraphReq")
	proto.RegisterType((*WalkGraphReq_Mode)(nil), "dm.WalkGraphReq.Mode")
	proto.RegisterType((*WalkGraphReq_Limit)(nil), "dm.WalkGraphReq.Limit")
	proto.RegisterType((*WalkGraphReq_Include)(nil), "dm.WalkGraphReq.Include")
	proto.RegisterType((*WalkGraphReq_Include_Options)(nil), "dm.WalkGraphReq.Include.Options")
	proto.RegisterType((*WalkGraphReq_Exclude)(nil), "dm.WalkGraphReq.Exclude")
	proto.RegisterEnum("dm.WalkGraphReq_Mode_Direction", WalkGraphReq_Mode_Direction_name, WalkGraphReq_Mode_Direction_value)
}

func init() { proto.RegisterFile("walk_graph.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0x6d, 0x6b, 0x13, 0x4d,
	0x14, 0x7d, 0xd2, 0xbc, 0xec, 0xe6, 0xb6, 0xe9, 0x13, 0x2f, 0x58, 0xd2, 0x15, 0x6c, 0x29, 0x2a,
	0x15, 0x65, 0x0b, 0x51, 0xfc, 0x20, 0x28, 0xa4, 0xa6, 0xbe, 0x60, 0xb5, 0x38, 0x2d, 0xf4, 0x63,
	0x98, 0x64, 0xa7, 0xe9, 0xd2, 0xdd, 0xec, 0x76, 0x67, 0x96, 0xa6, 0x82, 0xfe, 0x01, 0x7f, 0x81,
	0xf8, 0x67, 0x9d, 0xb9, 0x33, 0x9b, 0x8a, 0xb5, 0xf4, 0xdb, 0xdc, 0x73, 0xce, 0xbd, 0x73, 0xe7,
	0xcc, 0x81, 0xee, 0x05, 0x4f, 0xce, 0x46, 0xd3, 0x82, 0xe7, 0xa7, 0x61, 0x5e, 0x64, 0x2a, 0xc3,
	0xa5, 0x28, 0x0d, 0xee, 0x4f, 0xb3, 0x6c, 0x9a, 0x88, 0x1d, 0x42, 0xc6, 0xe5, 0xc9, 0x4e, 0x54,
	0x16, 0x5c, 0xc5, 0xd9, 0xcc, 0x6a, 0x82, 0x8d, 0xbf, 0x79, 0x15, 0xa7, 0x42, 0x2a, 0x9e, 0xe6,
	0x4e, 0xd0, 0xa5, 0x89, 0xa3, 0x88, 0x2b, 0xee, 0x90, 0x3b, 0x16, 0x39, 0x2f, 0x45, 0x71, 0xe9,
	0xa0, 0x65, 0x75, 0x99, 0x0b, 0x69, 0x8b, 0xad, 0x9f, 0x3e, 0xac, 0x1c, 0xeb, 0x5d, 0xde, 0x19,
	0x19, 0x13, 0xe7, 0xf8, 0x08, 0x1a, 0xbc, 0x54, 0xa7, 0xbd, 0xda, 0x66, 0x6d, 0x7b, 0xb9, 0x8f,
	0x61, 0x94, 0x86, 0x7b, 0x73, 0x31, 0x29, 0x69, 0x8d, 0x81, 0x66, 0x18, 0xf1, 0xf8, 0x00, 0x9a,
	0x34, 0xb4, 0xb7, 0x44, 0xc2, 0x55, 0x23, 0xa4, 0x21, 0x5f, 0x0c, 0xca, 0x2c, 0x89, 0x8f, 0xa1,
	0x91, 0x66, 0x91, 0xe8, 0xd5, 0x49, 0x74, 0xd7, 0x88, 0xfe, 0xbc, 0x2d, 0xfc, 0xa4, 0x49, 0x46,
	0x12, 0x7c, 0x0a, 0xcd, 0x24, 0x4e, 0x63, 0xd5, 0x6b, 0x90, 0x76, 0xed, 0x9a, 0x76, 0xdf, 0xb0,
	0xcc, 0x8a, 0xb0, 0x0f, 0x5e, 0x3c, 0x9b, 0x24, 0xa5, 0x9e, 0xdd, 0x24, 0x7d, 0xef, 0x9a, 0xfe,
	0x83, 0xe5, 0x59, 0x25, 0x34, 0x3d, 0x62, 0x6e, 0x7b, 0x5a, 0x37, 0xf4, 0xec, 0xcd, 0x5d, 0x8f,
	0x13, 0x06, 0x3f, 0x6a, 0xd0, 0x30, 0x4b, 0x62, 0x17, 0xea, 0xd1, 0x89, 0x24, 0x5b, 0x7c, 0x66,
	0x8e, 0xf8, 0x0a, 0xda, 0x51, 0x5c, 0x88, 0x89, 0x71, 0x86, 0x5c, 0x58, 0xed, 0x6f, 0xfc, 0xf3,
	0x81, 0xe1, 0xb0, 0x92, 0xb1, 0xab, 0x8e, 0xad, 0x3e, 0xb4, 0x17, 0x38, 0xae, 0x80, 0xff, 0xf6,
	0x80, 0x1d, 0x0f, 0xd8, 0xf0, 0xb0, 0xfb, 0x1f, 0x76, 0xa0, 0xbd, 0x3b, 0x78, 0xf3, 0xd1, 0x96,
	0x35, 0xf4, 0xa1, 0xb1, 0x7b, 0x70, 0xf4, 0xbe, 0xbb, 0x14, 0x7c, 0x87, 0x26, 0xb9, 0x80, 0xf7,
	0xa0, 0x9d, 0xf2, 0xf9, 0x28, 0x12, 0xb9, 0xfb, 0xaa, 0x3a, 0xf3, 0x35, 0x30, 0x34, 0x35, 0x3e,
	0x07, 0x73, 0x1e, 0x99, 0x70, 0xb8, 0xdf, 0x59, 0x0f, 0x6d, 0x72, 0xc2, 0x2a, 0x39, 0xe1, 0xd0,
	0x25, 0x8b, 0x79, 0x5a, 0x7a, 0xa4, 0x95, 0xb8, 0x05, 0x1d, 0x1a, 0xa9, 0xb3, 0x33, 0x92, 0xf1,
	0x57, 0xfb, 0x67, 0x1d, 0xb6, 0x6c, 0xc6, 0x6a, 0xec, 0x50, 0x43, 0xc1, 0xaf, 0x3a, 0x78, 0xce,
	0x56, 0x7c, 0x41, 0x01, 0x90, 0xca, 0x25, 0x65, 0xf3, 0x26, 0xff, 0xc3, 0x83, 0xdc, 0x5c, 0x24,
	0x99, 0x95, 0xe3, 0x4b, 0xf0, 0xb8, 0x52, 0x22, 0xcd, 0x95, 0x5b, 0xee, 0xf6, 0xce, 0xaa, 0x01,
	0x5f, 0x43, 0x5b, 0x54, 0x61, 0x74, 0x99, 0xba, 0xbd, 0xfb, 0xaa, 0x05, 0x1f, 0xc2, 0xea, 0xac,
	0x4c, 0x47, 0x0b, 0x40, 0x52, 0xd8, 0x3a, 0xac, 0xa3, 0xd1, 0x45, 0xca, 0x25, 0xae, 0x83, 0x7f,
	0x72, 0x11, 0x19, 0x77, 0x25, 0xa5, 0xcb, 0x67, 0x9e, 0xae, 0xb5, 0xb9, 0xd2, 0x18, 0x3f, 0xe6,
	0x93, 0x33, 0xcb, 0xb5, 0x88, 0xf3, 0x0d, 0x60, 0xc8, 0xe0, 0x1b, 0x78, 0xee, 0x52, 0x13, 0x97,
	0x38, 0x5a, 0xc4, 0x45, 0x1f, 0x11, 0xa1, 0x61, 0xbc, 0xa5, 0x47, 0xfb, 0x8c, 0xce, 0xb8, 0x06,
	0xad, 0x42, 0xc8, 0x32, 0x51, 0xf4, 0x18, 0x9f, 0xb9, 0x0a, 0x03, 0xf0, 0xf9, 0x78, 0x96, 0x15,
	0x29, 0x4f, 0x68, 0x43, 0x7d, 0x49, 0x55, 0x63, 0xcf, 0xa4, 0x38, 0xd7, 0xc9, 0x89, 0xaa, 0xdd,
	0x5c, 0x19, 0x7c, 0x06, 0xcf, 0xe5, 0xd7, 0x0c, 0x26, 0xb7, 0xcd, 0x06, 0xf5, 0xed, 0x36, 0x73,
	0x15, 0x3e, 0xd1, 0x83, 0xad, 0x97, 0xd2, 0xb9, 0xff, 0xbf, 0xf1, 0x6f, 0x60, 0xb1, 0xfd, 0x58,
	0x2a, 0xb6, 0x10, 0x8c, 0x5b, 0x94, 0x96, 0x67, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0xb6,
	0x57, 0x2c, 0xad, 0x04, 0x00, 0x00,
}