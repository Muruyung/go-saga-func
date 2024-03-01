package sagafunc

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
)

func (sg *sagaInteractor) execCompensation(index int) {
	for i := index; i >= 0; i-- {
		logrus.WithFields(logrus.Fields{
			"step":        i,
			"processName": sg.steps[i].ProcessName,
		}).Warn("compensation process")

		err := sg.steps[i].CompensateFunc()
		if err != nil {
			err = fmt.Errorf("step: %d, process_name: %s, err: %v", i, sg.steps[i].ProcessName, err)
			sg.result.CompensateErrors = multierror.Append(sg.result.CompensateErrors, err)
		}
	}

	if sg.result.CompensateErrors != nil {
		logrus.Error(sg.result.CompensateErrors.Error())
	}
}
