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

	
	NEW_USER //Se envia a todos los clientes cuando se unio un nuevo usuario
	NEW_STATUS //Se envia a todos los clientes cuado se cambia el estado de un usuario
	USER_LIST //Regresa la lista de usuarios del chat
	NEW_USER //Se envia a todos los clientes cuando se unio un nuevo usuario
	TEXT_FROM //Cuando otro usuario recibe un mensaje privado
	PUBLIC_TEXT_FROM //Algún usuario mando un mensaje publico
	INVITATION //Algún usuario invita a otro a una sala
	JOINED_ROOM //Le avisa a los usuarios de la sala que alguien se unio
	ROOM_USER_LIST //Regresa la lista de usuarios del cuarto
	ROOM_TEXT_FROM //Se envia un mensaje en especifico a algun cuarto
	LEFT_ROOM //Cuando un usuario abandona un cuarto
	DISCONNECTED //Avisa que un usuario se desconecto
	
}
