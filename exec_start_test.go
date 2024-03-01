package sagafunc_test

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	saga "github.com/Muruyung/go-saga-func"
)

func TestExecStart(t *testing.T) {
	type args struct {
		skipCompensateError bool
		execErr             error
		compErr             error
	}

	var (
		defaultErr = errors.New("")
		tests      = []struct {
			name    string
			args    args
			wantErr bool
			want    int
		}{
			{
				name: "success",
				args: args{
					skipCompensateError: false,
					execErr:             nil,
					compErr:             nil,
				},
				want:    20,
				wantErr: false,
			},
			{
				name: "failed-error_exec",
				args: args{
					skipCompensateError: false,
					execErr:             defaultErr,
					compErr:             nil,
				},
				want:    5,
				wantErr: true,
			},
			{
				name: "failed-error_exec_and_skip_comp_error",
				args: args{
					skipCompensateError: true,
					execErr:             defaultErr,
					compErr:             nil,
				},
				want:    15,
				wantErr: true,
			},
			{
				name: "failed-error_exec_and_comp_error",
				args: args{
					skipCompensateError: false,
					execErr:             defaultErr,
					compErr:             defaultErr,
				},
				want:    5,
				wantErr: true,
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				sagaTest = saga.NewSaga(tt.args.skipCompensateError)
				value    = 5
				err      error
				mu       sync.Mutex
			)

			err = sagaTest.AddStep(&saga.Step{
				ProcessName: "step 1",
				ExecutionFunc: func() error {
					value += 5
					return nil
				},
				CompensateFunc: func() error {
					value -= 5
					return nil
				},
			})
			assert.Nil(t, err)

			err = sagaTest.AddStep(&saga.Step{
				ProcessName: "step 2",
				ExecutionFunc: func() error {
					value += 10
					return tt.args.execErr
				},
				CompensateFunc: func() error {
					value -= 10
					return tt.args.compErr
				},
			})
			assert.Nil(t, err)

			go func() {
				err = sagaTest.ExecStart()
			}()

			time.Sleep(50 * time.Millisecond)

			mu.Lock()
			if tt.wantErr {
				if err == nil {
					fmt.Printf("%v condition is not expected\n", tt.name)
				}
				assert.NotNil(t, err)
				assert.Equal(t, tt.want, value)
			} else {
				if err != nil {
					fmt.Printf("%v condition is not expected\nerror: %v\n", tt.name, err)
				}
				assert.Nil(t, err)
				assert.Equal(t, tt.want, value)
			}
			mu.Unlock()
		})
	}
}
