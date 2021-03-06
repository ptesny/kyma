// Code generated by mockery v1.1.2. DO NOT EDIT.

package automock

import (
	internal "github.com/kyma-project/kyma/components/application-broker/internal"
	mock "github.com/stretchr/testify/mock"
)

// operationInserter is an autogenerated mock type for the operationInserter type
type OperationInserter struct {
	mock.Mock
}

// Insert provides a mock function with given fields: io
func (_m *OperationInserter) Insert(io *internal.InstanceOperation) error {
	ret := _m.Called(io)

	var r0 error
	if rf, ok := ret.Get(0).(func(*internal.InstanceOperation) error); ok {
		r0 = rf(io)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
