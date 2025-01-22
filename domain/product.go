package domain

type Product struct{
	id int32 
	name string
	price float32
}

func NewProduct(name string, price float32) *Product{
	return &Product{id:1, name: name, price: price}
} // visualizar como constuctor

func (p*Product)GetName() string{
	return p.name                    // asi se crean metodos nuevos "bindeando" y este es un getter
}

func (p*Product) SetName(name string) {
	p.name = name                     // setter
}

// interfaces para plantillas de datos (revisar reglas del negocio)
// clases para lo mismo pero con metodos 
// repository=
