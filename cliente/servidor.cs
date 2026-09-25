using System;
using System.Collections.Generic;

public class Servidor {

    Cliente c;
    Dictionary<string, Estados> clientes;
    Dictionary<string, List<string>> salas;
    List<string> invitaciones;
    private object candado = new Object();

    //Metodo para crear un nuevo servidor
    public Servidor(Cliente c, Dictionary<string, Estados> clientes ) {
	this.c = c;
	this.clientes = clientes ?? new Dictionary<string, Estados>();
	this.salas = new Dictionary<string, List<string>>();
	this.invitaciones = new List<string>();
    }

    //Metodo para unirse a una sala
    public void UnirseSala(string nombre){
	Mensaje cs = Mensaje.FabricaMensaje(Tipo.ROOM_USERS);
	c.Envia(cs);
	return;
    }

    //Metodo para interpretar los mensajes del cliente y convertirlos en un json
    
    public void Traductor(string s){
	
	string[] separada = s.Split('-');

	//Para que no afecten los espacios antes o despues
	for(int i = 0; i<separada.Length; i++)
	    separada[i] = separada[i].Trim();
	
	switch (separada[0]){
	    case "Actualiza estado":
	    case "actualiza estado":
	    case "ACTUALIZA ESTADO":
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
		Console.WriteLine("Clientes " + " - " + " Estados");
		if(separada[1] == "General")
		    foreach(var cliente in clientes)
			
			Console.WriteLine(cliente.Key + " - " + cliente.Value);
		else{
		    foreach(var sala in salas)
			if(sala.Key == separada[1])
			    foreach(var cliente in sala.Value){
				Estados e = clientes[cliente];
				Console.WriteLine(cliente + "-" + e );
			    }
		    
		    Console.WriteLine("Es necesario pertenecer a la sala para acceder a la lista de usuarios");
		}		
		break;
	    case "Crea":
	    case "CREA":
	    case "crea":
		Mensaje mc = Mensaje.FabricaMensaje(Tipo.NEW_ROOM).roomname(separada[1]);
		c.Envia(mc);
		break;
	    case "Invita":
	    case "INVITA":
	    case "invita":
		string[] invitados = separada[1].Split(',');
		foreach(string invitado in invitados)
		    if (!clientes.ContainsKey(invitado)){
			Console.WriteLine("El usuario "+ invitado+ " no existe");
			break;
		    }
		Mensaje mi = Mensaje.FabricaMensaje(Tipo.INVITE).roomname(separada[1]).usernames(invitados);
		c.Envia(mi);
		break;
	    case "UNIRSE":
	    case "Unirse":
	    case "unirse":
		foreach(string invitado in invitaciones)
		    if(invitado == separada[1]){
			Mensaje mu = Mensaje.FabricaMensaje(Tipo.JOIN_ROOM).roomname(separada[1]);
			c.Envia(mu);
			break;
		    }
		Console.Write("Es necesario que estes invitado a la sala para unirtea la sala " + separada[1]);
		break;
	    case "ABANDONA":
	    case "Abandona":
	    case "abandona":
		foreach(string sala in salas.Keys)
		    if(sala == separada[1]){
			Mensaje ma = Mensaje.FabricaMensaje(Tipo.LEAVE_ROOM).roomname(separada[1]);
			c.Envia(ma);
			break;
		    }
		Console.WriteLine("No eres miembro de la sala "+separada[1]);
		break;
	    case "DESCONECTA":
	    case "Desconecta":
	    case "desconecta":
	    Desconecta();
	    break;
	    default:
		if (separada.Length == 2)
		    if(separada[0] == "TODOS" || separada[0] == "Todos" || separada[0] == "todos"){
			Mensaje mt = Mensaje.FabricaMensaje(Tipo.PUBLIC_TEXT).text(separada[1]);
			c.Envia(mt);
		    }else {
			foreach (string usuario in clientes.Keys)
			    if (usuario == separada[0]){
				Mensaje mu = Mensaje.FabricaMensaje(Tipo.TEXT).username(usuario).text(separada[1]);
				c.Envia(mu);
				break;
			    }
			
			foreach (string sala in salas.Keys)
			    if(sala == separada[0]){
				Mensaje ms = Mensaje.FabricaMensaje(Tipo.ROOM_TEXT).roomname(separada[1]);
				c.Envia(ms);
				break;
			    }
			//No se encontro el usuario o sala
			Console.WriteLine("El usuario o la sala "+ separada[1]+ " no existe");
		    }
		else{
		    Console.WriteLine("El formato no coincide con el uso del chat");
		    UsoServidor();
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
		    Console.WriteLine("El usuario " + m.Username + " te ha invitado a la sala "+ m.Roomname);
		    invitaciones.Add(m.Roomname);
		    break;
		case Tipo.JOINED_ROOM:
		    salas[m.Roomname].Add(m.Username);
		    break;
		case Tipo.ROOM_USERS_LIST:
		    salas.Add(m.Roomname, new List<string>(m.Users.Keys));
		    salas[m.Roomname].Add(c.Username);
		    break;
		case Tipo.ROOM_TEXT_FROM:
		    Console.Write(m.Roomname + "-" + m.Username + ":" + m.Text);
		    break;
		case Tipo.LEFT_ROOM:
		    salas[m.Roomname].Remove(m.Username);
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
		Console.WriteLine("No fue posible unirse a la sala " + m.Extra + " debido a que no has sido invitado");
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
		salas.Add(m.Extra, new List<string> {c.Username});
		Console.WriteLine("La sala " + m.Extra + " ha sido creada correctamente");
		break;
	    case Tipo.JOIN_ROOM:
		UnirseSala(m.Extra);
		invitaciones.Remove(m.Extra);
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
			  "Para mandar un mensaje a todos escriba:\n" +
			  "                        Todos - mensaje\n" +
			  "Para mandar un mensaje a un usuario escriba\n" +
			  "                        nombre_del_usuario - mensaje\n" +
			  "Para mandar mensaje en una sala escriba\n" +
			  "                        nombre_de_la_sala - mensaje\n"+
			  "Para crear una sala escriba:\n" +
			  "                        Crea - nombre_de_la_sala\n" +
			  "Para invitar a usuarios a la sala escriba:\n" +
			  "                        Invita - usuario1, usuario2, usuario3\n" +
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
