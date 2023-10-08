package db

import (
	uuid "github.com/satori/go.uuid"
)

type SCHEDULE struct {
	ID            uuid.UUID      `gorm:"primaryKey" json:"id"`
	Key           string         `gorm:"not null" json:"nyckel"`
}
