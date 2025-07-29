package note

import (
	"os"
	"testing"

	"github.com/dyxgou/notas/pkg/domain"
	"github.com/dyxgou/notas/pkg/repositories/sqlite"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/grade"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/student"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/subject"
)

var r Repository
var subjectRepo subject.Repository
var gradeRepo grade.Repository
var studentRepo student.Repository

func TestMain(m *testing.M) {
	path := os.Getenv("DB_TEST_PATH")
	db := sqlite.ConnectClient(path)

	r.Db = db
	subjectRepo.Db = db
	gradeRepo.Db = db
	studentRepo.Db = db

	code := m.Run()

	db.Close()
	os.Exit(code)
}

func createNote(
	student *domain.Student,
	subject *domain.Subject,
	grade *domain.Grade,
) (*domain.Note, error) {
	subjectId, err := subjectRepo.Insert(subject)
	if err != nil {
		return nil, err
	}

	grade.SubjectId = subjectId
	gradeId, err := gradeRepo.Insert(grade)
	if err != nil {
		return nil, err
	}

	studentId, err := studentRepo.Insert(student)
	if err != nil {
		return nil, err
	}

	note := &domain.Note{
		GradeId:   gradeId,
		StudentId: studentId,
	}

	noteId, err := r.Insert(note)
	if err != nil {
		return nil, err
	}

	note.Id = noteId
	const defaultValue = 10
	note.Value = defaultValue

	return note, nil
}

func TestInsertNote(t *testing.T) {
	tt := struct {
		student      *domain.Student
		subject      *domain.Subject
		grade        *domain.Grade
		expectedNote *domain.Note
	}{
		student: &domain.Student{
			Name:        "AlejandroTest",
			Course:      5,
			ParentPhone: "1231231231",
		},
		subject: &domain.Subject{
			Name:   "Math",
			Course: 5,
			Period: 2,
			Grades: 1,
		},
		grade: &domain.Grade{
			Name: "First Exam",
		},
	}

	_, err := createNote(tt.student, tt.subject, tt.grade)
	if err != nil {
		t.Fatal(err)
	}

}

func TestGetByGradeIdAndStudentIdNote(t *testing.T) {
	tt := struct {
		student *domain.Student
		subject *domain.Subject
		grade   *domain.Grade
	}{
		student: &domain.Student{
			Name:        "AlejandroTest",
			Course:      5,
			ParentPhone: "1231231231",
		},
		subject: &domain.Subject{
			Name:   "Spanish",
			Period: 4,
			Course: 9,
			Grades: 1,
		},
		grade: &domain.Grade{
			Name: "First Exam",
		},
	}

	n1, err := createNote(tt.student, tt.subject, tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	n2, err := r.GetByGradeIdAndStudentId(n1.GradeId, n1.StudentId)
	if err != nil {
		t.Fatal(err)
	}

	if n1.Value != n2.Value {
		t.Fatalf("note value expected=%d. got=%d", n1.Value, n2.Value)
	}

	if n1.GradeId != n2.GradeId {
		t.Fatalf("note gradeId expected=%d. got=%d", n1.GradeId, n2.GradeId)
	}

	if n1.StudentId != n2.StudentId {
		t.Fatalf("note studentId expected=%d. got=%d", n1.StudentId, n2.StudentId)
	}
}

func TestGetNote(t *testing.T) {
	tt := struct {
		student *domain.Student
		subject *domain.Subject
		grade   *domain.Grade
	}{
		student: &domain.Student{
			Name:        "AlejandroTest2",
			Course:      9,
			ParentPhone: "1231231231",
		},
		subject: &domain.Subject{
			Name:   "Math",
			Period: 4,
			Course: 9,
			Grades: 1,
		},
		grade: &domain.Grade{
			Name: "First Exam",
		},
	}

	n1, err := createNote(tt.student, tt.subject, tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	n2, err := r.Get(n1.Id)
	if err != nil {
		t.Fatal(err)
	}

	if n1.Value != n2.Value {
		t.Fatalf("note value expected=%d. got=%d", n1.Value, n2.Value)
	}

	if n1.GradeId != n2.GradeId {
		t.Fatalf("note gradeId expected=%d. got=%d", n1.GradeId, n2.GradeId)
	}

	if n1.StudentId != n2.StudentId {
		t.Fatalf("note studentId expected=%d. got=%d", n1.StudentId, n2.StudentId)
	}

}

func TestChangeValue(t *testing.T) {
	tt := struct {
		student      *domain.Student
		subject      *domain.Subject
		grade        *domain.Grade
		expectedNote *domain.Note
	}{
		student: &domain.Student{
			Name:        "AlejandroTest",
			Course:      5,
			ParentPhone: "1231231231",
		},
		subject: &domain.Subject{
			Name:   "Math",
			Period: 3,
			Course: 10,
			Grades: 1,
		},
		grade: &domain.Grade{
			Name: "First Exam",
		},
	}

	note, err := createNote(tt.student, tt.subject, tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	if err := r.ChangeValue(note.Id, 50); err != nil {
		t.Fatal(err)
	}
}
