package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct{
	EngineID uuid.UUID `json:"engineid"`
	Displacement int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfCylinders"`
	CarRange int64 `json:"carRange"`

}
type RequestEngine struct{
	Displacement int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfCylinders"`
	CarRange int64 `json:"carRange"`
}

func ValidateEngineRequest(engine Engine)error{
	if err :=validateDisplacement(engine.Displacement); err!=nil{
		return err
	}
	if err :=validateNoOfCylinders(engine.NoOfCylinders); err!=nil{
		return err
	}
	if err :=validateCarRange(engine.CarRange); err!=nil{
		return err
	}
	return nil
}

func validateDisplacement(displacement int64) error {
	if displacement <= 0{
		return errors.New("displacement must be greater than zero")
	}
	return nil
}

func validateNoOfCylinders(NoOfCylinders int64) error {
	if NoOfCylinders <= 0{
		return errors.New("noOfCylinder must be greater than zero")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange must be greater than zero")
	}
	return nil
}
