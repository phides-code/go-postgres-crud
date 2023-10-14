package store

import (
	"time"
)

type Car struct {
	ID         int       `json:"id"`
	Make       string    `json:"make" binding:"required"`
	Model      string    `json:"model" binding:"required"`
	Year       int       `json:"year" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type PartialCar struct {
	ID         int       `json:"id"`
	Make       string    `json:"make"`
	Model      string    `json:"model"`
	Year       int       `json:"year"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

func AddCar(car *Car) error {
	_, err := db.Model(car).Returning("*").Insert()
	if err != nil {
		return err
	}
	return nil
}

func GetMatchingCars(partialCar *PartialCar) ([]*Car, error) {
	var matchingCars []*Car

	query := db.Model(&matchingCars)

	// Only include conditions if the fields in the partialCar are set.
	if partialCar.ID != 0 {
		query.Where("id = ?", partialCar.ID)
	}
	if partialCar.Make != "" {
		query.Where("make = ?", partialCar.Make)
	}
	if partialCar.Model != "" {
		query.Where("model = ?", partialCar.Model)
	}
	if partialCar.Year != 0 {
		query.Where("year = ?", partialCar.Year)
	}

	if err := query.Select(); err != nil {
		return nil, err
	}

	return matchingCars, nil
}

func UpdateCar(partialCar *PartialCar) (*Car, error) {
	// Ensure that the car with the specified ID exists in the database.
	existingCar := &Car{ID: partialCar.ID}
	err := db.Model(existingCar).WherePK().Select()
	if err != nil {
		return nil, err
	}

	// Apply updates to the existing car based on the provided partialCar.
	if partialCar.Make != "" {
		existingCar.Make = partialCar.Make
	}
	if partialCar.Model != "" {
		existingCar.Model = partialCar.Model
	}
	if partialCar.Year != 0 {
		existingCar.Year = partialCar.Year
	}

	// Update the car in the database.
	_, err = db.Model(existingCar).WherePK().Update()
	if err != nil {
		return nil, err
	}

	return existingCar, nil
}

func DeleteCar(carID int) error {
	// Ensure that the car with the specified ID exists in the database.
	existingCar := &Car{ID: carID}
	err := db.Model(existingCar).WherePK().Select()
	if err != nil {
		return err
	}

	// Delete the car from the database.
	_, err = db.Model(existingCar).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}
