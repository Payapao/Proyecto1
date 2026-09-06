package main

import (
	"fmt"
	//manejo de errores
	"errors"
	//randoms de math
	"math/rand"
)


//Saludo es publica porque comienza con mayuscula
//Función sin manejo de errores
/*func Saludo(name string)string{
	var mensaje string
	mensaje = fmt.Sprintf("hola, %v. Bienvenido", name)
	return mensaje
        } */

//Función con manejo de errores
func Saludo(name string)(string, error){
	if name == "" {
		return "", errors.New("nombre vacio")
	}
	//es equivalente a la de arriba
	/*mensaje := fmt.Sprintf("hola, %v. Bienvenido", name)
	return mensaje, nil*/

	//Con formato aleatorio
	mensaje := fmt.Sprintf(randomFormat(), name)
	return mensaje, nil
}

func Saludos(names []string)(map[string]string, error){
	//map[tipollave]tipovalor
	//mensajes[tipollave] puede regresar 2 valores; valor y existe
	//mensaje, existe := mensajes[Vania] --La llave puede ser utilizada como indice
	//delete(mensajes, "Vania") elimina el saludo hacia Vania
	mensajes := make(map[string]string)

	//range regresa 2 valores, indice y valor
	//for index, name := range names
	for _, name := range names{
		mensaje, err := Saludo(name)
		if err != nil{
			return nil, err
		}

		mensajes[name] = mensaje
	}
	return mensajes, nil
}

//randomFormat es privada porque empieza con minuscula
func randomFormat() string{
	//Crea un slice(similar a arreglo) --%v = placeholder
	//Para hacer uso de %v es necesario utilizar fmt.Sprintf
	//Con fmt.Sprint no son sustituidos los %v
	formats := []string{
		"Hola, %v. Bienvenido",
		"Que gusto verte, %v",
		"Saludos! Un placer conocerte %v",
	}

	return formats[rand.Intn(len(formats))]
}
