package sagafunc

import (
	"github.com/hashicorp/go-multierror"
)

type Step struct {
	ProcessName   string
	ExecutionFunc func() error
	RollbackFunc  func() error
}

type Result struct {
	ExecutionError error
	RollbackErrors *multierror.Error
}
