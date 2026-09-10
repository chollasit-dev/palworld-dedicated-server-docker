package palworld

type GetPlayersResponse struct {
	Players []*Player `json:"players"`
}

func (c *HTTPClient) GetPlayers() (*GetPlayersResponse, error) {
	return &GetPlayersResponse{}, nil
}

type SaveWorldRequest struct{}

func (c *HTTPClient) SaveWorld(req SaveWorldRequest) error {
	return nil
}

type ShutdownServerRequest struct {
	// Time to wait before shutdown
	Waittime int    `json:"waittime"`
	Message  string `json:"message"`
}

func (c *HTTPClient) ShutdownServer() error {
	return nil
}
