package main

import "testing"

func TestSumar(t *testing.T) {
	//Arrange
	var a, b int
	a = 1
	b = 2
	want := 3

	//Act
	res := Sumar(a, b)
	
	//Assert
	assertSuma(res, want, t)
	
}

func assertSuma(res, want int, t *testing.T) {
	if res != want {
		t.Error("no devolvio el valor esperado")
	} 
}


func TestDividirFallaDivicionPorCero(t *testing.T) {
	//Arrange
	var a, b int
	a = 1
	b = 0

	//Act
	_, err := Dividir(a, b)
	
	//assert
	if err.Error() != "no se puede dividir por 0" {
		t.Errorf("error inesperado, se obtubo %s", err.Error())
	}
}

func TestDividirFunciona(t *testing.T) {
	//Arrange
	var a, b int
	a = 2
	b = 2
	var want float64 = 1

	//Act
	res, _ := Dividir(a, b)
	
	//Assert
	assertDivicion(res, want, t)
	
}


func assertDivicion(res float64, want float64, t *testing.T) {
	if res != want {
		t.Error("no devolvio el valor esperado")

	} 
}