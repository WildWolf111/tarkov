package store

import (
	"fmt"
	"log"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
)

type Warehouses_CellsRepository struct {
	store *Store
}

var (
	tablewarehouses_cells string = "warehouses_cells"
)

//For Post request
func (warc *Warehouses_CellsRepository) Post(a *models.Warehouses_cells) (*models.Warehouses_cells, error) {
	query := fmt.Sprintf("INSERT INTO %s ( name, slug, warehouses_id) VALUES ( $1, $2, $3) RETURNING id", tablewarehouses_cells)
	log.Println(query)
	if err := warc.store.db.QueryRow(query, a.Name, a.Slug, a.Warehouse_id).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For Update request
func (warc *Warehouses_CellsRepository) UpdateWarehouses_cells(a *models.Warehouses_cells) (*models.Warehouses_cells, error) {
	query := fmt.Sprintf("UPDATE %s SET ( name, slug, warehouses_id) VALUES ($2, $3,$4) WHERE id=$1 RETURNING id", tablewarehouses_cells)
	if err := warc.store.db.QueryRow(query, a.ID, a.Name, a.Slug, a.Warehouse_id).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For delete
func (warc *Warehouses_CellsRepository) DeleteCompanies_WarehousesById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d ", tablewarehouses_cells, id)
	if _, err := warc.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}

//GET ALL

func (warc *Warehouses_CellsRepository) GetAll() ([]*models.Warehouses_cells, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablewarehouses_cells)
	rows, err := warc.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	warehouses_cells := make([]*models.Warehouses_cells, 0)
	for rows.Next() {
		a := models.Warehouses_cells{}

		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.Warehouse_id)
		if err != nil {
			log.Println(err)
			continue
		}

		warehouses_cells = append(warehouses_cells, &a)
	}
	return warehouses_cells, nil
}

//GET BY ID
func (warc *Warehouses_CellsRepository) GetByID(id int) ([]*models.Warehouses_cells, error) {
	query := fmt.Sprintf("SELECT * FROM %s Where warehouse_id = $id", tablewarehouses_cells)
	rows, err := warc.store.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	warehouses_cells := make([]*models.Warehouses_cells, 0)
	for rows.Next() {
		a := models.Warehouses_cells{}

		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.Warehouse_id)
		if err != nil {
			log.Println(err)
			continue
		}

		warehouses_cells = append(warehouses_cells, &a)
	}

	return warehouses_cells, nil
}

//DleteWarehousecellsById

func (warc *Warehouses_CellsRepository) DeleteWarehouses_cellsById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 ", tablewarehouses_cells)
	if _, err := warc.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
