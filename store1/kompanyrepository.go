package store

import (
	"fmt"
	"log"
)

type KompanyRepository struct {
	store *Store
}

var (
	tableKompanies string = "kompanies"
)

//For Post request
func (co *CompanyRepository) Create(a *models.kompanies) (*models.kompanies, error) {
	query := fmt.Sprintf("INSERT INTO %s (id, name, slug, inn, kpp) VALUES ($1, $2, $3,$4,$5) RETURNING id", tablecompanies)
	if err := co.store.db.QueryRow(query, a.ID, a.Name, a.Slug, a.INN, a.KPP).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

// For Put request

func (co *CompanyRepository) Update(a *models.kompanies) (*models.kompanies, error) {
	query := fmt.Sprintf("UPDATE %s SET (id, name, slug, inn, kpp) VALUES ($1, $2, $3,$4,$5)WHERE id=$1 RETURNING id", tablecompanies)
	if err := co.store.db.QueryRow(query, a.ID, a.Name, a.Slug, a.INN, a.KPP).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil

}

//For DELETE request
func (co *CompanyRepository) DeleteById(id int) (*models.kompanies, error) {
	kompanies, ok, err := co.FindCompanyById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("delete from %s where id=$1", tablekompanies)
		_, err = co.store.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}

	return kompanies, nil
}

//Helper for Delete by id and GET by id request
func (co *CompanyRepository) FindCompanyById(id int) (*models.kompanies, bool, error) {
	kompanies, err := co.SelectAll()
	founded := false
	if err != nil {
		return nil, founded, err
	}
	var companyFinded *models.kompanies
	for _, a := range kompanies {
		if a.ID == id {
			companyFinded = a
			founded = true
		}
	}

	return companyFinded, founded, nil

}

//Get all request and helper for FindByID
func (co *CompanyRepository) SelectAll() ([]*models.kompanies, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablecompanies)
	rows, err := co.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	kompanies := make([]*models.kompanies, 0)
	for rows.Next() {
		a := models.kompanies{}
		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.INN, &a.KPP)
		if err != nil {
			log.Println(err)
			continue
		}

		w, ok, err := co.store.Warehouse().GetWarehouseByCompanyId(a.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		if !ok {
			log.Printf("Company with id %d not found", a.ID)
		}
		a.Warehouses = w

		kompanies = append(kompanies, &a)
	}
	return kompanies, nil
}

//Get all request and helper for FindByID
func (co *CompanyRepository) GetKompanyById(id int) ([]*models.kompany, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tablecompanies)
	rows, err := co.store.db.Query(query, id)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	kompanies := make([]*models.kompanies, 0)
	for rows.Next() {
		a := models.kompanies{}

		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.INN, &a.KPP)
		if err != nil {
			log.Println(err)
			continue
		}
		var ok bool
		a.Warehouses, ok, err = co.store.Warehouse().GetWarehouseByCompanyId(a.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		if !ok {
			log.Printf("Company warehouse with id %d not found", a.ID)
			continue
		}

		kompanies = append(kompanies, &a)
	}
	if len(kompanies) == 0 {
		return kompanies, false, nil
	}
	return kompanies, true, nil
}
