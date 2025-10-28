package grade

import (
	"log/slog"
	"os"
	"testing"

	"github.com/dyxgou/notas/pkg/domain"
	"github.com/dyxgou/notas/pkg/repositories/sqlite"
	"github.com/dyxgou/notas/pkg/repositories/sqlite/subject"
)

var r Repository
var subjectRepo subject.Repository

func TestMain(m *testing.M) {
	path := os.Getenv("DB_TEST_PATH")
	db := sqlite.ConnectClient(path)
	r.Db = db
	subjectRepo.Db = db

	code := m.Run()

	if err := db.Close(); err != nil {
		slog.Error("testing closing database", "err", err)
	}

	os.Exit(code)
}

func TestInsertGrade(t *testing.T) {
	tt := struct {
		subject *domain.Subject
		grade   *domain.Grade
	}{
		subject: &domain.Subject{
			Name:   "Math",
			Course: 4,
			Period: 4,
			Grades: 0,
		},
		grade: &domain.Grade{
			Name: "first exam",
		},
	}

	id, err := subjectRepo.Insert(tt.subject)
	if err != nil {
		t.Fatal(err)
	}

	tt.grade.SubjectId = id

	id, err = r.Insert(tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	row := r.Db.QueryRow("SELECT * FROM grade WHERE id = ?", id)

	var g domain.Grade
	if err := row.Scan(&g.Id, &g.Name, &g.SubjectId, &g.IsFinalExam); err != nil {
		t.Fatal(err)
	}

	if g.Name != tt.grade.Name {
		t.Fatalf("grade name expected=%q. got=%q", tt.grade.Name, g.Name)
	}

	if g.SubjectId != tt.grade.SubjectId {
		t.Fatalf("grade subjectId expected=%q. got=%q", tt.grade.SubjectId, tt.grade.SubjectId)
	}

	row = r.Db.QueryRow("SELECT grades FROM subject WHERE id = ?", tt.grade.SubjectId)

	var grades byte
	if err := row.Scan(&grades); err != nil {
		t.Fatal(err)
	}

	if grades != tt.subject.Grades+1 {
		t.Fatalf(
			"subject grades did not increased. expected=%d. got=%d",
			tt.subject.Grades+1,
			grades,
		)
	}
}

func TestGetGrade(t *testing.T) {
	tt := struct {
		subject *domain.Subject
		grade   *domain.Grade
	}{
		subject: &domain.Subject{
			Name:   "Math",
			Course: 7,
			Grades: 1,
			Period: 1,
		},
		grade: &domain.Grade{
			Name: "first exam",
		},
	}

	id, err := subjectRepo.Insert(tt.subject)
	if err != nil {
		t.Fatal(err)
	}

	tt.grade.SubjectId = id

	id, err = r.Insert(tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	g, err := r.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if g.Name != tt.grade.Name {
		t.Fatalf("grade name expected=%q. got=%q", tt.grade.Name, g.Name)
	}

	if g.SubjectId != tt.grade.SubjectId {
		t.Fatalf("grade subjectId expected=%q. got=%q", tt.grade.SubjectId, g.SubjectId)
	}
}

func TestInsertGradeWhenSubjectFull(t *testing.T) {
	tt := struct {
		subject *domain.Subject
		grade   *domain.Grade
	}{
		subject: &domain.Subject{
			Name:   "Math",
			Course: 0,
			Grades: 10,
			Period: 1,
		},
		grade: &domain.Grade{
			Name: "first exam",
		},
	}

	id, err := subjectRepo.Insert(tt.subject)
	if err != nil {
		t.Fatal(err)
	}

	q := "UPDATE subject SET grades = 10 WHERE id = ?;"
	if _, err := r.Db.Exec(q, id); err != nil {
		t.Fatal(err)
	}

	tt.grade.SubjectId = id

	id, err = r.Insert(tt.grade)
	if err == nil {
		t.Fatalf("subject inserted when grades were full. grades=%d", tt.subject.Grades)
	}
}

func TestInsertFinalExamGrade(t *testing.T) {
	tt := struct {
		subject *domain.Subject
		grade   *domain.Grade
	}{
		subject: &domain.Subject{
			Name:   "Math",
			Course: 2,
			Grades: 10,
			Period: 1,
		},
		grade: &domain.Grade{
			Name: "first exam",
		},
	}

	id, err := subjectRepo.Insert(tt.subject)
	if err != nil {
		t.Fatal(err)
	}

	tt.grade.SubjectId = id

	id, err = r.Insert(tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	tt.grade.Id = id
}

func TestChangeName(t *testing.T) {
	tt := struct {
		subject      *domain.Subject
		grade        *domain.Grade
		expectedName string
	}{
		subject: &domain.Subject{
			Name:   "Math",
			Course: 9,
			Grades: 1,
			Period: 1,
		},
		grade: &domain.Grade{
			Name: "first exam",
		},

		expectedName: "Second exam",
	}

	id, err := subjectRepo.Insert(tt.subject)
	if err != nil {
		t.Fatal(err)
	}

	tt.grade.SubjectId = id

	id, err = r.Insert(tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	if err := r.ChangeName(id, tt.expectedName); err != nil {
		t.Fatal(err)
	}

	row := r.Db.QueryRow("SELECT name FROM grade WHERE id = ?", id)

	var name string
	if err := row.Scan(&name); err != nil {
		t.Fatal(err)
	}

	if name != tt.expectedName {
		t.Fatalf("name has not changed. expected=%q. got=%q", tt.expectedName, name)
	}
}

func TestDeleteGrade(t *testing.T) {
	tt := struct {
		subject *domain.Subject
		grade   *domain.Grade
	}{
		subject: &domain.Subject{
			Name:   "spanish",
			Course: 7,
			Period: 4,
			Grades: 1,
		},
		grade: &domain.Grade{
			Name: "Semana 1",
		},
	}

	id, err := subjectRepo.Insert(tt.subject)
	if err != nil {
		t.Fatal(err)
	}

	tt.grade.SubjectId = id

	id, err = r.Insert(tt.grade)
	if err != nil {
		t.Fatal(err)
	}

	gradeId, err := r.Delete(id, tt.grade.SubjectId)
	if err != nil {
		t.Fatal(err)
	}

	if gradeId == 0 {
		t.Fatalf("grade does not exists. id=%d", id)
	}

	s, err := subjectRepo.GetSubjectById(tt.grade.SubjectId)
	if err != nil {
		t.Fatal(err)
	}

	if s.Grades != 0 {
		t.Fatalf("subject grades amount expected=%d, got=%d", 0, s.Grades)
	}
}
