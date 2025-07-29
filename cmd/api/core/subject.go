package core

type CreateSubjectParams struct {
	Name   string `json:"name" validate:"required,min=3,max=20"`
	Course byte   `json:"course" validate:"gte=0,lte=11"`
	Period byte   `json:"period" validate:"required,gte=1,lte=4"`
}

type GetByNameAndCourseQuery struct {
	Name   string `query:"name" validate:"required,max=15"`
	Course byte   `query:"course" validate:"required,gte=0,lte=11"`
	Period byte   `query:"period" validate:"required,gte=1,lte=4"`
}
