package main

import "fmt"

// Iterator Interface
type Iterator interface {
	Valor() string
	Siguiente()
	Anterior()
}

// Agregado Interface
type Agregado interface {
	CrearIterador() Iterator
}

// Agregado Concreto
type Radio struct {
	Estaciones []string
}

func (r *Radio) CrearIterador() Iterator {
	return &RadioIterator{radio: r}
}

func (r *Radio) Registrar(estacion string) {
	r.Estaciones = append(r.Estaciones, estacion)
}

// Iterator Concreto
type RadioIterator struct {
	radio  *Radio
	indice int
}

func (ri *RadioIterator) Valor() string {
	return ri.radio.Estaciones[ri.indice]
}

func (ri *RadioIterator) Siguiente() {
	ri.indice++
}

func (ri *RadioIterator) Anterior() {
	ri.indice--
}

func main() {
	radio := &Radio{}
	radio.Registrar("FM100")
	radio.Registrar("FM200")
	radio.Registrar("FM300")

	iterador := radio.CrearIterador()

	fmt.Printf("Escuhando la radio %s\n", iterador.Valor())

	iterador.Siguiente()
	fmt.Printf("Escuhando la radio %s\n", iterador.Valor())

	iterador.Siguiente()
	fmt.Printf("Escuhando la radio %s\n", iterador.Valor())

	iterador.Anterior()
	fmt.Printf("Escuhando nuevamente la radio %s\n", iterador.Valor())
}
