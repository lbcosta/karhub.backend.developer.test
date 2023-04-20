package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const authUrl = "https://accounts.spotify.com/api/token"

type GetTokenResult struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func Authenticate(c *fiber.Ctx) error {
	data := url.Values{}
	data.Add("grant_type", "client_credentials")
	data.Add("client_id", os.Getenv("SPOTIFY_CLIENT_ID"))
	data.Add("client_secret", os.Getenv("SPOTIFY_CLIENT_SECRET"))

	encodedData := data.Encode()

	req, err := http.NewRequest("POST", authUrl, strings.NewReader(encodedData))
	if err != nil {
		return fiber.NewError(http.StatusBadGateway, err.Error())
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/x-www-form-urlencoded"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return fiber.NewError(http.StatusBadGateway, fmt.Errorf("failed to get token: %s", resp.Status).Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fiber.NewError(http.StatusBadGateway, fmt.Errorf("failed to read token: %w", err).Error())
	}
	defer resp.Body.Close()

	var result GetTokenResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fiber.NewError(http.StatusBadGateway, fmt.Errorf("failed to parse token: %w", err).Error())
	}

	c.Locals("token", result.AccessToken)

	return c.Next()
}
