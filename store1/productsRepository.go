package store

import (
	"fmt"
	"log"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
)

type ProductsRepository struct {
	store *Store
}

var (
	tableproducts string = "products"
)

//For Post request
func (cou *CountriesRepository) PostProduct(a *models.Product) (*models.Product, error) {
	query := fmt.Sprintf("INSERT INTO %s (name,slug, sku,shortDesc,fullDesc) VALUES ( $1, $2, $3, $4, $5) RETURNING id", tableproducts)
	log.Println(query)
	if err := cou.store.db.QueryRow(query, a.Name, a.Slug, a.Sku, a.ShortDesc, a.FullDesc).Scan(&a.Id); err != nil {
		return nil, err
	}
	return a, nil
}

//For delete
func (cou *CountriesRepository) DeleteProductsById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = %d ", tableproducts, id)
	if _, err := cou.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}

//GET ALL

func (cou *CountriesRepository) GetAllProducts() ([]*models.Product, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableproducts)
	rows, err := cou.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Countries := make([]*models.Product, 0)
	for rows.Next() {
		a := models.Product{}

		err := rows.Scan(&a.Id, &a.Name, &a.Slug, &a.Sku, &a.ShortDesc, &a.FullDesc)
		if err != nil {
			log.Println(err)
			continue
		}

		Countries = append(Countries, &a)
	}
	return Countries, nil
}

//GET BY ID
func (cou *CountriesRepository) GetProductByID(id int) ([]*models.Product, bool, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableproducts)
	log.Println(query)
	rows, err := cou.store.db.Query(query, id)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()
	Countries := make([]*models.Product, 0)
	for rows.Next() {
		a := models.Product{}
		err := rows.Scan(&a.Id, &a.Name, &a.Slug, &a.Sku, &a.ShortDesc, &a.FullDesc)
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

//DleteProductcsById

func (cou *CountriesRepository) DeleteProductById(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 ", tableproducts)
	if _, err := cou.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
