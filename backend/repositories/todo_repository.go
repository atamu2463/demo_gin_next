package repositories

import (
	"github.com/atamu2463/demo_gin_next/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DB初期化
func InitDB() {
	var err error
	dsn := "host=localhost user=postgres dbname=demo_gin_next password=demo_gin_next port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB接続に失敗しました")
	}
	//テーブルがなければ自動作成
	DB.AutoMigrate(&domain.Todo{})
}
