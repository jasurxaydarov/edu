package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/edu/storage/postgres"
	"github.com/jasurxaydarov/edu/storage/repoi"
)

type StorageI interface{

	TeacherRepo() repoi.TeacherRepoI
}


type storage struct{
	teacherRepo repoi.TeacherRepoI
}

func NewStorage( db *pgx.Conn)StorageI{

	return &storage{
		teacherRepo: postgres.NewTeacherRepo(db),
	}
}


func (s *storage)TeacherRepo()repoi.TeacherRepoI{
	return s.teacherRepo
}