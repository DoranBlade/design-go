package openClosePrinciple

import "fmt"

type Client2 struct {
	compute Compute
}

func (c *Client2) print() {
	fmt.Println(c.compute.Compute(2, 2))
}
