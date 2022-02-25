package store

import (
	"fmt"
	"log"

	"github.com/vlasove/8.HandlerImpl2/internal/app/models"
)

type StocRepository struct {
	store *Store
}

var (
	tablestocs string = "stocs"
)

//For Post request
func (st *StocRepository) Create(a *models.Stocs) (*models.Stocs, error) {
	query := fmt.Sprintf("INSERT INTO %s (id,company_sender_id, company_recipient_id, product_id, quantity, warehouse_cell_id, gtd_id) VALUES ($1, $2, $3,$4,$5,$6,$7) RETURNING id", tablestocs)
	log.Println(query)
	if err := st.store.db.QueryRow(query, a.ID, a.Company_sender_id, a.Company_recipient_id, a.Product_id, a.Quantity, a.Warehouse_cell_id, a.GTD_id).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

// For Put request

func (st *StocRepository) UpdateStocById(a *models.Stocs) (*models.Stocs, error) {
	query := fmt.Sprintf("UPDATE %s SET (id, company_sender_id, company_recipient_id, product_id, quantity, warehouse_cell_id, gtd_id) VALUES ($1, $2, $3,$4,$5,$6,&7) WHERE id=$1 RETURNING id", tablestocs)
	if err := st.store.db.QueryRow(query, a.ID, a.Company_sender_id, a.Company_recipient_id, a.Product_id, a.Quantity, a.Warehouse_cell_id, a.GTD_id).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil

}

//For DELETE request
func (st *StocRepository) DeleteById(id int) (*models.Stocs, error) {
	stocs, ok, err := st.FindStocById(id)
	if err != nil {
		return nil, err
	}
	if ok {
		query := fmt.Sprintf("delete from %s where id=$1", tablestocs)
		_, err = st.store.db.Exec(query, id)
		if err != nil {
			return nil, err
		}
	}

	return stocs, nil
}

//Helper for Delete by id and GET by id request
func (st *StocRepository) FindStocById(id int) (*models.Stocs, bool, error) {
	stocs, err := st.SelectAll()
	founded := false
	if err != nil {
		return nil, founded, err
	}
	var stocFinded *models.Stocs
	for _, a := range stocs {
		if a.ID == id {
			stocFinded = a
			founded = true
		}
	}

	return stocFinded, founded, nil

}

//Get all request and helper for FindByID
func (st *StocRepository) SelectAll() ([]*models.Stocs, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablestocs)
	log.Println(query)
	rows, err := st.store.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	stocs := make([]*models.Stocs, 0)
	for rows.Next() {
		a := models.Stocs{}
		log.Println(rows)
		err := rows.Scan(&a.ID, &a.Company_sender_id, &a.Company_recipient_id, &a.Product_id, &a.Quantity, &a.Warehouse_cell_id, &a.GTD_id)
		if err != nil {
			log.Println(err)
			continue
		}
		stocs = append(stocs, &a)
	}
	return stocs, nil
}
