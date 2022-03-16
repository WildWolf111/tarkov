package store

import (
	"fmt"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
)

type Warehouses_CellsRepository struct {
	store *Store
}

var (
	tablewarehouses_cells string = "warehouses_cells"
)

//For Post request
func (warc *WarehouseRepository) Post(a *models.Warehouses_cells) (*models.Warehouses_cells, error) {
	query := fmt.Sprintf("INSERT INTO %s (id, name, slug, company_id) VALUES ($1, $2, $3, $4) RETURNING id", tablewarehouses_cells)
	if err := warc.store.db.QueryRow(query, a.Name, a.Slug, a.Warehouse_id).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}
