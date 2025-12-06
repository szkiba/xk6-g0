package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

func Default() error {
	resp, err := http.Get("https://httpbin.org")
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	logrus.Info(doc.Find("div.info h2.title").Text())

	return nil
}
