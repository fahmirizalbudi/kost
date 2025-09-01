package responses

import "time"

type RentalResponse struct {
	ID             int       `json:"id"`
	RoomID         int       `json:"room_id"`
	TenantID       int       `json:"tenant_id"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	DurationMonths int       `json:"duration_months"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

type RentalWithRoomAndTenantResponse struct {
	ID             int          `json:"id"`
	RoomID         int          `json:"room_id"`
	TenantID       int          `json:"tenant_id"`
	StartDate      string       `json:"start_date"`
	EndDate        string       `json:"end_date"`
	DurationMonths int          `json:"duration_months"`
	Status         string       `json:"status"`
	Room           RoomResponse `json:"room"`
	Tenant         UserResponse `json:"tenant"`
	CreatedAt      time.Time    `json:"created_at"`
}