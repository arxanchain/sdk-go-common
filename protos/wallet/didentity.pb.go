// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet/didentity.proto

package wallet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DidType decentralized identity types
// Decentralized Identity
// ├── Asset: 资产类型
//    └── SerialNumber: 序列号资产类型
// └── Entity: 账户类型
//    ├── DApp: DAPP账户类型
//    ├── APPChain: 应用链账户类型
//    ├── Organization: 组织账户类型
//    └── Person: 个人账户类型
//        ├── Dependent: 非独立个人账户类型
//        └── Independent: 独立个人账户类型
type DidType int32

const (
	DidType_ASSET        DidType = 0
	DidType_ENTITY       DidType = 1
	DidType_ORGANIZATION DidType = 2
	DidType_PERSON       DidType = 3
	DidType_INDEPENDENT  DidType = 4
	DidType_DEPENDENT    DidType = 5
	DidType_SERIALNUMBER DidType = 6
	DidType_SWCASH       DidType = 11
	DidType_SWFEE        DidType = 12
	DidType_SWLOAN       DidType = 13
	DidType_SWINTEREST   DidType = 14
	DidType_DAPP         DidType = 21
	DidType_APPCHAIN     DidType = 22
)

var DidType_name = map[int32]string{
	0:  "ASSET",
	1:  "ENTITY",
	2:  "ORGANIZATION",
	3:  "PERSON",
	4:  "INDEPENDENT",
	5:  "DEPENDENT",
	6:  "SERIALNUMBER",
	11: "SWCASH",
	12: "SWFEE",
	13: "SWLOAN",
	14: "SWINTEREST",
	21: "DAPP",
	22: "APPCHAIN",
}
var DidType_value = map[string]int32{
	"ASSET":        0,
	"ENTITY":       1,
	"ORGANIZATION": 2,
	"PERSON":       3,
	"INDEPENDENT":  4,
	"DEPENDENT":    5,
	"SERIALNUMBER": 6,
	"SWCASH":       11,
	"SWFEE":        12,
	"SWLOAN":       13,
	"SWINTEREST":   14,
	"DAPP":         21,
	"APPCHAIN":     22,
}

func (x DidType) String() string {
	return proto.EnumName(DidType_name, int32(x))
}
func (DidType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// DidControlType defines the control type of decentralized identity controllers
// or, and, m of n
type DidControlType int32

const (
	DidControlType_OR   DidControlType = 0
	DidControlType_AND  DidControlType = 1
	DidControlType_MOFN DidControlType = 2
)

var DidControlType_name = map[int32]string{
	0: "OR",
	1: "AND",
	2: "MOFN",
}
var DidControlType_value = map[string]int32{
	"OR":   0,
	"AND":  1,
	"MOFN": 2,
}

func (x DidControlType) String() string {
	return proto.EnumName(DidControlType_name, int32(x))
}
func (DidControlType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// DidOwner defines owner information of decentralized identity
type DidOwner struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type      []string                   `protobuf:"bytes,2,rep,name=type" json:"type,omitempty"`
	Expires   *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=expires" json:"expires,omitempty"`
	PublicKey *PublicKey                 `protobuf:"bytes,4,opt,name=publicKey" json:"publicKey,omitempty"`
}

func (m *DidOwner) Reset()                    { *m = DidOwner{} }
func (m *DidOwner) String() string            { return proto.CompactTextString(m) }
func (*DidOwner) ProtoMessage()               {}
func (*DidOwner) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *DidOwner) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DidOwner) GetType() []string {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DidOwner) GetExpires() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expires
	}
	return nil
}

func (m *DidOwner) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

// DidControl defines controllers information of decentralized identity
type DidControl struct {
	Type              DidControlType `protobuf:"varint,1,opt,name=type,enum=wallet.DidControlType" json:"type,omitempty"`
	MinimumSignatures int32          `protobuf:"varint,2,opt,name=minimumSignatures" json:"minimumSignatures,omitempty"`
	Signers           []string       `protobuf:"bytes,3,rep,name=signers" json:"signers,omitempty"`
}

func (m *DidControl) Reset()                    { *m = DidControl{} }
func (m *DidControl) String() string            { return proto.CompactTextString(m) }
func (*DidControl) ProtoMessage()               {}
func (*DidControl) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *DidControl) GetType() DidControlType {
	if m != nil {
		return m.Type
	}
	return DidControlType_OR
}

