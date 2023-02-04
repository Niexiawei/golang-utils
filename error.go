package golangutils

import (
	"fmt"
	"io"
	"reflect"
	"runtime/debug"
)

type errorWithStack struct {
	error
	stack
}

type stack string

func callers() stack {
	return stack(fmt.Sprintf("%s", debug.Stack()))
}

func NewStackErr(err error) error {
	if err == nil {
		return nil
	}
	return &errorWithStack{
		err,
		callers(),
	}
}

func (w *errorWithStack) Unwrap() error { return w.error }

func (w *errorWithStack) Is(err error) bool {
	return reflect.TypeOf(*w).Name() == reflect.TypeOf(err).Name()
}

func (w *errorWithStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = fmt.Fprintf(s, "%+v\n", w.Unwrap())
			_, _ = fmt.Fprintln(s, w.stack)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, w.Error())
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", w.Error())
	}
}