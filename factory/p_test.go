package factory

import (
	"testing"
	"fmt"
)

func TestSimpleFactory(t *testing.T) {
	computer := createCompute("add")
	fmt.Println(computer.compute(2, 2))
}

func TestFactory(t *testing.T) {
	computerFactory := AddComputerFactory{}
	computer := computerFactory.create()
	fmt.Println(computer.compute(2, 2))
	computerFactory1 := MinusComputerFactory{}
	computer1 := computerFactory1.create()
	fmt.Println(computer1.compute(2 ,2))
}

