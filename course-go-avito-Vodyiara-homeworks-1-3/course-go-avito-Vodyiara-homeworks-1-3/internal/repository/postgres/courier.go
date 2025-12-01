package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"avito_project/course-go-avito-Vodyiara/internal/model"
)

type courierRepository struct {
	db *sql.DB
}

func NewCourierRepository(db *sql.DB) *courierRepository {
	return &courierRepository{db: db}
}

func (r *courierRepository) Create(ctx context.Context, req *model.CreateCourierRequest) (*model.Courier, error) {
	query := `
		INSERT INTO couriers (name, phone, status)
		VALUES ($1, $2, $3)
		RETURNING id, name, phone, status, created_at, updated_at
	`

	courier := &model.Courier{}
	err := r.db.QueryRowContext(ctx, query, req.Name, req.Phone, req.Status).Scan(
		&courier.ID,
		&courier.Name,
		&courier.Phone,
		&courier.Status,
		&courier.CreatedAt,
		&courier.UpdatedAt,
	)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, model.ErrPhoneAlreadyExists
		}
		return nil, fmt.Errorf("failed to create courier: %w", err)
	}

	return courier, nil
}

func (r *courierRepository) GetByID(ctx context.Context, id int64) (*model.Courier, error) {
	query := `
		SELECT id, name, phone, status, created_at, updated_at
		FROM couriers
		WHERE id = $1
	`

	courier := &model.Courier{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&courier.ID,
		&courier.Name,
		&courier.Phone,
		&courier.Status,
		&courier.CreatedAt,
		&courier.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrCourierNotFound
		}
		return nil, fmt.Errorf("failed to get courier: %w", err)
	}

	return courier, nil
}

func (r *courierRepository) GetAll(ctx context.Context) ([]*model.Courier, error) {
	query := `
		SELECT id, name, phone, status, created_at, updated_at
		FROM couriers
		ORDER BY id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get couriers: %w", err)
	}
	defer rows.Close()

	couriers := []*model.Courier{}
	for rows.Next() {
		courier := &model.Courier{}
		if err := rows.Scan(
			&courier.ID,
			&courier.Name,
			&courier.Phone,
			&courier.Status,
			&courier.CreatedAt,
			&courier.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan courier: %w", err)
		}
		couriers = append(couriers, courier)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return couriers, nil
}

func (r *courierRepository) Update(ctx context.Context, req *model.UpdateCourierRequest) error {
	query := "UPDATE couriers SET updated_at = now()"
	args := []interface{}{}
	argIndex := 1

	if req.Name != nil {
		query += fmt.Sprintf(", name = $%d", argIndex)
		args = append(args, *req.Name)
		argIndex++
	}
	if req.Phone != nil {
		query += fmt.Sprintf(", phone = $%d", argIndex)
		args = append(args, *req.Phone)
		argIndex++
	}
	if req.Status != nil {
		query += fmt.Sprintf(", status = $%d", argIndex)
		args = append(args, *req.Status)
		argIndex++
	}

	query += fmt.Sprintf(" WHERE id = $%d", argIndex)
	args = append(args, *req.ID)

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return model.ErrPhoneAlreadyExists
		}
		return fmt.Errorf("failed to update courier: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return model.ErrCourierNotFound
	}

	return nil
}

func (r *courierRepository) Exists(ctx context.Context, id int64) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM couriers WHERE id = $1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check courier existence: %w", err)
	}

	return exists, nil
}
