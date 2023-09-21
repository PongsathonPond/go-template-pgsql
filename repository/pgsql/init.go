package pgsql

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"idev-cms-service/domain"
	"time"
)

type Repository struct {
	DB      *gorm.DB
	TBName  string
	Timeout time.Duration
}

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	TBName   string
	Timeout  int
	Timezone string
}

func New(_ context.Context, config *Config, tbName string) (repo *Repository, err error) {
	dsn := "host=localhost user=adminidev password=g9Npp7VTaCb4z0WjMs7TNOVUc3SC7ZD9xmJaLqCDef9jdy2X8y dbname=postgres  port=5433 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	db.AutoMigrate(&domain.Users{})

	if err != nil {
		return nil, err
	}

	repo = &Repository{
		DB: db,

		TBName: tbName,

		Timeout: time.Duration(config.Timeout) * time.Second,
	}
	return repo, nil
}
