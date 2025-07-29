package domain

type Student struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Course      byte   `json:"course,omitempty"`
	ParentPhone string `json:"parent_phone,omitempty"`
}
