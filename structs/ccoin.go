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

// 染色币状态定义
const (
	CCOIN_STATUS_ISSUE   int = iota //发行
	CCOIN_STATUS_INUSE              //流通
	CCOIN_STATUS_REVOKE             //回收
	CCOIN_STATUS_EXPIRED            //过期
)

// 染色币颜色结构定义
type CoinColor struct {
	Id         string `json:"id"`          //染色币的ID
	Ancestor   string `json:"ancestor"`    //发行染色币的数字资产的ID
	Issuer     string `json:"issuer"`      //发行人
	IssueTime  int64  `json:"issue_time"`  //发行时间
	ExpireTime int64  `json:"expire_time"` //过期时间
}

// 染色币结构定义
type ColoredCoin struct {
	Amount    int64      `json:"amount"`     //染色币的数量
	Status    int        `json:"status"`     //染色币的状态
	CoinColor *CoinColor `json:"coin_color"` //染色币的颜色属相
}

// 染色币币种和数量
type CoinAmount struct {
	CoinId string `json:"coin_id"` //染色币ID
	Amount int64  `json:"amount"`  //染色币数量
}

//审计相关的账户差额定义
type AuditAmount struct {
	Amount int64 `json:"amount"` //审计出来的染色币差额
}

// 手续费结构定义
type Fees struct {
	Accounts []string      `json:"accounts"` //收取手续费的账户列表
	Coins    []*CoinAmount `json:"coins"`    //收取手续费的币种和数量
}

// 发行染色币的请求Body结构定义
type IssueBody struct {
	EnrollmentId string `json:"enrollment_id"`
	CallbackUrl  string `json:"callback_url"`
	Issuer       string `json:"issuer"`
	Owner        string `json:"owner"`
	AssetId      string `json:"asset_id"`
	Amount       int64  `json:"amount"`
	Fees         *Fees  `json:"fees"`
}

//发行的Payload结构定义
type IssuePayload struct {
	Issuer     string                  `json:"issuer"`      //发行人ID
	Owner      string                  `json:"owner"`       //发行的数字资产的所属人ID
	AssetId    string                  `json:"asset_id"`    //发行的数字资产ID
	Coins      map[string]*ColoredCoin `json:"coins"`       //发行的染色币和数量
	Fees       *Fees                   `json:"fees"`        //发行的手续费
	UpdateTime int64                   `json:"update_time"` //更新时间
}

// 转账交易的请求Body结构定义
type TransferBody struct {
	EnrollmentId string        `json:"enrollment_id"`
	CallbackUrl  string        `json:"callback_url"`
	From         string        `json:"from"`
	To           string        `json:"to"`
	AssetId      string        `json:"asset_id"`
	Coins        []*CoinAmount `json:"coins"`
	Fees         *Fees         `json:"fees"`
}

//转账的Payload结构定义
type TransferPayload struct {
	From       string        `json:"from"`        //转账交易的发起人ID
	To         string        `json:"to"`          //转账交易接收人ID
	AssetId    string        `json:"asset_id"`    //交易的数字资产的ID
	Coins      []*CoinAmount `json:"coins"`       //转账交易的染色币及数量
	Fees       *Fees         `json:"fees"`        //转账交易的手续费
	UpdateTime int64         `json:"update_time"` //更新时间
}

// 提现的请求Body结构定义
type WithdrawBody struct {
	EnrollmentId string        `json:"enrollment_id"`
	CallbackUrl  string        `json:"callback_url"`
	Funder       string        `json:"funder"`
	FunderCoins  []*CoinAmount `json:"funder_coins"`
	Holder       string        `json:"holder"`
	HolderCoins  []*CoinAmount `json:"holder_coins"`
	Fees         *Fees         `json:"fees"`
	IsDue        bool          `json:"is_due"`
}

//提现的Payload结构定义
type WithdrawPayload struct {
	Funder      string        `json:"funder"`       //提现的资金方ID
	FunderCoins []*CoinAmount `json:"funder_coins"` //用于兑现币种和数量
	Holder      string        `json:"holder"`       //染色币的持有人ID
	HolderCoins []*CoinAmount `json:"holder_coins"` //要兑现的染色币
	Fees        *Fees         `json:"fees"`         //提现的手续费
	IsDue       bool          `json:"is_due"`       //是否已到期
	UpdateTime  int64         `json:"update_time"`  //更新时间
}

// 平账的请求Body结构定义
type ReverseBody struct {
	EnrollmentId string                             `json:"enrollment_id"`
	CallbackUrl  string                             `json:"callback_url"`
	Audits       map[string]map[string]*AuditAmount `json:"audits"` //用于平账的数据信息
}

//平账的Payload结构定义
type ReversePayload struct {
	EntityId   string                  `json:"entity_id"`   //平账的账户ID
	Audits     map[string]*AuditAmount `json:"audits"`      //要平账的染色币信息
	UpdateTime int64                   `json:"update_time"` //更新时间
}

// 计息的请求Body结构定义
type InterestBody struct {
	EnrollmentId string                   `json:"enrollment_id"`
	CallbackUrl  string                   `json:"callback_url"`
	Interests    map[string][]*CoinAmount `json:"interests"`
}

// 交易回滚的请求Body结构定义
type RollbackBody struct {
	EnrollmentId   string   `json:"enrollment_id"`
	CallbackUrl    string   `json:"callback_url"`
	TransactionIds []string `json:"transaction_ids"`
}
