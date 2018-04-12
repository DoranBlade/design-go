package factory

type Computer interface {
	compute(a, b int) int
}

type AddCompute struct {
}

func (add *AddCompute) compute(a, b int) int {
	return a + b
}

type MinusCompute struct {
}

func (minus *MinusCompute) compute(a, b int) int {
	return a - b
}

type ComputerFactory interface {
	create() Computer
}

type AddComputerFactory struct {
}

func (add *AddComputerFactory) create() Computer {
	return &AddCompute{}
}

type MinusComputerFactory struct {
}

func (minus *MinusComputerFactory) create() Computer {
	return &MinusCompute{}
}
