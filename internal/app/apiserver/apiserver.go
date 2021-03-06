package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	store "github.com/vlasove/8.HandlerImpl2/store1"
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
	s.router.HandleFunc(prefix+"/articles", s.GetAllArticles).Methods("GET")
	s.router.HandleFunc(prefix+"/articles"+"/{id}", s.GetArticleById).Methods("GET")
	s.router.HandleFunc(prefix+"/articles"+"/{id}", s.DeleteArticleById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/articles", s.PostArticle).Methods("POST")

	s.router.HandleFunc(prefix+"/user/register", s.PostUserRegister).Methods("POST")

	//router companies
	s.router.HandleFunc(prefix+"/companies", s.GetAllCompany).Methods("GET")
	s.router.HandleFunc(prefix+"/companies"+"/{id}", s.GetCompanyById).Methods("GET")
	s.router.HandleFunc(prefix+"/companies"+"/{id}", s.DeleteCompanyById).Methods("DELETE")
	s.router.HandleFunc(prefix+"/companies", s.PostCompany).Methods("POST")
	s.router.HandleFunc(prefix+"/{id}", s.UpdateCompanyById).Methods("PUT")

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
