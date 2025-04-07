package handlers

import (
	"encoding/json"
	"net/http"
	"restaurant-management-system/internal/models"
	"restaurant-management-system/internal/repository/security"
	"restaurant-management-system/pkg/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// HandlerSecurity maneja las solicitudes relacionadas con seguridad
type HandlerSecurity struct {
	SecurityService security.SecurityOperations
}

// NewHandlerSecurity crea una nueva instancia de HandlerSecurity
func NewHandlerSecurity(service security.SecurityOperations) *HandlerSecurity {
	return &HandlerSecurity{SecurityService: service}
}

// Permisos
func (h *HandlerSecurity) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var permission models.Permission
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		http.Error(w, "Error al decodificar el permiso", http.StatusBadRequest)
		return
	}
	if permission.Name == "" {
		http.Error(w, "El nombre del permiso es obligatorio", http.StatusBadRequest)
		return
	}

	if err := h.SecurityService.CreatePermission(permission); err != nil {
		http.Error(w, "Error al crear el permiso", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(permission)
}

// GetAllPermissions handles the HTTP request to retrieve all permissions.
// It interacts with the SecurityService to fetch the list of permissions
// and returns them as a JSON response. If an error occurs during the process,
// it responds with an HTTP 500 status code and an error message.
//
// @param w http.ResponseWriter - The HTTP response writer to send the response.
// @param r *http.Request - The HTTP request received by the handler.
func (h *HandlerSecurity) GetAllPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.SecurityService.GetAllPermissions()
	if err != nil {
		http.Error(w, "Error al obtener los permisos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permissions)
}

// Roles
// CreateRole handles the HTTP request to create a new role.
// It expects a JSON payload in the request body containing the role details.
// The role must have a non-empty name and at least one permission.
// If the input is invalid, it responds with a 400 Bad Request status.
// If the role creation fails due to a server error, it responds with a 500 Internal Server Error status.
// On successful creation, it responds with a 201 Created status and returns the created role in the response body.
//
// Request Body:
//
//	{
//	  "name": "string",          // Name of the role (required)
//	  "permissions": ["string"]  // List of permissions (required)
//	}
//
// Responses:
// 400 Bad Request: If the input is invalid (e.g., missing name or permissions).
// 500 Internal Server Error: If there is an error while creating the role.
// 201 Created: If the role is successfully created, returns the created role in JSON format.
func (h *HandlerSecurity) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, "Error al decodificar el rol", http.StatusBadRequest)
		return
	}
	if role.Name == "" {
		http.Error(w, "El nombre del rol es obligatorio", http.StatusBadRequest)
		return
	}
	if len(role.Permissions) == 0 {
		http.Error(w, "Los permisos son obligatorios", http.StatusBadRequest)
		return
	}

	if err := h.SecurityService.CreateRole(role); err != nil {
		http.Error(w, "Error al crear el rol", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(role)
}

func (h *HandlerSecurity) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.SecurityService.GetAllRoles()
	if err != nil {
		http.Error(w, "Error al obtener los roles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}

// Perfiles
// CreateProfile handles the HTTP request for creating a new profile.
// It expects a JSON payload in the request body containing the profile details.
// The function performs the following steps:
//  1. Decodes the JSON payload into a models.Profile struct.
//  2. Validates that the profile has a non-empty Name and at least one Role.
//  3. Calls the SecurityService to create the profile in the system.
//  4. Returns an HTTP 201 Created status with the created profile in the response body
//     if the operation is successful.
//  5. Returns appropriate HTTP error responses for invalid input or internal errors.
//
// Parameters:
// - w: The HTTP response writer used to send the response.
// - r: The HTTP request containing the profile data.
//
// Possible HTTP Responses:
// - 400 Bad Request: If the JSON payload is invalid or required fields are missing.
// - 500 Internal Server Error: If an error occurs while creating the profile.
// - 201 Created: If the profile is successfully created.
func (h *HandlerSecurity) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Error al decodificar el perfil", http.StatusBadRequest)
		return
	}
	if profile.Name == "" {
		http.Error(w, "El nombre del perfil es obligatorio", http.StatusBadRequest)
		return
	}
	if len(profile.Roles) == 0 {
		http.Error(w, "Los roles son obligatorios", http.StatusBadRequest)
		return
	}

	if err := h.SecurityService.CreateProfile(profile); err != nil {
		http.Error(w, "Error al crear el perfil", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

func (h *HandlerSecurity) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := h.SecurityService.GetAllProfiles()
	if err != nil {
		http.Error(w, "Error al obtener los perfiles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

// Usuarios
// CreateUser handles the HTTP request for creating a new user.
// It expects a JSON payload in the request body containing the user's details.
// The function performs the following validations:
// - Ensures the username, email, and password fields are not empty.
// - Validates the email format using a utility function.
// - Ensures the profile ID is provided.
// If any validation fails, it responds with an appropriate HTTP error status and message.
// On successful validation, it calls the SecurityService to create the user.
// If the user is created successfully, it responds with HTTP status 201 (Created)
// and returns the created user in the response body as JSON.
// If an error occurs during user creation, it responds with HTTP status 500 (Internal Server Error).
func (h *HandlerSecurity) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error al decodificar el usuario", http.StatusBadRequest)
		return
	}
	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "El nombre, email y password son obligatorios", http.StatusBadRequest)
		return
	}
	if !utils.IsValidEmail(user.Email) {
		http.Error(w, "El email no es válido", http.StatusBadRequest)
		return
	}
	if user.ProfileID == 0 {
		http.Error(w, "El perfil es obligatorio", http.StatusBadRequest)
		return
	}
	if err := h.SecurityService.CreateUser(user); err != nil {
		http.Error(w, "Error al crear el usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserByID handles the HTTP request to retrieve a user by ID.
func (h *HandlerSecurity) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	user, err := h.SecurityService.GetUserByID(id)
	if err != nil {
		http.Error(w, "Error al obtener el usuario", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *HandlerSecurity) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.SecurityService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error al obtener los usuarios", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
