package main

import (
	//Para el protocolo
	//"encoding/json"
)

type Mensaje struct{
	Type int `json: "type"`
	Roomname string `json: roomname`
	Username string `json: username`
	Usernames string `json: usernames`
	Status int `json: status`
	Operation Tipo `json: operation`
	Result int `json: result`
	Text string `json: text`
	Users string `json: users`
	Extra string `json: extra`
}


