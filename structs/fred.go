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

const (
	FredAPIParam_APIKey      = "api_key"
	FredAPIParam_Type        = "type"
	FredAPIParam_Identifier  = "identifier"
	FredAPIParam_ACLGroupID  = "id"
	FredUploadFileFormHeader = "file"
	FredAPIParam_Page        = "page"
	FredAPIParam_Num         = "num"
)

// 企业认证信息
const (
	FredCompanyNameFormHeader         = "company_name"          // 企业名称
	FredCompanyAddrFormHeader         = "company_address"       // 企业地址
	FredTaxCodeFormHeader             = "tax_code"              // 企业税号
	FredJuridicalPersonNameFormHeader = "juridical_person_name" // 法人姓名
	FredJuridicalPersonTelFormHeader  = "juridical_person_tel"  // 法人电话
	FredJuridicalPersonIDFormHeader   = "juridical_person_id"   // 法人身份证号
	FredLicenceFormHeader             = "licence"               // 营业执照文件
	FredIdentityTopFormHeader         = "identity_top"          // 法人身份证正面
	FredIdentityBottomFormHeader      = "identity_bottom"       // 法人身份证背面
	FredLicenceNumberHeader           = "licence_number"        // 营业执照编号
)

// 实名认证状态
const (
	FredVerificationStatusNotUploaded  = 1 // 身份文件未上传
	FredVerificationStatusUnverified   = 2 // 待审核
	FredVerificationStatusVerifyFailed = 3 // 审核失败
	FredVerificationStatusVerified     = 4 // 认证成功
)

const (
	FredUserType_Service    = 1 // 系统服务用户, 内部使用 不允许外部创建
	FredUserType_Enterprise = 2 // 企业用户
	FredUserType_AppChain   = 3 // ChainAPP用户, 用于链群网关使用
	FredUserType_Normal     = 4 // 普通用户，企业DAPP创建的普通用户
	FredUserType_DApp       = 5 // DAPP用户,用于DAPP应用

)

// 权限级别
const (
	FredGroupType_Service     = 1 // 系统服务权限, 可以访问所有API
	FredGroupType_ServicePart = 2 // 系统服务权限, 可以访问部分API
	FredGroupType_Normal      = 3 // 非实名认证用户
	FredGroupType_Superadmin  = 4 // 超级管理员用户
	FredGroupType_Civil       = 5 // 实名认证用户
)

type IFredClient interface {
	GetUserClient() IUserClient
	GetEdkeyClient() IEdkeyClient
	GetCertsClient() ICertsClient
	GetACLClient() IACLClient
	GetSubwalletClient() ISubwalletClient
}

type IUserClient interface {
	// register new user
	RegisterUser(*RegisterRequest, http.Header) (*ResponseStruct, error)
	// revoke user
	Revoke(*RevokeRequest, http.Header) error
	// login
	Login(*LoginRequest) (*GetTokenResponse, error)
	// get user info with api key
	GetUserInfoWithAPIKey(string, http.Header) (*UserInfo, error)
	// get user info with DID
	GetUserInfoWithDID(string, http.Header) (*UserInfo, error)
	// reset user password
	ResetUserPassword(*UpdatePasswordRequest, http.Header) error
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
	// update identity status
	UpdateIdentityStatus(UserInfo, http.Header) error
	// query identity status
	QueryIdentityStatus(string, http.Header) (*UserInfo, error)
	// upload identity file
	UploadIdentity(string, string, http.Header) error
	// query usertype users count. userType 1: Enterprise 2: Persional 0: Enterprise+Persional
	QueryUsersCount(int, http.Header) (*UsersNumResponse, error)
	// query dapp users count creted by did
	QueryDappUsersCount(string, http.Header) (*DappUsersNumResponse, error)
	// query userType users list
	QueryUsersList(userType, page, num int, head http.Header) ([]DappInfoResponse, error)
	// query dapp users list created by did
	QueryDappUsersList(did string, page, num int, head http.Header) ([]DappInfoResponse, error)
	// query the last num days or months(depend on querytype) users growth
	QueryUsersGrowth(queryType, num int, head http.Header) ([]UsersGrowthResponse, error)
	// update user ACL by feature list
}

