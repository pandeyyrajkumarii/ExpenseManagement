package model

type UserDb struct {
	Id        string `gorm:"not null" json:"id"`
	Name      string `gorm:"not null" json:"name"`
	Age       int    `gorm:"not null" json:"age"`
	Gender    string `gorm:"not null" json:"gender"`
	CreatedAt int64  `gorm:"not null" json:"created_at"`
	UpdatedAt int64  `gorm:"not null" json:"updated_at"`
}

func (UserDb) TableName() string {
	return "users"
}

type PasswordDb struct {
	UserId    string `gorm:"column:user_id;not null" json:"user_id"`
	Password  string `gorm:"not null" json:"password"`
	Token     []byte `gorm:"not null" json:"token"`
	CreatedAt int64  `gorm:"not null" json:"created_at"`
	UpdatedAt int64  `gorm:"not null" json:"updated_at"`
}

func (PasswordDb) TableName() string {
	return "passwords"
}
