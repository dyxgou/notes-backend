package student

import (
	"log/slog"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func userExists(t *testing.T, id int64) bool {
	query := "SELECT EXISTS(SELECT 1 FROM student WHERE id = ?)"

	row := r.Db.QueryRow(query, id)
	var exists bool

	err := row.Scan(&exists)
	if err != nil {
		t.Fatal(err)
	}

	return exists
}

func generateRandomName(prefix string) string {
	return prefix + strconv.Itoa(rand.Intn(10_000))
}

func TestInsertStudent(t *testing.T) {
	tt := &domain.Student{
		Name:        generateRandomName("Alejandro"),
		Course:      7,
		ParentPhone: "1231231231",
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	if !userExists(t, id) {
		t.Fatalf("user was not created. name=%q", tt.Name)
	}
}

func TestInsertWrongStudent(t *testing.T) {
	test := []struct {
		name string
		s    *domain.Student
	}{
		{
			name: "Name is longer than 30 chars",
			s: &domain.Student{
				Name:        strings.Repeat("A", 31),
				Course:      5,
				ParentPhone: "1231231231",
			},
		},
		{
			name: "Course is out of range [0, 11]",
			s: &domain.Student{
				Name:        generateRandomName("Alejandro"),
				Course:      12,
				ParentPhone: "1231231231",
			},
		},
		{
			name: "ParentPhone is invalid",
			s: &domain.Student{
				Name:        generateRandomName("Alejandro"),
				Course:      12,
				ParentPhone: "",
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			_, err := r.Insert(tt.s)
			if err == nil {
				t.Fatalf("invalid insertion expected an error")
			}

			slog.Error("inserting invalid students", "err", err)
		})
	}
}

func TestGetStudent(t *testing.T) {
	tt := &domain.Student{
		Name:        generateRandomName("Alejandro"),
		Course:      7,
		ParentPhone: "1231231231",
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	s, err := r.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	if tt.Name != s.Name {
		t.Fatalf("s.Name expected=%q. got=%q", tt.Name, s.Name)
	}
}

func TestDeleteStudent(t *testing.T) {
	tt := &domain.Student{
		Name:        generateRandomName("Alejandro"),
		Course:      10,
		ParentPhone: "1231231231",
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	studentId, err := r.Delete(id)

	if studentId == 0 {
		t.Fatalf("no student was deleted. id=%d", studentId)
	}

	if err != nil {
		t.Fatal(err)
	}

	if userExists(t, id) {
		t.Fatalf("user exists after being deleted. name=%q", tt.Name)
	}
}

func TestGetStudentsByCourse(t *testing.T) {
	test := []*domain.Student{
		{
			Name:        generateRandomName("Alejandro1-"),
			Course:      10,
			ParentPhone: "1231231231",
		},
		{
			Name:        generateRandomName("Alejandro2-"),
			Course:      10,
			ParentPhone: "1231231234",
		},
		{
			Name:        generateRandomName("Alejandro5-"),
			Course:      10,
			ParentPhone: "1231231237",
		},
		{
			Name:        generateRandomName("Alejandro8-"),
			Course:      10,
			ParentPhone: "1231231240",
		},
		{
			Name:        generateRandomName("Alejandro11-"),
			Course:      10,
			ParentPhone: "1231231243",
		},
		{
			Name:        generateRandomName("Alejandro14-"),
			Course:      10,
			ParentPhone: "1231231246",
		},
	}

	if _, err := r.Db.Exec("DELETE FROM student;"); err != nil {
		t.Fatal(err)
	}

	for _, tt := range test {
		_, err := r.Insert(tt)
		if err != nil {
			t.Fatal(err)
		}
	}

	students, err := r.GetStudentsByCourse(10)
	if err != nil {
		t.Fatal(err)
	}

	for i, tt := range test {
		s := students[i]

		if tt.Name != s.Name {
			t.Errorf("student name expected=%q. got=%q", tt.Name, s.Name)
		}
	}
}

func TestChangeName(t *testing.T) {
	tt := &domain.Student{
		Name:        generateRandomName("Alejandro"),
		Course:      7,
		ParentPhone: "1231231231",
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	tt.Id = id

	newName := generateRandomName("NewAlejandro-")
	_, err = r.ChangeName(tt.Id, newName)
	if err != nil {
		t.Fatal(err)
	}

	s, err := r.Get(tt.Id)
	if err != nil {
		t.Fatal(err)
	}

	if s.Name != newName {
		t.Fatalf("student name was not changed. expected=%q. got=%q", newName, s.Name)
	}
}

func TestChangeParentPhone(t *testing.T) {
	tt := &domain.Student{
		Name:        generateRandomName("Alejandro"),
		Course:      7,
		ParentPhone: "1231231231",
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	tt.Id = id

	newParentPhone := "3213213211"
	_, err = r.ChangeParentPhone(tt.Id, newParentPhone)
	if err != nil {
		t.Fatal(err)
	}

	s, err := r.Get(tt.Id)
	if err != nil {
		t.Fatal(err)
	}

	if s.ParentPhone != newParentPhone {
		t.Fatalf(
			"student parent_phone was not changed. expected=%q. got=%q",
			newParentPhone,
			s.ParentPhone,
		)
	}
}

func TestGetParentPhone(t *testing.T) {
	tt := &domain.Student{
		Name:        generateRandomName("Alejandro"),
		Course:      7,
		ParentPhone: "1231231231",
	}

	id, err := r.Insert(tt)
	if err != nil {
		t.Fatal(err)
	}

	tel, err := r.GetParentPhone(id)
	if err != nil {
		t.Fatal(err)
	}

	if tt.ParentPhone != tel {
		t.Fatalf("ParentPhone expected=%q. got=%q", tt.ParentPhone, tel)
	}
}
