package errs

import (
	"fmt"

	"github.com/pkg/errors"
)

func Example4() {
	err := errors.Errorf("whoops: %s %s %s", "foo", "bar", "baz")
	fmt.Printf("%+v", err)
}
