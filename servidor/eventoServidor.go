package main

import(
	//Errores
	"errors"
	//Formato del texto
	"fmt"
	//Protocolo
	"encoding/json"
)

type Tipo int

const(
	//Del usuario
	IDENTIFY Tipo = iota + 1 //Identifica al usuario con el servidor
	STATUS //Cambia el estado de un usuario
	USERS //Solicita la lista de usuarios del chat
	TEXT //Manda un texto privado a otro usuario
	PUBLIC_TEXT //Manda un mensaje en el chat grupal
	NEW_ROOM //Crea una sala
	INVITE //Invita los usuarios a una sala
	JOIN_ROOM //El usuario se unio a la sala
	ROOM_USERS //Solicita la lista de usuarios en un cuarto
	ROOM_TEXT //Manda mensaje dentro de un cuarto
	LEAVE_ROOM //Para abandonar un cuarto
	DISCONNECT //Desconecta al usuario
	INVALID_TIPO //Operación no reconocida

	//Del servidor
	NEW_USER //Se envia a todos los clientes cuando se unio un nuevo usuario
	NEW_STATUS //Se envia a todos los clientes cuado se cambia el estado de un usuario
	USER_LIST //Regresa la lista de usuarios del chat
	TEXT_FROM //Cuando otro usuario recibe un mensaje privado
	PUBLIC_TEXT_FROM //Algún usuario mando un mensaje publico
	INVITATION //Algún usuario invita a otro a una sala
	JOINED_ROOM //Le avisa a los usuarios de la sala que alguien se unio
	ROOM_USERS_LIST //Regresa la lista de usuarios en un cuarto
	ROOM_TEXT_FROM //Se envia un mensaje en especifico a algun cuarto
	LEFT_ROOM //Cuando un usuario abandona un cuarto
	DISCONNECTED //Avisa que un usuario se desconecto
	RESPONSE //Respuesta del servidor
)

//Define como se va a escribir el tipo Tipo
func (t Tipo) MarshalJSON()([]byte, error){
	tip, err := t.toStringTipos()
	//No debe pasar
	if err != nil {
		return nil, err
	}
	//Le agrega las comillas para el json
	tipo := fmt.Sprintf(`"%s"`, tip)
	
	return []byte(tipo), nil
}

//Define como se va a leer el tipo Tipo
func (t *Tipo) UnmarshalJSON(b []byte) error {
	//Le quitamos las comillas
	var s string
	err := json.Unmarshal(b, &s)
	//Pasa si no es una string o es un json invalido
	if err != nil {
		return err
	}
	//Buscamos la string y lo pasamos a un elemento de nuestra iota
	tip, err := toIntTipos(s)
	//Elemento invalido
	if err != nil {
		return err
	}
	
	//Asignamos el valor
	*t = tip
	
	return nil
}

//Funcion to String para los tipos
func (t Tipo) toStringTipos()(string, error){
	s := ""
	switch t {
	case IDENTIFY: //Identifica al usuario con el servidor
		s = "IDENTIFY"
	case STATUS: //Cambia el estado de un usuario
		s = "STATUS"
	case USERS: //Solicita la lista de usuarios del chat
		s = "USERS"
	case TEXT: //Manda un texto privado a otro usuario
		s = "TEXT"
	case PUBLIC_TEXT: //Manda un mensaje en el chat grupal
		s = "PUBLIC_TEXT"
	case NEW_ROOM: //Crea una sala
		s = "NEW_ROOM"
	case INVITE: //Invita los usuarios a una sala
		s = "INVITE"
	case JOIN_ROOM: //El usuario se unio a la sala
		s = "JOIN_ROOM"
	case ROOM_USERS: //Solicita la lista de usuarios en un cuarto
		s = "ROOM_USERS"
	case ROOM_TEXT: //Manda mensaje dentro de un cuarto
		s = "ROOM_TEXT"
	case LEAVE_ROOM: //Para abandonar un cuarto
		s = "LEAVE_ROOM"
	case DISCONNECT://Desconecta al usuario
		s = "DISCONNECT"
	case INVALID_TIPO: //Cuando la operación no se reconoce
		s = "INVALID"
	case NEW_USER: //Se unio un nuevo usuario
		s = "NEW_USER"
	case NEW_STATUS: //Se cambia el estado de un usuario
		s = "NEW_STATUS"
	case USER_LIST: //Regresa la lista de usuarios del chat
		s = "USER_LIST"
	case TEXT_FROM: //Cuando otro usuario recibe un mensaje privado
		s = "TEXT_FROM"
	case PUBLIC_TEXT_FROM: //Algún usuario mando un mensaje publico
		s = "PUBLIC_TEXT_FROM"
	case INVITATION: //Algún usuario invita a otro a una sala
		s = "INVITATION"
	case JOINED_ROOM: //Le avisa a los usuarios de la sala que alguien se unio
		s = "JOINED_ROOM"
	case ROOM_USERS_LIST: //Solicita la lista de usuarios en un cuarto
		s = "ROOM_USERS_LIST"
	case ROOM_TEXT_FROM: //Se envia un mensaje en especifico a algun cuarto
		s = "ROOM_TEXT_FROM"
	case LEFT_ROOM: //Cuando un usuario abandona un cuarto
		s = "LEFT_ROOM"
	case DISCONNECTED: //Operación no reconocida
		s = "DISCONNECTED"
	case RESPONSE: //Respuesta del servidor
		s = "RESPONSE" 
	default:
		return s, errors.New("El tipo es invalido")
	}
	return s, nil
}

//Funcion toInt para los tipos
func toIntTipos(s string)(Tipo, error){
	var i Tipo
	switch s {
	case "IDENTIFY":
		i = IDENTIFY
	case "STATUS":
		i = STATUS
	case "USERS":
		i = USERS
	case "TEXT":
		i = TEXT
	case "PUBLIC_TEXT":
		i = PUBLIC_TEXT
	case "NEW_ROOM":
		i = NEW_ROOM
	case "INVITE":
		i = INVITE
	case "JOIN_ROOM":
		i = JOIN_ROOM
	case "ROOM_USERS":
		i = ROOM_USERS
	case "ROOM_TEXT":
		i = ROOM_TEXT
	case "LEAVE_ROOM":
		i = LEAVE_ROOM
	case "DISCONNECT":
		i = DISCONNECT
	case "INVALID":
		i = INVALID_TIPO
	case "NEW_USER":
		i = NEW_USER
	case "NEW_STATUS":
		i = NEW_STATUS
	case "USER_LIST":
		i = USER_LIST
	case "TEXT_FROM":
		i = TEXT_FROM
	case "PUBLIC_TEXT_FROM":
		i = PUBLIC_TEXT_FROM
	case "INVITATION":
		i = INVITATION
	case "JOINED_ROOM":
		i = JOINED_ROOM
	case "ROOM_USERS_LIST":
		i = ROOM_USERS_LIST
	case "ROOM_TEXT_FROM":
		i = ROOM_TEXT_FROM
	case "LEFT_ROOM":
		i = LEFT_ROOM
	case "DISCONNECTED":
		i = DISCONNECTED
	case "RESPONSE":
		i = RESPONSE
	default:
		return i, errors.New("El tipo es invalido")
	}
	return i, nil
}
