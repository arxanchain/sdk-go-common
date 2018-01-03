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

package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

func TestArraytoString(t *testing.T) {
	data := []string{"aa", "bb"}
	value := ArraytoString(data)
	t.Logf("value: %s", value)
	var expected []string
	json.Unmarshal([]byte(value), &expected)

	AssertSame(t, fmt.Sprintf("%v", data), fmt.Sprintf("%v", expected))
}

func TestIsHexdigest(t *testing.T) {
	actual, err := IsHexdigest("test", "sha512")
	AssertError(t, err, fmt.Sprintf("Unsupported type: %s", "sha512"))

	sha256 := "8b27626debf300bf598e7818fc8650ce1a1fbd920e3903e48961228143806e7c"
	actual, _ = IsHexdigest(sha256, "sha256")
	AssertEquals(t, actual, true)

	name := "高爽"
	actual, _ = IsHexdigest(name, "sha256")
	AssertEquals(t, actual, false)

	hasher := md5.New()
	hasher.Write([]byte("test"))
	str := hex.EncodeToString(hasher.Sum(nil))
	actual, _ = IsHexdigest(str, "md5")
	AssertEquals(t, actual, true)

	actual, _ = IsHexdigest(str, "sha256")
	AssertEquals(t, actual, false)

	h := sha1.New()
	h.Write([]byte("test"))
	str = hex.EncodeToString(h.Sum(nil))
	actual, _ = IsHexdigest(str, "sha1")
	AssertEquals(t, actual, true)

	actual, _ = IsHexdigest("test", "")
	AssertEquals(t, actual, false)
}

func TestStringInSlice(t *testing.T) {
	types := []string{"01", "02", "03"}
	a := "01"
	actual := StringInSlice(a, types)
	AssertEquals(t, actual, true)

	a = "099"
	actual = StringInSlice(a, types)
	AssertEquals(t, actual, false)

}
func TestComputeSha256(t *testing.T) {
	name := "高爽"
	expected := "8b27626debf300bf598e7818fc8650ce1a1fbd920e3903e48961228143806e7c"
	actual := ComputeSha256(name)
	AssertEquals(t, actual, expected)

	name = "王鹏"
	expected = "e034ca73746c2fdef37770e08522dbecc8be86e7f855fd15cbb0aedcd1727f13"
	actual = ComputeSha256(name)
	AssertEquals(t, actual, expected)

	id := "110104198511241221"
	expected = "e961c7d840ae13c85ca6071ad591f4c69f650a1dea4d43e14f9c9a33785951e4"
	actual = ComputeSha256(id)
	AssertEquals(t, actual, expected)

	id = "340223198209050819"
	expected = "df3553bae6333f22a7253877f0fafa5d281baa943e67b51d6b109b4f0b382fd4"
	actual = ComputeSha256(id)
	AssertEquals(t, actual, expected)
}

func TestComputeBitcoinAddress(t *testing.T) {
	pk := []byte("test1")
	addr := "2a3S9kvDfu5rCd9pJkda6bh5LMNj2PdwywwZ9jpGyEwwtKi1y1xjW4AUo4PBSP3ASTHFjLVBLep8jzpmuk51"
	actual := ComputeBitcoinAddress(pk)
	AssertEquals(t, actual, addr)

	pk = []byte("test2")
	addr = "2a3S9kvLc18APNCXMXxyc3W1eoxjjG7YLuTQHimbytm65m1sekQhWcmeuwp9dEcPmNCA8YbGgumUHMurJTxX"
	actual = ComputeBitcoinAddress(pk)
	AssertEquals(t, actual, addr)

	pk = []byte("abcdefghijk")
	addr = "3tRMvDs5QdEsDdHQJzQwcWEiJawDnvet6HD4WBwXpewSQXitKFompp5pZDwaU8ebkLL8xrDX7jgEKBv4GfrZZeFah5j1"
	actual = ComputeBitcoinAddress(pk)
	AssertEquals(t, actual, addr)
}

