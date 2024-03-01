package sagafunc

type sagaInteractor struct {
	result            *Result
	steps             map[int]*Step
	countStep         int
	skipRollbackError bool
}

// NewSaga initialize new saga (default FALSE)
//   - Set parameter TRUE if you won't execute rollback on error func
//   - Set parameter FALSE (or empty this parameter) if you want execute all rollback
func NewSaga(skipRollbackError ...bool) Saga {
	skip := false
	if len(skipRollbackError) > 0 {
		skip = skipRollbackError[0]
	}

	return &sagaInteractor{
		result:            nil,
		steps:             make(map[int]*Step),
		countStep:         0,
		skipRollbackError: skip,
	}
}
