package handler

import (
	"fmt"
	"fuckregex/internal/db"
	"fuckregex/model/api_model"
	"fuckregex/model/db_model"
	"github.com/gin-gonic/gin"
	"time"
)

func Latest(ctx *gin.Context) {

	var responses []db_model.Responses
	if err := db.DB.Order("created_date DESC").Limit(25).Find(&responses).Error; err != nil {
		ctx.JSON(500, api_model.APIError{
			ErrorCode:    500,
			ErrorMessage: "Internal server error",
			ResponseTime: time.Now().Unix(),
		})
		return
	}

	list := ""

	for _, response := range responses {
		list += fmt.Sprintf(
			"<li class='flex items-center'><a href='/?id=%s' class='truncate'>%s</a> <span class='text-gray-700 ml-auto'>%d</span></li>", response.ID, response.InputText, response.CreatedDate)
	}

	ctx.JSON(200, api_model.APIResponse{
		ResponseText: list,
	})
}
