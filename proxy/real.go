package proxy

import "fmt"

type Shower interface {
	show()
}

type Real struct {
	name string
}

func (r *Real) show()  {
	fmt.Println("real show")
	fmt.Println("name: " + r.name)
}


