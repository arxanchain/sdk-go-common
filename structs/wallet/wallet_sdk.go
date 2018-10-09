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
	"net/http"

	"github.com/arxanchain/sdk-go-common/protos/wallet"
	commdid "github.com/arxanchain/sdk-go-common/structs/did"
	"github.com/arxanchain/sdk-go-common/structs/pki"
)

// IWalletClient defines the behaviors implemented by wallet sdk
type IWalletClient interface {
	// Register is used to register user wallet.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will return the key pair.
	// If you want to trust the key pair, it will return the security code.
	//
	Register(http.Header, *RegisterWalletBody) (*WalletResponse, error)

	// RegisterSubWallet is used to register user subwallet.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will return the key pair.
	// If you want to trust the key pair, it will return the security code.
	//
	RegisterSubWallet(http.Header, *RegisterSubWalletBody) (*WalletResponse, error)

	// GetWalletBalance is used to get wallet balances.
	//
	GetWalletBalance(http.Header, commdid.Identifier) (*WalletBalance, error)

	// GetWalletInfo is used to get wallet base information.
	//
	GetWalletInfo(http.Header, commdid.Identifier) (*WalletInfo, error)

	// CreatePOE is used to create POE digital asset.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	CreatePOE(http.Header, *POEBody, *pki.SignatureParam) (*WalletResponse, error)

	// UpdatePOE is used to update POE digital asset.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	UpdatePOE(http.Header, *POEBody, *pki.SignatureParam) (*WalletResponse, error)

	// QueryPOE is used to query POE digital asset.
	//
	QueryPOE(http.Header, commdid.Identifier) (*POEPayload, error)

	// UploadPOEFile is used to upload file for specified POE digital asset
	//
	// poeID parameter is the POE digital asset ID pre-created using CreatePOE API.
	//
	// poeFile parameter is the path to file to be uploaded.
	//
	UploadPOEFile(header http.Header, poeID string, poeFile string, readOnly bool) (*UploadResponse, error)

	// IssueCToken is used to issue colored token.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	IssueCToken(http.Header, *IssueBody, *pki.SignatureParam) (*WalletResponse, error)

	// IssueAsset is used to issue digital asset.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	IssueAsset(http.Header, *IssueAssetBody, *pki.SignatureParam) (*WalletResponse, error)

	// TransferCToken is used to transfer colored tokens from one user to another.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	TransferCToken(http.Header, *TransferCTokenBody, *pki.SignatureParam) (*WalletResponse, error)

	// TransferAsset is used to transfer assets from one user to another.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	TransferAsset(http.Header, *TransferAssetBody, *pki.SignatureParam) (*WalletResponse, error)

	// SendIssueCTokenProposal is used to send issue ctoken proposal to get wallet.Tx to be signed.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	SendIssueCTokenProposal(http.Header, *IssueBody) (*IssueCTokenPrepareResponse, error)

	// SendIssueAssetProposal is used to send issue asset proposal to get wallet.Tx to be signed.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	SendIssueAssetProposal(http.Header, *IssueAssetBody) ([]*wallet.TX, error)

	// SendTransferCTokenProposal is used to send transfer colored tokens proposal to get wallet.Tx to be signed.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	SendTransferCTokenProposal(http.Header, *TransferCTokenBody) ([]*wallet.TX, error)

	// SendTransferAssetProposal is used to send transfer asset proposal to get wallet.Tx to be signed.
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The default key pair trust mode does not trust, it will required key pair.
	// If you had trust the key pair, it will required security code.
	//
	SendTransferAssetProposal(http.Header, *TransferAssetBody) ([]*wallet.TX, error)

	// ProcessTx is used to transfer formally with signature TX .
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	ProcessTx(http.Header, []*wallet.TX) (*WalletResponse, error)

	// QueryTransactionLogs is used to query transaction logs.
	//
	// txType:
	// in: query income type transaction
	// out: query spending type transaction
	// num, page: count, page of logs that to be returned
	//
	QueryTransactionLogs(http.Header, commdid.Identifier, string, int32, int32) ([]*wallet.UTXO, error)

	// QueryTransactionUTXO is used to query transaction utxo logs.
	// num, page: count, page of logs that to be returned
	//
	QueryTransactionUTXO(http.Header, commdid.Identifier, int32, int32) ([]*wallet.UTXO, error)

	// QueryTransactionUTXO is used to query transaction stxo logs.
	// num, page: count, page of logs that to be returned
	//
	QueryTransactionSTXO(http.Header, commdid.Identifier, int32, int32) ([]*wallet.UTXO, error)

	// IndexSet is used to create indexs for object-id
	//
	// The default invoking mode is asynchronous, it will return
	// without waiting for blockchain transaction confirmation.
	//
	// If you want to switch to synchronous invoking mode, set
	// 'BC-Invoke-Mode' header to 'sync' value. In synchronous mode,
	// it will not return until the blockchain transaction is confirmed.
	//
	// The response result is the blockchain transaction id list
	//
	IndexSet(http.Header, *IndexSetPayload) ([]string, error)

	// IndexGet is used to query object-id via indexs
	//
	// The response result is the query hit object-id list
	//
	IndexGet(http.Header, *IndexGetPayload) ([]string, error)
}
