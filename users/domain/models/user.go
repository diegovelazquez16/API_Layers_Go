package models

type User struct {
	ID        uint   `gorm:"primaryKey"`
    Nombre    string `gorm:"not null"`
    Email     string `gorm:"unique;not null"`
    Telefono  string
    Direccion string
}
