package student

import (
	"testing"

	"github.com/dyxgou/notas/pkg/domain"
)

func BenchmarkInsertingStudents(b *testing.B) {
	for b.Loop() {
		_, err := r.Insert(&domain.Student{
			Name:        "Bench",
			Course:      10,
			ParentPhone: "1231231231",
		})

		b.StopTimer()

		if err != nil {
			b.Errorf("benchmark err=%s", err.Error())
		}

		b.StartTimer()
	}
}