func (m *DidControl) GetMinimumSignatures() int32 {
	if m != nil {
		return m.MinimumSignatures
	}
	return 0
}

func (m *DidControl) GetSigners() []string {
	if m != nil {
		return m.Signers
	}
	return nil
}

// OffchainMetadata defines offchain metadata of decentralized identity
type OffchainMetadata struct {
	Filename    string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Endpoint    string `protobuf:"bytes,2,opt,name=endpoint" json:"endpoint,omitempty"`
	StorageType string `protobuf:"bytes,3,opt,name=storageType" json:"storageType,omitempty"`
	ContentHash string `protobuf:"bytes,4,opt,name=contentHash" json:"contentHash,omitempty"`
	Size        int32  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
	ReadOnly    bool   `protobuf:"varint,6,opt,name=readOnly" json:"readOnly,omitempty"`
}

func (m *OffchainMetadata) Reset()                    { *m = OffchainMetadata{} }
func (m *OffchainMetadata) String() string            { return proto.CompactTextString(m) }
func (*OffchainMetadata) ProtoMessage()               {}
func (*OffchainMetadata) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *OffchainMetadata) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *OffchainMetadata) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *OffchainMetadata) GetStorageType() string {
	if m != nil {
		return m.StorageType
	}
	return ""
}

func (m *OffchainMetadata) GetContentHash() string {
	if m != nil {
		return m.ContentHash
	}
	return ""
}

func (m *OffchainMetadata) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *OffchainMetadata) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

// DDO defines an object to describe the decentralized identity
type DDO struct {
	Context          string                           `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	Id               string                           `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Type             []DidType                        `protobuf:"varint,3,rep,packed,name=type,enum=wallet.DidType" json:"type,omitempty"`
	Owners           []*DidOwner                      `protobuf:"bytes,4,rep,name=owners" json:"owners,omitempty"`
	Guardian         string                           `protobuf:"bytes,5,opt,name=guardian" json:"guardian,omitempty"`
	Controls         []*DidControl                    `protobuf:"bytes,6,rep,name=controls" json:"controls,omitempty"`
	Claims           map[string]*VerifiableClaimBrief `protobuf:"bytes,7,rep,name=claims" json:"claims,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Services         map[string]string                `protobuf:"bytes,8,rep,name=services" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Endpoint         string                           `protobuf:"bytes,9,opt,name=endpoint" json:"endpoint,omitempty"`
	Metadata         *Metadata                        `protobuf:"bytes,10,opt,name=metadata" json:"metadata,omitempty"`
	Created          *google_protobuf.Timestamp       `protobuf:"bytes,11,opt,name=created" json:"created,omitempty"`
	Updated          *google_protobuf.Timestamp       `protobuf:"bytes,12,opt,name=updated" json:"updated,omitempty"`
	Expires          *google_protobuf.Timestamp       `protobuf:"bytes,13,opt,name=expires" json:"expires,omitempty"`
	Signatures       []*Signature                     `protobuf:"bytes,14,rep,name=signatures" json:"signatures,omitempty"`
	Status           Status                           `protobuf:"varint,15,opt,name=status,enum=wallet.Status" json:"status,omitempty"`
	Name             string                           `protobuf:"bytes,16,opt,name=name" json:"name,omitempty"`
	ParentId         string                           `protobuf:"bytes,17,opt,name=parentId" json:"parentId,omitempty"`
	OffchainMetadata *OffchainMetadata                `protobuf:"bytes,18,opt,name=offchainMetadata" json:"offchainMetadata,omitempty"`
}

func (m *DDO) Reset()                    { *m = DDO{} }
func (m *DDO) String() string            { return proto.CompactTextString(m) }
func (*DDO) ProtoMessage()               {}
func (*DDO) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *DDO) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *DDO) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DDO) GetType() []DidType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DDO) GetOwners() []*DidOwner {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *DDO) GetGuardian() string {
	if m != nil {
		return m.Guardian
	}
	return ""
}

func (m *DDO) GetControls() []*DidControl {
	if m != nil {
		return m.Controls
	}
	return nil
}

func (m *DDO) GetClaims() map[string]*VerifiableClaimBrief {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *DDO) GetServices() map[string]string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *DDO) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *DDO) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *DDO) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *DDO) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *DDO) GetExpires() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expires
	}
	return nil
}

func (m *DDO) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *DDO) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_INVALID
}

func (m *DDO) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DDO) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *DDO) GetOffchainMetadata() *OffchainMetadata {
	if m != nil {
		return m.OffchainMetadata
	}
	return nil
}

