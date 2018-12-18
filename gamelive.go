package gonhl

import (
	"fmt"
	"time"
)

const endpointGameLive = "/game/%d/feed/live"
const endpointGameLiveDiff = "/game/%d/feed/live/diffPatch"

// GetGameLiveFeed retrieves the live feed from a specific NHL game.
// LiveFeed contains all information from Boxscore and Linescore.
// LiveFeed also contains play by play information
func (c *Client) GetGameLiveFeed(id int) (LiveFeed, int) {
	var live LiveFeed
	status := c.makeRequest(fmt.Sprintf(endpointGameLive, id), nil, &live)
	return live, status
}

// GetGameLiveData retrieves only the live data portion from a specific NHL game.
// This should be a faster call than GetGameLiveFeed due to the smaller unmarshalling.
// The endpoint call is still the exact same as GetGameLiveFeed due to the limitations of the API
func (c *Client) GetGameLiveData(id int) (LiveData, int) {
	var live BasicLiveFeed
	status := c.makeRequest(fmt.Sprintf(endpointGameLive, id), nil, &live)
	return live.LiveData, status
}

// GetGamePlays retrieves only the newest plays from a specific NHL game.
// All plays after a certain time (not relative to game) are retrieved.
func (c *Client) GetGamePlays(id int, time time.Time) (Plays, int) {
	var live LiveFeed
	status := c.makeRequest(fmt.Sprintf(endpointGameLiveDiff, id), map[string]string{
		"startTimecode": createTimeStamp(time),
	}, &live)
	return live.LiveData.Plays, status
}

type LiveFeed struct {
	GamePk    int    `json:"gamePk"`
	Link      string `json:"link"`
	MetaData  struct {
		Wait      int    `json:"wait"`
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
	GameData struct {
		Game GameHeader `json:"game"`
		Datetime struct {
			DateTime    time.Time `json:"dateTime"`
			EndDateTime time.Time `json:"endDateTime"`
		} `json:"datetime"`
		Status GameStatus `json:"status"`
		Teams  struct {
			Away Team `json:"away"`
			Home Team `json:"home"`
		} `json:"teams"`
		Players map[string]Skater `json:"players"`
		Venue Venue `json:"venue"`
	} `json:"gameData"`
	LiveData LiveData `json:"liveData"`
}

type BasicLiveFeed struct {
	LiveData LiveData `json:"liveData"`
}

type LiveData struct {
	Plays Plays `json:"plays"`
	Linescore Linescore `json:"linescore"`
	Boxscore Boxscore `json:"boxscore"`
	Decisions Decisions `json:"decisions"`
}

type GameHeader struct {
	Pk     int    `json:"pk"`
	Season string `json:"season"`
	Type   string `json:"type"`
}

type Plays struct {
	AllPlays []Play `json:"allPlays"`
	ScoringPlays  []int `json:"scoringPlays"`
	PenaltyPlays  []int `json:"penaltyPlays"`
	PlaysByPeriod []struct {
		StartIndex int   `json:"startIndex"`
		Plays      []int `json:"plays"`
		EndIndex   int   `json:"endIndex"`
	} `json:"playsByPeriod"`
	CurrentPlay Play `json:"currentPlay"`
}

type Play struct {
	Players []struct {
		Player     Player `json:"player"`
		PlayerType string `json:"playerType"`
	} `json:"players"`
	Result      Result `json:"result"`
	About       About  `json:"about"`
	Coordinates struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"coordinates"`
	Team Team `json:"team"`
}

type Decisions struct {
	Winner     BasicPerson `json:"winner"`
	Loser      BasicPerson `json:"loser"`
	FirstStar  BasicPerson `json:"firstStar"`
	SecondStar BasicPerson `json:"secondStar"`
	ThirdStar  BasicPerson `json:"thirdStar"`
}

type Result struct {
	Event       string `json:"event"`
	EventCode   string `json:"eventCode"`
	EventTypeID string `json:"eventTypeId"`
	Description string `json:"description"`
}

type About struct {
	EventIdx            int       `json:"eventIdx"`
	EventID             int       `json:"eventId"`
	Period              int       `json:"period"`
	PeriodType          string    `json:"periodType"`
	OrdinalNum          string    `json:"ordinalNum"`
	PeriodTime          string    `json:"periodTime"`
	PeriodTimeRemaining string    `json:"periodTimeRemaining"`
	DateTime            time.Time `json:"dateTime"`
	Goals               struct {
		Away int `json:"away"`
		Home int `json:"home"`
	} `json:"goals"`
}