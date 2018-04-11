package openClosePrinciple

type Compute interface {
	Compute(int, int) int
}

type AddCompute struct {
}

func (c *AddCompute) Compute(a, b int) int {
	return a + b
}

type MinusCompute struct {
}

func (c *MinusCompute)Minus(a, b int) int {
	return a - b
}