package repoi

import (
	"context"

	"github.com/jasurxaydarov/edu/modles"
)

type TeacherRepoI interface {
	CreateTeacher(ctx context.Context, req *modles.Teacher) error
	GetTeacherById(ctx context.Context, id string) (*modles.Teacher, error)
	GetTeacher(ctx context.Context, req modles.GetListReq) (*[]modles.Teacher, error)
	UpdateTeacher(ctx context.Context ,id string , req modles.TeacherReq )error
	DeleteTeacher(ctx context.Context,id string)error
}
