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

package structs

import (
	"time"

	cerr "github.com/arxanchain/sdk-go-common/errors"
)

// MessageType defines message type
type MessageType uint8

const (
	// IgnoreUnknownTypeFlag is set along with a MessageType
	// to indicate that the message type can be safely ignored
	// if it is not recognized. This is for future proofing, so
	// that new commands can be added in a way that won't cause
	// old servers to crash when the FSM attempts to process them.
	IgnoreUnknownTypeFlag MessageType = 128

	// APIMajorVersion is returned as part of the Status.Version request.
	// It should be incremented anytime the APIs are changed in a way
	// that would break clients for sane client versioning.
	APIMajorVersion = 1

	// APIMinorVersion is returned as part of the Status.Version request.
	// It should be incremented anytime the APIs are changed to allow
	// for sane client versioning. Minor changes should be compatible
	// within the major version.
	APIMinorVersion = 1

	ProtocolVersionFieldName = "protocol"
	APIMajorVersionFieldName = "api.major"
	APIMinorVersionFieldName = "api.minor"

	RESTGET    = "GET"
	RESTPOST   = "POST"
	RESTPUT    = "PUT"
	RESTPATCH  = "PATCH"
	RESTDELETE = "DELETE"
)

// RPCInfo is used to describe common information about query
type RPCInfo interface {
	RequestRegion() string
	IsRead() bool
	AllowStaleRead() bool
}

// QueryOptions is used to specify various flags for read queries
type QueryOptions struct {
	// The target region for this query
	Region string

	// If set, wait until query exceeds given index. Must be provided
	// with MaxQueryTime.
	MinQueryIndex uint64

	// Provided with MinQueryIndex to wait for change.
	MaxQueryTime time.Duration

	// If set, any follower can service the request. Results
	// may be arbitrarily stale.
	AllowStale bool

	// If set, used as prefix for resource list searches
	Prefix string
}

func (q QueryOptions) RequestRegion() string {
	return q.Region
}

// IsRead, QueryOptions only applies to reads, so always true
func (q QueryOptions) IsRead() bool {
	return true
}

func (q QueryOptions) AllowStaleRead() bool {
	return q.AllowStale
}

type WriteRequest struct {
	// The target region for this write
	Region string
}

func (w WriteRequest) RequestRegion() string {
	// The target region for this request
	return w.Region
}

// WriteRequest only applies to writes, always false
func (w WriteRequest) IsRead() bool {
	return false
}

func (w WriteRequest) AllowStaleRead() bool {
	return false
}

// QueryMeta allows a query response to include potentially
// useful metadata about a query
type QueryMeta struct {
	// This is the index associated with the read
	Index uint64

	// If AllowStale is used, this is time elapsed since
	// last contact between the follower and leader. This
	// can be used to gauge staleness.
	LastContact time.Duration

	// Used to indicate if there is a known leader node
	KnownLeader bool
}

// WriteMeta allows a write response to include potentially
// useful metadata about the write
type WriteMeta struct {
	// This is the index associated with the write
	Index uint64
}

// GenericRequest is used to request where no
// specific information is needed.
type GenericRequest struct {
	QueryOptions
}

// GenericResponse is used to respond to a request where no
// specific response information is needed.
type GenericResponse struct {
	WriteMeta
}

// VersionResponse is used for the Status.Version response
type VersionResponse struct {
	Build   string
	Version string
	QueryMeta
}

// Response http/rpc response
type Response struct {
	Method     string
	ErrCode    cerr.ErrCodeType
	ErrMessage string
	Payload    interface{}
}

// DefaultResponse creates the default response
func DefaultResponse() *Response {
	return &Response{
		ErrCode:    cerr.SuccCode,
		ErrMessage: "",
	}
}
