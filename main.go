package main

import (
	"holamundo/aplication"
	"holamundo/infraestructure"
)

func main() {
	mysql := infraestructure.NewMySQL()
	cp := aplication.NewCreateProduct(mysql)
	cp.Run()
}

//:= operador de referencia
//minuscula =>privado, Mayuscula => publico