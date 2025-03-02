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
