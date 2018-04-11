package openClosePrinciple

import "fmt"

type Client1 struct {
	addComputer AddComputer
}

func (c *Client1)print() {
	fmt.Println(c.addComputer.Compute(2, 2))
}
