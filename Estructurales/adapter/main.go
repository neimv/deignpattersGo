package main

import "fmt"

// Objetivo
type Guerrero interface {
	UsarArma() string
}

type Elfo struct{}

func (e *Elfo) UsarArma() string {
	return "Atacando con arco y flecha"
}

// Adaptable
type GuerreroMagico interface {
	UsarMagia() string
}

type Mago struct{}

func (m *Mago) UsarMagia() string {
	return "atacando con magia"
}

// Adaptador
type MagoAdaptador struct {
	guerrero GuerreroMagico
}

func (ma *MagoAdaptador) UsarArma() string {
	return ma.guerrero.UsarMagia()
}

// Cliente
type Jugador struct {
	guerrero Guerrero
}

func (j *Jugador) Atacar() string {
	return j.guerrero.UsarArma()
}

func main() {
	jugadorA := &Jugador{&Elfo{}}
	fmt.Printf("Jugador A: %s\n", jugadorA.Atacar())

	jugadorB := &Jugador{&MagoAdaptador{&Mago{}}}
	fmt.Printf("Jugador B: %s\n", jugadorB.Atacar())
}
