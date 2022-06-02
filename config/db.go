package config

import (
	"database/sql"
	"fmt"
    
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func init() {
	LoadDbConfig()
}

var dbConfig DbConfig

func LoadDbConfig() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	dbConfig.Host = viper.GetString("MYSQL_HOST")
	dbConfig.Port = viper.GetString("MYSQL_PORT")
	dbConfig.Username = viper.GetString("MYSQL_USER_NAME")
	dbConfig.Password = viper.GetString("MYSQL_PASSWORD")
	dbConfig.Database = viper.GetString("MYSQL_DATABASE")
}
func GetDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
