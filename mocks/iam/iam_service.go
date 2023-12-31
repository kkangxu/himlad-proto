// Code generated by mockery v2.9.4. DO NOT EDIT.

package iam

import (
	context "context"

	client "github.com/asim/go-micro/v3/client"

	iam "github.com/kkangxu/himlad-proto/iam"

	mock "github.com/stretchr/testify/mock"
)

// IamService is an autogenerated mock type for the IamService type
type IamService struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: ctx, in, opts
func (_m *IamService) GetUser(ctx context.Context, in *iam.ReqGetUser, opts ...client.CallOption) (*iam.RspGetUser, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *iam.RspGetUser
	if rf, ok := ret.Get(0).(func(context.Context, *iam.ReqGetUser, ...client.CallOption) *iam.RspGetUser); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*iam.RspGetUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *iam.ReqGetUser, ...client.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
