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

package did

type (
	// IDIdentity format "did:<provider short id>:id_string"
	// IVerifiableClaim format "vcm:<adapter shrot id>:id_string"
	Identifier  string
	DidEndpoint string
	DidContext  string
)

// Did type define
type DidType string

const (
	// Common did type (Main wallet type)
	DTAsset        DidType = "Asset"
	DTEntity       DidType = "Entity"
	DTOrganization DidType = "Organization"
	DTPerson       DidType = "Person"      //Individual
	DTIndependent  DidType = "Independent" //Independent person
	DTDependent    DidType = "Dependent"   //Dependent person
	DTSerialNumber DidType = "SerialNumber"

	// Sub wallet did type
	DTCash     DidType = "cash"
	DTFee      DidType = "fee"
	DTLoan     DidType = "loan"
	DTInterest DidType = "interest"
)

// Did status define
type DidStatus string

const (
	DSValid   DidStatus = "Valid"
	DSInvalid DidStatus = "Invalid"
	DSIssued  DidStatus = "Issued"
)
