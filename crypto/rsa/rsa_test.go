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

package rsa

import (
	"testing"
)

func TestRSASignAndVerify(t *testing.T) {
	t.Log("running TestRSASignAndVerify")

	message := []byte("hello world")

	crypto, err := NewRSACryptoLib("./key.pem", "./cert.pem")
	if err != nil {
		t.Fatalf("new rsa crypto error: %v", err)
	}
	if crypto == nil {
		t.Fatal("rsa crypto instance invalide")
	}

	sign, err := crypto.Sign(message)
	if err != nil {
		t.Fatalf("rsa sign error: %v", err)
	}

	t.Logf("sign result: %v", sign)

	err = crypto.Verify(message, sign)
	if err != nil {
		t.Fatalf("rsa verify error: %v", err)
	}
}

func TestRSAEnryptAndDecrypt(t *testing.T) {
	t.Log("running TestRSAEnryptAndDecrypt")

	message := []byte("hello world")

	crypto, err := NewRSACryptoLib("./key.pem", "./cert.pem")
	if err != nil {
		t.Fatalf("new rsa crypto error: %v", err)
	}
	if crypto == nil {
		t.Fatal("rsa crypto instance invalide")
	}

	cipher, err := crypto.Encrypt(message)
	if err != nil {
		t.Fatalf("rsa encrypt error: %v", err)
	}

	t.Logf("encrypt result: %v", cipher)

	oriData, err := crypto.Decrypt(cipher)
	if err != nil {
		t.Fatalf("rsa decrypt error: %v", err)
	}

	t.Logf("decrypt result: %s", oriData)
}
