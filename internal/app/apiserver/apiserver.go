package apiserver

import (
	"net/http"

	store "github.com/WildWolf111/StandarWebSrver2/store1"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
)

// type for APIServer object for instancing server
type APIServer struct {
	//Unexported field
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

//APIServer constructor
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start http server and connection to db and logger confs
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting api server at port :", s.config.BindAddr)
	s.configureRouter()
	if err := s.configureStore(); err != nil {
		return err
	}
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

//func for configureate logger, should be unexported
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return nil
	}
	s.logger.SetLevel(level)

	return nil
}

//func for configure Router
func (s *APIServer) configureRouter() {

	//router warehouses
	s.router.HandleFunc(prefix+"/warehouses", s.GetAllWarehouse).Methods("GET")
	s.router.HandleFunc(prefix+"/warehouses"+"/{id}", s.GetWarehouseById).Methods("GET")
	s.router.HandleFunc(prefix+"/warehouses"+"/{id}", s.DeleteWarehouseById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/warehouses", s.PostWarhouse).Methods("POST")
	s.router.HandleFunc(prefix+"/{id}", s.UpdateWarehouseById).Methods("PUT")

	//router stocs

	s.router.HandleFunc(prefix+"/stocs", s.GetAllStocs).Methods("GET")
	s.router.HandleFunc(prefix+"/stocs"+"/{id}", s.GetStocById).Methods("GET")
	s.router.HandleFunc(prefix+"/stocs"+"/{id}", s.DeleteStocById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/stocs", s.PostStocs).Methods("POST")
	s.router.HandleFunc(prefix+"/{id}", s.UpdateStocById).Methods("PUT")

	//router kompanies
	s.router.HandleFunc(prefix+"/companies", s.GetAllCompany).Methods("GET")
	s.router.HandleFunc(prefix+"/companies"+"/{id}", s.GetCompanyById).Methods("GET")
	s.router.HandleFunc(prefix+"/companies"+"/{id}", s.DeleteCompanyById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/companies", s.PostCompany).Methods("POST")
	s.router.HandleFunc(prefix+"/{id}", s.UpdateCompanyById).Methods("PUT")

	//router companies_warhouses
	s.router.HandleFunc(prefix+"/companies_warehouses", s.PostCompanies_Warehouses).Methods("POST")
	s.router.HandleFunc(prefix+"/companies_warehouses", s.GetAllCompanies_Warehouses).Methods("GET")
	s.router.HandleFunc(prefix+"/companies_warehouses"+"/companies"+"/{id}", s.GetWarehousesByCompanyId).Methods("GET")
	s.router.HandleFunc(prefix+"/companies_warehouses"+"/warehouses"+"/{id}", s.GetCompanyByWarehousesId).Methods("GET")
	s.router.HandleFunc(prefix+"/companies_warehouses/delete", s.DeleteCompaniesWarehousesById).Methods("DELETE")

	//router Warehouses_cells

	s.router.HandleFunc(prefix+"/Warehouses_cells", s.Post).Methods("POST")
	s.router.HandleFunc(prefix+"/Warehouses_cells"+"/delete"+"/{id}", s.DeleteWarehouses_cellsById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/Warehouses_cells"+"/update"+"/{id}", s.UpdateWarehouses_cellsById).Methods("PUT")
	s.router.HandleFunc(prefix+"/Warehouses_cells"+"/get", s.GetAllWarehouses_cells).Methods("GET")
	s.router.HandleFunc(prefix+"/Warehouses_cells"+"/{id}", s.GetWarehouses_cellsById).Methods("GET")

	//router GTD'S

	s.router.HandleFunc(prefix+"/gtds", s.PostGtds).Methods("POST")
	s.router.HandleFunc(prefix+"/gtds"+"/delete"+"/{id}", s.DeleteGtdById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/gtds"+"/get", s.GetAll_Gtds).Methods("GET")
	s.router.HandleFunc(prefix+"/gtds"+"/{id}", s.DeleteGtdById).Methods("GET")

	//Countries

	s.router.HandleFunc(prefix+"/countries", s.PostCountries).Methods("POST")
	s.router.HandleFunc(prefix+"/countries"+"/delete"+"/{id}", s.DeleteCountryById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/countries"+"/get", s.GetAll_Countries).Methods("GET")
	s.router.HandleFunc(prefix+"/countries"+"/{id}", s.GetCountryById).Methods("GET")

	//Product

	s.router.HandleFunc(prefix+"/products", s.PostProduct).Methods("POST")
	s.router.HandleFunc(prefix+"/products"+"/delete"+"/{id}", s.DeleteProductById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/products"+"/get", s.GetAllProducts).Methods("GET")
	s.router.HandleFunc(prefix+"/products"+"/{id}", s.GetProductById).Methods("GET")

}

//configureStore method
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}
