package services

import (
    "github.com/saarthi123/saarthi-backend/models"
    "gorm.io/gorm"
)

type DashboardService struct {
    db *gorm.DB
}

func NewDashboardService(db *gorm.DB) *DashboardService {
    return &DashboardService{db: db}
}

func (s *DashboardService) Create(dashboard *models.Dashboard) error {
    return s.db.Create(dashboard).Error
}

func (s *DashboardService) GetByUserName(userName string) (*models.Dashboard, error) {
    var dashboard models.Dashboard
    if err := s.db.First(&dashboard, "user_name = ?", userName).Error; err != nil {
        return nil, err
    }
    return &dashboard, nil
}

func (s *DashboardService) Update(dashboard *models.Dashboard) error {
    return s.db.Save(dashboard).Error
}

func (s *DashboardService) Delete(userName string) error {
    return s.db.Delete(&models.Dashboard{}, "user_name = ?", userName).Error
}

func (s *DashboardService) List() ([]models.Dashboard, error) {
    var dashboards []models.Dashboard
    err := s.db.Find(&dashboards).Error
    return dashboards, err
}
