package service

import (
	"context"

	"avito_project/course-go-avito-Vodyiara/internal/model"
)

type CourierService interface {
	CreateCourier(ctx context.Context, req *model.CreateCourierRequest) (*model.Courier, error)

	GetCourier(ctx context.Context, id int64) (*model.Courier, error)

	GetAllCouriers(ctx context.Context) ([]*model.Courier, error)

	UpdateCourier(ctx context.Context, req *model.UpdateCourierRequest) error
}
