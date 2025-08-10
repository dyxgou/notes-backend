package domain

type Subject struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Course       byte   `json:"course,omitempty"`
	HasFinalExam bool   `json:"has_final_exam"`
	Period       byte   `json:"period,omitempty"`
	Grades       byte   `json:"grades,omitempty"`
}
