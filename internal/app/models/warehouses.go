package models

//User model ...
type Warehouses struct {
	ID         int          `json:"id"`
	Name       string       `json:"login"`
	Slug       string       `json:"slug"`
	Company_id int          `json:"company_id"`
	Address    string       `json:"address"`
	Companies  []*Companies `json:"companies"`
}
