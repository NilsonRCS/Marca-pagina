package routes

import (
	"github.com/NilsonRCS/Marca-pagina/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/BOOKS/get", controller.FindAllBooks)
	r.GET("/BOOKS/get/:bookId", controller.FindBookByID)
	r.POST("/BOOKS/create", controller.CreateBook)
	r.PUT("/BOOKS/update/:bookId", controller.UpdateBookByID)
	r.DELETE("/BOOKS/delete/:bookId", controller.DeleteBookByID)
}
