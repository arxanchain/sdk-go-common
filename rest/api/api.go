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
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/arxanchain/sdk-go-common/crypto"
	"github.com/arxanchain/sdk-go-common/structs"
	"github.com/arxanchain/sdk-go-common/structs/pki"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
	rootcerts "github.com/hashicorp/go-rootcerts"
)

// HttpBasicAuth is used to authenticate http client with HTTP Basic Authentication
type HttpBasicAuth struct {
	// Username to use for HTTP Basic Authentication
	Username string

	// Password to use for HTTP Basic Authentication
	Password string
}

type ICacheClient interface {
	HGet(key, field string) ([]byte, error)
	HSet(key, field string, data []byte) error
	HDel(key string, fields ...string) error
}

// Config is used to configure the creation of a client
type Config struct {
	// Address is the address of the Rest server
	Address string

	// Scheme is the URI scheme for the Rest server
	Scheme string

	// RouteTag is the route tag used by fabio to discover service
	RouteTag string

	// Transport is the Transport to use for the http client.
	Transport *http.Transport

	// HttpClient is the client to use. Default will be
	// used if not provided.
	HttpClient *http.Client

	// HttpAuth is the auth info to use for http access.
	HttpAuth *HttpBasicAuth

	// Token is used to provide a per-request ACL token
	// which overrides the agent's default token.
	Token string

	// Username is the login username used to get token from Fred service
	Username string

	// Secret is the login password encrypted by utils/AES
	// used to get token from Fred service
	Secret string

	// SecretKey is the secret key used to encrypt/decrypt Secret field
	SecretKey string

	// ApiKey is the access key for ACL access api
	ApiKey string

	// EnterpriseSignParam is enterprise sign parameters
	EnterpriseSignParam *pki.SignatureParam

	// CallbackUrl is used to receive asynchronous event notification
	// which will notify if the request succeeded or failed
	CallbackUrl string

	// TLS config
	TLSConfig TLSConfig

	// CryptoCfg is used to crypto transation between wasabi and client
	CryptoCfg *CryptoConfig

	// CacheClient is used to cache auth token
	CacheClient ICacheClient

	// TrusteeKeyPairEnable is used to set the flag of trust key pair,
	// if you want to trust, set the flag is true.
	TrusteeKeyPairEnable bool
}

// TLSConfig is used to generate a TLSClientConfig that's useful for talking to
// Rest using TLS.
type TLSConfig struct {
	// Address is the optional address of the Rest server. The port, if any
	// will be removed from here and this will be set to the ServerName of the
	// resulting config.
	Address string

	// CAFile is the optional path to the CA certificate used for Rest
	// communication, defaults to the system bundle if not specified.
	CAFile string

	// CAPath is the optional path to a directory of CA certificates to use for
	// Rest communication, defaults to the system bundle if not specified.
	CAPath string

	// CertFile is the optional path to the certificate for Rest
	// communication. If this is set then you need to also set KeyFile.
	CertFile string

	// KeyFile is the optional path to the private key for Rest communication.
	// If this is set then you need to also set CertFile.
	KeyFile string

	// InsecureSkipVerify if set to true will disable TLS host verification.
	InsecureSkipVerify bool
}

// CryptoConfig is used to crypto transation between wasabi and client
type CryptoConfig struct {
	// Enable flag of transation mode
	// true is crypto transation, false is origin transation
	Enable bool

	// CertsStorePath path of certificate file
	CertsStorePath string

	// EncryptType crypto mothed
	// 0: ecc, 1: rsa
	EncryptType crypto.EncryptType

	// SecurityLevel sets the security level
	SecurityLevel int

	// HashAlgorithm hash method
	// default is  "sha256"
	HashAlgorithm string
}

// DefaultConfig returns a default configuration for the client. By default this
// will pool and reuse idle connections to Rest. If you have a long-lived
// client object, this is the desired behavior and should make the most efficient
// use of the connections to Rest. If you don't reuse a client object , which
// is not recommended, then you may notice idle connections building up over
// time. To avoid this, use the DefaultNonPooledConfig() instead.
func DefaultConfig() *Config {
	return defaultConfig(cleanhttp.DefaultPooledTransport)
}

// DefaultNonPooledConfig returns a default configuration for the client which
// does not pool connections. This isn't a recommended configuration because it
// will reconnect to Rest on every request, but this is useful to avoid the
// accumulation of idle connections if you make many client objects during the
// lifetime of your application.
func DefaultNonPooledConfig() *Config {
	return defaultConfig(cleanhttp.DefaultTransport)
}

