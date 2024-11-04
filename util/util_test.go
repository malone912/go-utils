package util

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("test", func(t *testing.T) {

		var number = 0
		assert.NoError(t, Run(func() error {
			number = 1
			return nil
		}))

		assert.Equal(t, 1, number)

		assert.Error(t, Run(func() error {
			number = 2
			return errors.New("test")
		}))

		assert.Equal(t, 2, number)

	})
}

func TestSet(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		var myFuture = 0

		err := Run(Set(&myFuture, func() (int, error) {
			return 1, nil
		}))
		assert.NoError(t, err)
		assert.Equal(t, 1, myFuture)

	})
}

func TestMulti(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		var a = 0
		var b = 0
		var c = 0

		err :=
			Run(func() error { return nil },
				Set(&a, func() (int, error) { return 1, nil }),
				Void(func() {}),
				Set(&b, func() (int, error) { return 2, nil }),
				func() error { return errors.New("test") },
				Set(&c, func() (int, error) { return 3, nil }))

		assert.Error(t, err)
		assert.Equal(t, 1, a)
		assert.Equal(t, 2, b)
		assert.Equal(t, 0, c)
	})
}

