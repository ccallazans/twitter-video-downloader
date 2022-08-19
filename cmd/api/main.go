package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/labstack/echo/v4"
)

type Application struct {
	logger *log.Logger
	server *echo.Echo
}

func main() {

	logger := log.Default()

	app := Application{
		logger: logger,
	}

	app.NewRouter()
	
	err := app.server.Start(":5000")
	if err != nil {
		log.Fatalln(err)
	}
}

func parse() {
	userUrl := "https://twitter.com/zlatansincero/status/1560312147007180800"
	queryParse := strings.Split(userUrl, "status/")[1]
	postId := strings.Split(queryParse, "?")[0]

	log.Println(postId)

	// url := "https://twitter.com/i/api/graphql/UrnAThw_XuPHLdlayDWxfQ/TweetDetail?variables=%7B%22focalTweetId%22%3A%221560416964228546560%22%2C%22with_rux_injections%22%3Afalse%2C%22includePromotedContent%22%3Atrue%2C%22withCommunity%22%3Atrue%2C%22withQuickPromoteEligibilityTweetFields%22%3Atrue%2C%22withBirdwatchNotes%22%3Afalse%2C%22withSuperFollowsUserFields%22%3Atrue%2C%22withDownvotePerspective%22%3Afalse%2C%22withReactionsMetadata%22%3Afalse%2C%22withReactionsPerspective%22%3Afalse%2C%22withSuperFollowsTweetFields%22%3Atrue%2C%22withVoice%22%3Atrue%2C%22withV2Timeline%22%3Atrue%7D&features=%7B%22dont_mention_me_view_api_enabled%22%3Atrue%2C%22interactive_text_enabled%22%3Atrue%2C%22responsive_web_uc_gql_enabled%22%3Atrue%2C%22vibe_api_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Afalse%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D"
	url := "https://twitter.com/i/api/graphql/UrnAThw_XuPHLdlayDWxfQ/TweetDetail?variables=%7B%22focalTweetId%22%3A%22" + postId + "%22%2C%22with_rux_injections%22%3Afalse%2C%22includePromotedContent%22%3Atrue%2C%22withCommunity%22%3Atrue%2C%22withQuickPromoteEligibilityTweetFields%22%3Atrue%2C%22withBirdwatchNotes%22%3Afalse%2C%22withSuperFollowsUserFields%22%3Atrue%2C%22withDownvotePerspective%22%3Afalse%2C%22withReactionsMetadata%22%3Afalse%2C%22withReactionsPerspective%22%3Afalse%2C%22withSuperFollowsTweetFields%22%3Atrue%2C%22withVoice%22%3Atrue%2C%22withV2Timeline%22%3Atrue%7D&features=%7B%22dont_mention_me_view_api_enabled%22%3Atrue%2C%22interactive_text_enabled%22%3Atrue%2C%22responsive_web_uc_gql_enabled%22%3Atrue%2C%22vibe_api_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Afalse%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D"

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("error creating request client")
	}

	req.Header = http.Header{
		"accept":          {"*/*"},
		"accept-language": {"en-US,en;q=0.9"},
		"authorization":   {"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"},
		"content-type":    {"application/json"},
		"x-guest-token":   {"1560434423581007875"},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer res.Body.Close()

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	videoUrl, err := jsonparser.GetString(
		bodyByte,
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

	c := string(videoUrl)
	log.Println(c)
}
