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
	"crypto/rand"
	"fmt"

	logging "github.com/op/go-logging"
	edAlg "golang.org/x/crypto/ed25519"
)

const (
	// KeyType_PubRSA public key of rsa
	KeyType_PubRSA = "RsaPublicKey"
	// KeyType_PriRSA private key of rsa
	KeyType_PriRSA = "RsaPrivateKey"
	// KeyType_PubED25119 public key of ed25519
	KeyType_PubED25119 = "EdDsaPublicKey"
	// KeyType_PriED25119 private key of ed25519
	KeyType_PriED25119 = "EdDsaPrivateKey"
)

const (
	// Usage_EncDec used for encrypt and decrypt
	Usage_EncDec = "EncryptDecrypt"
	// Usage_SignVerify used for signature and verify
	Usage_SignVerify = "SignVerify"
	// Usage_Other used for all other unused crypto option
	Usage_Other = "Other"
)

const (
	// PublicKeySize is the size, in bytes, of public keys as used in this package.
	PublicKeySize = 32
	// PrivateKeySize is the size, in bytes, of private keys as used in this package.
	PrivateKeySize = 64
	// SignatureSize is the size, in bytes, of signatures generated and verified by this package.
	SignatureSize = 64
)

var (
	logger = logging.MustGetLogger("ed25519")
)

// PublicKey store public key and its information
type PublicKey struct {
	Usage         string `json:"usage"`
	KeyType       string `json:"key_type"`
	PublicKeyData []byte `json:"public_key_data"`
}

//PrivateKey store private key and its information
type PrivateKey struct {
	Usage          string `json:"usage"`
	KeyType        string `json:"key_type"`
	PrivateKeyData []byte `json:"private_key_data"`
}

// KeyPair generate keypair of ed25519
func Keypair() (pub *PublicKey, pri *PrivateKey, err error) {
	publicKey, privateKey, err := edAlg.GenerateKey(rand.Reader)
	if nil != err {
		logger.Errorf("Failed to generate ed25519 keypair: %v", err)
		return pub, pri, err
	}

	pub = &PublicKey{
		PublicKeyData: publicKey,
	}

	pri = &PrivateKey{
		PrivateKeyData: privateKey,
	}

	return pub, pri, nil
}

// GetUsage return usage of the key
func (privateKey *PrivateKey) GetUsage() string {
	return privateKey.Usage
}

// GetType return type of the key
func (privateKey *PrivateKey) GetType() string {
	return privateKey.KeyType
}

// GetRawData return the key
func (privateKey *PrivateKey) GetRawData() []byte {
	return privateKey.PrivateKeyData
}

// Sign data with the privateKey
func (privateKey *PrivateKey) Sign(message []byte) ([]byte, error) {
	return edAlg.Sign(privateKey.PrivateKeyData, message), nil
}

// GetUsage return usage of the key
func (publicKey *PublicKey) GetUsage() string {
	return publicKey.Usage
}

// GetType return type of the key
func (publicKey *PublicKey) GetType() string {
	return publicKey.KeyType
}

// GetRawData return the key
func (publicKey *PublicKey) GetRawData() []byte {
	return publicKey.PublicKeyData
}

// Verify signature with the publicKey
func (publicKey *PublicKey) Verify(message []byte, signedMessage []byte) error {
	if len(signedMessage) != SignatureSize || signedMessage[63]&224 != 0 {
		logger.Error("signedMessage is wrong format")
		return fmt.Errorf("signedMessage is wrong format")
	}

	result := edAlg.Verify(publicKey.PublicKeyData, message, signedMessage)
	if !result {
		logger.Error("Failed to verify the signature")
		return fmt.Errorf("failed to verify the signature")
	}
	return nil
}
