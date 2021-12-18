package main

import "fmt"

// Interface producto
type Entrevistador interface {
	RealizarPreguntas()
}

// Producto Concreto
type EntrevistadorIT struct{}

func (e *EntrevistadorIT) RealizarPreguntas() {
	fmt.Println("Nombre 5 patrones de dise;o?")
}

// Producto Concreto
type EntrevistadorFinanzas struct{}

func (e *EntrevistadorFinanzas) RealizarPreguntas() {
	fmt.Println("Cual es la alicouta del IVA?")
}

// creador interface
type RecursosHumanosInterface interface {
	LLamarEntrevistador() Entrevistador
}

// Creador abstracto
type RecursosHumanos struct{}

func (rh *RecursosHumanos) TomarEntrevista(self RecursosHumanosInterface) {
	entrevistador := self.LLamarEntrevistador()
	entrevistador.RealizarPreguntas()
}

// Creador concreto
type RecursosHumanosIT struct {
	*RecursosHumanos
}

func (rhi *RecursosHumanosIT) LLamarEntrevistador() Entrevistador {
	return &EntrevistadorIT{}
}

// Creador Concreto
type RecursosHumanosFinanzas struct {
	*RecursosHumanos
}

func (rhf *RecursosHumanosFinanzas) LLamarEntrevistador() Entrevistador {
	return &EntrevistadorFinanzas{}
}

func main() {
	fmt.Println("El entrevisatador de IT pregunta:")
	recursosHumanosIT := &RecursosHumanosIT{&RecursosHumanos{}}
	recursosHumanosIT.TomarEntrevista(recursosHumanosIT)

	fmt.Println("\nEl entrevisatador de Finanzas pregunta:")
	recursosHumanosFinanzas := &RecursosHumanosFinanzas{&RecursosHumanos{}}
	recursosHumanosFinanzas.TomarEntrevista(recursosHumanosFinanzas)
}
