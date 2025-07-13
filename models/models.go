package models

// Provincia structure
type Provincia struct {
	Id     int    `json:"Id"`
	Nombre string `json:"Nombre"`
}

// Distrito structure
type Distrito struct {
	Id          int    `json:"Id"`
	ProvinciaId int    `json:"ProvinciaId"`
	Nombre      string `json:"Nombre"`
}

// Distrito request structure
type DistritoRequest struct {
	ProvinciaId int `json:"ProvinciaId"`
}

// Corregimiento structure
type Corregimiento struct {
	Id         int    `json:"Id"`
	DistritoId int    `json:"DistritoId"`
	Nombre     string `json:"Nombre"`
}

// Corregimiento request structure
type CorregimientoRequest struct {
	DistritoId int `json:"DistritoId"`
}

// Barrio structure
type Barrio struct {
	Id              int    `json:"Id"`
	CorregimientoId int    `json:"CorregimientoId"`
	Nombre          string `json:"Nombre"`
	FullName        string `json:"FullName"`
}

// Barrio request structure
type BarrioRequest struct {
	CorregimientoId int `json:"CorregimientoId"`
}
