// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "idev-cms-service/domain"

	mock "github.com/stretchr/testify/mock"
)

// Permission is an autogenerated mock type for the Permission type
type Permission struct {
	mock.Mock
}

// ConvertCRUDToLevel provides a mock function with given fields: isCreate, isRead, isUpdate, isDelete
func (_m *Permission) ConvertCRUDToLevel(isCreate bool, isRead bool, isUpdate bool, isDelete bool) uint8 {
	ret := _m.Called(isCreate, isRead, isUpdate, isDelete)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(bool, bool, bool, bool) uint8); ok {
		r0 = rf(isCreate, isRead, isUpdate, isDelete)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	return r0
}

// ConvertLevelToCRUD provides a mock function with given fields: level
func (_m *Permission) ConvertLevelToCRUD(level uint8) *domain.RolesPermission {
	ret := _m.Called(level)

	var r0 *domain.RolesPermission
	if rf, ok := ret.Get(0).(func(uint8) *domain.RolesPermission); ok {
		r0 = rf(level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.RolesPermission)
		}
	}

	return r0
}

// ConvertSliceLevelToSliceCRUD provides a mock function with given fields: input
func (_m *Permission) ConvertSliceLevelToSliceCRUD(input []*domain.Permissions) []*domain.PermissionsCRUD {
	ret := _m.Called(input)

	var r0 []*domain.PermissionsCRUD
	if rf, ok := ret.Get(0).(func([]*domain.Permissions) []*domain.PermissionsCRUD); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.PermissionsCRUD)
		}
	}

	return r0
}

// ConvertSubMenuToSubMenuView provides a mock function with given fields: input
func (_m *Permission) ConvertSubMenuToSubMenuView(input []*domain.SubMenu) []*domain.SubMenuView {
	ret := _m.Called(input)

	var r0 []*domain.SubMenuView
	if rf, ok := ret.Get(0).(func([]*domain.SubMenu) []*domain.SubMenuView); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.SubMenuView)
		}
	}

	return r0
}