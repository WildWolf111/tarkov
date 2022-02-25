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

func (api *APIServer) GetAllStocs(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	stoc, err := api.store.Stoc().SelectAll()

	if err != nil {
		api.logger.Info(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to accessing stocs in database. Try later",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	api.logger.Info("Get All Stocs GET /stocs")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(stoc)
}

func (api *APIServer) PostStocs(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Stocs POST /stocs")
	var stoc models.Stocs
	err := json.NewDecoder(req.Body).Decode(&stoc)
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
	fmt.Println(stoc)
	a, err := api.store.Stoc().Create(&stoc)
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

func (api *APIServer) GetStocById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Stoc by ID /api/v1/stocs/{id}")
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
	stoc, ok, err := api.store.Stoc().FindStocById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (stocs) with id. err:", err)
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
		api.logger.Info("Can not find stoc with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "Stoc with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(stoc)

}

func (api *APIServer) DeleteStocById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Stoc by Id DELETE /api/v1/stocs/{id}")
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

	_, ok, err := api.store.Stoc().FindStocById(id)
	if err != nil {
		api.logger.Info("Troubles while accessing database table (stocs) with id. err:", err)
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
		api.logger.Info("Can not find stoc with that ID in database")
		msg := Message{
			StatusCode: 404,
			Message:    "stoc with that ID does not exists in database.",
			IsError:    true,
		}

		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	_, err = api.store.Stoc().DeleteById(id)
	if err != nil {
		api.logger.Info("Troubles while deleting database elemnt from table (stocs) with id. err:", err)
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
		Message:    fmt.Sprintf("Stoc with ID %d successfully deleted.", id),
		IsError:    false,
	}
	json.NewEncoder(writer).Encode(msg)
}

//warehouse update
func (api *APIServer) UpdateStocById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating Stoc ...")
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

	var newStoc models.Stocs

	err = json.NewDecoder(request.Body).Decode(&newStoc)
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
	newStoc.ID = id
	a, err := api.store.Stoc().UpdateStocById(&newStoc)
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
