package store

import (
	"log"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"

	"fmt"
)

type Companies_WarehousesRepository struct {
	store *Store
}

var (
	tablewarehouses           string = "warehouses"
	tablecompanies_warehouses string = "companies_warehouses"
)

func (cwo *Companies_WarehousesRepository) Create(a *models.Companies_Warehouses) error {
	query := fmt.Sprintf("INSERT INTO %s (company_id, warehouses_id) VALUES ($1, $2)", tablecompanies_warehouses)
	if _, err := cwo.store.db.Exec(query, a.Companies_id, a.Warehouses_id); err != nil {
		return err
	}
	return nil
}

func (cwo *Companies_WarehousesRepository) SelectAllCompanies_Warehouses(id int) ([]*models.Warehouses, error) {
	query := fmt.Sprintf("SELECT %s.* FROM %s JOIN  %s ON  %s.id = %s.warehouses_id WHERE %s.company_id = id ", tablewarehouses, tablecompanies_warehouses, tablewarehouses, tablewarehouses, tablecompanies_warehouses, tablecompanies_warehouses)
	log.Println(query)
	rows, err := cwo.store.db.Query(query)
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	warehouses := make([]*models.Warehouses, 0)
	log.Println(warehouses)
	for rows.Next() {
		a := models.Warehouses{}
		log.Println(rows)
		err := rows.Scan(&a)
		if err != nil {
			log.Println(err)
			continue
		}
		warehouses = append(warehouses, &a)
	}
	return warehouses, nil
}
