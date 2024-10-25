package auth

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go-starter/pkg/hash"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Register(email, password string) error {
	hashedPassword, err := hash.Password(password)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	_, err = s.db.Exec(query, email, hashedPassword)
	return err
}

func (s *Service) Login(email, password string) (string, error) {
	var user User
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := s.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return "", err
	}

	if !hash.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
