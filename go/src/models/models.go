package models

type People struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	YearBorn  string `json:"yearBorn"`
	Age       int    `json:"age"`
}
