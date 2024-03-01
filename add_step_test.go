package sagafunc_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	saga "github.com/Muruyung/go-saga-func"
)

func TestAddStep(t *testing.T) {
	type args struct {
		step *saga.Step
	}

	var (
		tests = []struct {
			name    string
			args    args
			wantErr bool
		}{
			{
				name: "success",
				args: args{
					step: &saga.Step{
						ProcessName: "test 1",
						ExecutionFunc: func() error {
							return nil
						},
						RollbackFunc: func() error {
							return nil
						},
					},
				},
				wantErr: false,
			},
			{
				name: "failed-empty_rollback_func",
				args: args{
					step: &saga.Step{
						ProcessName: "test 2",
						ExecutionFunc: func() error {
							return nil
						},
					},
				},
				wantErr: true,
			},
			{
				name: "failed-empty_execution_func",
				args: args{
					step: &saga.Step{
						ProcessName: "test 3",
						RollbackFunc: func() error {
							return nil
						},
					},
				},
				wantErr: true,
			},
			{
				name: "failed-empty_step",
				args: args{
					step: nil,
				},
				wantErr: true,
			},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				sagaTest = saga.NewSaga()
				err      = sagaTest.AddStep(tt.args.step)
			)

			if tt.wantErr {
				if err == nil {
					fmt.Printf("%v condition is not expected\n", tt.name)
				}
				assert.NotNil(t, err)
			} else {
				if err != nil {
					fmt.Printf("%v condition is not expected\nerror: %v\n", tt.name, err)
				}
				assert.Nil(t, err)
			}
		})
	}
}
