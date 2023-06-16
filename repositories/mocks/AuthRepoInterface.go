// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entities "user-management-backend/entities"

	mock "github.com/stretchr/testify/mock"
)

// AuthRepoInterface is an autogenerated mock type for the AuthRepoInterface type
type AuthRepoInterface struct {
	mock.Mock
}

// CreateActorSession provides a mock function with given fields: session
func (_m *AuthRepoInterface) CreateActorSession(session *entities.ActorSession) error {
	ret := _m.Called(session)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.ActorSession) error); ok {
		r0 = rf(session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetActorByUsername provides a mock function with given fields: username
func (_m *AuthRepoInterface) GetActorByUsername(username *string) (entities.Actor, error) {
	ret := _m.Called(username)

	var r0 entities.Actor
	var r1 error
	if rf, ok := ret.Get(0).(func(*string) (entities.Actor, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(*string) entities.Actor); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(entities.Actor)
	}

	if rf, ok := ret.Get(1).(func(*string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastActorSessionByToken provides a mock function with given fields: token
func (_m *AuthRepoInterface) GetLastActorSessionByToken(token *string) (entities.ActorSession, error) {
	ret := _m.Called(token)

	var r0 entities.ActorSession
	var r1 error
	if rf, ok := ret.Get(0).(func(*string) (entities.ActorSession, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(*string) entities.ActorSession); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(entities.ActorSession)
	}

	if rf, ok := ret.Get(1).(func(*string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthRepoInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthRepoInterface creates a new instance of AuthRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthRepoInterface(t mockConstructorTestingTNewAuthRepoInterface) *AuthRepoInterface {
	mock := &AuthRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
