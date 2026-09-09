package main

type Operacion int

const{
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
	INVALID //Cuando la operación no se reconoce
}
