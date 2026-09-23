using System;

public enum Tipo {

    SUCCESS,//La operación se ejecuto exitosamente
    USER_ALREADY_EXISTS, //El nombre se usuario ya exite
    NO_SUCH_USER, //El usuario no existe
    ROOM_ALREADY_EXISTS, //El nombre de la sala ya existe
    NO_SUCH_ROOM, //La sala no exite
    NOT_INVITED, //El usuario no ha sido invitado a la sala
    NOT_JOINED, //El usuario no se ha unido a la sala
    NOT_IDENTIFIED, //Usuario no identificado
    INVALID_RESPUESTA //Jason invalido
    
}
