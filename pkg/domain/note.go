package domain

type Note struct {
	Id        int64 `json:"id,omitempty"`
	GradeId   int64 `json:"grade_id,omitempty"`
	StudentId int64 `json:"student_id,omitempty"`
	Value     byte  `json:"value"`
}
