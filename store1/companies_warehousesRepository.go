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
		tablecompanies, tablecompanies, tablewarehouses, tablewarehouses, tablecompanies,
		tablecompanies_warehouses, tablecompanies, tablecompanies_warehouses, tablewarehouses, tablewarehouses,
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

//SelectWarehousesByCompanyId
func (cwo *Companies_WarehousesRepository) SelectWarehousesByCompanyId(id int) ([]*models.Warehouses, error) {
	query := fmt.Sprintf("SELECT %s.* FROM %s JOIN %s ON %s.id = %s.warehouses_id WHERE %s.company_id = %d",
		tablewarehouses, tablecompanies_warehouses, tablewarehouse, tablewarehouse, tablecompanies_warehouses, tablecompanies_warehouses, id)

	log.Println(query)
	rows, err := cwo.store.db.Query(query)
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Warehouses := make([]*models.Warehouses, 0)
	log.Println(Warehouses)
	for rows.Next() {
		a := models.Warehouses{}

		log.Println(rows)
		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.Company_id, &a.Address)
		if err != nil {
			log.Println(err)
			continue
		}
		Warehouses = append(Warehouses, &a)
	}
	return Warehouses, nil
}

//SelectCompaniesByWarehouseId
func (cwo *Companies_WarehousesRepository) SelectCompaniesByWarehouseId(id int) ([]*models.Company, error) {
	query := fmt.Sprintf("SELECT %s.* FROM %s JOIN %s ON %s.id = %s.warehouses_id WHERE %s.warehouses_id = %d",
		tablecompanies, tablecompanies_warehouses, tablecompanies, tablecompanies, tablecompanies_warehouses, tablecompanies_warehouses, id)

	log.Println(query)
	rows, err := cwo.store.db.Query(query)
	log.Println(rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Company := make([]*models.Company, 0)
	for rows.Next() {
		a := models.Company{}

		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.INN, &a.KPP)
		if err != nil {
			log.Println(err)
			continue
		}
		Company = append(Company, &a)
	}
	return Company, nil
}

//Delete From companies_warehouses

func (cwo *Companies_WarehousesRepository) DeleteCompanies_WarehousesById(compwar *models.Companies_Warehouses) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE company_id = %d and warehouses_id = %d", tablecompanies_warehouses, compwar.Companies_id, compwar.Warehouses_id)
	if _, err := cwo.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
