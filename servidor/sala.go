package main

import(
	//Sincronización de goroutines
	"sync"
	//Manejo de errores
	"errors"
)


//Estructura de las salas
type Sala struct{
	//Nombre de la sala
	nombre string
	//Clientes en la sala
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
		estado, _ := cliente.getEstado().toStringEstados()
		
		arregloClientes[usuario] = estado

	}
	return arregloClientes
}

//Nos dice si un usuario se encuentra en la sala
func (s *Sala) Existe(cliente *Cliente) bool {
	clientes := s.getClientes()
	_, existe := clientes[cliente.getUsuario()]
	if existe {
		return true
	}
	return false
}

//Función para agregar clientes a la sala -Es necesario pasar un usuario valido
func (s *Sala) AgregaCliente(cliente *Cliente) (error){
	s.candado.Lock()
	defer s.candado.Unlock()

	//Verifica si el usuario ya esta en la sala y lo ignora
	if s.Existe(cliente){
		return nil
	}

	//Agrega al usuario en la sala
	s.clientes[cliente.getUsuario()] = cliente
	return nil
}

//Función para eliminar un cliente de una sala
func (s *Sala) EliminaCliente(cliente *Cliente){
	s.candado.Lock()
	defer s.candado.Unlock()
	
	//Elimina al usuario de la sala

	//Envia mensaje de que el usuario ah dejado la sala
	for _ , cliente := range s.clientes {
		cliente.EnviaMensaje()
	}
}

