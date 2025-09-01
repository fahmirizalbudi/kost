package requests

type RentalRequest struct {
	RoomID         *int   `json:"room_id"`
	TenantID       *int   `json:"tenant_id"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	DurationMonths *int   `json:"duration_months"`
	Status         string `json:"status"`
}
