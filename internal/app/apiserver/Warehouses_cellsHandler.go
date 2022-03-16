package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WildWolf111/StandarWebSrver2/internal/app/models"
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
	a, err := api.store.Warehouse().Post(&Warehouses_cells)
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
