package decorator

import "testing"

func TestClient(t *testing.T) {
	c := Client{}
	c.show()
	c.showWear()
}

func TestDecorator(t *testing.T) {
	person := Person{}
	wear := WearDecorator{&person}
	tshirt := TShirtDecorator{&wear}
	tshirt.show()
}

