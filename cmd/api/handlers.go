package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ccallazans/twitter-video-downloader/internal/config"
	"github.com/ccallazans/twitter-video-downloader/internal/helpers"
	"github.com/labstack/echo/v4"
)

func (app *Application) GetHome(c echo.Context) error {
	return c.JSON(http.StatusOK, "Download Twitter Videos")
}

func (app *Application) GetUrl(c echo.Context) error {
	// Get url param
	args := c.Param("url")
	if args == "" {
		return c.JSON(http.StatusBadRequest, helpers.ErrorEmptyValue.Error())
	}

	// Validate url
	err := helpers.ValidateUrl(&args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Load page resource
	err = helpers.LoadPageResource(args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Parse url
	parsedUrl, err := helpers.ParseUrl(&args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Send Request
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
	
	// Read response
	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Get url video
	videoUrl, err := helpers.ParseJsonResponse(&bodyByte)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Download video
	downloadResponse, err := helpers.DownloadFile(*videoUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "erro video")
	}
	defer downloadResponse.Body.Close()

	// Writer the body to file
	c.Response().Header().Set("Content-Disposition", "attachment; filename=twitter-video.mp4")
	_, err = io.Copy(c.Response().Writer, downloadResponse.Body)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, "Download Successful!")
}
