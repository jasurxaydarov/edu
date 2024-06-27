package repoi

import "github.com/gin-gonic/gin"

type TeacherRepoI interface{
	CreateTeacher(ctx *gin.Context)
	GetTeacherById(ctx *gin.Context)
	GetTeachers(ctx *gin.Context) 
	UpdateTeacher(ctx *gin.Context)
	DeleteTeacher(ctx *gin.Context)

}
type CourseRepoI interface{
	CreateCourse(ctx *gin.Context)
	GetCourseById(ctx *gin.Context)
	GetCourses(ctx *gin.Context) 
	UpdateCourse(ctx *gin.Context)
	DeleteCourse(ctx *gin.Context)

}