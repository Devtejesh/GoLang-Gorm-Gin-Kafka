package models

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm :"not null" josn:"name"`
	Email string `gorm : "unique; not null" json:"email"`
	Age   int    `gorm : "not null" json:"age"`
}
