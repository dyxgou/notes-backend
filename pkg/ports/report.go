package ports

import "github.com/dyxgou/notas/pkg/apperrors"

type ReportService interface {
	GetSubjectReport(studentId int64, name string, course byte) ([]float64, *apperrors.Error)
	GetSubjectsAverage(studentId int64, course byte, names []string) (float64, *apperrors.Error)
}

type ReportRepository interface {
	GetSubjectReport(studentId int64, name string, course byte) ([]float64, error)
	GetSubjectAverage(studentId int64, name string, course byte) (float64, error)
}
