package helpers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/ccallazans/twitter-video-downloader/internal/config"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func ValidateUrl(rawUrl *string) error {
	parsedUrl, err := url.ParseRequestURI(*rawUrl)
	if err != nil {
		return ErrorNotValidUrl
	}

	if parsedUrl.Host != "twitter.com" {
		return ErrorNotTwitterUrl
	}

	if parsedUrl.Path == "/" {
		return ErrorEmptyValue
	}

	return nil
}

func ParseUrl(rawUrl *string) (*string, error) {
	parsedUrl, err := url.ParseRequestURI(*rawUrl)
	if err != nil {
		return nil, err
	}

	splitUrl := strings.Split(parsedUrl.Path, "/status/")
	if len(splitUrl) != 2 {
		return nil, ErrorPath
	}

	_, err = strconv.Atoi(splitUrl[1])
	if err != nil {
		return nil, ErrorPath
	}

	return &splitUrl[1], nil
}

func ParseJsonResponse(bodyByte *[]byte) (*string, error) {

	type Content struct {
		Bitrate     int    `json:"bitrate" validate:"required"`
		ContentType string `json:"content_type" validate:"required"`
		Url         string `json:"url" validate:"required"`
	}

	newData, _, _, err := jsonparser.Get(
		*bodyByte,
		"data",
		"threaded_conversation_with_injections_v2",
		"instructions",
		"[0]",
		"entries",
		"[0]",
		"content",
		"itemContent",
		"tweet_results",
		"result",
		"legacy",
		"extended_entities",
		"media",
		"[0]",
		"video_info",
		"variants",
	)
	if err != nil {
		return nil, err
	}

	var myContent []Content
	err = json.Unmarshal(newData, &myContent)
	if err != nil {
		return nil, err
	}

	higher := 0
	for i, cont := range myContent {
		if cont.Bitrate > higher {
			higher = i
		}
	}

	return &myContent[higher].Url, nil
}

func DownloadFile(url string) (*http.Response, error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func LoadPageResource(url string) error {

	// Load page resource
	u := launcher.New().
		Set("user-data-dir", "path").
		Set("headless").
		Set("no-sandbox").
		Delete("--headless").
		MustLaunch()

	page := rod.New().ControlURL(u).NoDefaultDevice().MustConnect().MustPage(url)
	err := page.WaitLoad()
	if err != nil {
		return err
	}

	getCookies := page.MustCookies()
	sessionId := getCookies[0]
	config.RequestHeader.Set("x-guest-token", sessionId.Value)

	return nil
}
