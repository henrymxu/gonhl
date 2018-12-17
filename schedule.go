package gonhl

import (
	"fmt"
	"strings"
	"time"
)

const endpointSchedule = "/schedule"

// GetSchedule retrieves the NHL schedule based on ScheduleParams.
// If no parameters are passed, the NHL schedule for the current^ day is retrieved.
// ^ This may not be true, seems like sometimes the previous day schedule is returned, safest option is to pass a time.
func (c *Client) GetSchedule(params *ScheduleParams) (Schedule, int) {
	var schedule Schedule
	status := c.makeRequest(endpointSchedule, parseScheduleParams(params), &schedule)
	return schedule, status
}

// Create a neatly formatted string for a schedule.
// The strings for each game for each day of the schedule are included.
func (s Schedule) String() string {
	var sb strings.Builder
	for _, date := range s.Dates {
		sb.WriteString(fmt.Sprintf("Date: %s\n", date.Date.Format("2016-01-02")))
		for _, game := range date.Games {
			sb.WriteString(fmt.Sprintf("%s\n", game))
		}
	}
	return sb.String()
}

// Create a neatly formatted string for a game.
// Not all information is included, key information included are: status, start time, teams, and id.
func (g Game) String() string {
	return fmt.Sprintf("%s: %s -> %s(H) vs %s(A) | %d:%d -> %d", g.Status.DetailedState, g.GameDate.Format("3:04PM MST"), g.Teams.Home.Team.Name, g.Teams.Away.Team.Name, g.Teams.Home.Score, g.Teams.Away.Score, g.GamePk)
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
	Date         JsonDate      `json:"date"`
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
	CodedGameState    int    `json:"codedGameState,string"`
	DetailedState     string `json:"detailedState"`
	StatusCode        int    `json:"statusCode,string"`
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
