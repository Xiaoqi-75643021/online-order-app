package database_test

import (
	"fmt"
	"online-ordering-app/internal/config"
	"online-ordering-app/internal/model"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestConnectToMySQL(t *testing.T) {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Host, config.Cfg.Database.Port, config.Cfg.Database.Name)

	fmt.Println("dsn:" + dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*model.User, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range data {
		fmt.Printf("User ==> [%v] \n", v)
	}
}
