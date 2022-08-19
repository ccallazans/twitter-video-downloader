package helpers

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
)

var (
	ErrorEmptyValue    = errors.New("empty value")
	ErrorNotValidUrl   = errors.New("not a valid url")
	ErrorNotTwitterUrl = errors.New("not a twitter url")
	ErrorPath          = errors.New("not a valid video url")
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

	videoUrl, err := jsonparser.GetString(
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
		"[0]",
		"url",
	)
	if err != nil {
		return nil, err
	}

	return &videoUrl, nil
}
