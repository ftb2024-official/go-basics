/*
	Cause
*/

package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

func fn2() error {
	e1 := errors.New("error")
	e2 := errors.Wrap(e1, "inner")
	e3 := errors.Wrap(e2, "middle")
	return errors.Wrap(e3, "outer")
}

func Example2() {
	err := fn2()
	fmt.Println(err)
	fmt.Println(errors.Cause(err))
}
