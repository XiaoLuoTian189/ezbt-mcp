package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

var (
	ApiToken  string
	BaseURL   string
	Timestamp string
)

func SetApiToken(apiToken string) {
	Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	ApiToken = md5Sum(Timestamp + md5Sum(apiToken))
}

func GetApiToken() string {
	if ApiToken != "" {
		return ApiToken
	}
	if shellToken := os.Getenv("BT_API_TOKEN"); shellToken != "" {
		SetApiToken(shellToken)
		return ApiToken
	}
	return ApiToken
}

func GetBaseURL() string {
	if BaseURL != "" {
		return BaseURL
	}
	BaseURL = os.Getenv("BT_BASE_URL")
	return BaseURL
}

type BTPanel struct {
	BaseURL  string
	APIToken string
}

func NewBTPanel(baseURL string, apiToken string) *BTPanel {
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}
	return &BTPanel{
		BaseURL:  baseURL,
		APIToken: apiToken,
	}
}

func md5Sum(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func (bt *BTPanel) Request(path string, params map[string]string) (*mcp.CallToolResult, error) {
	path = strings.TrimPrefix(path, "/")
	fullURL := fmt.Sprintf("%s%s?request_time=%s&request_token=%s", bt.BaseURL, path, Timestamp, ApiToken)
	if strings.Contains(path, "?") {
		fullURL = fmt.Sprintf("%s%s&request_time=%s&request_token=%s", bt.BaseURL, path, Timestamp, ApiToken)
	}

	formData := strings.Builder{}
	first := true
	for key, value := range params {
		if !first {
			formData.WriteString("&")
		}
		formData.WriteString(key)
		formData.WriteString("=")
		formData.WriteString(value)
		first = false
	}

	req, err := http.NewRequest("POST", fullURL, strings.NewReader(formData.String()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(body)), nil
}
