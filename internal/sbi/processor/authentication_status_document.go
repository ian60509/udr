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
	"github.com/free5gc/udr/internal/util"
)

// HTTPCreateAuthenticationStatus - To store the Authentication Status data of a UE
func (p *Processor) HandleCreateAuthenticationStatus(c *gin.Context) {
	var authEvent models.AuthEvent

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

	err = openapi.Deserialize(&authEvent, requestBody, "application/json")
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

	logger.DataRepoLog.Infof("Handle CreateAuthenticationStatus")

	putData := util.ToBsonM(authEvent)
	ueId := c.Params.ByName("ueId")
	collName := "subscriptionData.authenticationData.authenticationStatus"

	datarepository.CreateAuthenticationStatusProcedure(collName, ueId, putData)

	c.Status(http.StatusNoContent)
}

// HTTPQueryAuthenticationStatus - Retrieves the Authentication Status of a UE
func (p *Processor) HandleQueryAuthenticationStatus(c *gin.Context) {
	logger.DataRepoLog.Infof("Handle QueryAuthenticationStatus")

	ueId := c.Params.ByName("ueId")
	collName := "subscriptionData.authenticationData.authenticationStatus"

	data, problemDetails := datarepository.QueryAuthenticationStatusProcedure(collName, ueId)

	if data == nil && problemDetails == nil {
		pd := util.ProblemDetailsUpspecified("")
		c.JSON(int(pd.Status), pd)
		return
	} else if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
		return
	}
	c.JSON(http.StatusOK, data)
}