package responses

import "time"

type UserResponse struct {
	ID        int		`json:"id"`
	Name      string	`json:"name"`
	Email     string	`json:"email"`
	Role      string	`json:"role"`
	Phone     string	`json:"phone"`
	Address   string	`json:"address"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}