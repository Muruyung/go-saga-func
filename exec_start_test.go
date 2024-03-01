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
		skipRollbackError bool
		execErr           error
		rollbackErr       error
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
					skipRollbackError: false,
					execErr:           nil,
					rollbackErr:       nil,
				},
				want:    20,
				wantErr: false,
			},
			{
				name: "failed-error_exec",
				args: args{
					skipRollbackError: false,
					execErr:           defaultErr,
					rollbackErr:       nil,
				},
				want:    5,
				wantErr: true,
			},
			{
				name: "failed-error_exec_and_skip_rollback_error",
				args: args{
					skipRollbackError: true,
					execErr:           defaultErr,
					rollbackErr:       nil,
				},
				want:    15,
				wantErr: true,
			},
			{
				name: "failed-error_exec_and_rollback_error",
				args: args{
					skipRollbackError: false,
					execErr:           defaultErr,
					rollbackErr:       defaultErr,
				},
				want:    5,
				wantErr: true,
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				sagaTest = saga.NewSaga(tt.args.skipRollbackError)
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
				RollbackFunc: func() error {
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
				RollbackFunc: func() error {
					value -= 10
					return tt.args.rollbackErr
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
