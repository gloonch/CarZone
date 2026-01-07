package engine

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gloonch/CarZone/models"
	"github.com/gloonch/CarZone/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel"
)

type EngineHandler struct {
	service service.EngineServiceInterface
}

func NewEngineHandler(service service.EngineServiceInterface) *EngineHandler {
	return &EngineHandler{
		service: service,
	}
}

func (handler *EngineHandler) GetEngineByID(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("engine-handler")
	ctx, span := tracer.Start(r.Context(), "GetEngineByID-Handler")
	defer span.End()

	//ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]

	res, err := handler.service.EngineByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}
	body, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(body)
	if err != nil {
		log.Println("Error writing response: ", err)
		return
	}
}

func (handler *EngineHandler) CreateEngine(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("engine-handler")
	ctx, span := tracer.Start(r.Context(), "CreateEngine-Handler")
	defer span.End()

	//ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)

		return
	}
	defer r.Body.Close()
	var engine models.Engine
	err = json.Unmarshal(body, &engine)
	if err != nil {
		log.Println("Error unmarshalling body: ", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	var engineReq *models.EngineRequest
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		log.Println("Error unmarshalling body: ", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	createdEngine, err := handler.service.CreateEngine(ctx, engineReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error unmarshalling body: ", err)

		return
	}

	res, err := json.Marshal(createdEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error marshalling body: ", err)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)

}

func (handler *EngineHandler) UpdateEngine(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("engine-handler")
	ctx, span := tracer.Start(r.Context(), "UpdateEngine-Handler")
	defer span.End()

	//ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error reading request body: ", err)

		return
	}
	defer r.Body.Close()

	var engineReq models.EngineRequest
	err = json.Unmarshal(body, &engineReq)
	if err != nil {
		log.Println("Error unmarshalling body: ", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	updateEngine, err := handler.service.UpdateEngine(ctx, id, &engineReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error unmarshalling body: ", err)

		return
	}

	res, err := json.Marshal(updateEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error marshalling body: ", err)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (handler *EngineHandler) DeleteEngine(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("engine-handler")
	ctx, span := tracer.Start(r.Context(), "DeleteEngine-Handler")
	defer span.End()

	//ctx := r.Context()
	params := mux.Vars(r)
	id := params["id"]

	deletedEngine, err := handler.service.DeleteEngine(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error unmarshalling body: ", err)

		response := map[string]string{"error": "Invalid ID or Engine not found"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)

		return
	}

	// check if deleted
	if deletedEngine.EngineID == uuid.Nil {
		w.WriteHeader(http.StatusNotFound)
		response := map[string]string{"error": "Engine not found"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)

		return
	}

	jsonResponse, err := json.Marshal(deletedEngine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Error marshalling deleted engine response: ", err)

		response := map[string]string{"error": "Internal server error"}
		jsonResponse, _ := json.Marshal(response)
		_, _ = w.Write(jsonResponse)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)

}
