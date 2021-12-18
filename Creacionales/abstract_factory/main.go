package main

import "fmt"

// Producto Abstracto interface
type Puerta interface {
	VerMaterial() string
}

// Producto concreto
type PuertaMadera struct{}

func (pm *PuertaMadera) VerMaterial() string {
	return "Madera"
}

// Producto concreto
type PuertaMetal struct{}

func (pm *PuertaMetal) VerMaterial() string {
	return "Metal"
}

// Fabrica abstracta interface
type FabricaPuerta interface {
	ConstruirPuerta() Puerta
}

// fabrica concreta
type FabricaPuertaMadera struct{}

func (fpm *FabricaPuertaMadera) ConstruirPuerta() Puerta {
	return &PuertaMadera{}
}

// fabrica concreta
type FabricaPuertaMetal struct{}

func (fpm *FabricaPuertaMetal) ConstruirPuerta() Puerta {
	return &PuertaMetal{}
}

func main() {
	fabricaPuertaMadera := &FabricaPuertaMadera{}
	puertaMadera := fabricaPuertaMadera.ConstruirPuerta()
	fmt.Printf("Se construyo un puerta de: %s\n", puertaMadera.VerMaterial())

	fabricaPuertaMetal := &FabricaPuertaMetal{}
	puertaMetal := fabricaPuertaMetal.ConstruirPuerta()
	fmt.Printf("Se construyo un puerta de: %s\n", puertaMetal.VerMaterial())
}
