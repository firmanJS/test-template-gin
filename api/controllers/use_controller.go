package controllers

import (
	"net/http"
	"strconv"

	"github.com/dipeshdulal/clean-gin/constants"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UseController data type
type UseController struct {
	service services.UseService
	logger  lib.Logger
}

// NewUseController creates new use controller
func NewUseController(useService services.UseService, logger lib.Logger) UseController {
	return UseController{
		service: useService,
		logger:  logger,
	}
}

// GetOneUse gets one use
func (u UseController) GetOneUse(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	use, err := u.service.GetOneUse(uint(id))

	if err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": use,
	})

}

// GetUse gets the use
func (u UseController) GetUse(c *gin.Context) {
	uses, err := u.service.GetAllUse()
	if err != nil {
		u.logger.Zap.Error(err)
	}
	c.JSON(200, gin.H{"data": uses})
}

// SaveUse saves the use
func (u UseController) SaveUse(c *gin.Context) {
	use := models.User{}
	trxHandle := c.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.ShouldBindJSON(&use); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.WithTrxUse(trxHandle).CreateUse(use); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "use created"})
}

// UpdateUse updates use
func (u UseController) UpdateUse(c *gin.Context) {
	c.JSON(200, gin.H{"data": "use updated"})
}

// DeleteUse deletes use
func (u UseController) DeleteUse(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	if err := u.service.DeleteUse(uint(id)); err != nil {
		u.logger.Zap.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "use deleted"})
}
