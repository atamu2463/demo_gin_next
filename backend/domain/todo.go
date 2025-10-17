package domain

// Todoモデル
type Todo struct {
	ID        uint   `gorm: primaryKey; autoIncrement`
	Title     string `gorm: type:string; not null`
	Completed bool   `gorm: type:boolean; not null; default:false`
}
