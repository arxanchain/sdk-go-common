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

import (
	"net/http"

	reststructs "github.com/arxanchain/sdk-go-common/rest/structs"
)

type IFredClient interface {
	GetUserClient() IUserClient
	GetEdkeyClient() IEdkeyClient
	GetCertsClient() ICertsClient
}

type IUserClient interface {
	// register new user
	RegisterUser(*RegisterRequest, http.Header) (*ResponseStruct, error)
	// revoke user
	Revoke(*RevokeRequest, http.Header) error
	// login
	Login(*LoginRequest) (*GetTokenResponse, error)
	// get access info
	GetAccessInfo(string, http.Header) (interface{}, error)
	// update user password
	UpdateUserPassword(*UpdatePasswordRequest, http.Header) error
	// revoke token
	RevokeToken(http.Header) error
	// validate token
	ValidateToken(http.Header) (*TokenInfo, error)
	// get version
	Version() (*reststructs.VersionResponse, error)
	// update channel id
	UpdateChannelID(*UpdateChannelIDRequest, http.Header) error
	// get accesskey
	Accesskey(*AccesskeyRequest, http.Header) (*AccesskeyResponse, error)
	// validate acl
	ACLVerification(http.Header) (*ACLResponse, error)
}

type IEdkeyClient interface {
	// generate key
	Generate(string, http.Header) ([]byte, []byte, error)
	// download key
	Download(string, string, string, http.Header) error
}

type ICertsClient interface {
	// create cert
	CreateCerts(*CertCreateReqBody, http.Header) (*CertCreateRespBody, error)
	// disable cert
	DisableCerts(*CertCreateReqBody, http.Header) (*CertCreateRespBody, error)
	// recover cert
	RecoverCerts(*CertCreateReqBody, http.Header) (*CertCreateRespBody, error)
	// delete cert
	DeleteCerts(*CertCreateReqBody, http.Header) (*CertCreateRespBody, error)
	// get cert status
	StatusCerts(string, http.Header) (*GetCertStatusResp, error)
	// retrieve cert
	RetrieveCerts(string, string, http.Header) ([]byte, error)
}

// IOauth2Client is oauth2 function
type IOauth2Client interface {
	// validate oauth2 token
	ValidateOauth2Token(*Oauth2AccessToken) (bool, error)
}

type RegisterRequest struct {
	Credential RegisterBody `json:"credential"`
}

type RegisterBody struct {
	Role        string `json:"role"`
	Description string `json:"description"`
	ChannelId   string `json:"channel_id"`
	Value       User   `json:"value"`
}

type User struct {
	Identifier string `json:"identifier"`
	Access     string `json:"access"` // username
	Secret     string `json:"secret"` // password
}

type UserInfo struct {
	Id         string `json:"id"`
	Access     string `json:"access"`
	Status     string `json:"status"`
	Roles      string `json:"roles"`
	Issued_at  int64  `json:"issued_at"`
	Channel_id string `json:"channel_id"`
}

type LoginAccessSecret struct {
	Access string `json:"access"`
	Secret string `json:"secret"`
}

type LoginAccess struct {
	Value LoginAccessSecret
}

type UpdatePasswordRequest struct {
	Credential *UpdatePasswordBody `json:"credential"`
}

type UpdatePasswordBody struct {
	OriginalSecret string `json:"origin_secret"`
	NewSecret      string `json:"new_secret"`
	Access         string `json:"access"`
}

// LoginRequest is token request
type LoginRequest struct {
	Credential LoginAccess `json:"credential"`
}

type RevokeRequest struct {
	Credential *User `json:"credential"`
}

type ChannelIDAccess struct {
	ChannelID string `json:"channel_id"`
	Access    string `json:"access"`
}

type UpdateChannelIDRequest struct {
	Credential *ChannelIDAccess `json:"credential"`
}

type CertCreateReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Access      string `json:"access"`
}

type CertCreateReqBody struct {
	Certificate CertCreateReq `json:"certificate"`
}

type Cert struct {
	Key    string `json:"key"`
	Cert   string `json:"cert"`
	CACert string `json:"ca_cert"`
}

// cert resp info
type CertCreateResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Access      string `json:"access"`
	IssuedAt    int64  `json:"issued_at"`
	Hash        string `json:"hash"`
	Status      string `json:"status"`
	Value       *Cert  `json:"value"`
}

// CertCreateRespBody is cert resp info
type CertCreateRespBody struct {
	Certificate *CertCreateResp `json:"certificate"`
}

// GetCertStutasResp is cert resp status
type GetCertStatusResp struct {
	Certificate GetCertStatusRespInner `json:"certificate"`
}

// GetCertStutasRespInner is cert resp issued_at
type GetCertStatusRespInner struct {
	Isued_at int64 `json:"issued_at"`
}

type Oauth2AccessToken struct {
	AccessToken string `json:"access_token"`
}

// 获取token成功时返回的token详情
type TokenDetail struct {
	IssuedAt     int64  `json:"issue_at"`
	ExpiresAt    int64  `json:"expires_at"`
	Value        string `json:"value"`
	CredentialId string `json:"credential_id"`
	Roles        string `json:"roles"`
	ChannelId    string `json:"channel_id"`
}

// 获取token成功时返回的response
type GetTokenResponse struct {
	Token TokenDetail `json:"token"`
}

// 通过token能获取的用户信息
type TokenInfo struct {
	CredentialId string `json:"credential_id"`
	Access       string `json:"access"`
	Roles        string `json:"roles"`
}

type ResponseStruct struct {
	Credentials CredentialsStruct `json:"credentials"`
}
type CredentialsStruct struct {
	Id          string      `json:"id"`
	Role        string      `json:"role"`
	Description string      `json:"description"`
	Channel_id  string      `json:"channel_id"`
	Issued_at   int64       `json:"issued_at"`
	Value       ValueStruct `json:"value"`
}

type ValueStruct struct {
	Access string `json:"access"`
	//Secret string `json:"secret"`
}

type AccesskeyRequest struct {
	Accesskey AccesskeyRequestInner `json:"accesskey"`
}

type AccesskeyRequestInner struct {
	Description string `json:"description"`
}

type AccesskeyResponse struct {
	Accesskey AccesskeyResponseInner `json:"accesskey"`
}

type AccesskeyResponseInner struct {
	Id          string         `json:"id"`
	Description string         `json:"description"`
	Issued_at   int64          `json:"issued_at"`
	Value       AccessKeyValue `json:"value"`
}

type AccessKeyValue struct {
	Access    string `json:"access"`
	Requester string `json:"requester"`
}

type ACLResponse struct {
	Access    string `json:"access"`
	Roles     string `json:"roles"`
	ChannelId string `json:"channel_id"`
}
