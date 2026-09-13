package main


//Estructura que contenga todas las salas
type Servidor struct{
	Salas map[string]*Sala
	candado sync.Mutex
}

//Crea el servidor
func NuevoServidor(nombre string) *Servidor{
	return &Servidor{
		salas: make(map[string]*Sala),
	}
}

//Función para crear y agregar salas al servidor
func AgregaSala(nombre string) error{
	candado.Lock()
	defer candado.Unlock()

	_, existe := salas[nombre]
	if existe{
		return nil, errors.New("El nombre de la sala ya existe")
	}

	salas[nombre] = &Sala{
		nombre: nombre,
		clientes: make(map[net.Conn]*Cliente),
	}
	
	return nil
	
}
