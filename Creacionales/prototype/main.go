package main

import "fmt"

//Prototype interface
type Prototipo interface {
	Clonar() *Elemento
}

// Prototipo concreto
type Elemento struct {
	Material string
	Copias   int
}

func (e *Elemento) Clonar() *Elemento {
	return &Elemento{
		Material: e.Material,
		Copias:   e.Copias + 1,
	}
}

func main() {
	elementoA := &Elemento{"Azufre", 1}

	elementoB := elementoA.Clonar()
	elementoB.Material = elementoB.Material + " (fortificado)"

	elementoC := elementoB.Clonar()

	fmt.Printf("El elemento A es de %s y se clono %d veces\n", elementoA.Material, elementoA.Copias)
	fmt.Printf("El elemento B es de %s y se clono %d veces\n", elementoB.Material, elementoB.Copias)
	fmt.Printf("El elemento C es de %s y se clono %d veces\n", elementoC.Material, elementoC.Copias)
}
