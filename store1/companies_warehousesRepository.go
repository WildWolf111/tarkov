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

func (cwo *Companies_WarehousesRepository) SelectAllCompanies_Warehouses() ([]*models.Companies_Warehouses_Qwery, error) {

	query := fmt.Sprintf("SELECT %s.id, %s.slug, %s.id, %s.slug FROM %s RIGHT JOIN  %s ON  %s.id = %s.company_id LEFT JOIN %s ON %s.id = %s.warehouses_id",
		tablecompanies, tablecompanies, tablewarehouses, tablewarehouses, tablecompanies_warehouses,
		tablecompanies, tablecompanies, tablecompanies_warehouses, tablewarehouses, tablewarehouses,
		tablecompanies_warehouses)

	log.Println(query)
	rows, err := cwo.store.db.Query(query)
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Companies_Warehouses_Qwery := make([]*models.Companies_Warehouses_Qwery, 0)
	log.Println(Companies_Warehouses_Qwery)
	for rows.Next() {
		var (
			comp models.Company
			warh models.Warehouses
		)
		a := models.Companies_Warehouses_Qwery{
			Companies:  &comp,
			Warehouses: &warh,
		}
		log.Println(rows)
		err := rows.Scan(&a.Companies.ID, &a.Companies.Slug, &a.Warehouses.ID, &a.Warehouses.Slug)
		if err != nil {
			log.Println(err)
			continue
		}
		Companies_Warehouses_Qwery = append(Companies_Warehouses_Qwery, &a)
	}
	return Companies_Warehouses_Qwery, nil
}

//gresgtsrhgsrgh
