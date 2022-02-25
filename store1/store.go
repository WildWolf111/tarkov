package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Instance of store
type Store struct {
	config              *Config
	db                  *sql.DB
	userRepository      *UserRepository
	articleRepository   *ArticleRepository
	companyRepository   *CompanyRepository
	warehouseRepository *WarehouseRepository
	stocRepository      *StocRepository
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

//Public for UserRepo
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

//Public for ArticleRepo
func (s *Store) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		store: s,
	}
	return s.articleRepository
}

//Public for CompanyRepo
func (s *Store) Company() *CompanyRepository {
	if s.companyRepository != nil {
		return s.companyRepository
	}
	s.companyRepository = &CompanyRepository{
		store: s,
	}
	return s.companyRepository
}

//Public for WarehouseRepo
func (s *Store) Warehouse() *WarehouseRepository {
	if s.companyRepository != nil {
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
