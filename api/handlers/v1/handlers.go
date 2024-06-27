package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/jasurxaydarov/edu/api/handlers/repoi"
	"github.com/jasurxaydarov/edu/storage"
)

type HandlerI interface{

	TeacherRepo() repoi.TeacherRepoI
	CourseRepo() repoi.CourseRepoI
	
}

type Handler struct {
	Storage 	storage.StorageI
	teacherRepo repoi.TeacherRepoI
	courseRepo	repoi.CourseRepoI
}

func NewHandler(storage storage.StorageI) HandlerI {

	return &Handler{Storage: storage}
}

func (h *Handler) Ping(ctx *gin.Context) {

	ctx.JSON(200, map[string]string{"massage": "pong"})
}

func(h *Handler)TeacherRepo()repoi.TeacherRepoI{
	return h.teacherRepo
}

func(h *Handler)CourseRepo()repoi.CourseRepoI{

	return h.courseRepo
}