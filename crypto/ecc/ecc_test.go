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

package ecc

import (
	"testing"

	"github.com/arxanchain/sdk-go-common/crypto/ecc/primitives"
)

func TestECCSignAndVerify(t *testing.T) {
	t.Log("running TestECCSignAndVerify")

	message := []byte("hello world")

	primitives.InitSecurityLevel("SHA3", 256)
	crypto, err := NewECCCryptoLib("./test/enrollment.key", "./test/enrollment.cert")
	if err != nil {
		t.Fatalf("new ecc crypto error: %v", err)
	}
	if crypto == nil {
		t.Fatal("ecc crypto instance invalide")
	}

	sign, err := crypto.Sign(message)
	if err != nil {
		t.Fatalf("ecc sign error: %v", err)
	}

	t.Logf("sign result: %v", sign)

	err = crypto.Verify(message, sign)
	if err != nil {
		t.Fatalf("ecc verify error: %v", err)
	}
}

func TestECCEnryptAndDecrypt(t *testing.T) {
	t.Log("running TestECCEnryptAndDecrypt")

	message := []byte("hello world")

	primitives.InitSecurityLevel("SHA3", 256)
	crypto, err := NewECCCryptoLib("./test/enrollment.key", "./test/enrollment.cert")
	if err != nil {
		t.Fatalf("new ecc crypto error: %v", err)
	}
	if crypto == nil {
		t.Fatal("ecc crypto instance invalide")
	}

	cipher, err := crypto.Encrypt(message)
	if err != nil {
		t.Fatalf("ecc encrypt error: %v", err)
	}

	t.Logf("encrypt result: %v", cipher)

	oriData, err := crypto.Decrypt(cipher)
	if err != nil {
		t.Fatalf("ecc decrypt error: %v", err)
	}

	t.Logf("decrypt result: %s", oriData)
}
