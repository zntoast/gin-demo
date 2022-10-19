package system

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID
	Name     string
	Phone    string
	Password string
	Addr     string
}

func (u User) TableName() string {
	return "users"
}

type LoginRecord struct {
	ID        uint `gorm:"primarykey"`
	Uid       uint
	UserName  string
	LoginTime time.Time
	IP        string
}
