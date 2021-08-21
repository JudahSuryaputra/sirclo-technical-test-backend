package requests

type UpdateWeightRequest struct {
	Max        int    `json:"max"`
	Min        int    `json:"min"`
	Difference int    `json:"difference,omitempty"`
	Date       string `json:"date"`
}