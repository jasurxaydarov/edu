package v1

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/edu/modles"
)

func ( h *Handler)CreateGroup(ctx *gin.Context){

	var group modles.GradeReq
	var resp =&modles.Group{}


	ctx.BindJSON(group)
	DataParser(group,resp)


	err:=h.Storage.GroupRepo().CreateGroup(context.Background(),resp)

	if err!= nil{

		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"sucessfully created")

}

func(h *Handler)GetGrupById(ctx *gin.Context){

	req:=ctx.Param("id")

	err,resp:=h.Storage.GroupRepo().GetGrupById(context.Background(),req)

	if err!= nil{
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(200,resp)
}

func (h *Handler)GetGroups(ctx *gin.Context){

	var req modles.GetListReq

	limit:=ctx.Query("limit")
	page:=ctx.Query("page")

	req.Limit,_=strconv.Atoi(limit)
	req.Page,_=strconv.Atoi(page)

	

	resp,err:=h.Storage.GroupRepo().GetGroups(context.Background(),req)

	if err!= nil{

		ctx.JSON(500,err)
		return
	}

	ctx.JSON(200,resp)

}

func(h *Handler)DeleteGroup(ctx *gin.Context){

	id:=ctx.Param("id")

	err:=h.Storage.GroupRepo().DeleteGroup(context.Background(),id)
	if err!= nil{

		ctx.JSON(201,err)
		return
	}

	ctx.JSON(201,"sucessfully deleted")
}

func ( h *Handler)UpdateGroup(ctx *gin.Context){
	var req modles.GroupReq

	id:=ctx.Param("id")

	ctx.BindJSON(req)

	err:=h.Storage.GroupRepo().UpdateGroup(ctx,id,req)

	if err!= nil {

		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,"sucessfully updated")
}