/*
Copyright ArxanFintech Technology Ltd. 2017-2018 All Rights Reserved.

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

package wallet

import (
	pw "github.com/arxanchain/sdk-go-common/protos/wallet"
	commdid "github.com/arxanchain/sdk-go-common/structs/did"
)

////////////////////////////////////////////////////////////////////////////////
// POE Structs

const (
	// OffchainPOEID poe did when invoke upload file api formdata param
	OffchainPOEID = "poe_id"
	// OffchainPOEFile poe binary file when invoke upload file api formdata param
	OffchainPOEFile = "poe_file"
	// OffchainPOEFile poe file read/write attribute when upload file api formdata param
	OffchainReadOnly = "read_only"
	// DownloadUrl url link to download poe file related to POE instance
	DownloadUrl = "url"
)

// signature const for the formdata
const (
	// SignatureCreator issuer did
	SignatureCreator = "signature.creator"
	// SignatureCreated timestamp
	SignatureCreated = "signature.created"
	// SignatureNonce nonce
	SignatureNonce = "signature.nonce"
	// SignatureSignatureValue sign value
	SignatureSignatureValue = "signature.signatureValue"
)

// POEBody POE request body structure definition
type POEBody struct {
	Id         commdid.Identifier `json:"id,omitempty"`
	Name       string             `json:"name,omitempty"`
	ParentId   commdid.Identifier `json:"parent_id,omitempty"`
	Owner      commdid.Identifier `json:"owner,omitempty"`
	ExpireTime int64              `json:"expire_time,omitempty"`
	Hash       string             `json:"hash,omitempty"`
	Metadata   []byte             `json:"metadata,omitempty"`
	Indexs     *IndexTags         `json:"indexs,omitempty"`
}

// POEPayload POE query payload structure definition
type POEPayload struct {
	Id               commdid.Identifier `json:"id,omitempty"`
	Name             string             `json:"name,omitempty"`
	ParentId         commdid.Identifier `json:"parent_id,omitempty"`
	Owner            commdid.Identifier `json:"owner,omitempty"`
	ExpireTime       int64              `json:"expire_time,omitempty"`
	Hash             string             `json:"hash,omitempty"`
	Metadata         []byte             `json:"metadata,omitempty"`
	OffchainMetadata OffchainMetadata   `json:"offchain_metadata,omitempty"`
	Created          int64              `json:"created,omitempty"`
	Updated          int64              `json:"updated,omitempty"`
	Status           pw.Status          `json:"status,omitempty"`
}

// OffchainMetadata offchain storage metadata
type OffchainMetadata struct {
	Filename    string `json:"filename,omitempty"`
	Endpoint    string `json:"endpoint,omitempty"`
	StorageType string `json:"storageType,omitempty"`
	ContentHash string `json:"contentHash,omitempty"`
	Size        int    `json:"size,omitempty"`
	ReadOnly    bool   `json:"read_only,omitempty"`
}

// UploadResponse Upload POE file response payload
type UploadResponse struct {
	Id               commdid.Identifier `json:"id,omitempty"`
	OffchainMetadata *OffchainMetadata  `json:"offchain_metadata,omitempty"`
	TransactionIds   []string           `json:"transaction_ids,omitempty"`
}
