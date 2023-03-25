package handler

import (
	"fuckregex/internal"
	"fuckregex/internal/db"
	"fuckregex/model/api_model"
	"fuckregex/model/db_model"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Generate(ctx *gin.Context) {
	var req api_model.Generate
	err := ctx.BindJSON(&req)
	if err != nil /*|| req.Desc == "" || req.Type == 0*/ {
		ctx.JSON(400, api_model.APIError{
			ErrorCode:    400,
			ErrorMessage: "Invalid format",
			ResponseTime: time.Now().Unix(),
		})
		return
	}

	aiRequest, err := internal.OpenAIRequest(req)
	if err != nil {
		ctx.JSON(500, api_model.APIError{
			ErrorCode:    500,
			ErrorMessage: "Internal server error",
			ResponseTime: time.Now().Unix(),
		})
		log.Println(err)
		return
	}

	if aiRequest.Error.Message != "" {
		ctx.JSON(500, api_model.APIError{
			ErrorCode:    500,
			ErrorMessage: aiRequest.Error.Message,
			ResponseTime: time.Now().Unix(),
		})
		return
	}

	createResponse := db_model.Responses{
		ID:            aiRequest.ID,
		InputText:     req.Desc,
		ResponseText:  aiRequest.Choices[0].Message.Content,
		CreatedDate:   aiRequest.ResponseTime,
		ResponseToken: aiRequest.Usage.TotalToken,
	}
	result := db.DB.Create(&createResponse)
	if result.Error != nil {
		ctx.JSON(500, api_model.APIError{
			ErrorCode:    500,
			ErrorMessage: "Internal server error",
			ResponseTime: time.Now().Unix(),
		})
		log.Println(result.Error)
		return
	}

	ctx.JSON(200, api_model.APIResponse{
		ResponseID:    aiRequest.ID,
		ResponseText:  aiRequest.Choices[0].Message.Content,
		ResponseToken: aiRequest.Usage.TotalToken,
		ResponseTime:  aiRequest.ResponseTime,
	})
}
