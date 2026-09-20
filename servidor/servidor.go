package main

import(
	//Sincronización de goroutines
	"sync"
	//Manejo de errores
	"errors"
	//Conexión en la red
	"net"
	//para usar el paquete de strings
	"strings"
	//Protocolo
	"encoding/json"
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
		usuarios: make(map[string]*Cliente),
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

//Fucnión par eliminar salas, unicamente se llama con salas existentes
func (s *Servidor) EliminaSala(sala *Sala){
	s.candado.Lock()
	defer s.candado.Unlock()

	//Elimina la sala del servidor
	delete(s.salas, sala.getNombre())
	return	
}

//Metodo para agregar un usuario al servidor
func (s *Servidor) NuevoCliente(nombre string, conexion net.Conn, c *json.Encoder, d *json.Decoder) (*Cliente, error) {
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
		codificador: c,
		decodificador: d,
	}
	s.usuarios[nombre] = nuevoCliente
	return nuevoCliente, nil
}

//Metodo para eliminar un cliente
func (s *Servidor) EliminaCliente(cliente *Cliente){
	s.candado.Lock()
	defer s.candado.Unlock()

	//Elimina al usuario del servidor

	//Envia mensje de usuario desconectado a todos los clientes
	for _, c := range s.usuarios {
		c.EnviaMensaje(FabricaMensaje(DISCONNECT).username(cliente.getUsuario()))
	}

	//Verifica si el usuario estaba en alguna sala
	for _, sala := range s.salas{
		if sala.Existe(cliente) {
			//Si era el unico en la sala elimina la sala
			if len(sala.clientes) == 1 {
				s.EliminaSala(sala)
			}else{
				//Si no manda llamar el metodo correspondiente
				sala.EliminaCliente(cliente)
			}
		}
	}
}

func (s *Servidor) Temporal(conexion net.Conn){

	//Decodificador y codificador para que cada cliente pueda leer y escribir en su propio hilo
	decodificador := json.NewDecoder(conexion)
	codificador := json.NewEncoder(conexion)

	//Creamos la variable que va a guardar el json
	var identificacion Mensaje
	err := decodificador.Decode(&identificacion)
	if err != nil{
		//Manda mensaje de Json invalido
		codificador.Encode(Mensaje{Type: RESPONSE, Operation: INVALID_TIPO, Result: INVALID_RESPUESTA})
		//Desconecta al cliente
		conexion.Close()
		return	
	}

	//Nombre que se quiere poner el usuario
	nombre := strings.TrimSpace(identificacion.Username)

	if identificacion.Type != IDENTIFY || nombre == "" || len(nombre) > 8 {
		//Manda mensaje de usuario no identificado
		codificador.Encode(Mensaje{Type: RESPONSE, Operation: INVALID_TIPO, Result: NOT_IDENTIFIED})
		//Desconecta al cliente
		conexion.Close()
		return
	}

	//Manda llamar el metodo NuevoCliente para agregarlo a la lista de usuarios del servidor
	c, e := s.NuevoCliente(nombre, conexion, codificador, decodificador)

	if e != nil {
		//Manda mensaje de usuario repetido
		codificador.Encode(Mensaje{Type: RESPONSE, Operation: IDENTIFY, Result: USER_ALREADY_EXISTS, Extra: nombre })
		//Desconecta al cliente
		conexion.Close()
		return
	}

	mc := FabricaMensaje(RESPONSE).operation(IDENTIFY).result(SUCCESS).extra(nombre)
	mt := FabricaMensaje(NEW_USER).username(nombre)

	s.EnviaTodos(c, mc, mt)

	//Lamar al ciclo infinito del cliente
	s.EscuchaCliente(c)
	
}

//Ciclo para que se escuchen todos los mensajes de un cliente
func (s *Servidor) EscuchaCliente(c *Cliente){
	for{
		var mensaje Mensaje
		err := c.decodificador.Decode(&mensaje)
		if err != nil {
			//Manda mensaje de Json invalido
			c.codificador.Encode(Mensaje{Type: RESPONSE, Operation: INVALID_TIPO, Result: INVALID_RESPUESTA})
			//Desconecta al cliente
			c.conexion.Close()
			return
		}

		s.ProcesaMensaje(mensaje, c)	
	}
	
}

