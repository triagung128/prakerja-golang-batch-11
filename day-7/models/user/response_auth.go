package user

type ResponseAuth struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
