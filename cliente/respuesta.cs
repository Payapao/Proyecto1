using System;

public enum Respuesta {
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
    RESPONSE, //Respuesta del servidor
}
