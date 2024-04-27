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
	"github.com/free5gc/udr/internal/util"
)

// HTTPQuerySmsMngData - Retrieves the SMS management subscription data of a UE
func (p *Processor) HandleQuerySmsMngData(c *gin.Context) {
	logger.DataRepoLog.Infof("Handle QuerySmsMngData")

	collName := "subscriptionData.provisionedData.smsMngData"
	ueId := c.Params.ByName("ueId")
	servingPlmnId := c.Params.ByName("servingPlmnId")
	response, problemDetails := datarepository.QuerySmsMngDataProcedure(collName, ueId, servingPlmnId)

	if response == nil && problemDetails == nil {
		pd := util.ProblemDetailsUpspecified("")
		c.JSON(int(pd.Status), pd)
	} else if problemDetails != nil {
		c.JSON(int(problemDetails.Status), problemDetails)
	}
	c.JSON(http.StatusOK, response)
}
