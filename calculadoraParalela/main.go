package main

import (
	"fmt"
	"time"
)


type Historial struct {
	Bug bool
	Offset int
	Operandos []float64
	Operacion string
	Resultado float64
}

var c1 chan Historial =  make(chan Historial)
var c2 chan Historial = make(chan Historial)
var c3 chan Historial =  make(chan Historial)
var c4 chan Historial =  make(chan Historial)

func main() {

	var num1, num2 float64
	var operandos []float64
	var respuesta string
	var bug bool
	var offset int

	fmt.Printf("Ingrese dos operandos: \n")
	fmt.Scanln(&num1, &num2)
	operandos = []float64{num1, num2}
	fmt.Printf("Desea buguear la calculadora?: (S/N)")
	fmt.Scanln(&respuesta)
	if (respuesta == "S") {
		bug = true
		fmt.Printf("Ingrese OFFSET del Bug:")
		fmt.Scanln(&offset)
	}

	go func() {

		resultado := operandos[0] + operandos[1]
		if (bug) {
			resultado += float64(offset)
		}
		c1 <- Historial{
			Bug:       bug,
			Offset:    offset,
			Operandos: operandos,
			Operacion: "suma",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {
		resultado := operandos[0] - operandos[1]
		if (bug) {
			resultado += float64(offset)
		}
		c2 <- Historial{
			Bug:       bug,
			Offset:    offset,
			Operandos: operandos,
			Operacion: "resta",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)
	}()

	go func() {

		resultado := operandos[0] * operandos[1]
		if (bug) {
			resultado += float64(offset)
		}
		c3 <- Historial{
			Bug:       bug,
			Offset:    offset,
			Operandos: operandos,
			Operacion: "multiplicacion",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {

		resultado := operandos[0] / operandos[1]
		if (bug) {
			resultado += float64(offset)
		}
		c4<- Historial{
			Bug:       bug,
			Offset:    offset,
			Operandos: operandos,
			Operacion: "Division",
			Resultado: resultado,
		}
		time.Sleep(time.Second * 2)

	}()

	go func() {
		for {
			select {
			case suma := <-c1:
				fmt.Println(suma)
			case resta := <-c2:
				fmt.Println(resta)
			case multiplicar := <-c3:
				fmt.Println(multiplicar)
			case dividir := <-c4:
				fmt.Println(dividir)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}