package store

import (
	"fmt"
	"log"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
)

type KompanyRepository struct {
	store *Store
}

var (
	tablekompany string = "kompany"
)

//For Post request
func (ko *KompanyRepository) Create(a *models.Kompany) (*models.Kompany, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, slug, inn, kpp) VALUES ($1, $2, $3,$4) RETURNING id", tablekompany)
	if err := ko.store.db.QueryRow(query, a.Name, a.Slug, a.INN, a.KPP).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For Update request
func (ko *KompanyRepository) UpdateKompanyById(a *models.Kompany) (*models.Kompany, error) {
	query := fmt.Sprintf("UPDATE %s SET (name, slug, inn, kpp) VALUES ($2, $3,$4,$5)WHERE id=$1 RETURNING id", tablekompany)
	if err := ko.store.db.QueryRow(query, a.ID, a.Name, a.Slug, a.INN, a.KPP).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For DELETE request
func (ko *KompanyRepository) DeleteById(id int) (*models.Kompany, error) {
	kompanies, ok, err := ko.FindKompanyById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("delete from %s where id=$1", tablekompany)
		_, err = ko.store.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}

	return kompanies, nil
}

//Helper for Delete by id and GET by id request
func (ko *KompanyRepository) FindKompanyById(id int) (*models.Kompany, bool, error) {
	kompanies, err := ko.SelectAll()
	found := false
	if err != nil {
		return nil, found, err
	}
	var kompanyFound *models.Kompany
	for _, a := range kompanies {
		if a.ID == id {
			kompanyFound = a
			found = true
		}
	}

	return kompanyFound, found, nil

}

//Get all request
func (ko *KompanyRepository) SelectAll() ([]*models.Kompany, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablekompany)
	rows, err := ko.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	kompanies := make([]*models.Kompany, 0)
	for rows.Next() {
		a := models.Kompany{}

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
		kompanies = append(kompanies, &a)
	}
	return kompanies, nil
}

//Get  request dByID
func (ko *KompanyRepository) GetKompanyById(id int) ([]*models.Kompany, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tablekompany)
	rows, err := ko.store.db.Query(query, id)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	kompanies := make([]*models.Kompany, 0)
	for rows.Next() {
		a := models.Kompany{}

		err := rows.Scan(&a.ID, &a.Name, &a.Slug, &a.INN, &a.KPP)
		if err != nil {
			log.Println(err)
			continue
		}
		var ok bool
		a.Warehouses, ok, err = ko.store.Warehouse().GetWarehouseByCompanyId(a.ID)
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
