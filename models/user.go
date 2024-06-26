package models

type User struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(300)" json:"name"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	Role string `gorm:"type:varchar(20)" json:"role"`
	Nomor string `gorm:"type:varchar(25)" json:"nomor"`
	Email string `gorm:"type:varchar(100)" json:"email"`
}