package engine_service

import (
	"context"
	"errors"

	"carazone/models"
	"carazone/store"
)

type EngineService struct{
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService{
	return &EngineService{
		store: store,
	}
}

func (s *EngineService) GetEngineByID(ctx context.Context, id string) (*models.Engine, error){
	engine, err := s.store.EngineById(ctx, id)
	if err != nil{
		return nil, err
	}
	return &engine, nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error) {
	engine := models.Engine{
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}

	if err := models.ValidateEngineRequest(engine); err != nil {
		return nil, err
	}

	if s.store == nil {
		return nil, errors.New("store is not initialized")
	}
	createdEngine, err := s.store.EngineCreated(ctx, engineReq) 
	if err != nil {
		return nil, err
	}

	return &createdEngine, nil 
}

func (s *EngineService) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest)(*models.Engine, error){
	engine := models.Engine{
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}
	if err:= models.ValidateEngineRequest(engine); err != nil{
		return nil, err
	}
	updatedEngine, err := s.store.EngineUpdate(ctx, id, engineReq)
	if err != nil{
		return nil, err 
	}
	return &updatedEngine, nil
}

func (s * EngineService) DeleteEngine(ctx context.Context, id string) (*models.Engine, error){
	deletedEngine, err := s.store.EngineDelete(ctx, id)
	if err != nil{
		return nil, err
	}
	return &deletedEngine, nil
}