package news

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type News struct {
	Title   string
	Content string
}

func getSummary() (string, error) {
	allNews, err := getAllNews()
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	sb.WriteString("TODAY'S NEWS (based on expat.cz)\n")
	for i, n := range allNews {
		sb.WriteString(fmt.Sprintf("%d. %s\n", i+1, n.Title))
	}
	sb.WriteString("Use /details {news} or /sarcasm {news} for more")
	return sb.String(), nil
}

func getAllNews() ([]*News, error) {
	file, err := os.Open("news.json")
	if err != nil {
		return nil, err
	}
	var allNews []*News
	err = json.NewDecoder(file).Decode(&allNews)
	if err != nil {
		return nil, err
	}
	return allNews, nil
}

func GetNewsById(id int) (*News, error) {
	allNews, err := getAllNews()
	if err != nil {
		return nil, err
	}
	if len(allNews) < id {
		return nil, nil
	}
	return allNews[id-1], nil
}
