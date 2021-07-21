package services

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

// UseService service layer
type UseService struct {
	logger     lib.Logger
	repository repository.UseRepository
}

// NewUseService creates a new useservice
func NewUseService(logger lib.Logger, repository repository.UseRepository) UseService {
	return UseService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s UseService) WithTrxUse(trxHandle *gorm.DB) UseService {
	s.repository = s.repository.WithTrxUse(trxHandle)
	return s
}

// GetOneUse gets one use
func (s UseService) GetOneUse(id uint) (use models.User, err error) {
	return use, s.repository.Find(&use, id).Error
}

// GetAllUse get all the use
func (s UseService) GetAllUse() (uses []models.User, err error) {
	return uses, s.repository.Find(&uses).Error
}

// CreateUse call to create the use
func (s UseService) CreateUse(use models.User) error {
	return s.repository.Create(&use).Error
}

// UpdateUse updates the use
func (s UseService) UpdateUse(use models.User) error {
	return s.repository.Save(&use).Error
}

// DeleteUse deletes the use
func (s UseService) DeleteUse(id uint) error {
	return s.repository.Delete(&models.User{}, id).Error
}
