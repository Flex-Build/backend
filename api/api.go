// Package api provide support to create /api group
package api

import (
	"github.com/Flexi-Build/backend/api/publish"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies the /api group and v1 routes to given gin Engine
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		publish.ApplyRoutes(api)
	}
}
