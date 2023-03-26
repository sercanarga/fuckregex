package handler

import (
	"fuckregex/internal"
	"fuckregex/internal/db"
	"fuckregex/model/api_model"
	"fuckregex/model/db_model"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"time"
)

func Generate(ctx *gin.Context) {
	var req api_model.Generate
	err := ctx.BindJSON(&req)
	p := bluemonday.StripTagsPolicy()
	userInput := p.Sanitize(req.Desc)

	if err != nil || userInput == "" {
		ctx.JSON(400, api_model.APIError{
			ErrorCode:    400,
			ErrorMessage: "Invalid format",
			ResponseTime: time.Now().Unix(),
		})
		return
	}

	var response db_model.Responses
	db.DB.Raw("SELECT * FROM responses WHERE input_text = ?", userInput).Scan(&response)

	if response.ID != "" {
		ctx.JSON(200, api_model.APIResponse{
			ResponseID:    response.ID,
			InputText:     response.InputText,
			ResponseText:  response.ResponseText,
			ResponseToken: response.ResponseToken,
			ResponseTime:  response.CreatedDate,
		})
		return
	}

	validateRequest, err := internal.OpenAIRequest(req, true, 10)
	if err != nil {
		ctx.JSON(500, api_model.APIError{
			ErrorCode:    500,
			ErrorMessage: "Internal server error",
			ResponseTime: time.Now().Unix(),
		})
		log.Println(err)
		return
	}

	if validateRequest.Error.Message != "" {
		ctx.JSON(500, api_model.APIError{
			ErrorCode:    500,
			ErrorMessage: validateRequest.Error.Message,
			ResponseTime: time.Now().Unix(),
		})
		return
	}

	if validateRequest.Choices[0].Message.Content != "true" {
		ctx.JSON(400, api_model.APIError{
			ErrorCode:    400,
			ErrorMessage: "Invalid request",
			ResponseTime: time.Now().Unix(),
		})
		return
	}

	aiRequest, err := internal.OpenAIRequest(req, false, 200)
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
		InputText:     userInput,
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