// defaultConfig returns the default configuration for the client, using the
// given function to make the transport.
func defaultConfig(transportFn func() *http.Transport) *Config {
	config := &Config{
		Address:   "127.0.0.1:8500",
		Scheme:    "http",
		Transport: transportFn(),
		CryptoCfg: &CryptoConfig{Enable: false},
	}

	return config
}

// TLSConfig is used to generate a TLSClientConfig that's useful for talking to
// Rest using TLS.
func SetupTLSConfig(tlsConfig *TLSConfig) (*tls.Config, error) {
	tlsClientConfig := &tls.Config{
		InsecureSkipVerify: tlsConfig.InsecureSkipVerify,
	}

	if tlsConfig.Address != "" {
		server := tlsConfig.Address
		hasPort := strings.LastIndex(server, ":") > strings.LastIndex(server, "]")
		if hasPort {
			var err error
			server, _, err = net.SplitHostPort(server)
			if err != nil {
				return nil, err
			}
		}
		tlsClientConfig.ServerName = server
	}

	if tlsConfig.CertFile != "" && tlsConfig.KeyFile != "" {
		tlsCert, err := tls.LoadX509KeyPair(tlsConfig.CertFile, tlsConfig.KeyFile)
		if err != nil {
			return nil, err
		}
		tlsClientConfig.Certificates = []tls.Certificate{tlsCert}
	}

	rootConfig := &rootcerts.Config{
		CAFile: tlsConfig.CAFile,
		CAPath: tlsConfig.CAPath,
	}
	if err := rootcerts.ConfigureTLS(tlsClientConfig, rootConfig); err != nil {
		return nil, err
	}

	return tlsClientConfig, nil
}

// Client provides a client to the Rest API
type Client struct {
	config *Config
}

// NewClient returns a new client
func NewClient(config *Config) (*Client, error) {
	// bootstrap the config
	defConfig := DefaultConfig()

	if len(config.Address) == 0 {
		config.Address = defConfig.Address
	}

	if len(config.Scheme) == 0 {
		config.Scheme = defConfig.Scheme
	}

	if config.Transport == nil {
		config.Transport = defConfig.Transport
	}

	if config.TLSConfig.Address == "" {
		config.TLSConfig.Address = defConfig.TLSConfig.Address
	}

	if config.TLSConfig.CAFile == "" {
		config.TLSConfig.CAFile = defConfig.TLSConfig.CAFile
	}

	if config.TLSConfig.CAPath == "" {
		config.TLSConfig.CAPath = defConfig.TLSConfig.CAPath
	}

	if config.TLSConfig.CertFile == "" {
		config.TLSConfig.CertFile = defConfig.TLSConfig.CertFile
	}

	if config.TLSConfig.KeyFile == "" {
		config.TLSConfig.KeyFile = defConfig.TLSConfig.KeyFile
	}

	if !config.TLSConfig.InsecureSkipVerify {
		config.TLSConfig.InsecureSkipVerify = defConfig.TLSConfig.InsecureSkipVerify
	}

	if config.CryptoCfg == nil {
		config.CryptoCfg = defConfig.CryptoCfg
	}

	if config.CryptoCfg.Enable {
		// when config.CryptoCfg.Enable is true, the certificate path must be set
		if config.CryptoCfg.CertsStorePath == "" {
			return nil, fmt.Errorf("ERROR: certificate stroe path must be set")
		}

		// initialize struct
		if config.CryptoCfg.SecurityLevel == 0 {
			config.CryptoCfg.SecurityLevel = 256
		}
		if config.CryptoCfg.HashAlgorithm == "" {
			config.CryptoCfg.HashAlgorithm = "SHA3"
		}
		if config.CryptoCfg.EncryptType == 0 {
			config.CryptoCfg.EncryptType = crypto.ECC_TYPE
		}

		// Init Crypto Lib
		crypto.SetSecurityLevel(config.CryptoCfg.SecurityLevel, config.CryptoCfg.HashAlgorithm)
		// Infact, at SDK client the ClientMode/SignFlag/EncryptFlag all have default value, not need change anymore
		crypto.SetServerClientMode(crypto.ServerClientMode(crypto.CLIENT_MODE))
		crypto.SetEncryptFlag(true)
		crypto.SetSignFlag(true)
		crypto.SetEncryptType(crypto.EncryptType(config.CryptoCfg.EncryptType))
		_, err := crypto.NewCertsStore(config.CryptoCfg.CertsStorePath)
		if err != nil {
			return nil, fmt.Errorf("ERROR: failed to initialize the certs store: %s", err)
		}
	}

	if config.HttpClient == nil {
		var err error
		config.HttpClient, err = NewHttpClient(config.Transport, config.TLSConfig)
		if err != nil {
			return nil, err
		}
	}

	parts := strings.SplitN(config.Address, "://", 2)
	if len(parts) == 2 {
		switch parts[0] {
		case "http":
			config.Scheme = "http"
		case "https":
			config.Scheme = "https"
		case "unix":
			trans := cleanhttp.DefaultTransport()
			trans.DialContext = func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", parts[1])
			}
			config.HttpClient = &http.Client{
				Transport: trans,
			}
		default:
			return nil, fmt.Errorf("Unknown protocol scheme: %s", parts[0])
		}
		config.Address = parts[1]
	}

	if config.Token == "" {
		config.Token = defConfig.Token
	}

	return &Client{config: config}, nil
}

