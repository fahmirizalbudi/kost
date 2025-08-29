package requests

type DormitoryRequest struct {
	Name		string	`json:"name"`
	Address		string	`json:"address"`
	Description	string	`json:"description"`
	Price		*int	`json:"price"`
	Facilities	string	`json:"facilities"`
	GoogleMaps	string	`json:"google_maps"`
}