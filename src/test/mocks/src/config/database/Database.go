// Code generated by mockery v2.25.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// Database is an autogenerated mock type for the Database type
type Database struct {
	mock.Mock
}

// connect provides a mock function with given fields: dns
func (_m *Database) connect(dns string) *gorm.DB {
	ret := _m.Called(dns)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(string) *gorm.DB); ok {
		r0 = rf(dns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

type mockConstructorTestingTNewDatabase interface {
	mock.TestingT
	Cleanup(func())
}

// NewDatabase creates a new instance of Database. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDatabase(t mockConstructorTestingTNewDatabase) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
