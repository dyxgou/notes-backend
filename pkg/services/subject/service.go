package subject

import "github.com/dyxgou/notas/pkg/ports"

var _ ports.SubjectService = &Service{}

type Service struct {
	Repo ports.SubjectRepository
}
