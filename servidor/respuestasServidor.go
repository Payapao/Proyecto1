package main

import(
	//Manejo de errores
	"errors"
	//Formato del texto
	"fmt"
	//Protocolo
	"encoding/json"
)

type Respuesta int

const(
	//RESPONSE
	SUCCESS Respuesta = iota + 1//La operación se ejecuto exitosamente
	USER_ALREADY_EXISTS //El nombre se usuario ya exite
	NO_SUCH_USER //El usuario no existe
	ROOM_ALREADY_EXISTS //El nombre de la sala ya existe
	NO_SUCH_ROOM //La sala no exite
	NOT_INVITED //El usuario no ha sido invitado a la sala
	NOT_JOINED //El usuario no se ha unido a la sala
	NOT_IDENTIFIED //Usuario no identificado
)

//Define como se va a escribir el tipo Respuesta
func (r Respuesta) MarshalJSON()([]byte, error){
	resp, err := r.toStringRespuestas()
	//No debe pasar
	if err != nil {
		return nil, err
	}
	//Le agrega las comillas para el json
	respuesta := fmt.Sprintf(`"%s"`, resp)
	
	return []byte(respuesta), nil
}

//Define como se va a leer el tipo Respuesta
func (r *Respuesta) UnmarshalJSON(b []byte) error {
	//Le quitamos las comillas
	var s string
	err := json.Unmarshal(b, &s)
	//Pasa si no es una string o es un json invalido
	if err != nil {
		return err
	}
	//Buscamos la string y lo pasamos a un elemento de nuestra iota
	resp, err := toIntRespuestas(s)
	//Elemento invalido
	if err != nil {
		return err
	}
	
	//Asignamos el valor
	*r = resp
	return nil
}

//Funcion to String para las respuestas
func (r Respuesta) toStringRespuestas()(string, error){
	s := ""
	switch r {
	case SUCCESS: //La operación se ejecuto exitosamente
		s = "SUCCESS"
	case USER_ALREADY_EXISTS: //El nombre se usuario ya exite
		s = "USER_ALREADY_EXISTS"
	case NO_SUCH_USER: //El usuario no existe
		s = "NO_SUCH_USER"
	case ROOM_ALREADY_EXISTS: //El nombre de la sala ya existe
		s = "ROOM_ALREADY_EXISTS"
	case NO_SUCH_ROOM: //La sala no exite
		s = "NO_SUCH_ROOM"
	case NOT_INVITED: //El usuario no ha sido invitado a la sala
		s = "NOT_INVITED"
	case NOT_JOINED: //El usuario no se ha unido a la sala
		s = "NOT_JOINED"
	case NOT_IDENTIFIED: //Usuario no identificado
		s = "NOT_IDENTIFIED"
	default:
		return s , errors.New("La respuesta es invalida")
	}
	return s, nil
}

//Funcion to Int para las respuestas
func toIntRespuestas(s string)(Respuesta, error){
	var est Respuesta
	switch s {
	case "SUCCESS":
		est = SUCCESS
	case "USER_ALREADY_EXISTS":
		est = USER_ALREADY_EXISTS
	case "NO_SUCH_USER":
		est = NO_SUCH_USER
	case "ROOM_ALREADY_EXISTS":
		est = ROOM_ALREADY_EXISTS
	case "NO_SUCH_ROOM":
		est = NO_SUCH_ROOM
	case "NOT_INVITED":
		est = NOT_INVITED
	case "NOT_JOINED":
		est = NOT_JOINED
	case "NOT_IDENTIFIED":
		est = NOT_IDENTIFIED
	default:
		return est , errors.New("La respuesta es invalida")
	}
	return est, nil
}
