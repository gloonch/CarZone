package engine

import (
	"context"

	"github.com/gloonch/CarZone/models"
	"github.com/gloonch/CarZone/store"
	"go.opentelemetry.io/otel"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}
}

func (e *EngineService) EngineByID(ctx context.Context, id string) (*models.Engine, error) {
	tracer := otel.Tracer("engine-service")
	ctx, span := tracer.Start(ctx, "EngineByID-Service")
	defer span.End()

	engine, err := e.store.EngineByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func (e *EngineService) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error) {
	tracer := otel.Tracer("engine-service")
	ctx, span := tracer.Start(ctx, "CreateEngine-Service")
	defer span.End()

	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}

	createdEngine, err := e.store.CreateEngine(ctx, engineReq)
	if err != nil {
		return nil, err
	}
	return &createdEngine, nil
}

func (e *EngineService) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error) {

	tracer := otel.Tracer("engine-service")
	ctx, span := tracer.Start(ctx, "UpdateEngine-Service")
	defer span.End()

	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}
	updatedEngine, err := e.store.UpdateEngine(ctx, id, engineReq)
	if err != nil {
		return nil, err
	}
	return &updatedEngine, nil
}

func (e *EngineService) DeleteEngine(ctx context.Context, id string) (*models.Engine, error) {
	tracer := otel.Tracer("engine-service")
	ctx, span := tracer.Start(ctx, "DeleteEngine-Service")
	defer span.End()

	deletedEngine, err := e.store.DeleteEngine(ctx, id)
	if err != nil {
		return nil, err
	}
	return &deletedEngine, nil
}
