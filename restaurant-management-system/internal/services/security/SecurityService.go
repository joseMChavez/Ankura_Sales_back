package services

import (
	"restaurant-management-system/internal/models"
	"restaurant-management-system/internal/repository/security"
)

// SecurityServiceImpl implementa SecurityOperations
type SecurityServiceImpl struct {
	Repo security.SecurityOperations
}

// NewSecurityService crea una nueva instancia de SecurityServiceImpl
func NewSecurityService(repo security.SecurityOperations) security.SecurityOperations {
	return &SecurityServiceImpl{Repo: repo}
}

// Implementación de los métodos

// Permisos
func (s *SecurityServiceImpl) CreatePermission(permission models.Permission) error {
	return s.Repo.CreatePermission(permission)
}

func (s *SecurityServiceImpl) GetAllPermissions() ([]models.Permission, error) {
	return s.Repo.GetAllPermissions()
}

// Roles
func (s *SecurityServiceImpl) CreateRole(role models.Role) error {
	return s.Repo.CreateRole(role)
}

func (s *SecurityServiceImpl) GetAllRoles() ([]models.Role, error) {
	return s.Repo.GetAllRoles()
}

// Perfiles
func (s *SecurityServiceImpl) CreateProfile(profile models.Profile) error {
	return s.Repo.CreateProfile(profile)
}

func (s *SecurityServiceImpl) GetAllProfiles() ([]models.Profile, error) {
	return s.Repo.GetAllProfiles()
}

// Usuarios
func (s *SecurityServiceImpl) CreateUser(user models.User) error {
	return s.Repo.CreateUser(user)
}

func (s *SecurityServiceImpl) GetUserByID(id int) (models.User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *SecurityServiceImpl) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}
