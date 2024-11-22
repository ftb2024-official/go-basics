/*
	Cause (Printf)
*/

package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

func Example3() {
	err := errors.Wrap(func() error {
		return func() error {
			return errors.New("hello, world")
		}()
	}(), "failed")

	fmt.Printf("%v", err)
}
