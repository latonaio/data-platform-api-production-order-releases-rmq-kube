package requests

type Header struct {
	ProductionOrder            int     `json:"ProductionOrder"`
	IsReleased          	   *bool   `json:"IsReleased"`
}
