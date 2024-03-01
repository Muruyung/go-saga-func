package sagafunc

import (
	"github.com/hashicorp/go-multierror"
)

type Step struct {
	ProcessName    string
	ExecutionFunc  func() error
	CompensateFunc func() error
}

type Result struct {
	ExecutionError   error
	CompensateErrors *multierror.Error
}
