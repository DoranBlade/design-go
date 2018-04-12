package proxy

import "fmt"

type Client struct {
	name string
}

func (c *Client) show() {
	fmt.Println("show client")
}

func (c *Client) newShow() {
	fmt.Println("newShow do something")
	c.show()
}
