package database

import (
	"github.com/markitos/markitos-svc-boilerplate/internal/domain"
	"gorm.io/gorm"
)

type BoilerPostgresRepository struct {
	db *gorm.DB
}

func NewBoilerPostgresRepository(db *gorm.DB) BoilerPostgresRepository {
	return BoilerPostgresRepository{db: db}
}

func (r *BoilerPostgresRepository) Create(boilerplate *domain.Boilerplate) error {
	return r.db.Create(boilerplate).Error
}

func (r *BoilerPostgresRepository) Delete(id *domain.BoilerplateId) error {
	_, err := r.One(id)
	if err != nil {
		return domain.BoilerplateNotFoundError
	}

	return r.db.Delete(&domain.Boilerplate{}, "id = ?", id.Value()).Error
}

func (r *BoilerPostgresRepository) Update(boilerplate *domain.Boilerplate) error {
	return r.db.Save(boilerplate).Error
}

func (r *BoilerPostgresRepository) One(id *domain.BoilerplateId) (*domain.Boilerplate, error) {
	var boiler domain.Boilerplate
	if err := r.db.First(&boiler, "id = ?", id.Value()).Error; err != nil {
		return nil, domain.BoilerplateNotFoundError
	}
	return &boiler, nil
}
