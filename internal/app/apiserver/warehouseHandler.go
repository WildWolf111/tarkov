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

func (api *APIServer) GetAllWarehouse(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	warehouse, err := api.store.Warehouse().SelectAll()
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
	api.logger.Info("Get All Warehouses GET /warehouses")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(warehouse)
}

func (api *APIServer) PostWarhouse(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Warehouse POST /warehouses")
	var warehouse models.Warehouses
	err := json.NewDecoder(req.Body).Decode(&warehouse)
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
	fmt.Println(warehouse)
	a, err := api.store.Warehouse().Create(&warehouse)
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
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)

}

func (api *APIServer) GetWarehouseById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Warehouse by ID /api/v1/warehouses/{id}")
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
	article, ok, err := api.store.Warehouse().FindWarehouseById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (warehouse) with id. err:", err)
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

func (api *APIServer) DeleteWarehouseById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Warehouse by Id DELETE /api/v1/warehouses/{id}")
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

	_, ok, err := api.store.Warehouse().FindWarehouseById(id)
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

	_, err = api.store.Warehouse().DeleteById(id)
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
		Message:    fmt.Sprintf("Warehouses with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

//warehouse update
func (api *APIServer) UpdateWarehouseById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating Warehouse ...")
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

	var newWarehouse models.Warehouses

	err = json.NewDecoder(request.Body).Decode(&newWarehouse)
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
	newWarehouse.ID = id
	a, err := api.store.Warehouse().UpdateWarehouseById(&newWarehouse)
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
