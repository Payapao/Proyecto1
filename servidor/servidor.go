package main

import(
	//Sincronización de goroutines
	"sync"
	//Manejo de errores
	"errors"
)

//Estructura que contenga todas las salas
type Servidor struct{
	nombre string
	salas map[string]*Sala
	candado sync.RWMutex
}

//Regresa el nombre del servidor
func (s *Servidor)getNombre() string {
	return s.nombre
}

func (s *Servidor) getSalas() map[string]*Sala {
	s.candado.RLock()
	defer s.candado.RUnlock()
	return s.salas
}

//Crea el servidor
func NuevoServidor(s string) *Servidor {
	return &Servidor{
		nombre: s,
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
