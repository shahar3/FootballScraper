package main

import (
	footballScraper "footballScraper/scraper"
	"footballScraper/utils/logger"
)

func main() {
	scraper := footballScraper.Scraper{
		Logger: &logger.Logger{
			Tag: "NBA Scraper 1.0",
		},
	}
	scraper.Init()
}
