package main

import(
	//Sincronización de goroutines
	"sync"
	//Manejo de errores
	"errors"
	//Conexión en la red
	"net"
)

//Estructura que contenga todas las salas
type Servidor struct{
	nombre string
	salas map[string]*Sala
	usuarios map[string]*Cliente
	candado sync.RWMutex
}

//Regresa el nombre del servidor
func (s *Servidor) getNombre() string {
	return s.nombre
}

func (s *Servidor) getSalas() map[string]*Sala {
	s.candado.RLock()
	defer s.candado.RUnlock()
	return s.salas
}

func (s *Servidor) getUsuarios() map[string]*Cliente {
	s.candado.RLock()
	defer s.candado.RUnlock()
	return s.usuarios
}


//Crea el servidor
func NuevoServidor(s string) *Servidor {
	return &Servidor{
		nombre: s,
		salas: make(map[string]*Sala),
	}
}

//Metodo para crear y agregar salas al servidor
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

//Metodo para agregar un usuario al servidor
func (s *Servidor) NuevoCliente(nombre string, conexion net.Conn) (*Cliente, error) {
	s.candado.Lock()
	defer s.candado.Unlock()

	_, existe := s.usuarios[nombre]
	if existe{
		return nil, errors.New("El nombre de usuario ya existe")
	}

	nuevoCliente := &Cliente{
		conexion: conexion,
		usuario: nombre,
		estado: ACTIVE,
	}
	s.usuarios[nombre] = nuevoCliente
	return nuevoCliente, nil
}

func (s *Servidor) Temporal(conexion net.Conn){
	//Asegura desconectar al cliente
	defer conexion.Close()

	//Recibe el json de identificación

	//Manda llamar el metodo NuevoCliente para agregarlo a la lista de usuarios del servidor
	
}
