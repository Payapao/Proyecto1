using System;

public class Servidor{

    Cliente c;
    Dictionary[string]Estado clientes;
    List<string> salas;
    string[] invitaciones;

    public NuevoServidor(Cliente c, Dictionary[string] clientes ){
	this.c = c;
	this.clientes = clientes;
	this.salas = new List<string>;
	this.invitaciones = new List<string>;
    }

    //Escucha los mensajes del servidor
    public EscuchaServidor(){
	//Mientras siga conectado con el servidor
	while (true) {
	    //Decodifica el mensaje
	    Mensaje mensaje = c.Recibe();

	    //Si el mensaje no es valido
	    if(mensaje == null)
		return;

	    switch (m.Type){
		case RESPONSE:
		    ManejaRespuestas(mensaje);
		case NEW_USER:
		    clientes.Add(m.Username);
		case NEW_STATUS:
		    
		default:
		    
	    }
	}
    }

    public Desconecta(c *Cliente){
	Mensaje m = FabricaMensaje(DISCONNECT);
	c.Envia(m);
    }
}
