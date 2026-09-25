using System;

public class ConvierteRespuestas{

    public string toStringRespuestas(Respuesta r){
	string s = "";
	switch (r) {
	    case Respuesta.SUCCESS: //La operación se ejecuto exitosamente
		s = "SUCCESS";
		break;
	    case Respuesta.USER_ALREADY_EXISTS: //El nombre se usuario ya exite
		s = "USER_ALREADY_EXISTS";
		break;
	    case Respuesta.NO_SUCH_USER: //El usuario no existe
		s = "NO_SUCH_USER";
		break;
	    case Respuesta.ROOM_ALREADY_EXISTS: //El nombre de la sala ya existe
		s = "ROOM_ALREADY_EXISTS";
		break;
	    case Respuesta.NO_SUCH_ROOM: //La sala no exite
		s = "NO_SUCH_ROOM";
		break;
	    case Respuesta.NOT_INVITED: //El usuario no ha sido invitado a la sala
		s = "NOT_INVITED";
		break;
	    case Respuesta.NOT_JOINED: //El usuario no se ha unido a la sala
		s = "NOT_JOINED";
		break;
	    case Respuesta.NOT_IDENTIFIED: //Usuario no identificado
		s = "NOT_IDENTIFIED";
		break;
	    case Respuesta.INVALID_RESPUESTA:
		s = "INVALID";
		break;
	    default:
		break;
	}
	return s;
    }
    
    public Respuesta toRespuesta(string s){
	Respuesta est = 0;
	switch (s) {
	    case "SUCCESS":
		est = Respuesta.SUCCESS;
		break;
	case "USER_ALREADY_EXISTS":
	    est = Respuesta.USER_ALREADY_EXISTS;
	    break;
	    case "NO_SUCH_USER":
		est = Respuesta.NO_SUCH_USER;
		break;
	    case "ROOM_ALREADY_EXISTS":
		est = Respuesta.ROOM_ALREADY_EXISTS;
		break;
	    case "NO_SUCH_ROOM":
		est = Respuesta.NO_SUCH_ROOM;
		break;
	    case "NOT_INVITED":
		est = Respuesta.NOT_INVITED;
		break;
	    case "NOT_JOINED":
		est = Respuesta.NOT_JOINED;
		break;
	    case "NOT_IDENTIFIED":
		est = Respuesta.NOT_IDENTIFIED;
		break;
	    case "INVALID":
		est = Respuesta.INVALID_RESPUESTA;
		break;
	    default:
		break;
	}
	return est;
    }
    
}
