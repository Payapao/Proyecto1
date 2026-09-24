using System;

public enum Tipo {
    IDENTIFY,//Identifica al usuario con el servidor
    STATUS, //Cambia el estado del usuario
    USERS, //Solicita la lista de usuarios del chat
    TEXT, //Manda un texto privado a otro usuario
    PUBLIC_TEXT, //Manda un mensaje en el chat grupal
    NEW_ROOM, //Crea una sala
    INVITE, //Invita los usuarios a una sala
    JOIN_ROOM, //Unirse a la sala
    ROOM_USERS, //Solicita la lista de usuarios en una sala
    ROOM_TEXT, //Manda mensaje dentro de una sala
    LEAVE_ROOM, //Para abandonar un cuarto
    DISCONNECT, //Desconecta al usuario
    INVALID_TIPO, //Operación no reconocida

    NEW_USER, //Se unio un nuevo usuario
    NEW_STATUS, //Se cambia el estado de un usuario
    USER_LIST, //Regresa la lista de usuarios del chat
    TEXT_FROM, //Recibe un mensaje privado
    PUBLIC_TEXT_FROM, //Algún usuario mando un mensaje publico
    INVITATION, //Invitación a una sala
    JOINED_ROOM, //Alguien se unio a una sala
    ROOM_USERS_LIST, //Regresa la lista de usuarios de una sala
    ROOM_TEXT_FROM, //Se envió un mensaje a una sala
    LEFT_ROOM, //Un usuario abandonó una sala
    DISCONNECTED, //Avisa que un usuario se desconecto
    RESPONSE //Respuesta del servidor
}

public string toStringTipos(Tipo t){
    string s = "";
    switch (t) {
	case IDENTIFY: //Identifica al usuario con el servidor
	    s = "IDENTIFY";
	    break;
	case STATUS: //Cambia el estado de un usuario
	    s = "STATUS";
	    break;
	case USERS: //Solicita la lista de usuarios del chat
	    s = "USERS";
	    break;
	case TEXT: //Manda un texto privado a otro usuario
	    s = "TEXT";
	    break;
	case PUBLIC_TEXT: //Manda un mensaje en el chat grupal
	    s = "PUBLIC_TEXT";
	    break;
	case NEW_ROOM: //Crea una sala
	    s = "NEW_ROOM";
	    break;
	case INVITE: //Invita los usuarios a una sala
	    s = "INVITE";
	    break;
	case JOIN_ROOM: //El usuario se unio a la sala
	    s = "JOIN_ROOM";
	    break;
	case ROOM_USERS: //Solicita la lista de usuarios en un cuarto
	    s = "ROOM_USERS";
	    break;
	case ROOM_TEXT: //Manda mensaje dentro de un cuarto
	    s = "ROOM_TEXT";
	    break;
	case LEAVE_ROOM: //Para abandonar un cuarto
	    s = "LEAVE_ROOM";
	    break;
	case DISCONNECT://Desconecta al usuario
	    s = "DISCONNECT";
	    break;
	case INVALID_TIPO: //Cuando la operación no se reconoce
	    s = "INVALID";
	    break;
	case NEW_USER: //Se unio un nuevo usuario
	    s = "NEW_USER";
	    break;
	case NEW_STATUS: //Se cambia el estado de un usuario
	    s = "NEW_STATUS";
	    break;
	case USER_LIST: //Regresa la lista de usuarios del chat
	    s = "USER_LIST";
	    break;
	case TEXT_FROM: //Cuando otro usuario recibe un mensaje privado
	    s = "TEXT_FROM";
	    break;
	case PUBLIC_TEXT_FROM: //Algún usuario mando un mensaje publico
	    s = "PUBLIC_TEXT_FROM";
	    break;
	case INVITATION: //Algún usuario invita a otro a una sala
	    s = "INVITATION";
	    break;
	case JOINED_ROOM: //Le avisa a los usuarios de la sala que alguien se unio
	    s = "JOINED_ROOM";
	    break;
	case ROOM_USERS_LIST: //Solicita la lista de usuarios en un cuarto
	    s = "ROOM_USERS_LIST";
	    break;
	case ROOM_TEXT_FROM: //Se envia un mensaje en especifico a algun cuarto
	    s = "ROOM_TEXT_FROM";
	    break;
	case LEFT_ROOM: //Cuando un usuario abandona un cuarto
	    s = "LEFT_ROOM";
	    break;
	case DISCONNECTED: //Operación no reconocida
	    s = "DISCONNECTED";
	    break;
	case RESPONSE: //Respuesta del servidor
	    s = "RESPONSE";
	    break;
	default:
	}
    return s;
}


public Tipo toTipo(string s){
    Tipo i = null;
    switch (s) {
	case "IDENTIFY":
	    i = IDENTIFY;
	    break;
	case "STATUS":
	    i = STATUS;
	    break;
	case "USERS":
	    i = USERS;
	    break;
	case "TEXT":
	    i = TEXT;
	    break;
	case "PUBLIC_TEXT":
	    i = PUBLIC_TEXT;
	    break;
	case "NEW_ROOM":
	    i = NEW_ROOM;
	    break;
	case "INVITE":
	    i = INVITE;
	    break;
	case "JOIN_ROOM":
	    i = JOIN_ROOM;
	    break;
	case "ROOM_USERS":
	    i = ROOM_USERS;
	    break;
	case "ROOM_TEXT":
	    i = ROOM_TEXT;
	    break;
	case "LEAVE_ROOM":
	    i = LEAVE_ROOM;
	    break;
	case "DISCONNECT":
	    i = DISCONNECT;
	    break;
	case "INVALID":
	    i = INVALID_TIPO;
	    break;
	case "NEW_USER":
	    i = NEW_USER;
	    break;
	case "NEW_STATUS":
	    i = NEW_STATUS;
	    break;
	case "USER_LIST":
	    i = USER_LIST;
	    break;
	case "TEXT_FROM":
	    i = TEXT_FROM;
	    break;
	case "PUBLIC_TEXT_FROM":
	    i = PUBLIC_TEXT_FROM;
	    break;
	case "INVITATION":
	    i = INVITATION;
	    break;
	case "JOINED_ROOM":
	    i = JOINED_ROOM;
	    break;
	case "ROOM_USERS_LIST":
	    i = ROOM_USERS_LIST;
	    break;
	case "ROOM_TEXT_FROM":
	    i = ROOM_TEXT_FROM;
	    break;
	case "LEFT_ROOM":
	    i = LEFT_ROOM;
	    break;
	case "DISCONNECTED":
	    i = DISCONNECTED;
	    break;
	case "RESPONSE":
	    i = RESPONSE;
	    break;
    }
    return i;
}
