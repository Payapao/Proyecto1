package main

import(
	//Sincronización de goroutines
	"sync"
)


//Estructura de las salas
type Sala struct{
	//Nombre de la sala
	nombre string
	//Clientes en la sala
	clientes map[string]*Cliente
	//Clientes invitados a la sala
	invitados map[string]*Cliente
	//Permite lecturas simultaneas
	candado sync.RWMutex
}

//Getters y setters

func (s Sala) getNombre() string{
	return s.nombre
}

func (s Sala) getClientes() map[string]*Cliente{
	s.candado.RLock();
	defer s.candado.RUnlock()

	copia := make(map[string]*Cliente)
	for nombre, cliente := range s.clientes{
		copia[nombre] = cliente
	}
	return copia
}

func (s Sala) getInvitados() map[string]*Cliente{
	return s.invitados
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

//Nos dice si un usuario esta invitado en la sala
func (s *Sala) Invitado(cliente *Cliente) bool {
	invitados := s.getInvitados()
	_, existe := invitados[cliente.getUsuario()]
	if existe {
		return true
	}
	return false
}

//Para invitar a los usuarios a la sala
func (s *Sala) Invita(c *Cliente) bool {
	if s.Invitado(c) || s.Existe(c){
		return false
	}

	s.candado.Lock()
	defer s.candado.Unlock()
	
	s.invitados[c.getUsuario()] = c
	return true
}

//Función para agregar clientes a la sala -Es necesario pasar un usuario valido
func (s *Sala) AgregaCliente(cliente *Cliente) {
	s.candado.Lock()
	defer s.candado.Unlock()

	//Elimina al cliente de la lista de invitados
	delete(s.invitados, cliente.getUsuario())

	//Agrega al usuario en la sala
	s.clientes[cliente.getUsuario()] = cliente
	return 
}

//Función para eliminar un cliente de una sala
func (s *Sala) EliminaCliente(cliente *Cliente){
	if !s.Existe(cliente){
		return
	}
	s.candado.Lock()
	
	//Elimina al usuario de la sala
	delete(s.clientes, cliente.getUsuario())

	s.candado.Unlock()

	//Envia el mensaje de usuario eliminado al resto de la sala
	mt := FabricaMensaje(LEFT_ROOM).roomname(s.getNombre()).username(cliente.getUsuario())
	s.EnviaSala(cliente, nil, mt)
	
	return
}

//Envia el mensaje a todos los miembros de la sala
func (s *Sala) EnviaSala(c *Cliente, mc *Mensaje, mt *Mensaje){
	s.candado.RLock()
	//Si no hay mensaje al propio usuario
	if mc == nil{
		for usuario, cliente := range s.clientes {
			if usuario != c.getUsuario() {
				cliente.EnviaMensaje(mt)
			}
		}
	}else{
		//Hay mensaje para ambos
		for usuario, cliente := range s.clientes {
			if usuario == c.getUsuario() {
				cliente.EnviaMensaje(mc)
			}else{
				cliente.EnviaMensaje(mt)
			}
		}
	}
	s.candado.RUnlock()
}
