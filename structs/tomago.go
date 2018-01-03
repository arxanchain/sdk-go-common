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

package structs

import "net/http"

//Tomago API Response Struct Define
type TomagoResponse struct {
	Code           int      `json:"code"`
	Message        string   `json:"message"`
	Id             string   `json:"id"`
	CoinId         string   `json:"coin_id"`
	TransactionIds []string `json:"transaction_ids"`
}

//TomagoClient Interface
type ITomagoClient interface {
	GetEntityClient() IEntityClient
	GetAssetClient() IAssetClient
	GetCCoinClient() ICCoinClient
}

//EntityClient Interface
type IEntityClient interface {
	CreateEntity(http.Header, *EntityBody) (*TomagoResponse, error)
	UpdateEntity(http.Header, string, *EntityBody) (*TomagoResponse, error)
	QueryEntity(http.Header, string) (*EntityPayload, error)
}

//AssetClient Interface
type IAssetClient interface {
	CreateAsset(http.Header, *AssetBody) (*TomagoResponse, error)
	UpdateAsset(http.Header, string, *AssetBody) (*TomagoResponse, error)
	QueryAsset(http.Header, string) (*AssetPayload, error)
	TransferAsset(http.Header, *TransferAssetBody) (*TomagoResponse, error)
}

//CCoinClient Interface
type ICCoinClient interface {
	Issue(http.Header, *IssueBody) (*TomagoResponse, error)
	Transfer(http.Header, *TransferBody) (*TomagoResponse, error)
	Rollback(http.Header, *RollbackBody) (*TomagoResponse, error)
	Interest(http.Header, *InterestBody) (*TomagoResponse, error)
	Withdraw(http.Header, *WithdrawBody) (*TomagoResponse, error)
}
