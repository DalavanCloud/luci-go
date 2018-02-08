// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/machines.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "go.chromium.org/luci/machine-db/api/common/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A machine in the database.
type Machine struct {
	// The name of this machine. Uniquely identifies this machine.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The type of platform this machine is.
	Platform string `protobuf:"bytes,2,opt,name=platform" json:"platform,omitempty"`
	// The rack this machine belongs to.
	Rack string `protobuf:"bytes,3,opt,name=rack" json:"rack,omitempty"`
	// A description of this machine.
	Description string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// The asset tag associated with this machine.
	AssetTag string `protobuf:"bytes,5,opt,name=asset_tag,json=assetTag" json:"asset_tag,omitempty"`
	// The service tag associated with this machine.
	ServiceTag string `protobuf:"bytes,6,opt,name=service_tag,json=serviceTag" json:"service_tag,omitempty"`
	// The deployment ticket associated with this machine.
	DeploymentTicket string `protobuf:"bytes,7,opt,name=deployment_ticket,json=deploymentTicket" json:"deployment_ticket,omitempty"`
	// The state of this machine.
	State common.State `protobuf:"varint,8,opt,name=state,enum=common.State" json:"state,omitempty"`
}

func (m *Machine) Reset()                    { *m = Machine{} }
func (m *Machine) String() string            { return proto.CompactTextString(m) }
func (*Machine) ProtoMessage()               {}
func (*Machine) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Machine) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Machine) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *Machine) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *Machine) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Machine) GetAssetTag() string {
	if m != nil {
		return m.AssetTag
	}
	return ""
}

func (m *Machine) GetServiceTag() string {
	if m != nil {
		return m.ServiceTag
	}
	return ""
}

func (m *Machine) GetDeploymentTicket() string {
	if m != nil {
		return m.DeploymentTicket
	}
	return ""
}

func (m *Machine) GetState() common.State {
	if m != nil {
		return m.State
	}
	return common.State_STATE_UNSPECIFIED
}

// A request to create a new machine in the database.
type CreateMachineRequest struct {
	// The machine to create in the database.
	Machine *Machine `protobuf:"bytes,1,opt,name=machine" json:"machine,omitempty"`
}

func (m *CreateMachineRequest) Reset()                    { *m = CreateMachineRequest{} }
func (m *CreateMachineRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateMachineRequest) ProtoMessage()               {}
func (*CreateMachineRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *CreateMachineRequest) GetMachine() *Machine {
	if m != nil {
		return m.Machine
	}
	return nil
}

