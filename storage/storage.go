package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/edu/storage/postgres"
	"github.com/jasurxaydarov/edu/storage/repoi"
)

type StorageI interface{

	TeacherRepo() repoi.TeacherRepoI
	CourseRepo() repoi.CourseRepoI
	GroupRepo() repoi.GroupRepoI
	
}


type storage struct{
	teacherRepo repoi.TeacherRepoI
	courseRepo repoi.CourseRepoI
	groupRepo repoi.GroupRepoI
}

func NewStorage( db *pgx.Conn)StorageI{

	return &storage{
		teacherRepo: postgres.NewTeacherRepo(db),
		courseRepo: postgres.NewCourseRepo(db),
		groupRepo: postgres.NewGroupRepo(db),
	}
}


func (s *storage)TeacherRepo()repoi.TeacherRepoI{
	return s.teacherRepo
}

func (s *storage)CourseRepo()repoi.CourseRepoI{
	return s.courseRepo
}

func(s *storage)GroupRepo()repoi.GroupRepoI{

	return s.groupRepo
}