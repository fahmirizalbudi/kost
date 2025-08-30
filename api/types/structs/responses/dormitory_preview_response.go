package responses

import "time"

type DormitoryPreviewResponse struct {
	ID          int       `json:"id"`
	DormitoryID int       `json:"formitory_id"`
	Url         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
}
