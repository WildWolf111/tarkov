package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Instance of store
type Store struct {
	config                         *Config
	db                             *sql.DB
	warehouseRepository            *WarehouseRepository
	stocRepository                 *StocRepository
	CompanyRepository              *CompanyRepository
	Companies_WarehousesRepository *Companies_WarehousesRepository
}

// Constructor for store
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//Open store method
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	//Проверим, что все ок. Реально соединение тут не создается. Соединение только при первом вызове
	//db.Ping() // Пустой SELECT *
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("Connection to db successfully")
	return nil
}

//Close store method
func (s *Store) Close() {
	s.db.Close()
}

//Public for WarehouseRepo
func (s *Store) Warehouse() *WarehouseRepository {
	if s.CompanyRepository != nil {
		return s.warehouseRepository
	}
	s.warehouseRepository = &WarehouseRepository{
		store: s,
	}
	return s.warehouseRepository
}

//Stoc for StocsRepo
func (s *Store) Stoc() *StocRepository {
	if s.stocRepository != nil {
		return s.stocRepository
	}
	s.stocRepository = &StocRepository{
		store: s,
	}
	return s.stocRepository
}

//Public for CompanyRepo
func (s *Store) Company() *CompanyRepository {
	if s.CompanyRepository != nil {
		return s.CompanyRepository
	}
	s.CompanyRepository = &CompanyRepository{
		store: s,
	}
	return s.CompanyRepository
}

//Public for Companies_warehouses Repo
func (s *Store) Companies_Warehouses() *Companies_WarehousesRepository {
	if s.Companies_WarehousesRepository != nil {
		return s.Companies_WarehousesRepository
	}
	s.Companies_WarehousesRepository = &Companies_WarehousesRepository{
		store: s,
	}
	return s.Companies_WarehousesRepository
}