// IACLClient ...
type IACLClient interface {
	// create acl group
	CreateACLGroup(*ACLGroup, http.Header) (*ACLGroup, error)
	// update acl group
	UpdateACLGroup(*ACLGroup, http.Header) (*ACLGroup, error)
	// delete acl group
	DeleteACLGroup(*ACLGroups, http.Header) (*ACLGroups, error)
	// get acl groups
	GetACLGroups(uint, uint, http.Header) (*ACLGroups, error)
	// get acl group
	GetACLGroup(uint, http.Header) (*ACLGroupResource, error)
	// add acl resource
	AddACLResource(*ACLResourceRequest, http.Header) (*ACLGroup, error)
	// remove acl resource
	RemoveACLResource(*ACLResourceRequest, http.Header) (*ACLGroup, error)
	// get user acl group
	GetUserACLGroup(uint, http.Header) (*ACLGroup, error)
	// update user acl group
	UpdateUserACLGroup(*UpdateUserGroupRequest, http.Header) (*UpdateUserGroupRequest, error)
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

// ISubwalletClient is DApp function
type ISubwalletClient interface {
	RegisterSubwallet(*SubwalletRegisterRequest, http.Header) error
	UpdateDApp(*DAppUpdateRequest, http.Header) error
	DAppList(identifier, page, num string, header http.Header) (*DAppListResp, error)
	DAppinfo(identifier string, header http.Header) (*DAppListResp, error)
}

type RegisterRequest struct {
	Credential RegisterBody `json:"credential,omitempty"`
}

type RegisterBody struct {
	Role        string `json:"role,omitempty"`
	Description string `json:"description,omitempty"`
	ChannelId   string `json:"channel_id,omitempty"`
	Value       User   `json:"value,omitempty"`
}

type User struct {
	Access     string      `json:"access,omitempty"` // username
	Phone      string      `json:"phone,omitempty"`
	Email      string      `json:"email,omitempty"`
	Secret     string      `json:"secret,omitempty"` // password
	Identifier string      `json:"identifier,omitempty"`
	UserType   uint        `json:"user_type,omitempty"`
	Metadata   interface{} `json:"meta_data,omitempty"`
}

type UserInfo struct {
	Id                 string          `json:"id,omitempty"`
	GroupID            string          `json:"group_id,omitempty"`
	Access             string          `json:"access,omitempty"`
	Phone              string          `json:"phone,omitempty"`
	Email              string          `json:"email,omitempty"`
	Identifier         string          `json:"identifier,omitempty"`
	Status             string          `json:"status,omitempty"`
	Roles              string          `json:"roles,omitempty"`
	VerificationStatus uint            `json:"verification_status,omitempty"`
	Remark             string          `json:"remark,omitempty"`
	UserType           uint            `json:"user_type,omitempty"`
	EnterpriseInfo     *EnterpriseInfo `json:"enterprise_info,omitempty"`
	Issued_at          int64           `json:"issued_at,omitempty"`
	Channel_id         string          `json:"channel_id,omitempty"`
	CreatedAt          int64           `json:"created_at,omitempty"`
	UpdateAt           int64           `json:"update_at,omitempty"`
}

type EnterpriseInfo struct {
	CompanyName            string `json:"company_name,omitempty"`
	CompanyAddr            string `json:"company_address,omitempty"`
	TaxCode                string `json:"tax_code,omitempty"`
	JuridicalPersonName    string `json:"juridical_person_name,omitempty"`
	JuridicalPersonTel     string `json:"juridical_person_tel,omitempty"`
	JuridicalPersonID      string `json:"juridical_person_id,omitempty"`
	Licence                string `json:"licence,omitempty"`
	LicenceFileName        string `json:"licence_file_name,omitempty"`
	LicenceNumber          string `json:"licence_number,omitempty"`
	IdentityTop            string `json:"identity_top,omitempty"`
	IdentityTopFileName    string `json:"identity_top_file_type,omitempty"`
	IdentityBottom         string `json:"identity_bottom,omitempty"`
	IdentityBottomFileName string `json:"identity_bottom_file_type,omitempty"`
}

type LoginAccessSecret struct {
	Access string `json:"access,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Email  string `json:"email,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type LoginAccess struct {
	Value LoginAccessSecret
}

type UpdatePasswordRequest struct {
	Credential *UpdatePasswordBody `json:"credential,omitempty"`
}

type UpdatePasswordBody struct {
	OriginalSecret string `json:"original_secret,omitempty"`
	Identifier     string `json:"identifier,omitempty"`
	NewSecret      string `json:"new_secret,omitempty"`
	Access         string `json:"access,omitempty"`
	Email          string `json:"email,omitempty"`
	Phone          string `json:"phone,omitempty"`
}

// LoginRequest is token request
type LoginRequest struct {
	Credential LoginAccess `json:"credential,omitempty"`
}

type RevokeRequest struct {
	Credential *User `json:"credential,omitempty"`
}

type ChannelIDAccess struct {
	ChannelID string `json:"channel_id,omitempty"`
	//Access     string `json:"access,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type UpdateChannelIDRequest struct {
	Credential *ChannelIDAccess `json:"credential,omitempty"`
}

type CertCreateReq struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ApiKey      string `json:"api_key,omitempty"`
}

