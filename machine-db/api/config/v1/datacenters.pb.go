// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/config/v1/datacenters.proto

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/machine-db/api/config/v1/datacenters.proto
	go.chromium.org/luci/machine-db/api/config/v1/oses.proto
	go.chromium.org/luci/machine-db/api/config/v1/platforms.proto
	go.chromium.org/luci/machine-db/api/config/v1/vlans.proto

It has these top-level messages:
	Switch
	Rack
	Datacenter
	Datacenters
	OS
	OSes
	Platform
	Platforms
	VLAN
	VLANs
*/
package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Switch describes a switch.
type Switch struct {
	// The name of this switch.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this switch.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// The number of ports this switch has.
	Ports int32 `protobuf:"varint,3,opt,name=ports" json:"ports,omitempty"`
}

func (m *Switch) Reset()                    { *m = Switch{} }
func (m *Switch) String() string            { return proto.CompactTextString(m) }
func (*Switch) ProtoMessage()               {}
func (*Switch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Switch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Switch) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Switch) GetPorts() int32 {
	if m != nil {
		return m.Ports
	}
	return 0
}

// Rack describes a rack.
type Rack struct {
	// The name of this rack.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this rack.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// The switches belonging to this rack.
	Switch []*Switch `protobuf:"bytes,3,rep,name=switch" json:"switch,omitempty"`
}

func (m *Rack) Reset()                    { *m = Rack{} }
func (m *Rack) String() string            { return proto.CompactTextString(m) }
func (*Rack) ProtoMessage()               {}
func (*Rack) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Rack) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rack) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Rack) GetSwitch() []*Switch {
	if m != nil {
		return m.Switch
	}
	return nil
}

// Datacenter describes a datacenter.
type Datacenter struct {
	// The name of this datacenter.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this datacenter.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// The racks belonging to this datacenter.
	Rack []*Rack `protobuf:"bytes,3,rep,name=rack" json:"rack,omitempty"`
}

func (m *Datacenter) Reset()                    { *m = Datacenter{} }
func (m *Datacenter) String() string            { return proto.CompactTextString(m) }
func (*Datacenter) ProtoMessage()               {}
func (*Datacenter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Datacenter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Datacenter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Datacenter) GetRack() []*Rack {
	if m != nil {
		return m.Rack
	}
	return nil
}

// Datacenters enumerates datacenter config files.
type Datacenters struct {
	// A list of names of datacenter config files.
	Datacenter []string `protobuf:"bytes,1,rep,name=datacenter" json:"datacenter,omitempty"`
}

func (m *Datacenters) Reset()                    { *m = Datacenters{} }
func (m *Datacenters) String() string            { return proto.CompactTextString(m) }
func (*Datacenters) ProtoMessage()               {}
func (*Datacenters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Datacenters) GetDatacenter() []string {
	if m != nil {
		return m.Datacenter
	}
	return nil
}

func init() {
	proto.RegisterType((*Switch)(nil), "config.Switch")
	proto.RegisterType((*Rack)(nil), "config.Rack")
	proto.RegisterType((*Datacenter)(nil), "config.Datacenter")
	proto.RegisterType((*Datacenters)(nil), "config.Datacenters")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/config/v1/datacenters.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x41, 0x4e, 0xc4, 0x20,
	0x14, 0x40, 0x53, 0xdb, 0x69, 0x32, 0xbf, 0xc6, 0x05, 0x71, 0xc1, 0xca, 0x90, 0x2e, 0x4c, 0x37,
	0x03, 0x51, 0x0f, 0xe0, 0xc6, 0x13, 0xa0, 0x17, 0x60, 0x3e, 0xd8, 0x92, 0xb1, 0xd0, 0x00, 0xa3,
	0xd7, 0x37, 0xc2, 0x38, 0x33, 0xeb, 0xee, 0xe0, 0xbf, 0x84, 0xf7, 0xf8, 0xf0, 0x3a, 0x7a, 0x8e,
	0x53, 0xf0, 0xb3, 0x3d, 0xce, 0xdc, 0x87, 0x51, 0x7c, 0x1d, 0xd1, 0x8a, 0x59, 0xe1, 0x64, 0x9d,
	0xd9, 0xe9, 0xbd, 0x50, 0x8b, 0x15, 0xe8, 0xdd, 0xa7, 0x1d, 0xc5, 0xf7, 0x93, 0xd0, 0x2a, 0x29,
	0x34, 0x2e, 0x99, 0x10, 0xf9, 0x12, 0x7c, 0xf2, 0xa4, 0x2d, 0xb0, 0xff, 0x80, 0xf6, 0xfd, 0xc7,
	0x26, 0x9c, 0x08, 0x81, 0xc6, 0xa9, 0xd9, 0xd0, 0x8a, 0x55, 0xc3, 0x56, 0xe6, 0x33, 0x61, 0xd0,
	0x69, 0x13, 0x31, 0xd8, 0x25, 0x59, 0xef, 0xe8, 0x4d, 0x46, 0xd7, 0x23, 0x72, 0x0f, 0x9b, 0xc5,
	0x87, 0x14, 0x69, 0xcd, 0xaa, 0x61, 0x23, 0xcb, 0xa5, 0xd7, 0xd0, 0x48, 0x85, 0x87, 0x95, 0x6f,
	0x3e, 0x42, 0x1b, 0x73, 0x13, 0xad, 0x59, 0x3d, 0x74, 0xcf, 0x77, 0xbc, 0xc4, 0xf2, 0x52, 0x2a,
	0x4f, 0xb4, 0xd7, 0x00, 0x6f, 0xe7, 0x8f, 0xad, 0x74, 0x31, 0x68, 0x82, 0xc2, 0xc3, 0xc9, 0x74,
	0xfb, 0x6f, 0xfa, 0xab, 0x97, 0x99, 0xf4, 0x3b, 0xe8, 0x2e, 0x96, 0x48, 0x1e, 0x00, 0x2e, 0xdb,
	0xa4, 0x15, 0xab, 0x87, 0xad, 0xbc, 0x9a, 0xec, 0xdb, 0xbc, 0xdf, 0x97, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0x73, 0x36, 0x7c, 0xa2, 0x01, 0x00, 0x00,
}
