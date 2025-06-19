package services

import (
    "github.com/saarthi123/saarthi-backend/models"
    "gorm.io/gorm"
)

type RoleService struct {
    db *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
    return &RoleService{db: db}
}

func (s *RoleService) CreateRole(role *models.Role) error {
    return s.db.Create(role).Error
}

func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
    var role models.Role
    if err := s.db.Preload("Permissions").First(&role, id).Error; err != nil {
        return nil, err
    }
    return &role, nil
}

func (s *RoleService) UpdateRole(role *models.Role) error {
    return s.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(role).Error
}

func (s *RoleService) DeleteRole(id uint) error {
    return s.db.Delete(&models.Role{}, id).Error
}

func (s *RoleService) ListRoles() ([]models.Role, error) {
    var roles []models.Role
    err := s.db.Preload("Permissions").Find(&roles).Error
    return roles, err
}
