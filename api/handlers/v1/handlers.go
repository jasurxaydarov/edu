package v1

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

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
	err:= h.Storage.TeacherRepo().CreateTeacher(ctx, teacher)

	if err!= nil{

		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"succesfully created")

}

func (h *Handler) GetTeacherById(ctx *gin.Context) {

	id:=ctx.Param("id")

	teacher, err := h.Storage.TeacherRepo().GetTeacherById(context.Background(), id)

	if err != nil {
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(200, teacher)

}

func (h *Handler) GetTeachers(ctx *gin.Context) {

	var req modles.GetListReq

	Limit := ctx.Query("limit")
	Page := ctx.Query("page")

	req.Limit,_=strconv.Atoi(Limit)
	req.Page,_=strconv.Atoi(Page)
	
	teacher, err := h.Storage.TeacherRepo().GetTeachers(context.Background(), req)

	if err != nil {
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(200, teacher)

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

	id:=ctx.Param("id")

	err := h.Storage.TeacherRepo().DeleteTeacher(context.Background(), id)

	if err != nil {
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"succesfully deleted")
}


func(h *Handler)UpdateTeacher(ctx *gin.Context){

	var req modles.TeacherReq

	id:=ctx.Param("id")

	ctx.BindJSON(&req)

	err:=h.Storage.TeacherRepo().UpdateTeacher(context.Background(),id,req)

	if err!= nil{
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"succesfully updated")

}