package config

import (
	"fmt"

	"gorm.io/gorm"
)

// por algum motivo, se eu der CNTRL+S no main o erro de logger desaparece

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	// inicializando postgresConfig
	db, err = InitializePostgres()

	if err != nil {
		return fmt.Errorf("error initialize postgresql: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
