package repoi

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jasurxaydarov/edu/modles"
)

type TeacherRepoI interface{

	CreateTeacher(ctx context.Context, req *modles.Teacher)error
	GetTeacherById(ctx context.Context, id uuid.UUID)(*modles.Teacher, error)
	GetTeacher(ctx context.Context, req modles.GetListReq)(*modles.GetTeachersListResp, error)
}