package core

type GetSubjectsReportParams struct {
	StudentId int64  `query:"student_id" validate:"required"`
	Name      string `query:"name" validate:"required,min=2,max=30"`
	Course    byte   `query:"course" validate:"gte=0,lte=11"`
}

type GetSubjectsAverageParams struct {
	StudentId int64    `query:"student_id" validate:"required"`
	Names     []string `query:"names" validate:"required"`
	Course    byte     `query:"course" validate:"gte=0,lte=11"`
}
