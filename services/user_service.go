package services

import (

    "github.com/saarthi123/saarthi-backend/models"
    "gorm.io/gorm"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db: db}
}

func (s *UserService) CreateUser(user *models.User) error {
    return s.db.Create(user).Error
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    var user models.User
    if err := s.db.Preload("Role.Permissions").First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (s *UserService) UpdateUser(user *models.User) error {
    return s.db.Save(user).Error
}

func (s *UserService) DeleteUser(id uint) error {
    return s.db.Delete(&models.User{}, id).Error
}

func (s *UserService) ListUsers() ([]models.User, error) {
    var users []models.User
    err := s.db.Preload("Role").Find(&users).Error
    return users, err
}
