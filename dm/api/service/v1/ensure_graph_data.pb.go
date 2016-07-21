// Code generated by protoc-gen-go.
// source: ensure_graph_data.proto
// DO NOT EDIT!

package dm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import template "github.com/luci/luci-go/common/api/template"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TemplateInstantiation struct {
	// project is the luci-config project which defines the template.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// ref is the git ref of the project that defined this template. If omitted,
	// this will use the template definition from the project-wide configuration
	// and not the configuration located on a particular ref (like
	// 'refs/heads/master').
	Ref string `protobuf:"bytes,2,opt,name=ref" json:"ref,omitempty"`
	// specifier specifies the actual template name, as well as any substitution
	// parameters which that template might require.
	Specifier *template.Specifier `protobuf:"bytes,4,opt,name=specifier" json:"specifier,omitempty"`
}

func (m *TemplateInstantiation) Reset()                    { *m = TemplateInstantiation{} }
func (m *TemplateInstantiation) String() string            { return proto.CompactTextString(m) }
func (*TemplateInstantiation) ProtoMessage()               {}
func (*TemplateInstantiation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TemplateInstantiation) GetSpecifier() *template.Specifier {
	if m != nil {
		return m.Specifier
	}
	return nil
}

// EnsureGraphDataReq allows you to assert some things about the state of DM's
// graph:
//   * That 0 or more quest descriptions exist (the `quests` and
//     `template_quest` field).
//   * That 0 or more attempts exist (the `attempts` field)
//     * That those `attempts` are dependencies of a particular execution
//
// One of quests or attempts MUST be provided, it's an error for them to both
// be empty. Any quest description must have at least one corresponding Attempt
// number (i.e. you are not permitted to add a Quest without also adding at
// least one Attempt of it).
//
// In response, DM will tell you what the IDs of all supplied Quests/Attempts
// are. If the `attempts` were being created as dependencies, and were already
// in the Finished state, this request can also opt to include the
// AttemptResults directly.
type EnsureGraphDataReq struct {
	// Quest is a list of quest descriptors. DM will ensure that the
	// corresponding Quests exist. If they don't, they'll be created.
	Quest []*Quest_Desc `protobuf:"bytes,1,rep,name=quest" json:"quest,omitempty"`
	// Attempts is a list that asserts that the following attempts should
	// exist. The quest ids in this list must either be known to DM, or have
	// their descriptions included in the quests field above.
	Attempts *AttemptList `protobuf:"bytes,2,opt,name=attempts" json:"attempts,omitempty"`
	// TemplateQuest allows the addition of quests which are derived from
	// Templates, as defined on a per-project basis.
	TemplateQuest []*TemplateInstantiation `protobuf:"bytes,3,rep,name=template_quest,json=templateQuest" json:"template_quest,omitempty"`
	// TemplateAttempt allows the addition of attempts which are derived from
	// Templates. This must be equal in length to template_quest.
	// Each entry here maps 1:1 with the equivalent quest in template_quest.
	TemplateAttempt []*AttemptList_Nums `protobuf:"bytes,4,rep,name=template_attempt,json=templateAttempt" json:"template_attempt,omitempty"`
	// ForExecution is an authentication pair (Execution_ID, Token).
	//
	// If this is provided then it will serve as authorization for the creation of
	// any `quests` included, and any `attempts` indicated will be set as
	// dependencies for the execution.
	//
	// If this omitted, then the request requires some user/bot authentication,
	// and any quests/attempts provided will be made standalone (e.g. nothing will
	// depend on them).
	ForExecution *Execution_Auth             `protobuf:"bytes,5,opt,name=for_execution,json=forExecution" json:"for_execution,omitempty"`
	Limit        *EnsureGraphDataReq_Limit   `protobuf:"bytes,6,opt,name=limit" json:"limit,omitempty"`
	Include      *EnsureGraphDataReq_Include `protobuf:"bytes,7,opt,name=include" json:"include,omitempty"`
}

func (m *EnsureGraphDataReq) Reset()                    { *m = EnsureGraphDataReq{} }
func (m *EnsureGraphDataReq) String() string            { return proto.CompactTextString(m) }
func (*EnsureGraphDataReq) ProtoMessage()               {}
func (*EnsureGraphDataReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *EnsureGraphDataReq) GetQuest() []*Quest_Desc {
	if m != nil {
		return m.Quest
	}
	return nil
}

func (m *EnsureGraphDataReq) GetAttempts() *AttemptList {
	if m != nil {
		return m.Attempts
	}
	return nil
}

func (m *EnsureGraphDataReq) GetTemplateQuest() []*TemplateInstantiation {
	if m != nil {
		return m.TemplateQuest
	}
	return nil
}

