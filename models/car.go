package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuelType"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuelType"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func ValidateCarRequest(carRequest CarRequest) error {
	if err := ValidateName(carRequest.Name); err != nil {
		return err
	}
	if err := ValidateYear(carRequest.Year); err != nil {
		return err
	}
	if err := ValidateBrand(carRequest.Brand); err != nil {
		return err
	}
	if err := ValidateFuelType(carRequest.FuelType); err != nil {
		return err
	}
	if err := ValidateEngine(carRequest.Engine); err != nil {
		return err
	}
	if err := ValidatePrice(carRequest.Price); err != nil {
		return err
	}
	return nil
}

func ValidateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

func ValidateYear(year string) error {
	if year == "" {
		return errors.New("year is required")
	}
	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a valid number")
	}
	currentYear := time.Now().Year()
	yearInt, err := strconv.Atoi(year)
	if yearInt > currentYear || yearInt < 1900 {
		return errors.New(fmt.Sprintf("year must be between 1900 and %d", currentYear))
	}
	return nil
}

func ValidateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

func ValidateFuelType(fuelType string) error {
	validateFuelType := []string{"Persol", "Diesel", "Electric", "Hybrid"}
	for _, validType := range validateFuelType {
		if fuelType == validType {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("FuelType must be one of %v", validateFuelType))
}

func ValidateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("EngineID is required")
	}
	if engine.Displacement <= 0 {
		return errors.New("displacement must be greater than zero")
	}
	if engine.NoOfCylinders <= 0 {
		return errors.New("NoOfCylinders must be greater than zero")
	}
	if engine.CarRange <= 0 {
		return errors.New("CarRange must be greater than zero")
	}
	return nil
}

func ValidatePrice(price float64) error {
	if price <= 0 {
		return errors.New("Price must be greater than zero")
	}
	return nil
}
