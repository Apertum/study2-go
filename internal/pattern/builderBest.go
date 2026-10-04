package pattern

import "fmt"

// WithMode — пример функции, которая присваивает поле Mode.
func WithMode(mode int) func(*Object) {
	return func(o *Object) {
		o.Mode = mode
	}
}

// WithPath — пример функции, которая присваивает поле Path.
func WithPath(path string) func(*Object) {
	return func(o *Object) {
		o.Path = path
	}
}

// NewObject — функция-конструктор объекта.
func NewObjectBest(opts ...func(*Object)) *Object {
	o := &Object{}

	// вызываем все указанные функции для установки параметров
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func main() {
	o := NewObjectBest(WithMode(10), WithPath(`root`))
	fmt.Println(o)
}
