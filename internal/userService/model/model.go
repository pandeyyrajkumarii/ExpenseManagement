package model

type UserDb struct {
	Id        string `gorm:"not null" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Age       int    `gorm:"not null" json:"age"`
	Gender    string `gorm:"not null" json:"gender"`
	CreatedAt int64  `gorm:"not null" json:"createdAt"`
	UpdatedAt int64  `gorm:"not null" json:"updatedAt"`
}

func (UserDb) TableName() string {
	return "users"
}

type PasswordDb struct {
	UserId    string `gorm:"column:userid;not null" json:"userid"`
	Password  string `gorm:"not null" json:"password"`
	Token     []byte `gorm:"not null" json:"token"`
	CreatedAt int64  `gorm:"not null" json:"createdAt"`
	UpdatedAt int64  `gorm:"not null" json:"updatedAt"`
}

func (PasswordDb) TableName() string {
	return "passwords"
}