// A request to delete a machine from the database.
type DeleteMachineRequest struct {
	// The name of the machine to delete.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteMachineRequest) Reset()                    { *m = DeleteMachineRequest{} }
func (m *DeleteMachineRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteMachineRequest) ProtoMessage()               {}
func (*DeleteMachineRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *DeleteMachineRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request to list machines in the database.
type ListMachinesRequest struct {
	// The names of machines to get.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
	// The platforms to filter returned machines on.
	Platforms []string `protobuf:"bytes,2,rep,name=platforms" json:"platforms,omitempty"`
	// The racks to filter returned machines on.
	Racks []string `protobuf:"bytes,3,rep,name=racks" json:"racks,omitempty"`
	// The states to filter returned machines on.
	States []common.State `protobuf:"varint,4,rep,packed,name=states,enum=common.State" json:"states,omitempty"`
}

func (m *ListMachinesRequest) Reset()                    { *m = ListMachinesRequest{} }
func (m *ListMachinesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListMachinesRequest) ProtoMessage()               {}
func (*ListMachinesRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ListMachinesRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListMachinesRequest) GetPlatforms() []string {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func (m *ListMachinesRequest) GetRacks() []string {
	if m != nil {
		return m.Racks
	}
	return nil
}

func (m *ListMachinesRequest) GetStates() []common.State {
	if m != nil {
		return m.States
	}
	return nil
}

// A response containing a list of machines in the database.
type ListMachinesResponse struct {
	// The machines matching this request.
	Machines []*Machine `protobuf:"bytes,1,rep,name=machines" json:"machines,omitempty"`
}

func (m *ListMachinesResponse) Reset()                    { *m = ListMachinesResponse{} }
func (m *ListMachinesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListMachinesResponse) ProtoMessage()               {}
func (*ListMachinesResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *ListMachinesResponse) GetMachines() []*Machine {
	if m != nil {
		return m.Machines
	}
	return nil
}

func init() {
	proto.RegisterType((*Machine)(nil), "crimson.Machine")
	proto.RegisterType((*CreateMachineRequest)(nil), "crimson.CreateMachineRequest")
	proto.RegisterType((*DeleteMachineRequest)(nil), "crimson.DeleteMachineRequest")
	proto.RegisterType((*ListMachinesRequest)(nil), "crimson.ListMachinesRequest")
	proto.RegisterType((*ListMachinesResponse)(nil), "crimson.ListMachinesResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/machines.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0xaf, 0xd3, 0x30,
	0x14, 0x85, 0x95, 0x97, 0xb6, 0x69, 0x6e, 0x05, 0x7a, 0x98, 0x0c, 0xd6, 0x03, 0x89, 0x28, 0x08,
	0xa9, 0x2a, 0x90, 0x88, 0xb2, 0x21, 0xb1, 0x40, 0x47, 0x58, 0x42, 0xf7, 0xca, 0x75, 0x2f, 0xa9,
	0xd5, 0x38, 0x0e, 0xb6, 0x5b, 0x89, 0x3f, 0xc0, 0xdf, 0x06, 0xc5, 0x4e, 0x5a, 0x50, 0x3b, 0xb0,
	0xe5, 0x9e, 0xf3, 0xf9, 0xc8, 0xf7, 0xc4, 0xf0, 0xb1, 0x52, 0x39, 0xdf, 0x6b, 0x25, 0xc5, 0x51,
	0xe6, 0x4a, 0x57, 0x45, 0x7d, 0xe4, 0xa2, 0x90, 0x8c, 0xef, 0x45, 0x83, 0x6f, 0x77, 0xdb, 0x82,
	0xb5, 0xa2, 0xe0, 0x5a, 0x48, 0xa3, 0x9a, 0xe2, 0xf4, 0x6e, 0x70, 0x4c, 0xde, 0x6a, 0x65, 0x15,
	0x89, 0x7a, 0xeb, 0xe1, 0xc3, 0x7f, 0xe5, 0x28, 0x29, 0x7d, 0x8c, 0xb1, 0xcc, 0x0e, 0x21, 0xd9,
	0xef, 0x00, 0xa2, 0xaf, 0x9e, 0x24, 0x04, 0x46, 0x0d, 0x93, 0x48, 0x83, 0x34, 0x98, 0xc7, 0xa5,
	0xfb, 0x26, 0x0f, 0x30, 0x6d, 0x6b, 0x66, 0xbf, 0x2b, 0x2d, 0xe9, 0x9d, 0xd3, 0xcf, 0x73, 0xc7,
	0x6b, 0xc6, 0x0f, 0x34, 0xf4, 0x7c, 0xf7, 0x4d, 0x52, 0x98, 0xed, 0xd0, 0x70, 0x2d, 0x5a, 0x2b,
	0x54, 0x43, 0x47, 0xce, 0xfa, 0x5b, 0x22, 0xcf, 0x20, 0x66, 0xc6, 0xa0, 0xdd, 0x58, 0x56, 0xd1,
	0xb1, 0x8f, 0x74, 0xc2, 0x9a, 0x55, 0xe4, 0x05, 0xcc, 0x0c, 0xea, 0x93, 0xe0, 0xe8, 0xec, 0x89,
	0xb3, 0xa1, 0x97, 0x3a, 0xe0, 0x35, 0x3c, 0xd9, 0x61, 0x5b, 0xab, 0x9f, 0x12, 0x1b, 0xbb, 0xb1,
	0x82, 0x1f, 0xd0, 0xd2, 0xc8, 0x61, 0xf7, 0x17, 0x63, 0xed, 0x74, 0xf2, 0x12, 0xc6, 0x6e, 0x59,
	0x3a, 0x4d, 0x83, 0xf9, 0xe3, 0xe5, 0xa3, 0xdc, 0x97, 0x90, 0x7f, 0xeb, 0xc4, 0xd2, 0x7b, 0xd9,
	0x27, 0x48, 0x3e, 0x6b, 0x64, 0x16, 0xfb, 0x1a, 0x4a, 0xfc, 0x71, 0x44, 0x63, 0xc9, 0x02, 0xa2,
	0xbe, 0x42, 0x57, 0xc8, 0x6c, 0x79, 0x9f, 0xf7, 0x85, 0xe7, 0x03, 0x39, 0x00, 0xd9, 0x02, 0x92,
	0x15, 0xd6, 0x78, 0x95, 0x71, 0xa3, 0xd1, 0xec, 0x57, 0x00, 0x4f, 0xbf, 0x08, 0x63, 0x7b, 0xd4,
	0x0c, 0x6c, 0x02, 0xe3, 0xce, 0x37, 0x34, 0x48, 0xc3, 0x79, 0x5c, 0xfa, 0x81, 0x3c, 0x87, 0x78,
	0xe8, 0xdb, 0xd0, 0x3b, 0xe7, 0x5c, 0x84, 0xee, 0x4c, 0xd7, 0xba, 0xa1, 0xa1, 0x3f, 0xe3, 0x06,
	0xf2, 0x0a, 0x26, 0xfe, 0x1f, 0xd3, 0x51, 0x1a, 0x5e, 0xef, 0xdd, 0x9b, 0xd9, 0x0a, 0x92, 0x7f,
	0xef, 0x61, 0x5a, 0xd5, 0x18, 0x24, 0x6f, 0x60, 0x3a, 0xbc, 0x34, 0x77, 0x97, 0x5b, 0x9b, 0x9f,
	0x89, 0xed, 0xc4, 0xbd, 0xa3, 0xf7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x33, 0x1e, 0xce, 0x2a,
	0xcd, 0x02, 0x00, 0x00,
}
