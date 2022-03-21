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

//POST
func (api *APIServer) Post(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Warehouse_cells POST /warehouses_cells")
	var Warehouses_cells models.Warehouse_cell
	err := json.NewDecoder(req.Body).Decode(&Warehouses_cells)
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
	fmt.Println(Warehouses_cells)
	a, err := api.store.Warehouses_cells().Post(&Warehouses_cells)
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

//delete
func (api *APIServer) DeleteWarehouses_cellsById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete CompaniesWarehousesById DELETE /warehouses_cells/delete/{id}")
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
	var Warehouse_cells models.Warehouse_cell
	fmt.Println(Warehouse_cells)
	err = api.store.Warehouses_cells().DeleteCompanies_WarehousesById(id)
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

//GET ALL

func (api *APIServer) GetAllWarehouses_cells(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)

	Warehouse_cells, err := api.store.Warehouses_cells().GetAll()
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
	api.logger.Info("Get All Company GET /Warehouses_cells")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(Warehouse_cells)
}

//GET Warehouses_cells By ID

func (api *APIServer) GetWarehouses_cellsById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Warehouse by ID /api/v1/warehouses_cells/{id}")
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
	warehouses_cells, ok, err := api.store.Warehouses_cells().GetByID(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (warehouses_cells) with id. err:", err)
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
	json.NewEncoder(writer).Encode(warehouses_cells)

}

//UPDATE

func (api *APIServer) UpdateWarehouses_cellsById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating Warehouses_cells ...")
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

	var newWarehouse_cells models.Warehouse_cell

	err = json.NewDecoder(request.Body).Decode(&newWarehouse_cells)
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
	newWarehouse_cells.ID = id
	a, err := api.store.Warehouses_cells().UpdateWarehouses_cells(&newWarehouse_cells)
	if err != nil {
		api.logger.Info("Troubles while connections to the stoc database:", err)
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
