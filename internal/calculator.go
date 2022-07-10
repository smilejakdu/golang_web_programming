package internal

import "fmt"

type Calculator struct {
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c Calculator) Add(num1, num2 int) int {
	fmt.Println(num1 + num2)
	panic("implement me")
}
