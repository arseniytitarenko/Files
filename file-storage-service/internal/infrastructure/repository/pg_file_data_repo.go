package repository

import (
	"files/internal/domain"
	"gorm.io/gorm"
)

type PgFileDataRepo struct {
	db *gorm.DB
}

func NewPgFileDataRepo(db *gorm.DB) *PgFileDataRepo {
	return &PgFileDataRepo{db: db}
}

func (r *PgFileDataRepo) GetFileData(fileDataID string) (*domain.FileData, error) {
	var fileData domain.FileData
	if err := r.db.Where("id = ?", fileDataID).First(&fileData).Error; err != nil {
		return nil, err
	}
	return &fileData, nil
}

func (r *PgFileDataRepo) NewFileData(fileData *domain.FileData) error {
	return r.db.Create(fileData).Error
}
