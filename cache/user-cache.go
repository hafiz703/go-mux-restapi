package cache

import (
	"golang-mux-api/entity"
)

type UserCache interface {
	Set(key string, value *entity.User)
	Get(key string) *entity.User
	GetAll() ([]entity.User)
	Validate(user *entity.User) error
	Delete(key string)
}
