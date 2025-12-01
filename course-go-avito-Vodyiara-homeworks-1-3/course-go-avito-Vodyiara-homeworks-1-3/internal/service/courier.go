package service

import (
	"avito_project/course-go-avito-Vodyiara/internal/repository"
	"context"

	"avito_project/course-go-avito-Vodyiara/internal/model"
)

type courierService struct {
	repo repository.CourierRepository
}

func NewCourierService(repo repository.CourierRepository) CourierService {
	return &courierService{
		repo: repo,
	}
}

func (s *courierService) CreateCourier(ctx context.Context, req *model.CreateCourierRequest) (*model.Courier, error) {
	// Валидация
	if err := req.Validate(); err != nil {
		return nil, err
	}

	courier, err := s.repo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return courier, nil
}

func (s *courierService) GetCourier(ctx context.Context, id int64) (*model.Courier, error) {
	if id <= 0 {
		return nil, model.ErrInvalidID
	}
	courier, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return courier, nil
}

func (s *courierService) GetAllCouriers(ctx context.Context) ([]*model.Courier, error) {
	couriers, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return couriers, nil
}

func (s *courierService) UpdateCourier(ctx context.Context, req *model.UpdateCourierRequest) error {

	if err := req.Validate(); err != nil {
		return err
	}

	exists, err := s.repo.Exists(ctx, *req.ID)
	if err != nil {
		return err
	}
	if !exists {
		return model.ErrCourierNotFound
	}

	if err := s.repo.Update(ctx, req); err != nil {
		return err
	}

	return nil
}
