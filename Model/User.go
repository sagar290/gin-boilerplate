package Model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint
	Name         string  `gorm:"size:255"`
	Email        *string `gorm:"size:255"`
	Age          int16   `gorm:"size:11"`
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func init() {
	Db.AutoMigrate(User{})
}
