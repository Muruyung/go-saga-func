package sagafunc

import "errors"

// AddStep add saga step
func (sg *sagaInteractor) AddStep(step *Step) error {
	if step == nil {
		return errors.New("invalid step, cannot be nil")
	}

	if step.ExecutionFunc == nil {
		return errors.New("invalid execution function, cannot be empty")
	}

	if step.RollbackFunc == nil {
		return errors.New("invalid rollback function, cannot be empty")
	}

	sg.steps[sg.countStep] = step
	sg.countStep++
	return nil
}
