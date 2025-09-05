package responses

import "time"

type TransactionResponse struct {
	ID             string    `json:"id"`
	RentalID       int       `json:"rental_id"`
	DormitoryPrice int       `json:"dormitory_price"`
	MonthPaid      int       `json:"month_paid"`
	Amount         int       `json:"amount"`
	Method         string    `json:"method"`
	Purpose        string    `json:"purpose"`
	Status         string    `json:"status"`
	Proof          *string   `json:"proof"`
	CreatedAt      time.Time `json:"created_at"`
}

type TransactionWithRentalResponse struct {
	ID             string                          `json:"id"`
	RentalID       int                             `json:"rental_id"`
	DormitoryPrice int                             `json:"dormitory_price"`
	MonthPaid      int                             `json:"month_paid"`
	Amount         int                             `json:"amount"`
	Method         string                          `json:"method"`
	Purpose        string                          `json:"purpose"`
	Status         string                          `json:"status"`
	Proof          *string                         `json:"proof"`
	Rental         RentalWithRoomAndTenantResponse `json:"rental"`
	CreatedAt      time.Time                       `json:"created_at"`
}
