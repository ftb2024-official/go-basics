package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

func fn() error {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}

func Example1() {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	err, ok := errors.Cause(fn()).(stackTracer)
	if !ok {
		panic("oops, err doesn't implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Println(st)
	fmt.Printf("%+v", st[:2]) // top two frames
}
