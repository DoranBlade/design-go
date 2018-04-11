package openClosePrinciple

import "testing"

func TestClient1(t *testing.T) {
	c := &Client1{addComputer: AddComputer{}}
	c.print()
}

func TestClient2(t *testing.T) {
	c1 := &Client2{computer: &AddComputer{}}
	c1.print()
	c2 := &Client2{computer: &MinusComputer{}}
	c2.print()
}

