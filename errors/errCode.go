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

package errors

type ErrCodeType uint32

//Success code
const SuccCode ErrCodeType = 0

// common error code and message defined
const (
	InvalidParamsErrCode    ErrCodeType = 1000 //参数无效
	MissingParamsErrCode    ErrCodeType = 1001 //缺少参数
	DatabaseOperationFailed ErrCodeType = 1002 //数据库操作失败
	ParseRequestParamsError ErrCodeType = 1003 //解析请求体失败
	SerializeDataFail       ErrCodeType = 1004 //序列化数据失败
	DeserializeDataFail     ErrCodeType = 1005 //反序列化(解析)数据失败
	GetServerContextFail    ErrCodeType = 1006 //获取服务的上下文失败
	DatabaseUnavailable     ErrCodeType = 1007 //数据库不可用
	DatabaseDisabled        ErrCodeType = 1008 //数据库已禁用
	PermissionDenied        ErrCodeType = 1009 // 没有权限
	ED25519SignFail         ErrCodeType = 1010 //ED25519签名失败
	ED25519VerifyFail       ErrCodeType = 1011 //ED25519验签失败
	InternalServerFailure   ErrCodeType = 1012 //服务内部错误
)

// ccsandbox
const (
	InvalidRequestBody         ErrCodeType = 3000
	UnavailableServerContext   ErrCodeType = 3001
	FailedChaincodeRetrieval   ErrCodeType = 3002
	FailedFabricCCInstallation ErrCodeType = 3003
	FailedFabricCCInvocation   ErrCodeType = 3004
	FailedFabricCCQuery        ErrCodeType = 3005
	FailedFabricCCStop         ErrCodeType = 3006
)

//fred
const (
	RepeatRegistration       ErrCodeType = 4000 //重复注册
	UnmarshalFailed          ErrCodeType = 4003 //Unmarshal失败
	DataTypeIsIncorrect      ErrCodeType = 4004 //数据类型有误
	OriginalSecretError      ErrCodeType = 4005 //原始Secret错误
	UsersDoNotHavePermission ErrCodeType = 4006 //用户无权限
	CertificateUnavailable   ErrCodeType = 4007 //证书不可用
	NoSuchTypeOfUser         ErrCodeType = 4008 //无此类型用户
	InvalidAccessOrSecret    ErrCodeType = 4009 //无效的用户名或密码
	UserNotExist             ErrCodeType = 4010 //无此用户
	NoSuchTypeOfACLGroup     ErrCodeType = 4011 //无此类型权限组
	DeleteNotAllowed         ErrCodeType = 4012 //不允许删除
	FredUploadIdentityFailed ErrCodeType = 4013 //上传实名认证失败
	AuthTokenInvalid         ErrCodeType = 4014 //认证Token无效
)

//tomago
const (
	RegisterEntityFail      ErrCodeType = 5000 //注册实体失败
	UpdateEntityFail        ErrCodeType = 5001 //更新实体失败
	QueryEntityFail         ErrCodeType = 5002 //查询实体失败
	RegisterAssetFail       ErrCodeType = 5003 //注册资产失败
	UpdateAssetFail         ErrCodeType = 5004 //更新资产失败
	QueryAssetFail          ErrCodeType = 5005 //查询资产失败
	QueryAuditInfoFail      ErrCodeType = 5006 //查询审计信息失败
	AuditReverseFail        ErrCodeType = 5007 //审计平账操作失败
	ChargeInterestFail      ErrCodeType = 5008 //计息失败
	IssueCTokenFail         ErrCodeType = 5009 //发行Token失败
	RollbackTransactionFail ErrCodeType = 5010 //回滚交易失败
	TransferCTokenFail      ErrCodeType = 5011 //转账交易失败
	TransferAssetFail       ErrCodeType = 5012 //转移资产失败
	WithdrawFail            ErrCodeType = 5013 //取现失败
	ColoredCoinNotFound     ErrCodeType = 5014 //Token没有找到
	BalancesNotSufficient   ErrCodeType = 5015 //余额不足
	AssetHasNotOwner        ErrCodeType = 5016 //资产没有所属人
	AssetHasBeenIssued      ErrCodeType = 5017 //资产已发行
	AssetOwnerNotMatch      ErrCodeType = 5018 //资产所属人不匹配
	WalletStatusInvalid     ErrCodeType = 5019 //钱包状态异常
	EntityNotFound          ErrCodeType = 5020 //组织实体找不到
	AssetNotFound           ErrCodeType = 5021 //数字资产没有找到
	CTokenStatusNotInUse    ErrCodeType = 5022 //Token已回收
	CTokenAmountInvalid     ErrCodeType = 5023 //Token数额无效
)

//wallet-webserver
const (
	FailedToGenerateQRCode    ErrCodeType = 6000 //生成二维码失败
	ImgFormatConversionFailed ErrCodeType = 6001 //图片转码失败
	InvalidAssetPassword      ErrCodeType = 6002 //无效的资产密码
	ErrorMatchRulesEngine     ErrCodeType = 6003 //匹配资产规则失败
	NoMatchRulesEngine        ErrCodeType = 6004 //没有匹配资产规则
	InvalidPrivateKey         ErrCodeType = 6005 //无效的私钥
	InvalidSecurityCode       ErrCodeType = 6006 //无效的安全码
)

