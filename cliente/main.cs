using System;
using System.Threading;
using System.Net.Sockets;
using System.Collections.Generic;

class Proyecto1 {
    
    static void UsoBanderas(){
	Console.WriteLine("Uso de las banderas\n" +
	                  "'-h' muestra este mensaje de ayuda\n" +
			  "'-i' para indicar la dirección ip\n" +
			  "'-p' para indicar el puerto (1024-65535)\n" +
	                  "Ejemplo: ./build/cliente.exe -i localhost -p 1234");
    }
    
    static void Main(string[] args){
	
	string ip = "";
	int puerto = 0;
	
	for(int i = 0; i<args.Length; i++)
	    if (args[i] == "-h"){
		UsoBanderas();
		return;
	    }
	//Verifica que la cantidad de banderas sea exactamente las necesarias
	if (args.Length < 4 || args.Length > 4){
	    UsoBanderas();
	    return;
	}
	
	//Lee las banderas para guardar los valores
	for(int i = 0; i<args.Length; i++){
	    switch(args[i]){
		case "-i":
		    if(i+1 >= args.Length){
			Console.Write("La bandera -i necesita tener una dirección ip");
			UsoBanderas();
			return;
		    }
		    ip = args[i+1];
		    i++;
		    break;
		case "-p":
		    if(i+1 >= args.Length){
			Console.Write("La bandera -p necesita tener un puerto");
			UsoBanderas();
			return;
		    }
		    //Intenta convertir el argumento a un entero
		    bool numero = int.TryParse(args[i+1], out puerto);
		    if (!numero){
			Console.Write("La bandera -p necesita tener un puerto");
			UsoBanderas();
			return;
		    }
		    i++;
		    break;
		default:
		    UsoBanderas();
		    return;
	    }
	}
	
	//Checa que el puerto sea valido
	if(puerto < 1024 || puerto > 65535){
	    Console.Write("El puerto no es valido");
	    UsoBanderas();
	    return;
	}
	
	//Pide el nombre de usuario antes de conectarse
	Console.WriteLine("Para continuar es necesario agregar un username");
	Console.Write("Username: ");
	
	string username = Console.ReadLine();
	username = username.Trim();
	
	if (username == ("")){
	    Console.WriteLine("El nombre de usuario no es válido");
	    return;
	}
	
	if(username.Length > 8){
	    Console.WriteLine("El nombre de usuario debe tener una longitud menor a 9");
	    return;
	}
	
	//Crea al clinete
	Cliente c = null;
	
	//Checa que la dirección ip sea valida
	try{
	    c = new Cliente(username, ip, puerto);
	}catch(SocketException){
	    Console.Write("No fue posible conectarse al servidor");
	    return;
	}
	
	
	
	//Es necesario que escucha y escribe esten en hilos separados
	Servidor s = c.Temporal();
	if(s != null){
	    s.UsoServidor();
	    Thread EscuchaServidor = new Thread(() => {s.EscuchaServidor();});
	    EscuchaServidor.Start();

	    //Para leer la terminal
	    while(true){
		string entrada = Console.ReadLine();
		if(entrada == null)
		    break;
		s.Traductor(entrada);
	    }
	}
	
    }
}
