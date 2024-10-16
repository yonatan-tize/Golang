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

type RunnersController struct {
	runnersService *services.RunnersService
}

func NewRunnersController( runnersService *services.RunnersService) *RunnersController {
	return &RunnersController{
		runnersService: runnersService,
	}
}

func (rc RunnersController) CreateRunner(ctx *gin.Context){
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	var runner models.Runner

	err = json.Unmarshal(body, &runner)
	if err != nil{
		log.Println("error while unmarshaling create runner request body")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return 
	}

	response, responseErr := rc.runnersService.CreateRunner(&runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status,
		responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
		
}

func (rc RunnersController) UpdateRunner(ctx *gin.Context){
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create runner request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}	

	var runner models.Runner

	err = json.Unmarshal(body, &runner)
	if err != nil{
		log.Println("error while unmarshaling update runner request body")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return 
	}

	runnerID := ctx.Param("id")

	responseErr := rc.runnersService.UpdateRunner(runnerID, &runner)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return 
	}

	ctx.Status(http.StatusNoContent)
}

func (rc RunnersController) DeleteRunner(ctx *gin.Context){
	runnerID := ctx.Param("id")
	responseErr := rc.runnersService.DeleteRunner(runnerID)

	if responseErr != nil{
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return 
	}

	ctx.Status(http.StatusNoContent)
}

func (rc RunnersController) GetRunner(ctx *gin.Context){
	runnerID := ctx.Param("id")
	response, responseErr := rc.runnersService.GetRunner(runnerID)

	if responseErr != nil{
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return 
	}
	ctx.JSON(http.StatusOK, response)
}

func (rc RunnersController) GetRunnersBatch(ctx *gin.Context){
	params := ctx.Request.URL.Query()
	country := params.Get("country")
	year := params.Get("year")

	response, responseErr := rc.runnersService.GetRunnersBatch(country, year)

	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	ctx.JSON(http.StatusOK, response)	
}