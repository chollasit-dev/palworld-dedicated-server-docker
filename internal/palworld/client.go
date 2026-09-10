package palworld

type HTTPClient struct {
	Username string
	Password string
	URL      string
}

// TODO:
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{}
}
