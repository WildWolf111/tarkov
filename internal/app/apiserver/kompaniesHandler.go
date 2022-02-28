package apiserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vlasove/8.HandlerImpl2/internal/app/models"
)

func (api *APIServer) GetAllKompany(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	company, err := api.store.Kompany().SelectAll()
	if err != nil {
		api.logger.Info(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing companies in database. Try later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	api.logger.Info("Get All Companies GET /companies")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(company)
}

func (api *APIServer) Postkompany(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Kompany POST /kompanies")
	var company models.Companies
	err := json.NewDecoder(req.Body).Decode(&company)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	fmt.Println(company)
	a, err := api.store.Kompany().Create(&company)
	if err != nil {
		api.logger.Info("Troubles while connections to the company database:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)

}

func (api *APIServer) GetKompanyById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Company by ID /api/v1/companies/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, ok, err := api.store.Company().FindCompanyById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (companies) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find article with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Article with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}

func (api *APIServer) DeleteCompanyById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Company by Id DELETE /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Troubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unapropriate id value. don't use ID as uncasting to int value.",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, ok, err := api.store.Company().FindCompanyById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(500)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("Can not find company with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Company with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, err = api.store.Company().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles while deleting database elemnt from table (articles) with id. err:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    fmt.Sprintf("Company with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

//warehouse update
func (api *APIServer) UpdateCompanyById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating Company ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := Message{
			StatusCode: 400,
			Message:    "do not use parameter ID as uncasted to int type",
			IsError:    true,
		}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newCompany models.Warehouses

	err = json.NewDecoder(request.Body).Decode(&newCompany)
	if err != nil {
		msg := Message{
			StatusCode: 400,
			Message:    "provideed json file is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	newCompany.ID = id
	a, err := api.store.Warehouse().UpdateWarehouseById(&newCompany)
	if err != nil {
		api.logger.Info("Troubles while connections to the company database:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}
