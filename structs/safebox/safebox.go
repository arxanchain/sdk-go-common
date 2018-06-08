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

// DeleteKeyPairRequestBody ...
type DeleteKeyPairRequestBody struct {
	UserDid string `json:"user_did"`
	Code    string `json:"code"`
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
