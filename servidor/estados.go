package estado

//Enum para go
type Estados int

//Estados validos para los clientes
const{
	ACTIVE Estados = iota
	AWAY //Estados = iota
	BUSY //Puede omitirse el = iota para los siguientes elementos
}
