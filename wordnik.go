package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	apiRoot     = "https://api.wordnik.com/v4/"
	wodEndPoint = apiRoot + "words.json/wordOfTheDay?api_key="
)

type wodJson struct {
	Word        string  `json:"word"`
	Definitions []child `json:"definitions"`
	Examples    []child `json:"examples"`
	Note        string  `json:"note"`
}

func (j wodJson) Definition() string {
	sl := []string{}
	for _, s := range j.Definitions {
		sl = append(sl, s.Text)
	}
	return strings.Join(sl, " / ")
}

func (j wodJson) Example() string {
	sl := []string{}
	for _, s := range j.Examples {
		sl = append(sl, fmt.Sprintf(`"%s"`, s.Text))
	}
	return strings.Join(sl, "\n")
}

func (j wodJson) String() (out string) {
	return strings.Join([]string{
		fmt.Sprintf("%s - %s", j.Word, j.Definition()),
		j.Example(),
		j.Note,
	}, "\n")
}

type child struct {
	Text string `json:"text"`
}

func wod(apiKey string) (string, error) {
	url := wodEndPoint + apiKey
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	j := &wodJson{}

	err = json.Unmarshal(body, &j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}
