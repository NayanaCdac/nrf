// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package accesstoken

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/http_wrapper"
	"github.com/free5gc/nrf/logger"
	"github.com/free5gc/nrf/producer"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

// AccessTokenRequest - Access Token Request
func HTTPAccessTokenRequest(c *gin.Context) {
	logger.AccessTokenLog.Infoln("In HTTPAccessTokenRequest")
	var accessTokenReq models.AccessTokenReq

	// logger.AccessTokenLog.Infoln("Content Type: ", c.ContentType())
	err := c.Bind(&accessTokenReq)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.AccessTokenLog.Warnln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, accessTokenReq)
	req.Params["paramName"] = c.Params.ByName("paramName")

	httpResponse := producer.HandleAccessTokenRequest(req)

	responseBody, err := openapi.Serialize(httpResponse.Body, "application/json")

	if err != nil {
		logger.AccessTokenLog.Warnln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.JSON(httpResponse.Status, responseBody)
	}
}
