// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ses "github.com/applike/gosoline/pkg/ses"
	mock "github.com/stretchr/testify/mock"
)

// SimpleMailer is an autogenerated mock type for the SimpleMailer type
type SimpleMailer struct {
	mock.Mock
}

// Send provides a mock function with given fields: ctx, message
func (_m *SimpleMailer) Send(ctx context.Context, message ses.Message) error {
	ret := _m.Called(ctx, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ses.Message) error); ok {
		r0 = rf(ctx, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
