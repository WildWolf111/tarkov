package store

import (
	"fmt"
	"log"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
)

type GtdRepository struct {
	store *Store
}

var (
	tablegtd string = "gtd"
)

//For Post request
func (gtd *GtdRepository) PostGtds(a *models.Gtd) (*models.Gtd, error) {
	query := fmt.Sprintf("INSERT INTO %s (  coutry,number) VALUES ( $1, $2,) RETURNING id", tablegtd)
	log.Println(query)
	if err := gtd.store.db.QueryRow(query, a.Country, a.Number).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

//For delete
func (gtd *GtdRepository) DeleteGtdsById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d ", tablegtd, id)
	if _, err := gtd.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}

//GET ALL

func (gtd *GtdRepository) GetAllGtds() ([]*models.Gtd, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablegtd)
	rows, err := gtd.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Gtd := make([]*models.Gtd, 0)
	for rows.Next() {
		a := models.Gtd{}

		err := rows.Scan(&a.ID, &a.Country, &a.Number)
		if err != nil {
			log.Println(err)
			continue
		}

		Gtd = append(Gtd, &a)
	}
	return Gtd, nil
}

//GET BY ID
func (gtd *GtdRepository) GetGtdByID(id int) ([]*models.Gtd, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tablegtd)
	log.Println(query)
	rows, err := gtd.store.db.Query(query, id)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	Gtd := make([]*models.Gtd, 0)
	for rows.Next() {
		a := models.Gtd{}
		err := rows.Scan(&a.ID, &a.Country, &a.Number)
		if err != nil {
			log.Println(err)
			continue
		}
		Gtd = append(Gtd, &a)
	}
	if len(Gtd) == 0 {
		return Gtd, false, nil
	}
	return Gtd, true, nil
}

//DleteWarehousecellsById

func (gtd *GtdRepository) DeleteGtdById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 ", tablegtd)
	if _, err := gtd.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
