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

// 资产状态定义
const (
	ASSET_STATUS_INIT    int = iota //未发行
	ASSET_STATUS_ISSUED             //已发行
	ASSET_STATUS_EXPIRED            //已过期
)

// 数字资产结构定义
type DigitalAsset struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// 数字资产的请求Body结构定义
type AssetBody struct {
	Id           string      `json:"id"`
	EnrollmentId string      `json:"enrollment_id"`
	CallbackUrl  string      `json:"callback_url"`
	Name         string      `json:"name"`
	Hash         string      `json:"hash"`
	ParentId     string      `json:"parent_id"`
	Owner        string      `json:"owner"`
	ExpireTime   int64       `json:"expire_time"`
	Metadata     interface{} `json:"metadata"`
}

//数字资产的Payload定义
type AssetPayload struct {
	Id         string                  `json:"id"`          //数字资产ID
	Name       string                  `json:"name"`        //数字资产的名称
	Hash       string                  `json:"hash"`        //数字资产Hash
	ParentId   string                  `json:"parent_id"`   //父资产ID
	Owner      string                  `json:"owner"`       //数字资产所属人ID
	ExpireTime int64                   `json:"expire_time"` //数字资产过期时间
	Metadata   interface{}             `json:"metadata"`    //数字资产元数据
	CreateTime int64                   `json:"create_time"` //创建时间
	UpdateTime int64                   `json:"update_time"` //更新时间
	IssueCoins map[string]*ColoredCoin `json:"issue_coins"` //发行的染色币
	Status     int                     `json:"status"`      //数字资产状态
}

// 转移资产请求Body结构定义
type TransferAssetBody struct {
	EnrollmentId string   `json:"enrollment_id"`
	CallbackUrl  string   `json:"callback_url"`
	From         string   `json:"from"`
	To           string   `json:"to"`
	Assets       []string `json:"assets"`
	Fees         *Fees    `json:"fees"`
}

// 转移资产的Payload结构定义
type TransferAssetPayload struct {
	From       string   `json:"from"`        //交易发起人ID
	To         string   `json:"to"`          //交易接收人ID
	Assets     []string `json:"assets"`      //交易的资产
	Fees       *Fees    `json:"fees"`        //交易的手续费
	UpdateTime int64    `json:"update_time"` //当前时间
}
