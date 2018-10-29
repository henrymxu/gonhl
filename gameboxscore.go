package gonhl

import (
	"fmt"
)

const endpointGameBoxscore = "/game/%d/boxscore"

// GetGameBoxscore retrieves the boxscore from a specific NHL game.
// Boxscore contains less information than GetGameLiveFeed.
func (c *Client) GetGameBoxscore(id int) Boxscore {
	var boxscore Boxscore
	status := c.makeRequest(fmt.Sprintf(endpointGameBoxscore, id), nil, &boxscore)
	fmt.Println(status)
	return boxscore
}

// API endpoint
type Boxscore struct {
	Teams struct {
		Away BoxscoreTeam `json:"away"`
		Home BoxscoreTeam `json:"home"`
	} `json:"teams"`
	Officials []Official `json:"officials"`
}

type BoxscoreTeam struct {
	Team      Team `json:"team"`
	TeamStats struct {
		TeamSkaterStats TeamSkaterStats `json:"teamSkaterStats"`
	} `json:"teamStats"`
	Players    map[string]Skater `json:"players"`
	Goalies    []int             `json:"goalies"`
	Skaters    []int             `json:"skaters"`
	OnIce      []int             `json:"onIce"`
	OnIcePlus  []OnIcePlus       `json:"onIcePlus"`
	Scratches  []int             `json:"scratches"`
	PenaltyBox []PenaltyBox      `json:"penaltyBox"`
	Coaches    []Coach           `json:"coaches"`
}

type Skater struct {
	Person       Player   `json:"person"`
	JerseyNumber string   `json:"jerseyNumber"`
	Position     Position `json:"position"`
	Stats        struct {
		SkaterStats SkaterGameStats `json:"skaterStats"`
		GoalieStats GoalieGameStats `json:"goalieStats"`
	} `json:"stats"`
}

type Coach struct {
	Person   BasicPerson `json:"person"`
	Position Position    `json:"position"`
}

type Official struct {
	Official     BasicPerson `json:"official"`
	OfficialType string      `json:"officialType"`
}

type BasicPerson struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Link     string `json:"link"`
}

type OnIcePlus struct {
	PlayerID      int `json:"playerId"`
	ShiftDuration int `json:"shiftDuration"`
	Stamina       int `json:"stamina"`
}

type PenaltyBox struct {
	ID            int    `json:"id"`
	TimeRemaining string `json:"timeRemaining"`
	Active        bool   `json:"active"`
}

type TeamSkaterStats struct {
	Goals                  int     `json:"goals"`
	Pim                    int     `json:"pim"`
	Shots                  int     `json:"shots"`
	PowerPlayPercentage    string  `json:"powerPlayPercentage"`
	PowerPlayGoals         float64 `json:"powerPlayGoals"`
	PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
	FaceOffWinPercentage   string  `json:"faceOffWinPercentage"`
	Blocked                int     `json:"blocked"`
	Takeaways              int     `json:"takeaways"`
	Giveaways              int     `json:"giveaways"`
	Hits                   int     `json:"hits"`
}

type SkaterGameStats struct {
	TimeOnIce            string `json:"timeOnIce"`
	Assists              int    `json:"assists"`
	Goals                int    `json:"goals"`
	Shots                int    `json:"shots"`
	Hits                 int    `json:"hits"`
	PowerPlayGoals       int    `json:"powerPlayGoals"`
	PowerPlayAssists     int    `json:"powerPlayAssists"`
	PenaltyMinutes       int    `json:"penaltyMinutes"`
	FaceOffWins          int    `json:"faceOffWins"`
	FaceoffTaken         int    `json:"faceoffTaken"`
	Takeaways            int    `json:"takeaways"`
	Giveaways            int    `json:"giveaways"`
	ShortHandedGoals     int    `json:"shortHandedGoals"`
	ShortHandedAssists   int    `json:"shortHandedAssists"`
	Blocked              int    `json:"blocked"`
	PlusMinus            int    `json:"plusMinus"`
	EvenTimeOnIce        string `json:"evenTimeOnIce"`
	PowerPlayTimeOnIce   string `json:"powerPlayTimeOnIce"`
	ShortHandedTimeOnIce string `json:"shortHandedTimeOnIce"`
}

type GoalieGameStats struct {
	TimeOnIce                  string  `json:"timeOnIce"`
	Assists                    int     `json:"assists"`
	Goals                      int     `json:"goals"`
	Pim                        int     `json:"pim"`
	Shots                      int     `json:"shots"`
	Saves                      int     `json:"saves"`
	PowerPlaySaves             int     `json:"powerPlaySaves"`
	ShortHandedSaves           int     `json:"shortHandedSaves"`
	EvenSaves                  int     `json:"evenSaves"`
	ShortHandedShotsAgainst    int     `json:"shortHandedShotsAgainst"`
	EvenShotsAgainst           int     `json:"evenShotsAgainst"`
	PowerPlayShotsAgainst      int     `json:"powerPlayShotsAgainst"`
	SavePercentage             float64 `json:"savePercentage"`
	ShortHandedSavePercentage  float64 `json:"shortHandedSavePercentage"`
	EvenStrengthSavePercentage float64 `json:"evenStrengthSavePercentage"`
}
