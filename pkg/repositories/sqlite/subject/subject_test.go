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

	subject, err := r.GetByNameAndCourse(tt.Name, tt.Course, tt.Period)
	if err != nil {
		t.Fatal(err)
	}

	if subject.Name != tt.Name {
		t.Fatalf("subject name expected=%q. got=%q", tt.Name, subject.Name)
	}

	if subject.Course != tt.Course {
		t.Fatalf("subject course expected=%q. got=%q", tt.Course, subject.Course)
	}

	if subject.Period != tt.Period {
		t.Fatalf("subject period expected=%q. got=%q", tt.Period, subject.Period)
	}
}
