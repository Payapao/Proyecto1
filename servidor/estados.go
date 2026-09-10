package estado

//Enum para go
type Estados int

//Estados validos para los clientes
const{
	ACTIVE Estados = iota
	AWAY //Estados = iota
	BUSY //Puede omitirse el = iota para los siguientes elementos
}

//Funcion to String para los estados
func toString(e Estado)(string, error){
	s := ""
	switch e {
	case "ACTIVE":
		s = "ACTIVE"
	case "AWAY":
		s = "AWAY"
	case "BUSY":
		s = "BUSY"
	default:
		return s , errors.New("El estado no es valido")
	}
	return s, nil
}
