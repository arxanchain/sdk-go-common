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
package pki

// ICryptoLib function set of public and private keys
type ICryptoLib interface {
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, sig []byte) error
	Encrypt(data []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

// IPublicKey function sets of public key
type IPublicKey interface {
	GetUsage() string
	GetType() string
	GetRawData() []byte
	Verify(data []byte, signature []byte) error
}

// IPrivateKey function sets of private key
type IPrivateKey interface {
	GetUsage() string
	GetType() string
	GetRawData() []byte
	Sign(data []byte) ([]byte, error)
}
