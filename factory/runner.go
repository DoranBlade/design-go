package factory

import "fmt"

type Runner interface {
	run()
}

type Dog struct {
}

func (dog *Dog)run()  {
	fmt.Println("dog is run")
}

type Cat struct {
}

func (cat *Cat)run()  {
	fmt.Println("cat is run")
}
