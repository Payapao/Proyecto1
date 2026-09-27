using System;
using System.Collections.Generic;

public class Servidor {

    Cliente c;
    public Dictionary<string, Estados> clientes;
    Dictionary<string, List<string>> salas;
    List<string> invitaciones;
    private object candado = new Object();

    //Constructor de Servidor
    public Servidor(Cliente c, Dictionary<string, Estados> clientes ) {
	this.c = c;
	this.clientes = clientes;
	this.salas = new Dictionary<string, List<string>>();
	this.invitaciones = new List<string>();
    }

    //Metodo para unirse a una sala
    public void UnirseSala(string nombre){
	Mensaje cs = Mensaje.FabricaMensaje(Tipo.ROOM_USERS).roomname(nombre);
	c.Envia(cs);
	return;
    }

    //Metodo para interpretar los mensajes del cliente y convertirlos en un json
    
    public void Traductor(string s){
	
	string[] separada = s.Split('-');

	//Para que no afecten los espacios antes o despues
	for(int i = 0; i<separada.Length; i++){
	    separada[i] = separada[i].Trim();
	    //Si alguna de las cadenas es vacia lo rechaza
	    if(separada[i] == ""){
		Console.WriteLine("Ninguno de los campos puede estar vacio");
		return;
	    }
		
	}
	
	switch (separada[0]){
	    case "Actualiza estado":
	    case "actualiza estado":
	    case "ACTUALIZA ESTADO":
		if(separada.Length != 2){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
		if(separada[1] == ("ACTIVE") || separada[1] == ("active") || separada[1] ==("Active")){
		    Mensaje me = Mensaje.FabricaMensaje(Tipo.STATUS).status(Estados.ACTIVE);
		    c.Envia(me);
		}else if(separada[1] == ("AWAY") || separada[1] == ("away") || separada[1] == ("Away")){
		    Mensaje me = Mensaje.FabricaMensaje(Tipo.STATUS).status(Estados.AWAY);
		    c.Envia(me);
		}else if(separada[1] == ("BUSY") || separada[1] == ("busy") || separada[1] == ("Busy")){
		    Mensaje me = Mensaje.FabricaMensaje(Tipo.STATUS).status(Estados.BUSY);
		    c.Envia(me);
		}else
		    Console.WriteLine("El estado no es válido");
		break;
	    case "usuarios":
	    case "Usuarios":
	    case "USUARIOS":
		if(separada.Length != 2){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
		if(separada[1] == "General"){
		    lock(candado){
			Console.WriteLine("Clientes " + " - " + " Estados");
			foreach(var cliente in clientes)
			    Console.WriteLine(cliente.Key + " - " + cliente.Value);
		    }
		    break;
		}else{
		    lock(candado){
			foreach(var sala in salas)
			    if(sala.Key == separada[1]){
				Console.WriteLine("Clientes " + " - " + " Estados");
				foreach(var cliente in sala.Value){
				    Estados e = clientes[cliente];
				    Console.WriteLine(cliente + " - " + e );
				    
				}
				return;
			    }
		    }
		    Console.WriteLine("Es necesario pertenecer a la sala para acceder a la lista de usuarios");
		}		
		break;
	    case "salas":
	    case "Salas":
	    case "SALAS":
		if(separada.Length != 1){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
		if(salas.Count == 0)
		    Console.WriteLine("Aún no has sido invitado a ninguna sala");
		else{
		    Console.WriteLine("Las salas a las que perteneces son:");
		    lock(candado)
			foreach(var sala in salas.Keys)
			    Console.WriteLine(sala);    
		}
		break;
	    case "invitaciones":
	    case "Invitaciones":
	    case "INVITACIONES":
		if(separada.Length != 1){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
		if(invitaciones.Count == 0)
		    Console.WriteLine("Aún no has sido invitado a ninguna sala");
		else{
		    Console.WriteLine("Las salas a las que has sido invitado son:");
		    lock(candado)
			foreach(string sala in invitaciones)
			    Console.WriteLine(sala);    
		}
		break;
	    case "Crea":
	    case "CREA":
	    case "crea":
		if(separada.Length != 2){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
		Mensaje mc = Mensaje.FabricaMensaje(Tipo.NEW_ROOM).roomname(separada[1]);
		c.Envia(mc);
		break;
	    case "Invita":
	    case "INVITA":
	    case "invita":
		if(separada.Length != 3){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
		string[] invitados = separada[2].Split(',');
		lock(candado){
		    foreach(string invitado in invitados)
			if (!clientes.ContainsKey(invitado)){
			    Console.WriteLine("El usuario "+ invitado+ " no existe");
			    break;
			}
		    }
		Mensaje mi = Mensaje.FabricaMensaje(Tipo.INVITE).roomname(separada[1]).usernames(invitados);
		c.Envia(mi);
		break;
	    case "UNIRSE":
	    case "Unirse":
	    case "unirse":
		if(separada.Length != 2){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
		lock(candado){
		    foreach(string invitado in invitaciones)
			if(invitado == separada[1]){
			    Mensaje mu = Mensaje.FabricaMensaje(Tipo.JOIN_ROOM).roomname(separada[1]);
			    c.Envia(mu);
			    invitaciones.Remove(invitado);
			    break;
			}
		    }
		Console.WriteLine("Es necesario que estes invitado a la sala para unirtea la sala "+  separada[1]);
		break;
	    case "ABANDONA":
	    case "Abandona":
	    case "abandona":
		if(separada.Length != 2){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
		lock(candado){
		    foreach(string sala in salas.Keys)
			if(sala == separada[1]){
			    Mensaje ma = Mensaje.FabricaMensaje(Tipo.LEAVE_ROOM).roomname(separada[1]);
			    c.Envia(ma);
			    break;
			}
		}
		Console.WriteLine("No eres miembro de la sala "+separada[1]);
		break;
	    case "DESCONECTA":
	    case "Desconecta":
	    case "desconecta":
		if(separada.Length != 1){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
	    Desconecta();
	    break;
	    default:
		if(separada.Length != 2){
		    Console.WriteLine("El formato no es el correcto para la operación");
		    UsoServidor();
		    break;
		}
		if(separada[0] == "TODOS" || separada[0] == "Todos" || separada[0] == "todos"){
		    Mensaje mt = Mensaje.FabricaMensaje(Tipo.PUBLIC_TEXT).text(separada[1]);
		    c.Envia(mt);
		    return;
		}else {
		    lock(candado){
			foreach (string usuario in clientes.Keys)
			    if (usuario == separada[0]){
				Mensaje mu = Mensaje.FabricaMensaje(Tipo.TEXT).username(usuario).text(separada[1]);
				c.Envia(mu);
				return;
				}
			
			foreach (string sala in salas.Keys)
			    if(sala == separada[0]){
				Mensaje ms = Mensaje.FabricaMensaje(Tipo.ROOM_TEXT).roomname(separada[1]);
				c.Envia(ms);
				return;
			    }
			}
		    //No se encontro el usuario o sala
		    Console.WriteLine("El usuario o la sala "+ separada[0]+ " no existe");
		}
		break;	    
	}
	return;
    }

    //Escucha los mensajes del servidor
    public void EscuchaServidor(){
	//Mientras siga conectado con el servidor
	while (true) {
	    //Decodifica el mensaje
	    Mensaje m = c.Recibe();

	    //Si el mensaje no es valido
	    if(m == null)
		return;

	    switch (m.Type){
		case Tipo.RESPONSE:
		    ManejaRespuestas(m);
		    break;
		case Tipo.NEW_USER:
		    lock(candado){
			clientes.Add(m.Username, Estados.ACTIVE);
		    }
		    break;
		case Tipo.NEW_STATUS:
		    lock(candado){
			clientes[m.Username] = m.Status;
		    }
		    break;
		case Tipo.TEXT_FROM:
		    Console.WriteLine(m.Username + "- " + m.Text);
		    break;
		case Tipo.PUBLIC_TEXT_FROM:
		    Console.WriteLine("General - "+ m.Username +" - " + m.Text);
		    break;
		case Tipo.INVITATION:
		    Console.WriteLine("El usuario " + m.Username + " te ha invitado a la sala " + m.Roomname);
		    lock(candado){
			invitaciones.Add(m.Roomname);
		    }
		    break;
		case Tipo.JOINED_ROOM:
		    lock(candado){
			salas[m.Roomname].Add(m.Username);
		    }
		    break;
		case Tipo.ROOM_USERS_LIST:
		    lock(candado){
			salas.Add(m.Roomname, new List<string>(m.Users.Keys));
		    }
		    break;
		case Tipo.ROOM_TEXT_FROM:
		    Console.WriteLine(m.Roomname + "-" + m.Username + ":" + m.Text);
		    break;
		case Tipo.LEFT_ROOM:
		    lock(candado){
			salas[m.Roomname].Remove(m.Username);
		    }
		    break;
		case Tipo.DISCONNECTED:
		    lock(candado){
			clientes.Remove(m.Username);
		    }
		    break;
		default:
		    Console.WriteLine("Error al recibir la respuesta");
		    Desconecta();
		    break;
	    }
	}
    }

    public void ManejaRespuestas(Mensaje m){
	switch (m.Result){
	    case Respuesta.SUCCESS:
		ManejaOperacion(m);
		break;
	    case Respuesta.NO_SUCH_USER:
		Console.WriteLine("El usuario "+ m.Extra + " no existe");
		break;
	    case Respuesta.ROOM_ALREADY_EXISTS:
		Console.WriteLine("El nombre de sala " + m.Extra + " ya existe en el servidor");
		break;
	    case Respuesta.NO_SUCH_ROOM:
		Console.WriteLine("La sala "+ m.Extra + " no existe");
		break;
	    case Respuesta.NOT_INVITED:
		Console.WriteLine("No fue posible unirse a la sala " + m.Extra + " debido a que no has sido invitado o no existe");
		break;
	    case Respuesta.NOT_JOINED:
		Console.WriteLine("No formas parte de la sala " + m.Extra);
		break;
	    default:
		Console.WriteLine("Algo salio mal");
		Desconecta();
		break;
	}
	return;
    }
    
    public void ManejaOperacion(Mensaje m){
	switch(m.Operation){
	    case Tipo.NEW_ROOM:
		lock(candado){
		    salas.Add(m.Extra, new List<string> {c.Username});
			     }
		Console.WriteLine("La sala " + m.Extra + " ha sido creada correctamente");
		break;
	    case Tipo.JOIN_ROOM:
		UnirseSala(m.Extra);
		break;
		//Este caso se agrego por una extencion al protocolo
	    case Tipo.STATUS:
		lock(candado)
		    clientes[c.Username] = m.Status;
		break;
	    default:
		Console.WriteLine("Ocurrio un error inesperado");
		Desconecta();
		break;
		}
	return;
    }

    //Metodo que explica como usar el chat
    public void UsoServidor(){
	Console.WriteLine("Uso del chat:\n" +
			  "Para actalizar un estado escriba AWAY, BUSY o ACTIVE:\n " +
			  "                        Actaliza estado - estado\n" +
			  "Para obtener la lista de usuarios del servidor escriba:\n" +
			  "                        Usuarios - General\n" +
			  "Para obtener la lista de las salas a las que perteneces:\n" +
			  "                        Salas\n" +
			  "Para obtener la lista de las salas a las que has sido invitado:\n" +
			  "                        Invitaciones\n" +
			  "Para mandar un mensaje a todos escriba:\n" +
			  "                        Todos - mensaje\n" +
			  "Para mandar un mensaje a un usuario escriba\n" +
			  "                        nombre_del_usuario - mensaje\n" +
			  "Para mandar mensaje en una sala escriba\n" +
			  "                        nombre_de_la_sala - mensaje\n"+
			  "Para crear una sala escriba:\n" +
			  "                        Crea - nombre_de_la_sala\n" +
			  "Para invitar a usuarios a la sala escriba:\n" +
			  "                        Invita - nombre_de_la_sala - usuario1, usuario2, usuario3\n" +
			  "Para unirse a una sala a la que fue invitado escriba\n" +
			  "                        Unirse - nombre_de_la_sala\n" +
			  "Para obtener la lista de usuarios de una sala en especifico escrba\n" +
			  "                        Usuarios - nombre_de_la_sala\n" +
			  "Para abandonar una sala escriba:\n"+
			  "                        Abandona - nombre_de_la_sala\n" +
			  "Para desconectarse del servidor escriba:\n"+
			  "                        Desconecta");
	return;
    }

    //Metodo para enviar el mensaje de desconección  a los usuarios
    public void Desconecta(){
	Mensaje m = Mensaje.FabricaMensaje(Tipo.DISCONNECT);
	c.Envia(m);
	Environment.Exit(0);
	return;
    }
}
