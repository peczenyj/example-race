package foo_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/peczenyj/example-race/foo"
)

func TestBatcher(t *testing.T) {
	t.Parallel()

	mockBucket := NewBucket(t)

	mockBucket.On("Do", mock.AnythingOfType("[]string")).Return(nil).Twice()

	instance := foo.New(mockBucket)

	require.NoError(t, instance.Get("test_key1"))
	require.NoError(t, instance.Get("test_key2"))
}

// mock, usually self generated

type Bucket struct {
	mock.Mock
}

func NewBucket(t interface {
	mock.TestingT
	Cleanup(func())
}) *Bucket {
	mock := &Bucket{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

func (_m *Bucket) Do(ops []string) error {
	ret := _m.Called(ops)

	var err error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		err = rf(ops)
	} else {
		err = ret.Error(0)
	}

	return err //nolint: wrapcheck
}
