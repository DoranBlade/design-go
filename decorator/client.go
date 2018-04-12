package decorator

import "fmt"

type Client struct {
}

func (c *Client) show() {
	fmt.Println("show body")
}

func (c *Client) showWear() {
	fmt.Println("show wear")
}
