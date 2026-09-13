package main

import(
	"errors"
)

type Tipo int

const(
	//Del usuario
	IDENTIFY Tipo = iota //Identifica al usuario con el servidor
	STATUS //Cambia el estado de un usuario
	USERS //Solicita la lista de usuarios del chat
	TEXT //Manda un texto privado a otro usuario
	PUBLIC_TEXT //Manda un mensaje en el chat grupal
	NEW_ROOM //Crea una sala
	INVITE //Invita los usuarios a una sala
	JOIN_ROOM //El usuario se unio a la sala
	ROOM_USER_LIST //Solicita la lista de usuarios en un cuarto
	ROOM_TEXT //Manda mensaje dentro de un cuarto
	LEAVE_ROOM //Para abandonar un cuarto
	DISCONNECT //Desconecta al usuario

	//Del servidor
	NEW_USER //Se envia a todos los clientes cuando se unio un nuevo usuario
	NEW_STATUS //Se envia a todos los clientes cuado se cambia el estado de un usuario
	USER_LIST //Regresa la lista de usuarios del chat
	TEXT_FROM //Cuando otro usuario recibe un mensaje privado
	PUBLIC_TEXT_FROM //Algún usuario mando un mensaje publico
	INVITATION //Algún usuario invita a otro a una sala
	JOINED_ROOM //Le avisa a los usuarios de la sala que alguien se unio
	//ROOM_USER_LIST //Regresa la lista de usuarios del cuarto
	ROOM_TEXT_FROM //Se envia un mensaje en especifico a algun cuarto
	LEFT_ROOM //Cuando un usuario abandona un cuarto
	DISCONNECTED //Avisa que un usuario se desconecto
)

//Funcion to String para los tipos
func toStringTipos(t Tipo)(string, error){
	s := ""
	switch t {
	case 0: //Identifica al usuario con el servidor
		s = "IDENTIFY"
	case 1: //Cambia el estado de un usuario
		s = "STATUS"
	case 2: //Solicita la lista de usuarios del chat
		s = "USERS"
	case 3: //Manda un texto privado a otro usuario
		s = "TEXT"
	case 4: //Manda un mensaje en el chat grupal
		s = "PUBLIC_TEXT"
	case 5: //Crea una sala
		s = "NEW_ROOM"
	case 6: //Invita los usuarios a una sala
		s = "INVITE"
	case 7: //El usuario se unio a la sala
		s = "JOIN_ROOM"
	case 8: //Solicita la lista de usuarios en un cuarto
		s = "ROOM_USER_LIST"
	case 9: //Manda mensaje dentro de un cuarto
		s = "ROOM_TEXT"
	case 10: //Para abandonar un cuarto
		s = "LEAVE_ROOM"
	case 11://Desconecta al usuario
		s = "DISCONNECT"
	case 12: //Cuando la operación no se reconoce
		s = "INVALID"
	case 13: //Se unio un nuevo usuario
		s = "NEW_USER"
	case 14: //Se cambia el estado de un usuario
		s = "NEW_STATUS"
	case 15: //Regresa la lista de usuarios del chat
		s = "USER_LIST"
	case 16: //Cuando otro usuario recibe un mensaje privado
		s = "TEXT_FROM"
	case 17: //Algún usuario mando un mensaje publico
		s = "PUBLIC_TEXT_FROM"
	case 18: //Algún usuario invita a otro a una sala
		s = "INVITATION"
	case 19: //Le avisa a los usuarios de la sala que alguien se unio
		s = "JOINED_ROOM"
	case 20: //Se envia un mensaje en especifico a algun cuarto
		s = "ROOM_TEXT_FROM"
	case 21: //Cuando un usuario abandona un cuarto
		s = "LEFT_ROOM"
	case 22: //Avisa que un usuario se desconecto un usuario
		s = "DISCONNECTED"
	default:
		return s, errors.New("El tipo es invalido")
	}
	return s, nil
}
