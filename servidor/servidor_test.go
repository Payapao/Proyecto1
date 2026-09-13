package main

import(
	//test
	"testing"
	//Conversión de valores
	"strconv"
)

var nombre = "Prueba"
var servidorPrueba = NuevoServidor(nombre)

func TestGetNombre(t *testing.T){

	if servidorPrueba.getNombre() != nombre {
		t.Fatal("El nombre del servidor no coincide con el nombre esperado")
	}
}

func TestGetSalas(t *testing.T){
	
	if servidorPrueba.getSalas() == nil  {
		t.Fatal("No se inicializarón correctamente las salas")
	}

	for i := 0; i<10; i++{
		//Convierte i a un string
		nombreSala := "Sala " + strconv.Itoa(i)

		_, err := servidorPrueba.NuevaSala(nombreSala)
		if err != nil{
			t.Fatal("La sala no se creo correctamente")
		}
		if len(servidorPrueba.getSalas()) != i+1 {
			t.Fatal("La salas no se agregaron correctamente")
		}

		salas := servidorPrueba.getSalas()
		_, existe := salas[nombreSala]
		if !existe {
			t.Fatal("El nombre de la sala es incorrecto")
		}	
		
	}
	
}

func TestNuevoServidor(t *testing.T){

	if servidorPrueba == nil {
		t.Fatal("El servidor no se creo correcamente")
	}
}

func TestNuevaSala(t *testing.T){
	_, err := servidorPrueba.NuevaSala("Sala de prueba")
	if err != nil {
		t.Fatal("Error al crear la sala")
	}
	
	_, error := servidorPrueba.NuevaSala("Sala de prueba")
	if error == nil {
		t.Fatal("Sala duplicada")
	}
}
