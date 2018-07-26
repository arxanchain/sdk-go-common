// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet/colored_token.proto

/*
Package wallet is a generated protocol buffer package.

It is generated from these files:
	wallet/colored_token.proto
	wallet/common.proto
	wallet/didentity.proto
	wallet/pki.proto
	wallet/query.proto
	wallet/services.proto
	wallet/tx.proto
	wallet/verifiable_claim.proto

It has these top-level messages:
	ColoredToken
	Metadata
	DidOwner
	DidControl
	OffchainMetadata
	DDO
	PublicKey
	PrivateKey
	SignatureHeader
	Signature
	SignatureSimple
	AssetsGrowth
	UsersGrowth
	DAppAxtConsume
	DAppAxtTotal
	DAppIssuedToken
	DAppIssuedTokens
	DAppList
	DAppNum
	DAppTotal
	DAppUsersList
	DAppUsersTotal
	HotAssets
	HotTokens
	TopAssetUsers
	TopTokenUsers
	Total
	TxIN
	TxOUT
	UTXO
	TX
	SingleColorResult
	QueryTxRequest
	QueryUTXORequest
	QueryUTXOResponse
	TxScript
	ProcessTxResponse
	UTXOSignature
	VerifiableClaim
	VerifiableClaimBrief
*/
package wallet

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

type CTStatus int32

const (
	CTStatus_INUSE   CTStatus = 0
	CTStatus_REVOKED CTStatus = 1
	CTStatus_EXPIRED CTStatus = 2
)

var CTStatus_name = map[int32]string{
	0: "INUSE",
	1: "REVOKED",
	2: "EXPIRED",
}
var CTStatus_value = map[string]int32{
	"INUSE":   0,
	"REVOKED": 1,
	"EXPIRED": 2,
}

