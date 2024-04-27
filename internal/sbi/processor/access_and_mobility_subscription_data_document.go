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

	"github.com/free5gc/udr/internal/logger"
	datarepository "github.com/free5gc/udr/internal/sbi/datarepository"
)

// HTTPQueryAmData - Retrieves the access and mobility subscription data of a UE
func (p *Processor) HandleQueryAmData(c *gin.Context) {
	//--------
	logger.DataRepoLog.Infof("Handle QueryAmData")

	collName := "subscriptionData.provisionedData.amData"
	ueId := c.Params.ByName("ueId")
	servingPlmnId := c.Params.ByName("servingPlmnId")
	response, problemDetails := datarepository.QueryAmDataProcedure(collName, ueId, servingPlmnId)

	if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
		return 
	}
	c.JSON(http.StatusOK, response)
}
