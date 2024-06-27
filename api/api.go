package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jasurxaydarov/edu/api/handlers/v1"
	"github.com/jasurxaydarov/edu/storage"
)
func Api(storage storage.StorageI){

	router:= gin.Default()

	h:=v1.NewHandler(storage)



	router.GET("/ping",)

	router.POST("/teacher",h.TeacherRepo().GetTeacherById)
	router.GET("/teacher/:id",h.TeacherRepo().GetTeacherById)
	router.GET("/teachers",h.TeacherRepo().GetTeachers)
	router.DELETE("/teacher/:id",h.TeacherRepo().DeleteTeacher)
	router.PUT("/teacher/:id",h.TeacherRepo().UpdateTeacher)

	router.POST("/course",h.CourseRepo().CreateCourse)
	router.GET("/course/:id",h.CourseRepo().GetCourseById)
	router.GET("/courses",h.CourseRepo().GetCourses)
	router.DELETE("/course/:id",h.CourseRepo().DeleteCourse)
	router.PUT("/course/:id",h.CourseRepo().UpdateCourse)

	router.Run(":8080")

}