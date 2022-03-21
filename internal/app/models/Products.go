package models

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Sku       string `json:"sku"`
	ShortDesc string `json:"short_desc"`
	FullDesc  string `json:"full_desc"`
	Sort      int    `json:"sort"`
}
