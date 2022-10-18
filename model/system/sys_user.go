package system

import (
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
