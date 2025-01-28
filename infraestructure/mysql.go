package infraestructure

import "fmt"

// El nombre de este archivo corresponde a la tecnologia que se esta usando
// Puertos & adaptadores: ej. varias tecnologias (adaptadores) pueden estar conectadas a una sola interfaz (puerto)
// Handler: manipulador/funcion que se solo ejecuta cuando se solicita
// Aqui simplemente vas a configurar para imprimir mensajes cuando realizaces operaciones crud
type MySQL struct {}

func NewMySQL () * MySQL {
	return & MySQL{}
}

func (mysql *MySQL) Save(){
	fmt.Println("Producto salvado")

}

func (mysql *MySQL) GetAll(){
	fmt.Print("Lista de productos")
}

