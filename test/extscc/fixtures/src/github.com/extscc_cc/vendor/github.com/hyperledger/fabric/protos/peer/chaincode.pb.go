// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Confidentiality Levels
type ConfidentialityLevel int32

const (
	ConfidentialityLevel_PUBLIC       ConfidentialityLevel = 0
	ConfidentialityLevel_CONFIDENTIAL ConfidentialityLevel = 1
)

var ConfidentialityLevel_name = map[int32]string{
	0: "PUBLIC",
	1: "CONFIDENTIAL",
}
var ConfidentialityLevel_value = map[string]int32{
	"PUBLIC":       0,
	"CONFIDENTIAL": 1,
}

func (x ConfidentialityLevel) String() string {
	return proto.EnumName(ConfidentialityLevel_name, int32(x))
}
func (ConfidentialityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
	ChaincodeSpec_BINARY    ChaincodeSpec_Type = 5
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
	5: "BINARY",
}
var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
	"BINARY":    5,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}
func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type ChaincodeDeploymentSpec_ExecutionEnvironment int32

const (
	ChaincodeDeploymentSpec_DOCKER     ChaincodeDeploymentSpec_ExecutionEnvironment = 0
	ChaincodeDeploymentSpec_SYSTEM     ChaincodeDeploymentSpec_ExecutionEnvironment = 1
	ChaincodeDeploymentSpec_SYSTEM_EXT ChaincodeDeploymentSpec_ExecutionEnvironment = 2
)

var ChaincodeDeploymentSpec_ExecutionEnvironment_name = map[int32]string{
	0: "DOCKER",
	1: "SYSTEM",
	2: "SYSTEM_EXT",
}
var ChaincodeDeploymentSpec_ExecutionEnvironment_value = map[string]int32{
	"DOCKER":     0,
	"SYSTEM":     1,
	"SYSTEM_EXT": 2,
}

func (x ChaincodeDeploymentSpec_ExecutionEnvironment) String() string {
	return proto.EnumName(ChaincodeDeploymentSpec_ExecutionEnvironment_name, int32(x))
}
func (ChaincodeDeploymentSpec_ExecutionEnvironment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{3, 0}
}

// ChaincodeID contains the path as specified by the deploy transaction
// that created it as well as the hashCode that is generated by the
// system for the path. From the user level (ie, CLI, REST API and so on)
// deploy transaction is expected to provide the path and other requests
// are expected to provide the hashCode. The other value will be ignored.
// Internally, the structure could contain both values. For instance, the
// hashCode will be set when first generated using the path
type ChaincodeID struct {
	// deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// all other requests will use the name (really a hashcode) generated by
	// the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// user friendly version name for the chaincode
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *ChaincodeID) Reset()                    { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()               {}
func (*ChaincodeID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ChaincodeID) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ChaincodeID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChaincodeID) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args [][]byte `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

func (m *ChaincodeInput) Reset()                    { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()               {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ChaincodeInput) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type        ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeId *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId" json:"chaincode_id,omitempty"`
	Input       *ChaincodeInput    `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Timeout     int32              `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ChaincodeSpec) Reset()                    { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()               {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ChaincodeSpec) GetType() ChaincodeSpec_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeSpec_UNDEFINED
}

func (m *ChaincodeSpec) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeSpec) GetInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *ChaincodeSpec) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// Controls when the chaincode becomes executable.
	EffectiveDate *google_protobuf1.Timestamp                  `protobuf:"bytes,2,opt,name=effective_date,json=effectiveDate" json:"effective_date,omitempty"`
	CodePackage   []byte                                       `protobuf:"bytes,3,opt,name=code_package,json=codePackage,proto3" json:"code_package,omitempty"`
	ExecEnv       ChaincodeDeploymentSpec_ExecutionEnvironment `protobuf:"varint,4,opt,name=exec_env,json=execEnv,enum=protos.ChaincodeDeploymentSpec_ExecutionEnvironment" json:"exec_env,omitempty"`
}

