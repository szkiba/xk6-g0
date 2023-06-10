package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

func Default() error {
	doc, err := goquery.NewDocument("https://test.k6.io")
	if err != nil {
		return err
	}

	logrus.Info(doc.Find("h1.title span.text-blue").Text())

	return nil
}
