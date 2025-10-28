package report

import (
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/dyxgou/notas/pkg/domain"
	"github.com/dyxgou/notas/pkg/repositories/sqlite"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/grade"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/note"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/student"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/subject"
)

var r Repository
var subjectRepo subject.Repository
var gradeRepo grade.Repository
var studentRepo student.Repository
var noteRepo note.Repository

func TestMain(m *testing.M) {
	path := os.Getenv("DB_TEST_PATH")
	db := sqlite.ConnectClient(path)

	r.Db = db
	subjectRepo.Db = db
	gradeRepo.Db = db
	studentRepo.Db = db
	noteRepo.Db = db

	code := m.Run()

	db.Close()
	os.Exit(code)
}

func TestGetSubjectReport(t *testing.T) {
	tt := struct {
		studentName    string
		subjectName    string
		noteValue      byte
		subjectAverage byte
		course         byte
	}{
		studentName: "Diego",
		subjectName: "math",
		noteValue:   20,
		course:      0,
	}

	studentId, err := studentRepo.Insert(&domain.Student{
		Name:        tt.studentName,
		Course:      tt.course,
		ParentPhone: "1231231231",
	})

	if err != nil {
		t.Fatal(err)
	}

	for i := range byte(3) {
		id, err := subjectRepo.Insert(&domain.Subject{
			Name:   tt.subjectName,
			Course: tt.course,
			Period: i + 1,
		})

		if err != nil {
			t.Fatal(err)
		}

		for j := range byte(9) {
			gradeId, err := gradeRepo.Insert(&domain.Grade{
				SubjectId: id,
				Name:      fmt.Sprintf("Semana %d", j),
			})

			if err != nil {
				t.Fatal(err)
			}

			noteId, err := noteRepo.Insert(&domain.Note{
				GradeId:   gradeId,
				StudentId: studentId,
			})

			if err != nil {
				t.Fatal(err)
			}

			if err := noteRepo.ChangeValue(noteId, tt.noteValue); err != nil {
				t.Fatal(err)
			}
		}
	}

	averages, err := r.GetSubjectReport(studentId, tt.subjectName, tt.course)
	if err != nil {
		t.Fatal(err)
	}

	for _, avg := range averages {
		epsilon := 0.01
		if avg-float64(tt.noteValue)*0.7 > epsilon {
			t.Fatalf("subject average expected=%f. got=%f", float32(tt.noteValue), avg)
		}
	}
}