// NewClient returns a new client
func (c *Client) GetEnterpriseSignParam() (*pki.SignatureParam, error) {
	if c.config.EnterpriseSignParam == nil {
		return nil, fmt.Errorf("enterprise signature params is nil")
	}
	return c.config.EnterpriseSignParam, nil
}

// NewHttpClient returns an http client configured with the given Transport and TLS
// config.
func NewHttpClient(transport *http.Transport, tlsConf TLSConfig) (*http.Client, error) {
	client := &http.Client{
		Transport: transport,
	}

	if transport.TLSClientConfig == nil {
		tlsClientConfig, err := SetupTLSConfig(&tlsConf)

		if err != nil {
			return nil, err
		}

		transport.TLSClientConfig = tlsClientConfig
	}

	return client, nil
}

// Request is used to help build up a request
type Request struct {
	config *Config
	method string
	url    *url.URL
	params url.Values
	body   io.Reader
	header http.Header
	obj    interface{}
	ctx    context.Context
}

// SetBody is used to set Request body.
//
// The input 'obj' arg can be '[]byte' type binary data,
// also it can be structure object. When it is the structure
// object, it will be converted to binary data as JSON using
// json.Marshal.
//
// Once the crypto mode enabled, the binary data will
// be signed and encrypted using crypto libary, then the
// final result set to the request body.
//
func (r *Request) SetBody(obj interface{}) error {
	if r.body != nil {
		return nil
	}
	if obj == nil {
		log.Println("Body object is nil")
		return fmt.Errorf("body object is nil")
	}

	// Check if we should encode the body
	var b io.Reader
	var objData []byte
	var err error
	var ok bool
	if r.config.CryptoCfg.Enable {
		log.Println("Encryption transmission mode enabled")

		apiKey := r.header.Get(structs.APIKeyHeader)
		if apiKey == "" {
			log.Println("API-Key header must be set when enable crypto mode")
			return fmt.Errorf("API-Key header must be set when enable crypto mode")
		}

		objData, ok = obj.([]byte)
		if !ok {
			objData, err = json.Marshal(obj)
			if err != nil {
				log.Printf("json.Marshal input object fail: %v", err)
				return err
			}
		}

		var cipherObjData string
		cipherObjData, err = crypto.SignAndEncrypt(objData, apiKey)
		if err != nil {
			log.Printf("SignAndEncrypt body data fail: %v", err)
			return err
		}

		b = bytes.NewBufferString(cipherObjData)
	} else {
		objData, ok = obj.([]byte)
		if ok && objData != nil {
			b = bytes.NewBuffer(objData)
		} else {
			b, err = EncodeBody(obj)
			if err != nil {
				log.Printf("EncodeBody fail: %v", err)
				return err
			}
		}
	}

	r.body = b

	return nil
}

// SeHeaders is used to add multiple header KV pairs
func (r *Request) SetHeaders(headers http.Header) {
	for k, list := range headers {
		for _, v := range list {
			r.SetHeader(k, v)
		}
	}
}

// SetHeader is used to set one header KV pair
func (r *Request) SetHeader(k, v string) {
	if strings.ToLower(k) == "accept-encoding" {
		return
	}
	r.header.Set(k, v)
}

// GetHeader is used to get one header value
func (r *Request) GetHeader(k string) string {
	return r.header.Get(k)
}

