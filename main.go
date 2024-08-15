package main

import "fmt"

func main() {
	res := Sumar(1, 2)

	fmt.Println("la suma es ", res)
}

func Sumar(a, b int) int {
	return a + b
}

func Dividir(a, b int) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por 0")
	}

	return float64(a)/float64(b), nil
}