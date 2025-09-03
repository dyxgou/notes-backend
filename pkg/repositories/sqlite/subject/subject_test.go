package subject

import (
	"os"
	"testing"

	"github.com/dyxgou/notas/pkg/domain"
	"github.com/dyxgou/notas/pkg/repositories/sqlite"
)

var r Repository

func TestMain(m *testing.M) {
	path := os.Getenv("DB_TEST_PATH")
	db := sqlite.ConnectClient(path)
	r.Db = db

	code := m.Run()

	db.Close()
	os.Exit(code)
}

func subjectExists(t *testing.T, id int64) bool {
	query := "SELECT EXISTS(SELECT 1 FROM subject WHERE id = ?)"

	row := r.Db.QueryRow(query, id)
	var exists bool

	err := row.Scan(&exists)
	if err != nil {
		t.Fatal(err)
	}

	return exists
}

func TestInsertSubject(t *testing.T) {
	tt := &domain.Subject{
		Name:   "Math",
		Course: 11,
		Period: 3,
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	if !subjectExists(t, id) {
		t.Fatalf("subject was not created. name=%q", tt.Name)
	}
}

func TestGetByNameAndCourse(t *testing.T) {
	tt := &domain.Subject{
		Name:   "Social",
		Course: 2,
		Period: 2,
		Grades: 1,
	}

	_, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	id, err := r.GetByNameAndCourse(tt.Name, tt.Course, tt.Period)
	if err != nil {
		t.Fatal(err)
	}

	if id == 0 {
		t.Fatalf("subject id invalid. got=%d", id)
	}
}

func TestGetByCourseAndPeriod(t *testing.T) {
	tt := &domain.Subject{
		Name:   "Project",
		Course: 0,
		Period: 2,
		Grades: 1,
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	tt.Id = id

	subjects, err := r.GetByCourseAndPeriod(tt.Course, tt.Period)
	if err != nil {
		t.Fatal(err)
	}

	if len(subjects) != 1 {
		t.Fatalf("invalid subjects length. expected=%d. got=%d", 1, len(subjects))
	}

	s := subjects[0]

	if s.Name != tt.Name {
		t.Fatalf("subject name expected=%q. got=%q", tt.Name, s.Name)
	}
}
