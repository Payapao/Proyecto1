package main

import(
	//Pra manejo de errores
	"errors"
)

//Enum para go
type Estados int

//Estados validos para los clientes
const (
	ACTIVE Estados = iota
	AWAY //Estados = iota
	BUSY //Puede omitirse el = iota para los siguientes elementos
)

//Funcion to String para los estados
func toStringEstados(e Estados)(string, error){
	s := ""
	switch e {
	case 0:
		s = "ACTIVE"
	case 1:
		s = "AWAY"
	case 2:
		s = "BUSY"
	default:
		return s , errors.New("El estado no es valido")
	}
	return s, nil
}
