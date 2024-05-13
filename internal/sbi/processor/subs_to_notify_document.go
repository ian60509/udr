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
	"github.com/free5gc/util/httpwrapper"
)

// HTTPRemovesubscriptionDataSubscriptions - Deletes a subscriptionDataSubscriptions
func (p *Processor) HandleRemovesubscriptionDataSubscriptions(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	rsp := datarepository.HandleRemovesubscriptionDataSubscriptions(req)

	logger.DataRepoLog.Infof("Handle RemovesubscriptionDataSubscriptions")

	subsId := c.Params.ByName("subsId")

	problemDetails := datarepository.RemovesubscriptionDataSubscriptionsProcedure(subsId)

	if problemDetails != nil {
		c.JSON(http.StatusInternalServerError, problemDetails)
		return
	}
	c.JSON(http.StatusNoContent, rsp)
}