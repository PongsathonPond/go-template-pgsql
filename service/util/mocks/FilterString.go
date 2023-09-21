// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FilterString is an autogenerated mock type for the FilterString type
type FilterString struct {
	mock.Mock
}

// MakeAccessToken provides a mock function with given fields: accessToken
func (_m *FilterString) MakeAccessToken(accessToken string) string {
	ret := _m.Called(accessToken)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(accessToken)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeAccessTokenExpiredAtGreaterThan provides a mock function with given fields: time
func (_m *FilterString) MakeAccessTokenExpiredAtGreaterThan(time int64) string {
	ret := _m.Called(time)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(time)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeActivateExpiredAtGreaterThan provides a mock function with given fields: time
func (_m *FilterString) MakeActivateExpiredAtGreaterThan(time int64) string {
	ret := _m.Called(time)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(time)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeActivateTokenString provides a mock function with given fields: activateTokenString
func (_m *FilterString) MakeActivateTokenString(activateTokenString string) string {
	ret := _m.Called(activateTokenString)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(activateTokenString)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeDeletedAtIsNull provides a mock function with given fields:
func (_m *FilterString) MakeDeletedAtIsNull() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeEmail provides a mock function with given fields: email
func (_m *FilterString) MakeEmail(email string) string {
	ret := _m.Called(email)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeEqualUserType provides a mock function with given fields: userType
func (_m *FilterString) MakeEqualUserType(userType string) string {
	ret := _m.Called(userType)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userType)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeID provides a mock function with given fields: id
func (_m *FilterString) MakeID(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeIDFilters provides a mock function with given fields: id
func (_m *FilterString) MakeIDFilters(id string) []string {
	ret := _m.Called(id)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MakeKeySlug provides a mock function with given fields: keySlug
func (_m *FilterString) MakeKeySlug(keySlug string) string {
	ret := _m.Called(keySlug)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(keySlug)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeNotEqualID provides a mock function with given fields: id
func (_m *FilterString) MakeNotEqualID(id string) string {
	ret := _m.Called(id)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakePassword provides a mock function with given fields: password
func (_m *FilterString) MakePassword(password string) string {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakePermission provides a mock function with given fields: key
func (_m *FilterString) MakePermission(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeRefreshToken provides a mock function with given fields: refreshToken
func (_m *FilterString) MakeRefreshToken(refreshToken string) string {
	ret := _m.Called(refreshToken)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(refreshToken)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeRefreshTokenExpiredAtLessThan provides a mock function with given fields: time
func (_m *FilterString) MakeRefreshTokenExpiredAtLessThan(time int64) string {
	ret := _m.Called(time)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(time)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeRoleID provides a mock function with given fields: roleID
func (_m *FilterString) MakeRoleID(roleID string) string {
	ret := _m.Called(roleID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(roleID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeStatus provides a mock function with given fields: status
func (_m *FilterString) MakeStatus(status string) string {
	ret := _m.Called(status)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeStatusIn provides a mock function with given fields: status
func (_m *FilterString) MakeStatusIn(status ...string) string {
	_va := make([]interface{}, len(status))
	for _i := range status {
		_va[_i] = status[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(...string) string); ok {
		r0 = rf(status...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeSubMenuIn provides a mock function with given fields: subMenu
func (_m *FilterString) MakeSubMenuIn(subMenu []string) string {
	ret := _m.Called(subMenu)

	var r0 string
	if rf, ok := ret.Get(0).(func([]string) string); ok {
		r0 = rf(subMenu)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeUserAgent provides a mock function with given fields: userAgent
func (_m *FilterString) MakeUserAgent(userAgent string) string {
	ret := _m.Called(userAgent)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userAgent)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeUserGroup provides a mock function with given fields: status
func (_m *FilterString) MakeUserGroup(status string) string {
	ret := _m.Called(status)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(status)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeUserGroupRoleID provides a mock function with given fields: roleID
func (_m *FilterString) MakeUserGroupRoleID(roleID string) string {
	ret := _m.Called(roleID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(roleID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeUserID provides a mock function with given fields: userID
func (_m *FilterString) MakeUserID(userID string) string {
	ret := _m.Called(userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeUserName provides a mock function with given fields: username
func (_m *FilterString) MakeUserName(username string) string {
	ret := _m.Called(username)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeVerified provides a mock function with given fields: val
func (_m *FilterString) MakeVerified(val bool) string {
	ret := _m.Called(val)

	var r0 string
	if rf, ok := ret.Get(0).(func(bool) string); ok {
		r0 = rf(val)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MakeVerifiedBool provides a mock function with given fields: val
func (_m *FilterString) MakeVerifiedBool(val bool) string {
	ret := _m.Called(val)

	var r0 string
	if rf, ok := ret.Get(0).(func(bool) string); ok {
		r0 = rf(val)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
