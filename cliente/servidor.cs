using System;

public class Servidor{

    Cliente c;
    Dictionary<string, Estados>() clientes;
    Dictionary<string, List<string>>() salas;
    string[] invitaciones;

    //Metodo para crear un nuevo servidor
    public NuevoServidor(Cliente c, Dictionary[string] clientes ){
	this.c = c;
	this.clientes = clientes;
	this.salas = new Dictionary<string, List<string>>();
	this.invitaciones = new List<string>;
    }

    //Metodo para unirse a una sala
    public UnirseSala(string nombre){
	cs = FabricaMensaje(ROOM_USERS);
	c.Envia(cs);
    }

    //Metodo para interpretar los mensajes del cliente y convertirlos en un json
    
    public Traductor(string s){
	
	string[] separada = s.Split('-');
	
	switch (separada[0]){
	    case "Actualiza estado":
	    case "actualiza estado";
	    case "ACTUALIZA ESTADO";
	    if(separada[1].equals("ACTIVE") || separada[1].equals("active") || separada[1].equals("Active")){
		Mensaje me = FabricaMensaje(STATUS).status(ACTIVE);
		c.Envia(me);
	    }else if(separada[1].equals("AWAY") || separada[1].equals("away") || separada[1].equals("Away")){
		me = FabricaMensaje(STATUS).status(AWAY);
		c.Envia(me);
	    }else if(separada[1].equals("BUSY") || separada[1].equals("busy") || separada[1].equals("Busy")){
		me = FabricaMensaje(STATUS).status(BUSY);
		c.Envia(me);
	    }else
		Console.WriteLine("El estado no es válido");
	    break;
	    case "usuarios":
	    case "Usuarios":
	    case "USUARIOS":
		Console.WriteLine("Clientes " + "-" + " Estados");
		if(string[2].equals("General"))
		    foreach(var cliente in clientes)
			Console.WriteLine(cliente.Key + "-" + cliente.Value);
		else{
		    foreach(var sala in salas)
			if(sala.Key.equals(string[1]))
			    foreach(var cliente in sala.Value)
				Console.WriteLine(cliente.Key + "-" + cliente.Value);
		    
		    Console.WriteLine("Es necesario pertenecer a la sala para acceder a la lista de usuarios");
		}		
		break;
	    case "Crea":
	    case "CREA":
	    case "crea":
		Mensaje mc = FabricaMensaje(NEW_ROOM).romname(separada[1]);
	    case "Invita":
	    case "INVITA":
	    case "invita":
		string[] invitados = separada[2].Split(',');
		foreach(string invitado in initados)
		    if (!clientes.ContainsKey(invitado)){
			Console.WriteLine("El usuario"+ invitado+ "no existe");
			break;
		    }
		Mensaje mi = FabricaMensaje(INVITE).roomname(separada[1]).usernames(invitados);
		c.Envia(mi);
		break;
	    case "UNIRSE":
	    case "Unirse":
	    case "unirse":
		foreach(string invitado in invitaciones)
		    if(invitado.equals(separada[1])){
			Mensaje mu = FabricaMensaje(JOIN_ROOM).roomname(separada[1]);
			c.Envia(mu);
			break;
		    }
		Console.Write("Es necesario que estes invitado a la sala para unirtea la sala" + separada[1]);
	    case "ABANDONA":
	    case "Abandona":
	    case "abandona":
		foreach(string sala in salas.Keks)
		    if(sala.equals(separada[1])){
			Mensaje ma = FabricaMensaje(LEAVE_ROOM).roomname(separada[1]);
			c.Envia(ma);
			break;
		    }
		Console.Write("No eres miembro de la sala"+separada[1]);
		break;
	    case "DESCONECTA";
	    case "Desconecta";
	    case "desconecta";
	    Desconecta();
	    break;
	    default:
		if (separada.Lenght == 2)
		    if(string[0].equals("TODOS") || string[0].equals("Todos") || string[0].equals("todos")){
			mt = FabricaMensaje(PUBLIC_TEXT).text(separada[1]);
			c.Envia(mt);
		    }else {
			foreach (string usuario in clientes.Keys){
			    if (usuario.equals(string[0])){
				mu = FabricaMensaje(TEXT).username(usuario).text(separada[1]);
				c.Envia(mu);
				break;
			    }
			}
			foreach (string sala in salas.Keys)
			    if(sala.equals(string[0])){
				ms = FabricaMensaje(ROOM_TEXT).roomname(separa[1]);
				c.Envia(ms);
				break
			    }
			Console.Write("El usuario o la sala "+ separada[1]+ no existe);
		    }
		else
		    Console.Write("El usuario o la sala "+ separada[1]+ "no existe");
	}
	
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
		    break;
		case NEW_USER:
		    clientes.Add(m.Username, ACTIVE);
		    break;
		case NEW_STATUS:
		    clientes[m.Username] = m.Status;
		    break;
		case TEXT_FROM:
		    Console.Write(m.Username + "- " + m.Text);
		    break;
		case PUBLIC_TEXT_FROM:
		    Console.Write("General -" m.Username "- " + m.Text);
		    break;
		case INVITATION:
		    Console.Write("El usuario" + m.Username + " te ha invitado a la sala "+ m.Roomname);
		    invitaciones.Add(m.Roomname);
		    break;
		case JOINED_ROOM:
		    salas[m.Roomname].Add(m.Username);
		    break;
		case ROOM_USER_LIST:
		    salas.Add(m.Roomname, m.Users);
		    salas[m.Roomname].Add(c.username);
		    break;
		case ROOM_TEXT_FROM:
		    Console.Write(m.Roomname + "-" + m.Username + ":" + m.Text);
		    break;
		case LEFT_ROOM:
		    salas[m.Roomname].Remove(m.Username);
		    break;
		case DISCONNECTED:
		    cientes.Remove(m.Username);
		    break;
		default:
		    Console.Write("Error al recibir la respuesta");
		    Desconecta();
	    }
	}
    }

    public ManejaRespuestas(m Mensaje){
	switch (m.Result){
	    case SUCCESS:
		ManejaOperacion(m);
		break;
	    case NO_SUCH_USER:
		Console.Write("El usuario "+ m.Extra + " no existe");
		break;
	    case ROOM_ALREADY_EXISTS:
		Console.Write("El nombre de sala " + m.Extra + " ya existe en el servidor");
		break;
	    case NO_SUCH_USER:
		Console.Write("El usuario"+ m.Extra + "no puede ser invitado porque no existe");
		break;
	    case NO_SUCH_ROOM:
		Console.Write("La sala"+ m.Extra + "no existe");
		break;
	    case NOT_INVITED:
		Console.Write("No fue posible unirse a la sala" + m.Extra + "debido a que no has sido invitado");
		break;
	    case NOT_JOINED:
		Console.Write("No formas parte de la sala" + m.Extra);
		break;
	    default:
		Console.Write("Algo salio mal");
		Desconecta();
	}
    }

    public ManejaOperacion(m Mensaje){
	switch(m.Operation){
	    case NEW_ROOM:    
		salas.Add(m.Extra, new List<string> {c.username});
		Console.Write("La sala" + m.Extra + " ha sido creada correctamente");
		break;
	    case JOIN_ROOM:
		UnirseSala(m.Extra);
		invitaciones.Remove(m.Extra);
		break;
	    default:
		Console.Write("Ocurrio un error inesperado");
		Desconecta();
		}
    }


    //Metodo para enviar el mensaje de desconección  a los usuarios
    public Desconecta(){
	Mensaje m = FabricaMensaje(DISCONNECT);
	c.Envia(m);
    }
}
