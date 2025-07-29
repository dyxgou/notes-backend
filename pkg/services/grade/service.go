package grade

import "github.com/dyxgou/notas/pkg/ports"

var _ ports.GradeService = &Service{}

type Service struct {
	Repo ports.GradeRespository
}
