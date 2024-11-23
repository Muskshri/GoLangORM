package databases

import (
	"fmt"
	"gorm_demo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConfig holds db configrations
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Database connection instance
var dbConfig = DBConfig{
	Host:     "127.0.0.1",
	Port:     "3306",
	User:     "root",
	Password: "abc123",
	DBName:   "store",
}


func InitDatabase() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	//to open db

	db, err := gorm.Open(mysql.Open(connectionString))
	if err != nil{
		return nil, fmt.Errorf("failed to connect to Mysql database : %v", err)
	}


    //to automate the creation of tables.
	err = db.AutoMigrate(&models.User{})

	if err != nil {
		return nil, fmt.Errorf("failed to migrate databse: %v", err)
	}

	return db, nil
}