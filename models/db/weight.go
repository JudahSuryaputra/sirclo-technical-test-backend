package db

type Weight struct {
	ID         int       	`db:"id" json:"id"`
	Date       string		`db:"date" json:"date"`
	Max        int       	`db:"max" json:"max"`
	Min        int       	`db:"min" json:"min"`
	Difference int       	`db:"difference" json:"difference"`
}