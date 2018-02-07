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
	"encoding/json"
	"flag"
	"fmt"

	"github.com/arxanchain/sdk-go-common/crypto/sign/ed25519"
	"github.com/arxanchain/sdk-go-common/structs"
	"github.com/arxanchain/sdk-go-common/utils"
)

func main() {
	//get args
	//.eg: ./sign-util -key '0lxEFzMQhn68vY2F0f+nOwP7kl5zjahjPcfyMAJVmzn0HNQssIIYh+c2CgCKEHeUvxqCu6W/sJKqKt2DLJnKpw==' -nonce 'nonce' -did '123456' -data 'SGVsbG8gd29ybGQh'
	// Returns base64 encoded signed data

	privateKey := flag.String("key", "", "private key")
	data := flag.String("data", "", "base64 encoded data string to be signed")
	nonce := flag.String("nonce", "nonce", "a random string")
	id := flag.String("did", "", "creator id")

	flag.Parse()

	// Validate
	if len(*privateKey) <= 0 {
		fmt.Printf("[ERROR]: key cannot be empty!\n")
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
	private_key, err := utils.DecodeBase64(*privateKey)
	if err != nil {
		fmt.Printf("[ERROR]: DecodeBase64 failed: %s\n", err)
		return
	}

	pri := &ed25519.PrivateKey{
		PrivateKeyData: []byte(private_key),
	}

	ddata, err := utils.DecodeBase64(*data)
	if err != nil {
		fmt.Printf("[ERROR]: DecodeBase64 failed: %s\n", err)
	}
	sh := &structs.SignatureHeader{
		Creator: structs.Identifier(*id),
		Nonce:   []byte(*nonce),
	}

	sd := &structs.SignedData{
		Data:   []byte(ddata),
		Header: sh,
	}
	signData, err := sd.DoSign(pri)
	if err != nil {
		fmt.Printf("[ERROR]: DoSign failed: %s\n", err)
		return
	}
	result, err := json.Marshal(signData)
	if err != nil {
		fmt.Printf("[ERROR]: Unmarshal signed data failed: %s\n", err)
		return
	}
	resultB64 := utils.EncodeBase64(result)
	fmt.Printf("%s\n", resultB64)

	return

}
