/**
 * Licensed Materials - Property of Arxan Fintech
 *
 * (C) Copyright Arxan Fintech. 2018 All Rights Reserved
 *
 * Contributors:
 *    Wang Zhongzhi
**/

package common

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"database/sql/driver"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
)

// Value implements the driver.Valuer interface.
func (p TimestampWrapper) Value() (driver.Value, error) {
	tempTime := p.GetTime()
	if tempTime == nil {
		return "", nil
	}
	nanos := strconv.FormatInt((int64)(tempTime.GetNanos()), 10)
	res := time.Unix(tempTime.GetSeconds(), 0).UTC().Format("2006-01-02 15:04:05")
	res += "."
	res += nanos
	return res, nil
}

// Scan implements the sql.Scanner interface.
func (p *TimestampWrapper) Scan(src interface{}) error {
	source, ok := src.(string)
	if !ok {
		return errors.New("type assertion .(string) failed")
	}
	// empty should not be error
	if source == "" {
		return nil
	}
	fields := strings.Split(source, ".")
	if len(fields) < 2 {
		return errors.New("split nanos from source failed")
	}
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", fields[0], time.UTC)
	if err != nil {
		return err
	}

	timestampSec := theTime.Unix()
	tempTime := &google_protobuf.Timestamp{}
	tempTime.Seconds = timestampSec
	nanosStr := strings.Trim(fields[1], "Z")
	nanos, err := strconv.Atoi(nanosStr)
	if err != nil {
		return err
	}
	tempTime.Nanos = int32(nanos)
	p.Time = tempTime

	return nil
}
