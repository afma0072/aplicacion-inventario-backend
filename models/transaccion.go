package models

type Transaccion struct {
	ID              int64  `json:"id" gorm:"primary_key;auto_increment"`
	Tipo            string `json:"tipo"`
	Id_producto     int    `json:"id_producto"`
	Cantidad        int    `json:"cantidad"`
	Nro_transaccion int    `json:"nro_transaccion" gorm:"auto_increment"`
	Fecha           string `json:"fecha"`
	Descripcion     string `json:"descripcion"`
	Status          int    `json:"status"`
}
