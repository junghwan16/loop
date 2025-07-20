package user

import "github.com/google/uuid"

type UserRepository interface {
	Save(user *User) error
	FindByID(id uuid.UUID) (*User, error)
	FindByEmail(email string) (*User, error) // 가입 시 중복 체크용
}
