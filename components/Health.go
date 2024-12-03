package components

type Health struct {
	HP int
}

func (c Health) ID() int {

	return HealthComponentId
}

func NewHealth() *Health {
	return &Health{}
}
