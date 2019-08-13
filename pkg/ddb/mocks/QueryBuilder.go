// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import ddb "github.com/applike/gosoline/pkg/ddb"
import expression "github.com/aws/aws-sdk-go/service/dynamodb/expression"
import mock "github.com/stretchr/testify/mock"

// QueryBuilder is an autogenerated mock type for the QueryBuilder type
type QueryBuilder struct {
	mock.Mock
}

// Build provides a mock function with given fields: result
func (_m *QueryBuilder) Build(result interface{}) (*ddb.QueryOperation, error) {
	ret := _m.Called(result)

	var r0 *ddb.QueryOperation
	if rf, ok := ret.Get(0).(func(interface{}) *ddb.QueryOperation); ok {
		r0 = rf(result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.QueryOperation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableTtlFilter provides a mock function with given fields:
func (_m *QueryBuilder) DisableTtlFilter() ddb.QueryBuilder {
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

// WithDescendingOrder provides a mock function with given fields:
func (_m *QueryBuilder) WithDescendingOrder() ddb.QueryBuilder {
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

// WithFilter provides a mock function with given fields: filter
func (_m *QueryBuilder) WithFilter(filter expression.ConditionBuilder) ddb.QueryBuilder {
	ret := _m.Called(filter)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(expression.ConditionBuilder) ddb.QueryBuilder); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithHash provides a mock function with given fields: value
func (_m *QueryBuilder) WithHash(value interface{}) ddb.QueryBuilder {
	ret := _m.Called(value)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.QueryBuilder); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithIndex provides a mock function with given fields: name
func (_m *QueryBuilder) WithIndex(name string) ddb.QueryBuilder {
	ret := _m.Called(name)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(string) ddb.QueryBuilder); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithLimit provides a mock function with given fields: limit
func (_m *QueryBuilder) WithLimit(limit int) ddb.QueryBuilder {
	ret := _m.Called(limit)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(int) ddb.QueryBuilder); ok {
		r0 = rf(limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithPageSize provides a mock function with given fields: size
func (_m *QueryBuilder) WithPageSize(size int) ddb.QueryBuilder {
	ret := _m.Called(size)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(int) ddb.QueryBuilder); ok {
		r0 = rf(size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithProjection provides a mock function with given fields: projection
func (_m *QueryBuilder) WithProjection(projection interface{}) ddb.QueryBuilder {
	ret := _m.Called(projection)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.QueryBuilder); ok {
		r0 = rf(projection)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithRange provides a mock function with given fields: comp, values
func (_m *QueryBuilder) WithRange(comp string, values ...interface{}) ddb.QueryBuilder {
	var _ca []interface{}
	_ca = append(_ca, comp)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(string, ...interface{}) ddb.QueryBuilder); ok {
		r0 = rf(comp, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithRangeBeginsWith provides a mock function with given fields: prefix
func (_m *QueryBuilder) WithRangeBeginsWith(prefix string) ddb.QueryBuilder {
	ret := _m.Called(prefix)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(string) ddb.QueryBuilder); ok {
		r0 = rf(prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithRangeBetween provides a mock function with given fields: lower, upper
func (_m *QueryBuilder) WithRangeBetween(lower interface{}, upper interface{}) ddb.QueryBuilder {
	ret := _m.Called(lower, upper)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) ddb.QueryBuilder); ok {
		r0 = rf(lower, upper)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithRangeEq provides a mock function with given fields: value
func (_m *QueryBuilder) WithRangeEq(value interface{}) ddb.QueryBuilder {
	ret := _m.Called(value)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.QueryBuilder); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithRangeGt provides a mock function with given fields: value
func (_m *QueryBuilder) WithRangeGt(value interface{}) ddb.QueryBuilder {
	ret := _m.Called(value)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.QueryBuilder); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithRangeGte provides a mock function with given fields: value
func (_m *QueryBuilder) WithRangeGte(value interface{}) ddb.QueryBuilder {
	ret := _m.Called(value)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.QueryBuilder); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithRangeLt provides a mock function with given fields: value
func (_m *QueryBuilder) WithRangeLt(value interface{}) ddb.QueryBuilder {
	ret := _m.Called(value)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.QueryBuilder); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// WithRangeLte provides a mock function with given fields: value
func (_m *QueryBuilder) WithRangeLte(value interface{}) ddb.QueryBuilder {
	ret := _m.Called(value)

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.QueryBuilder); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}
