package store

import (
	"fmt"
	"log"

	"github.com/vlasove/8.HandlerImpl2/internal/app/models"
)

type WarehouseRepository struct {
	store *Store
}

var (
	tablewarehouse string = "warehouses"
)

//For Post request
func (wa *WarehouseRepository) Create(a *models.Warehouses) (*models.Warehouses, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, slug, company_id, address, id) VALUES ($1, $2, $3,$4,$5) RETURNING id", tablewarehouse)
	if err := wa.store.db.QueryRow(query, a.Name, a.Slug, a.Company_id, a.Address, a.ID).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For Update request
func (wa *WarehouseRepository) UpdateWarehouseById(a *models.Warehouses) (*models.Warehouses, error) {
	query := fmt.Sprintf("UPDATE %s SET name=$1, slug=$2, company_id=$3, address=$4) WHERE id=$5 RETURNING id", tablewarehouse)
	if err := wa.store.db.QueryRow(query, a.Name, a.Slug, a.Company_id, a.Address, a.ID).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For DELETE request
func (wa *WarehouseRepository) DeleteById(id int) (*models.Warehouses, error) {
	warehouses, ok, err := wa.FindWarehouseById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("delete from %s where id=$1", tablewarehouse)
		_, err = wa.store.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}

	return warehouses, nil
}

//Helper for Delete by id and GET by id request
func (wa *WarehouseRepository) FindWarehouseById(id int) (*models.Warehouses, bool, error) {
	warehouses, err := wa.SelectAll()
	founded := false
	if err != nil {
		return nil, founded, err
	}
	var warehouseFinded *models.Warehouses
	for _, a := range warehouses {
		if a.ID == id {
			warehouseFinded = a
			founded = true
		}
	}

	return warehouseFinded, founded, nil

}

//Get all request and helper for FindByID
func (wa *WarehouseRepository) SelectAll() ([]*models.Warehouses, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablewarehouse)
	rows, err := wa.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	warehouses := make([]*models.Warehouses, 0)
	for rows.Next() {
		a := models.Warehouses{}

		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.Company_id, &a.Address)
		if err != nil {
			log.Println(err)
			continue
		}

		w, ok, err := wa.store.Company().GetCompanyById(a.Company_id)
		if err != nil {
			log.Println(err)
			continue
		}
		if !ok {
			log.Printf("Company with id %d not found", a.Company_id)
		}
		a.Companies = w
		warehouses = append(warehouses, &a)
	}
	return warehouses, nil
}

//Get all request and helper for FindByID
func (wa *WarehouseRepository) GetWarehouseByCompanyId(id int) ([]*models.Warehouses, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE company_id = $1", tablewarehouse)
	log.Println(query)
	rows, err := wa.store.db.Query(query, id)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	warehouses := make([]*models.Warehouses, 0)
	for rows.Next() {
		a := models.Warehouses{}
		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.Company_id, &a.Address)
		if err != nil {
			log.Println(err)
			continue
		}
		warehouses = append(warehouses, &a)
	}
	if len(warehouses) == 0 {
		return warehouses, false, nil
	}
	return warehouses, true, nil
}
