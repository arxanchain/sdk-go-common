/*
Copyright ArxanFintech Technology Ltd. 2018 All Rights Reserved.

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

package safebox

import (
	"net/http"

	commdid "github.com/arxanchain/sdk-go-common/structs/did"
)

// SaveKeyPairRequetBody define save KeyPair request body
type SaveKeyPairRequetBody struct {
	UserDid    string `json:"user_did"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

// SaveKeyPairReply define save KeyPair reply
type SaveKeyPairReply struct {
	Code string `json:"code"`
}

// PrivateKeyReply ...
type PrivateKeyReply struct {
	PrivateKey string `json:"private_key"`
}

// PublicKeyReply ...
type PublicKeyReply struct {
	PublicKey string `json:"public_key"`
}

// UpdateSecurityCodeRequestBody ...
type UpdateSecurityCodeRequestBody struct {
	UserDid      string `json:"user_did"`
	OriginalCode string `json:"original_code"`
	NewCode      string `json:"new_code"`
}

// CodeInfoReply ...
type CodeInfoReply struct {
	Code string `json:"code"`
}

// ISafeboxClient defines the behaviors implemented by safebox sdk
type ISafeboxClient interface {
	// TrusteeKeyPair is used to trutee keypair.
	//
	// API-Key must set to header.
	TrusteeKeyPair(http.Header, *SaveKeyPairRequetBody) (*SaveKeyPairReply, error)

	// QueryPrivateKey is used to query private key.
	//
	// API-Key must set to header.
	QueryPrivateKey(http.Header, *OperateKeyInfo) (*PrivateKeyReply, error)

	// QueryPublicKey is used to query public key.
	//
	// API-Key must set to header.
	QueryPublicKey(http.Header, *OperateKeyInfo) (*PublicKeyReply, error)

	// DeleteKeyPair is used to delete keypair.
	//
	// API-Key must set to header.
	DeleteKeyPair(http.Header, *OperateKeyInfo) error

	// UpdateAssistCode is used to update assist code.
	//
	// API-Key must set to header.
	UpdateAssistCode(http.Header, *UpdateSecurityCodeRequestBody) error

	// RecoverAssistCode is used to recover assist code when user has forgot.
	//
	// API-Key must set to header.
	RecoverAssistCode(http.Header, commdid.Identifier) (*CodeInfoReply, error)
}

// OperateKeyInfo ...
type OperateKeyInfo struct {
	UserDid string `json:"user_did"`
	Code    string `json:"code"`
}
