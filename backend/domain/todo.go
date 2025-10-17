package domain

// Todoモデル
type Todo struct {
	id        uint   `gorm: primaryKey; autoIncrement`
	titie     string `gorm: type:string; not null`
	completed bool   `gorm: type:boolean; not null; default:false`
}
