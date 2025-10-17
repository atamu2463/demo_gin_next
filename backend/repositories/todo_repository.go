package repositories

import (
	"github.com/atamu2463/demo_gin_next/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// DB初期化
func InitDB() {
	var err error
	db, err = gorm.Open(postgres.Open("demo_gin_next"), &gorm.Config{})
	if err != nil {
		panic("DB接続に失敗しました")
	}
	//テーブルがなければ自動作成
	db.AutoMigrate(&domain.Todo{})
}
