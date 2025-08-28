package structs

type Payload struct {
	Message string      `json:"message"`
	Error   any `json:"error"`
	Data    any `json:"data"`
}