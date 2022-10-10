package config

import (
	"assigment-2/server/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysqlGORM() (*gorm.DB, error) {
	dsn := "mysql://root:59Ll41dBZ2omHpm44vZq@containers-us-west-99.railway.app:7446/railway"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(models.Item{}, models.Order{})

	return db, nil
}
