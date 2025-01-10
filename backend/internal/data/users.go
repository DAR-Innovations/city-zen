package data

type User struct {
	BaseEntity
	FirstName  string `gorm:"size:255;not null"`
	LastName   string `gorm:"size:255;not null"`
	Phone      string `gorm:"size:20;unique;not null"`
	IsVerified bool   `gorm:"default:false"`
	Role       string `gorm:"size:50;not null;check:role IN ('ADMIN', 'USER')"`
	Password   string `gorm:"size:255;not null"`
}
