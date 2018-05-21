# Crypto tools
Common tools to generate ECC encryption and ED25519 signatures tool adapted to your OS. You can eigther using binary executables or dynamic link library.

## 1. Using dynamic link library

### Build

After successfully installed **sdk-go-common**, you should've configured your **GOPAH** environment variable, use the following command to build crypto-util and sign-util dynamic link library.

```sh
$ cd $GOPATH/src/github.com/arxanchain/sdk-go-common/crypto/tools/library
$ make
```
The dynamic link libraries will be built in path `$GOPATH/src/github.com/arxanchain/sdk-go-common/crypto/tools/library/build`

### Usage

#### crypto-util

Crypto-util is used to implement signing, encryption, decryption and verification process.

```python
>>> import ctypes
>>> func = ctypes.CDLL("./crypto-util.so").encrypt
>>> func.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p]
>>> func.restype = ctypes.c_char_p
>>> result = func("1", "alice", "/your/cert/path", "aGVsbG8gd29ybGQh")
>>> print result
```
Params(in sequence):
  -mode string
      1: sign and encrypt(default); 2: decrypt and verify
  -apikey string
      api key
  -path string
      cert path
  -data string
      base64 encoded data string

Returns: mode 1: base64 encoded cipher; mode 2: plain text

#### sign-util

Sign-util is used to implement ED25519 signing process.

```python
>>> import ctypes
>>> func = ctypes.CDLL("./sign-util.so").sign
>>> func.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p]
>>> func.restype = ctypes.c_char_p
>>> result = func("0lxEFzMQhn68vY2F0f+nOwP7kl5zjahjPcfyMAJVmzn0HNQssIIYh+c2CgCKEHeUvxqCu6W/sJKqKt2DLJnKpw==", "nonce", "123456", "SGVsbG8gd29ybGQh")
>>> print result
```
./sign-util -help
Params(in sequence):
  -key string
        private key
  -data string
        base64 encoded data string to be signed
  -nonce string
        a random string (default "nonce")
  -did string
        creator id

Returns: base64 encoded signed data

## 2. Using binaries

### Build

Use the following command to build crypto-util and sign-util executables.

```sh
$ cd $GOPATH/src/github.com/arxanchain/sdk-go-common/crypto/tools
$ make
```
The executables will be built in path `$GOPATH/src/github.com/arxanchain/sdk-go-common/crypto/tools/build/bin`

### Usage

#### crypto-util

Crypto-util is used to implement signing, encryption, decryption and verification process.

```sh
$ cd build/bin
$ ./crypto-util -help
Usage of ./crypto-util:
  -apikey string
        api key
  -data string
        base64 encoded data string
  -mode string
        1: sign and encrypt(default); 2: decrypt and verify
  -path string
        cert path


$ ./crypto-util -mode 1 \
    -apikey alice \
    -path '~/sdk-go-common/rest/api/client_certs/' \
    -data "aGVsbG8="
```
Returns: mode 1: base64 encoded cipher; mode 2: plain text

#### sign-util

Sign-util is used to implement ED25519 signing process.

```sh
./sign-util -help
Usage of ./sign-util:
  -data string
        base64 encoded data string to be signed
  -did string
        creator id
  -key string
        private key
  -nonce string
        a random string (default "nonce")

$ ./sign-util -key '0lxEFzMQhn68vY2F0f+nOwP7kl5zjahjPcfyMAJVmzn0HNQssIIYh+c2CgCKEHeUvxqCu6W/sJKqKt2DLJnKpw==' \
    -nonce 'nonce' \
    -did '123456' \
    -data 'SGVsbG8gd29ybGQh'
```
Returns: base64 encoded signed data
