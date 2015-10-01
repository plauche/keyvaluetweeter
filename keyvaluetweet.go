package keyvaluetweeter

import (
	"fmt"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

type KeyValueTweeter struct {
	AccountName string
	Api         *anaconda.TwitterApi
}

func NewKeyValueTweeter(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string, accountName string) *KeyValueTweeter {
	k := new(KeyValueTweeter)
	k.AccountName = accountName
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	k.Api = anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	return k
}

func (k KeyValueTweeter) Insert(key string, value string) (anaconda.Tweet, error) {
	tweetString := fmt.Sprintf("%s #%s", value, key)
	tweet, err := k.Api.PostTweet(tweetString, nil)
	return tweet, err
}

func (k KeyValueTweeter) Find(key string) (string, error) {
	searchString := fmt.Sprintf("%s #%s", k.AccountName, key)
	searchResult, err := k.Api.GetSearch(searchString, nil)
	if err != nil {
		return "", err
	}

	value := ""

	if len(searchResult.Statuses) == 1 {
		value = searchResult.Statuses[0].Text
		value = strings.Replace(value, fmt.Sprintf("#%s", key), "", 1)
		value = strings.TrimSpace(value)
	}

	return value, nil
}

func (k KeyValueTweeter) Delete(api *anaconda.TwitterApi, key string) (err error) {
	searchString := fmt.Sprintf("%s #%s", k.AccountName, key)
	searchResult, err := k.Api.GetSearch(searchString, nil)
	if err != nil {
		return err
	}

	for _, tweet := range searchResult.Statuses {
		_, err := k.Api.DeleteTweet(tweet.Id, true)
		return err
	}

	return nil
}
