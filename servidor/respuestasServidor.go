package main

type Respuesta int

const{
	//RESPONSE
	SUCCESS Operacion = iota //La operación se ejecuto exitosamente
	USER_ALREADY_EXISTS //El nombre se usuario ya exite
	NO_SUCH_USER //El usuario no existe
	ROOM_ALREADY_EXISTS //El nombre de la sala ya existe
	NO_SUCH_ROOM //La sala no exite
	NOT_INVITED //El usuario no ha sido invitado a la sala
	NOT_JOINED //El usuario no se ha unido a la sala
	NOT_IDENTIFIED //Usuario no identificado
	INVALID //Operación no reconocida
}

//Funcion to String para las respuestas
func toString(r Respuesta)(string, error){
	s := ""
	switch r {
	case "SUCCESS": //La operación se ejecuto exitosamente
		r = "SUCCESS"
	case "USER_ALREADY_EXISTS": //El nombre se usuario ya exite
		r = "USER_ALREADY_EXISTS"
	case "NO_SUCH_USER": //El usuario no existe
		r = "NO_SUCH_USER"
	case "ROOM_ALREADY_EXISTS": //El nombre de la sala ya existe
		r = "ROOM_ALREADY_EXISTS"
	case "NO_SUCH_ROOM": //La sala no exite
		r = "NO_SUCH_ROOM"
	case "NOT_INVITED": //El usuario no ha sido invitado a la sala
		r = "NOT_INVITED"
	case "NOT_JOINED": //El usuario no se ha unido a la sala
		r = "NOT_JOINED"
	case "NOT_IDENTIFIED": //Usuario no identificado
		r = "NOT_IDENTIFIED"
	case "INVALID": //Operación no reconocida
		r = "INVALID"
	default:
		return s , error.New("El tipo es invalido")
	}
	return s, nil
}
