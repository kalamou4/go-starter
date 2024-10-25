package auth

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
    Email     string `json:"email" binding:"required,email"`
    Password  string `json:"password" binding:"required,min=8"`
    FirstName string `json:"first_name" binding:"required"`
    LastName  string `json:"last_name" binding:"required"`
}

type AuthResponse struct {
    Token string `json:"token"`
    User  interface{} `json:"user"`
}
