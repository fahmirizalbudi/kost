package requests

type TransactionRequest struct {
	RentalID  *int   `json:"rental_id"`
	MonthPaid *int   `json:"month_paid"`
	Method    string `json:"method"`
	Purpose   string `json:"purpose"`
	Status    string `json:"status"`
	Proof     string `json:"proof"`
}
