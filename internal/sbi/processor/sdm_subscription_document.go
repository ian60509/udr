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

// HTTPRemovesdmSubscriptions - Deletes a sdmsubscriptions
func (p *Processor) HandleRemovesdmSubscriptions(c *gin.Context) {

	logger.DataRepoLog.Infof("Handle RemovesdmSubscriptions")

	ueId := c.Params.ByName("ueId")
	subsId := c.Params.ByName("subsId")

	problemDetails := datarepository.RemovesdmSubscriptionsProcedure(ueId, subsId)
	if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
		return
	}
	c.Status(http.StatusNoContent)
}

// HTTPUpdatesdmsubscriptions - Stores an individual sdm subscriptions of a UE
func (p *Processor) HandleUpdatesdmsubscriptions(c *gin.Context) {
	var sdmSubscription models.SdmSubscription

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

	err = openapi.Deserialize(&sdmSubscription, requestBody, "application/json")
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

	logger.DataRepoLog.Infof("Handle Updatesdmsubscriptions")

	ueId := c.Params.ByName("ueId")
	subsId := c.Params.ByName("subsId")

	problemDetails := datarepository.UpdatesdmsubscriptionsProcedure(ueId, subsId, sdmSubscription)

	if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
		return
	}
	c.Status(http.StatusNoContent)
}
