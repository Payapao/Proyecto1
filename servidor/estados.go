package main

import(
	//Formato del texto
	"fmt"
	//Pra manejo de errores
	"errors"
	//Protocolo
	"encoding/json"
)

//Enum para go
type Estados int

//Estados validos para los clientes
const (
	ACTIVE Estados = iota + 1
	AWAY //Estados = iota
	BUSY //Puede omitirse el = iota para los siguientes elementos
)

//Define como se va a escribir el tipo Estados
func (e Estados) MarshalJSON()([]byte, error){
	est, err := e.toStringEstados()
	//No debe pasar
	if err != nil {
		return nil, err
	}
	//Le agrega las comillas para el json
	estado := fmt.Sprintf(`"%s"`, est)
	
	return []byte(estado), nil
}

//Define como se va a leer el tipo Estados
func (e *Estados) UnmarshalJSON(b []byte) error {

	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	est, err := toIntEstados(s)
	//No debe pasar
	if err != nil {
		return err
	}
	
	//Asignamos el valor
	*e = est
	
	return nil
}

//Funcion to String para los estados
func (e Estados) toStringEstados()(string, error){
	s := ""
	switch e {
	case ACTIVE:
		s = "ACTIVE"
	case AWAY:
		s = "AWAY"
	case BUSY:
		s = "BUSY"
	default:
		return s , errors.New("El estado no es valido")
	}
	return s, nil
}

func toIntEstados(s string)(Estados, error){
	var e Estados
	switch s {
	case "ACTIVE":
		e = ACTIVE
	case "AWAY":
		e = AWAY
	case "BUSY":
		e = BUSY
	default:
		return e, errors.New("El estado no es valido")
	}
	return e, nil
}
