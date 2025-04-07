package models

// Tabla: permissions
type Permission struct {
	ID          int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"type:varchar(100);unique;not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null"`
}

// Tabla: roles
type Role struct {
	ID          int          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string       `json:"name" gorm:"type:varchar(100);unique;not null"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`
}

// Tabla intermedia: role_permissions
type RolePermission struct {
	RoleID       int `json:"role_id" gorm:"primaryKey"`
	PermissionID int `json:"permission_id" gorm:"primaryKey"`
}

// Tabla: profiles
type Profile struct {
	ID          int          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string       `json:"name" gorm:"type:varchar(100);unique;not null"`
	Roles       []Role       `json:"roles" gorm:"many2many:profile_roles;"`
	Permissions []Permission `json:"permissions" gorm:"many2many:profile_permissions;"`
}

// Tabla intermedia: profile_roles
type ProfileRole struct {
	ProfileID int `json:"profile_id" gorm:"primaryKey"`
	RoleID    int `json:"role_id" gorm:"primaryKey"`
}

// Tabla intermedia: profile_permissions
type ProfilePermission struct {
	ProfileID    int `json:"profile_id" gorm:"primaryKey"`
	PermissionID int `json:"permission_id" gorm:"primaryKey"`
}

// Tabla: users
type User struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string  `json:"username" gorm:"type:varchar(100);unique;not null"`
	Email     string  `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password  string  `json:"password" gorm:"type:varchar(255);not null"`
	ProfileID int     `json:"profile_id" gorm:"not null"`
	Profile   Profile `json:"profile" gorm:"foreignKey:ProfileID"`
}
