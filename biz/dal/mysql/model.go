package mysql

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserId         string         `gorm:"primaryKey;size:50;column:user_id"`
	Username       string         `gorm:"size:50;not null;column:username"`
	Password       string         `gorm:"size:100;not null;column:password"`
	Phone          string         `gorm:"size:20;not null;uniqueIndex;column:phone"`
	BudgetMin      float64        `gorm:"type:decimal(10,2);default:0.00;column:budget_min"`
	BudgetMax      float64        `gorm:"type:decimal(10,2);default:0.00;column:budget_max"`
	PreferredType  string         `gorm:"size:20;default:'';column:preferred_type"`
	PreferredBrand string         `gorm:"size:50;default:'';column:preferred_brand"`
	Status         int8           `gorm:"default:1;column:status"`
	Address        string         `gorm:"size:255;column:address"`
	CreatedAt      time.Time      `gorm:"column:created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
