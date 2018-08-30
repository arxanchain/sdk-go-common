# Crypto tools

Common tools to generate ECC encryption and ED25519 signatures tool adapted to your OS. You can eigther using binary executables or dynamic link library.

## 1. Using dynamic link library

### Build

After successfully installed **sdk-go-common**, you should've configured your **GOPAH** environment variable, use the following command to build dynamic link library.

```sh
$ cd $GOPATH/src/github.com/arxanchain/sdk-go-common/crypto/tools/library
$ make
```

The dynamic link libraries will be built in path `$GOPATH/src/github.com/arxanchain/sdk-go-common/crypto/tools/library/build`

### Usage

#### ECC Signing and Encryption

Function `encrypt` is used to implement ecc signing and encryption, or ecc decryption and verification process.

- Ecc Signing and Encryption

	Use mode 1 to invoke `encrypt` function.

	```python
	>>> import ctypes
	>>> func = ctypes.CDLL("./utils.so").encrypt
	>>> func.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p]
	>>> func.restype = ctypes.c_char_p
	>>> result = func("1", "your-api-key", "/your/cert/path", "your-base64-encoded-data-to-be-signed-and-encrypted")
	>>> print result
	```

	Params(in sequence):

		- mode string
			- 1: signing and encryption (default)
			- 2: decryption and signature verification
		- apikey string
			- your real API Key
		- path string
			- your ecc cert path
		- data string
			- at mode 1: your base64 encoded data to be signed and encrypted
			- at mode 2: your base64 encoded data to be decrypted and verified

	Returns:

		- at mode 1: base64 encoded cipher text
		- at mode 2: plain text

- Ecc Decryption and Verification

	Use mode 2 to invoke `encrypt` function.

	```python
	>>> import ctypes
	>>> func = ctypes.CDLL("./utils.so").encrypt
	>>> func.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p]
	>>> func.restype = ctypes.c_char_p
	>>> result = func("2", "your-api-key", "/your/cert/path", "your-base64-encoded-data-to-be-decrypted-and-verified")
	>>> print result
	```

	Params(in sequence):

		- mode string
			- 1: signing and encryption (default)
			- 2: decryption and signature verification
		- apikey string
			- your real API Key
		- path string
			- your ecc cert path
		- data string
			- at mode 1: your base64 encoded data to be signed and encrypted
			- at mode 2: your base64 encoded data to be decrypted and verified

	Returns:

		- at mode 1: base64 encoded cipher text
		- at mode 2: plain text

#### ED25519 Signing

Function `sign` is used to implement ED25519 signing process.

```python
>>> import ctypes
>>> func = ctypes.CDLL("./utils.so").sign
>>> func.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p, ctypes.c_char_p]
>>> func.restype = ctypes.c_char_p
>>> result = func("your-base64-encoded-ed25519-private-key", "nonce", "your-creator-did-for-signing", "your-base64-encoded-data-to-be-signed")
>>> print result
```

Params(in sequence):

	- privateKey string
		- your base64 encoded ed25519 private key
	- nonce string
		- a random string (default "nonce")
	- did string
		- your creator did for signing
	- data string
		- base64 encoded data string to be signed

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

#### ECC crypto util

`crypto-util` is used to implement ecc signing and encryption, or ecc decryption and verification process.

- Usage

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

	```

	Returns:

		- mode 1: base64 encoded cipher text
		- mode 2: plain text


- Ecc Signing and Encryption Example

	Use mode 1 to invoke crypto util.

	```
	$ ./crypto-util \
		-mode 1 \
		-apikey your-api-key \
		-path /your/ecc/cert/path \
		-data your-base64-encoded-data-to-be-signed-and-encrypted
	```

- ECC Decryption and Verification Example

	Use mode 2 to invoke crypto util.

	```
	$ ./crypto-util \
		-mode 2 \
		-apikey your-api-key \
		-path /your/ecc/cert/path \
		-data your-base64-encoded-data-to-be-decrypted-and-verified
	```

#### ED25519 signing util

`sign-util` is used to implement ED25519 signing process.

- Usage

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
	```

	Returns: base64 encoded signed data

- ED25519 Signing Example

	```
	$ ./sign-util \
		-key your-base64-encoded-ed25519-private-key \
		-nonce 'nonce' \
		-did your-creator-did-for-signing \
		-data your-base64-encoded-data-to-be-signed
	```
