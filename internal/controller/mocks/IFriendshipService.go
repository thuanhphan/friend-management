// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	model "friend-management-go/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// IFriendshipService is an autogenerated mock type for the IFriendshipService type
type IFriendshipService struct {
	mock.Mock
}

// FriendshipExists provides a mock function with given fields: email1, eamil2
func (_m *IFriendshipService) FriendshipExists(email1 string, eamil2 string) (bool, error) {
	ret := _m.Called(email1, eamil2)

	if len(ret) == 0 {
		panic("no return value specified for FriendshipExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(email1, eamil2)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(email1, eamil2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email1, eamil2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommonFriends provides a mock function with given fields: email1, email2
func (_m *IFriendshipService) GetCommonFriends(email1 string, email2 string) ([]string, error) {
	ret := _m.Called(email1, email2)

	if len(ret) == 0 {
		panic("no return value specified for GetCommonFriends")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]string, error)); ok {
		return rf(email1, email2)
	}
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(email1, email2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email1, email2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFriends provides a mock function with given fields: email
func (_m *IFriendshipService) GetFriends(email string) ([]string, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetFriends")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReceivableUpdates provides a mock function with given fields: email
func (_m *IFriendshipService) GetReceivableUpdates(email string) ([]string, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetReceivableUpdates")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeFriend provides a mock function with given fields: friendship
func (_m *IFriendshipService) MakeFriend(friendship model.Friendship) error {
	ret := _m.Called(friendship)

	if len(ret) == 0 {
		panic("no return value specified for MakeFriend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Friendship) error); ok {
		r0 = rf(friendship)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFriendshipStatus provides a mock function with given fields: friendship
func (_m *IFriendshipService) UpdateFriendshipStatus(friendship model.Friendship) error {
	ret := _m.Called(friendship)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFriendshipStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Friendship) error); ok {
		r0 = rf(friendship)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIFriendshipService creates a new instance of IFriendshipService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIFriendshipService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IFriendshipService {
	mock := &IFriendshipService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
