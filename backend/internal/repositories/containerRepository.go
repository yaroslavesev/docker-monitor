package repositories

import (
	"backend/internal/db"
	"backend/internal/models"
	"errors"
)

// CRUD-repository

// Create - creates new record in DB and returns error
func Create(container *models.Container) error {
	err := db.DB.Create(container).Error
	return err
}

// GetAll - returns slice of containers and error
func GetAll() (*[]models.Container, error) {
	var containers []models.Container
	result := db.DB.Find(&containers)
	if result.Error != nil {
		return nil, result.Error
	}
	return &containers, nil
}

// FindById - returns container by id and error
func FindById(id int) (*models.Container, error) {
	var container models.Container
	result := db.DB.First(&container, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &container, nil
}

// Save - saves container to DB and returns error
func Save(container *models.Container) error {
	result := db.DB.Save(container)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete - deletes a container by its ID
func Delete(id int) error {
	result := db.DB.Delete(models.Container{}, id)
	if result.RowsAffected == 0 {
		return errors.New("no container found with the given ID")
	}
	return result.Error
}
