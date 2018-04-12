package decorator

import "fmt"

type Shower interface {
	show()
}

type Person struct {
}

func (p *Person) show() {
	fmt.Println("show body")
}

type WearDecorator struct {
	component Shower
}

func (w *WearDecorator) decorate(shower Shower) {
	w.component = shower
}

func (w *WearDecorator) show() {
	fmt.Println("show wear")
	w.component.show()
}

type TShirtDecorator struct {
	component Shower
}

func (t *TShirtDecorator) show() {
	fmt.Println("show tshirt")
	t.component.show()
}

