package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
	"github.com/gorilla/mux"
)

func (api *APIServer) Post(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Warehouse_cells POST /warehouses_cells")
	var Warehouses_cells models.Warehouses_cells
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
	var Warehouse_cells models.Warehouses_cells
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
