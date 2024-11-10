package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract from http headers
// Authorization: ApiKey {insert apikey}
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil
}
