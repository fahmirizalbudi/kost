package requests

type RoomRequest struct {
	DormitoryID *int   		`json:"dormitory_id"`
	RoomNumber  string 		`json:"room_number"`
	Status	  	string 		`json:"status"`
}