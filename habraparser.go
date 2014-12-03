package habraparser

import (
	"fmt"
	"github.com/ungerik/go-rss"
	"log"
)

type Item struct {
	Title string
	Link  string
}

func Read(hub string, count int) ([]Item, error) {
	if count == 0 {
		count = 20
	}
	result, err := rss.Read("http://habrahabr.ru/rss/hub/" + hub)
	if err != nil {
		log.Printf("Error getting habrahabr posts from hab %s", hub)
		return nil, err
	}
	items := make([]Item, count)
	for i := 0; i < count; i++ {
		items[i] = Item{Title: result.Item[i].Title, Link: result.Item[i].Link}
	}
	return items, nil
}
