package services

import (
    "github.com/saarthi123/saarthi-backend/models"
    "gorm.io/gorm"
)

type StudentProgressService struct {
    db *gorm.DB
}

func NewStudentProgressService(db *gorm.DB) *StudentProgressService {
    return &StudentProgressService{db: db}
}

func (s *StudentProgressService) Create(sp *models.StudentProgress) error {
    return s.db.Create(sp).Error
}

func (s *StudentProgressService) GetByID(id string) (*models.StudentProgress, error) {
    var sp models.StudentProgress
    if err := s.db.First(&sp, "student_id = ?", id).Error; err != nil {
        return nil, err
    }
    return &sp, nil
}

func (s *StudentProgressService) Update(sp *models.StudentProgress) error {
    return s.db.Save(sp).Error
}

func (s *StudentProgressService) Delete(id string) error {
    return s.db.Delete(&models.StudentProgress{}, "student_id = ?", id).Error
}

func (s *StudentProgressService) List() ([]models.StudentProgress, error) {
    var sps []models.StudentProgress
    err := s.db.Find(&sps).Error
    return sps, err
}
