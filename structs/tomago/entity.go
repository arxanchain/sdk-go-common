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

package tomago

//钱包的定义
type Wallet struct {
	ColoredCoins  map[string]*ColoredCoin  `json:"colored_coins"`  //钱包中所有Token
	DigitalAssets map[string]*DigitalAsset `json:"digital_assets"` //钱包中的数字资产
}

// 组织实体的请求Body结构定义
type EntityBody struct {
	Id           string      `json:"id"`            //组织实体ID
	EnrollmentId string      `json:"enrollment_id"` //Enrollment id
	CallbackUrl  string      `json:"callback_url"`  //用于接收异步通知的回调URL
	Metadata     interface{} `json:"metadata"`      //元数据
}

//组织实体的Payload定义
type EntityPayload struct {
	Id         string      `json:"id"`          //组织实体的ID
	Wallet     *Wallet     `json:"wallet"`      //钱包
	CreateTime int64       `json:"create_time"` //创建时间
	UpdateTime int64       `json:"update_time"` //更新时间
	Metadata   interface{} `json:"metadata"`    //元数据
}
