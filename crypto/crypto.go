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
	"encoding/json"

	commdid "github.com/arxanchain/sdk-go-common/structs/did"
	"github.com/arxanchain/sdk-go-common/structs/pki"
	"github.com/arxanchain/sdk-go-common/structs/wallet"
	"github.com/arxanchain/sdk-go-common/utils"
)

type ServerClientMode int

const (
	SERVER_MODE ServerClientMode = iota
	CLIENT_MODE
)

type EncryptType int

const (
	ECC_TYPE EncryptType = iota
	RSA_TYPE
)

var (
	g_SecurityLevel    int              = 256
	g_HashAlgorithm    string           = "SHA3"
	g_ServerClientMode ServerClientMode = SERVER_MODE
	g_IsEncrypt        bool             = true
	g_IsSign           bool             = true
	g_EncryptType      EncryptType      = ECC_TYPE
)

type SignedData struct {
	Data      string `json:"data" yaml:"data"`
	Signature string `json:"signature" yaml:"signature"`
}

// SetSecurityLevel sets the security level and hash algorithm of cryptp library
//
// If not set, the default value is:
// 	 SecurityLevel: 256
//	 HashAlgorithm: "SHA3"
func SetSecurityLevel(securityLevel int, hashAlgorithm string) {
	if securityLevel != 0 {
		g_SecurityLevel = securityLevel
	}
	if hashAlgorithm != "" {
		g_HashAlgorithm = hashAlgorithm
	}
}

// SetServerClientMode sets the server or client mode
// 0: server mode
// 1: client mode
//
// If not set, the default value is 0.
func SetServerClientMode(mode ServerClientMode) {
	switch mode {
	case SERVER_MODE:
		fallthrough
	case CLIENT_MODE:
		g_ServerClientMode = mode
	}
}

// SetEncryptFlag sets the encryption flag
//
// If not set, the default value is "true".
func SetEncryptFlag(isEncrypt bool) {
	g_IsEncrypt = isEncrypt
}

// SetSignFlag sets the signature falg
//
// If not set, the default value is "true".
func SetSignFlag(isSign bool) {
	g_IsSign = isSign
}

// SetEncryptType sets the encryption type
// 0: ecc
// 1: rsa
//
// If not set, the default value is 0.
func SetEncryptType(encryptType EncryptType) {
	switch encryptType {
	case ECC_TYPE:
		fallthrough
	case RSA_TYPE:
		g_EncryptType = encryptType
	}
}

// DecryptAndVerify accepts data bytes and the specified enrollment id,
// decrypt the data and verify the signature according to the enrollmentId's certs,
// if success, return the raw plaintext data。
func DecryptAndVerify(dataBytes []byte, enrollmentID string) ([]byte, error) {
	var err error

	cryptoLib, err := GetCryptoLib(enrollmentID)
	if err != nil {
		return nil, err
	}

	cryptedDataBytes, err := utils.DecodeBase64(string(dataBytes))
	if err != nil {
		return nil, err
	}

	var rawDataWithSign []byte
	if g_IsEncrypt {
		logger.Debug("Decrypt data...")

		rawDataWithSign, err = cryptoLib.Decrypt(cryptedDataBytes)
		if err != nil {
			return nil, err
		}

		logger.Debug("Decrypt data succ")
	} else {
		rawDataWithSign = cryptedDataBytes
	}

	logger.Debugf("Data after decrypt: %s", string(rawDataWithSign))

	var signedData SignedData
	err = json.Unmarshal(rawDataWithSign, &signedData)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	logger.Debugf("dataBase64: [%s]", signedData.Data)
	logger.Debugf("signBase64: [%s]", signedData.Signature)

	rawData, err := utils.DecodeBase64(signedData.Data)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	rawSignature, err := utils.DecodeBase64(signedData.Signature)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	//verify signature
	if g_IsSign {
		err = cryptoLib.Verify(rawData, rawSignature)
		if err != nil {
			return nil, err
		}
	}

	return rawData, nil
}

// SignAndEncrypt accepts data bytes and the specified enrollment id,
// sign and encrypt the data according to the enrollmentId's certs,
// if success, return the encrypted data.
func SignAndEncrypt(databytes []byte, enrollmentID string) (string, error) {
	var err error

	cryptoLib, err := GetCryptoLib(enrollmentID)
	if err != nil {
		return "", err
	}

	//sign the data
	var signature []byte
	if g_IsSign {
		signature, err = cryptoLib.Sign(databytes)
		if err != nil {
			logger.Errorf(" - ERROR: Failed to sign the data: %v", err)
			return "", err
		}
	}

	// to base64
	dataBase64 := utils.EncodeBase64(databytes)
	signBase64 := utils.EncodeBase64(signature)
	logger.Debugf("signBase64: [%s]", signBase64)
	logger.Debugf("dataBase64: [%s]", dataBase64)

	signedData := new(SignedData)
	signedData.Data = dataBase64
	signedData.Signature = signBase64

	data, _ := json.Marshal(signedData)
	var out []byte

	if g_IsEncrypt {
		out, err = cryptoLib.Encrypt(data)
		if err != nil {
			logger.Errorf("RSA Encrypt error: %v", err)
			return "", err
		}
	} else {
		out = data
	}

	return utils.EncodeBase64(out), nil
}

// VerifySignatureED25519 verify signature of the given data
func VerifySignatureED25519(wr *wallet.WalletRequest, ipk pki.IPublicKey) error {
	var header = &pki.SignatureHeader{
		Creator: commdid.Identifier(wr.Signature.Creator),
		Nonce:   []byte(wr.Signature.Nonce),
	}

	signed, err := utils.DecodeBase64(wr.Signature.SignatureValue)
	if err != nil {
		logger.Error(err)
		return err
	}
	var sd = &pki.SignedData{
		Data:   []byte(wr.Payload),
		Header: header,
		Sign:   signed,
	}
	return sd.Verify(ipk)
}