//Define que hace el mensaje que envio el cliente
func (s *Servidor) ProcesaMensaje(m Mensaje, c *Cliente){
	switch m.Type {
	case STATUS:
		//Revisa si el estado es distinto al que tiene el cliente
		if m.Status == c.getEstado(){
			return
		}
		//Actualiza el estado
		c.setEstado(m.Status)
		//No es posible que suceda un error 

		//Fabrica el mensaje
		mt := FabricaMensaje(NEW_STATUS).username(c.getUsuario()).status(m.Status)
		//Envia el mensaje
		s.EnviaTodos(c, nil, mt)
		
	case USERS:
		//Crea el map de los usuarios
		users := make(map[string]string)
		s.candado.RLock()
		for n, cliente := range s.usuarios {
			estado, _ := cliente.getEstado().toStringEstados()
			users[n] = estado
		}
		s.candado.RUnlock()

		//Fabrica el mensaje y lo envia
		mc := FabricaMensaje(USER_LIST).users(users)
		s.EnviaTodos(c, mc, nil)
		
	case TEXT:
		//Checa si el usuario existe en el servidor
		destino , existe := s.usuarios[m.Username]
		if !existe{
			//Si no existe envia el mensaje correspondiente
			mc := FabricaMensaje(RESPONSE).operation(TEXT). result(NO_SUCH_USER).extra(m.Username)
			s.EnviaTodos(c, mc, nil)
			return
		}

		//Fabrica y envia el mensaje al destinatario
		md := FabricaMensaje(TEXT_FROM).username(c.getUsuario()).extra(m.Extra)
		s.EnviaTodos(destino, md, nil)

	case PUBLIC_TEXT:
		//Fabrica y envia el mensaje a todos
		mt := FabricaMensaje(PUBLIC_TEXT_FROM).username(c.getUsuario()).extra(m.Extra)
		s.EnviaTodos(c, nil, mt)

	case NEW_ROOM:
		//Verifica que el nombre de la sala tenga una longitud menor o igual a 16
		if len(m.Roomname) > 16{
			return
		}
		//Crea la sala
		sala, err := s.NuevaSala(m.Roomname)
		//Significa que la sala ya existe y manda el mensaje correspondiente
		if err != nil {
			mc := FabricaMensaje(RESPONSE).operation(NEW_ROOM).result(ROOM_ALREADY_EXISTS).extra(m.Roomname)
			s.EnviaTodos(c, mc, nil)
			return
		}

		//Agrega al cliente
		sala.AgregaCliente(c)
		//Manda el mensaje de confirmación
		mc := FabricaMensaje(RESPONSE).operation(NEW_ROOM).result(SUCCESS).extra(m.Roomname)
		s.EnviaTodos(c, mc, nil)
		
	case INVITE:
		//Verifica
		sala := s.Verifica(c, m, INVITE)
		if sala == nil {
			return
		}
		//Si el cliente no pertenece a la sala ignora el json
		if !sala.Existe(c){
			return
		}

		//Agrega al cliente en la lista de invitados y envia el mensaje
		for _, cliente := range m.Usernames {
			sala.Invita(cliente)
			cliente.EnviaMensaje(FabricaMensaje(INVITATION).username(c.getUsuario()).roomname(m.Roomname))
		}

	case JOIN_ROOM:
		//Verifica que la sala exista
		sala := s.Verifica(c, m, JOIN_ROOM)
		if sala == nil{
			return
		}
		
		//Agrega el usuario a la sala
		sala.AgregaCliente(c)

		//Fabrica y manda el mensaje correspondiente 
		mc := FabricaMensaje(RESPONSE).operation(JOIN_ROOM).result(SUCCESS).extra(m.Roomname)
		mt := FabricaMensaje(JOIN_ROOM).roomname(m.Roomname).username(c.getUsuario())
		sala.EnviaSala(c, mc, mt)

	case ROOM_USERS:
		//Verifica
		sala := s.Verifica(c, m, ROOM_USERS)
		if sala == nil{
			return
		}
		
		//Crea el map de los usuarios
		users := make(map[string]string)
		sala.candado.RLock()
		for n, cliente := range sala.getClientes() {
			estados, _ := cliente.getEstado().toStringEstados()
			users[n] = estados
		}
		sala.candado.RUnlock()

		//Fabrica y envia el mensaje
		mc := FabricaMensaje(ROOM_USERS_LIST).users(users)
		s.EnviaTodos(c, mc, nil)

	case ROOM_TEXT:
		//Verifica
		sala := s.Verifica(c, m, ROOM_TEXT)
		if sala == nil {
			return
		}

		//Fabrica y envia el mensaje a todos en la sala
		mt := FabricaMensaje(ROOM_TEXT_FROM).roomname(m.Roomname).username(c.getUsuario()).extra(m.Extra)
		sala.EnviaSala(c, nil, mt)

	case LEAVE_ROOM:
		//Verifica
		sala := s.Verifica(c, m, LEAVE_ROOM)
		if sala == nil{
			return	
		}

		//Si es el unico usuario en la sala elimina la sala
		if len(sala.getClientes()) == 1 {
			s.EliminaSala(sala)
		}else{
			//Si no, elimina al usuario de la sala
			sala.EliminaCliente(c)
		}
		
		//Por practicidad el mensaje lo envia el metodo EliminaCliente

	case DISCONNECT:
		//Manda el mensaje correspondiente a todos los usuarios
		mt := FabricaMensaje(DISCONNECTED).username(c.getUsuario())
		s.EnviaTodos(c, nil, mt)

		//Elimina al cliente de las salas
		s.candado.RLock()
		for _, sala := range s.salas {
			if len(sala.getClientes()) == 1 {
				s.EliminaSala(sala)
			} 
			sala.EliminaCliente(c)
		}
		s.candado.RUnlock()
		
		//Elimina al cliente del servidor
		s.EliminaCliente(c)
		
		//Desconecta al usuario
		c.conexion.Close()
		
	default:
		//Nunca deberia pasar este default
		return
	}
	return
}

