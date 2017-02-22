// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/vpython/api/env/spec.proto
// DO NOT EDIT!

/*
Package env is a generated protocol buffer package.

It is generated from these files:
	github.com/luci/luci-go/vpython/api/env/spec.proto

It has these top-level messages:
	Spec
*/
package env

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

type Spec struct {
	// The Python version to use. This should be of the form:
	// "Major[.Minor[.Patch]]"
	//
	// If specified,
	// - The Major version will be enforced absolutely. Python 3 will not be
	//   preferred over Python 2 because '3' is greater than '2'.
	// - The remaining versions, if specified, will be regarded as *minimum*
	//   versions. In other words, if "2.7.4" is specified and the system has
	//   "2.7.12", that will suffice. Similarly, "2.6" would accept a "2.7"
	//   interpreter.
	//
	// If empty, the default Python interpreter ("python") will be used.
	PythonVersion string          `protobuf:"bytes,1,opt,name=python_version,json=pythonVersion" json:"python_version,omitempty"`
	Wheel         []*Spec_Package `protobuf:"bytes,2,rep,name=wheel" json:"wheel,omitempty"`
	// The VirtualEnv package.
	//
	// This should be left empty to use the `vpython` default package
	// (recommended).
	Virtualenv *Spec_Package `protobuf:"bytes,3,opt,name=virtualenv" json:"virtualenv,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Spec) GetPythonVersion() string {
	if m != nil {
		return m.PythonVersion
	}
	return ""
}

func (m *Spec) GetWheel() []*Spec_Package {
	if m != nil {
		return m.Wheel
	}
	return nil
}

func (m *Spec) GetVirtualenv() *Spec_Package {
	if m != nil {
		return m.Virtualenv
	}
	return nil
}

// A definition for a remote package. The type of package depends on the
// configured package resolver.
type Spec_Package struct {
	// The path of the package.
	//
	// - For CIPD, this is the package name.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// The package version.
	//
	// - For CIPD, this will be any recognized CIPD version (i.e., ID, tag, or
	//   ref).
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *Spec_Package) Reset()                    { *m = Spec_Package{} }
func (m *Spec_Package) String() string            { return proto.CompactTextString(m) }
func (*Spec_Package) ProtoMessage()               {}
func (*Spec_Package) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Spec_Package) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Spec_Package) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*Spec)(nil), "env.Spec")
	proto.RegisterType((*Spec_Package)(nil), "env.Spec.Package")
}

func init() { proto.RegisterFile("github.com/luci/luci-go/vpython/api/env/spec.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x49, 0x5b, 0x2d, 0x8e, 0x28, 0x98, 0x53, 0xf0, 0x54, 0x04, 0xb1, 0x17, 0x13, 0xac,
	0x07, 0x5f, 0x43, 0x2a, 0x78, 0x95, 0x34, 0x0c, 0x4d, 0xb0, 0x26, 0xa1, 0x4d, 0x23, 0xbe, 0x9e,
	0x4f, 0xb6, 0x6c, 0xd2, 0xc2, 0x1e, 0xf6, 0x32, 0xcc, 0x7c, 0xf3, 0x31, 0xfc, 0x03, 0xdd, 0x68,
	0x82, 0x5e, 0x07, 0xae, 0xdc, 0x8f, 0x98, 0x56, 0x65, 0x52, 0x79, 0x1e, 0x9d, 0x88, 0xfe, 0x2f,
	0x68, 0x67, 0x85, 0xf4, 0x46, 0xa0, 0x8d, 0x62, 0xf1, 0xa8, 0xb8, 0x9f, 0x5d, 0x70, 0xb4, 0x44,
	0x1b, 0x1f, 0xfe, 0x09, 0x54, 0x1f, 0x1e, 0x15, 0x7d, 0x84, 0xdb, 0xac, 0x7e, 0x45, 0x9c, 0x17,
	0xe3, 0x2c, 0x23, 0x0d, 0x69, 0xaf, 0xfa, 0x9b, 0x4c, 0x3f, 0x33, 0xa4, 0x4f, 0x70, 0xf1, 0xab,
	0x11, 0x27, 0x56, 0x34, 0x65, 0x7b, 0xdd, 0xdd, 0x71, 0xb4, 0x91, 0x1f, 0x0f, 0xf0, 0x77, 0xa9,
	0xbe, 0xe5, 0x88, 0x7d, 0xde, 0xd3, 0x17, 0x80, 0x68, 0xe6, 0xb0, 0xca, 0x09, 0x6d, 0x64, 0x65,
	0x43, 0xce, 0xdb, 0x27, 0xd2, 0xfd, 0x1b, 0xd4, 0x1b, 0xa6, 0x14, 0x2a, 0x2f, 0x83, 0xde, 0x32,
	0xa4, 0x9e, 0x32, 0xa8, 0xf7, 0x68, 0x45, 0xc2, 0xfb, 0x38, 0x5c, 0xa6, 0x87, 0x5e, 0x0f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x2a, 0x9f, 0x43, 0x06, 0x01, 0x00, 0x00,
}