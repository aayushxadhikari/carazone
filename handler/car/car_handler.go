package car

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/aayushxadhikari/carazone/models"
	"github.com/aayushxadhikari/carazone/service"
	"github.com/gorilla/mux"
)
 

type CarHandler struct{
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler{
	return &CarHandler{
		service: service,
	}
}

func (h *CarHandler) GetCarById(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	vars := mux.Vars(r)
	id, exists := vars["id"]

	// Check if the ID exists
	if !exists || id == "" {
		http.Error(w, `{"error": "car ID is required"}`, http.StatusBadRequest)
		return
	}

	resp, err := h.service.GetCarById(ctx, id)
	if err != nil{
		http.Error(w, `{"error": "failed to retrieve car"}`, http.StatusInternalServerError)
		log.Println("Error:", err)
		return 
	}
	body, err := json.Marshal(resp)
	if err !=nil{
		http.Error(w, `{"error": "failed to encode response"}`, http.StatusInternalServerError)
		log.Println("Error:", err)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	// write the response body
	if _, err := w.Write(body);err != nil{
		log.Println("Error Writing Response:", err)
	}
}

func (h *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	brand := r.URL.Query().Get("brand")
	isEngine := r.URL.Query().Get("isEngine") == "true"

	resp, err := h.service.GetCarByBrand(ctx, brand, isEngine)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error:", err)
		return
	}

	body, err := json.Marshal(resp)
	if err!= nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write the response body
	if _, err := w.Write(body);err!= nil{
		log.Println("Error Writing Response:", err)
	}
}

func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil{
		log.Println("Err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return	
	}


	var carReq models.CarRequest
	err = json.Unmarshal(body, carReq)
	if err != nil{
		log.Println("Error while Unmarshalling request.")
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	createdCar, err := h.service.CreateCar(ctx, &carReq)
	if err != nil{
		log.Println("Error Creating Car:",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(createdCar)
	w.WriteHeader(http.StatusInternalServerError)

	// write the reponse body
	_,_ = w.Write(responseBody)
}

func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil{
		log.Println("Error Reading Request body:",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var carReq models.CarRequest
	err = json.Unmarshal(body, &carReq)
	if err != nil{
		log.Println("Error while Unmarshalling Request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedCar, err := h.service.UpdateCar(ctx, id, &carReq)
	if err != nil{
		log.Println("Error Updating car:",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody, err:= json.Marshal(updatedCar)
	if err != nil{
		log.Println("Error while Marshalling Response body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	//write the response body
	_,_ = w.Write(resBody)
}

func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	deletedCar , err := h.service.DeleteCar(ctx, id)
	if err!= nil{
		log.Println("Error while deleting the car:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	body, err := json.Marshal(deletedCar)
	if err != nil{
		log.Println("Error while Marshalling the response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	// write the response body
	if _, err := w.Write(body); err != nil{
		log.Println("Error Writing Response:", err)
	}

}