package main

import(
	//Manejo de errores
	"errors"
	//Formato del texto
	"fmt"
)

type Respuesta int

const(
	//RESPONSE
	SUCCESS Respuesta = iota //La operación se ejecuto exitosamente
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

//Funcion to String para las respuestas
func (r Respuesta) toStringRespuestas()(string, error){
	s := ""
	switch r {
	case 0: //La operación se ejecuto exitosamente
		s = "SUCCESS"
	case 1: //El nombre se usuario ya exite
		s = "USER_ALREADY_EXISTS"
	case 2: //El usuario no existe
		s = "NO_SUCH_USER"
	case 3: //El nombre de la sala ya existe
		s = "ROOM_ALREADY_EXISTS"
	case 4: //La sala no exite
		s = "NO_SUCH_ROOM"
	case 5: //El usuario no ha sido invitado a la sala
		s = "NOT_INVITED"
	case 6: //El usuario no se ha unido a la sala
		s = "NOT_JOINED"
	case 7: //Usuario no identificado
		s = "NOT_IDENTIFIED"
	default:
		return s , errors.New("El tipo es invalido")
	}
	return s, nil
}
