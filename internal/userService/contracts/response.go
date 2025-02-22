package contracts

type UserResponse struct {
	UserId string `json:"user_id"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