type CertCreateReqBody struct {
	Certificate CertCreateReq `json:"certificate,omitempty"`
}

type Cert struct {
	Key    string `json:"key,omitempty"`
	Cert   string `json:"cert,omitempty"`
	CACert string `json:"ca_cert,omitempty"`
}

// cert resp info
type CertCreateResp struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ApiKey      string `json:"api_key,omitempty"`
	IssuedAt    int64  `json:"issued_at,omitempty"`
	Hash        string `json:"hash,omitempty"`
	Status      string `json:"status,omitempty"`
	Value       *Cert  `json:"value,omitempty"`
}

// CertCreateRespBody is cert resp info
type CertCreateRespBody struct {
	Certificate *CertCreateResp `json:"certificate,omitempty"`
}

// CertInfo
type CertInfo struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ApiKey      string `json:"api_key,omitempty"`
	IssuedAt    int64  `json:"issued_at,omitempty"`
	Hash        string `json:"hash,omitempty"`
	Status      string `json:"status,omitempty"`
}

// CertListRespBody return all certs list
type CertListRespBody struct {
	Certificates []*CertInfo `json:"certificates,omitempty"`
}

// GetCertStutasResp is cert resp status
type GetCertStatusResp struct {
	Certificate GetCertStatusRespInner `json:"certificate,omitempty"`
}

// GetCertStutasRespInner is cert resp issued_at
type GetCertStatusRespInner struct {
	Isued_at int64 `json:"issued_at,omitempty"`
}

type Oauth2AccessToken struct {
	AccessToken string `json:"access_token,omitempty"`
}

// 获取token成功时返回的token详情
type TokenDetail struct {
	IssuedAt     int64  `json:"issue_at,omitempty"`
	ExpiresAt    int64  `json:"expires_at,omitempty"`
	Value        string `json:"value,omitempty"`
	CredentialId string `json:"credential_id,omitempty"`
	Roles        string `json:"roles,omitempty"`
	Identifier   string `json:"identifier,omitempty"`
	GroupID      uint   `json:"group_id,omitempty"`
	ChannelId    string `json:"channel_id,omitempty"`
}

// 获取token成功时返回的response
type GetTokenResponse struct {
	Token TokenDetail `json:"token,omitempty"`
}

// 通过token能获取的用户信息
type TokenInfo struct {
	CredentialId string `json:"credential_id,omitempty"`
	Access       string `json:"access,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
	Roles        string `json:"roles,omitempty"`
	Identifier   string `json:"identifier,omitempty"`
	GroupID      uint   `json:"group_id,omitempty"`
	ChannelID    string `json:"channel_id,omitempty"`
}

type ResponseStruct struct {
	Credentials CredentialsStruct `json:"credentials,omitempty"`
}
type CredentialsStruct struct {
	Id          string      `json:"id,omitempty"`
	Role        string      `json:"role,omitempty"`
	Description string      `json:"description,omitempty"`
	Channel_id  string      `json:"channel_id,omitempty"`
	Issued_at   int64       `json:"issued_at,omitempty"`
	Value       ValueStruct `json:"value,omitempty"`
}

type CredentialsListResponse struct {
	Credentials []UserInfo `json:"credentials,omitempty"`
}

type ValueStruct struct {
	Access     string `json:"access,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	//Secret string `json:"secret"`
}

type AccesskeyRequest struct {
	Accesskey AccesskeyRequestInner `json:"accesskey,omitempty"`
}

