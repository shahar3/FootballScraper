package utils

import (
	"fmt"
	"footballScraper/utils/logger"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetDocument(url string, myLogger *logger.Logger) *goquery.Document {
	myLogger.Write(fmt.Sprintf("Get the DOM of the url: %s", url), logger.LogTypeDebug)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return dom
}
