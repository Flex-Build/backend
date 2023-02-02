package publish

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Flexi-Build/backend/pkg/caddy"
	"github.com/Flexi-Build/backend/pkg/envconfig"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/cloudflare/cloudflare-go"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/publish")
	{
		g.POST("", publish)
	}
}

func publish(c *gin.Context) {
	var body PublishRequest
	err := c.BindJSON(&body)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "failed to validate body").
			Send(c, http.StatusBadRequest)
		return
	}
	api, err := cloudflare.NewWithAPIToken(envconfig.EnvVars.CLOUDFLARE_API_TOKEN)
	if err != nil {
		logo.Error("failed to create client", err)
		httpo.NewErrorResponse(500, "failed to deploy site").SendD(c)
		return
	}
	ctx := context.Background()
	res, err := api.CreateDNSRecord(ctx, cloudflare.ResourceIdentifier(envconfig.EnvVars.ZONE_ID),
		cloudflare.CreateDNSRecordParams{Name: body.Name, Content: envconfig.EnvVars.SERVER_IP, Type: "A"})
	if err != nil {
		logo.Error("failed to create dns", err)
		httpo.NewErrorResponse(500, "failed to deploy site").SendD(c)
		return
	}
	domain := fmt.Sprintf("%s.ommore.me", body.Name)
	folder_name := fmt.Sprintf("/static/%s", domain)
	file_name := fmt.Sprintf("%s/index.html", folder_name)
	err = os.Mkdir(folder_name, os.ModePerm)
	if err != nil {
		logo.Error("failed to create directory", err)
		httpo.NewErrorResponse(500, "failed to deploy site").SendD(c)
		return
	}
	file, err := os.Create(file_name)
	if err != nil {
		logo.Error("failed to create file", err)
		httpo.NewErrorResponse(500, "failed to deploy site").SendD(c)
		return
	}
	_, err = file.WriteString(body.HtmlString)
	if err != nil {
		logo.Error("failed to write file", err)
		httpo.NewErrorResponse(500, "failed to deploy site").SendD(c)
		return
	}

	err = caddy.UpdateCaddy(domain)
	if err != nil {
		logo.Error("failed to update caddy config", err)
		httpo.NewErrorResponse(500, "failed to deploy site").SendD(c)
		return
	}
	if res.Success {
		httpo.NewSuccessResponse(http.StatusOK, "site deployed successfully").
			SendD(c)
	} else {
		httpo.NewErrorResponse(500, "failed to deploy site").SendD(c)
	}

}
