package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type IdentityPreset int

const (
	SneakyDetector IdentityPreset = iota
	Backup
)

// TODO: Add avatar URL
func GetIdentityByPreset(preset IdentityPreset) (username, avatarUrl string) {
	switch preset {
	case SneakyDetector:
		username = "Sneaky Detector Bot"
		avatarUrl = ""
	case Backup:
		username = "Backup Bot"
		avatarUrl = ""
	}
	return username, avatarUrl
}

type DiscordClient struct {
	Username  string
	AvatarURL string
	Content   *string
}

func NewClient() *DiscordClient {
	http.DefaultClient.Timeout = time.Second * 3 // 3 seconds
	return &DiscordClient{}
}

// Use provided profile data instead of default from at the time the webhook was created
func (c *DiscordClient) OverrideIdentity(username, avatarUrl string) {
	c.Username = username
	c.AvatarURL = avatarUrl
}

func (c *DiscordClient) CreateMessageContent(content string) error {
	msgLens := len(content)
	if msgLens > 2000 {
		return fmt.Errorf("Content message must be less than 2000 characters; Received: %d", msgLens)
	}
	c.Content = &content
	return nil
}

func (c *DiscordClient) CreateMessageCardContent() {
}

func (c *DiscordClient) Send() error {
	url := os.Getenv("WEBHOOK_URL")
	if url == "" {
		return fmt.Errorf("Invalid WEBHOOK_URL; Received: %s", url)
	}

	var req *bytes.Buffer
	if err := json.NewEncoder(req).Encode(*c); err != nil {
		return err
	}

	resp, err := http.Post(url, "applicaion/json", req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Request failed with status %d", resp.StatusCode)
	}

	return nil
}
