package strategy

type Compute interface {
	compute(int ,int) int
}

type Add struct {
}

func (add *Add)compute(a, b int) int {
	return a + b
}

type Minus struct {
}

func (minus *Minus) compute(a, b int) int {
	return a - b
}
