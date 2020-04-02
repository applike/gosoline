// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import cfg "github.com/applike/gosoline/pkg/cfg"
import context "context"

import mock "github.com/stretchr/testify/mock"
import mon "github.com/applike/gosoline/pkg/mon"

// Module is an autogenerated mock type for the Module type
type Module struct {
	mock.Mock
}

// Boot provides a mock function with given fields: config, logger
func (_m *Module) Boot(config cfg.Config, logger mon.Logger) error {
	ret := _m.Called(config, logger)

	var r0 error
	if rf, ok := ret.Get(0).(func(cfg.Config, mon.Logger) error); ok {
		r0 = rf(config, logger)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields: ctx
func (_m *Module) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
