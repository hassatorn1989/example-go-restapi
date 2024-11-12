package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:varchar(100);primaryKey"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil { // Check if ID is not already set
		u.ID = uuid.New() // Generate a new UUID
	}
	u.CreatedBy = "1"
	u.UpdatedBy = "1"
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedBy = "1"
	u.UpdatedAt = time.Now()
	return nil
}
