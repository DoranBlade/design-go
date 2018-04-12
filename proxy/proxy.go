package proxy

import "fmt"

type Proxy struct {
	real Real
}

func CreateProxy(name string) Proxy {
	return Proxy{Real{name}}
}

func (p *Proxy)show()  {
	fmt.Println("proxy do something")
	p.real.show()
}