func init() {
	proto.RegisterType((*DidOwner)(nil), "wallet.DidOwner")
	proto.RegisterType((*DidControl)(nil), "wallet.DidControl")
	proto.RegisterType((*OffchainMetadata)(nil), "wallet.OffchainMetadata")
	proto.RegisterType((*DDO)(nil), "wallet.DDO")
	proto.RegisterEnum("wallet.DidType", DidType_name, DidType_value)
	proto.RegisterEnum("wallet.DidControlType", DidControlType_name, DidControlType_value)
}

func init() { proto.RegisterFile("wallet/didentity.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 918 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x17, 0x5e, 0x27, 0x69, 0x12, 0x9f, 0xb4, 0xe9, 0x74, 0xde, 0x97, 0x62, 0x22, 0x10, 0x51, 0x91,
	0x50, 0x54, 0x76, 0x13, 0x6d, 0x16, 0x24, 0x04, 0x57, 0x69, 0xe3, 0xa5, 0x11, 0x5b, 0x3b, 0x1a,
	0x07, 0x2a, 0xf6, 0x06, 0x4d, 0xe2, 0x49, 0x3a, 0xaa, 0xbf, 0x64, 0x4f, 0xda, 0x86, 0x2b, 0x6e,
	0xf8, 0x09, 0xfc, 0x1a, 0x24, 0x7e, 0x1b, 0x9a, 0xf1, 0xd8, 0x75, 0x5b, 0xa4, 0xe5, 0x6e, 0xce,
	0x39, 0xcf, 0x33, 0x73, 0xbe, 0x1e, 0x1b, 0x8e, 0xef, 0x68, 0x10, 0x30, 0x31, 0xf2, 0xb9, 0xcf,
	0x22, 0xc1, 0xc5, 0x6e, 0x98, 0xa4, 0xb1, 0x88, 0x71, 0x33, 0xf7, 0xf7, 0x3e, 0xdf, 0xc4, 0xf1,
	0x26, 0x60, 0x23, 0xe5, 0x5d, 0x6e, 0xd7, 0x23, 0xc1, 0x43, 0x96, 0x09, 0x1a, 0x26, 0x39, 0xb0,
	0xf7, 0x3f, 0x7d, 0xc1, 0x2a, 0x0e, 0xc3, 0x38, 0xd2, 0x4e, 0xa4, 0x9d, 0xc9, 0x0d, 0xd7, 0x9e,
	0xcf, 0xb4, 0xe7, 0x96, 0xa5, 0x7c, 0xcd, 0xe9, 0x32, 0x60, 0xbf, 0xae, 0x02, 0xca, 0xc3, 0x3c,
	0x7c, 0xf2, 0xa7, 0x01, 0xed, 0x29, 0xf7, 0xdd, 0xbb, 0x88, 0xa5, 0xb8, 0x0b, 0x35, 0xee, 0x5b,
	0x46, 0xdf, 0x18, 0x98, 0xa4, 0xc6, 0x7d, 0x8c, 0xa1, 0x21, 0x76, 0x09, 0xb3, 0x6a, 0xfd, 0xfa,
	0xc0, 0x24, 0xea, 0x8c, 0xbf, 0x86, 0x16, 0xbb, 0x4f, 0x78, 0xca, 0x32, 0xab, 0xde, 0x37, 0x06,
	0x9d, 0x71, 0x6f, 0x98, 0x67, 0x3a, 0x2c, 0x32, 0x1d, 0x2e, 0x8a, 0x4c, 0x49, 0x01, 0xc5, 0x23,
	0x30, 0x93, 0xed, 0x32, 0xe0, 0xab, 0x1f, 0xd9, 0xce, 0x6a, 0x28, 0xde, 0xd1, 0x30, 0xcf, 0x6c,
	0x38, 0x2f, 0x02, 0xe4, 0x01, 0x73, 0xf2, 0xbb, 0x01, 0x30, 0xe5, 0xfe, 0x79, 0x1c, 0x89, 0x34,
	0x0e, 0xf0, 0xa9, 0xce, 0x44, 0xe6, 0xd6, 0x1d, 0x1f, 0x17, 0xd4, 0x07, 0xc4, 0x62, 0x97, 0x30,
	0x9d, 0xe1, 0x4b, 0x38, 0x0a, 0x79, 0xc4, 0xc3, 0x6d, 0xe8, 0xf1, 0x4d, 0x44, 0xc5, 0x56, 0xe6,
	0x5a, 0xeb, 0x1b, 0x83, 0x3d, 0xf2, 0x3c, 0x80, 0x2d, 0x68, 0x65, 0x7c, 0x13, 0xb1, 0x54, 0xd6,
	0x23, 0xcb, 0x2c, 0xcc, 0x93, 0xbf, 0x0d, 0x40, 0xee, 0x7a, 0xbd, 0xba, 0xa6, 0x3c, 0xba, 0x64,
	0x82, 0xfa, 0x54, 0x50, 0xdc, 0x83, 0xf6, 0x9a, 0x07, 0x2c, 0xa2, 0x21, 0xd3, 0x8d, 0x2a, 0x6d,
	0x19, 0x63, 0x91, 0x9f, 0xc4, 0x3c, 0x12, 0xea, 0x3d, 0x93, 0x94, 0x36, 0xee, 0x43, 0x27, 0x13,
	0x71, 0x4a, 0x37, 0x4c, 0x66, 0xaa, 0x5a, 0x67, 0x92, 0xaa, 0x4b, 0x22, 0x56, 0x71, 0x24, 0x58,
	0x24, 0x2e, 0x68, 0x76, 0xad, 0x9a, 0x64, 0x92, 0xaa, 0x4b, 0x8e, 0x23, 0xe3, 0xbf, 0x31, 0x6b,
	0x4f, 0xd5, 0xa2, 0xce, 0xf2, 0xcd, 0x94, 0x51, 0xdf, 0x8d, 0x82, 0x9d, 0xd5, 0xec, 0x1b, 0x83,
	0x36, 0x29, 0xed, 0x93, 0x3f, 0x5a, 0x50, 0x9f, 0x4e, 0x5d, 0x59, 0xa2, 0xba, 0xe6, 0x5e, 0xe8,
	0x94, 0x0b, 0x53, 0x0f, 0xbc, 0x56, 0x0e, 0xfc, 0x0b, 0xdd, 0x66, 0xd9, 0x89, 0xee, 0xf8, 0xb0,
	0xd2, 0xe6, 0x4a, 0x7f, 0x07, 0xd0, 0x8c, 0xef, 0x54, 0xc3, 0x1a, 0xfd, 0xfa, 0xa0, 0x33, 0x46,
	0x15, 0x98, 0xda, 0x23, 0xa2, 0xe3, 0x32, 0xb9, 0xcd, 0x96, 0xa6, 0x3e, 0xa7, 0x91, 0x4a, 0xda,
	0x24, 0xa5, 0x8d, 0x87, 0xd0, 0x5e, 0xe5, 0xa3, 0xcb, 0xac, 0xa6, 0xba, 0x07, 0x3f, 0x9f, 0x2a,
	0x29, 0x31, 0x78, 0x04, 0x4d, 0xb5, 0xb7, 0x99, 0xd5, 0x52, 0xe8, 0x8f, 0x4b, 0xf4, 0xd4, 0x1d,
	0x9e, 0xab, 0x88, 0x1d, 0x89, 0x74, 0x47, 0x34, 0x0c, 0x7f, 0x03, 0xed, 0x8c, 0xa5, 0xb7, 0x7c,
	0xc5, 0x32, 0xab, 0xad, 0x28, 0x9f, 0x54, 0x29, 0x9e, 0x8e, 0xe5, 0xa4, 0x12, 0xfa, 0x68, 0x88,
	0xe6, 0x93, 0x21, 0xbe, 0x84, 0x76, 0xa8, 0x17, 0xc1, 0x02, 0xb5, 0xc4, 0x65, 0xed, 0xc5, 0x82,
	0x90, 0x12, 0x21, 0x95, 0xb2, 0x4a, 0x19, 0x15, 0xcc, 0xb7, 0x3a, 0x1f, 0x56, 0x8a, 0x86, 0x4a,
	0xd6, 0x36, 0xf1, 0x15, 0x6b, 0xff, 0xc3, 0x2c, 0x0d, 0xad, 0xaa, 0xf2, 0xe0, 0xbf, 0xab, 0xf2,
	0x35, 0x40, 0xf6, 0x20, 0x91, 0xae, 0x6a, 0x52, 0x29, 0xcb, 0x52, 0x23, 0xa4, 0x02, 0xc2, 0x5f,
	0x42, 0x33, 0x13, 0x54, 0x6c, 0x33, 0xeb, 0x50, 0x49, 0xb1, 0x5b, 0xc2, 0x95, 0x97, 0xe8, 0xa8,
	0xdc, 0x55, 0xa5, 0x11, 0xa4, 0x5a, 0xd8, 0x28, 0xf4, 0x91, 0xd0, 0x94, 0x45, 0x62, 0xe6, 0x5b,
	0x47, 0x79, 0x6b, 0x0b, 0x1b, 0x4f, 0x01, 0xc5, 0x4f, 0xb4, 0x66, 0x61, 0x55, 0x89, 0x55, 0xbc,
	0xf0, 0x54, 0x8b, 0xe4, 0x19, 0xa3, 0x77, 0x05, 0x9d, 0xca, 0x2a, 0x60, 0x04, 0xf5, 0x1b, 0xb6,
	0xd3, 0x4b, 0x2f, 0x8f, 0x78, 0x0c, 0x7b, 0xb7, 0x34, 0xd8, 0x32, 0xb5, 0xf3, 0x9d, 0xf1, 0xa7,
	0xc5, 0xdd, 0x3f, 0x97, 0x5f, 0x47, 0xc5, 0x3f, 0x4b, 0x39, 0x5b, 0x93, 0x1c, 0xfa, 0x5d, 0xed,
	0x5b, 0xa3, 0xf7, 0x3d, 0x1c, 0x3c, 0x5a, 0x98, 0x7f, 0xb9, 0xfa, 0xff, 0xd5, 0xab, 0xcd, 0x0a,
	0xf9, 0xf4, 0x2f, 0x03, 0x5a, 0x5a, 0x42, 0xd8, 0x84, 0xbd, 0x89, 0xe7, 0xd9, 0x0b, 0xf4, 0x02,
	0x03, 0x34, 0x6d, 0x67, 0x31, 0x5b, 0xfc, 0x82, 0x0c, 0x8c, 0x60, 0xdf, 0x25, 0x3f, 0x4c, 0x9c,
	0xd9, 0xfb, 0xc9, 0x62, 0xe6, 0x3a, 0xa8, 0x26, 0xa3, 0x73, 0x9b, 0x78, 0xae, 0x83, 0xea, 0xf8,
	0x10, 0x3a, 0x33, 0x67, 0x6a, 0xcf, 0x6d, 0x67, 0x6a, 0x3b, 0x0b, 0xd4, 0xc0, 0x07, 0x60, 0x3e,
	0x98, 0x7b, 0x92, 0xed, 0xd9, 0x64, 0x36, 0x79, 0xe7, 0xfc, 0x74, 0x79, 0x66, 0x13, 0xd4, 0x94,
	0x6c, 0xef, 0xea, 0x7c, 0xe2, 0x5d, 0xa0, 0x8e, 0x7c, 0xd2, 0xbb, 0x7a, 0x6b, 0xdb, 0x68, 0x3f,
	0x77, 0xbf, 0x73, 0x27, 0x0e, 0x3a, 0xc0, 0x5d, 0x00, 0xef, 0x6a, 0xe6, 0x2c, 0x6c, 0x62, 0x7b,
	0x0b, 0xd4, 0xc5, 0x6d, 0x68, 0x4c, 0x27, 0xf3, 0x39, 0xfa, 0x08, 0xef, 0x43, 0x7b, 0x32, 0x9f,
	0x9f, 0x5f, 0x4c, 0x66, 0x0e, 0x3a, 0x3e, 0xfd, 0x0a, 0xba, 0x8f, 0x3f, 0xb3, 0xb8, 0x09, 0x35,
	0x97, 0xa0, 0x17, 0xb8, 0x05, 0xf5, 0x89, 0x33, 0x45, 0x86, 0xa4, 0x5e, 0xba, 0x6f, 0x1d, 0x54,
	0x3b, 0x7b, 0xf3, 0xfe, 0xf5, 0x86, 0x8b, 0xeb, 0xed, 0x72, 0xb8, 0x8a, 0xc3, 0x11, 0x4d, 0xef,
	0x69, 0xa4, 0x26, 0x34, 0xca, 0xfc, 0x9b, 0x57, 0x9b, 0xf8, 0x55, 0xfe, 0xb3, 0xca, 0x7f, 0x6a,
	0xd9, 0x28, 0x6f, 0xfe, 0xb2, 0xa9, 0xcc, 0x37, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xda, 0x75,
	0x14, 0x23, 0x13, 0x07, 0x00, 0x00,
}
