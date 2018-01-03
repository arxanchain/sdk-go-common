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
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	logging "github.com/op/go-logging"
)

var (
	logger = logging.MustGetLogger("rsa")
)

// RSACryptoLib represents the crypto util.
type RSACryptoLib struct {
	privateKey *rsa.PrivateKey
	peerCert   *x509.Certificate
}

// NewRSACryptoLib create a new RSACryptoLib, and load private key and certificate from java PKCS12 keystore
func NewRSACryptoLib(privateKeyFile string, publicCertFile string) (*RSACryptoLib, error) {
	privateKey, err := loadPrivateKey(privateKeyFile)
	if err != nil {
		logger.Errorf("Failed to load private key(%s): %+v", privateKeyFile, err)
		return nil, err
	}

	cert, err := loadX509Cert(publicCertFile)
	if err != nil {
		logger.Errorf("Failed to load certificate(%s): %+v", publicCertFile, err)
		return nil, err
	}

	return &RSACryptoLib{privateKey, cert}, nil
}

func (c *RSACryptoLib) Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)
	rng := rand.Reader
	return rsa.SignPKCS1v15(rng, c.privateKey, crypto.SHA256, hashed[:])
}

func (c *RSACryptoLib) Verify(data []byte, sig []byte) error {
	hashed := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(c.peerCert.PublicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], sig)
}

// RSA Encrypt
func (c *RSACryptoLib) Encrypt(data []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, c.peerCert.PublicKey.(*rsa.PublicKey), data, nil)
}

// RSA Decrypt
func (c *RSACryptoLib) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, c.privateKey, ciphertext, nil)
}

func loadPrivateKey(filename string) (*rsa.PrivateKey, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(raw)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func loadX509Cert(filename string) (*x509.Certificate, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Errorf("Failed to read x509 certificate file %s! %+v", filename, err)
		return nil, err
	}

	block, _ := pem.Decode(raw)
	return x509.ParseCertificate(block.Bytes)
}
