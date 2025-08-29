package responses

import "time"

type DormitoryResponse struct {
	ID          int    		`json:"id"`
	Name        string 		`json:"name"`
	Address     string 		`json:"address"`
	Description string 		`json:"description"`
	Price       int    		`json:"price"`
	Facilities  string 		`json:"facilities"`
	GoogleMaps  string 		`json:"google_maps"`
	CreatedAt   time.Time	`json:"created_at"`
	UpdatedAt   time.Time	`json:"updated_at"`
}