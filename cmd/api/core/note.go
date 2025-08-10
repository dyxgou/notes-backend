package core

type CreateNoteParams struct {
	GradeId   int64 `json:"grade_id" validate:"required"`
	StudentId int64 `json:"student_id" validate:"required"`
}

type GetNoteParams struct {
	GradeId   int64 `query:"grade_id"`
	StudentId int64 `query:"student_id"`
}

type GetAllStudentNotes struct {
	StudentId int64 `query:"student_id" validate:"required"`
	SubjectId int64 `query:"subject_id" validate:"required"`
}

type ChangeNoteValue struct {
	Id    int64 `json:"id" validate:"required"`
	Value byte  `json:"value" validate:"required,gte=10,lte=50"`
}
