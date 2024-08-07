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

	"github.com/free5gc/openapi/models"
	udr_context "github.com/free5gc/udr/internal/context"
	"github.com/free5gc/udr/internal/logger"
	"github.com/free5gc/udr/internal/util"
)

func (p *Processor) CreateAMFSubscriptionsProcedure(c *gin.Context, subsId string, ueId string,
	AmfSubscriptionInfo []models.AmfSubscriptionInfo,
) {
	udrSelf := udr_context.GetSelf()
	value, ok := udrSelf.UESubsCollection.Load(ueId)
	if !ok {
		pd := util.ProblemDetailsNotFound("USER_NOT_FOUND")
		logger.DataRepoLog.Errorf("CreateAMFSubscriptionsProcedure err: %s", pd.Detail)
		c.JSON(int(pd.Status), pd)
		return
	}
	UESubsData := value.(*udr_context.UESubsData)

	_, ok = UESubsData.EeSubscriptionCollection[subsId]
	if !ok {
		pd := util.ProblemDetailsNotFound("SUBSCRIPTION_NOT_FOUND")
		logger.DataRepoLog.Errorf("CreateAMFSubscriptionsProcedure err: %s", pd.Detail)
		c.JSON(int(pd.Status), pd)
		return
	}

	UESubsData.EeSubscriptionCollection[subsId].AmfSubscriptionInfos = AmfSubscriptionInfo
	c.Status(http.StatusNoContent)
}

func (p *Processor) RemoveAmfSubscriptionsInfoProcedure(c *gin.Context, subsId string, ueId string) {
	udrSelf := udr_context.GetSelf()
	value, ok := udrSelf.UESubsCollection.Load(ueId)
	var pd *models.ProblemDetails = nil

	if !ok {
		pd = util.ProblemDetailsNotFound("USER_NOT_FOUND")
		logger.DataRepoLog.Errorf("RemoveAmfSubscriptionsInfoProcedure err: %s", pd.Detail)
	}

	UESubsData := value.(*udr_context.UESubsData)
	_, ok = UESubsData.EeSubscriptionCollection[subsId]

	if !ok {
		pd = util.ProblemDetailsNotFound("SUBSCRIPTION_NOT_FOUND")
		logger.DataRepoLog.Errorf("RemoveAmfSubscriptionsInfoProcedure err: %s", pd.Detail)
	}

	if UESubsData.EeSubscriptionCollection[subsId].AmfSubscriptionInfos == nil {
		pd = util.ProblemDetailsNotFound("AMFSUBSCRIPTION_NOT_FOUND")
	}

	if pd != nil {
		logger.DataRepoLog.Errorf("RemoveAmfSubscriptionsInfoProcedure err: %s", pd.Detail)
		c.JSON(int(pd.Status), pd)
		return
	}

	UESubsData.EeSubscriptionCollection[subsId].AmfSubscriptionInfos = nil
	c.Status(http.StatusNoContent)
}
