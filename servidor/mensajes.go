package main

import (
	//Para el protocolo
	//"encoding/json"
)

type Mensaje struct{
	Type Tipo `json:"type,omitempty"`
	Roomname string `json:"roomname,omitempty"`
	Username string `json:"username,omitempty"`
	Usernames string `json:"usernames,omitempty"`
	Status Estados `json:"status,omitempty"`
	Operation Tipo `json:"operation,omitempty"`
	Result Respuesta `json:"result,omitempty"`
	Text string `json:"text,omitempty"`
	Users string `json:"users,omitempty"`
	Extra string `json:"extra,omitempty"`
}


