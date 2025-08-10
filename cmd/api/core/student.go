package core

type CreateStudentParams struct {
	Name        string `json:"name" validate:"required,min=2,max=30"`
	Course      byte   `json:"course" validate:"gte=0,lte=11"`
	ParentPhone string `json:"parent_phone" validate:"required,len=10"`
}

type GetStudentAverageQuery struct {
	StudentId int64 `query:"student_id" validate:"required"`
	SubjectId int64 `query:"subject_id" validate:"required"`
}

type StudentResponse struct {
	Id          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Course      byte   `json:"course"`
	ParentPhone string `json:"parent_phone,omitempty"`
}

type ChangeStudentName struct {
	Id   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=2,max=30"`
}

type ChangeStudentParentPhone struct {
	Id          int64  `json:"id" validate:"required"`
	ParentPhone string `json:"parent_phone" validate:"required,min=2,max=30"`
}
