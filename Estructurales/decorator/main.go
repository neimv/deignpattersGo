package main

import "fmt"

// Componente interface
type CafeInterface interface {
	GetCosto() int
	GetDetalle() string
}

// Componente concreto
type Cafe struct{}

func (c *Cafe) GetCosto() int {
	return 30
}

func (c *Cafe) GetDetalle() string {
	return "Cafe"
}

// Decorador
type CafeDecorador struct {
	CafeInterface
}

func (cd *CafeDecorador) GetCosto() int {
	return cd.CafeInterface.GetCosto()
}

func (cd *CafeDecorador) GetDetalle() string {
	return cd.CafeInterface.GetDetalle()
}

// Decorador concreto
type CafeConCrema struct {
	*CafeDecorador
}

func (ccc *CafeConCrema) GetCosto() int {
	return ccc.CafeDecorador.GetCosto() + 15
}

func (ccc *CafeConCrema) GetDetalle() string {
	return ccc.CafeDecorador.GetDetalle() + " con crema"
}

// Decorador concreto
type CafeConCanela struct {
	*CafeDecorador
}

func (ccc *CafeConCanela) GetCosto() int {
	return ccc.CafeDecorador.GetCosto() + 10
}

func (ccc *CafeConCanela) GetDetalle() string {
	return ccc.CafeDecorador.GetDetalle() + " con canela"
}

func main() {
	cafe := &Cafe{}
	fmt.Printf("Detalle: %s - Importe $%d\n", cafe.GetDetalle(), cafe.GetCosto())

	cafeConCrema := &CafeConCrema{&CafeDecorador{cafe}}
	fmt.Printf("Detalle: %s - Importe $%d\n", cafeConCrema.GetDetalle(), cafeConCrema.GetCosto())

	cafeConCremaConCanela := &CafeConCanela{&CafeDecorador{cafeConCrema}}
	fmt.Printf("Detalle: %s - Importe $%d\n", cafeConCremaConCanela.GetDetalle(), cafeConCremaConCanela.GetCosto())
}
