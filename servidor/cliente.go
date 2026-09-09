package main

import{
	//manejo de errores
	"errors"
	
}

//Creamos la estructura del cliente
type Cliente struct{
	conexion net.Conn
	usuario string
	estado Estados	
}

//Setters y getters para clientes

func (c Cliente) getConexion() net.Conn{
	return c.conexion
}

func (c Cliente) getUsuario() string{
	return c.usuario
}

func (c Cliente) getEstado() estados{
	return c.estado
}

func (c *Cliente) setEstado(s string) error{
	switch s {
	case "ACTIVE":
		c.estado* = ACTIVE
	case "AWAY":
		c.estado* = AWAY
	case "BUSY":
		c.estado* = BUSY
	default:
		return errors.New("El estado no es valido")
	}
	return nil
}

//Función para crear nuevos clientes

func NuevoCliente(nombre string, conexion net.Conn) (*Cliente, error){
	_, existe := servidor[nombre]
	if existe{
		return nil, errors.New("El nombre de usuario ya esta ocupado")
	}

	return &Cliente{
		conexion: conexion,
		nombre: nombre,
		estado: ACTIVE
	}, nil
}


