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
	StocksRepository               *StocksRepository
	CompanyRepository              *CompanyRepository
	Companies_WarehousesRepository *Companies_WarehousesRepository
	Warehouses_CellsRepository     *Warehouses_CellsRepository
	GtdRepository                  *GtdRepository
	CountriesRepository            *CountriesRepository
	ProductsRepository             *ProductsRepository
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
func (s *Store) Stoc() *StocksRepository {
	if s.StocksRepository != nil {
		return s.StocksRepository
	}
	s.StocksRepository = &StocksRepository{
		store: s,
	}
	return s.StocksRepository
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

//Public for Warehouses_cells Repo
func (s *Store) Warehouses_cells() *Warehouses_CellsRepository {
	if s.Warehouses_CellsRepository != nil {
		return s.Warehouses_CellsRepository
	}
	s.Warehouses_CellsRepository = &Warehouses_CellsRepository{
		store: s,
	}
	return s.Warehouses_CellsRepository
}

//Public for GTD Repo
func (s *Store) Gtd() *GtdRepository {
	if s.GtdRepository != nil {
		return s.GtdRepository
	}
	s.GtdRepository = &GtdRepository{
		store: s,
	}
	return s.GtdRepository
}

//Coutries Repo
func (s *Store) Countries() *CountriesRepository {
	if s.CountriesRepository != nil {
		return s.CountriesRepository
	}
	s.CountriesRepository = &CountriesRepository{
		store: s,
	}
	return s.CountriesRepository
}

//Products Repo
func (s *Store) Product() *ProductsRepository {
	if s.ProductsRepository != nil {
		return s.ProductsRepository
	}
	s.ProductsRepository = &ProductsRepository{
		store: s,
	}
	return s.ProductsRepository
}
