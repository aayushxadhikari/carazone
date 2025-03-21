package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/aayushxadhikari/carazone/models"
	"github.com/aayushxadhikari/carazone/service"
	"github.com/gorilla/mux"
)

type EngineHandler struct{
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler{
	return &EngineHandler{
		service: service,
	}
}

func (h *EngineHandler) GetEngineByID(w http.ResponseWriter, r *http.Request){
	ctx:= r.Context()
	vars := mux.Vars(r)
	id, exists := vars["id"]

	// check if the ID exists
	if !exists || id == ""{
		http.Error(w, `{"error":"car ID is required"}`, http.StatusBadRequest)
		return 
	}

	resp, err := h.service.GetEngineByID(ctx, id)
	if err!= nil{
		http.Error(w, `{"error": "failed to retrieve car"}`, http.StatusInternalServerError)
		log.Println("Error:", err)
		return 
	}
	body, err:= json.Marshal(resp)
	if err!=nil{
		http.Error(w, `{"error": "failed to encode response"}`, http.StatusInternalServerError)
		log.Println("Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// write the response body
	if _, err := w.Write(body); err!= nil{
		log.Println("Error writing response:", err)
	}
}

func (h *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err!= nil{
		log.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, engineReq)
	if err != nil{
		log.Println("Error while Unmarshalling request.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	createdEngine, err := h.service.CreateEngine(ctx, &engineReq)
	if err != nil{
		log.Println("Error Creating engine:",err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseBody, err := json.Marshal(createdEngine)
	w.WriteHeader(http.StatusInternalServerError)

	// write the response body
	_, _ = w.Write(responseBody)

}

func (h *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil{
		log.Println("Error Reading Request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, &engineReq)
	if err != nil{
		log.Println("Error while unmarshalling Request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	updatedEngine, err := h.service.UpdateEngine(ctx, id, &engineReq)
	if err != nil{
		log.Println("Error while Marshalling response body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resBody, err := json.Marshal(updatedEngine)
	if err != nil{
		log.Println("Error while Marshalling Response body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	// write the response body
	_,_= w.Write(resBody)
}

func (h *EngineHandler) DeleteCar(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	deletedEngine, err := h.service.DeleteEngine(ctx, id)
	if err!= nil{
		log.Println("Error while deleting the car:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	body, err := json.Marshal(deletedEngine)
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
