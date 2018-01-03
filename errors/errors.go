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

package errors

import (
	"bytes"
	"errors"
)

var (
	// ErrRegistrationRequired Registration to the Membership Service required.
	ErrRegistrationRequired = errors.New("Registration to the Membership Service required.")

	// ErrNotInitialized Initialization required
	ErrNotInitialized = errors.New("Initialization required.")

	// ErrAlreadyInitialized Already initialized
	ErrAlreadyInitialized = errors.New("Already initialized.")

	// ErrAlreadyRegistered Already registered
	ErrAlreadyRegistered = errors.New("Already registered.")

	// ErrTransactionMissingCert Transaction missing certificate or signature
	ErrTransactionMissingCert = errors.New("Transaction missing certificate or signature.")

	// ErrInvalidTransactionSignature Invalid Transaction Signature
	ErrInvalidTransactionSignature = errors.New("Invalid Transaction Signature.")

	// ErrTransactionCertificate Missing Transaction Certificate
	ErrTransactionCertificate = errors.New("Missing Transaction Certificate.")

	// ErrTransactionSignature Missing Transaction Signature
	ErrTransactionSignature = errors.New("Missing Transaction Signature.")

	// ErrInvalidSignature Invalid Signature
	ErrInvalidSignature = errors.New("Invalid Signature.")

	// ErrInvalidKey Invalid key
	ErrInvalidKey = errors.New("Invalid key.")

	// ErrInvalidReference Invalid reference
	ErrInvalidReference = errors.New("Invalid reference.")

	// ErrNilArgument Invalid reference
	ErrNilArgument = errors.New("Nil argument.")

	// ErrNotImplemented Not implemented
	ErrNotImplemented = errors.New("Not implemented.")

	// ErrKeyStoreAlreadyInitialized Keystore already Initilized
	ErrKeyStoreAlreadyInitialized = errors.New("Keystore already Initilized.")

	// ErrEncrypt Encryption failed
	ErrEncrypt = errors.New("Encryption failed.")

	// ErrDecrypt Decryption failed
	ErrDecrypt = errors.New("Decryption failed.")

	// ErrDifferentChaincodeID ChaincodeIDs are different
	ErrDifferentChaincodeID = errors.New("ChaincodeIDs are different.")

	// ErrDifferrentConfidentialityProtocolVersion different confidentiality protocol versions
	ErrDifferrentConfidentialityProtocolVersion = errors.New("Confidentiality protocol versions are different.")

	// ErrInvalidConfidentialityLevel Invalid confidentiality level
	ErrInvalidConfidentialityLevel = errors.New("Invalid confidentiality level")

	// ErrInvalidConfidentialityProtocol Invalid confidentiality level
	ErrInvalidConfidentialityProtocol = errors.New("Invalid confidentiality protocol")

	// ErrInvalidTransactionType Invalid transaction type
	ErrInvalidTransactionType = errors.New("Invalid transaction type")

	// ErrInvalidProtocolVersion Invalid protocol version
	ErrInvalidProtocolVersion = errors.New("Invalid protocol version")
)

func NoRecvHandlerEnabled() error {
	var text bytes.Buffer
	text.WriteString("No receive handler enabled in configuration file.")
	return errors.New(text.String())
}

func RedisSaveError(message string) error {
	var text bytes.Buffer
	text.WriteString("Fail to save data into redis: ")
	text.WriteString(message)
	return errors.New(text.String())
}

func BlockchainSaveError(message string) error {
	var text bytes.Buffer
	text.WriteString("Fail to save data into blockchain: ")
	text.WriteString(message)
	return errors.New(text.String())
}

func BlockchainResponseError(message string) error {
	var text bytes.Buffer
	text.WriteString("Blockchain response error: ")
	text.WriteString(message)
	return errors.New(text.String())
}

func DataValidationError(message string) error {
	var text bytes.Buffer
	text.WriteString("Data validation fail: ")
	text.WriteString(message)
	return errors.New(text.String())
}

func BCSrvStartError(message string) error {
	var text bytes.Buffer
	text.WriteString("Fail to start Blockchain service: ")
	text.WriteString(message)
	return errors.New(text.String())
}

func MQNotInitialzedError() error {
	var text bytes.Buffer
	text.WriteString("ZeroMQ havn't been initialized!")
	return errors.New(text.String())
}

func MQInvalidModeError(mode string, message string) error {
	var text bytes.Buffer
	text.WriteString("Invalid ZeroMQ mode " + mode + ", ")
	text.WriteString(message)
	return errors.New(text.String())
}
