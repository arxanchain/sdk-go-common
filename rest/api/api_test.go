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

package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/arxanchain/sdk-go-common/crypto"
	"github.com/arxanchain/sdk-go-common/structs"
	"gopkg.in/h2non/gock.v1"
)

func TestSetHeaders(t *testing.T) {
	var r Request
	r.header = make(http.Header)

	head := http.Header{}
	head.Add("X-Auth-Token", "2k9AlO2IynLpsGTRHHVI3TTm/HCkjprGzwlUxbvXayTrj/Lyjwlsty3XzlKldiiGGxJjiSG+54f2CbYCvcLifA==")

	if head.Get("X-Auth-Token") != "2k9AlO2IynLpsGTRHHVI3TTm/HCkjprGzwlUxbvXayTrj/Lyjwlsty3XzlKldiiGGxJjiSG+54f2CbYCvcLifA==" {
		t.Errorf("get headers is failed")
	}

	r.SetHeaders(head)

	token := r.GetHeader("X-Auth-Token")
	if token != "2k9AlO2IynLpsGTRHHVI3TTm/HCkjprGzwlUxbvXayTrj/Lyjwlsty3XzlKldiiGGxJjiSG+54f2CbYCvcLifA==" {
		t.Errorf("get headers is failed")
	}

	if r.header["X-Auth-Token"][0] != "2k9AlO2IynLpsGTRHHVI3TTm/HCkjprGzwlUxbvXayTrj/Lyjwlsty3XzlKldiiGGxJjiSG+54f2CbYCvcLifA==" {
		t.Errorf("get headers is failed")
	}
}

func TestSetHeader(t *testing.T) {
	var r Request
	r.header = make(http.Header)

	r.SetHeader("X-Auth-Token", "2k9AlO2IynLpsGTRHHVI3TTm/HCkjprGzwlUxbvXayTrj/Lyjwlsty3XzlKldiiGGxJjiSG+54f2CbYCvcLifA==")

	token := r.GetHeader("X-Auth-Token")
	if token != "2k9AlO2IynLpsGTRHHVI3TTm/HCkjprGzwlUxbvXayTrj/Lyjwlsty3XzlKldiiGGxJjiSG+54f2CbYCvcLifA==" {
		t.Errorf("get headers is failed")
	}
}

func TestSetParam(t *testing.T) {
	var r Request

	r.params = make(map[string][]string)
	r.SetParam("access", "m1HMDTepp1510733539")

	if r.GetParam("access") != "m1HMDTepp1510733539" {
		t.Errorf("set param is failed")
	}
}

func TestSetParams(t *testing.T) {
	var r Request
	v := url.Values{}

	s := []string{"m1HMDTepp1510733539"}
	v["access"] = s

	r.params = make(map[string][]string)
	r.SetParams(v)

	if r.GetParam("access") != "m1HMDTepp1510733539" {
		t.Errorf("get params is failed")
	}

	if r.params["access"][0] != "m1HMDTepp1510733539" {
		t.Errorf("get params is failed")
	}
}

func TestSetBodyBytes(t *testing.T) {
	// Initialize clien
	c, err := NewClient(&Config{
		ApiKey: "alice",
	})
	if err != nil {
		t.Errorf("NewClient error: %s", err)
	}

	// Set body
	body := []byte("test body #@!234123ssssssssss test body")
	r := c.NewRequest("", "")
	err = r.SetBody(body)
	if err != nil {
		t.Errorf("SetBody error: %s", err)
	}
	req, err := r.ToHTTP()
	if err != nil {
		t.Errorf("ToHttp error: %s", err)
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Get requst.Body error: %s", err)
	}

	if !bytes.Equal(body, result) {
		t.Errorf("Verify error: %s", string(result))
	}
}

func TestSetBodyJson(t *testing.T) {
	// Initialize clien
	c, err := NewClient(&Config{
		ApiKey: "alice",
	})
	if err != nil {
		t.Errorf("NewClient error: %s", err)
	}

	// Set body
	body := "test body #@!234123ssssssssss test body"
	r := c.NewRequest("", "")
	err = r.SetBody(body)
	if err != nil {
		t.Errorf("SetBody error: %s", err)
	}
	req, err := r.ToHTTP()
	if err != nil {
		t.Errorf("ToHttp error: %s", err)
	}

	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Get requst.Body error: %s", err)
	}

	var bodyRes string
	err = json.Unmarshal(result, &bodyRes)
	if err != nil {
		t.Errorf("UnMarshal error: %s", err)
	}

	if body != bodyRes {
		t.Errorf("Verify error: %s", string(result))
	}
}

