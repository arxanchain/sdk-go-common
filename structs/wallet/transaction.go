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
	"github.com/arxanchain/sdk-go-common/protos/wallet"
)

////////////////////////////////////////////////////////////////////////////////
// Transaction Structs

type AXTUnit int64

const (
	ATOM     AXTUnit = 1
	MicroAXT         = 1000 * ATOM
	AXT              = 1000 * MicroAXT
)

// Colored Token Amount Structure
type TokenAmount struct {
	TokenId string `json:"token_id,omitempty"`
	Amount  int64  `json:"amount,omitempty"`
}

// Transaction Fee Structure
type Fee struct {
	Amount AXTUnit `json:"amount,omitempty"`
}

// Issue Asset Request Structure
type IssueAssetBody struct {
	Issuer  string `json:"issuer,omitempty"`
	Owner   string `json:"owner,omitempty"`
	AssetId string `json:"asset_id,omitempty"`
	Fee     *Fee   `json:"fee,omitempty"`
}

// Issue Colored Token Request Structure
type IssueBody struct {
	Issuer  string `json:"issuer,omitempty"`
	Owner   string `json:"owner,omitempty"`
	AssetId string `json:"asset_id,omitempty"`
	TokenId string `json:"token_id,omitempty"`
	Amount  int64  `json:"amount,omitempty"`
	Fee     *Fee   `json:"fee,omitempty"`
}

// Transfer Colored Token Request Structure
type TransferCTokenBody struct {
	From    string         `json:"from,omitempty"`
	To      string         `json:"to,omitempty"`
	AssetId string         `json:"asset_id,omitempty"`
	Tokens  []*TokenAmount `json:"tokens,omitempty"`
	Fee     *Fee           `json:"fee,omitempty"`
}

// Transfer to process Tx Request Structure
type ProcessTxBody struct {
	Txs []*wallet.TX `json:"txs,omitempty"`
}

// Transfer Asset Request Structure
type TransferAssetBody struct {
	From   string   `json:"from,omitempty"`
	To     string   `json:"to,omitempty"`
	Assets []string `json:"assets,omitempty"`
	Fee    *Fee     `json:"fee,omitempty"`
}

type IssueCTokenPrepareResponse struct {
	TokenId string       `json:"token_id,omitempty"`
	Txs     []*wallet.TX `json:"txs,omitempty"`
}
