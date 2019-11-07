// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autopilot-operator.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// MeshProviders provide an interface to monitoring and managing a specific
// mesh.
// AutoPilot does not abstract the mesh API - AutoPilot developers must
// still reason able about Provider-specific CRDs. AutoPilot's job is to
// abstract operational concerns such as discovering control plane configuration
// and monitoring metrics.
type MeshProvider int32

const (
	// the Operator will utilize the Service Mesh Interface (SMI) for metrics and configuration.
	// Compatible with multiple meshes (may require installation of an SMI Adapter).
	MeshProvider_SMI MeshProvider = 0
	// the Operator will utilize Istio mesh for metrics and configuration
	MeshProvider_Istio MeshProvider = 1
)

var MeshProvider_name = map[int32]string{
	0: "SMI",
	1: "Istio",
}

var MeshProvider_value = map[string]int32{
	"SMI":   0,
	"Istio": 1,
}

func (x MeshProvider) String() string {
	return proto.EnumName(MeshProvider_name, int32(x))
}

func (MeshProvider) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56f975433f2c607a, []int{0}
}

// The AutoPilotOperator file is the bootstrap
// Configuration file for the Operator.
// It is stored and mounted to the operator as a Kubernetes ConfigMap.
// The Operator will hot-reload when the configuration file changes.
// Default name is 'autopilot-operator.yaml' and should be stored in the project root.
type AutoPilotOperator struct {
	// Version of the operator
	// used for reporting, metrics, etc (can be any format)
	// default is "0.0.1"
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// MeshProvider determines how the operator will connect to a service mesh
	// Default is "SMI"
	MeshProvider MeshProvider `protobuf:"varint,2,opt,name=MeshProvider,proto3,enum=autopilot.MeshProvider" json:"MeshProvider,omitempty"`
	// WorkInterval to sets the interval at which CRD workers resync.
	// Default is 5s
	WorkInterval         *duration.Duration `protobuf:"bytes,3,opt,name=WorkInterval,proto3" json:"WorkInterval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AutoPilotOperator) Reset()         { *m = AutoPilotOperator{} }
func (m *AutoPilotOperator) String() string { return proto.CompactTextString(m) }
func (*AutoPilotOperator) ProtoMessage()    {}
func (*AutoPilotOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f975433f2c607a, []int{0}
}

func (m *AutoPilotOperator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoPilotOperator.Unmarshal(m, b)
}
func (m *AutoPilotOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoPilotOperator.Marshal(b, m, deterministic)
}
func (m *AutoPilotOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoPilotOperator.Merge(m, src)
}
func (m *AutoPilotOperator) XXX_Size() int {
	return xxx_messageInfo_AutoPilotOperator.Size(m)
}
func (m *AutoPilotOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoPilotOperator.DiscardUnknown(m)
}

var xxx_messageInfo_AutoPilotOperator proto.InternalMessageInfo

func (m *AutoPilotOperator) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AutoPilotOperator) GetMeshProvider() MeshProvider {
	if m != nil {
		return m.MeshProvider
	}
	return MeshProvider_SMI
}

func (m *AutoPilotOperator) GetWorkInterval() *duration.Duration {
	if m != nil {
		return m.WorkInterval
	}
	return nil
}

func init() {
	proto.RegisterEnum("autopilot.MeshProvider", MeshProvider_name, MeshProvider_value)
	proto.RegisterType((*AutoPilotOperator)(nil), "autopilot.AutoPilotOperator")
}

func init() { proto.RegisterFile("autopilot-operator.proto", fileDescriptor_56f975433f2c607a) }

var fileDescriptor_56f975433f2c607a = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0x87, 0x8d, 0x43, 0x47, 0xe3, 0x90, 0x99, 0x8b, 0xd1, 0x83, 0x94, 0x89, 0x50, 0x84, 0x25,
	0x38, 0x8f, 0xe2, 0x41, 0xf1, 0xd2, 0xc3, 0x70, 0xd4, 0x83, 0xe0, 0x2d, 0x75, 0xb1, 0x0b, 0xd6,
	0xfe, 0x4a, 0xfa, 0xa6, 0xdf, 0xc9, 0x6f, 0x29, 0xee, 0x4f, 0x59, 0x8f, 0x2f, 0xbf, 0x87, 0xf7,
	0x79, 0xb8, 0x34, 0x81, 0x50, 0xbb, 0x12, 0x34, 0x45, 0x6d, 0xbd, 0x21, 0x78, 0x55, 0x7b, 0x10,
	0x44, 0xd4, 0x2d, 0x97, 0x57, 0x05, 0x50, 0x94, 0x56, 0xaf, 0x87, 0x3c, 0x7c, 0xe9, 0x65, 0xf0,
	0x86, 0x1c, 0xaa, 0x0d, 0x3a, 0xf9, 0x65, 0xfc, 0xec, 0x29, 0x10, 0x16, 0xff, 0xf4, 0xeb, 0xf6,
	0x8d, 0x90, 0x7c, 0xd8, 0x5a, 0xdf, 0x38, 0x54, 0x92, 0xc5, 0x2c, 0x89, 0xb2, 0xdd, 0x29, 0x1e,
	0xf8, 0x68, 0x6e, 0x9b, 0xd5, 0xc2, 0xa3, 0x75, 0x4b, 0xeb, 0xe5, 0x61, 0xcc, 0x92, 0xd3, 0xd9,
	0xb9, 0xea, 0x8c, 0x6a, 0x7f, 0xce, 0x7a, 0xb0, 0x78, 0xe4, 0xa3, 0x77, 0xf8, 0xef, 0xb4, 0x22,
	0xeb, 0x5b, 0x53, 0xca, 0x41, 0xcc, 0x92, 0x93, 0xd9, 0x85, 0xda, 0x34, 0xaa, 0x5d, 0xa3, 0x7a,
	0xd9, 0x36, 0x66, 0x3d, 0xfc, 0x76, 0xd2, 0x77, 0x8b, 0x21, 0x1f, 0xbc, 0xcd, 0xd3, 0xf1, 0x81,
	0x88, 0xf8, 0x51, 0xda, 0x90, 0xc3, 0x98, 0x3d, 0xdf, 0x7c, 0x5c, 0x17, 0x8e, 0x56, 0x21, 0x57,
	0x9f, 0xf8, 0xd1, 0x0d, 0x4a, 0x4c, 0x1d, 0x74, 0x57, 0xa7, 0x4d, 0xed, 0x74, 0x7b, 0x97, 0x1f,
	0xaf, 0x5d, 0xf7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x61, 0x58, 0x1b, 0xa4, 0x44, 0x01, 0x00,
	0x00,
}