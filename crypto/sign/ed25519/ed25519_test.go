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

package ed25519

import (
	"testing"
)

func TestSignVerify(t *testing.T) {
	//generate keypair
	pub, pri, err := Keypair()
	if nil != err {
		t.Errorf("ed25519 generate keypair error")
	}

	message := "hello world"

	//signature
	signedData, _ := pri.Sign([]byte(message))

	//verify
	result := pub.Verify([]byte(message), signedData)

	if nil != result {
		t.Errorf("ed25519 verify error")
	}
}
