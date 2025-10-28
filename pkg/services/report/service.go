package report

import "github.com/dyxgou/notas/pkg/ports"

var _ ports.ReportService = &Service{}

type Service struct {
	Repo ports.ReportRepository
}
