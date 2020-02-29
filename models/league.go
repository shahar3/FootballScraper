package models

type League struct {
	Name  string
	Teams map[string]*Team
	Urls  *LeagueUrls
}

type LeagueUrls struct {
	Fixtures string
	Table    string
	Teams    string
	Stats    string
}
