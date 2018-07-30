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

import (
	commdid "github.com/arxanchain/sdk-go-common/structs/did"
)

// IndexTags ...
type IndexTags struct {
	CombinedIndex   []string `json:"combined_index,omitempty"`
	IndividualIndex []string `json:"individual_index,omitempty"`
}

// IndexSetPayload represents the payload of IndexSet interface
type IndexSetPayload struct {
	Id     commdid.Identifier `json:"id,omitempty"`
	Indexs *IndexTags         `json:"indexs,omitempty"`
}

// IndexGetPayload represents the payload of IndexGet interface
type IndexGetPayload struct {
	Indexs *IndexTags `json:"indexs,omitempty"`
}
