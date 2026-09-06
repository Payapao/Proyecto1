package main

import(
	"testing"
	"regexp"
)

//Prueba para la funcion saludo
func TestSaludoName(t *testing.T){
	name := "Vania"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Saludo("Vania")
	if !want.MatchString(msg) || err != nil{
		t.Errorf(`Saludo("Vania") = %q, %v, no coincide con %#q, nil`, msg, err, want)
	}
	
}

//Prueba para la funcion saludo en estado de error
func TestSaludoVacia(t *testing.T){
	msg, err := Saludo("")
	if msg != "" || err == nil{
		t.Errorf(`Saludo("") = %q, %v, no coincide con "", error`, msg, err)
	}
}
