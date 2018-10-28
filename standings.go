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

// GetStandings retrieves the NHL schedule based on StandingsParams.
// If no parameters are passed, the current NHL standings are retrieved.
func GetStandings(c *Client, params *StandingsParams) []Standings {
	var standings standings
	endpointCall := endpointStandings
	if params.standingsType != "" {
		endpointCall = fmt.Sprintf(endpointStanding, params.standingsType)
	}
	status := c.makeRequest(endpointCall, parseStandingsParams(params), &standings)
	fmt.Println(status)
	return standings.Records
}

// GetConferences retrieves information about the NHL conferences.
func GetConferences(c *Client) []Conference {
	var conferences conferences
	status := c.makeRequest(endpointConferences, nil, &conferences)
	fmt.Println(status)
	return conferences.Conferences
}

// GetConference retrieves information about a specific NHL conference using a conference ID.
func GetConference(c *Client, id int) Conference {
	var conferences conferences
	status := c.makeRequest(fmt.Sprintf(endpointConference, id), nil, &conferences)
	fmt.Println(status)
	return conferences.Conferences[0]
}

// GetDivisions retrieves information about the NHL divisions.
func GetDivisions(c *Client) []Division {
	var divisions divisions
	status := c.makeRequest(endpointDivisions, nil, &divisions)
	fmt.Println(status)
	return divisions.Divisions
}

// GetDivision retreives information about a specific NHL division using a division ID.
func GetDivision(c *Client, id int) Division {
	var divisions divisions
	status := c.makeRequest(fmt.Sprintf(endpointDivision, id), nil, &divisions)
	fmt.Println(status)
	return divisions.Divisions[0]
}

// GetStandingsTypes retrieves information about the various enums that can be used when retrieving NHL standings.
// Pass values retrieved from here to SetStandingsType for StandingsParams.
func GetStandingsTypes(c *Client) []StandingsType {
	var standingsTypes []StandingsType
	status := c.makeRequest(endpointStandingTypes, nil, &standingsTypes)
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
