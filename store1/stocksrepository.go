package store

import (
	"fmt"
	"log"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
)

type StocksRepository struct {
	store *Store
}

var (
	tablestocs string = "stocs"
)

//For Post request
func (st *StocksRepository) Create(a *models.Stocks) (*models.Stocks, error) {
	query := fmt.Sprintf("INSERT INTO %s (id,company_sender_id, company_recipient_id, product_id, quantity, warehouse_cell_id, gtd_id) VALUES ($1, $2, $3,$4,$5,$6,$7) RETURNING id", tablestocs)
	log.Println(query)
	if err := st.store.db.QueryRow(query, a.ID, a.Company_sender_id, a.Company_recipient_id, a.Product_id, a.Quantity, a.Warehouse_cell_id, a.GTD_id).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil
}

// For Put request

func (st *StocksRepository) UpdateStocById(a *models.Stocks) (*models.Stocks, error) {
	query := fmt.Sprintf("UPDATE %s SET (id, company_sender_id, company_recipient_id, product_id, quantity, warehouse_cell_id, gtd_id) VALUES ($1, $2, $3,$4,$5,$6,&7) WHERE id=$1 RETURNING id", tablestocs)
	if err := st.store.db.QueryRow(query, a.ID, a.Company_sender_id, a.Company_recipient_id, a.Product_id, a.Quantity, a.Warehouse_cell_id, a.GTD_id).Scan(&a.ID); err != nil {
		return nil, err
	}
	return a, nil

}

//For DELETE request
func (st *StocksRepository) DeleteById(id int) (*models.Stocks, error) {
	Stocks, ok, err := st.FindStocById(id)
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

	return Stocks, nil
}

//Helper for Delete by id and GET by id request
func (st *StocksRepository) FindStocById(id int) (*models.Stocks, bool, error) {
	stocs, err := st.SelectAll()
	founded := false
	if err != nil {
		return nil, founded, err
	}
	var stocFinded *models.Stocks
	for _, a := range stocs {
		if a.ID == id {
			stocFinded = a
			founded = true
		}
	}

	return stocFinded, founded, nil

}

//Get all request and helper for FindByID
func (st *StocksRepository) SelectAll() ([]*models.Stocks, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tablestocs)
	log.Println(query)
	rows, err := st.store.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Stocks := make([]*models.Stocks, 0)
	for rows.Next() {
		a := models.Stocks{}
		log.Println(rows)
		err := rows.Scan(&a.ID, &a.Company_sender_id, &a.Company_recipient_id, &a.Product_id, &a.Quantity, &a.Warehouse_cell_id, &a.GTD_id)
		if err != nil {
			log.Println(err)
			continue
		}
		Stocks = append(Stocks, &a)
	}
	return Stocks, nil
}
