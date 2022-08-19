package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ccallazans/twitter-video-downloader/internal/config"
	"github.com/ccallazans/twitter-video-downloader/internal/helpers"
	"github.com/labstack/echo/v4"
)

func (app *Application) GetUrl(c echo.Context) error {
	args := c.Param("url")
	if args == "" {
		return c.JSON(http.StatusBadRequest, helpers.ErrorEmptyValue.Error())
	}

	err := helpers.ValidateUrl(&args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	parsedUrl, err := helpers.ParseUrl(&args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	sendUrl := strings.Replace(config.TwitterGraphQLString, "{replaceID}", *parsedUrl, 1)

	client := http.Client{}
	req, err := http.NewRequest("GET", sendUrl, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	req.Header = config.RequestHeader
	res, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer res.Body.Close()

	// Read Req
	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	videoUrl, err := helpers.ParseJsonResponse(&bodyByte)

	return c.JSON(http.StatusOK, videoUrl)
}
