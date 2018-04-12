package simpleFactory

func CreateRunner(t string) Runner {
	switch t {
	case "dog":
		return &Dog{}
	case "cat":
		return &Cat{}
	}
	return nil
}
