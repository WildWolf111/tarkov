package models

//Article models...
type Companies struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	INN  uint64 `json:"inn"`
	KPP  uint64 `json:"kpp"`
}