// SetParams is used to set multiple query params
func (r *Request) SetParams(params url.Values) {
	for k, list := range params {
		for _, v := range list {
			r.params.Set(k, v)
		}
	}
}

// SetParam is used to set one query param
func (r *Request) SetParam(k, v string) {
	r.params.Set(k, v)
}

// GetParam is used to get one query param
func (r *Request) GetParam(k string) string {
	return r.params.Get(k)
}

// ToHTTP is used to convert the Request object to an standard HTTP request object.
//
func (r *Request) ToHTTP() (*http.Request, error) {
	// Encode the query parameters
	r.url.RawQuery = r.params.Encode()

	// Create the HTTP request
	req, err := http.NewRequest(r.method, r.url.RequestURI(), r.body)
	if err != nil {
		log.Printf("New http request fail: %v", err)
		return nil, err
	}

	req.URL.Host = r.url.Host
	req.URL.Scheme = r.url.Scheme
	req.Host = r.url.Host
	if r.config.RouteTag != "" {
		req.Host = r.config.RouteTag
	}
	req.Header = r.header

	// Setup auth
	if r.config.HttpAuth != nil {
		req.SetBasicAuth(r.config.HttpAuth.Username, r.config.HttpAuth.Password)
	}
	if r.ctx != nil {
		return req.WithContext(r.ctx), nil
	} else {
		return req, nil
	}
}

// NewRequest is used to create a new request
func (c *Client) NewRequest(method, path string) *Request {
	r := &Request{
		config: c.config,
		method: method,
		url: &url.URL{
			Scheme: c.config.Scheme,
			Host:   c.config.Address,
			Path:   path,
		},
		params: make(map[string][]string),
		header: make(http.Header),
	}
	if c.config.Token != "" {
		r.header.Set(structs.XAuthTokenHeader, r.config.Token)
	}
	if c.config.RouteTag != "" {
		r.header.Set(structs.FabioRouteTagHeader, r.config.RouteTag)
		r.header.Set(structs.RouteTagHeader, r.config.RouteTag)
	}
	if c.config.ApiKey != "" {
		r.header.Set(structs.APIKeyHeader, r.config.ApiKey)
	}
	if c.config.CallbackUrl != "" {
		r.header.Set(structs.CallbackUrlHeader, r.config.CallbackUrl)
	}
	// Prevent data being compressed by gzip
	r.header.Set("Accept-Encoding", "*")
	return r
}

// DoRequest does an HTTP request.
//
// Once the crypto mode enabled, the response result will
// be decrypted and verified signature using crypto libary,
// then the final result will be return to end client.
//
func (c *Client) DoRequest(r *Request) (time.Duration, *http.Response, error) {
	req, err := r.ToHTTP()
	if err != nil {
		return 0, nil, err
	}
	start := time.Now()

	resp, err := c.config.HttpClient.Do(req)
	diff := time.Since(start)

	if !r.config.CryptoCfg.Enable {
		return diff, resp, err
	}

	// decrypt and verify resp.Body
	if err != nil {
		return diff, resp, err
	}
	if resp == nil {
		return diff, nil, fmt.Errorf("DoRequest http.Response is nil")
	}
	apiKey := r.header.Get(structs.APIKeyHeader)
	if apiKey == "" {
		log.Println("API-Key header must be set when enable crypto mode")
		return diff, resp, fmt.Errorf("API-Key header must be set when enable crypto mode")
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return diff, resp, err
	}

	result, err := crypto.DecryptAndVerify(buf, apiKey)
	if err != nil {
		log.Printf("DecryptAndVerify fail: %v", err)
		return diff, resp, err
	}

	// close the body, then set a new one
	resp.Body.Close()
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(result))
	resp.ContentLength = int64(len(result))

	return diff, resp, err
}

// DecodeBody is used to JSON decode a body
func DecodeBody(resp *http.Response, out interface{}) error {
	dec := json.NewDecoder(resp.Body)
	return dec.Decode(out)
}

// EncodeBody is used to encode a request body
func EncodeBody(obj interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(obj); err != nil {
		return nil, err
	}
	return buf, nil
}

// RequireOK is used to wrap DoRequest
func RequireOK(d time.Duration, resp *http.Response, e error) (time.Duration, *http.Response, error) {
	if e != nil {
		if resp != nil {
			resp.Body.Close()
		}
		return d, nil, e
	}
	return d, resp, nil
}
