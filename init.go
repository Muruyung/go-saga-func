package sagafunc

type sagaInteractor struct {
	result              *Result
	steps               map[int]*Step
	countStep           int
	skipCompensateError bool
}

// NewSaga initialize new saga (default FALSE)
//   - Set parameter TRUE if you won't execute compensate on error func
//   - Set parameter FALSE (or empty this parameter) if you want execute all compensate
func NewSaga(skipCompensateError ...bool) Saga {
	skip := false
	if len(skipCompensateError) > 0 {
		skip = skipCompensateError[0]
	}

	return &sagaInteractor{
		result:              nil,
		steps:               make(map[int]*Step),
		countStep:           0,
		skipCompensateError: skip,
	}
}
