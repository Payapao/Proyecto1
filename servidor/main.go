package main

import (
	"fmt"
	//si estuvieran en carpetas distintas
	//nombreDeLaCarpeta/NombreDeLaFuncion(parametrosDeLaFuncion)

	//Para manejo de errores
	"log"
	//Para la entrada estandar
	//"bufio"
	//Para leer argumentos iniciales
	"os"
	//Conversión de valores
	"strconv"
	//Comunicación con la red
	"net"
)

func uso(){
	fmt.Println("go run main.go" + "puerto")
	os.Exit(1)
}

func main(){
	//agrega prefijos antes de cada texto hecho por log
	log.SetPrefix("[ERROR]")
	//Quita la fecha y hora de los textos
	log.SetFlags(0)
	
	//indica si el número de parametros recividos es exactamente 1
	if len(os.Args) != 2 {
		log.Fatal("Es necesario incluir unicamente un puerto")
		uso()
	}
	
	//intenta convertir el parametro recibido en un entero
	puerto, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("El parametro recivido no es un puerto")
		uso()
	}
	
	//Intenta conectar con el puerto dado
	escucha, err := net.Listen("tcp", puerto)
	if err != nil{
		log.Fatal("No se pudo abrir el puerto: " puerto)
	}

	//Asegura de que se libere el puerto cuando se termine el programa
	defer escucha.Close()

	ftm.Println("Servidor encendido correctamente")

	//Crea la sala donde estan todos los clientes
	clientes := NuevaSala()

	
	//Activar servidor para aceptar clientes
	for{
		//Intenta conectar un cliente
		cliente, err := escucha.Accept()
		if err != nil{
			log.Fatal("No se pudo recibir al cliente")
			//Para que no intente agregar al cliente en estado de error
			continue
		}

		clientes := sala

		//Agrega el cliente a un map
		clientes.AgregaCliente(cliente)
		
	}
	
	
	//personas := []string{"Vania", "Diego", "Francisco"}
	
	//De igual manera si esta en otra carpeta
	/*mensaje := servidor/Saludo("Vania ")*/
	//Sin manejo de errores
	/*mensaje := Saludo("Vania")*/
	/*mensaje, err := Saludos(personas)
	//si un error pasa
	if err != nil{
		log.Fatal("Se detuvo el problema devido al error; ", err)
	}
	fmt.Println(mensaje)
	*/
}
