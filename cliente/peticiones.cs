using System;

public enum Peticion {
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
    INVALID_TIPO //Operación no reconocida
}
