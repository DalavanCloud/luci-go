// Code generated by protoc-gen-go.
// source: distributor.proto
// DO NOT EDIT!

/*
Package distributor is a generated protocol buffer package.

It is generated from these files:
	distributor.proto

It has these top-level messages:
	Alias
	Distributor
	Config
*/
package distributor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import "github.com/luci/luci-go/dm/api/distributor/jobsim"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Alias struct {
	OtherConfig string `protobuf:"bytes,1,opt,name=other_config,json=otherConfig" json:"other_config,omitempty"`
}

func (m *Alias) Reset()                    { *m = Alias{} }
func (m *Alias) String() string            { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()               {}
func (*Alias) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Distributor struct {
	// TODO(iannucci): Maybe something like Any or extensions would be a better
	// fit here? The ultimate goal is that users will be able to use the proto
	// text format for luci-config. I suspect that Any or extensions would lose
	// the ability to validate such text-formatted protobufs, but maybe that's
	// not the case.
	//
	// Types that are valid to be assigned to DistributorType:
	//	*Distributor_Alias
	//	*Distributor_Jobsim
	DistributorType isDistributor_DistributorType `protobuf_oneof:"distributor_type"`
}

func (m *Distributor) Reset()                    { *m = Distributor{} }
func (m *Distributor) String() string            { return proto.CompactTextString(m) }
func (*Distributor) ProtoMessage()               {}
func (*Distributor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isDistributor_DistributorType interface {
	isDistributor_DistributorType()
}

type Distributor_Alias struct {
	Alias *Alias `protobuf:"bytes,1,opt,name=alias,oneof"`
}
type Distributor_Jobsim struct {
	Jobsim *jobsim.Config `protobuf:"bytes,2048,opt,name=jobsim,oneof"`
}

func (*Distributor_Alias) isDistributor_DistributorType()  {}
func (*Distributor_Jobsim) isDistributor_DistributorType() {}

func (m *Distributor) GetDistributorType() isDistributor_DistributorType {
	if m != nil {
		return m.DistributorType
	}
	return nil
}

func (m *Distributor) GetAlias() *Alias {
	if x, ok := m.GetDistributorType().(*Distributor_Alias); ok {
		return x.Alias
	}
	return nil
}

func (m *Distributor) GetJobsim() *jobsim.Config {
	if x, ok := m.GetDistributorType().(*Distributor_Jobsim); ok {
		return x.Jobsim
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Distributor) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Distributor_OneofMarshaler, _Distributor_OneofUnmarshaler, _Distributor_OneofSizer, []interface{}{
		(*Distributor_Alias)(nil),
		(*Distributor_Jobsim)(nil),
	}
}

func _Distributor_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Distributor)
	// distributor_type
	switch x := m.DistributorType.(type) {
	case *Distributor_Alias:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Alias); err != nil {
			return err
		}
	case *Distributor_Jobsim:
		b.EncodeVarint(2048<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Jobsim); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Distributor.DistributorType has unexpected type %T", x)
	}
	return nil
}

