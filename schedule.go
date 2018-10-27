package gonhl

import (
	"fmt"
	"time"
)

const endpointSchedule = "/schedule"

func GetSchedule(c *Client, params *scheduleParams) Schedule {
	var schedule Schedule
	status := c.MakeRequest(endpointSchedule, parseScheduleParams(params), &schedule)
	fmt.Println(status)
	return schedule
}

//API endpoint
type Schedule struct {
	TotalItems   int    `json:"totalItems"`
	TotalEvents  int    `json:"totalEvents"`
	TotalGames   int    `json:"totalGames"`
	TotalMatches int    `json:"totalMatches"`
	Wait         int    `json:"wait"`
	Dates        []Date `json:"dates"`
}

type Date struct {
	Date         string        `json:"date"`
	TotalItems   int           `json:"totalItems"`
	TotalEvents  int           `json:"totalEvents"`
	TotalGames   int           `json:"totalGames"`
	TotalMatches int           `json:"totalMatches"`
	Games        []Game        `json:"games"`
	Events       []interface{} `json:"events"`
	Matches      []interface{} `json:"matches"`
}

type Game struct {
	GamePk   int       `json:"gamePk"`
	Link     string    `json:"link"`
	GameType string    `json:"gameType"`
	Season   string    `json:"season"`
	GameDate time.Time `json:"gameDate"`
	Status   Status    `json:"status"`
	Teams    struct {
		Away GameTeam `json:"away"`
		Home GameTeam `json:"home"`
	} `json:"teams"`
	Linescore  Linescore   `json:"linescore"`
	Venue      Venue       `json:"venue"`
	Tickets    []Ticket    `json:"tickets"`
	Broadcasts []Broadcast `json:"broadcasts"`
	Content    Content     `json:"content"`
}

type Status struct {
	AbstractGameState string `json:"abstractGameState"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	StartTimeTBD      bool   `json:"startTimeTBD"`
}

type GameTeam struct {
	LeagueRecord LeagueRecord `json:"leagueRecord"`
	Score        int          `json:"score"`
	Team         Team         `json:"team"`
}

type Ticket struct {
	TicketType string `json:"ticketType"`
	TicketLink string `json:"ticketLink"`
}

type Broadcast struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Site     string `json:"site"`
	Language string `json:"language"`
}

type Content struct {
	Link string `json:"link"`
}