func (x CTStatus) String() string {
	return proto.EnumName(CTStatus_name, int32(x))
}
func (CTStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ColoredToken struct {
	// ColoredToken ID
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// ColoredToken issued base on which ancestor
	// this should be a didentity
	Ancestor string `protobuf:"bytes,2,opt,name=ancestor" json:"ancestor,omitempty"`
	// issue timestamp
	IssueTime int64 `protobuf:"varint,3,opt,name=issueTime" json:"issueTime,omitempty"`
	// expire timestamp, -1 indicates no expire
	ExpireTime int64 `protobuf:"varint,4,opt,name=expireTime" json:"expireTime,omitempty"`
	// ColoredToken status
	Status CTStatus `protobuf:"varint,5,opt,name=status,enum=wallet.CTStatus" json:"status,omitempty"`
	// issuers didentity and signature map
	Issuers map[string]*Signature `protobuf:"bytes,6,rep,name=issuers" json:"issuers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// endorsers didentity and signature map
	Endorsers map[string]*Signature `protobuf:"bytes,7,rep,name=endorsers" json:"endorsers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// metadata
	Metadata []byte `protobuf:"bytes,8,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *ColoredToken) Reset()                    { *m = ColoredToken{} }
func (m *ColoredToken) String() string            { return proto.CompactTextString(m) }
func (*ColoredToken) ProtoMessage()               {}
func (*ColoredToken) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ColoredToken) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ColoredToken) GetAncestor() string {
	if m != nil {
		return m.Ancestor
	}
	return ""
}

func (m *ColoredToken) GetIssueTime() int64 {
	if m != nil {
		return m.IssueTime
	}
	return 0
}

func (m *ColoredToken) GetExpireTime() int64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *ColoredToken) GetStatus() CTStatus {
	if m != nil {
		return m.Status
	}
	return CTStatus_INUSE
}

func (m *ColoredToken) GetIssuers() map[string]*Signature {
	if m != nil {
		return m.Issuers
	}
	return nil
}

func (m *ColoredToken) GetEndorsers() map[string]*Signature {
	if m != nil {
		return m.Endorsers
	}
	return nil
}

func (m *ColoredToken) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*ColoredToken)(nil), "wallet.ColoredToken")
	proto.RegisterEnum("wallet.CTStatus", CTStatus_name, CTStatus_value)
}

func init() { proto.RegisterFile("wallet/colored_token.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x5f, 0xab, 0xd3, 0x30,
	0x14, 0xb7, 0xad, 0xeb, 0xd6, 0xb3, 0x31, 0x6a, 0x9e, 0x4a, 0x11, 0xa9, 0xfa, 0x60, 0x11, 0x6e,
	0x8b, 0xbb, 0x2f, 0xa2, 0x4f, 0x7a, 0x6f, 0x1f, 0x86, 0xe8, 0x95, 0x6c, 0x8a, 0xf8, 0x22, 0xb9,
	0x6d, 0xd8, 0x42, 0xdb, 0x64, 0x24, 0xa9, 0x6e, 0x5f, 0xd6, 0xcf, 0x22, 0x4d, 0xd6, 0x6d, 0x82,
	0x6f, 0xbe, 0xe5, 0xf7, 0xef, 0x9c, 0xc3, 0x8f, 0x40, 0xfc, 0x8b, 0x34, 0x0d, 0xd5, 0x79, 0x29,
	0x1a, 0x21, 0x69, 0xf5, 0x43, 0x8b, 0x9a, 0xf2, 0x6c, 0x27, 0x85, 0x16, 0xc8, 0xb7, 0x5a, 0x1c,
	0x1e, 0x3d, 0xbb, 0x9a, 0x59, 0xe5, 0xd9, 0x6f, 0x0f, 0x66, 0x37, 0x36, 0xb1, 0xee, 0x03, 0x68,
	0x0e, 0x2e, 0xab, 0x22, 0x27, 0x71, 0xd2, 0x00, 0xbb, 0xac, 0x42, 0x31, 0x4c, 0x08, 0x2f, 0xa9,
	0xd2, 0x42, 0x46, 0xae, 0x61, 0x4f, 0x18, 0x3d, 0x86, 0x80, 0x29, 0xd5, 0xd1, 0x35, 0x6b, 0x69,
	0xe4, 0x25, 0x4e, 0xea, 0xe1, 0x33, 0x81, 0x9e, 0x00, 0xd0, 0xfd, 0x8e, 0x49, 0x2b, 0x3f, 0x34,
	0xf2, 0x05, 0x83, 0x52, 0xf0, 0x95, 0x26, 0xba, 0x53, 0xd1, 0x28, 0x71, 0xd2, 0xf9, 0x22, 0xcc,
	0xec, 0x75, 0xd9, 0xcd, 0x7a, 0x65, 0x78, 0x7c, 0xd4, 0xd1, 0x5b, 0x18, 0x9b, 0xb1, 0x52, 0x45,
	0x7e, 0xe2, 0xa5, 0xd3, 0xc5, 0xd3, 0x93, 0xf5, 0xe2, 0xf4, 0x6c, 0x69, 0x3d, 0x05, 0xd7, 0xf2,
	0x80, 0x87, 0x04, 0x7a, 0x07, 0x01, 0xe5, 0x95, 0x90, 0xaa, 0x8f, 0x8f, 0x4d, 0xfc, 0xf9, 0x3f,
	0xe3, 0xc5, 0xe0, 0xb2, 0x03, 0xce, 0xa9, 0xbe, 0x83, 0x96, 0x6a, 0x52, 0x11, 0x4d, 0xa2, 0x49,
	0xe2, 0xa4, 0x33, 0x7c, 0xc2, 0xf1, 0x47, 0x98, 0x5d, 0xee, 0x45, 0x21, 0x78, 0x35, 0x3d, 0x1c,
	0x0b, 0xec, 0x9f, 0xe8, 0x05, 0x8c, 0x7e, 0x92, 0xa6, 0xa3, 0xa6, 0xbe, 0xe9, 0xe2, 0xd1, 0xb0,
	0x7c, 0xc5, 0x36, 0x9c, 0xe8, 0x4e, 0x52, 0x6c, 0xf5, 0x37, 0xee, 0x6b, 0x27, 0xbe, 0x83, 0xf9,
	0xdf, 0x77, 0xfc, 0xe7, 0xc0, 0x97, 0x39, 0x4c, 0x86, 0x3e, 0x51, 0x00, 0xa3, 0xe5, 0xa7, 0x2f,
	0xab, 0x22, 0x7c, 0x80, 0xa6, 0x30, 0xc6, 0xc5, 0xd7, 0xbb, 0x0f, 0xc5, 0x6d, 0xe8, 0xf4, 0xa0,
	0xf8, 0xf6, 0x79, 0x89, 0x8b, 0xdb, 0xd0, 0x7d, 0x7f, 0xfd, 0xfd, 0xd5, 0x86, 0xe9, 0x6d, 0x77,
	0x9f, 0x95, 0xa2, 0xcd, 0x89, 0xdc, 0x13, 0x5e, 0x6e, 0x09, 0xe3, 0xb9, 0xaa, 0xea, 0xab, 0x8d,
	0xb8, 0x2a, 0x45, 0xdb, 0x0a, 0x9e, 0x9b, 0xef, 0xa3, 0x72, 0xbb, 0xfa, 0xde, 0x37, 0xf0, 0xfa,
	0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x8b, 0x1d, 0x65, 0x85, 0x02, 0x00, 0x00,
}
