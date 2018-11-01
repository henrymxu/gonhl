package gonhl

import (
	"time"
)

const endpointSchedule = "/schedule"

// GetSchedule retrieves the NHL schedule based on ScheduleParams.
// If no parameters are passed, the NHL schedule for the current day is retrieved.
func (c *Client) GetSchedule(params *ScheduleParams) (Schedule, int) {
	var schedule Schedule
	status := c.makeRequest(endpointSchedule, parseScheduleParams(params), &schedule)
	return schedule, status
}

//API endpoint
type Schedule struct {
	TotalItems   int       `json:"totalItems"`
	TotalEvents  int       `json:"totalEvents"`
	TotalGames   int       `json:"totalGames"`
	TotalMatches int       `json:"totalMatches"`
	Wait         int       `json:"wait"`
	Dates        []GameDay `json:"dates"`
}

type GameDay struct {
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
	GamePk   int        `json:"gamePk"`
	Link     string     `json:"link"`
	GameType string     `json:"gameType"`
	Season   string     `json:"season"`
	GameDate time.Time  `json:"gameDate"`
	Status   GameStatus `json:"status"`
	Teams    struct {
		Away GameTeam `json:"away"`
		Home GameTeam `json:"home"`
	} `json:"teams"`
	Linescore       Linescore      `json:"linescore"`
	Venue           Venue          `json:"venue"`
	Tickets         []Ticket       `json:"tickets"`
	Broadcasts      []Broadcast    `json:"broadcasts"`
	RadioBroadcasts RadioBroadcast `json:"radioBroadcasts"`
	Content         Content        `json:"content"`
	Metadata        MetaData       `json:"metadata"`
}

type GameStatus struct {
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

type RadioBroadcast struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Content struct {
	Link string `json:"link"`
}

type MetaData struct {
	IsManuallyScored bool `json:"isManuallyScored"`
	IsSplitSquad     bool `json:"isSplitSquad"`
}
