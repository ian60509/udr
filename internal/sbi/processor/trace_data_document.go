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
	"go.mongodb.org/mongo-driver/bson"

	"github.com/free5gc/udr/internal/logger"
)

func (p *Processor) QueryTraceDataProcedure(c *gin.Context, collName string, ueId string,
	servingPlmnId string,
) {
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}
	data, pd := p.GetDataFromDB(collName, filter)
	if pd != nil {
		logger.DataRepoLog.Errorf("QueryTraceDataProcedure err: %s", pd.Detail)
		c.JSON(int(pd.Status), pd)
		return
	}
	c.JSON(http.StatusOK, data)
}
