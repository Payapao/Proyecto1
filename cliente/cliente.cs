using System;

public class Cliente{

    string username;
    Estado estado;
    TCPClient conexion;
    StreamReader decodificador;
    StreamWriter codificador;
    

    //Es necesario hacer un constructor de cliente
    public Cliente(string u, string ip, int puerto){
	this.username = u;
	this.estado = ACTIVE;
	this.conexion = new TcpClient(ip, puerto);
	NetworkStream flujo = new conexion.GetStream;

	decodificador = new StreamReader(flujo, Encoding.UTF8);
	codificador = new StreamWriter(flujo, Encoding.Utf8)
    }

    //Metodo envia para codificar los mensajes al servidor
    public Envia(Mensaje m){
	string json = JsonSerializer.Serialize(m);
	codificador.WriteLine(json);
    }

    //Metodo envia para decodificar los mensajes al servidor
    public Mensaje Recibe(){
	string json = decodificador.ReadLine();
	if(json == null){
	    Console.WriteLine("El servidor cerro la conexión");
	    return null;
	}
	return JsonSerializer.Deserialize<Mensaje>(json);
    }

    //Metodo temporal para esuchar la lista de usuarios y crear el servidor
    public Temporal()(Servidor){
	Mensaje mi = FabricaMensaje(IDENTIFY).username(username);
	Envia(mi);
	Mensaje m = c.Recibe();
	
	switch (m.Result){
	    case SUCCESS:
		Console.WriteLine("Se ha unido al servidor de manera exitosa");
		break;
	    case USER_ALREADY_EXISTS:
		Console.WriteLine("El nombre de usuario ya existia en el servidor, intente volver a unirse con un nombre distinto");
		return null;
	    default:
		Console.WriteLine("No fue posible unirse al servidor");
		return null;
	}
	
	Mensaje mu = FabricaMensaje(USERS);
	Envia(mu);

	while(true){
	    	m = c.Recibe();
		if (m.Type == USER_LIST)
		    return NuevoServidor(this, m.Users);
	}
	
	return s;
    }

}
