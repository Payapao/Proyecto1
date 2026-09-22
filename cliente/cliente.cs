using System;

public class Cliente{

    string username;
    Estado estado;
    

    //Es necesario hacer un constructor de cliente

    //Metodo envia para codificar los mensajes al servidor
    public Envia(Mensaje m){
	
    }

    //Metodo envia para decodificar los mensajes al servidor
    public Recibe(Mensaje m){
	
    }

    //Metodo temporal para esuchar la lista de usuarios y crear el servidor
    public Temporal()(Servidor){
	Servidor s = null;
	while(s == null){
	    Mensaje m = c.Recibe();
	    
	    switch (m.Result){
		case SUCCESS:
		    break;
		case USER_LIST:
		    s = NuevoServidor(c, m.Users);
		    break;
	    }
	}
	return s;
    }
}
