package requests

type InsertWeightRequest struct {
	Date string `json:"date" validation:"required"`
	Max  int    `json:"max" validation:"required"`
	Min  int    `json:"min" validation:"required"`
}