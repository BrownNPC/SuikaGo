package components

type Input struct {
	Left   bool
	Right  bool
	Action bool
}

func NewInput() *Input {
	return &Input{}
}

func (c Input) ID() int {
	return InputComponentId
}
