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

package structs

// PayloadWithTags Defines Invoke and Query API Request Struct
type PayloadWithTags struct {
	Payload    *ChaincodeRequest `json:"payload"`
	QueryTags  []string          `json:"queryTags"`
	UpdateTags []string          `json:"updateTags"`
}

// ChaincodeRequest Struct Define
type ChaincodeRequest struct {
	Channel     string   `json:"channel"`
	ChaincodeID string   `json:"chaincode_id"`
	Args        []string `json:"args"`
}

// ChaincodeResponse Defines Invoke and Query API Response Struct
type ChaincodeResponse struct {
	Result  string `json:"result"`
	Code    int64  `json:"Code"`
	Message string `json:"Message"`
}

// TransactionResponse Defines QueryTxn API Response Struct
type TransactionResponse struct {
	Code          int64     `json:"Code"`
	Message       string    `json:"Message"`
	ChannelID     string    `json:"channel_id"`
	ChaincodeID   string    `json:"chaincode_id"`
	TransactionID string    `json:"transaction_id"`
	Timestamp     Timestamp `json:"timestamp"`
	CreatorID     []byte    `json:"creator_id"`
	PayloadSize   uint64    `json:"payload_size"`
	IsInvalID     bool      `json:"is_invalid"`
	Payload       string    `json:"payload"`
}

// Timestamp Struct Defines
type Timestamp struct {
	Seconds int64 `json:"seconds"`
	Nanos   int32 `json:"nanos"`
}
