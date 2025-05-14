package config

import (
	"github.com/rochaandrey/goproject.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgresConfig")

	// String de conexão para PostgreSQL
	dbpath := "host=localhost user=postgres password=postgres dbname=mydatabase port=5432 sslmode=disable"

	// Conectar primeiro ao banco postgres padrão para criar nosso banco de dados se necessário
	dbRoot, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		logger.Errorf("falha ao conectar ao PostgreSQL: %v", err)
		return nil, err
	}

	//Se o banco não existir, cria ele
	var exists bool
	dbRoot.Raw("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = ?)", "mydatabase").Scan(&exists)
	if !exists {
		if err := dbRoot.Exec("CREATE DATABASE mydatabase WITH ENCODING 'UTF8'").Error; err != nil {
			logger.Errorf("erro ao criar database: %v", err)
			return nil, err
		}
	}

	// Fechar conexão com banco postgres padrão
	sqlDB, _ := dbRoot.DB()
	sqlDB.Close()

	// Cria e conecta ao banco de dados mydatabase
	db, err := gorm.Open(postgres.Open(dbpath), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgresql opening error: %v", err)
		return nil, err
	}

	// Migrar schemas
	logger.Info("Executando migrações...")
	if err = db.AutoMigrate(&schemas.Opening{}); err != nil {
		return nil, err
	}

	logger.Info("Banco de dados PostgreSQL inicializado com sucesso")
	return db, nil
}
