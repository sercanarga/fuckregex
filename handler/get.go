package handler

import (
	"fuckregex/internal/db"
	"fuckregex/model/api_model"
	"fuckregex/model/db_model"
	"github.com/gin-gonic/gin"
	"time"
)

func Get(ctx *gin.Context) {
	var req api_model.Get
	err := ctx.BindJSON(&req)
	if err != nil || req.ID == "" {
		ctx.JSON(400, api_model.APIError{
			ErrorCode:    400,
			ErrorMessage: "Invalid format",
			ResponseTime: time.Now().Unix(),
		})
		return
	}

	var response db_model.Responses
	db.DB.Raw("SELECT * FROM responses WHERE id = ?", req.ID).Scan(&response)

	if response.ID == "" {
		ctx.JSON(400, api_model.APIError{
			ErrorCode:    400,
			ErrorMessage: "Invalid ID",
			ResponseTime: time.Now().Unix(),
		})
		return
	}

	ctx.JSON(200, api_model.APIResponse{
		ResponseID:    response.ID,
		InputText:     response.InputText,
		ResponseText:  response.ResponseText,
		ResponseToken: response.ResponseToken,
		ResponseTime:  response.CreatedDate,
	})
}
