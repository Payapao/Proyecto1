package main

import {
	"encoding/json"
}

type Mensaje struct{
	Type Tipo `json: "type"`
	Roomname string `json: roomname`
	Username string `json: username`
	Usernames string `json: usernames`
	Status Estados `json: status`
	Operation Tipo `json: operation`
	Result Respuesta `json: result`
	Text string `json: text`
	Users string `json: users`
	Extra string `json: extra`
}