type AccesskeyRequestInner struct {
	Description string `json:"description,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
}

type AccesskeyResponse struct {
	Accesskey AccesskeyResponseInner `json:"accesskey,omitempty"`
}

type AccesskeyResponseInner struct {
	Id          string         `json:"id,omitempty"`
	Description string         `json:"description,omitempty"`
	Issued_at   int64          `json:"issued_at,omitempty"`
	Value       AccessKeyValue `json:"value,omitempty"`
}

type AccessKeyValue struct {
	ApiKey     string `json:"api_key,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type ACLResponse struct {
	Access     string `json:"access,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Roles      string `json:"roles,omitempty"`
	ChannelId  string `json:"channel_id,omitempty"`
}

// ACLGroups ...
type ACLGroups struct {
	Groups []ACLGroup `json:"groups,omitempty"`
}

// ACLGroup ...
type ACLGroup struct {
	ID          *uint  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   int64  `json:"created_at,omitempty"`
	UpdateAt    int64  `json:"updated_at,omitempty"`
}

// ACLResourceRequest ...
type ACLResourceRequest struct {
	GroupID   uint   `json:"group_id,omitempty"`
	Resources []uint `json:"resources,omitempty"`
}

// ACLGroupResource ...
type ACLGroupResource struct {
	Group     ACLGroup      `json:"group,omitempty"`
	Resources []ACLResource `json:"resources,omitempty"`
}

// ACLResource ...
type ACLResource struct {
	ID      uint   `json:"id,omitempty"`
	Service string `json:"service,omitempty"`
	Path    string `json:"path,omitempty"`
	Method  string `json:"method,moitempty"`
}

// UpdateUserGroupRequest ...
type UpdateUserGroupRequest struct {
	Users []UserInfo `json:"users,omitempty"`
}

// UsersNumResponse ...
type UsersNumResponse struct {
	UsersNum int `json:"users_num,omitempty"`
}

// DappUsersNumResponse ...
type DappUsersNumResponse struct {
	DappUsersNum int `json:"dapp_users_num,omitempty"`
}

// DappInfoResponse slice should be return when query users info list
type DappInfoResponse struct {
	Identifier         string `json:"id,omitempty"`
	Access             string `json:"access,omitempty"`
	Email              string `json:"email,omitempty"`
	Phone              string `json:"phone,omitempty"`
	GroupID            uint   `json:"group_id,omitempty"`
	UserType           uint   `json:"user_type,omitempty"`
	ChannelID          string `json:"channel_id,omitempty"`
	Description        string `json:"description,omitempty"`
	VerificationStatus string `json:"verification_status"`
	MetaData           string `json:"meta_data,omitempty"`
}

// UsersGrowthResponse  slice should be return when query users growth
type UsersGrowthResponse struct {
	DateTime string `json:"datetime,omitempty"`
	// the num of new users created during datetime
	GrowthAmount int `json:"growth_amount,omitempty"`
}

// SubwalletRegisterRequest ...
type SubwalletRegisterRequest struct {
	Identifier      string `json:"identifier,omitempty"`
	ParentIdentifer string `json:"parent_identifier,omitempty"`
	SubwalletType   int32  `json:"subwallet_type,omitempty"`
}

// DAppUpdateRequest ...
type DAppUpdateRequest struct {
	Identifier  string `json:"identifier,omitempty"`
	Name        string `json:"name,omitepmty"`
	Type        int    `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	IconFile    []byte `json:"icon,omitempty"`
	IconName    string `json:"icon_name" gorm:"type:varchar(32)"`
}

// DAppListResp ...
type DAppListResp struct {
	ID               int64  `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"`
	Name             string `json:"name.omitempty" gorm:"type:varchar(32)"`
	ApiKey           string `json:"api_key,omitempty" gorm:"type:varchar(128)"`
	Identifier       string `json:"identifier,omitempty" gorm:"type:varchar(128);"`
	ParentIdentifier string `json:"parent_identifier,omitempty" gorm:"type:varchar(128)"`
	Type             int    `json:"type,omitempty"` //包括去中心化交易所，游戏，其他等
	Description      string `json:"description,omitempty" gorm:"type:varchar(128)"`
	IconFile         []byte `json:"icon_file" gorm:"type:bytea"`
	IconName         string `json:"icon_name" gorm:"type:varchar(32)"`
	IssuedAt         int64  `json:"issued_at,omitempty"`
}
