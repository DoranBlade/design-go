package openClosePrinciple

import "fmt"

type Client1 struct {
}

func (c *Client1)print() {
	fmt.Println(c.add(2, 2))
}

func (c *Client1)add(a, b int) int {
	return a + b
}
