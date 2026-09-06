package main

import (
	"fmt"
	//si estuvieran en carpetas distintas
	//nombreDeLaCarpeta/NombreDeLaFuncion(parametrosDeLaFuncion)

	//Para manejo de errores
	"log"
)

func main(){
	//agrega prefijos antes de cada texto hecho por log
	log.SetPrefix("[ERROR]")
	//Quita la fecha y hora de los textos
	log.SetFlags(0)

	personas := []string{"Vania", "Diego", "Francisco"}
	
	//De igual manera si esta en otra carpeta
	/*mensaje := servidor/Saludo("Vania ")*/
	//Sin manejo de errores
	/*mensaje := Saludo("Vania")*/
	mensaje, err := Saludos(personas)
	//si un error pasa
	if err != nil{
		log.Fatal("Se detuvo el problema devido al error; ", err)
	}
	fmt.Println(mensaje)
}
