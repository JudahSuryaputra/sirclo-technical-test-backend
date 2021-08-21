package responses

import "sirclo-technical-test-backend/models/db"

type WeightIndexResponse struct {
	Weights []db.Weight `json:"weights"`
	Average Average		`json:"average"`
}

type Average struct {
	Max        float64 	`json:"max"`
	Min        float64 	`json:"min"`
	Difference float64 	`json:"difference"`
}
