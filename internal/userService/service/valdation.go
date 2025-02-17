package service

import (
	"ExpenseManagement/internal/userService/contracts"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (s *Service) ValidateCreateUserRequest(u contracts.User) error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required, validation.Length(5, 14)),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 50)),
		validation.Field(&u.Name, validation.Required, validation.Length(5, 50)),
		validation.Field(&u.Age, validation.Min(10).Error("age must be 10 or older")),
		validation.Field(&u.Gender, validation.Required, validation.Length(1, 1)),
	)
}

func (s *Service) ValidateUserLoginRequest(u contracts.UserLogin) error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required, validation.Length(5, 14)),
		validation.Field(&u.Password, validation.Required, validation.Length(8, 50)),
	)
}
