package logr_test

import (
	"logur.dev/logur"

	logrintegration "logur.dev/integration/logr"
)

func ExampleNew() {
	logger := logrintegration.New(logur.NoopLogger{})

	// Output:
	_ = logger
}
