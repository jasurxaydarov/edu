package v1

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/edu/modles"
	"github.com/jasurxaydarov/edu/storage"
)

type Handler struct{
	Storage storage.StorageI
}

func NewHandler( storage storage.StorageI)Handler{

	return Handler{Storage: storage}
}

func (h *Handler)Ping(ctx *gin.Context){

	ctx.JSON(200,map[string]string{"massage":"pong"})
}
func (h *Handler)CreateTeacher(ctx *gin.Context){

	var reBody modles.TeacherReq
	var teacher =&modles.Teacher{}

	ctx.BindJSON(&reBody)

	DataParser(reBody,teacher)
	h.Storage.TeacherRepo().CreateTeacher(ctx,teacher)
	

}

func DataParser[t any,t2 any](src t,dst t2){
	byData,err:=json.Marshal(src)

	if err!= nil{
		log.Println("err on dataParser",err)
		return
	}

	json.Unmarshal(byData,dst)
}