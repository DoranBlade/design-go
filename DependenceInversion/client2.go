package openClosePrinciple

import "fmt"

type Client2 struct {
	computer Computer
}

func (c *Client2)print() {
	fmt.Println(c.computer.Compute(2, 2))
}
