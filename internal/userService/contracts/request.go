package contracts

type User struct {
	Id       string
	Password string
	Name     string
	Age      int
	Gender   string
}

type UserLogin struct {
	Id       string
	Password string
}
