/*
Copyright ArxanFintech Technology Ltd. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package pki

import (
	"encoding/json"

	commdid "github.com/arxanchain/sdk-go-common/structs/did"
)

// SignatureParam is used to pass signature params to SDK
// If you had trusted key pair, you should set 'SecurityCode',
// otherwish, you should set 'PrivateKey'.
type SignatureParam struct {
	Creator      commdid.Identifier `json:"creator"`       // signature creator
	Created      string             `json:"created"`       // signature created timestamp
	Nonce        string             `json:"nonce"`         // signature random string
	PrivateKey   string             `json:"private_key"`   // user ed25519 private key using for signing, base64 encoded
	SecurityCode string             `json:"security_code"` // user security code
}

// SignatureBody is used to pass the signature value to SDK
type SignatureBody struct {
	Creator        commdid.Identifier `json:"creator"`         // signature creator
	Created        string             `json:"created"`         // signature created timestamp
	Nonce          string             `json:"nonce"`           // signature random string
	SignatureValue string             `json:"signature_value"` // ed25519 signature value of the request payload, base64 encoded
}

// SignatureHeader ...
type SignatureHeader struct {
	SignType string             `json:"sign_type,omitempty"`
	Creator  commdid.Identifier `json:"creator,omitempty"`
	Created  int64              `json:"created,omitempty"`
	Nonce    []byte             `json:"nonce,omitempty"`
}

// Signature include signature header and body
type Signature struct {
	Header *SignatureHeader `json:"header,omitempty"`
	Sign   []byte           `json:"sign,omitempty"`
}

// SignedData is used to represent the general triplet required to verify a signature
// This is intended to be generic across crypto schemes, while most crypto schemes will
// include the signing identity and a nonce within the Data, this is left to the crypto
// implementation
type SignedData struct {
	Data   []byte           `json:"data,omitempty"`
	Header *SignatureHeader `json:"header,omitempty"`
	Sign   []byte           `json:"sign,omitempty"`
}

// Verify signature
func (sd *SignedData) Verify(ipk IPublicKey) error {
	message, err := json.Marshal(sd.Header)
	if nil != err {
		return err
	}
	message = append(message, sd.Data...)

	return ipk.Verify(message, sd.Sign)
}

// Sign message
func (sd *SignedData) DoSign(ipk IPrivateKey) (*Signature, error) {
	message, err := json.Marshal(sd.Header)
	if nil != err {
		return nil, err
	}
	message = append(message, sd.Data...)

	sd.Sign, err = ipk.Sign(message)
	if err != nil {
		return nil, err
	}
	return sd.GetSignature(), nil
}

// GetSignature generate signature struct from SignedData
func (sd *SignedData) GetSignature() *Signature {
	return &Signature{
		Header: sd.Header,
		Sign:   sd.Sign,
	}
}

// ISignable types are those which can map their contents to a set of SignedData
type ISignable interface {
	// 返回所有带签名数据的签名列表
	AsSignedData() ([]*SignedData, error)
	// 返回不带签名数据的签名
	NewSignedData() (*SignedData, error)
	// 附加签名
	AttachSignature(*Signature) error
}

// ISignable types are those which can map their contents to a set of SignedData
// ISigner ...
type ISigner interface {
	// NewSignatureHeader creates a SignatureHeader with the correct signing identity and a valid nonce
	NewSignatureHeader() (*SignatureHeader, error)
	// Sign a payload with creating a new signature header created by NewSignatureHeader
	Sign(signable ISignable) error
}

type ISignClient interface {
	// Sign a payload with creating a new signature header created by NewSignatureHeader
	DoSign(ipk IPrivateKey) (*Signature, error)

	// Verify signature
	Verify(ipk IPublicKey) error
	// GetSignature generate signature struct from SignedData
	GetSignature() *Signature
}
