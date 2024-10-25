package auth

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Handler struct {
    service *Service
}

func NewHandler(service *Service) *Handler {
    return &Handler{service: service}
}

func RegisterRoutes(r *gin.Engine, service *Service) {
    handler := NewHandler(service)

    auth := r.Group("/auth")
    {
        auth.POST("/login", handler.Login)
        auth.POST("/register", handler.Register)
    }
}

func (h *Handler) Login(c *gin.Context) {
    // Implementation
}

func (h *Handler) Register(c *gin.Context) {
    // Implementation
}
