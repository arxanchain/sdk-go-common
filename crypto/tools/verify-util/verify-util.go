/*
Copyright ArxanFintech Technology Ltd. 2018 All Rights Reserved.

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

package main

import (
	"flag"
	"fmt"

	"github.com/arxanchain/sdk-go-common/crypto/sign/ed25519"
	commdid "github.com/arxanchain/sdk-go-common/structs/did"
	"github.com/arxanchain/sdk-go-common/structs/pki"
	"github.com/arxanchain/sdk-go-common/utils"
)

func main() {
	//get args
	//.eg: ./sign-util  -data yourdata -did did:axn:63946328-58a9-43cf-b370-a6da3a9059b9 -pubkey MwlGskPlAadBU3rCt9V3sP4rzxANxrcL3IowfMi7lZU= -nonce nonce -signed izLVivkQkLsw7rG57Dd7oMFx0e8pb7MIxB9lHj6GflomK53Q/5E4mIUJ0xfpPjjAF8buHUpFEFa40wEOZ5lvCQ==
	// Returns verify error message

	signed := flag.String("signed", "", "base64 signed data string to be verify")
	nonce := flag.String("nonce", "nonce", "a random string")
	id := flag.String("did", "", "creator id")
	data := flag.String("data", "", "original base64 encoded data")
	publicKey := flag.String("pubkey", "", "base64 public key")

	flag.Parse()

	// Validate
	if len(*signed) <= 0 {
		fmt.Printf("[ERROR]: signed cannot be empty!\n")
		return
	}
	if len(*data) <= 0 {
		fmt.Printf("[ERROR]: data cannot be empty!\n")
		return
	}
	if len(*id) <= 0 {
		fmt.Printf("[ERROR]: did cannot be empty!\n")
		return
	}
	if len(*publicKey) <= 0 {
		fmt.Printf("[ERROR]: publicKey cannot be empty!\n")
		return
	}

	var header = &pki.SignatureHeader{
		Creator: commdid.Identifier(*id),
		Nonce:   []byte(*nonce),
	}

	ddata, err := utils.DecodeBase64(*data)
	if err != nil {
		fmt.Printf("[ERROR]: DecodeBase64 failed: %s\n", err)
	}
	signeddata, err := utils.DecodeBase64(*signed)
	if err != nil {
		fmt.Printf("[ERROR]: DecodeBase64 failed: %s\n", err)
	}
	var sd = &pki.SignedData{
		Data:   []byte(ddata),
		Header: header,
		Sign:   []byte(signeddata),
	}

	public_key, err := utils.DecodeBase64(*publicKey)
	if err != nil {
		fmt.Printf("[ERROR]: DecodeBase64 failed: %s\n", err)
		return
	}

	ipk := &ed25519.PublicKey{
		PublicKeyData: []byte(public_key)}

	err = sd.Verify(ipk)

	fmt.Printf("%s\n", err)

	return
}
