package initdb

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres(account, password, address, schema string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", account, password, address, schema)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
