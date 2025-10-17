package repositories

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Todoモデル
type Todo struct {
	gorm.Model
	Titie     string `gorm: type:string; not null`
	completed bool   `gorm: type:boolean; not null; default:false`
}

var db *gorm.DB

// DB初期化
func InitDB() {
	var err error
	db, err = gorm.Open(postgres.Open("demo_gin_next"), &gorm.Config{})
	if err != nil {
		panic("DB接続に失敗しました")
	}
	//テーブルがなければ自動作成
	db.AutoMigrate(&Todo{})
}
