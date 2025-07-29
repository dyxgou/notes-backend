package domain

type Note struct {
	Id        int64 `json:"id,omitempty"`
	GradeId   int64 `json:"grade_id"`
	StudentId int64 `json:"student_id"`
	Value     byte  `json:"value"`
}
