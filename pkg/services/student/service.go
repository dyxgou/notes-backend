package student

import "github.com/dyxgou/notas/pkg/ports"

var _ ports.StudentService = &Service{}

type Service struct {
	Repo ports.StudentRepository
}
