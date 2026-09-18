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
	Status string `json:"status,omitempty"`
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
func (m Mensaje) roomname(s string) Mensaje{
	m.Roomname = s
	return m
}

func (m Mensaje) username(s string) Mensaje{
	m.Username = s
	return m
}

func (m Mensaje) usernames(clientes map[string]*Cliente) Mensaje {
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

func (m Mensaje) status(e Estados) Mensaje {
	m.Status, _ = e.toStringEstados()
	return m
}

func (m Mensaje) operation(t Tipo) Mensaje {
	m.Operation = t
	return m
}

func (m Mensaje) result(r Respuesta) Mensaje {
	m.Result = r
	return m
}

func (m Mensaje) text(s string) Mensaje {
	m.Text = s
	return m
}

func (m Mensaje) users(clientes map[string]*Cliente) Mensaje {
	m.Users = clientes
	return m
}

func (m Mensaje) extra(s string) Mensaje {
	m.Extra = s
	return m
}


