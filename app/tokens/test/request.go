package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"idev-cms-service/service/tokens/inout"
)

const (
	METHOD_POST   = "POST"
	CONTENT_TYPE  = "Content-Type"
	CONTENT_VALUE = "application/json"
)

func (suite *PackageTestSuite) makeRefreshTokenRequest(input *inout.RefreshTokenInput) (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_POST, "/token/refresh", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeRefreshTokenRequestInvalidJson() (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes := []byte("{{{}}}")
	req, _ = http.NewRequest(METHOD_POST, "/token/refresh", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeRevokeTokenRequest(input *inout.RevokeTokenInput) (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_POST, "/token/revoke", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}
