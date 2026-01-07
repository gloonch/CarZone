package car

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gloonch/CarZone/models"
	"github.com/gloonch/CarZone/service"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
)

type CarHandler struct {
	service service.CarServiceInterface
}

func NewCarHandler(service service.CarServiceInterface) *CarHandler {
	return &CarHandler{
		service: service,
	}
}

func (handler *CarHandler) GetCarByID(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("car-handler")
	ctx, span := tracer.Start(r.Context(), "GetCarByID-Handler")
	defer span.End()

	//ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	res, err := handler.service.GetCarByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error getting car: %v", err)

		return
	}

	body, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error getting car: %v", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

func (handler *CarHandler) GetCarByBrand(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("car-handler")
	ctx, span := tracer.Start(r.Context(), "GetCarByBrand-Handler")
	defer span.End()

	//ctx := r.Context()
	brand := r.URL.Query().Get("brand")
	isEngine := r.URL.Query().Get("engine") == "true"

	res, err := handler.service.GetCarsByBrand(ctx, brand, isEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error getting car by brand: %v", err)

		return
	}
	body, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error getting car by brand: %v", err)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("Error writing response: %v", err)

		return
	}
}

func (handler *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("car-handler")
	ctx, span := tracer.Start(r.Context(), "CreateCar-Handler")
	defer span.End()
	//ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error reading body: %v", err)

		return
	}

	var carReq models.CarRequest
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error unmarshalling body: %v", err)

		return
	}

	createdCar, err := handler.service.CreateCar(ctx, &carReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error creating car: %v", err)

		return
	}
	body, err = json.Marshal(createdCar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshalling body: %v", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(body)
}

func (handler *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("car-handler")
	ctx, span := tracer.Start(r.Context(), "UpdateCar-Handler")
	defer span.End()

	//ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error reading body: %v", err)

		return
	}

	var carReq models.CarRequest
	err = json.Unmarshal(body, &carReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Error unmarshalling body: %v", err)

		return
	}
	updatedCar, err := handler.service.UpdateCar(ctx, id, &carReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error updating car: %v", err)

		return
	}
	body, err = json.Marshal(updatedCar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshalling body: %v", err)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func (handler *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("car-handler")
	ctx, span := tracer.Start(r.Context(), "DeleteCar-Handler")
	defer span.End()

	//ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	deletedCar, err := handler.service.DeleteCar(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error deleting car: %v", err)

		return
	}
	body, err := json.Marshal(deletedCar)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error marshalling body: %v", err)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func (handler *CarHandler) GetCar(w http.ResponseWriter, r *http.Request) {

}
