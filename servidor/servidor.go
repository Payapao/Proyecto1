package main

import(
	//Sincronización de goroutines
	"sync"
	//Manejo de errores
	"errors"
)

//Estructura que contenga todas las salas
type Servidor struct{
	salas map[string]*Sala
	candado sync.Mutex
}

//Crea el servidor
func NuevoServidor(nombre string) *Servidor{
	return &Servidor{
		salas: make(map[string]*Sala),
	}
}

//Función para crear y agregar salas al servidor
func (s *Servidor) NuevaSala(nombre string) (*Sala, error){
	s.candado.Lock()
	defer s.candado.Unlock()

	_, existe := s.salas[nombre]
	if existe {
		return nil, errors.New("El nombre de la sala ya existe")
	}

	nuevaSala := &Sala{
		nombre: nombre,
		clientes: make(map[string]*Cliente),
	}

	s.salas[nombre] = nuevaSala

	return nuevaSala, nil
	
}
