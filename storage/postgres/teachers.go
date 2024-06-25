package postgres

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/edu/modles"
	"github.com/jasurxaydarov/edu/storage/repoi"
)

type TeacherRepo struct{

	db *pgx.Conn
}

func NewTeacherRepo(db *pgx.Conn)repoi.TeacherRepoI{
	return &TeacherRepo{db: db}
}


func (t * TeacherRepo)CreatedTeacher(ctx context.Context, req modles.Teacher)error{
	return nil
}
func (t * TeacherRepo)GetTeacherById(ctx context.Context, id uuid.UUID)(*modles.Teacher, error){
	return  nil, nil
}
func (t * TeacherRepo)GetTeacher(ctx context.Context, req modles.GetListReq)(*modles.GetTeachersListResp, error){

	return nil,nil
}