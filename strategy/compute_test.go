package strategy

import (
	"testing"
	"fmt"
)

func TestStrategy(t *testing.T) {
	computeContext := CreateComputeContext("add")
	fmt.Println(computeContext.Result(2, 2))
}

