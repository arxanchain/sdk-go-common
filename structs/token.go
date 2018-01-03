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

type TokenCreateRequest struct {
	// The name of the token role
	Role string `json:"role,omitempty"`
	// A list of policies for the token
	Policies []string `json:"policies,omitempty"`
	NoParent bool     `json:"no_parent,omitempty"`
	//  A map of string to string valued metadata. This is passed through to the audit backends
	MetaData map[string]string `json:"meta_data,omitempty"`
	// If set, the token will have an explicit max TTL set upon it. This maximum token TTL cannot be changed later,
	// and unlike with normal tokens, updates to the system/mount max TTL value will have no effect at renewal time
	// the token will never be able to be renewed or used past the value set at issue time
	MaxTTL string `json:"max_ttl,omitempty"`
	// Set to false to disable the ability of the token to be renewed past its initial TTL.
	// Setting the value to true will allow the token to be renewable up to the system/mount maximum TTL
	Renewable *bool `json:"renewable,omitempty"`
	// The TTL period of the token, provided as "1h", where hour is the largest suffix.
	// If not provided, the token is valid for the default lease TTL, or indefinitely if the root policy is used
	TTL string `json:"ttl,omitempty"`
	// The maximum uses for the given token. This can be used to create a one-time-token or limited use token.
	// The value of 0 has no limit to the number of uses
	NumUsers int `json:"num_users,omitempty"`
}
