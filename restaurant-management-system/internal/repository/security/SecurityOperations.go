package security

import "restaurant-management-system/internal/models"

// SecurityOperations define las operaciones relacionadas con permisos, roles, perfiles y usuarios
// SecurityOperations defines the interface for security-related operations
// within the system, including management of permissions, roles, profiles,
// and users.
//
// Methods:
//
//   - Permissions:
//     - CreatePermission: Adds a new permission to the system.
//     - GetAllPermissions: Retrieves all permissions available in the system.
//
//   - Roles:
//     - CreateRole: Adds a new role to the system.
//     - GetAllRoles: Retrieves all roles available in the system.
//
//   - Profiles:
//     - CreateProfile: Adds a new profile to the system.
//     - GetAllProfiles: Retrieves all profiles available in the system.
//
//   - Users:
//     - CreateUser: Adds a new user to the system.
//     - GetUserByID: Retrieves a user by their unique identifier.
//     - GetAllUsers: Retrieves all users available in the system.
type SecurityOperations interface {
	// Permisos
	CreatePermission(permission models.Permission) error
	GetAllPermissions() ([]models.Permission, error)

	// Roles
	CreateRole(role models.Role) error
	GetAllRoles() ([]models.Role, error)

	// Perfiles
	CreateProfile(profile models.Profile) error
	GetAllProfiles() ([]models.Profile, error)

	// Usuarios
	CreateUser(user models.User) error
	GetUserByID(id int) (models.User, error)
	GetAllUsers() ([]models.User, error)
}
