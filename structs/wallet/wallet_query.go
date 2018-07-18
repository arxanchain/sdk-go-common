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

package wallet

//DApp issued/remain amount
type DAppIssuedTokens map[string]*DAppIssuedToken // key is token_id

//DAppTotal define DApp Total struct
type DAppTotal struct {
	TotalAssets         int64 `json:"total_assets"`           //企业用户拥有数字资产的数量
	TotalTokens         int64 `json:"total_tokens"`           //企业用户拥有数字凭证的数量
	TotalTxs            int64 `json:"total_txs"`              //企业用户交易的总量
	TotalIncomeTxs      int64 `json:"total_income_txs"`       //企业用户输入的交易总量
	TotalSpendingTxs    int64 `json:"total_spending_txs"`     //企业用户输出的交易总量
	TotalAxtConsumedTxs int64 `json:"total_axt_consumed_txs"` //企业用户消耗AXT的交易总量
	TotalIncomeAssets   int64 `json:"total_income_assets"`    //企业用户输入的数字资产总量
	TotalSpendingAssets int64 `json:"total_spending_assets"`  //企业用户输出的数字资产总量
	TotalIncomeTokens   int64 `json:"total_income_tokens"`    //企业用户输入的数字凭证总量
	TotalSpendingTokens int64 `json:"total_spending_tokens"`  //企业用户输出的数字凭证总量
}

//DAppUsersTotal define DApp Uesrs Total struct
type DAppUsersTotal struct {
	TotalUsers int64 `json:"total_users"` //应用用户的数量
}

//DAppAxtTotal define DApp Axt Total struct
type DAppAxtTotal struct {
	AssignAmount int64 `json:"assign_amount"` // AXT分配总量
	RemainAmount int64 `json:"remain_amount"` //AXT剩余总量
}

type DAppUsersLists []*DAppUsersList

//DAppUsersList define DApp Users List struct
type DAppUsersList struct {
	Id    int64  `json:"id"`    //用户Id
	Name  string `json:"name"`  //用户名称
	Email string `json:"email"` //用户邮箱
	Phone string `json:"phone"` //用户电话
}

type DAppAxtConsumes []*DAppAxtConsume

//DAppAxtConsume define DApp Axt Consume struct
type DAppAxtConsume struct {
	Datetime      string `json:"datetime"`       //日期
	ConsumeAmount int64  `json:"consume_amount"` //消耗AXT的数量
}

type TopAssetUsers []*TopAssetUser

//TopAssetUser define Top Asset User struct
type TopAssetUser struct {
	UserId    string `json:"user_id"`    //用户DID
	AssetsNum int64  `json:"assets_num"` //用户拥有的数字资产数量
}

type TopTokenUsers []*TopTokenUser

//TopTokenUser define Top Token User struct
type TopTokenUser struct {
	UserId    string `json:"user_id"`    //用户DID
	TokensNum int64  `json:"tokens_num"` //用户拥有的数字凭证数量
}

type HotAssets []*HotAsset

//HotAsset define Hot Asset stuct
type HotAsset struct {
	AssetId string `json:"asset_id"` //数字资产ID
	TxsNum  int64  `json:"txs_num"`  //该数字资产总的交易数量
}

type HotTokens []*HotToken

//HotToken define Hot Token struct
type HotToken struct {
	TokenId string `json:"token_id"` //数字凭证ID
	TxsNum  int64  `json:"txs_num"`  //该数字凭证总的交易数量
}

//Total define Total struct
type Total struct {
	TotalDApps  int64 `json:"total_dapps"`  //平台上所有DAPP的数量
	TotalUsers  int64 `json:"total_users"`  //平台上所有应用用户的数量
	TotalAssets int64 `json:"total_assets"` //平台上所有数字资产的数量
	TotalTokens int64 `json:"total_tokens"` //平台上所有数字凭证的数量
}

//DAppNum define DApp's Number struct
type DAppNum struct {
	Num int64 `json:"num"` //DAPP的数量
}

//DAppIssuedToken define DApp Issued Token struct
type DAppIssuedToken struct {
	IssuedAmount int64 `json:"issued_amount"` //发行的总量
	RemainAmount int64 `json:"remain_amount"` //剩余的总量
}

type DAppLists []*DAppList

//DAppList define DApp's List struct
type DAppList struct {
	Name        string `json:"name"`        //DAPP的名称
	Description string `json:"description"` //DAPP的描述
	Logo        string `json:"logo"`        //DAPP的Logo路径
}

type Growths []*Growth

//Growth define  Growth struct
type Growth struct {
	Datetime     string `json:"datetime"`      //日期
	GrowthAmount int64  `json:"growth_amount"` //增长量
}
