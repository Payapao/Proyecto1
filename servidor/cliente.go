package main

import(
	//Sincronización de goroutines
	"sync"
	//manejo de errores
	"errors"
	//Comunicación en red
	"net"
	//Protocolo
	"encoding/json"
)

//Creamos la estructura del cliente
type Cliente struct{
	conexion net.Conn
	usuario string
	estado Estados
	//Para que cada cliente envie sus propios mensajes
	codificador *json.Encoder
	decodificador *json.Decoder
	candado sync.RWMutex
}

//Setters y getters para clientes

func (c Cliente) getConexion() net.Conn{
	return c.conexion
}

func (c Cliente) getUsuario() string{
	return c.usuario
}

func (c Cliente) getEstado() Estados{
	return c.estado
}

//Define el estado de un cliente
func (c *Cliente) setEstado(s Estados) error{
	switch s {
	case ACTIVE:
		c.estado = ACTIVE
	case AWAY:
		c.estado = AWAY
	case BUSY:
		c.estado = BUSY
	default:
		return errors.New("El estado no es valido")
	}
	return nil
}

//Envia mensaje al cliente
func (c *Cliente) EnviaMensaje(mensaje *Mensaje) {
	c.candado.Lock()
	defer c.candado.Unlock()

	err := c.codificador.Encode(mensaje)
	if err != nil {
		c.conexion.Close()
	}
	
}
