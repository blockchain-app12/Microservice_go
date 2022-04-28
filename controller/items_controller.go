package controller

import(
	"net/http"
	"fmt"
	"Golang-Microservice/domain/items"
	"Golang-Microservice/service"
	"Golang-Microservice/utils"
	"Golang-Microservice/logger"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"log"
)

func Create() httprouter.Handle {
	return func (w http.ResponseWriter,r *http.Request,  p httprouter.Params) {
	  var item items.Item

		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			restErr := utils.NewBadRequestError("invalid request body")
			utils.ResponseError(w, *restErr)
			return
		}

		defer r.Body.Close()

		if err:= json.Unmarshal(requestBody, &item); err != nil {
			resErr := utils.NewBadRequestError("invalid json body")
			utils.ResponseError(w, *resErr)
			return
		}

		result, createErr := service.ItemsService.Create(item)
		if createErr != nil {
			fmt.Println("ERROR::::", err)
			utils.ResponseError(w, *createErr)
			return
		}
		
		response := utils.BuildResponse(w, http.StatusOK,result)
		json.NewEncoder(w).Encode(response)
	}
}

func  GetAll() httprouter.Handle {
	return func (w http.ResponseWriter,r *http.Request,  p httprouter.Params) {
		result, getErr := service.ItemsService.Get()
		if getErr != nil {
			utils.ResponseError(w, *getErr)
			return
		}
		response := utils.BuildResponse(w, http.StatusOK,result)
		json.NewEncoder(w).Encode(response)
	}
}

func GetItem() httprouter.Handle {
	return func (w http.ResponseWriter,r *http.Request,  p httprouter.Params) {
		queryParam := r.URL.Query().Get("item_id")
			if queryParam == "" {
				fmt.Println("empty query param", queryParam)
				w.WriteHeader(http.StatusNotFound)
				return
			}
		 item , getItemErr := service.ItemsService.GetByID(queryParam)
		 if getItemErr != nil {
			utils.ResponseError(w, *getItemErr)
			return
		 }
		 response := utils.BuildResponse(w, http.StatusOK,item)
		json.NewEncoder(w).Encode(response)
	}
}

func UpdateItem() httprouter.Handle{
	return func (w http.ResponseWriter,r *http.Request,  p httprouter.Params) {
		var items items.Item
		queryParam :=  p.ByName("item_id")
		if queryParam == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewDecoder(r.Body).Decode(&items)
		items.Item_ID = queryParam
		isPartial := r.Method == http.MethodPatch

		result, updateErr := service.ItemsService.Update(isPartial, items)
		if updateErr != nil {
			logger.Info("Error::::")
			utils.ResponseError(w, *updateErr)
			return
		}
		response := utils.BuildResponse(w, http.StatusOK,result)
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteItem() httprouter.Handle{
	return func(w http.ResponseWriter,r *http.Request,  p httprouter.Params) {
		queryParam := r.URL.Query().Get("item_id")
		if queryParam == "" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		result, err := service.ItemsService.Delete(queryParam)
		if err != nil {
			log.Println("Error::::", err)
		}
		response := utils.BuildResponse(w, http.StatusOK,result)
		json.NewEncoder(w).Encode(response)
	}
}