// wallet-ng error code and message defined
const (
	WalletNotFound             ErrCodeType = 8000 //钱包对象没有找到
	WalletGetCTokensFail       ErrCodeType = 8001 //获取钱包Token失败
	WalletGetAssetsFail        ErrCodeType = 8002 //获取钱包资产失败
	WalletTransferCTokensFail  ErrCodeType = 8003 //钱包转账失败
	WalletTransferAssetsFail   ErrCodeType = 8004 //钱包转移资产失败
	CreateMainWalletFail       ErrCodeType = 8005 //创建主钱包失败
	CreateSubWalletFail        ErrCodeType = 8006 //创建子钱包失败
	GetPublicKeyFail           ErrCodeType = 8007 //获取公钥失败
	OffchainReadUploadFileFail ErrCodeType = 8008 //获取上传的文件失败
	OffchainSaveFileFail       ErrCodeType = 8009 //保存文件失败
	OffchainDIDTypeInvalid     ErrCodeType = 8010 //不正确的资产类型
	OverMaxActivedCount        ErrCodeType = 8011 // SN超过最大激活次数
	SNExpired                  ErrCodeType = 8012 // SN已过期
	OverMaxUploadSize          ErrCodeType = 8013 // 上传文件超过最大文件大小
	OffchainReadOnly           ErrCodeType = 8014 // offchain metadata只允许读,不允许更新
	OffchainContentExist       ErrCodeType = 8015 // offchain file content已经存在
	TransactionInvalid         ErrCodeType = 8016 // Blockchain transaction invalid
	QueryTxStatusTimeout       ErrCodeType = 8017 // query transaction status timeout
)

// chain-mgmt error code and message defined
const (
	ChannelAlreadyCreated ErrCodeType = 9000 // channel已经创建成功
	ChaincodeIDMissing    ErrCodeType = 9001
	ChaincodePathMissing  ErrCodeType = 9002
	CTypeUnrecognized     ErrCodeType = 9003
	MPartUnrecognized     ErrCodeType = 9004
	FileNotAccessible     ErrCodeType = 9005
	CCAlreadyCreated      ErrCodeType = 9006 // chaincode已经创建
	CCDeployRecordExisted ErrCodeType = 9007 // chaincode deploy记录存在
	CCDeleteDeployed      ErrCodeType = 9008 // 删除已经部署的chaincode source
	ChaincodeNotExist     ErrCodeType = 9009 // chaincode记录不存在
	CCSandBoxFailed       ErrCodeType = 9010 // chaincode沙箱测试失败
	CCNotDeployed         ErrCodeType = 9011 // chaincode没有部署,不允许upgrade
	CCAlreadyDeployed     ErrCodeType = 9012 // chaincode已经部署过,不允许部署
	CCUpgradeNoResource   ErrCodeType = 9013 // 无更新的chaincode资源去更新
)

// payment-gateway error code and message defined
const (
	RechargeFailed ErrCodeType = 2000 // 充值失败
	SendSMSFailed  ErrCodeType = 2001 // 发送短信验证码失败
	BindCardFailed ErrCodeType = 2002 // 绑定银行卡失败
)

// SDK error code
const (
	SDKServerResponseInvalid  ErrCodeType = 10000 // 服务端响应数据无效
	SDKDecryptAndVerifyFailed ErrCodeType = 10001 // SDK解密验签失败
	SDKInvalidBase64Data      ErrCodeType = 10002 // 私钥base64解码失败
)

// safebox srvc error code
const (
	UserInfoIsExist         ErrCodeType = 11000 // 用户已经存在
	UserInfoNotExit         ErrCodeType = 11001 // 用户信息不存在
	SecurityCodeErr         ErrCodeType = 11002 // 安全码错误
	OriginalSecurityCodeErr ErrCodeType = 11003 // 原始安全码错误
	SecurityInvalidErr      ErrCodeType = 11004 // 安全码无效
)

// escrow srvc error code
const (
	CurrencyTypeErr          ErrCodeType = 12000 // currency类型错误
	AddressFormatErr         ErrCodeType = 12001 // currency地址格式错误
	AddressExists            ErrCodeType = 12002 // currency地址已存在
	AddressNotExist          ErrCodeType = 12003 // currency地址不存在
	SMSCodeerror             ErrCodeType = 12004 // 短信验证码错误
	DepositAmountErr         ErrCodeType = 12005 // 充值金额错误
	WithdrawAmountErr        ErrCodeType = 12006 // 提现金额错误
	WithdrawFeesAmountErr    ErrCodeType = 12007 // 提现手续费错误
	TradePwdErr              ErrCodeType = 12008 // 交易密码错误
	EmailCodeErr             ErrCodeType = 12009 // 邮件验证码错误
	CurrencyTypeExist        ErrCodeType = 12010 // currency已存在
	SignRawTransactionFailed ErrCodeType = 12011 // 交易签名失败
	GenerateAddressFailed    ErrCodeType = 12012 // 生成地址失败
	StorageKeysotreFailed    ErrCodeType = 12013 // key存储keystore失败
	InsufficientBalance      ErrCodeType = 12014 // 余额不足
)
