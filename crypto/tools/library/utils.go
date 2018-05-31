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

//util.go
package main

import (
	"C"
	"fmt"

	cpt "github.com/arxanchain/sdk-go-common/crypto"
	"github.com/arxanchain/sdk-go-common/crypto/sign/ed25519"
	"github.com/arxanchain/sdk-go-common/structs"
	"github.com/arxanchain/sdk-go-common/utils"
)

//export sign
func sign(cprivateKey, cnonce, cdid, cdata *C.char) *C.char {
	//
	// Returns base64 encoded signed data
	//.eg:
	// signed := sign("0lxEFzMQhn68vY2F0f+nOwP7kl5zjahjPcfyMAJVmzn0HNQssIIYh+c2CgCKEHeUvxqCu6W/sJKqKt2DLJnKpw==", "nonce", "123456", "SGVsbG8gd29ybGQh")
	// if cipher != nil{
	//     fmt.Printf("succeed")
	// }

	privateKey := C.GoString(cprivateKey)
	data := C.GoString(cdata)
	nonce := C.GoString(cnonce)
	id := C.GoString(cdid)

	// Validate
	if len(privateKey) <= 0 {
		fmt.Printf("[ERROR]: key cannot be empty!\n")
		return nil
	}
	if len(data) <= 0 {
		fmt.Printf("[ERROR]: data cannot be empty!\n")
		return nil
	}
	if len(id) <= 0 {
		fmt.Printf("[ERROR]: did cannot be empty!\n")
		return nil
	}
	privateKeyB64, err := utils.DecodeBase64(privateKey)
	if err != nil {
		fmt.Printf("[ERROR]: DecodeBase64 failed: %s\n", err)
		return nil
	}

	ddata, err := utils.DecodeBase64(data)
	if err != nil {
		fmt.Printf("[ERROR]: DecodeBase64 failed: %s\n", err)
		return nil
	}

	sd := &structs.SignedData{
		Data: []byte(ddata),
		Header: &structs.SignatureHeader{
			Creator: structs.Identifier(id),
			Nonce:   []byte(nonce),
		},
	}
	signData, err := sd.DoSign(&ed25519.PrivateKey{
		PrivateKeyData: []byte(privateKeyB64)},
	)
	if err != nil {
		fmt.Printf("[ERROR]: DoSign failed: %s\n", err)
		return nil
	}

	return C.CString(utils.EncodeBase64(signData.Sign))

}

//export encrypt
func encrypt(cmode, capiKey, cpath, cdata *C.char) *C.char {
	//
	// Returns base64 encoded cipher
	// eg.:
	// cipher := encrypt("1", "alice", /your/cert/path", "aGVsbG8gd29ybGQh")
	// if cipher != nil {
	//     fmt.Printf("succeed")
	// }
	mode := C.GoString(cmode)
	apiKey := C.GoString(capiKey)
	path := C.GoString(cpath)
	data := C.GoString(cdata)

	if len(apiKey) <= 0 {
		fmt.Printf("[ERROR]: apikey cannot be empty!\n")
		return nil
	}
	if len(path) <= 0 {
		fmt.Printf("[ERROR]: path cannot be empty!\n")
		return nil
	}
	if len(data) <= 0 {
		fmt.Printf("[ERROR]: data cannot be empty!\n")
		return nil
	}
	// Decode base64
	// Init Crypto Lib
	cpt.SetSecurityLevel(256, "SHA3")
	cpt.SetServerClientMode(cpt.ServerClientMode(cpt.CLIENT_MODE))
	_, err := cpt.NewCertsStore(path)
	if err != nil {
		fmt.Printf("[ERROR]: Init cert store fail: %v\n", err)
		return nil
	}

	switch mode {
	case "1":
		//sign and encrypt
		message, err := utils.DecodeBase64(data)
		if err != nil {
			fmt.Printf("[ERROR]: Decode base64 fail: %v\n", err)
			return nil
		}
		result, err := cpt.SignAndEncrypt(message, apiKey)
		if err != nil {
			fmt.Printf("[ERROR]: SignAndEncrypt fail: %v\n", err)
			return nil
		}
		return C.CString(result)
	case "2":
		//decrypt and verify
		result, err := cpt.DecryptAndVerify([]byte(data), apiKey)
		if err != nil {
			fmt.Printf("[ERROR]: DecryptAndVerify fail: %v\n", err)
			return nil
		}
		return C.CString(string(result))
	default:
		fmt.Printf("[ERROR]: Unsupported operation: %s", mode)
		return nil
	}
}

func main() {}
