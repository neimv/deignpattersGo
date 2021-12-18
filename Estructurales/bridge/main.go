package main

import "fmt"

// Abstraccion interface
type DispositivoInterface interface {
	ConectarInternet() string
	SetConexion(Conexion)
}

// Abstraccion abstracta
type Dispositivo struct {
	conexion Conexion
}

func (d *Dispositivo) SetConexion(conexion Conexion) {
	d.conexion = conexion
}

// Abstraccion refinada
type Telefono struct {
	numero string
	*Dispositivo
}

func (t *Telefono) ConectarInternet() string {
	return "Teléfono N° " + t.numero + " conectado a internet mediante " + t.conexion.Conectar()
}

// Abstraccion refinada
type Tablet struct {
	*Dispositivo
}

func (t *Tablet) ConectarInternet() string {
	return "Tablet conectada a internet mediante " + t.conexion.Conectar()
}

// Implementador interface
type Conexion interface {
	Conectar() string
}

// Implementar concreto
type Red4G struct{}

func (r *Red4G) Conectar() string {
	return "4G"
}

// Implementar concreto
type RedWiFi struct{}

func (r *RedWiFi) Conectar() string {
	return "WiFi"
}

func main() {
	telefonoA := &Telefono{"0115161", &Dispositivo{}}
	telefonoA.SetConexion(&Red4G{})
	fmt.Printf("%s\n", telefonoA.ConectarInternet())

	telefonoB := &Telefono{"0117854", &Dispositivo{}}
	telefonoB.SetConexion(&RedWiFi{})
	fmt.Printf("%s\n", telefonoB.ConectarInternet())

	tablet := &Tablet{&Dispositivo{}}
	tablet.SetConexion(&RedWiFi{})
	fmt.Printf("%s\n", tablet.ConectarInternet())
}
