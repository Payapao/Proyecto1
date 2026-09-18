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
		codificador.Encode(Mensaje{Type: RESPONSE, Operation: INVALID, Result: NOT_IDENTIFIED})
		//Desconecta al cliente
		conexion.Close()
		return	
	}

	//Nombre que se quiere poner el usuario
	nombre := strings.TrimSpace(identificacion.Username)

	if identificacion.Type != IDENTIFY || nombre == "" || len(nombre) > 8 {
		//Manda mensaje de usuario no identificado
		codificador.Encode(Mensaje{Type: RESPONSE, Operation: INVALID, Result: NOT_IDENTIFIED})
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

	s.candado.RLock()
	for usuario, cliente := range s.usuarios {
		if usuario != nombre {
			//Envia mensaje de nuevo usuario
			cliente.EnviaMensaje(FabricaMensaje(NEW_USER).username(nombre))
		}else {
			//Envia mensaje de identificación valida
			cliente.EnviaMensaje(FabricaMensaje(RESPONSE).operation(IDENTIFY).result(SUCCESS).extra(nombre))
		}
	}
	s.candado.RUnlock()

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
			c.codificador.Encode(Mensaje{Type: RESPONSE, Operation: INVALID, Result: NOT_IDENTIFIED})
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
		//Actualiza el estado
		errEstado := c.setEstado(m.Status)
		//Checar si el estado es valido --Esto no sale en el protocolo y talvez sea un error
		//Si es un error entonces regresar status a Estado y cambiar set estado para que reciba estados en vez de strings
		if errEstado != nil{
			c.EnviaMensaje(FabricaMensaje(RESPONSE).operation(STATUS).result(NOT_IDENTIFIED))
		}
		
		//Manda el mensaje valido para el cliente
		//Manda la actualización para el resto de clientes
		//Deberia hacer un metodo envia para no repetir el for?
	default:
		//Nunca deberia pasar este default
		return
	}
	return
}
