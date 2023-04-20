// Code generated by mockery v2.25.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "karhub.backend.developer.test/src/api/v1/model"
)

// BeerRepository is an autogenerated mock type for the BeerRepository type
type BeerRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *BeerRepository) Create(data model.Beer) (model.Beer, error) {
	ret := _m.Called(data)

	var r0 model.Beer
	var r1 error
	if rf, ok := ret.Get(0).(func(model.Beer) (model.Beer, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(model.Beer) model.Beer); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(model.Beer)
	}

	if rf, ok := ret.Get(1).(func(model.Beer) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *BeerRepository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *BeerRepository) GetAll() ([]model.Beer, error) {
	ret := _m.Called()

	var r0 []model.Beer
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Beer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Beer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Beer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *BeerRepository) GetByID(id int) (model.Beer, error) {
	ret := _m.Called(id)

	var r0 model.Beer
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (model.Beer, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) model.Beer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Beer)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, data
func (_m *BeerRepository) Update(id int, data model.Beer) (model.Beer, error) {
	ret := _m.Called(id, data)

	var r0 model.Beer
	var r1 error
	if rf, ok := ret.Get(0).(func(int, model.Beer) (model.Beer, error)); ok {
		return rf(id, data)
	}
	if rf, ok := ret.Get(0).(func(int, model.Beer) model.Beer); ok {
		r0 = rf(id, data)
	} else {
		r0 = ret.Get(0).(model.Beer)
	}

	if rf, ok := ret.Get(1).(func(int, model.Beer) error); ok {
		r1 = rf(id, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBeerRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBeerRepository creates a new instance of BeerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBeerRepository(t mockConstructorTestingTNewBeerRepository) *BeerRepository {
	mock := &BeerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}