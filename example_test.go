package logr_test

import (
	"github.com/goph/logur"

	logrintegration "logur.dev/integration/logr"
)

func ExampleNew() {
	logger := logrintegration.New(logur.NewNoopLogger())

	// Output:
	_ = logger
}
