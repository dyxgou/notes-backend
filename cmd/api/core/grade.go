package core

type CreateGradeParams struct {
	Name      string `json:"name" validate:"required,min=2,max=20"`
	SubjectId int64  `json:"subject_id" validate:"required"`
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
