package v1

import (
	"github.com/gin-gonic/gin"
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