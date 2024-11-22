package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

func Example6() {
	cause := errors.New("whoops")
	err := errors.Wrapf(cause, "oh noes #%d", 3)
	fmt.Println(err)
}
