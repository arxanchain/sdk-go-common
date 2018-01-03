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
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	ecv2 "github.com/arxanchain/sdk-go-common/crypto/ecc/eciesv2"
	"github.com/arxanchain/sdk-go-common/crypto/ecc/primitives"
	"github.com/arxanchain/sdk-go-common/crypto/ecc/primitives/ecies"
	logging "github.com/op/go-logging"
)

var (
	logger      = logging.MustGetLogger("ecc")
	encryptMode = "ecv2"
)

// ECCCryptoLib represents the crypto util.
type ECCCryptoLib struct {
	privateKey *ecdsa.PrivateKey
	peerCert   *x509.Certificate
}

// NewECCCryptoLib create a new ECCCryptoLib, and load private key and certificate from java PKCS12 keystore
func NewECCCryptoLib(privateKeyFile string, publicCertFile string) (*ECCCryptoLib, error) {
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

	return &ECCCryptoLib{privateKey, cert}, nil
}

func (c *ECCCryptoLib) Sign(data []byte) ([]byte, error) {
	signData, err := primitives.ECDSASign(c.privateKey, data)
	if err != nil {
		logger.Errorf("Failed to sign the data: %v", err)
		return nil, err
	}
	return signData, nil
}

func (c *ECCCryptoLib) Verify(data []byte, sig []byte) error {
	verified, err := primitives.ECDSAVerify(c.peerCert.PublicKey.(*ecdsa.PublicKey), data, sig)
	if err != nil {
		logger.Errorf("Failed to verify the data: %v", err)
		return err
	}

	if !verified {
		logger.Error("Failed to verify the signature")
		return fmt.Errorf("failed to verify the signature")
	}

	return nil
}

// RSA Encrypt
func (c *ECCCryptoLib) Encrypt(data []byte) (out []byte, err error) {
	if encryptMode == "ecv2" {
		pk := ecv2.ImportECDSAPublic(c.peerCert.PublicKey.(*ecdsa.PublicKey))
		out, err = ecv2.Encrypt(rand.Reader, pk, data, nil, nil)
		if err != nil {
			logger.Errorf("Failed to encrypt data: %v", err)
			return nil, err
		}
	} else {
		spi := ecies.NewSPI()
		eciesKey, err := spi.NewPublicKey(nil, c.peerCert.PublicKey.(*ecdsa.PublicKey))
		if err != nil {
			logger.Errorf("Failed to create SPI public key: %v", err)
			return nil, err
		}
		ecies, err := spi.NewAsymmetricCipherFromPublicKey(eciesKey)
		if err != nil {
			logger.Errorf("Failed to create Asymmetric Cipher: %v", err)
			return nil, err
		}
		out, err = ecies.Process(data)
		if err != nil {
			logger.Errorf("Failed to encrypt data: %v", err)
			return nil, err
		}
	}

	return out, nil
}

// RSA Decrypt
func (c *ECCCryptoLib) Decrypt(ciphertext []byte) (out []byte, err error) {
	if encryptMode == "ecv2" {
		privk := ecv2.ImportECDSA(c.privateKey)
		out, err = privk.Decrypt(rand.Reader, ciphertext, nil, nil)
		if err != nil {
			logger.Errorf("Failed to decrypt data: %v", err)
			return nil, err
		}
	} else {
		spi := ecies.NewSPI()
		eciesKey, err := spi.NewPrivateKey(nil, c.privateKey)
		if err != nil {
			logger.Errorf("Failed to create SPI private key: %v", err)
			return nil, err
		}
		ecies, err := spi.NewAsymmetricCipherFromPublicKey(eciesKey)
		if err != nil {
			logger.Errorf("Failed to create NewAsymmetricCipherFromPublicKey: %v", err)
			return nil, err
		}
		out, err = ecies.Process(ciphertext)
		if err != nil {
			logger.Errorf("Failed to decrypt data: %v", err)
			return nil, err
		}
	}

	return out, nil
}

func loadPrivateKey(filename string) (*ecdsa.PrivateKey, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(raw)
	if block == nil {
		logger.Error("failed to decode EC Private Key")
		return nil, fmt.Errorf("failed to decode EC Private Key")
	}

	return x509.ParseECPrivateKey(block.Bytes)
}

func loadX509Cert(filename string) (*x509.Certificate, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Errorf("Failed to read x509 certificate file %s! %+v", filename, err)
		return nil, err
	}

	block, _ := pem.Decode(raw)
	if block == nil {
		logger.Error("failed to decode EC CERTIFICATE")
		return nil, fmt.Errorf("failed to decode EC CERTIFICATE")
	}

	return x509.ParseCertificate(block.Bytes)
}
