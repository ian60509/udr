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
	"github.com/free5gc/util/mongoapi"
)

func (p *Processor) CreateAuthenticationSoRProcedure(c *gin.Context, collName string, ueId string, putData bson.M) {
	filter := bson.M{"ueId": ueId}
	putData["ueId"] = ueId

	if _, err := mongoapi.RestfulAPIPutOne(collName, filter, putData); err != nil {
		logger.DataRepoLog.Errorf("CreateAuthenticationSoRProcedure err: %+v", err)
	}

	c.Status(http.StatusNoContent)
}

func (p *Processor) QueryAuthSoRProcedure(c *gin.Context, collName string, ueId string) {
	filter := bson.M{"ueId": ueId}
	data, pd := p.GetDataFromDB(collName, filter)
	if pd != nil {
		logger.DataRepoLog.Errorf("QueryAuthSoRProcedure err: %s", pd.Detail)
		c.JSON(int(pd.Status), pd)
		return
	}
	c.JSON(http.StatusOK, data)
}
