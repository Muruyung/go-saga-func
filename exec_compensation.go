package sagafunc

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
)

func (sg *sagaInteractor) execRollback(index int) {
	for i := index; i >= 0; i-- {
		logrus.WithFields(logrus.Fields{
			"step":        i,
			"processName": sg.steps[i].ProcessName,
		}).Warn("rollback process")

		err := sg.steps[i].RollbackFunc()
		if err != nil {
			err = fmt.Errorf("step: %d, process_name: %s, err: %v", i, sg.steps[i].ProcessName, err)
			sg.result.RollbackErrors = multierror.Append(sg.result.RollbackErrors, err)
		}
	}

	if sg.result.RollbackErrors != nil {
		logrus.Error(sg.result.RollbackErrors.Error())
	}
}
