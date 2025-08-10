package domain

type Grade struct {
	Id          int64  `json:"id"`
	SubjectId   int64  `json:"subject_id,omitempty"`
	Name        string `json:"name"`
	IsFinalExam bool   `json:"is_final_exam"`
}
