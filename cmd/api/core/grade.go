package core

type CreateGradeParams struct {
	Name        string `json:"name" validate:"required,min=2,max=20"`
	SubjectId   int64  `json:"subject_id" validate:"required"`
	IsFinalExam bool   `json:"is_final_exam"`
}

type GetGradeByNameAndSubjectIdQuery struct {
	Name      string `query:"name"`
	SubjectId int64  `query:"subject"`
}

type ChangeGradeName struct {
	Id   int64  `json:"id"`
	Name string `json:"name" validate:"min=3,max=20"`
}

type GetGradeBySubjectId struct {
	SubjectId int64 `query:"subject_id" validate:"required"`
}

type DeleteGradeParams struct {
	GradeId   int64 `query:"grade_id" validate:"required"`
	SubjectId int64 `query:"subject_id" validate:"required"`
}
