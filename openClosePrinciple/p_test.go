package openClosePrinciple

import "testing"

func TestClient1(t *testing.T) {
	c := &Client1{}
	c.print()
}

func TestClient2(t *testing.T) {
	c := &Client2{compute: &AddCompute{}}
	c.print()
}


