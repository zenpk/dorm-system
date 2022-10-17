package dal

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
	// GORM connect to DB
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         200,   // default size for string fields
		SkipInitializeWithVersion: false, // autoconfigure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	// CREATE TABLE
	if err = DB.AutoMigrate(&UserCredential{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&UserInfo{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&Building{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&Dorm{}); err != nil {
		return err
	}
	return nil
}
