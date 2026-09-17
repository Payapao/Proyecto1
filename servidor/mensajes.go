package main

import (
	//Para el protocolo
	//"encoding/json"
)

type Mensaje struct{
	Type Tipo `json:"type,omitempty"`
	Roomname string `json:"roomname,omitempty"`
	Username string `json:"username,omitempty"`
	Usernames map[string]*Cliente `json:"usernames,omitempty"`
	Status Estados `json:"status,omitempty"`
	Operation Tipo `json:"operation,omitempty"`
	Result Respuesta `json:"result,omitempty"`
	Text string `json:"text,omitempty"`
	Users map[string]*Cliente `json:"users,omitempty"`
	Extra string `json:"extra,omitempty"`
}

//Funciòn base para crear el mensaje
func FabricaMensaje(tipo Tipo) Mensaje {
	return Mensaje{
		Type: tipo,
	}
}

//Metodos para agregar otros parametros de ser necesario
func (m Mensaje) Roomname(s string) Mensaje{
	m.Roomname = s
	return m
}

func (m Mensaje) Username(s string) Mensaje{
	m.Username = s
	return m
}

func (m Mensaje) Usernames(clientes map[string]*Cliente) Mensaje {
	/*	var s string.Builder
	s.WriteString("[")
	for user , _ := range clientes {
		s.WriteString(`"`)
		s.WriteString(user)
		s.WriteString(`", `)
	}
	s.WriteString("]") */
	m.Usernames = clientes
	return m
}

func (m Mensaje) Status(e Estados) Mensaje {
	m.Status = e
	return m
}

func (m Mensaje) Operation(t Tipo) Mensaje {
	m.Operation = t
	return m
}

func (m Mensaje) Result(r Respuesta) Mensaje {
	m.Result = r
	return m
}

func (m Mensaje) Text(s string) Mensaje {
	m.Text = s
	return m
}

func (m Mensaje) Users(clientes map[string]*Cliente) Mensaje {
	m.Users = clientes
	return m
}

func (m Mensje) Extra(s string) Mensaje {
	m.Extra = s
	return m
}