func TestSetBodyBytesCrypoMode(t *testing.T) {
	// Initialize clien
	c, err := NewClient(&Config{
		ApiKey:    "alice",
		CryptoCfg: &CryptoConfig{Enable: true, CertsStorePath: "./client_certs/"}},
	)
	if err != nil {
		t.Errorf("NewClient error: %s", err)
	}

	// Set body
	body := []byte("test body #@!234123ssssssssss test body")
	r := c.NewRequest("", "")
	err = r.SetBody(body)
	if err != nil {
		t.Errorf("SetBody error: %s", err)
	}
	req, err := r.ToHTTP()
	if err != nil {
		t.Errorf("ToHttp error: %s", err)
	}

	// Initialize server
	crypto.SetServerClientMode(crypto.ServerClientMode(crypto.SERVER_MODE))
	_, err = crypto.NewCertsStore("./server_certs/")
	if err != nil {
		t.Errorf("Initialize server certs store: %s", err)
	}

	chiper, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Get requst.Body error: %s", err)
	}

	// Verify body
	result, err := crypto.DecryptAndVerify(chiper, "alice")
	if err != nil {
		t.Errorf("DecryptoAndVerify error: %s", err)
	}

	if !bytes.Equal(body, result) {
		t.Errorf("Verify error: %s", string(result))
	}
}

func TestSetBodyJsonCrypoMode(t *testing.T) {
	// Initialize clien
	c, err := NewClient(&Config{
		ApiKey:    "alice",
		CryptoCfg: &CryptoConfig{Enable: true, CertsStorePath: "./client_certs/"}},
	)
	if err != nil {
		t.Errorf("NewClient error: %s", err)
	}

	// Set body
	body := "test body #@!234123ssssssssss test body"
	r := c.NewRequest("", "")
	err = r.SetBody(body)
	if err != nil {
		t.Errorf("SetBody error: %s", err)
	}
	req, err := r.ToHTTP()
	if err != nil {
		t.Errorf("ToHttp error: %s", err)
	}

	// Initialize server
	crypto.SetServerClientMode(crypto.ServerClientMode(crypto.SERVER_MODE))
	_, err = crypto.NewCertsStore("./server_certs/")
	if err != nil {
		t.Errorf("Initialize server certs store: %s", err)
	}

	chiper, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Get requst.Body error: %s", err)
	}

	// Verify body
	result, err := crypto.DecryptAndVerify(chiper, "alice")
	if err != nil {
		t.Errorf("DecryptoAndVerify error: %s", err)
	}
	var bodyRes string
	err = json.Unmarshal(result, &bodyRes)
	if err != nil {
		t.Errorf("UnMarshal error: %s", err)
	}

	if body != bodyRes {
		t.Errorf("Verify error: %s", string(result))
	}
}

func TestDoRequest(t *testing.T) {
	// create client
	c, err := NewClient(&Config{
		ApiKey:    "alice",
		CryptoCfg: &CryptoConfig{Enable: true, CertsStorePath: "./client_certs/"}},
	)
	if err != nil {
		t.Errorf("NewClient error: %s", err)
	}

	// set body
	body := "test body #@!234123ssssssssss test body"
	r := c.NewRequest("", "")
	err = r.SetBody(body)
	if err != nil {
		t.Errorf("SetBody error: %s", err)
	}

	// change to server mode, signatrue and encrypt data
	crypto.SetServerClientMode(crypto.ServerClientMode(crypto.SERVER_MODE))
	_, err = crypto.NewCertsStore("./server_certs/")
	if err != nil {
		t.Errorf("Initialize server certs store: %s", err)
	}
	result, err := crypto.SignAndEncrypt([]byte(body), "alice")
	if err != nil {
		t.Errorf("SignAndEncrypt error: %s", err)
	}

	// gock
	gock.InterceptClient(c.config.HttpClient)
	defer gock.Off()
	gock.New("http://127.0.0.1:8500").
		Reply(200).
		JSON(result)

	// change to client mode for verify and decrypt
	crypto.SetServerClientMode(crypto.ServerClientMode(crypto.CLIENT_MODE))
	_, err = crypto.NewCertsStore("./client_certs/")
	if err != nil {
		t.Errorf("Initialize client certs store: %s", err)
	}

	_, resp, err := c.DoRequest(r)
	if err != nil {
		t.Errorf("DoRequest error: %s", err)
	}

	// verify
	buf, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Errorf("ReadAll resp.Body error: %s", err)
	}

	if body != string(buf) {
		t.Errorf("data error, origin data: %s receive data:%s", body, string(buf))
	}
}

func TestDoRequestError(t *testing.T) {
	// create client
	c, err := NewClient(&Config{
		ApiKey:    "alice",
		CryptoCfg: &CryptoConfig{Enable: true, CertsStorePath: "./client_certs/"}},
	)
	if err != nil {
		t.Errorf("NewClient error: %s", err)
	}

	// set body
	body := "test body #@!234123ssssssssss test body"
	r := c.NewRequest("", "")
	err = r.SetBody(body)
	if err != nil {
		t.Errorf("SetBody error: %s", err)
	}

	// gock
	gock.InterceptClient(c.config.HttpClient)
	defer gock.Off()
	gock.New("http://127.0.0.1:8500").
		Reply(200).
		JSON(body)

	r.header.Set(structs.APIKeyHeader, "")
	_, _, err = c.DoRequest(r)
	if err == nil {
		t.Errorf("DoRequest error: %s", err)
	}
}
