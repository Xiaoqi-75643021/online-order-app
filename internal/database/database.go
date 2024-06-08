package database

import (
	"fmt"
	"log"
	"online-ordering-app/internal/config"
	"online-ordering-app/internal/model"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   = Init()
	once sync.Once
)

func Init() *gorm.DB {
	db, err := ConnectDatabase(config.Cfg)
	if err != nil {
		log.Fatalf("Failed to load database: %v", err)
	}

	// 自动迁移模型
	db.AutoMigrate(&model.User{}, &model.Dish{}, &model.Cart{}, &model.CartItem{}, &model.Category{}, &model.Coupon{}, &model.UserCoupon{}, &model.Order{}, &model.OrderItem{})
	return db
}

// ConnectDatabase 初始化连接和返回数据库的单例
func ConnectDatabase(cfg *config.AppConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})
	return db, err
}
