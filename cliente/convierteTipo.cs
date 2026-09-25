using System;

public class ConvierteTipo {
    
    public string toStringTipos(Tipo t){
	string s = "";
	switch (t) {
	    case Tipo.IDENTIFY: //Identifica al usuario con el servidor
		s = "IDENTIFY";
		break;
	    case Tipo.STATUS: //Cambia el estado de un usuario
		s = "STATUS";
		break;
	    case Tipo.USERS: //Solicita la lista de usuarios del chat
		s = "USERS";
		break;
	    case Tipo.TEXT: //Manda un texto privado a otro usuario
		s = "TEXT";
		break;
	    case Tipo.PUBLIC_TEXT: //Manda un mensaje en el chat grupal
		s = "PUBLIC_TEXT";
		break;
	    case Tipo.NEW_ROOM: //Crea una sala
		s = "NEW_ROOM";
		break;
	    case Tipo.INVITE: //Invita los usuarios a una sala
		s = "INVITE";
		break;
	    case Tipo.JOIN_ROOM: //El usuario se unio a la sala
		s = "JOIN_ROOM";
		break;
	    case Tipo.ROOM_USERS: //Solicita la lista de usuarios en un cuarto
		s = "ROOM_USERS";
		break;
	    case Tipo.ROOM_TEXT: //Manda mensaje dentro de un cuarto
		s = "ROOM_TEXT";
		break;
	    case Tipo.LEAVE_ROOM: //Para abandonar un cuarto
		s = "LEAVE_ROOM";
		break;
	    case Tipo.DISCONNECT://Desconecta al usuario
		s = "DISCONNECT";
		break;
	    case Tipo.INVALID_TIPO: //Cuando la operación no se reconoce
		s = "INVALID";
		break;
	    case Tipo.NEW_USER: //Se unio un nuevo usuario
		s = "NEW_USER";
		break;
	    case Tipo.NEW_STATUS: //Se cambia el estado de un usuario
		s = "NEW_STATUS";
		break;
	    case Tipo.USER_LIST: //Regresa la lista de usuarios del chat
		s = "USER_LIST";
		break;
	    case Tipo.TEXT_FROM: //Cuando otro usuario recibe un mensaje privado
		s = "TEXT_FROM";
		break;
	    case Tipo.PUBLIC_TEXT_FROM: //Algún usuario mando un mensaje publico
		s = "PUBLIC_TEXT_FROM";
		break;
	    case Tipo.INVITATION: //Algún usuario invita a otro a una sala
		s = "INVITATION";
		break;
	    case Tipo.JOINED_ROOM: //Le avisa a los usuarios de la sala que alguien se unio
		s = "JOINED_ROOM";
		break;
	    case Tipo.ROOM_USERS_LIST: //Solicita la lista de usuarios en un cuarto
		s = "ROOM_USERS_LIST";
		break;
	    case Tipo.ROOM_TEXT_FROM: //Se envia un mensaje en especifico a algun cuarto
		s = "ROOM_TEXT_FROM";
		break;
	    case Tipo.LEFT_ROOM: //Cuando un usuario abandona un cuarto
		s = "LEFT_ROOM";
		break;
	    case Tipo.DISCONNECTED: //Operación no reconocida
		s = "DISCONNECTED";
		break;
	    case Tipo.RESPONSE: //Respuesta del servidor
		s = "RESPONSE";
		break;
	}
	return s;
    }
    
    
    public Tipo toTipo(string s){
	Tipo i = 0;
	switch (s) {
	    case "IDENTIFY":
		i = Tipo.IDENTIFY;
		break;
	    case "STATUS":
		i = Tipo.STATUS;
		break;
	    case "USERS":
		i = Tipo.USERS;
		break;
	    case "TEXT":
		i = Tipo.TEXT;
		break;
	    case "PUBLIC_TEXT":
		i = Tipo.PUBLIC_TEXT;
		break;
	    case "NEW_ROOM":
		i = Tipo.NEW_ROOM;
		break;
	    case "INVITE":
		i = Tipo.INVITE;
		break;
	    case "JOIN_ROOM":
		i = Tipo.JOIN_ROOM;
		break;
	    case "ROOM_USERS":
		i = Tipo.ROOM_USERS;
		break;
	    case "ROOM_TEXT":
		i = Tipo.ROOM_TEXT;
		break;
	    case "LEAVE_ROOM":
		i = Tipo.LEAVE_ROOM;
		break;
	    case "DISCONNECT":
		i = Tipo.DISCONNECT;
		break;
	    case "INVALID":
		i = Tipo.INVALID_TIPO;
		break;
	    case "NEW_USER":
		i = Tipo.NEW_USER;
		break;
	    case "NEW_STATUS":
		i = Tipo.NEW_STATUS;
		break;
	    case "USER_LIST":
		i = Tipo.USER_LIST;
		break;
	    case "TEXT_FROM":
		i = Tipo.TEXT_FROM;
		break;
	    case "PUBLIC_TEXT_FROM":
		i = Tipo.PUBLIC_TEXT_FROM;
		break;
	    case "INVITATION":
		i = Tipo.INVITATION;
		break;
	    case "JOINED_ROOM":
		i = Tipo.JOINED_ROOM;
		break;
	    case "ROOM_USERS_LIST":
		i = Tipo.ROOM_USERS_LIST;
		break;
	    case "ROOM_TEXT_FROM":
		i = Tipo.ROOM_TEXT_FROM;
		break;
	    case "LEFT_ROOM":
		i = Tipo.LEFT_ROOM;
		break;
	    case "DISCONNECTED":
		i = Tipo.DISCONNECTED;
		break;
	    case "RESPONSE":
		i = Tipo.RESPONSE;
		break;
	}
	return i;
    }
    
}
