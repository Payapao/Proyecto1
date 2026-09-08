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

//Estructura que contenga todas las salas
type Servidor struct{
	salas map[string]*Sala
	candado sync.Mutex
}

//Crea el servidor
func NuevoServidor(nombre string) *Servidor{
	return &Servidor{
		salas: make(map[string]*Sala)
	}
}

//Función para indicar como usar el programa
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

	//Crea el servidor
	servidor := NuevoServidor("Payapao's server")
	
	//Crea la sala donde estan todos los clientes
	clientes, _ := NuevaSala("General")

	//Se agrega la sala grupal al servidor
	map[clientes.getNombre()]clientes

	
	//Activar servidor para aceptar clientes
	for{
		//Intenta conectar un cliente
		conexion, err := escucha.Accept()
		if err != nil{
			log.Fatal("No se pudo recibir al cliente")
			//Para que no intente agregar al cliente en estado de error
			continue
		}

		//Crea un cliente
		cliente := NuevoCliente("aqui va el nombre", conexion)
		//Agrega el cliente al servidor
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
