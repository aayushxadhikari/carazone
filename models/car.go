package models

import (
	"errors"
	"strconv"
	"time"
	"fmt"
	"github.com/google/uuid"
)

type Car struct{
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Year string `json:"year"`
	Brand string `json:"brand"`
	FuelType string `json:"fuelType"`
	Engine Engine  `json:"engine"`
	Price float64  `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type CarRequest struct{
	Name string `json:"name"`
	Year string `json:"year"`
	Brand string `json:"brand"`
	FuelType string `json:"fuelType"`
	Engine Engine  `json:"engine"`
	Price float64  `json:"price"`
}

func validateName(name string) error {
	if name == ""{
		return errors.New("name is required")
	}
	return nil

}

func validateYear(year string) error {
	if year == ""{
		return errors.New("year is required")
	}
	_, err := strconv.Atoi(year)
	if err != nil{
		return errors.New("year must be a valid number")
	}
	currentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)
	if yearInt < 1886 || yearInt > currentYear {
		return fmt.Errorf("year must be between 1886 and %d", currentYear)
	}
	return nil
}

func validateBrand(brand string) error {
	if brand == ""{
		return errors.New("brand is required")
	}
	return nil
}

func validateFuelType(fuelType string)error{
	validateFuelTypes:= []string{
		"Petrol",
		"Diesel",
		"Electric",
		"Hybrid",
	}
	for _, validType := range validateFuelTypes{
		if fuelType == validType{
			return nil
		}
	}

	return errors.New("FuelType must be one of : Petrol, Diesel, Electric, Hybrid")
}

func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil{
		return errors.New("EngineId is required")
	}
	if engine.Displacement <= 0 {
		return errors.New("displacement must be greated than zero")
	}
	if engine.NoOfCylinders <= 0 {
		return errors.New("noOfCylinders must be greater than zero")
	}
	if engine.CarRange <= 0 {
		return errors.New("carRange must be greater than zero")
	}
	return nil
}

func validatePrice(price float64) error {
	if price <= 0{
		return errors.New("price must be greater than zero")
	}
	return nil
}

// Run all validations
func ValidateCarRequest(car CarRequest) error{
	if err := validateName(car.Name); err != nil{
		return err
	}
	if err := validateYear(car.Year); err!= nil{
		return err
	}
	if err := validateBrand(car.Brand); err!= nil{
		return err
	}
	if err := validateFuelType(car.FuelType); err!= nil{
		return err
	}
	if err := validateEngine(car.Engine); err!= nil{
		return err
	}
	if err := validatePrice(car.Price); err!= nil{
		return err
	}
	return nil
}