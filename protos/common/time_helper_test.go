package common

import (
	"reflect"
	"testing"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

func TestValueSucc(t *testing.T) {
	timestampWrapper := TimestampWrapper{Time: &google_protobuf.Timestamp{Seconds: 1519701982, Nanos: 126091123}}
	res, err := timestampWrapper.Value()
	if err != nil {
		t.Errorf("err should be nil, not %v", err)
	}
	timestampWrapperStr, ok := res.(string)
	if !ok {
		t.Errorf("res should be string type, not %v", reflect.TypeOf(res))
	}
	if timestampWrapperStr != "2018-02-27 03:26:22.126091123" {
		t.Errorf("timestampWrapperStr is not right ")
	}
}

func TestScanSucc(t *testing.T) {
	timestampWrapperRef := TimestampWrapper{Time: &google_protobuf.Timestamp{Seconds: 1519701982, Nanos: 126091123}}
	timestampWrapper := TimestampWrapper{Time: &google_protobuf.Timestamp{Seconds: 0, Nanos: 0}}
	var input interface{} = "2018-02-27 03:26:22.126091123"
	err := timestampWrapper.Scan(input)
	if err != nil {
		t.Errorf("err should be nil, not %v", err)
	}

	if timestampWrapper.Time.Seconds != timestampWrapperRef.Time.Seconds ||
		timestampWrapper.Time.Nanos != timestampWrapperRef.Time.Nanos {
		t.Errorf("timestampWrapper should be equal with timestampWrapperRef")
	}
}

func TestScanTimeIsNilSucc(t *testing.T) {
	timestampWrapperRef := TimestampWrapper{Time: &google_protobuf.Timestamp{Seconds: 1519701982, Nanos: 126091123}}
	timestampWrapper := TimestampWrapper{}
	var input interface{} = "2018-02-27 03:26:22.126091123"
	err := timestampWrapper.Scan(input)
	if err != nil {
		t.Errorf("err should be nil, not %v", err)
	}

	if timestampWrapper.Time.Seconds != timestampWrapperRef.Time.Seconds ||
		timestampWrapper.Time.Nanos != timestampWrapperRef.Time.Nanos {
		t.Errorf("timestampWrapper should be equal with timestampWrapperRef")
	}
}

func TestScanSplitFailed(t *testing.T) {
	timestampWrapper := TimestampWrapper{Time: &google_protobuf.Timestamp{Seconds: 0, Nanos: 0}}
	var input interface{} = "2018-02-27 03:26:22"
	err := timestampWrapper.Scan(input)
	if err == nil {
		t.Errorf("err should not be nil")
	}
}

func TestValueTimeIsNil(t *testing.T) {
	timestampWrapper := TimestampWrapper{}
	val, err := timestampWrapper.Value()
	if err != nil {
		t.Errorf("err should be nil")
	}
	if val != "" {
		t.Errorf("val should not be empty")
	}
}
