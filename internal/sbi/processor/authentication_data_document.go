/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package processor

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/udr/internal/logger"
	datarepository "github.com/free5gc/udr/internal/sbi/datarepository"
)

// HTTPModifyAuthentication - modify the authentication subscription data of a UE
func (p *Processor) HandleModifyAuthentication(c *gin.Context) {
	var patchItemArray []models.PatchItem

	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.DataRepoLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&patchItemArray, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.DataRepoLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	logger.DataRepoLog.Infof("Handle ModifyAuthentication")

	collName := "subscriptionData.authenticationData.authenticationSubscription"
	ueId := c.Params.ByName("ueId")

	problemDetails := datarepository.ModifyAuthenticationProcedure(collName, ueId, patchItemArray)

	if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
		return
	}
	c.Status(http.StatusNoContent)
}

// HTTPQueryAuthSubsData - Retrieves the authentication subscription data of a UE
func (p *Processor) HandleQueryAuthSubsData(c *gin.Context) {
	logger.DataRepoLog.Infof("Handle QueryAuthSubsData")

	collName := "subscriptionData.authenticationData.authenticationSubscription"
	ueId :=  c.Params.ByName("ueId")

	data, problemDetails := datarepository.QueryAuthSubsDataProcedure(collName, ueId)

	if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
		return
	}
	c.JSON(http.StatusOK, data)
}