package store

import (
	"fmt"
	"log"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
)

type CompanyRepository struct {
	store *Store
}

var (
	tablecompanies string = "companies"
)

//For Post request
func (co *CompanyRepository) Create(a *models.Company) (*models.Company, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, slug, inn, kpp) VALUES ($1, $2, $3,$4) RETURNING id", tablecompanies)
	if err := co.store.db.QueryRow(query, a.Name, a.Slug, a.INN, a.KPP).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For Update request
func (co *CompanyRepository) UpdateCompanyById(a *models.Company) (*models.Company, error) {
	query := fmt.Sprintf("UPDATE %s SET (name, slug, inn, kpp) VALUES ($2, $3,$4,$5) WHERE id=$1 RETURNING id", tablecompanies)
	if err := co.store.db.QueryRow(query, a.ID, a.Name, a.Slug, a.INN, a.KPP).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For DELETE request
func (co *CompanyRepository) DeleteById(id int) (*models.Company, error) {
	Company, ok, err := co.FindCompanyById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("delete from %s where id=$1", tablecompanies)
		_, err = co.store.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}

	return Company, nil
}

//Helper for Delete by id and GET by id request
func (co *CompanyRepository) FindCompanyById(id int) (*models.Company, bool, error) {
	Company, err := co.SelectAll()
	found := false
	if err != nil {
		return nil, found, err
	}
	var companyFound *models.Company
	for _, a := range Company {
		if a.ID == id {
			companyFound = a
			found = true
		}
	}

	return companyFound, found, nil

}

//Get all request
func (co *CompanyRepository) SelectAll() ([]*models.Company, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablecompanies)
	rows, err := co.store.db.Query(query)
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
		/*
			w, ok, err := ko.store.Warehouse().GetWarehouseByCompanyId(a.ID)
			if err != nil {
				log.Println(err)
				continue
			}
			if !ok {
				log.Printf("Kompany with id %d not found", a.ID)
			}
			a.Warehouses = w
		*/
		Company = append(Company, &a)
	}
	return Company, nil
}

//Get  request dByID
func (co *CompanyRepository) GetCompanyById(id int) ([]*models.Company, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tablecompanies)
	rows, err := co.store.db.Query(query, id)
	if err != nil {
		return nil, false, err
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
		/*var ok bool
		a.Warehouses, ok, err = co.store.Warehouse().GetWarehouseByCompanyId(a.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		if !ok {
			log.Printf("Company warehouse with id %d not found", a.ID)
			continue
		}
		*/
		Company = append(Company, &a)
	}
	if len(Company) == 0 {
		return Company, false, nil
	}
	return Company, true, nil
}
