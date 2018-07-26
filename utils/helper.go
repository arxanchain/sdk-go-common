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
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

// GetFQDN Fully Qualified Domain Name
// returns "unknown" or hostanme in case of error
func GetFQDN() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}

	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return hostname
	}

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip, err := ipv4.MarshalText()
			if err != nil {
				return hostname
			}
			hosts, err := net.LookupAddr(string(ip))
			if err != nil || len(hosts) == 0 {
				return hostname
			}
			fqdn := hosts[0]
			return strings.TrimSuffix(fqdn, ".") // return fqdn without trailing dot
		}
	}
	return hostname
}

// validUUID is used to check if a given string looks like a UUID
var validUUID = regexp.MustCompile(`(?i)^[\da-f]{8}-[\da-f]{4}-[\da-f]{4}-[\da-f]{4}-[\da-f]{12}$`)

// IsUUID returns true if the given string is a valid UUID.
func IsUUID(str string) bool {
	const uuidLen = 36
	if len(str) != uuidLen {
		return false
	}

	return validUUID.MatchString(str)
}

// boolToPtr returns the pointer to a boolean
func BoolToPtr(b bool) *bool {
	return &b
}

// IntToPtr returns the pointer to an int
func IntToPtr(i int) *int {
	return &i
}

// UintToPtr returns the pointer to an uint
func Uint64ToPtr(u uint64) *uint64 {
	return &u
}

// StringToPtr returns the pointer to a string
func StringToPtr(str string) *string {
	return &str
}

// TimeToPtr returns the pointer to a time stamp
func TimeToPtr(t time.Duration) *time.Duration {
	return &t
}

// MapStringStringSliceValueSet returns the set of values in a map[string][]string
func MapStringStringSliceValueSet(m map[string][]string) []string {
	set := make(map[string]struct{})
	for _, slice := range m {
		for _, v := range slice {
			set[v] = struct{}{}
		}
	}

	flat := make([]string, 0, len(set))
	for k := range set {
		flat = append(flat, k)
	}
	return flat
}

func SliceStringToSet(s []string) map[string]struct{} {
	m := make(map[string]struct{}, (len(s)+1)/2)
	for _, k := range s {
		m[k] = struct{}{}
	}
	return m
}

// SliceStringIsSubset returns whether the smaller set of strings is a subset of
// the larger. If the smaller slice is not a subset, the offending elements are
// returned.
func SliceStringIsSubset(larger, smaller []string) (bool, []string) {
	largerSet := make(map[string]struct{}, len(larger))
	for _, l := range larger {
		largerSet[l] = struct{}{}
	}

	subset := true
	var offending []string
	for _, s := range smaller {
		if _, ok := largerSet[s]; !ok {
			subset = false
			offending = append(offending, s)
		}
	}

	return subset, offending
}

func SliceSetDisjoint(first, second []string) (bool, []string) {
	contained := make(map[string]struct{}, len(first))
	for _, k := range first {
		contained[k] = struct{}{}
	}

	offending := make(map[string]struct{})
	for _, k := range second {
		if _, ok := contained[k]; ok {
			offending[k] = struct{}{}
		}
	}

	if len(offending) == 0 {
		return true, nil
	}

	flattened := make([]string, 0, len(offending))
	for k := range offending {
		flattened = append(flattened, k)
	}
	return false, flattened
}

// Helpers for copying generic structures.
func CopyMapStringString(m map[string]string) map[string]string {
	l := len(m)
	if l == 0 {
		return nil
	}

	c := make(map[string]string, l)
	for k, v := range m {
		c[k] = v
	}
	return c
}

func CopyMapStringInt(m map[string]int) map[string]int {
	l := len(m)
	if l == 0 {
		return nil
	}

	c := make(map[string]int, l)
	for k, v := range m {
		c[k] = v
	}
	return c
}

func CopyMapStringFloat64(m map[string]float64) map[string]float64 {
	l := len(m)
	if l == 0 {
		return nil
	}

	c := make(map[string]float64, l)
	for k, v := range m {
		c[k] = v
	}
	return c
}

func CopySliceString(s []string) []string {
	l := len(s)
	if l == 0 {
		return nil
	}

	c := make([]string, l)
	for i, v := range s {
		c[i] = v
	}
	return c
}

func CopySliceInt(s []int) []int {
	l := len(s)
	if l == 0 {
		return nil
	}

	c := make([]int, l)
	for i, v := range s {
		c[i] = v
	}
	return c
}

// CleanEnvVar replaces all occurrences of illegal characters in an environment
// variable with the specified byte.
func CleanEnvVar(s string, r byte) string {
	b := []byte(s)
	for i, c := range b {
		switch {
		case c == '_':
		case c >= 'a' && c <= 'z':
		case c >= 'A' && c <= 'Z':
		case i > 0 && c >= '0' && c <= '9':
		default:
			// Replace!
			b[i] = r
		}
	}
	return string(b)
}

// Has returns true if the needle is in the haystack (case-sensitive)
func StringSliceHas(haystack []string, needle string) bool {
	for _, current := range haystack {
		if current == needle {
			return true
		}
	}
	return false
}

// HasI returns true if the needle is in the haystack (case-insensitive)
func StringSliceHasI(haystack []string, needle string) bool {
	for _, current := range haystack {
		if strings.ToLower(current) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}