func (m *ChaincodeDeploymentSpec) Reset()                    { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()               {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetEffectiveDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EffectiveDate
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetCodePackage() []byte {
	if m != nil {
		return m.CodePackage
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetExecEnv() ChaincodeDeploymentSpec_ExecutionEnvironment {
	if m != nil {
		return m.ExecEnv
	}
	return ChaincodeDeploymentSpec_DOCKER
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// This field can contain a user-specified ID generation algorithm
	// If supplied, this will be used to generate a ID
	// If not supplied (left empty), sha256base64 will be used
	// The algorithm consists of two parts:
	//  1, a hash function
	//  2, a decoding used to decode user (string) input to bytes
	// Currently, SHA256 with BASE64 is supported (e.g. idGenerationAlg='sha256base64')
	IdGenerationAlg string `protobuf:"bytes,2,opt,name=id_generation_alg,json=idGenerationAlg" json:"id_generation_alg,omitempty"`
}

func (m *ChaincodeInvocationSpec) Reset()                    { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()               {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeInvocationSpec) GetIdGenerationAlg() string {
	if m != nil {
		return m.IdGenerationAlg
	}
	return ""
}

func init() {
	proto.RegisterType((*ChaincodeID)(nil), "protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "protos.ChaincodeInput")
	proto.RegisterType((*ChaincodeSpec)(nil), "protos.ChaincodeSpec")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "protos.ChaincodeInvocationSpec")
	proto.RegisterEnum("protos.ConfidentialityLevel", ConfidentialityLevel_name, ConfidentialityLevel_value)
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterEnum("protos.ChaincodeDeploymentSpec_ExecutionEnvironment", ChaincodeDeploymentSpec_ExecutionEnvironment_name, ChaincodeDeploymentSpec_ExecutionEnvironment_value)
}

func init() { proto.RegisterFile("peer/chaincode.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x93, 0xf4, 0x6b, 0xf2, 0x81, 0x59, 0x0a, 0x44, 0xbd, 0x50, 0x2c, 0x0e, 0xa5, 0x42,
	0x8e, 0x14, 0x2a, 0x4e, 0x08, 0xc9, 0x8d, 0xdd, 0xca, 0x25, 0xd8, 0x95, 0x9b, 0x22, 0xca, 0x25,
	0x72, 0xec, 0x89, 0xb3, 0xc2, 0xd9, 0xb5, 0xec, 0x8d, 0xd5, 0x9c, 0x39, 0xf2, 0xdf, 0xf8, 0x4d,
	0x68, 0xd7, 0x4d, 0xda, 0xaa, 0x3d, 0x72, 0xf2, 0xec, 0xdb, 0x37, 0xb3, 0x6f, 0x9e, 0x66, 0x0c,
	0x7b, 0x19, 0x62, 0xde, 0x8b, 0x66, 0x21, 0x65, 0x11, 0x8f, 0xd1, 0xcc, 0x72, 0x2e, 0x38, 0xd9,
	0x52, 0x9f, 0x62, 0xff, 0x4d, 0xc2, 0x79, 0x92, 0x62, 0x4f, 0x1d, 0x27, 0x8b, 0x69, 0x4f, 0xd0,
	0x39, 0x16, 0x22, 0x9c, 0x67, 0x15, 0xd1, 0xf0, 0xa1, 0x39, 0x58, 0xe5, 0xba, 0x36, 0x21, 0xd0,
	0xc8, 0x42, 0x31, 0xeb, 0x6a, 0x07, 0xda, 0xe1, 0x6e, 0xa0, 0x62, 0x89, 0xb1, 0x70, 0x8e, 0xdd,
	0x5a, 0x85, 0xc9, 0x98, 0x74, 0x61, 0xbb, 0xc4, 0xbc, 0xa0, 0x9c, 0x75, 0xeb, 0x0a, 0x5e, 0x1d,
	0x8d, 0x77, 0xd0, 0xb9, 0x2b, 0xc8, 0xb2, 0x85, 0x90, 0xf9, 0x61, 0x9e, 0x14, 0x5d, 0xed, 0xa0,
	0x7e, 0xd8, 0x0a, 0x54, 0x6c, 0xfc, 0xa9, 0x41, 0x7b, 0x4d, 0xbb, 0xcc, 0x30, 0x22, 0x26, 0x34,
	0xc4, 0x32, 0x43, 0xf5, 0x72, 0xa7, 0xbf, 0x5f, 0xc9, 0x2b, 0xcc, 0x07, 0x24, 0x73, 0xb4, 0xcc,
	0x30, 0x50, 0x3c, 0xf2, 0x09, 0x5a, 0xeb, 0xa6, 0xc7, 0x34, 0x56, 0xea, 0x9a, 0xfd, 0x17, 0x8f,
	0xf2, 0x5c, 0x3b, 0x68, 0xae, 0x89, 0x6e, 0x4c, 0x3e, 0xc0, 0x26, 0x95, 0xb2, 0x94, 0xee, 0x66,
	0xff, 0xd5, 0xe3, 0x04, 0x79, 0x1b, 0x54, 0x24, 0xd9, 0xa7, 0x74, 0x8c, 0x2f, 0x44, 0xb7, 0x71,
	0xa0, 0x1d, 0x6e, 0x06, 0xab, 0xa3, 0x71, 0x0e, 0x0d, 0xa9, 0x86, 0xb4, 0x61, 0xf7, 0xca, 0xb3,
	0x9d, 0x53, 0xd7, 0x73, 0x6c, 0x7d, 0x83, 0x00, 0x6c, 0x9d, 0xf9, 0x43, 0xcb, 0x3b, 0xd3, 0x35,
	0xb2, 0x03, 0x0d, 0xcf, 0xb7, 0x1d, 0xbd, 0x46, 0xb6, 0xa1, 0x3e, 0xb0, 0x02, 0xbd, 0x2e, 0xa1,
	0x73, 0xeb, 0xbb, 0xa5, 0x37, 0x24, 0xf1, 0xc4, 0xf5, 0xac, 0xe0, 0x5a, 0xdf, 0x34, 0xfe, 0xd6,
	0xe0, 0xf5, 0xfa, 0x7d, 0x1b, 0xb3, 0x94, 0x2f, 0xe7, 0xc8, 0x84, 0xf2, 0xe5, 0x33, 0x74, 0xee,
	0xfa, 0x2c, 0x32, 0x8c, 0x94, 0x43, 0xcd, 0xfe, 0xcb, 0x27, 0x1d, 0x0a, 0xda, 0xd1, 0x03, 0x57,
	0x2d, 0xe8, 0xe0, 0x74, 0x8a, 0x91, 0xa0, 0x25, 0x8e, 0xe3, 0x50, 0xe0, 0xad, 0x4f, 0xfb, 0x66,
	0x35, 0x18, 0xe6, 0x6a, 0x30, 0xcc, 0xd1, 0x6a, 0x30, 0x82, 0xf6, 0x3a, 0xc3, 0x0e, 0x05, 0x92,
	0xb7, 0xd0, 0x52, 0x6f, 0x67, 0x61, 0xf4, 0x2b, 0x4c, 0x50, 0xf9, 0xd6, 0x0a, 0x9a, 0x12, 0xbb,
	0xa8, 0x20, 0xe2, 0xc3, 0x0e, 0xde, 0x60, 0x34, 0x46, 0x56, 0x2a, 0x9b, 0x3a, 0xfd, 0xe3, 0x47,
	0xea, 0x1e, 0xb6, 0x65, 0x3a, 0x37, 0x18, 0x2d, 0x04, 0xe5, 0xcc, 0x61, 0x25, 0xcd, 0x39, 0x93,
	0x17, 0xc1, 0xb6, 0xac, 0xe2, 0xb0, 0xd2, 0xf8, 0x02, 0x7b, 0x4f, 0x11, 0xa4, 0x69, 0xb6, 0x3f,
	0xf8, 0xea, 0x04, 0x95, 0xd3, 0x97, 0xd7, 0x97, 0x23, 0xe7, 0x9b, 0xae, 0x91, 0x0e, 0x40, 0x15,
	0x8f, 0x9d, 0x1f, 0x23, 0xbd, 0x66, 0xfc, 0xd6, 0xee, 0x19, 0xea, 0xb2, 0x92, 0x47, 0xa1, 0x2c,
	0xf5, 0x1f, 0x0c, 0x3d, 0x82, 0xe7, 0x34, 0x1e, 0x27, 0xc8, 0x30, 0x57, 0x25, 0xc7, 0x61, 0x9a,
	0xdc, 0x6e, 0xc6, 0x33, 0x1a, 0x9f, 0xad, 0x71, 0x2b, 0x4d, 0x8e, 0x8e, 0x61, 0x6f, 0xc0, 0xd9,
	0x94, 0xc6, 0xc8, 0x04, 0x0d, 0x53, 0x2a, 0x96, 0x43, 0x2c, 0x31, 0x95, 0xca, 0x2f, 0xae, 0x4e,
	0x86, 0xee, 0x40, 0xdf, 0x20, 0x3a, 0xb4, 0x06, 0xbe, 0x77, 0xea, 0xda, 0x8e, 0x37, 0x72, 0xad,
	0xa1, 0xae, 0x9d, 0xf8, 0x60, 0xf0, 0x3c, 0x31, 0x67, 0xcb, 0x0c, 0xf3, 0x14, 0xe3, 0x04, 0x73,
	0x73, 0x1a, 0x4e, 0x72, 0x1a, 0xad, 0xf4, 0xc9, 0x85, 0xff, 0xf9, 0x3e, 0xa1, 0x62, 0xb6, 0x98,
	0x98, 0x11, 0x9f, 0xf7, 0xee, 0x51, 0x7b, 0x15, 0xb5, 0xda, 0xf7, 0xa2, 0x27, 0xa9, 0x93, 0xea,
	0x5f, 0xf0, 0xf1, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xed, 0xd0, 0x0e, 0x2a, 0x04, 0x00,
	0x00,
}