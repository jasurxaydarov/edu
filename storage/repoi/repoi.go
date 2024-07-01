package repoi

import (
	"context"

	"github.com/jasurxaydarov/edu/modles"
)

type TeacherRepoI interface {
	CreateTeacher(ctx context.Context, req *modles.Teacher) error
	GetTeacherById(ctx context.Context, id string) (*modles.Teacher, error)
	GetTeachers(ctx context.Context, req modles.GetListReq) (*[]modles.Teacher, error)
	UpdateTeacher(ctx context.Context ,id string , req modles.TeacherReq )error
	DeleteTeacher(ctx context.Context,id string)error
}



type CourseRepoI interface{
	CreateCourse(ctx context.Context, req *modles.Course) error
	GetCourseById(ctx context.Context, id string) (*modles.Course, error)
	GetTCourses(ctx context.Context, req modles.GetListReq) (*[]modles.Course, error)
	UpdateCourse(ctx context.Context ,id string , req modles.CourseReq )error
	DeleteCourse(ctx context.Context,id string)error

}

type GroupRepoI interface{
	CreateGroup(ctx context.Context, req *modles.Group) error
	GetGrupById(ctx context.Context, id string) (*modles.Group, error)
	GetGroups(ctx context.Context, req modles.GetListReq) (*[]modles.Group, error)
	UpdateGroup(ctx context.Context ,id string , req modles.GroupReq )error
	DeleteGroup(ctx context.Context,id string)error

}