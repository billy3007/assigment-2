package config

import (
	"assigment-2/server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysqlGORM() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(models.Item{}, models.Order{})

	return db, nil
}