func _Distributor_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Distributor)
	switch tag {
	case 1: // distributor_type.alias
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Alias)
		err := b.DecodeMessage(msg)
		m.DistributorType = &Distributor_Alias{msg}
		return true, err
	case 2048: // distributor_type.jobsim
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(jobsim.Config)
		err := b.DecodeMessage(msg)
		m.DistributorType = &Distributor_Jobsim{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Distributor_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Distributor)
	// distributor_type
	switch x := m.DistributorType.(type) {
	case *Distributor_Alias:
		s := proto.Size(x.Alias)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Distributor_Jobsim:
		s := proto.Size(x.Jobsim)
		n += proto.SizeVarint(2048<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Config struct {
	DistributorConfigs map[string]*Distributor `protobuf:"bytes,1,rep,name=distributor_configs,json=distributorConfigs" json:"distributor_configs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Config) GetDistributorConfigs() map[string]*Distributor {
	if m != nil {
		return m.DistributorConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*Alias)(nil), "distributor.Alias")
	proto.RegisterType((*Distributor)(nil), "distributor.Distributor")
	proto.RegisterType((*Config)(nil), "distributor.Config")
}

func init() { proto.RegisterFile("distributor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0x4b, 0x4b, 0xfd, 0x30,
	0x10, 0xc5, 0xff, 0x7d, 0xf2, 0x77, 0x2a, 0x12, 0xe3, 0xc2, 0x72, 0x57, 0xda, 0x95, 0x0f, 0x4c,
	0x40, 0x37, 0xe2, 0xce, 0xc7, 0x05, 0x71, 0xd9, 0xb5, 0x50, 0xfa, 0xb2, 0x37, 0xda, 0x36, 0xa5,
	0x4d, 0x85, 0xee, 0x5c, 0xfb, 0xc5, 0xfc, 0x5a, 0xa6, 0x49, 0xc1, 0x5c, 0xc4, 0xcd, 0xb4, 0x9c,
	0xfc, 0xce, 0x99, 0x99, 0x04, 0xf6, 0x0b, 0x36, 0x88, 0x9e, 0x65, 0xa3, 0xe0, 0x3d, 0xe9, 0x7a,
	0x2e, 0x38, 0x0e, 0x0c, 0x69, 0xb5, 0xae, 0x98, 0xd8, 0x8c, 0x19, 0xc9, 0x79, 0x43, 0xeb, 0x31,
	0x67, 0xaa, 0x5c, 0x54, 0x9c, 0x4a, 0xa1, 0xe1, 0x2d, 0x4d, 0x3b, 0x46, 0x8b, 0x86, 0x1a, 0x16,
	0xfa, 0xca, 0xb3, 0x81, 0x35, 0xcb, 0x47, 0x67, 0x46, 0x67, 0xe0, 0xdd, 0xd6, 0x2c, 0x1d, 0xf0,
	0x31, 0xec, 0x72, 0xb1, 0x29, 0xfb, 0x24, 0xe7, 0xed, 0x0b, 0xab, 0x42, 0xeb, 0xc8, 0x3a, 0xd9,
	0x89, 0x03, 0xa5, 0xdd, 0x2b, 0x29, 0xfa, 0xb4, 0x20, 0x78, 0xf8, 0xc9, 0xc3, 0xd2, 0x9b, 0xce,
	0x5e, 0xc5, 0x06, 0x97, 0x98, 0x98, 0x23, 0xab, 0xd4, 0xc7, 0x7f, 0xb1, 0x46, 0xf0, 0x29, 0xf8,
	0xba, 0x6f, 0xf8, 0x81, 0x14, 0xbd, 0x47, 0x96, 0x39, 0x74, 0xb8, 0x24, 0x17, 0xe0, 0x0e, 0x03,
	0x32, 0x82, 0x12, 0x31, 0x75, 0xe5, 0x93, 0xfb, 0xdf, 0x46, 0x8e, 0xac, 0x0e, 0x72, 0x65, 0x75,
	0x91, 0x17, 0x7d, 0x59, 0xe0, 0x6b, 0x2b, 0x7e, 0x86, 0x03, 0xd3, 0xa0, 0x17, 0x98, 0xa7, 0x72,
	0x64, 0x9f, 0xf3, 0xad, 0xa9, 0xb4, 0x83, 0x18, 0x5b, 0x68, 0x65, 0x58, 0xb7, 0xa2, 0x9f, 0x62,
	0x5c, 0xfc, 0x3a, 0x58, 0x25, 0x70, 0xf8, 0x07, 0x8e, 0x11, 0x38, 0x6f, 0xe5, 0xb4, 0x5c, 0xd5,
	0xfc, 0x8b, 0x09, 0x78, 0xef, 0x69, 0x3d, 0x96, 0xa1, 0xad, 0x96, 0x0c, 0xb7, 0x9a, 0x1b, 0x31,
	0xb1, 0xc6, 0x6e, 0xec, 0x6b, 0x2b, 0xf3, 0xd5, 0x4b, 0x5c, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0x97, 0x84, 0x81, 0xf2, 0x01, 0x00, 0x00,
}