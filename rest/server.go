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

package rest

import (
	"bytes"
	ctls "crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/NYTimes/gziphandler"
	cerr "github.com/arxanchain/sdk-go-common/errors"
	"github.com/arxanchain/sdk-go-common/rest/structs"
	logging "github.com/op/go-logging"
	"github.com/ugorji/go/codec"
)

var (
	logger = logging.MustGetLogger("rest")
	// jsonHandle and jsonHandlePretty are the codec handles to JSON encode
	// structs. The pretty handle will add indents for easier human consumption.
	jsonHandle = &codec.JsonHandle{
		HTMLCharsAsIs: true,
	}
	jsonHandlePretty = &codec.JsonHandle{
		HTMLCharsAsIs: true,
		Indent:        4,
	}
)

// Server interface ...
type Server interface {
	Start() (err error)
	Shutdown()
}

// RequestHandler defines the function type for handling http request
type RequestHandler func(resp http.ResponseWriter, req *http.Request) (interface{}, error)

// HTTPServer is used to wrap an Agent and expose it over an HTTP interface
type HTTPServer struct {
	mux                    *http.ServeMux
	listener               net.Listener
	addr                   string
	HTTPAPIResponseHeaders map[string]string
}

// NewHTTPServer creates new HTTP server
func NewHTTPServer(listen string, tls *TLSConfig) (*HTTPServer, error) {
	logger.Infof("Creating HTTP listener on %s", listen)
	// Start the listener
	lnAddr, err := net.ResolveTCPAddr("tcp", listen)
	if err != nil {
		logger.Errorf("Failed to resolve TCP address %s, %+v", listen, err)
		return nil, err
	}
	ln, err := netListener("tcp", lnAddr.IP.String(), lnAddr.Port)
	if err != nil {
		return nil, fmt.Errorf("failed to start HTTP listener: %v", err)
	}

	// If TLS is enabled, wrap the listener with a TLS listener
	if tls != nil {
		tlsConf := &Config{
			VerifyIncoming:       false,
			VerifyOutgoing:       true,
			VerifyServerHostname: tls.VerifyServerHostname,
			CAFile:               tls.CAFile,
			CertFile:             tls.CertFile,
			KeyFile:              tls.KeyFile,
		}
		tlsConfig, err := tlsConf.IncomingTLSConfig()
		if err != nil {
			return nil, err
		}
		ln = ctls.NewListener(tcpKeepAliveListener{ln.(*net.TCPListener)}, tlsConfig)
	}

	// Create the mux
	mux := http.NewServeMux()

	// Create the server
	srv := &HTTPServer{
		mux:      mux,
		listener: ln,
		addr:     ln.Addr().String(),
		HTTPAPIResponseHeaders: make(map[string]string),
	}

	return srv, nil
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by NewHttpServer so
// dead TCP connections eventually go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(30 * time.Second)
	return tc, nil
}

func netListener(proto, addr string, port int) (net.Listener, error) {

	if 0 > port || port > 65535 {
		return nil, &net.OpError{
			Op:  "listen",
			Net: proto,
			Err: &net.AddrError{Err: "invalid port", Addr: fmt.Sprint(port)},
		}
	}
	return net.Listen(proto, net.JoinHostPort(addr, strconv.Itoa(port)))
}

// Serve is used to start the HTTP server
func (s *HTTPServer) Serve() {
	// Start the server
	go http.Serve(s.listener, gziphandler.GzipHandler(s.mux))
}

// Shutdown is used to shutdown the HTTP server
func (s *HTTPServer) Shutdown() {
	if s != nil {
		logger.Debugf("http: Shutting down http server")
		s.listener.Close()
	}
}

// RegisterHandler register the handler function for http server
func (s *HTTPServer) RegisterHandler(url string, rhandler RequestHandler) {
	s.mux.HandleFunc(url, s.wrap(rhandler))
}

// wrap is used to wrap functions to make them more convenient
func (s *HTTPServer) wrap(handler RequestHandler) func(resp http.ResponseWriter, req *http.Request) {
	f := func(resp http.ResponseWriter, req *http.Request) {
		setHeaders(resp, s.HTTPAPIResponseHeaders)
		// Invoke the handler
		reqURL := req.URL.String()
		start := time.Now()
		defer func() {
			logger.Debugf("http: Request %v (%v)", reqURL, time.Now().Sub(start))
		}()
		obj, err := handler(resp, req)

		response := &structs.Response{}
		// Check for an error
		if err != nil {
			logger.Errorf("http: Request %v, error: %v", reqURL, err)
			code := cerr.ErrCodeType(http.StatusBadRequest)
			if httpError, ok := err.(HTTPCodedError); ok {
				code = httpError.Code()
			}
			// check code is/not http standard status code, if is http code, we write code to resp header
			// else write code 400 to resp header
			if int(code) >= 400 && int(code) < 600 {
				resp.WriteHeader(int(code))
			} else {
				resp.WriteHeader(http.StatusBadRequest)
			}
			//resp.Write([]byte(err.Error()))
			response.ErrMessage = err.Error()
			response.ErrCode = code
		} else {
			if obj != nil {
				if res, ok := obj.(*structs.Response); ok {
					response = res
				} else {
					data, err := json.Marshal(obj)
					if err == nil {
						response.Payload = string(data)
					}
				}
			} else {
				// allow the handler function to handle the response on its own
				return
			}
		}

		prettyPrint := false
		if v, ok := req.URL.Query()["pretty"]; ok {
			if len(v) > 0 && (len(v[0]) == 0 || v[0] != "0") {
				prettyPrint = true
			}
		}

		// Write out the JSON object
		var buf bytes.Buffer
		if prettyPrint {
			enc := codec.NewEncoder(&buf, jsonHandlePretty)
			err = enc.Encode(response)
			if err == nil {
				buf.Write([]byte("\n"))
			}
		} else {
			enc := codec.NewEncoder(&buf, jsonHandle)
			err = enc.Encode(response)
		}
		if err == nil {
			resp.Header().Set("Content-Type", "application/json")
			resp.Header().Set("Access-Control-Allow-Origin", "*")
			resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			resp.Header().Set("Access-Control-Allow-Headers",
				"Action, Module")
			resp.Write(buf.Bytes())
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(err.Error()))
		}
	}
	return f
}

// DecodeBody is used to decode a JSON request body
func DecodeBody(req *http.Request, out interface{}) error {
	dec := json.NewDecoder(req.Body)
	return dec.Decode(&out)
}

// setHeaders is used to set canonical response header fields
func setHeaders(resp http.ResponseWriter, headers map[string]string) {
	for field, value := range headers {
		resp.Header().Set(http.CanonicalHeaderKey(field), value)
	}
}
