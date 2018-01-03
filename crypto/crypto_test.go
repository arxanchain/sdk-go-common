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

package crypto

import (
	"testing"

	"github.com/arxanchain/sdk-go-common/crypto/sign/ed25519"
	"github.com/arxanchain/sdk-go-common/structs"
)

func TestVerifySignatureED25519(t *testing.T) {
	sBody := &structs.SignatureBody{
		Creator:        "did:ara:8uQhQMGzWxR8vw5P3UWH1j",
		Created:        "ssss",
		Nonce:          "nonce",
		SignatureValue: []byte{144, 110, 161, 2, 122, 117, 218, 88, 94, 172, 199, 101, 32, 25, 130, 245, 116, 221, 51, 25, 14, 41, 123, 219, 134, 10, 32, 113, 127, 155, 62, 225, 139, 85, 105, 182, 86, 177, 35, 222, 116, 254, 252, 44, 100, 194, 97, 182, 215, 55, 184, 233, 120, 12, 34, 74, 218, 52, 225, 212, 165, 149, 43, 0},
	}

	wr := &structs.WalletRequest{
		Payload:   "{\"enrollment_id\":\"1f38a7a1-2c79-465e-a4c0-0038e25c7edg\",\"callback_url\":\"http://127.0.0.1\",\"from\":\"did:ara:8uQhQMGzWxR8vw5P3UWH1j\",\"to\":\"did:ara:21tDAKCERh95uGgKbJNHYp\",\"asset_id\":\"1f38a7a1-2c79-465e-a4c0-0038e25c7edg\",\"coins\":[{\"coin_id\":\"1f38a7a1-2c79-465e-a4c0-0038e25c7edg\",\"amount\":5}],\"fees\":{\"accounts\":[\"did:ara:8uQhQMGzWxR8vw5P3UWH1j\"],\"coins\":[{\"coin_id\":\"1f38a7a1-2c79-465e-a4c0-0038e25c7edg\",\"amount\":5}]}}",
		Signature: sBody,
	}

	ipk := &ed25519.PublicKey{
		PublicKeyData: []byte{244, 28, 212, 44, 176, 130, 24, 135, 231, 54, 10, 0, 138, 16, 119, 148, 191, 26, 130, 187, 165, 191, 176, 146, 170, 42, 221, 131, 44, 153, 202, 167},
	}

	err := VerifySignatureED25519(wr, ipk)
	if nil != err {
		t.Errorf("VerifySignatureED25519 fail")
	}

}
