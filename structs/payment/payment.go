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

package payment

import "net/http"

// IPaymentClient ...
type IPaymentClient interface {
	// racharge
	Recharge(http.Header, *RechargeBody) (interface{}, error)
	// bindcard
	BindCard(http.Header, *BindCardBody) (interface{}, error)
	// send bindcard sms
	SendBindSMS(http.Header, *SendSMSBody) (interface{}, error)
}

// RechargeBody ...
type RechargeBody struct {
	OrderID     string `json:"order_id"`
	DID         string `json:"did"`
	ChannelType string `json:"channel_type"`
	TxnAmt      int64  `json:"txn_amt"`
	AccNo       string `json:"acc_no"`
	UserMac     string `json:"user_mac"`
}

// SendSMSBody ...
type SendSMSBody struct {
	OrderID     string `json:"order_id"`
	ChannelType string `json:"channel_type"`
	Phone       string `json:"phone"`
	TxnAmt      int64  `json:"txn_amt"`
	AccNo       string `json:"acc_no"`
	UserMac     string `json:"user_mac"`
}

// BindCardBody ...
type BindCardBody struct {
	OrderID     string `json:"order_id,omitempty"`
	ChannelType string `json:"channel_type,omitempty"`
	TxnTime     string `json:"txn_time,omitempty"`
	AccNo       string `json:"acc_no,omitempty"`
	PhoneNo     string `json:"phone_no,omitempty"`
	CustomerNm  string `json:"customer_nm,omitempty"`
	CertifTp    string `json:"certif_tp,omitempty"`
	CertifID    string `json:"certif_id,omitempty"`
	SMSCode     string `json:"sms_code,omitempty"`
}
