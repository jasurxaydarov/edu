package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/jasurxaydarov/edu/modles"
	"github.com/jasurxaydarov/edu/storage"
)

type Handler struct {
	Storage storage.StorageI
}

func NewHandler(storage storage.StorageI) Handler {

	return Handler{Storage: storage}
}

func (h *Handler) Ping(ctx *gin.Context) {

	ctx.JSON(200, map[string]string{"massage": "pong"})
}

func (h *Handler) CreateTeacher(ctx *gin.Context) {

	var reBody modles.TeacherReq
	var teacher = &modles.Teacher{}

	ctx.BindJSON(&reBody)

	DataParser(reBody, teacher)
	h.Storage.TeacherRepo().CreateTeacher(ctx, teacher)

}

func (h *Handler) GetTeacherById(ctx *gin.Context) {

	var url = strings.Split(ctx.Request.URL.Path, "/")

	id := url[len(url)]

	fmt.Println("iddddddddddd", url)

	teacher, err := h.Storage.TeacherRepo().GetTeacherById(context.Background(), id)

	if err != nil {

		return
	}

	ctx.JSON(2001, teacher)

}

func (h *Handler) GetTeacher(ctx *gin.Context) {

	var req modles.GetListReq

	req.Limit = 5
	req.Page = 1
	teacher, err := h.Storage.TeacherRepo().GetTeacher(context.Background(), req)
	fmt.Println("teacherrrr", teacher)

	if err != nil {

		return
	}

	ctx.JSON(201, teacher)

}

func DataParser[t any, t2 any](src t, dst t2) {
	byData, err := json.Marshal(src)

	if err != nil {
		log.Println("err on dataParser", err)
		return
	}

	json.Unmarshal(byData, dst)
}

func (h *Handler) DeleteTeacher(ctx *gin.Context) {

	urlS := strings.Split(ctx.Request.URL.Scheme, "/")

	id := urlS[len(urlS)-1]

	err := h.Storage.TeacherRepo().DeleteTeacher(context.Background(), id)

	if err != nil {

		return
	}

}
