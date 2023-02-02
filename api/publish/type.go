package publish

type PublishRequest struct {
	Name       string `json:"name"`
	HtmlString string `json:"htmlString"`
	Components []string `json:"components"`
}
