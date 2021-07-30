package Model

import (
	"time"
)

type Product struct {
	ID        uint
	Name      string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	Db.AutoMigrate(Product{})
}
