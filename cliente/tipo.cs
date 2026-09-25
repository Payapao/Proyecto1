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
