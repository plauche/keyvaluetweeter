package keyvaluetweeter

import (
	"fmt"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

func Insert(a *anaconda.TwitterApi, key string, value string) anaconda.Tweet {
	tweetString := fmt.Sprintf("%s #%s", value, key)
	tweet, err := a.PostTweet(tweetString, nil)
	if err != nil {
		panic(err)
	}
	return tweet
}

func Find(api *anaconda.TwitterApi, key string) string {
	accountName := "welpwatchout"
	searchString := fmt.Sprintf("%s #%s", accountName, key)
	search_result, err := api.GetSearch(searchString, nil)
	if err != nil {
		panic(err)
	}

	value := ""

	if len(search_result.Statuses) == 1 {
		value = search_result.Statuses[0].Text
		value = strings.Replace(value, fmt.Sprintf("#%s", key), "", 1)
		value = strings.TrimSpace(value)
	}

	return value
}
