// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet/verifiable_claim.proto

package wallet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VerifiableClaim struct {
	Id         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Status     Status                     `protobuf:"varint,2,opt,name=status,enum=wallet.Status" json:"status,omitempty"`
	Name       string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type       string                     `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Metadata   *Metadata                  `protobuf:"bytes,5,opt,name=metadata" json:"metadata,omitempty"`
	Holder     string                     `protobuf:"bytes,6,opt,name=holder" json:"holder,omitempty"`
	Issuer     string                     `protobuf:"bytes,7,opt,name=issuer" json:"issuer,omitempty"`
	Signatures []*Signature               `protobuf:"bytes,8,rep,name=signatures" json:"signatures,omitempty"`
	Created    *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=created" json:"created,omitempty"`
	Updated    *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=updated" json:"updated,omitempty"`
	Expires    *google_protobuf.Timestamp `protobuf:"bytes,11,opt,name=expires" json:"expires,omitempty"`
}

func (m *VerifiableClaim) Reset()                    { *m = VerifiableClaim{} }
func (m *VerifiableClaim) String() string            { return proto.CompactTextString(m) }
func (*VerifiableClaim) ProtoMessage()               {}
func (*VerifiableClaim) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *VerifiableClaim) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VerifiableClaim) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_INVALID
}

func (m *VerifiableClaim) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VerifiableClaim) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VerifiableClaim) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *VerifiableClaim) GetHolder() string {
	if m != nil {
		return m.Holder
	}
	return ""
}

func (m *VerifiableClaim) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *VerifiableClaim) GetSignatures() []*Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *VerifiableClaim) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *VerifiableClaim) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *VerifiableClaim) GetExpires() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expires
	}
	return nil
}

type VerifiableClaimBrief struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Status Status `protobuf:"varint,2,opt,name=status,enum=wallet.Status" json:"status,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
}

func (m *VerifiableClaimBrief) Reset()                    { *m = VerifiableClaimBrief{} }
func (m *VerifiableClaimBrief) String() string            { return proto.CompactTextString(m) }
func (*VerifiableClaimBrief) ProtoMessage()               {}
func (*VerifiableClaimBrief) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *VerifiableClaimBrief) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VerifiableClaimBrief) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_INVALID
}

func (m *VerifiableClaimBrief) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VerifiableClaimBrief) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifiableClaim)(nil), "wallet.VerifiableClaim")
	proto.RegisterType((*VerifiableClaimBrief)(nil), "wallet.VerifiableClaimBrief")
}

func init() { proto.RegisterFile("wallet/verifiable_claim.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x6e, 0xe2, 0x30,
	0x10, 0x86, 0x15, 0x60, 0x03, 0x18, 0x89, 0x65, 0xbd, 0xab, 0x95, 0x85, 0x54, 0x35, 0xe2, 0x50,
	0xe5, 0x50, 0x12, 0x01, 0x7d, 0x02, 0x7a, 0xee, 0x25, 0xad, 0x7a, 0xe8, 0xa5, 0x72, 0x92, 0x21,
	0x58, 0xc4, 0x71, 0x64, 0x3b, 0x94, 0xbe, 0x4f, 0x1f, 0xb4, 0x8a, 0xed, 0xa0, 0xaa, 0x17, 0x4e,
	0xbd, 0x79, 0xbe, 0xf9, 0xff, 0x99, 0xd1, 0x78, 0xd0, 0xd5, 0x1b, 0x2d, 0x4b, 0xd0, 0xf1, 0x11,
	0x24, 0xdb, 0x31, 0x9a, 0x96, 0xf0, 0x9a, 0x95, 0x94, 0xf1, 0xa8, 0x96, 0x42, 0x0b, 0xec, 0xdb,
	0xf4, 0xfc, 0xba, 0x10, 0xa2, 0x28, 0x21, 0x36, 0x34, 0x6d, 0x76, 0xb1, 0x66, 0x1c, 0x94, 0xa6,
	0xbc, 0xb6, 0xc2, 0xf9, 0x5f, 0x57, 0x27, 0x13, 0x9c, 0x8b, 0xca, 0xc1, 0x99, 0x83, 0xf5, 0x81,
	0x59, 0xb2, 0xf8, 0xe8, 0xa3, 0xdf, 0xcf, 0xe7, 0x56, 0xf7, 0x6d, 0x27, 0x3c, 0x45, 0x3d, 0x96,
	0x13, 0x2f, 0xf0, 0xc2, 0x71, 0xd2, 0x63, 0x39, 0xbe, 0x41, 0xbe, 0xd2, 0x54, 0x37, 0x8a, 0xf4,
	0x02, 0x2f, 0x9c, 0xae, 0xa7, 0x91, 0x2d, 0x13, 0x3d, 0x1a, 0x9a, 0xb8, 0x2c, 0xc6, 0x68, 0x50,
	0x51, 0x0e, 0xa4, 0x6f, 0x9c, 0xe6, 0xdd, 0x32, 0xfd, 0x5e, 0x03, 0x19, 0x58, 0xd6, 0xbe, 0xf1,
	0x2d, 0x1a, 0x71, 0xd0, 0x34, 0xa7, 0x9a, 0x92, 0x5f, 0x81, 0x17, 0x4e, 0xd6, 0xb3, 0xae, 0xe2,
	0x83, 0xe3, 0xc9, 0x59, 0x81, 0xff, 0x23, 0x7f, 0x2f, 0xca, 0x1c, 0x24, 0xf1, 0x4d, 0x0d, 0x17,
	0xb5, 0x9c, 0x29, 0xd5, 0x80, 0x24, 0x43, 0xcb, 0x6d, 0x84, 0x57, 0x08, 0x29, 0x56, 0x54, 0x54,
	0x37, 0x12, 0x14, 0x19, 0x05, 0xfd, 0x70, 0xb2, 0xfe, 0x73, 0x9e, 0xb8, 0xcb, 0x24, 0x5f, 0x44,
	0xf8, 0x0e, 0x0d, 0x33, 0x09, 0x54, 0x43, 0x4e, 0xc6, 0x66, 0x9e, 0x79, 0x64, 0xd7, 0x1b, 0x75,
	0xeb, 0x8d, 0x9e, 0xba, 0xf5, 0x26, 0x9d, 0xb4, 0x75, 0x35, 0x75, 0x6e, 0x5c, 0xe8, 0xb2, 0xcb,
	0x49, 0x5b, 0x17, 0x9c, 0x6a, 0xd6, 0xce, 0x36, 0xb9, 0xec, 0x72, 0xd2, 0xc5, 0x11, 0xfd, 0xfb,
	0xf6, 0x4b, 0x5b, 0xc9, 0x60, 0xf7, 0xd3, 0x5f, 0xb5, 0xdd, 0xbc, 0xac, 0x0a, 0xa6, 0xf7, 0x4d,
	0x1a, 0x65, 0x82, 0xc7, 0x54, 0x9e, 0x68, 0x95, 0xed, 0x29, 0xab, 0x62, 0x95, 0x1f, 0x96, 0x85,
	0x58, 0xda, 0xeb, 0xb2, 0x57, 0xa8, 0x62, 0xdb, 0x2b, 0xf5, 0x4d, 0xb8, 0xf9, 0x0c, 0x00, 0x00,
	0xff, 0xff, 0x69, 0x35, 0x5b, 0x0a, 0xcb, 0x02, 0x00, 0x00,
}
