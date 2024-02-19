package foo_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/peczenyj/example-race/foo"
)

type nopBucket struct{}

func (*nopBucket) Do([]string) error { return nil }

func TestBatcher(t *testing.T) {
	t.Parallel()

	var mockBucket nopBucket

	instance := foo.New(&mockBucket)

	require.NoError(t, instance.Get("test_key1"))
	require.NoError(t, instance.Get("test_key2"))
}
