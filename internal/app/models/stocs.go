package models

//Article models...
type Stocs struct {
	ID                   int         `json:"id"`
	Company_sender_id    int         `json:"company_sender_id"`
	Company_recipient_id int         `json:"company_recipient_id"`
	Product_id           int         `json:"product_id"`
	Quantity             string      `json:"quantity"`
	Warehouse_cell_id    int         `json:"warehouse_cell_id"`
	GTD_id               int         `json:"gtd_id"`
	Companies            []Companies `json:"companies"`
}