//Verifica que el usuario pueda hacer lo que quiere hacer en las salas
func (s *Servidor) Verifica(c *Cliente, m Mensaje, tipo Tipo) *Sala {
	//Verifica que la sala exista
	sala, existe := s.salas[m.Roomname]
	if !existe{
		mc := FabricaMensaje(RESPONSE).operation(tipo).result(NO_SUCH_ROOM).extra(m.Roomname)
		s.EnviaTodos(c, mc, nil)
		return nil
	}
	switch tipo {
	case LEAVE_ROOM, ROOM_TEXT, ROOM_USERS:
		//Verifica que el usuario este en el cuarto
		if !sala.Existe(c) {
			mc := FabricaMensaje(RESPONSE).operation(tipo).result(NOT_JOINED).extra(m.Roomname)
			s.EnviaTodos(c, mc, nil)
			return nil
		}
	case JOIN_ROOM:
		//Verifica que el usuario este invitado en el cuarto
		if !sala.Invitado(c) {
			mc := FabricaMensaje(RESPONSE).operation(tipo).result(NOT_INVITED).extra(m.Roomname)
			s.EnviaTodos(c, mc, nil)
			return nil
		}
	case INVITE:
		for invitado, _ := range m.Usernames {
			_, e := s.usuarios[invitado]
			if ! e {
				mc := FabricaMensaje(RESPONSE).operation(INVITE).result(NO_SUCH_USER).extra(invitado)
				s.EnviaTodos(c, mc, nil)
				return nil
			}	
		}
		
	}
	return sala
}

func (s *Servidor) EnviaTodos(c *Cliente, mc *Mensaje, mt *Mensaje){

	//Si no hay mensaje al propio usuario
	if mc == nil{
		s.candado.RLock()
		for usuario, cliente := range s.usuarios {
			if usuario != c.getUsuario() {
				cliente.EnviaMensaje(mt)
			}
		}
		s.candado.RUnlock()
	}else if mt == nil {
		//Si no hay mensaje para el resto de usuarios
		c.EnviaMensaje(mc)
	}else {
		//Hay mensaje para ambos
		s.candado.RLock()
		for usuario, cliente := range s.usuarios {
			if usuario == c.getUsuario() {
				cliente.EnviaMensaje(mc)
			}else{
				cliente.EnviaMensaje(mt)
			}
		}
		s.candado.RUnlock()
	}
}
