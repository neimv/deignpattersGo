package main

import "fmt"

// Interface Sujeto
type Sujeto interface {
	Adquirir(Observador)
	notificar()
}

// Sujeto concreto
type AgenciaEmpleo struct {
	observadores []Observador
}

func (ae *AgenciaEmpleo) Adquirir(observador Observador) {
	ae.observadores = append(ae.observadores, observador)
}

func (ae *AgenciaEmpleo) notificar(oferta string) {
	for _, observador := range ae.observadores {
		observador.Actualizar(oferta)
	}
}

func (ae *AgenciaEmpleo) IngresarOfertaLaboral(oferta string) {
	ae.notificar(oferta)
}

// Interface Observador
type Observador interface {
	Actualizar(string)
}

// Observador concreto
type ObservadorEmpleo struct {
	nombre string
}

func (o *ObservadorEmpleo) Actualizar(oferta string) {
	fmt.Printf("Hola %s, existe la siguiente oferta de empleo: %s\n", o.nombre, oferta)
}

func main() {
	observadorA := &ObservadorEmpleo{"Juan"}
	observadorB := &ObservadorEmpleo{"Maria"}

	agencia := &AgenciaEmpleo{}
	agencia.Adquirir(observadorA)
	agencia.Adquirir(observadorB)

	agencia.IngresarOfertaLaboral("Programador JAVA Senior")
	fmt.Printf("\n")
	agencia.IngresarOfertaLaboral("Programador C# Junior")
}
