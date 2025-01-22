package domain

type IProduct interface {
	Save()
	GetAll()
}

// interface solo puede decir que va a hacer pero no como lo va ha hacer
// todos los metodos de la interfaz
// la logica de como van actuar los metodos de la interfaz van en infra
// patron repostitory separa el modelo de la capa de datos
