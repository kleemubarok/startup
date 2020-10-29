package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"startup/campaign"
	"startup/helper"
	"strconv"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		response := helper.APIRespose("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIRespose("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context)  {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response:=helper.APIRespose("Error to get campaignDetails details", http.StatusBadRequest,"error",nil)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	campaignDetails, errs := h.campaignService.GetCampaign(input)
	if errs != nil {
		response:=helper.APIRespose("Error to get campaignDetails details", http.StatusBadRequest,"error",errs)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	response:=helper.APIRespose("Campaign details",http.StatusOK,"success", campaign.FormatCampaignDetails(campaignDetails))
	c.JSON(http.StatusOK,response)
}
