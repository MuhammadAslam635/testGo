package storage

import "example/hello/cmd/internal/types"

type Storage interface {
	CreateStudent(Name string, Email string, Age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
}
