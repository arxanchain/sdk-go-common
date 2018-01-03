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
	"testing"

	"github.com/arxanchain/sdk-go-common/crypto/sign/ed25519"
)

func TestSignVerify(t *testing.T) {
	var sd = SignedData{
		Data: []byte("hello world"),
		Header: &SignatureHeader{
			SignType: "ed25519-with-sha512",
			Creator:  "did<21tDAKCERh95uGgKbJNHYp>",
			Nonce:    []byte("ssssss"),
		},
	}

	// generate key ed25519 pair
	pub, pri, err := ed25519.Keypair()
	if nil != err {
		t.Errorf("ed25519 generate keypair error")
	}

	// sign
	_, err = sd.DoSign(pri)
	if nil != err {
		t.Errorf("ed25519 sign error")
	}

	// verify
	err = sd.Verify(pub)
	if nil != err {
		t.Errorf("ed25519 verify error")
	}

	// faile verify
	for i, _ := range sd.Data {
		sd.Data[i] = 0
	}
	err = sd.Verify(pub)
	if nil == err {
		t.Errorf("ed25519 verify error")
	}

}
