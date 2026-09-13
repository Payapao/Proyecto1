package main


//Estructura de las salas
type Sala struct{
	nombre string
	clientes map[string]*Cliente
	candado sync.Mutex
}

//Getters y setters

func (s Sala) getNombre() string{
	return s.nombre
}

func (s Sala) getClientes() []string{
	//tipo, longitud, capacidad
	clientes := make([]string, 0, len(s))
	for cliente := range s{
		clientes = append(clientes, cliente) 
	}
	return clientes
}

func (s *Sala) setNombre(nuevoNombre string) {
	s.nombre = nuevoNombre
}

