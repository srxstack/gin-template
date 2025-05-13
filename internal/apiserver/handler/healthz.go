package handler

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/srxstack/srxstack/pkg/log"

	apiv1 "github.com/srxstack/gin-template/pkg/api/apiserver/v1"
)

// Healthz 服务健康检查.
func (h *Handler) Healthz(c *gin.Context) {
	log.W(c.Request.Context()).Infow("Healthz handler is called", "method", "Healthz", "status", "healthy")
	c.JSON(200, &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	})
}
