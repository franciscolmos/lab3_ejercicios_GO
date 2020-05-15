package operaciones

type Historial struct {
	Bug bool
	Offset int
	Operandos []float64
	Operacion string
	Resultado float64
}

func Sumar(bug bool, offset int, operandos ...float64) (Historial) {
	resultado :=operandos[0] + operandos[1]
	if(bug){
		resultado += float64(offset)
	}
	return  Historial{
		Bug: bug,
		Offset: offset,
		Operandos: operandos,
		Operacion: "suma",
		Resultado: resultado,
	}

}

func Restar(bug bool,offset int,operandos ...float64) (Historial) {
	resultado := operandos[0] - operandos[1]
	if(bug){
		resultado += float64(offset)
	}
	return  Historial{
		Operandos: operandos,
		Operacion: "resta",
		Resultado: resultado,
	}

}

func Multiplicar(bug bool,offset int,operandos ...float64) (Historial) {
	resultado := operandos[0] * operandos[1]
	if(bug){
		resultado += float64(offset)
	}
	return  Historial{
		Operandos: operandos,
		Operacion: "multiplicacion",
		Resultado: resultado,
	}

}

func Division(bug bool,offset int,operandos ...float64) (Historial) {
	resultado := operandos[0] / operandos[1]
	if(bug){
		resultado += float64(offset)
	}
	return  Historial{
		Operandos: operandos,
		Operacion: "Division",
		Resultado: resultado,
	}

}