func (m *EnsureGraphDataReq) GetTemplateAttempt() []*AttemptList_Nums {
	if m != nil {
		return m.TemplateAttempt
	}
	return nil
}

func (m *EnsureGraphDataReq) GetForExecution() *Execution_Auth {
	if m != nil {
		return m.ForExecution
	}
	return nil
}

func (m *EnsureGraphDataReq) GetLimit() *EnsureGraphDataReq_Limit {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *EnsureGraphDataReq) GetInclude() *EnsureGraphDataReq_Include {
	if m != nil {
		return m.Include
	}
	return nil
}

type EnsureGraphDataReq_Limit struct {
	// MaxDataSize sets the maximum amount of 'Data' (in bytes) that can be
	// returned, if include.attempt_result is set. If this limit is hit, then
	// the appropriate 'partial' value will be set for that object, but the base
	// object would still be included in the result.
	//
	// If this limit is 0, a default limit of 16MB will be used. If this limit
	// exceeds 30MB, it will be reduced to 30MB.
	MaxDataSize uint32 `protobuf:"varint,3,opt,name=max_data_size,json=maxDataSize" json:"max_data_size,omitempty"`
}

func (m *EnsureGraphDataReq_Limit) Reset()                    { *m = EnsureGraphDataReq_Limit{} }
func (m *EnsureGraphDataReq_Limit) String() string            { return proto.CompactTextString(m) }
func (*EnsureGraphDataReq_Limit) ProtoMessage()               {}
func (*EnsureGraphDataReq_Limit) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

type EnsureGraphDataReq_Include struct {
	// AttemptResult will include the Attempt result payloads for any Attempts
	// that it returns were already in the Finished state.
	//
	// If the request would return more than limit.max_data_size of data, the
	// remaining attempt results will have their Partial.Data field set to true.
	AttemptResult bool `protobuf:"varint,4,opt,name=attempt_result,json=attemptResult" json:"attempt_result,omitempty"`
}

func (m *EnsureGraphDataReq_Include) Reset()                    { *m = EnsureGraphDataReq_Include{} }
func (m *EnsureGraphDataReq_Include) String() string            { return proto.CompactTextString(m) }
func (*EnsureGraphDataReq_Include) ProtoMessage()               {}
func (*EnsureGraphDataReq_Include) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1} }

type EnsureGraphDataRsp struct {
	// accepted is true when all new graph data was journaled successfully. This
	// means that `quests`, `attempts`, `template_quest`, `template_attempt` were
	// all well-formed and are scheduled to be added. They will 'eventually' be
	// readable via other APIs (like WalkGraph), but when they are, they'll have
	// the IDs reflected in this response.
	//
	// If `attempts` referrs to quests that don't exist and weren't provided in
	// `quests`, those quests will be listed in `result` with the DNE flag set.
	//
	// If `template_quest` had errors (missing template, bad params, etc.), the
	// errors will be located in `template_error`. If all of the templates parsed
	// successfully, the quest ids for those rendered `template_quest` will be in
	// `template_ids`.
	Accepted bool `protobuf:"varint,1,opt,name=accepted" json:"accepted,omitempty"`
	// template_ids will be populated with the Quest.IDs of any templates defined
	// by template_quest in the initial request. Its length is guaranteed to match
	// the length of template_quest, if there were no errors.
	TemplateIds []*Quest_ID `protobuf:"bytes,2,rep,name=template_ids,json=templateIds" json:"template_ids,omitempty"`
	// template_error is either empty if there were no template errors, or the
	// length of template_quest. Non-empty strings are errors.
	TemplateError []string `protobuf:"bytes,3,rep,name=template_error,json=templateError" json:"template_error,omitempty"`
	// result holds the graph data pertaining to the request, containing any
	// graph state that already existed at the time of the call. Any new data
	// that was added to the graph state (accepted==true) will appear with
	// `DNE==true`.
	//
	// Quest data will always be returned for any Quests which exist.
	//
	// If accepted==false, you can inspect this to determine why:
	//   * Quests (without data) mentioned by the `attempts` field that do not
	//     exist will have `DNE==true`.
	//
	// This also can be used to make adding dependencies a stateless
	// single-request action:
	//   * Attempts requested (assuming the corresponding Quest exists) will
	//     contain their current state. If Include.AttemptResult was true, the
	//     results will be populated (with the size limit mentioned in the request
	//     documentation).
	Result *GraphData `protobuf:"bytes,4,opt,name=result" json:"result,omitempty"`
	// (if `for_execution` was specified) ShouldHalt indicates that the request
	// was accepted by DM, and the execution should halt (DM will re-execute the
	// Attempt when it becomes unblocked). If this is true, then the execution's
	// auth Token is also revoked and will no longer work for futher API calls.
	//
	// If `for_execution` was provided in the request and this is false, it means
	// that the execution may continue executing.
	ShouldHalt bool `protobuf:"varint,5,opt,name=should_halt,json=shouldHalt" json:"should_halt,omitempty"`
}

