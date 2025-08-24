package requests

type UserRequest struct {
	ID 			int		`json:"id"`
	Name		string	`json:"name"`
	Email		string	`json:"email"`
	Password	string	`json:"password"`
	Role		string	`json:"role"`
	Phone		string	`json:"phone"`
	Address		string	`json:"address"`
}