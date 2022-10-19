package system

import (
	"time"

	"github.com/gofrs/uuid"
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

type LoginRecord struct {
	ID        uint `gorm:"primarykey"`
	Uid       uint
	UserName  string
	LoginTime time.Time
	IP        string
}
