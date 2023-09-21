package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"idev-cms-service/service/users/inout"
)

const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PUT    = "PUT"
	METHOD_DELETE = "DELETE"
	CONTENT_TYPE  = "Content-Type"
	CONTENT_VALUE = "application/json"
)

func (suite *PackageTestSuite) makeListRequest() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, _ = http.NewRequest(METHOD_GET, "/users?page=1&per_page=15", nil)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeListRequestInvalidQueryParam() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, _ = http.NewRequest(METHOD_GET, "/users?page={{}}&per_page={{}}", nil)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeCreateRequest(input *inout.UserCreateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_POST, "/users", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeCreateRequestInvalidJSON() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")
	req, _ = http.NewRequest(METHOD_POST, "/users", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeUpdateRequest(input *inout.UserUpdateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_PUT, fmt.Sprintf("/users/%s", input.ID), bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeUpdateRequestInvalidJSON() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")
	req, _ = http.NewRequest(METHOD_PUT, "/users/test", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeReadRequest(input *inout.ReadInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, _ = http.NewRequest(METHOD_GET, fmt.Sprintf("/users/%s", input.ID), nil)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeDeleteRequest(input *inout.UserDeleteInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_DELETE, fmt.Sprintf("/users/%s", input.ID), bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeGetMeRequest() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	req, _ = http.NewRequest(METHOD_GET, "/me", nil)
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeUpdateMeRequest(input *inout.ProfileUpdateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_PUT, "/me", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeUpdateMeRequestInvalidJSON() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")
	req, _ = http.NewRequest(METHOD_PUT, "/me", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeResendActivateRequest(input *inout.ResendActivateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_POST, "/activate/resend", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeResendActivateRequestInvalidJSON() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")
	req, _ = http.NewRequest(METHOD_POST, "/activate/resend", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeActivateRequest(input *inout.UserActivateInput) (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes, _ := json.Marshal(input)
	req, _ = http.NewRequest(METHOD_POST, "/me/activate", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}

func (suite *PackageTestSuite) makeActivateRequestInvalidJSON() (req *http.Request, w *httptest.ResponseRecorder, err error) {
	jsonBytes := []byte("{{{}}}")
	req, _ = http.NewRequest(METHOD_POST, "/me/activate", bytes.NewBuffer(jsonBytes))
	req.Header.Set(CONTENT_TYPE, CONTENT_VALUE)
	return req, httptest.NewRecorder(), nil
}
