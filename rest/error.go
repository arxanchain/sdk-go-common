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
	cerr "github.com/arxanchain/sdk-go-common/errors"
)

const (
	// ErrInvalidMethod is used if the HTTP method is not supported
	ErrInvalidMethod = "Invalid method"

	// ErrInvalidParams is used if the HTTP request params is not valid
	ErrInvalidParams = "Invalid request params"

	// ErrInvalidRequest is used if the HTTP request is invalid for example: invalid parameter ...
	ErrInvalidRequest = "Invalid request"

	ErrPermissionDenied = "Permission deny"

	ErrInternalServerError = "Internal server error"

	ErrMethodNotImplemented = "Method not implementated"

	// ErrCodeDuplicateEntity is error code to indicate has duplicate resource
	ErrCodeDuplicateEntity cerr.ErrCodeType = 1001
)

// HTTPCodedError is used to provide the HTTP error code
type HTTPCodedError interface {
	error
	Code() cerr.ErrCodeType
}

func CodedError(c cerr.ErrCodeType, s string) HTTPCodedError {
	return &codedError{s, c}
}

type codedError struct {
	s    string
	code cerr.ErrCodeType
}

func (e *codedError) Error() string {
	return e.s
}

func (e *codedError) Code() cerr.ErrCodeType {
	return e.code
}
