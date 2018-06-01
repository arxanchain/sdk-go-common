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

// common http header define
const (
	UserIdHeader        = "User-Id"
	UserRoleHeader      = "User-Role"
	ChannelIdHeader     = "Channel-Id"
	ChaincodeIdHeader   = "Chaincode-Id"
	XAuthTokenHeader    = "X-Auth-Token"
	XSubjectTokenHeader = "X-Subject-Token"
	ACLActionHeader     = "ACL-Action"
	APIKeyHeader        = "API-Key"
	EnrollmentIdHeader  = "Enrollment-Id"
	CryptoModeHeader    = "Crypto-Mode"
	AuthModeHeader      = "Auth-Mode"
	// defined a new auth mode to auth channel/chaincode, for we need auth token and channel/chaincode
	AuthChannelModeHeader = "Auth-ChCC-Mode"
	// defined a new auth mode to auth SN
	AuthSNModeHeader    = "Auth-SN-Mode"
	AuthEmailHeader     = "Auth-Email"
	FileAuthTokenHeader = "Auth-Token"
	FabioRouteTagHeader = "Host"
	CallbackUrlHeader   = "Callback-Url"
	RouteTagHeader      = "Route-Tag"
	InvokeModeHeader    = "Bc-Invoke-Mode"
)

// User role Header value list
const (
	UserRoleSuperAdmin = "super_admin"
	UserRoleAdmin      = "admin"
	UserRoleNormal     = "normal"
	UserRoleDid        = "did"
)

// User GroupID value list
const (
	UserRoleGroupIDSuperAdmin uint = 1
)

// Blockchain Invoke Mode header value list
const (
	InvokeModeSync  = "sync"
	InvokeModeAsync = "async"
)
