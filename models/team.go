package models

type Team struct {
	Name string
	Urls *TeamUrls
}

type TeamUrls struct {
	Fixtures string
	Results  string
	Squad    string
	Stats    string
}
