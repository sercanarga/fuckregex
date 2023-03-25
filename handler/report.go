package handler

import (
	"fuckregex/internal/db"
	"fuckregex/model/api_model"
	"fuckregex/model/db_model"
	"github.com/gin-gonic/gin"
	"time"
)

func Report(ctx *gin.Context) {
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
	db.DB.Raw("UPDATE responses SET is_reported=true WHERE id = ?", req.ID).Scan(&response)

	ctx.JSON(200, api_model.APIResponse{
		ResponseText: "reported",
	})
}
