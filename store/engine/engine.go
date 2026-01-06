package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gloonch/CarZone/models"
	"github.com/google/uuid"
)

type EngineStore struct {
	db *sql.DB
}

func NewEngineStore(db *sql.DB) *EngineStore {
	return &EngineStore{
		db: db,
	}
}

func (e EngineStore) EngineByID(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				fmt.Printf("Transaction rollback error: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v", cmErr)
			}
		}
	}()

	tx.QueryRowContext(ctx, "SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id=$1", id).Scan(
		&engine.EngineID, &engine.Displacement, &engine.NoOfCylinders, &engine.CarRange,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, errors.New("Engine does not exist")
		}
		return engine, err
	}
	return engine, nil
}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				fmt.Printf("Transaction rollback error: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v", cmErr)
			}
		}
	}()

	engineID := uuid.New()

	_, err = tx.ExecContext(ctx,
		"INSERT INTO engine ( id, displacement, no_of_cylinders, car_range ) VALUES ( $1, $2, $3, $4 )",
		engineID, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange)
	if err != nil {
		return models.Engine{}, err
	}
	engine := models.Engine{
		EngineID:      engineID,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}

	return engine, nil
}

func (e EngineStore) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {

	engineID, err := uuid.Parse(id)
	if err != nil {
		return models.Engine{}, fmt.Errorf("Invalid engine ID: %s", id)
	}

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				fmt.Printf("Transaction rollback error: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v", cmErr)
			}
		}
	}()

	results, err := tx.ExecContext(ctx,
		"UPDATE engine SET displacement = $1, no_of_cylinders = $2, car_range = $3 WHERE id = $4",
		engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange, engineID)
	if err != nil {
		return models.Engine{}, err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("No rows updated")
	}

	engine := models.Engine{
		EngineID:      engineID,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}

	return engine, nil
}

func (e EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {

	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				fmt.Printf("Transaction rollback error: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				fmt.Printf("Transaction commit error: %v", cmErr)
			}
		}
	}()

	err = tx.QueryRowContext(ctx, "SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id=$1", id).Scan(
		&engine.EngineID, &engine.Displacement, &engine.NoOfCylinders, &engine.CarRange,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, errors.New("Engine does not exist")
		}
		return engine, err
	}

	result, err := tx.ExecContext(ctx, "DELETE FROM engine WHERE id=$1", id)
	if err != nil {
		return models.Engine{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Engine{}, err
	}
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("No rows deleted")
	}
	return engine, nil
}
