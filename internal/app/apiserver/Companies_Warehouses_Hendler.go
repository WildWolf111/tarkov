package apiserver

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"strconv"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
	"github.com/gorilla/mux"
)

//post
func (api *APIServer) PostCompanies_Warehouses(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Companies_Warehouses POST /Companies_Warehouses")

	var Companies_Warehouses models.Companies_Warehouses

	err := json.NewDecoder(req.Body).Decode(&Companies_Warehouses)
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
	fmt.Println(Companies_Warehouses)
	err = api.store.Companies_Warehouses().Create(&Companies_Warehouses)
	if err != nil {
		api.logger.Info("Troubles while connections to the warehouse database:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 200,
		Message:    "",
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}

//Get all requests

func (api *APIServer) GetAllCompanies_Warehouses(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	Companies_Warehouses_Qwery, err := api.store.Companies_Warehouses().SelectAllCompanies_Warehouses()
	if err != nil {
		api.logger.Info(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing companies_warehouses in database. Try later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	api.logger.Info("Get All Companies_warehouses GET /companies_warehouses")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(Companies_Warehouses_Qwery)
}

//GetWarehousesByCompanyId
func (api *APIServer) GetWarehousesByCompanyId(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Warehouses By Company Id /api/v1/companies_warehouses/companies/{id}")
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

	Companies_Warehouses_Qwery, err := api.store.Companies_Warehouses().SelectWarehousesByCompanyId(id)
	if err != nil {
		api.logger.Info(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing companies_warehouses in database. Try later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	api.logger.Info("GetWarehousesByCompanyId /companies_warehouses")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(Companies_Warehouses_Qwery)
}

//GetCmpanyByWarehouseId
func (api *APIServer) GetCompanyByWarehousesId(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Warehouses By Company Id /api/v1/companies_warehouses/warehouses/{id}")
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

	Companies_Warehouses, err := api.store.Companies_Warehouses().SelectCompaniesByWarehouseId(id)
	if err != nil {
		api.logger.Info(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing companies_warehouses in database. Try later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	api.logger.Info("GetCompanyByWarehouseId /companies_warehouses")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(Companies_Warehouses)
}

//Delete

func (api *APIServer) DeleteCompaniesWarehousesById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("DeleteCompaniesWarehousesById /api/v1/companies_warehouses/delete")

	var Companies_Warehouses models.Companies_Warehouses
	log.Println(Companies_Warehouses)

	err := json.NewDecoder(req.Body).Decode(&Companies_Warehouses)
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
	fmt.Println(Companies_Warehouses)
	err = api.store.Companies_Warehouses().DeleteCompanies_WarehousesById(&Companies_Warehouses)
	if err != nil {
		api.logger.Info("Troubles while connections to the warehouse database:", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 200,
		Message:    "delete complited",
		IsError:    false,
	}
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(msg)

}
