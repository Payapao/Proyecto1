using System;
using System.IO;
using System.Net.Sockets;
using System.Text;
using System.Runtime.Serialization.Json;
using System.Collections.Generic;

public class Cliente {

    string username;
    Estados estado;
    TcpClient conexion;
    StreamReader decodificador;
    StreamWriter codificador;
    
    //Getters y setters
    public string Username{
	get {return username;}
    }

    public Estados Estado{
	get{ return estado;}
	set{ estado = value;}
    }
    
    //Es necesario hacer un constructor de cliente
    public Cliente(string u, string ip, int puerto){
	this.username = u;
	this.estado = Estados.ACTIVE;
	this.conexion = new TcpClient(ip, puerto);
	NetworkStream flujo = conexion.GetStream();

	decodificador = new StreamReader(flujo, Encoding.UTF8);
	codificador = new StreamWriter(flujo, Encoding.UTF8) {AutoFlush = true};
    }

    //Metodo envia para codificar los mensajes al servidor
    public void Envia(Mensaje m){
	try{
	    var seria = new DataContractJsonSerializer(typeof(Mensaje));

	    //Se transforma a un json, se manda y fuerza la salida
	    using (var ms = new MemoryStream()){
		seria.WriteObject(ms, m);
		string json = Encoding.UTF8.GetString(ms.ToArray());
		Console.Write(json);
		codificador.WriteLine( json);
	    }

	}catch( Exception e){
	    Console.WriteLine($"Error: {e.Message}");
	}
	return;
    }

    //Metodo envia para decodificar los mensajes al servidor
    public Mensaje Recibe(){
	string json = decodificador.ReadLine();
	if(json == null){
	    Console.WriteLine("El servidor cerro la conexión");
	    return null;
	}

	//
	using (var ms = new MemoryStream(Encoding.UTF8.GetBytes(json))) {
            var serializador = new DataContractJsonSerializer(typeof(Mensaje));
            return (Mensaje)serializador.ReadObject(ms);
        }
    }

    //Metodo temporal para esuchar la lista de usuarios y crear el servidor
    public Servidor Temporal(){
	Mensaje mi = Mensaje.FabricaMensaje(Tipo.IDENTIFY).username(this.username);
	Envia(mi);
	Mensaje m = Recibe();
	if (m == null)
	    return null;
	
	switch (m.Result){
	    case Respuesta.SUCCESS:
		Console.WriteLine("Se ha unido al servidor de manera exitosa :)");
		break;
	    case Respuesta.USER_ALREADY_EXISTS:
		Console.WriteLine("El nombre de usuario ya existia en el servidor, intente volver a unirse con un nombre distinto");
		return null;
	    default:
		Console.WriteLine("No fue posible unirse al servidor");
		return null;
	}
	
	Mensaje mu = Mensaje.FabricaMensaje(Tipo.USERS);
	Envia(mu);

	while(true){
	    	m = Recibe();
		if(m == null){
		    Console.WriteLine("Error al recibir respuesta del servidor");
		    return null;
		}
		if (m.Type == Tipo.USER_LIST)
		    return new Servidor(this, m.Users);
	}
	
    }
    
}
