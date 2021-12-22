package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

type IConfigAPI interface {
	GetConfig(*gin.Context)
}

type ConfigAPI struct {
}

// GetConfig
// @Summary Get the full configuration file for the server.
// @Description Gets the app settings for the environment the server is running in.
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "The configuration file is returned."
// @Router /config [get]
func (configApi *ConfigAPI) GetConfig(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, viper.AllSettings())
}
