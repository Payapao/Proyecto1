package main

import(
	//Sincronización de goroutines
	"sync"
	//Conexión con red
	"net"
	//Manejo de errores
	"errors"
)


//Estructura de las salas
type Sala struct{
	nombre string
	clientes map[string]*Cliente
	//Permite lecturas simultaneas
	candado sync.RWMutex
}

//Getters y setters

func (s Sala) getNombre() string{
	return s.nombre
}

func (s Sala) getClientes() map[string]string{
	//Bloqueamos que se modifique la sala mientras la leemos
	s.candado.RLock()
	defer s.candado.RUnlock()
	//tipo, longitud, capacidad
	arregloClientes := make(map[string]string, len(s.clientes))
	
	for usuario, cliente := range s.clientes {
		//No puede pasar el error
		estado, _ := toStringEstados(cliente.getEstado())
		
		arregloClientes[usuario] = estado

	}
	return arregloClientes
}

func (s *Sala) setNombre(nuevoNombre string) {
	//Bloqueamos la escritura y lectura
	s.candado.Lock()
	defer s.candado.Unlock()
	s.nombre = nuevoNombre
}

//Función para crear nuevos clientes
func (s *Sala) NuevoCliente(nombre string, conexion net.Conn) (*Cliente, error){
	_, existe := s.clientes[nombre]
	if existe{
		return nil, errors.New("El nombre de usuario ya esta ocupado")
	}

	nuevoCliente := &Cliente{
		conexion: conexion,
		usuario: nombre,
		estado: ACTIVE,
	}

	s.clientes[nombre] = nuevoCliente
	return nuevoCliente, nil
}

