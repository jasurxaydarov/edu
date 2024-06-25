package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jasurxaydarov/edu/api/handlers/v1"
	"github.com/jasurxaydarov/edu/storage"
)
func Api(storage storage.StorageI){

	router:= gin.Default()

	h:=v1.NewHandler(storage)

	router.GET("/ping",h.Ping)
	router.POST("/teacher",h.CreateTeacher)

	router.Run(":8080")
}