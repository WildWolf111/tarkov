package models

//User model ...
type Warehouses struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Company_id int
	Address    string `json:"address"`
}
