package openClosePrinciple

type Computer interface {
	Compute(int, int) int
}

type AddComputer struct {
}

func (c *AddComputer) Compute(a, b int) int {
	return a + b
}

type MinusComputer struct {
}

func (c *MinusComputer)Compute(a, b int) int {
	return a - b
}