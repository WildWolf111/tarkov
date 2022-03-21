package store

import (
	"fmt"
	"log"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
)

type CountriesRepository struct {
	store *Store
}

var (
	tablecountries string = "countries"
)

//For Post request
func (cou *CountriesRepository) PostCountries(a *models.Countries) (*models.Countries, error) {
	query := fmt.Sprintf("INSERT INTO %s (code, country) VALUES ( $1, $2) RETURNING id", tablecountries)
	log.Println(query)
	if err := cou.store.db.QueryRow(query, a.Code, a.Country).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For delete
func (cou *CountriesRepository) DeleteCountryById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d ", tablecountries, id)
	if _, err := cou.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}

//GET ALL

func (cou *CountriesRepository) GetAllCountries() ([]*models.Countries, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablecountries)
	rows, err := cou.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Countries := make([]*models.Countries, 0)
	for rows.Next() {
		a := models.Countries{}

		err := rows.Scan(&a.ID, &a.Code, &a.Country)
		if err != nil {
			log.Println(err)
			continue
		}

		Countries = append(Countries, &a)
	}
	return Countries, nil
}

//GET BY ID
func (cou *CountriesRepository) GetCountryByID(id int) ([]*models.Countries, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tablecountries)
	log.Println(query)
	rows, err := cou.store.db.Query(query, id)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	Countries := make([]*models.Countries, 0)
	for rows.Next() {
		a := models.Countries{}
		err := rows.Scan(&a.ID, &a.Code, &a.Country)
		if err != nil {
			log.Println(err)
			continue
		}
		Countries = append(Countries, &a)
	}
	if len(Countries) == 0 {
		return Countries, false, nil
	}
	return Countries, true, nil
}

//DletecountriescellsById

func (cou *CountriesRepository) DeleteCountriesById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 ", tablecountries)
	if _, err := cou.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
