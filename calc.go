package calculadora

type calculadora struct {}

func NewCalculadora() *calculadora {
	return &calculadora{}
}

func (c *calculadora) Sum(num1,num2 int) int {
	return num1 + num2
}

func (c *calculadora) Sub(num1,num2 int) int {
	return num1 - num2
}