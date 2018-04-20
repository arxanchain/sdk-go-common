# Crypto tools
Common tools to generate ECC encryption and ED25519 signatures tool adapted to your OS.

## Build

After successfully installed **sdk-go-common**, you should've configured your **GOPAH** environment variable, use the following command to build crypto-util and sign-util executables.

```sh
$ cd $GOPATH/src/github.com/arxanchain/sdk-go-common/crypto/tools
$ make
```
The executables will be built in path `$GOPATH/src/github.com/arxanchain/sdk-go-common/crypto/tools/build/bin`

## Usage

### crypto-util

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

### sign-util

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
