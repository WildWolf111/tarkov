package models

type Warehouses_cells struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Warehouse_id int    `json:"warehouse_id"`
}
