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

	"github.com/arxanchain/sdk-go-common/crypto"
	"github.com/arxanchain/sdk-go-common/utils"
)

func main() {
	//get args
	//.eg: ./crypto-util -mode 1 -apikey alice -path '~/sdk-go-common/rest/api/client_certs/' -data "aGVsbG8="
	// Returns: mode 1: base64 encoded cipher; mode 2: plain text

	mode := flag.String("mode", "1", "1: sign and encrypt(default), 2: decrypt and verify")
	apiKey := flag.String("apikey", "", "api key.")
	path := flag.String("path", "", "cert path.")
	data := flag.String("data", "", "base64 encoded data string.")

	flag.Parse()
	// Validate
	if len(*apiKey) <= 0 {
		fmt.Printf("[ERROR]: apikey cannot be empty!\n")
		return
	}
	if len(*path) <= 0 {
		fmt.Printf("[ERROR]: path cannot be empty!\n")
		return
	}
	if len(*data) <= 0 {
		fmt.Printf("[ERROR]: data cannot be empty!\n")
		return
	}
	// Decode base64
	// Init Crypto Lib
	crypto.SetSecurityLevel(256, "SHA3")
	crypto.SetServerClientMode(crypto.ServerClientMode(crypto.CLIENT_MODE))
	_, err := crypto.NewCertsStore(*path)
	if err != nil {
		fmt.Printf("[ERROR]: Init cert store fail: %v\n", err)
		return
	}

	switch *mode {
	case "1":
		//sign and encrypt
		message, err := utils.DecodeBase64(*data)
		if err != nil {
			fmt.Printf("[ERROR]: Decode base64 fail: %v\n", err)
			return
		}
		result, err := crypto.SignAndEncrypt(message, *apiKey)
		if err != nil {
			fmt.Printf("[ERROR]: SignAndEncrypt fail: %v\n", err)
			return
		}
		fmt.Printf("%s\n", result)
		return
	case "2":
		//decrypt and verify
		result, err := crypto.DecryptAndVerify([]byte(*data), *apiKey)
		if err != nil {
			fmt.Printf("[ERROR]: DecryptAndVerify fail: %v\n", err)
			return
		}
		fmt.Printf("%s\n", result)
		return
	default:
		fmt.Printf("[ERROR]: Unsupported operation: %s", *mode)
		return
	}
}