func (m *EnsureGraphDataRsp) Reset()                    { *m = EnsureGraphDataRsp{} }
func (m *EnsureGraphDataRsp) String() string            { return proto.CompactTextString(m) }
func (*EnsureGraphDataRsp) ProtoMessage()               {}
func (*EnsureGraphDataRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *EnsureGraphDataRsp) GetTemplateIds() []*Quest_ID {
	if m != nil {
		return m.TemplateIds
	}
	return nil
}

func (m *EnsureGraphDataRsp) GetResult() *GraphData {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*TemplateInstantiation)(nil), "dm.TemplateInstantiation")
	proto.RegisterType((*EnsureGraphDataReq)(nil), "dm.EnsureGraphDataReq")
	proto.RegisterType((*EnsureGraphDataReq_Limit)(nil), "dm.EnsureGraphDataReq.Limit")
	proto.RegisterType((*EnsureGraphDataReq_Include)(nil), "dm.EnsureGraphDataReq.Include")
	proto.RegisterType((*EnsureGraphDataRsp)(nil), "dm.EnsureGraphDataRsp")
}

func init() { proto.RegisterFile("ensure_graph_data.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0x5d, 0x8b, 0xd4, 0x30,
	0x14, 0x65, 0x3e, 0x3a, 0xed, 0xa4, 0xd3, 0xd9, 0x12, 0x15, 0x6b, 0x11, 0x95, 0x41, 0x61, 0x41,
	0xec, 0xe0, 0xf8, 0xb0, 0xe2, 0x8b, 0x2e, 0xcc, 0xa0, 0x23, 0x8b, 0x60, 0xd6, 0x27, 0x5f, 0x4a,
	0xb7, 0xcd, 0x4c, 0x23, 0xfd, 0xda, 0x26, 0x81, 0x55, 0xfc, 0x3d, 0xfe, 0x1e, 0x7f, 0x92, 0x37,
	0xe9, 0xc7, 0x2c, 0xbb, 0xee, 0xcb, 0xa5, 0x39, 0x39, 0xe7, 0xdc, 0x9b, 0xe4, 0x14, 0x3d, 0xa4,
	0x05, 0x97, 0x35, 0x0d, 0xf7, 0x75, 0x54, 0xa5, 0x61, 0x12, 0x89, 0x28, 0xa8, 0xea, 0x52, 0x94,
	0x78, 0x98, 0xe4, 0xfe, 0xbb, 0x3d, 0x13, 0xa9, 0xbc, 0x08, 0xe2, 0x32, 0x5f, 0x66, 0x32, 0x66,
	0xba, 0xbc, 0xda, 0x97, 0x4b, 0x00, 0xf2, 0xb2, 0x58, 0x46, 0x15, 0x5b, 0x0a, 0x9a, 0x57, 0x59,
	0x24, 0x68, 0xff, 0xd1, 0xe8, 0x7d, 0xf7, 0xa6, 0xa3, 0x6f, 0x8b, 0x9f, 0x15, 0xe5, 0xcd, 0x62,
	0xf1, 0x1b, 0x3d, 0xf8, 0xd6, 0x0a, 0xb6, 0x05, 0x17, 0x51, 0x21, 0x58, 0x24, 0x58, 0x59, 0x60,
	0x0f, 0x99, 0xc0, 0xf8, 0x41, 0x63, 0xe1, 0x0d, 0x9e, 0x0d, 0x8e, 0xa7, 0xa4, 0x5b, 0x62, 0x17,
	0x8d, 0x6a, 0xba, 0xf3, 0x86, 0x1a, 0x55, 0x9f, 0xf8, 0x35, 0x9a, 0xf2, 0x8a, 0xc6, 0x6c, 0xc7,
	0x68, 0xed, 0x8d, 0x01, 0xb7, 0x57, 0xf7, 0x82, 0x7e, 0x8e, 0xf3, 0x6e, 0x8b, 0x1c, 0x58, 0x9f,
	0xc7, 0xd6, 0xc8, 0x1d, 0x2f, 0xfe, 0x8c, 0x11, 0xde, 0xe8, 0x83, 0x7f, 0x54, 0x53, 0xae, 0x61,
	0x48, 0x42, 0x2f, 0xf1, 0x73, 0x64, 0x5c, 0x4a, 0xca, 0x55, 0xe7, 0x11, 0x78, 0xcd, 0x83, 0x24,
	0x0f, 0xbe, 0x2a, 0x20, 0x58, 0x53, 0x1e, 0x93, 0x66, 0x13, 0xbf, 0x44, 0x56, 0x24, 0x54, 0x17,
	0xc1, 0xf5, 0x30, 0xf6, 0xea, 0x48, 0x11, 0x4f, 0x1b, 0xec, 0x8c, 0x71, 0x41, 0x7a, 0x02, 0xfe,
	0x80, 0xe6, 0xdd, 0x40, 0x61, 0xe3, 0x3d, 0xd2, 0xde, 0x8f, 0x94, 0xe4, 0xbf, 0x37, 0x40, 0x9c,
	0x4e, 0xa0, 0x5b, 0xe3, 0xf7, 0xc8, 0xed, 0x1d, 0x5a, 0x5b, 0x38, 0xab, 0xf2, 0xb8, 0x7f, 0xa3,
	0x6d, 0xf0, 0x45, 0xe6, 0x9c, 0x1c, 0x75, 0xec, 0x76, 0x07, 0x9f, 0x20, 0x67, 0x57, 0xd6, 0x21,
	0xbd, 0xa2, 0xb1, 0x54, 0x0d, 0x3c, 0x43, 0x0f, 0x8d, 0x95, 0x7a, 0xd3, 0x81, 0xc1, 0xa9, 0x14,
	0x29, 0x99, 0x01, 0xb1, 0x87, 0xf0, 0x0a, 0x19, 0x19, 0xcb, 0x99, 0xf0, 0x26, 0x5a, 0xf0, 0x58,
	0x0b, 0x6e, 0xdd, 0x5a, 0x70, 0xa6, 0x38, 0xa4, 0xa1, 0xe2, 0xb7, 0xc8, 0x64, 0x45, 0x9c, 0xc9,
	0x84, 0x7a, 0xa6, 0x56, 0x3d, 0xb9, 0x43, 0xb5, 0x6d, 0x58, 0xa4, 0xa3, 0xfb, 0x27, 0xc8, 0xd0,
	0x4e, 0x78, 0x81, 0x9c, 0x3c, 0xba, 0xd2, 0xc9, 0x09, 0x39, 0xfb, 0x45, 0xe1, 0xc6, 0x06, 0xc7,
	0x0e, 0xb1, 0x01, 0x54, 0xe2, 0x73, 0x80, 0xe0, 0x19, 0x07, 0xee, 0x10, 0xea, 0xd0, 0x1d, 0xf9,
	0xdf, 0x91, 0xd9, 0x9a, 0xe1, 0x17, 0x68, 0xde, 0x5e, 0x51, 0x58, 0x53, 0x2e, 0x33, 0xa1, 0x53,
	0x61, 0x11, 0xa7, 0x45, 0x89, 0x06, 0xaf, 0xab, 0x9b, 0x40, 0x40, 0x35, 0xdc, 0x09, 0xd4, 0x89,
	0x6b, 0x42, 0x35, 0x5d, 0x0b, 0xaa, 0xe5, 0x4e, 0x17, 0x7f, 0x07, 0xb7, 0x83, 0xc2, 0x2b, 0xec,
	0x43, 0x04, 0xe2, 0x98, 0x56, 0x82, 0x26, 0x3a, 0xa5, 0x16, 0xe9, 0xd7, 0x78, 0x89, 0x66, 0xfd,
	0x7b, 0xb1, 0x44, 0x45, 0x44, 0xbd, 0xd5, 0xec, 0x90, 0xa5, 0xed, 0x9a, 0xd8, 0x1d, 0x63, 0x9b,
	0x70, 0x35, 0x74, 0x2f, 0xa0, 0x75, 0x5d, 0xd6, 0x3a, 0x22, 0xd3, 0x43, 0x0e, 0x36, 0x0a, 0x04,
	0xda, 0xe4, 0xda, 0x99, 0xec, 0x95, 0xa3, 0x1c, 0x0f, 0x53, 0xb5, 0x9b, 0xf8, 0x29, 0xb2, 0x79,
	0x5a, 0xca, 0x2c, 0x09, 0xd3, 0x08, 0xb8, 0x86, 0x9e, 0x0e, 0x35, 0xd0, 0x27, 0x40, 0x2e, 0x26,
	0xfa, 0x07, 0x7c, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x16, 0x79, 0xec, 0xfa, 0x03, 0x00,
	0x00,
}