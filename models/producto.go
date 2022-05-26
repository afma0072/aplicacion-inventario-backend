package models

type Producto struct {
	ID          int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nombre      string `json:"nombre"`
	Existencia  string `json:"existencia"`
	Precio      string `json:"precio"`
	Status      string `json:"status"`
	Id_producto string `json:"id_producto"`
}
