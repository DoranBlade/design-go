package strategy

type ComputeContext struct {
	compute Compute
}

func (context *ComputeContext)Result(a, b int) int {
	return context.compute.compute(a, b)
}

func CreateComputeContext(t string) *ComputeContext {
	switch t {
	case "add":
		return &ComputeContext{compute: &Add{}}
	case "minus":
		return &ComputeContext{compute: &Minus{}}
	}
	return nil
}
