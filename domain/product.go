package domain

type Product struct {
	ID    int32   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func NewProduct(name string, price float32) *Product {
	return &Product{Name: name, Price: price}
}


// interfaces para plantillas de datos (revisar reglas del negocio)
// clases para lo mismo pero con metodos 
// repository=
