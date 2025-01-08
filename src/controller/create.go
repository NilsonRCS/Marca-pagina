package controller

import (
	"fmt"

	"github.com/NilsonRCS/Marca-pagina/src/controller/model/request"
	global_errors "github.com/NilsonRCS/Marca-pagina/src/utils"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var modelRequest request.ModelRequest

	if err := c.ShouldBindJSON(&modelRequest); err != nil {
		restErr := global_errors.NewBadRequestError(
			fmt.Sprintf("algum deu errado nos campos informado error=%s", err),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(modelRequest)
}
