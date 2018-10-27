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

func parseScheduleParams(params *scheduleParams) map[string]string {
	query := map[string]string{}
	expand := expandQuery("schedule", map[string]bool{
		"broadcasts": params.broadcasts,
		"ticket":     params.ticket,
		"linescore":  params.linescore,
	})
	query["expand"] = expand
	if params.teamId != -1 {
		query["teamId"] = string(params.teamId)
	}
	if &params.date != nil {
		query["date"] = CreateDateFromTime(params.date)
	} else if &params.startDate != nil {
		query["startDate"] = CreateDateFromTime(params.startDate)
		query["endDate"] = CreateDateFromTime(params.endDate)
	}
	return query
}

type scheduleParams struct {
	broadcasts bool      // Shows the broadcasts of the game
	linescore  bool      // Linescore for completed games
	ticket     bool      // Provides the different places to buy tickets for the upcoming games
	teamId     int       // Limit results to a specific team. Team ids can be found through the teams endpoint
	date       time.Time // Single defined date for the search
	startDate  time.Time // Start date for the search
	endDate    time.Time // End date for the search
}

func NewScheduleParams() *scheduleParams {
	return &scheduleParams{
		teamId: -1,
	}
}

func (sp *scheduleParams) ShowBroadcasts() *scheduleParams {
	sp.broadcasts = true
	return sp
}

func (sp *scheduleParams) ShowLinescore() *scheduleParams {
	sp.linescore = true
	return sp
}

func (sp *scheduleParams) ShowTicketRetailers() *scheduleParams {
	sp.ticket = true
	return sp
}

func (sp *scheduleParams) SetTeamId(teamId int) *scheduleParams {
	sp.teamId = teamId
	return sp
}

func (sp *scheduleParams) SetDate(date ...time.Time) *scheduleParams {
	if len(date) == 1 {
		sp.date = date[0]
	} else {
		sp.startDate = date[0]
		sp.endDate = date[1]
	}
	return sp
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
