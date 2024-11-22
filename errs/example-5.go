package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

func Example5() {
	cause := errors.New("whoops")
	err := errors.WithMessage(cause, "oh noes")
	fmt.Println(err)
}
