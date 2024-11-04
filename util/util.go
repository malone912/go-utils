package util

import (
	"github.com/pkg/errors"
	"log/slog"
)

func Run(f ...func() error) error {
	for _, function := range f {
		err := function()
		if err != nil {
			return errors.Wrap(err, "Function returned error")
		}
	}
	return nil
}

func Set[T any](toSet *T, f func() (T, error)) func() error {
	return func() error {
		result, err := f()
		*toSet = result
		if err != nil {
			return errors.Wrap(err, "Function returned error")
		}
		return nil
	}
}

type Closer interface {
	Close() error
}

func Close(c Closer) {
	if err := c.Close(); err != nil {
		slog.Warn("Failed to close struct", slog.Any("struct", c))
	}
}

func Void(f func()) func() error {
	return func() error {
		f()
		return nil
	}
}

func Try(f func() error) {
	err := f()
	slog.Warn("Got error when attempting to execute.", slog.Any("function", f), slog.Any("error", err))
}

