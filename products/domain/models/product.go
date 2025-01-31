package models


type Product struct {
	ID    uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string  `json:"name" gorm:"not null"`
	Price float64 `json:"price" gorm:"not null"`
}

// visualizar como constuctor



// interfaces para plantillas de datos (revisar reglas del negocio)
// clases para lo mismo pero con metodos 
// repository=