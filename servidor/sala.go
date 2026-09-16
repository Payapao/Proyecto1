package main

import(
	//Sincronización de goroutines
	"sync"
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

//Función para agregar clientes a la sala
func (s *Sala) AgregaCliente(cliente *Cliente) (error){
	s.candado.Lock()
	defer s.candado.Unlock()

	usuarios := servidor.getUsuarios()

	_, existe := usuarios[cliente.getUsuario()]
	
	if !existe{
		return errors.New("El usuario no se encuentra en el servidor")
	}


	s.clientes[cliente.getUsuario()] = cliente
	return nil
}

