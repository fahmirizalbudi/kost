package requests

type UserRequest struct {
	Name		string	`json:"name"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
	Role		string	`json:"role"`
	Phone		string	`json:"phone"`
	Address		string	`json:"address"`
}