using System;

public enum Respuesta {

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

public string toStringRespuestas(Respuesta r){
    string s = "";
    switch (r) {
	case SUCCESS: //La operación se ejecuto exitosamente
	    s = "SUCCESS";
	    break;
	case USER_ALREADY_EXISTS: //El nombre se usuario ya exite
	    s = "USER_ALREADY_EXISTS";
	    break;
	case NO_SUCH_USER: //El usuario no existe
	    s = "NO_SUCH_USER";
	    break;
	case ROOM_ALREADY_EXISTS: //El nombre de la sala ya existe
	    s = "ROOM_ALREADY_EXISTS";
	    break;
	case NO_SUCH_ROOM: //La sala no exite
	    s = "NO_SUCH_ROOM";
	    break;
	case NOT_INVITED: //El usuario no ha sido invitado a la sala
	    s = "NOT_INVITED";
	    break;
	case NOT_JOINED: //El usuario no se ha unido a la sala
	    s = "NOT_JOINED";
	    break;
	case NOT_IDENTIFIED: //Usuario no identificado
	    s = "NOT_IDENTIFIED";
	    break;
	case INVALID_RESPUESTA:
	    s = "INVALID";
	    break;
	default:
    }
    return s;
}

public Respuesta toRespuesta(string s){
    Respuesta est = null;
    switch (s) {
	case "SUCCESS":
	    est = SUCCESS;
	    break;
	case "USER_ALREADY_EXISTS":
	    est = USER_ALREADY_EXISTS;
	    break;
	case "NO_SUCH_USER":
	    est = NO_SUCH_USER;
	    break;
	case "ROOM_ALREADY_EXISTS":
	    est = ROOM_ALREADY_EXISTS;
	    break;
	case "NO_SUCH_ROOM":
	    est = NO_SUCH_ROOM;
	    break;
	case "NOT_INVITED":
	    est = NOT_INVITED;
	    break;
	case "NOT_JOINED":
	    est = NOT_JOINED;
	    break;
	case "NOT_IDENTIFIED":
	    est = NOT_IDENTIFIED;
	    break;
	case "INVALID":
	    est = INVALID_RESPUESTA;
	    break;
	default:
	}
    return est;
}
