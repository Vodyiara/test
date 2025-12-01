package repository

import (
	"context"

	"avito_project/course-go-avito-Vodyiara/internal/model"
)

type CourierRepository interface {
	Create(ctx context.Context, courier *model.CreateCourierRequest) (*model.Courier, error)

	GetByID(ctx context.Context, id int64) (*model.Courier, error)

	GetAll(ctx context.Context) ([]*model.Courier, error)

	Update(ctx context.Context, req *model.UpdateCourierRequest) error

	Exists(ctx context.Context, id int64) (bool, error)
}
