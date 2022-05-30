package models

type Transaccion struct {
	ID              int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nro_transaccion string `json:"nro_transaccion"`
	Fecha           string `json:"fecha"`
	Descripcion     string `json:"descripcion"`
	Status          int    `json:"status"`
}
