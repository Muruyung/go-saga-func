package sagafunc

type Saga interface {
	AddStep(step *Step) error
	ExecStart() error
}
