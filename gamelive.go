package gonhl

import (
	"fmt"
	"time"
)

const endpointGameLive = "/game/%d/feed/live"
const endpointGameLiveDiff = "/game/%d/feed/live/diffPatch"

func GetGameLive(c *Client, id int) LiveFeed {
	var live LiveFeed
	status := c.MakeRequest(fmt.Sprintf(endpointGameLive, id), nil, &live)
	fmt.Println(status)
	return live
}

func GetLiveData(c *Client, id int, time time.Time) Plays {
	var live LiveFeed
	status := c.MakeRequest(fmt.Sprintf(endpointGameLiveDiff, id), map[string]string{
		"startTimecode": createTimeStamp(time),
	}, &live)
	fmt.Println(status)
	return live.LiveData.Plays
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
		Status Status `json:"status"`
		Teams struct {
			Away Team `json:"away"`
			Home Team `json:"home"`
		} `json:"teams"`
		Players map[string]Skater `json:"players"`
		Venue Venue `json:"venue"`
	} `json:"gameData"`
	LiveData struct {
		Plays Plays `json:"plays"`
		Linescore Linescore `json:"linescore"`
		Boxscore Boxscore `json:"boxscore"`
		Decisions Decisions `json:"decisions"`
	} `json:"liveData"`
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
	} `json:"players,omitempty"`
	Result      Result `json:"result"`
	About       About  `json:"about"`
	Coordinates struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"coordinates"`
	Team Team `json:"team,omitempty"`
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