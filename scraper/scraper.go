package scraper

import (
	"fmt"
	"footballScraper/constants"
	"footballScraper/models"
	"footballScraper/utils"
	"footballScraper/utils/logger"
	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	Logger *logger.Logger
	//Teams  map[string]*models.Team
	Leagues map[string]*models.League
}

//Init is responsible of the main logic behind the scraper
func (s *Scraper) Init() {
	s.Logger.Write("Initiating the Football Scraper", logger.LogTypeHeader)

	s.getLeagues()
}

func (s *Scraper) getTeams(league *models.League) {
	s.Leagues[league.Name].Teams = make(map[string]*models.Team)

	startingUrl := fmt.Sprintf("%s%s", constants.BaseUrl, league.Urls.Teams)
	dom := utils.GetDocument(startingUrl, s.Logger)

	dom.Find(".TeamLinks").Each(func(_ int, sel *goquery.Selection) {
		team := &models.Team{Urls: &models.TeamUrls{}}
		team.Name = sel.Find("a h2").Text()
		s.Leagues[league.Name].Teams[team.Name] = team
		sel.Find("span a").Each(func(_ int, linkSel *goquery.Selection) {
			text := linkSel.Text()
			link, ok := linkSel.Attr("href")
			if ok {
				switch text {
				case "Fixtures":
					team.Urls.Fixtures = link
				case "Results":
					team.Urls.Results = link
				case "Squad":
					team.Urls.Squad = link
				case "Stats":
					team.Urls.Stats = link
				}
			}
		})

		s.Leagues[league.Name].Teams[team.Name] = team
		s.getSquad(team.Urls.Squad)
	})
}

func (s *Scraper) getLeagues() {
	s.Leagues = make(map[string]*models.League)
	startingUrl := fmt.Sprintf("%s%s", constants.BaseUrl, constants.LeaguesEP)

	dom := utils.GetDocument(startingUrl, s.Logger)

	//Get the top competitions (Leagues and Cups)
	dom.Find("#topCompetitions").Next().Find(".TeamLinks").Each(func(_ int, sel *goquery.Selection) {
		name := sel.Find("a h2").Text()
		s.Logger.Write(fmt.Sprintf("Extracting data of league: %s", name), logger.LogTypeDebug)
		league := &models.League{Name: name, Urls: &models.LeagueUrls{}}

		//Get links for the competitions
		sel.Find("span a").Each(func(_ int, linkSel *goquery.Selection) {
			text := linkSel.Text()
			link, ok := linkSel.Attr("href")
			if ok {
				switch text {
				case "Fixtures & Results":
					league.Urls.Fixtures = link
				case "Table":
					league.Urls.Table = link
				case "Teams":
					league.Urls.Teams = link
				case "Stats":
					league.Urls.Stats = link
				}
			} else {
				fmt.Println("ERROR")
			}
		})

		//Get teams info
		s.Leagues[league.Name] = league
		s.getTeams(league)
	})
}

func (s *Scraper) getSquad(squadUrl string) {
	startingUrl := fmt.Sprintf("%s%s", constants.BaseUrl, squadUrl)
	dom := utils.GetDocument(startingUrl, s.Logger)
	dom.Find("tbody tr td span").Each(func(i int, sel *goquery.Selection) {
		switch i {
		case 0:
			//get the name of the footballer
			name := sel.Find("a").Text()
			link, ok := sel.Find("a").Attr("href")
			if ok {
				fmt.Println(link)
			}
			number := sel.Find("span").Text()
			fmt.Println(name, number)
		case 1: //position
			position := sel.Text()
			fmt.Println(position)
		case 2: //age
			age := sel.Text()
			fmt.Println(age)
		case 3: //height
			height := sel.Text()
			fmt.Println(height)
		case 4: //weight
			weight := sel.Text()
			fmt.Println(weight)
		case 5: //nationality
			nationality := sel.Text()
			fmt.Println(nationality)
		case 6: //app
			appearances := sel.Text()
			fmt.Println(appearances)
		case 7: //sub app
			subOn := sel.Text()
			fmt.Println(subOn)
		case 8: //saves/goals
		case 9:
		case 10:
		}
	})
}
