package core

type CreateSubjectParams struct {
	Name   string `json:"name" validate:"required,min=3,max=20"`
	Course byte   `json:"course" validate:"gte=0,lte=11"`
	Period byte   `json:"period" validate:"required,gte=1,lte=4"`
}

type GetByNameAndCourseQuery struct {
	Name   string `query:"name" validate:"required,max=15"`
	Course byte   `query:"course" validate:"gte=0,lte=11"`
	Period byte   `query:"period" validate:"gte=1,lte=4"`
}

type GetByPeriodAndCourseQuery struct {
	Course byte `query:"course" validate:"gte=0,lte=11"`
	Period byte `query:"period" validate:"gte=1,lte=4"`
}

type SubjectResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Course byte   `json:"course"`
	Period byte   `json:"period"`
	Grades byte   `json:"grades"`
}
