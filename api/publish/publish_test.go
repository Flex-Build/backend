package publish

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Flexi-Build/backend/app/stage/appinit"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func Test_GetServices(t *testing.T) {
	appinit.Init()

	apiUrl := "/api/v1.0/orgadmin/org-service"
	reqBody := PublishRequest{
		Name:       "maxprerdomf78maaarrrrx",
		HtmlString: "<p>Hello Pro</p>",
		Components: []string{"3"},
	}
	jsonBytes, err := json.Marshal(reqBody)
	require.Nil(t, err)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonBytes))
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	publish(c)
	require.Equal(t, http.StatusOK, rr.Result().StatusCode, rr.Body.String())
}
