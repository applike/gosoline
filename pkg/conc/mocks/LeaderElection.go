// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// LeaderElection is an autogenerated mock type for the LeaderElection type
type LeaderElection struct {
	mock.Mock
}

// IsLeader provides a mock function with given fields: ctx, memberId
func (_m *LeaderElection) IsLeader(ctx context.Context, memberId string) (bool, error) {
	ret := _m.Called(ctx, memberId)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, memberId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, memberId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Resign provides a mock function with given fields: ctx, memberId
func (_m *LeaderElection) Resign(ctx context.Context, memberId string) error {
	ret := _m.Called(ctx, memberId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, memberId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
