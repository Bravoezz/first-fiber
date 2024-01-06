package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	FirstName string
	LastName  string
	Email     string
	Tasks     []Task

	CreatedAt   time.Time
    UpdatedAt   time.Time
}
