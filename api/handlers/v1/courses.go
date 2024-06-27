package v1

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/edu/modles"
)

func(h *Handler)CreateCourse(ctx *gin.Context){

	var (
		req modles.CourseReq
		course =&modles.Course{}
	)
	
	err:=ctx.BindJSON(&req)

	DataParser(req,course)

	if err!=nil{
		ctx.JSON(400,err)
		return
	}

	err= h.Storage.CourseRepo().CreateCourse(context.Background(),course)
	
	if err!=nil{
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"succesfully created")


}


func (h *Handler)GetCourseById(ctx *gin.Context){

	id:=ctx.Param("id")


	resp,err:=h.Storage.CourseRepo().GetCourseById(context.Background(),id)

	if err!= nil{
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(200,resp)
}

func (h *Handler)GetCourses(ctx *gin.Context){

	var req modles.GetListReq

	limit:=ctx.Query("limit")
	page:=ctx.Query("page")

	req.Limit,_=strconv.Atoi(limit)
	req.Page,_=strconv.Atoi(page)
	
	resp,err:=h.Storage.CourseRepo().GetTCourses(context.Background(),req)

	if err !=nil{

		ctx.JSON(500,err)
		return
	}

	ctx.JSON(200,resp)

}


func (h *Handler)DeleteCourse( ctx *gin.Context){

	id:=ctx.Param("id")

	err:=h.Storage.CourseRepo().DeleteCourse(context.Background(),id)

	if err != nil{

		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"deleted sucessfully")

}


func (h *Handler)UpdateCourse(ctx *gin.Context){
	
	var req modles.CourseReq 
	id:=ctx.Param("id")

	ctx.BindJSON(&req)

	err:=h.Storage.CourseRepo().UpdateCourse(context.Background(),id,req)

	if err!= nil{

		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"sucessfully updated")

}