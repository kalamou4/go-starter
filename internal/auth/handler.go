package auth

import (
	"encoding/json"
	"github.com/gorilla/mux" // Optional, for easier routing
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// RegisterRoutes registers auth routes using standard http
func RegisterRoutes(router *mux.Router, service *Service) {
	handler := NewHandler(service)

	// Create subrouter for /auth endpoints
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", handler.Login).Methods("POST")
	authRouter.HandleFunc("/register", handler.Register).Methods("POST")
}

// If you don't want to use gorilla/mux, here's how to use standard http.ServeMux:
func RegisterRoutesStd(mux *http.ServeMux, service *Service) {
	handler := NewHandler(service)

	mux.HandleFunc("/auth/login", handler.Login)
	mux.HandleFunc("/auth/register", handler.Register)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request body
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate request
	if err := validateLoginRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process login
	resp, err := h.service.Login(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request body
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate request
	if err := validateRegisterRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process registration
	resp, err := h.service.Register(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

// Helper functions for validation
func validateLoginRequest(req LoginRequest) error {
	if req.Email == "" {
		return &ValidationError{Field: "email", Message: "Email is required"}
	}
	if req.Password == "" {
		return &ValidationError{Field: "password", Message: "Password is required"}
	}
	return nil
}

func validateRegisterRequest(req RegisterRequest) error {
	if req.Email == "" {
		return &ValidationError{Field: "email", Message: "Email is required"}
	}
	if req.Password == "" {
		return &ValidationError{Field: "password", Message: "Password is required"}
	}
	if len(req.Password) < 8 {
		return &ValidationError{Field: "password", Message: "Password must be at least 8 characters"}
	}
	if req.FirstName == "" {
		return &ValidationError{Field: "firstName", Message: "First name is required"}
	}
	if req.LastName == "" {
		return &ValidationError{Field: "lastName", Message: "Last name is required"}
	}
	return nil
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

// ResponseWriter helper functions
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
