package security

import (
	"restaurant-management-system/internal/models"
	// "restaurant-management-system/internal/security" // Removed as it does not exist

	"gorm.io/gorm"
)

// SecurityRepositoryImpl implementa SecurityOperations
type SecurityRepositoryImpl struct {
	DB *gorm.DB
}

// NewSecurityRepository crea una nueva instancia de SecurityRepositoryImpl
func NewSecurityRepository(db *gorm.DB) SecurityOperations {
	return &SecurityRepositoryImpl{DB: db}
}

// Implementación de los métodos

// Permisos
func (r *SecurityRepositoryImpl) CreatePermission(permission models.Permission) error {
	return r.DB.Create(&permission).Error
}

func (r *SecurityRepositoryImpl) GetAllPermissions() ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.DB.Find(&permissions).Error
	return permissions, err
}

// Roles
func (r *SecurityRepositoryImpl) CreateRole(role models.Role) error {
	return r.DB.Create(&role).Error
}

func (r *SecurityRepositoryImpl) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	err := r.DB.Preload("Permissions").Find(&roles).Error
	return roles, err
}

// Perfiles
func (r *SecurityRepositoryImpl) CreateProfile(profile models.Profile) error {
	return r.DB.Create(&profile).Error
}

func (r *SecurityRepositoryImpl) GetAllProfiles() ([]models.Profile, error) {
	var profiles []models.Profile
	err := r.DB.Preload("Roles").Preload("Permissions").Find(&profiles).Error
	return profiles, err
}

// Usuarios
func (r *SecurityRepositoryImpl) CreateUser(user models.User) error {
	return r.DB.Create(&user).Error
}

func (r *SecurityRepositoryImpl) GetUserByID(id int) (models.User, error) {
	var user models.User
	err := r.DB.Preload("Profile.Roles").Preload("Profile.Permissions").First(&user, id).Error
	return user, err
}

func (r *SecurityRepositoryImpl) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Preload("Profile.Roles").Preload("Profile.Permissions").Find(&users).Error
	return users, err
}
