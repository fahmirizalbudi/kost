package requests

type DormitoryPreviewRequest struct {
	DomitoryID	int		`json:"dormitory_id"`
	Url			string	`json:"url"`
}