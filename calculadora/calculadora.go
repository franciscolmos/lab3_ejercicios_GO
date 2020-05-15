package main

import (
	"fmt"
	"calculadora/operaciones"
)

func main()  {
	var num1, num2 float64
	var operacion string
	var operandos []float64
	var respuesta string
	var bug bool
	var offset int
	fmt.Printf("Ingrese la operación a realizar: \n")
	fmt.Scanln(&operacion)
	fmt.Printf("Ingrese dos operandos: \n")
	fmt.Scanln(&num1, &num2)
	operandos = []float64{num1,num2}
	fmt.Printf("Desea buguear la calculadora?: (S/N)")
	fmt.Scanln(&respuesta)
	if(respuesta == "S"){
		bug = true
		fmt.Printf("Ingrese OFFSET del Bug:")
		fmt.Scanln(&offset)
	}
	switch operacion {
	case "sumar":
		fmt.Print(operaciones.Sumar(bug,offset,operandos...))
	case "restar":
		fmt.Print(operaciones.Restar(bug,offset,operandos...))
	case "multiplicar":
		fmt.Print(operaciones.Multiplicar(bug,offset,operandos...))
	case "dividir":
		fmt.Print(operaciones.Division(bug,offset,operandos...))
	default:
		fmt.Print("Operracion inválida.")
	}






}
