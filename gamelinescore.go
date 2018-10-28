package gonhl

import (
	"fmt"
	"time"
)

const endpointGameLinescore = "/game/%d/linescore"

// GetGameLinescore retrieves the linescore from a specific NHL game.
// Linescore contains less information than GetGameBoxscore.
func (c *Client) GetGameLinescore(id int) Linescore {
	var linescore Linescore
	status := c.makeRequest(fmt.Sprintf(endpointGameLinescore, id), nil, &linescore)
	fmt.Println(status)
	return linescore
}

type Linescore struct {
	CurrentPeriod              int      `json:"currentPeriod"`
	CurrentPeriodOrdinal       string   `json:"currentPeriodOrdinal"`
	CurrentPeriodTimeRemaining string   `json:"currentPeriodTimeRemaining"`
	Periods                    []Period `json:"periods"`
	ShootoutInfo               struct {
		Away TeamShootoutInfo `json:"away"`
		Home TeamShootoutInfo `json:"home"`
	} `json:"shootoutInfo"`
	Teams struct {
		Home LinescoreTeam `json:"home"`
		Away LinescoreTeam `json:"away"`
	} `json:"teams"`
	PowerPlayStrength string           `json:"powerPlayStrength"`
	HasShootout       bool             `json:"hasShootout"`
	IntermissionInfo  IntermissionInfo `json:"intermissionInfo"`
	PowerPlayInfo     PowerPlayInfo    `json:"powerPlayInfo"`
}

type Period struct {
	PeriodType string         `json:"periodType"`
	StartTime  time.Time      `json:"startTime"`
	EndTime    time.Time      `json:"endTime,omitempty"`
	Num        int            `json:"num"`
	OrdinalNum string         `json:"ordinalNum"`
	Home       PeriodTeamData `json:"home"`
	Away       PeriodTeamData `json:"away"`
}

type PeriodTeamData struct {
	Goals       int    `json:"goals"`
	ShotsOnGoal int    `json:"shotsOnGoal"`
	RinkSide    string `json:"rinkSide"`
}

type TeamShootoutInfo struct {
	Scores   int `json:"scores"`
	Attempts int `json:"attempts"`
}

type LinescoreTeam struct {
	Team         Team `json:"team"`
	Goals        int  `json:"goals"`
	ShotsOnGoal  int  `json:"shotsOnGoal"`
	GoaliePulled bool `json:"goaliePulled"`
	NumSkaters   int  `json:"numSkaters"`
	PowerPlay    bool `json:"powerPlay"`
}

type IntermissionInfo struct {
	IntermissionTimeRemaining int  `json:"intermissionTimeRemaining"`
	IntermissionTimeElapsed   int  `json:"intermissionTimeElapsed"`
	InIntermission            bool `json:"inIntermission"`
}

type PowerPlayInfo struct {
	SituationTimeRemaining int  `json:"situationTimeRemaining"`
	SituationTimeElapsed   int  `json:"situationTimeElapsed"`
	InSituation            bool `json:"inSituation"`
}
