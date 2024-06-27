package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/edu/storage/postgres"
	"github.com/jasurxaydarov/edu/storage/repoi"
)

type StorageI interface{

	TeacherRepo() repoi.TeacherRepoI
	CourseRepo() repoi.CourseRepoI
}


type storage struct{
	teacherRepo repoi.TeacherRepoI
	courseRepo repoi.CourseRepoI
}

func NewStorage( db *pgx.Conn)StorageI{

	return &storage{
		teacherRepo: postgres.NewTeacherRepo(db),
		courseRepo: postgres.NewCourseRepo(db),
	}
}


func (s *storage)TeacherRepo()repoi.TeacherRepoI{
	return s.teacherRepo
}

func (s *storage)CourseRepo()repoi.CourseRepoI{
	return s.courseRepo
}