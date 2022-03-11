package models

//Article models...
type Company struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	INN  int    `json:"inn"`
	KPP  int    `json:"kpp"`
}
