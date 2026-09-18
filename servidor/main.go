package main

import (
	//Formato del texto
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


//Función para indicar como usar el programa
func uso(){
	fmt.Println("go run main.go" + "puerto" + " (1024-65535)")
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

	//Verifica que sea un puerto valido
	if puerto < 1024 || puerto > 65535 {
		log.Fatal("El puerto no es valido")
		uso()
	}
	
	//Intenta conectar con el puerto dado
	escucha, err := net.Listen("tcp", ":"+os.Args[1])
	if err != nil{
		log.Fatal("No se pudo abrir el puerto: ",  puerto)
		uso()
	}

	//Asegura de que se libere el puerto cuando se termine el programa
	defer escucha.Close()


	//Crea el servidor
	servidor := NuevoServidor("Payapao's server")

	fmt.Println("Servidor", servidor.getNombre(),  "encendido correctamente en el puerto", puerto)
	

	
	//Activar servidor para aceptar clientes
	for{
		//Intenta conectar un cliente
		conexion, err := escucha.Accept()
		if err != nil{
			log.Fatal("No se pudo recibir al cliente")
			//Para que no intente agregar al cliente en estado de error
			continue
		}

		//Crea un hilo y pasa la conexión para esperar que se identifique
		go servidor.Temporal(conexion)
		
		
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
