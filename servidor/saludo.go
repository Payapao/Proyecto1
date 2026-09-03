package main

import (
	"fmt"
	//manejo de errores
	"errors"
)


//Saludo es publica
//saludo es privada
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
	mensaje := fmt.Sprintf("hola, %v. Bienvenido", name)
	return mensaje, nil
}
