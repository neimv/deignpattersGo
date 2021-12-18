package main

import (
	"fmt"
	"sync"
	"time"
)

// Singleton
type Singleton struct {
	Tiempo int64
}

// Crear "Estatico"
var instancia *Singleton
var once sync.Once

func GetInstancia() *Singleton {
	once.Do(func() {
		instancia = &Singleton{
			time.Now().Unix(),
		}
	})

	return instancia
}

func main() {
	fmt.Println("Todas las instancias Singleton tienen que tener el mismo número")

	fmt.Printf("Instancia Singleton: %d\n", GetInstancia().Tiempo)

	time.Sleep(1 * time.Second)

	fmt.Printf("Instancia Singleton: %d\n", GetInstancia().Tiempo)

	canalEspera := make(chan int64)

	go func() {
		time.Sleep(1 * time.Second)

		canalEspera <- GetInstancia().Tiempo
	}()

	fmt.Printf("Instancia Singleton: %d\n", <-canalEspera)
}
