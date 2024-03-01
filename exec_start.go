package sagafunc

import "github.com/sirupsen/logrus"

// ExecStart start execute saga pattern
func (sg *sagaInteractor) ExecStart() error {
	var indexErr int

	logrus.Info("SAGA EXECUTION START")

	for i := 0; i < sg.countStep; i++ {
		logrus.WithFields(logrus.Fields{
			"step":        i,
			"processName": sg.steps[i].ProcessName,
		}).Info("execution process")

		err := sg.steps[i].ExecutionFunc()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"step":        i,
				"processName": sg.steps[i].ProcessName,
				"error":       err.Error(),
			}).Error("execution process")

			sg.result = &Result{
				ExecutionError: err,
			}
			indexErr = i
			break
		}
	}

	if sg.result != nil {
		logrus.Warn("SAGA ROLLBACK PROCESS")
		if sg.skipCompensateError {
			indexErr -= 1
		}
		go sg.execCompensation(indexErr)

		return sg.result.ExecutionError
	}

	logrus.Info("SAGA EXECUTION DONE")

	return nil
}
