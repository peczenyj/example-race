package foo

import (
	"time"

	"github.com/eapache/go-resiliency/batcher"
)

type Batcher interface {
	Run(interface{}) error
}

type Bucket interface {
	Do([]string) error
}

type Foo struct {
	Bucket  Bucket
	batcher Batcher
	timeout time.Duration
}

func New(bucket Bucket) *Foo {
	foo := &Foo{ // nolint: exhaustruct
		Bucket:  bucket,
		timeout: time.Second,
	}

	foo.batcher = batcher.New(foo.timeout, foo.batch)

	return foo
}

// Get wraps a batcher around the standard Get method from gocb.Bucket.
func (b *Foo) Get(key string) error {
	return b.batcher.Run(key) //nolint: wrapcheck
}

func (b *Foo) batch(batch []interface{}) (err error) {
	ops := make([]string, len(batch))

	for i, op := range batch {
		ops[i], _ = op.(string)
	}

	return b.Bucket.Do(ops) //nolint: wrapcheck
}
