package factory

func createCompute(t string) Computer {
	switch t {
	case "add":
		return &AddCompute{}
	case "minus":
		return &MinusCompute{}
	}
	return nil
}
