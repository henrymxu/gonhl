package gonhl

import (
	"fmt"
	"time"
)

const endpointConferences = "/conferences"
const endpointConference = endpointConferences + "/%d"
const endpointDivisions = "/divisions"
const endpointDivision = endpointDivisions + "/%d"

const endpointStandings = "/standings"
const endpointStanding = endpointStandings + "/%s"
const endpointStandingTypes = "/standingsTypes"

func GetStandings(c *Client, params *standingsParams) []Standings {
	var standings standings
	endpointCall := endpointStandings
	if params.standingsType != "" {
		endpointCall = fmt.Sprintf(endpointStanding, params.standingsType)
	}
	status := c.MakeRequest(endpointCall, parseStandingsParams(params), &standings)
	fmt.Println(status)
	return standings.Records
}

func GetConferences(c *Client) []Conference {
	var conferences conferences
	status := c.MakeRequest(endpointConferences, nil, &conferences)
	fmt.Println(status)
	return conferences.Conferences
}

func GetConference(c *Client, id int) Conference {
	var conferences conferences
	status := c.MakeRequest(fmt.Sprintf(endpointConference, id), nil, &conferences)
	fmt.Println(status)
	return conferences.Conferences[0]
}

func GetDivisions(c *Client) []Division {
	var divisions divisions
	status := c.MakeRequest(endpointDivisions, nil, &divisions)
	fmt.Println(status)
	return divisions.Divisions
}

func GetDivision(c *Client, id int) Division {
	var divisions divisions
	status := c.MakeRequest(fmt.Sprintf(endpointDivision, id), nil, &divisions)
	fmt.Println(status)
	return divisions.Divisions[0]
}

func GetStandingsTypes(c *Client) []StandingsType {
	var standingsTypes []StandingsType
	status := c.MakeRequest(endpointStandingTypes, nil, &standingsTypes)
	fmt.Println(status)
	return standingsTypes
}

// API Endpoint (Each record represents a division)
type standings struct {
	Records []Standings `json:"records"`
}

// API Endpoint
type divisions struct {
	Divisions []Division `json:"divisions"`
}

// API Endpoint
type conferences struct {
	Conferences []Conference `json:"conferences"`
}

type Standings struct {
	StandingsType string       `json:"standingsType"`
	League        League       `json:"league"`
	Division      Division     `json:"division"`
	Conference    Conference   `json:"conference"`
	TeamRecords   []TeamRecord `json:"teamRecords"`
}

type Conference struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
	ShortName    string `json:"shortName"`
	Active       bool   `json:"active"`
}

type Division struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	NameShort    string `json:"nameShort"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
	Conference   Conference
	Active       bool `json:"active"`
}

type TeamRecord struct {
	Team           Team         `json:"team"`
	LeagueRecord   LeagueRecord `json:"leagueRecord"`
	GoalsAgainst   int          `json:"goalsAgainst"`
	GoalsScored    int          `json:"goalsScored"`
	Points         int          `json:"points"`
	DivisionRank   string       `json:"divisionRank"`
	ConferenceRank string       `json:"conferenceRank"`
	LeagueRank     string       `json:"leagueRank"`
	WildCardRank   string       `json:"wildCardRank"`
	Row            int          `json:"row"`
	GamesPlayed    int          `json:"gamesPlayed"`
	Streak         Streak       `json:"streak"`
	LastUpdated    time.Time    `json:"lastUpdated"`
}

type League struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type LeagueRecord struct {
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Ot     int    `json:"ot"`
	Type   string `json:"type"`
}

type Streak struct {
	StreakType   string `json:"streakType"`
	StreakNumber int    `json:"streakNumber"`
	StreakCode   string `json:"streakCode"`
}

type StandingsType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
