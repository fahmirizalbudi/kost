package responses

import "time"

type RoomResponse struct {
	ID			int			`json:"id"`
	DormitoryID int   		`json:"dormitory_id"`
	RoomNumber  string 		`json:"room_number"`
	Status  	string 		`json:"status"`
	CreatedAt   time.Time	`json:"created_at"`
	UpdatedAt   time.Time	`json:"updated_at"`
}

type RoomWithDormitoryResponse struct {
	ID			int					`json:"id"`
	DormitoryID int   				`json:"dormitory_id"`
	RoomNumber  string 				`json:"room_number"`
	Status  	string 				`json:"status"`
	Dormitory	DormitoryResponse	`json:"dormitory"`
	CreatedAt   time.Time			`json:"created_at"`
	UpdatedAt   time.Time			`json:"updated_at"`
}