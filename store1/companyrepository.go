package store

import (
	"fmt"
	"log"

	"github.com/vlasove/8.HandlerImpl2/internal/app/models"
)

type CompanyRepository struct {
	store *Store
}

var (
	tablecompanies string = "companies"
)

//For Post request
func (co *CompanyRepository) Create(a *models.Companies) (*models.Companies, error) {
	query := fmt.Sprintf("INSERT INTO %s (id, name, slug, inn, kpp) VALUES ($1, $2, $3,$4,$5) RETURNING id", tablecompanies)
	if err := co.store.db.QueryRow(query, a.ID, a.Name, a.Slug, a.INN, a.KPP).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For DELETE request
func (co *CompanyRepository) DeleteById(id int) (*models.Companies, error) {
	companies, ok, err := co.FindCompanyById(id)
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

	return companies, nil
}

//Helper for Delete by id and GET by id request
func (co *CompanyRepository) FindCompanyById(id int) (*models.Companies, bool, error) {
	companies, err := co.SelectAll()
	founded := false
	if err != nil {
		return nil, founded, err
	}
	var companyFinded *models.Companies
	for _, a := range companies {
		if a.ID == id {
			companyFinded = a
			founded = true
		}
	}

	return companyFinded, founded, nil

}

//Get all request and helper for FindByID
func (co *CompanyRepository) SelectAll() ([]*models.Companies, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablecompanies)
	rows, err := co.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	companies := make([]*models.Companies, 0)
	for rows.Next() {
		a := models.Companies{}
		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.INN, &a.KPP)
		if err != nil {
			log.Println(err)
			continue
		}
		companies = append(companies, &a)
	}
	return companies, nil
}