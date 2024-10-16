package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"project/models"
	"project/services"

	"github.com/gin-gonic/gin"
)

type ResultsController struct{
	resultsService *services.ResultsService
}

func NewResultController(resultsService *services.ResultsService) *ResultsController{
	return &ResultsController{
		resultsService: resultsService,
	}
}

func (rc ResultsController) CreateResult(ctx *gin.Context){
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create result request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	var result models.Result
	err = json.Unmarshal(body, &result)
	if err != nil{
		log.Println("error while un marshaling create result request body")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return 
	}

	response, responseErr := rc.resultsService.CreateResult(&result)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (rc ResultsController) DeleteResult(ctx *gin.Context) {
	resultId := ctx.Param("id")
	responseErr := rc.resultsService.DeleteResult(resultId)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.Status(http.StatusNoContent)
}