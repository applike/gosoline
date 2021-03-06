// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import ddb "github.com/applike/gosoline/pkg/ddb"
import mock "github.com/stretchr/testify/mock"

// QueryBuilderSimple is an autogenerated mock type for the QueryBuilderSimple type
type QueryBuilderSimple struct {
	mock.Mock
}

// Build provides a mock function with given fields:
func (_m *QueryBuilderSimple) Build() ddb.QueryBuilder {
	ret := _m.Called()

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func() ddb.QueryBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithHash provides a mock function with given fields: value
func (_m *QueryBuilderSimple) WithHash(value interface{}) ddb.QueryBuilderSimple {
	ret := _m.Called(value)

	var r0 ddb.QueryBuilderSimple
	if rf, ok := ret.Get(0).(func(interface{}) ddb.QueryBuilderSimple); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilderSimple)
		}
	}

	return r0
}

// WithRange provides a mock function with given fields: comp, values
func (_m *QueryBuilderSimple) WithRange(comp string, values ...interface{}) ddb.QueryBuilderSimple {
	var _ca []interface{}
	_ca = append(_ca, comp)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 ddb.QueryBuilderSimple
	if rf, ok := ret.Get(0).(func(string, ...interface{}) ddb.QueryBuilderSimple); ok {
		r0 = rf(comp, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilderSimple)
		}
	}

	return r0
}
