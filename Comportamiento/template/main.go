package main

import "fmt"

// Clase abstracta - interface
type DeployInterface interface {
	Testear()
	Compilar()
	Publicar()
}

// Clase abstracta
type Deploy struct{}

// Metodo plantilla
func (d Deploy) Construir(di DeployInterface) {
	fmt.Println("Ejecutando las siguientes acciones: ")

	di.Testear()
	di.Compilar()
	di.Publicar()
}

// Clase concreta - android
type DeployAndroid struct {
	Deploy
}

func (d DeployAndroid) Testear() {
	fmt.Println("Android: Testeando")
}

func (d DeployAndroid) Compilar() {
	fmt.Println("Android: Compilando")
}

func (d DeployAndroid) Publicar() {
	fmt.Println("Android: Publicando")
}

// Clase concreta - iOS
type DeployiOS struct {
	Deploy
}

func (d DeployiOS) Testear() {
	fmt.Println("iOS: Testeando")
}

func (d DeployiOS) Compilar() {
	fmt.Println("iOS: Compilando")
}

func (d DeployiOS) Publicar() {
	fmt.Println("iOS: Publicando")
}

func main() {
	deployAndroid := DeployAndroid{Deploy{}}
	deployAndroid.Construir(&deployAndroid)

	deployiOS := DeployiOS{Deploy{}}
	deployiOS.Construir(&deployiOS)
}
