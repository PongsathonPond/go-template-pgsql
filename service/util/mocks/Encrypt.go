// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Encrypt is an autogenerated mock type for the Encrypt type
type Encrypt struct {
	mock.Mock
}

// CheckPasswordHash provides a mock function with given fields: password, hash
func (_m *Encrypt) CheckPasswordHash(password string, hash string) bool {
	ret := _m.Called(password, hash)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(password, hash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HashPassword provides a mock function with given fields: password
func (_m *Encrypt) HashPassword(password string) string {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MD5 provides a mock function with given fields: str
func (_m *Encrypt) MD5(str string) string {
	ret := _m.Called(str)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}