func TestInt2Str(t *testing.T) {
	pType := 1
	wantLen := 2
	result := Int2Str(pType, wantLen)
	AssertEquals(t, result, "01")

	pType = 100
	wantLen = 2
	result = Int2Str(pType, wantLen)
	AssertEquals(t, result, "100")

	pType = 10
	wantLen = 2
	result = Int2Str(pType, wantLen)
	AssertEquals(t, result, "10")

	pType = 10
	wantLen = 3
	result = Int2Str(pType, wantLen)
	AssertEquals(t, result, "010")
}

func TestCreateUtcToday(t *testing.T) {
	today := CreateUtcToday()
	t.Log(today)
	match, _ := regexp.MatchString("[0-9]{4}-[0-9]{2}-[0-9]{2}", today)
	AssertEquals(t, match, true)
}

func TestGenerateShortID(t *testing.T) {
	id := GenerateShortID()
	t.Log(id)
	AssertNotEquals(t, id, "")
	AssertEquals(t, len(id), 9)
}

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("example key 1234")
	originalText := "encrypt this golang"
	cryptoText := Encrypt(key, originalText)
	decrypted := Decrypt(key, cryptoText)
	AssertEquals(t, decrypted, originalText)
}

func TestEncryptVerifyPass(t *testing.T) {
	pass := "verify pass"
	hash := EncryptPassword(pass)
	AssertNotEquals(t, hash, pass)
	AssertEquals(t, VerifyPassword(pass, hash), true)
}

func TestValidateUUID(t *testing.T) {
	invalid_uuid := "3e513e9d-2765-524b-8d31-2ac7af8bd465"
	ok := ValidateUUID(invalid_uuid)
	AssertEquals(t, ok, false)
	valid_uuid := "3e513e9d-2765-424b-8d31-2ac7af8bd465"
	ok = ValidateUUID(valid_uuid)
	AssertEquals(t, ok, true)
}

func TestValidateTXID(t *testing.T) {
	invalid_txid := "asdasdasdasd"
	ok := ValidateTXID(invalid_txid)
	AssertEquals(t, ok, false)
	valid_txid := "7b9a060fa4a710b383380c427377a710c0c87e33b76b5f30a47d007493c87a95"
	ok = ValidateTXID(valid_txid)
	AssertEquals(t, ok, true)
}

func TestValidateBase64EncodingString(t *testing.T) {
	invalid_string := "asdaed_asdasd"
	ok := ValidateBase64EncodingString(invalid_string)
	AssertEquals(t, ok, false)
	valid_string := "+dwGalE8Cjowt/7hQq64lvDvPaAvIxcOeioxmehr2dw="
	ok = ValidateBase64EncodingString(valid_string)
	AssertEquals(t, ok, true)
}

func TestBase58EncodeDecode(t *testing.T) {
	mssage := "00eb15231dfceb60925886b67d065299925915aeb172c06647"
	encode := EncodeBase58([]byte(mssage))
	decode := DecodeBase58(encode)
	AssertEquals(t, []byte(mssage), decode)
}

func TestRandNewStr(t *testing.T) {
	length := 20
	randString := RandNewStr(length)
	AssertEquals(t, len(randString), 20)
}

func TestCopyFile(t *testing.T) {
	str := "temporary file's content"
	content := []byte(str)
	tmpfile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err1 := tmpfile.Write(content); err1 != nil {
		t.Fatal(err1)
	}
	if err2 := tmpfile.Close(); err2 != nil {
		t.Fatal(err2)
	}

	dstLink := filepath.Join(os.TempDir(), "testdstlink")
	defer os.Remove(dstLink)
	err = CopyFile(tmpfile.Name(), dstLink)
	if err != nil {
		t.Fatal()
	}

	dstContent, err := ioutil.ReadFile(dstLink)
	AssertNil(t, err)
	AssertEquals(t, string(dstContent[:]), str)

	// test the error path: src does not exist
	err = CopyFile("null", dstLink)
	AssertNotNil(t, err)

	// test error path: src is not non-regular source file
	err = CopyFile("./", dstLink)
	AssertNotNil(t, err)

	// test copyFileContents
	dst := filepath.Join(os.TempDir(), "testdst")
	defer os.Remove(dst)
	err = copyFileContents(tmpfile.Name(), dst)
	AssertNil(t, err)
	dstContent, err = ioutil.ReadFile(dst)
	AssertNil(t, err)
	AssertEquals(t, string(dstContent[:]), str)
}
