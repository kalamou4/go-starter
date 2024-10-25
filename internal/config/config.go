package config

import (
	"bufio"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"os"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Config struct {
	ServerAddr   string
	DatabaseURL  string
	MigrationURL string
	JWTSecret    string
}

func Load() (*Config, error) {
	if err := SetEnvFile(".env"); err != nil {
		return nil, err
	}
	return &Config{
		ServerAddr:   os.Getenv("SERVER_ADDR"),
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		MigrationURL: os.Getenv("MIGRATION_URL"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
	}, nil
}

func SetEnvFile(filePath string) error {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open env file: %w", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Keep track of the line number for error reporting
	lineNum := 0

	// Scan the file line by line
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("malformed env variable at line %d: %s", lineNum, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove surrounding quotes if present
		value = strings.Trim(value, `"'`)

		// Set the environment variable
		err := os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("failed to set env variable %s: %w", key, err)
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading env file: %w", err)
	}

	return nil
}

func RunMigrations(migration_url, name_db string) error {
	m, err := migrate.New(
		"file://"+migration_url,
		"sqlite3://"+name_db,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
