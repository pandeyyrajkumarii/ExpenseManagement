package repo

import (
	"ExpenseManagement/internal/userService/contracts"
	"ExpenseManagement/internal/userService/model"
	"ExpenseManagement/packages/database"
	"fmt"
	"time"
)

type UserRepo struct {
	Dbs *database.DatabaseGorm
}

func NewUserRepo(db *database.DatabaseGorm) *UserRepo {
	return &UserRepo{Dbs: db}
}

func (u *UserRepo) Create(usr *contracts.User) (*model.UserDb, error) {

	currentTime := time.Now().Unix()
	userDb := ToUserDBModel(usr)
	userDb.CreatedAt = currentTime
	userDb.UpdatedAt = currentTime
	_, err := u.Dbs.Create(&userDb)
	if err != nil {
		return nil, err
	}
	passwordDb := ToPasswordModel(usr)
	passwordDb.CreatedAt = currentTime
	passwordDb.UpdatedAt = currentTime
	_, err = u.Dbs.Create(&passwordDb)
	if err != nil {
		return nil, err
	}

	return userDb, nil

}

func (u *UserRepo) FindByUserId(userId string) (*model.UserDb, error) {
	user := model.UserDb{}
	count, err := u.Dbs.FindByID(&user, userId)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("user_not_found")
	}
	return &user, nil
}

func (u *UserRepo) FindUserByIdForLogin(userId string) (*model.PasswordDb, error) {
	user := model.PasswordDb{}
	count, err := u.Dbs.FindWithQuery(&user, map[string]interface{}{"userid": userId})
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("user_not_found")
	}
	return &user, nil
}

func ToPasswordModel(usr *contracts.User) *model.PasswordDb {
	return &model.PasswordDb{
		UserId:   usr.Id,
		Password: usr.Password,
		Token:    []byte("dummy"),
	}
}

func ToUserDBModel(usr *contracts.User) *model.UserDb {
	return &model.UserDb{
		Id:     usr.Id,
		Name:   usr.Name,
		Age:    usr.Age,
		Gender: usr.Gender,
	}
}

func ToUserApiModel(usr *model.UserDb) *contracts.User {
	return &contracts.User{
		Id:     usr.Id,
		Name:   usr.Name,
		Age:    usr.Age,
		Gender: usr.Gender,
	}
}
