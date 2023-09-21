package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"idev-cms-service/service/authen/inout"
)

const (
	METHOD_POST   = "POST"
	CONTENT_TYPE  = "Content-Type"
	CONTENT_VALUE = "application/json"
)

func (suite *PackageTestSuite) makeLoginRequest(input *inout.LoginInput) (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_POST, "/login", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeLoginRequestInvalidJson() (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes := []byte("{{{}}}")
	req, _ = http.NewRequest(METHOD_POST, "/login", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}

func (suite *PackageTestSuite) makeLogoutRequest(input *inout.LogoutInput) (req *http.Request, w *httptest.ResponseRecorder) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_POST, "/logout", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder()
}
