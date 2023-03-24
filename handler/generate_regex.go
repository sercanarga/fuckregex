package handler

import (
	"fuckregex/model/api_model"
	"github.com/gin-gonic/gin"
	"time"
)

func GenerateRegex(ctx *gin.Context) {
	var req api_model.Generate
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(400, api_model.APIError{
			ErrorCode:    400,
			ErrorMessage: "Invalid JSON",
			ResponseTime: time.Now().Unix(),
		})
		return
	}

}
