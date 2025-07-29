package domain

type Subject struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Course byte   `json:"course"`
	Period byte   `json:"period"`
	Grades byte   `json:"grades"`
}
