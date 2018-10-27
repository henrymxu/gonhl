package gonhl

import "fmt"

const endpointPlayer = "/people/%d"
const endpointPlayerStats = "/people/%d/stats/"

const endpointStatTypes = "/statTypes"

func GetPlayer(c *Client, id int) Player {
	var player struct{ People []Player `json:"people"` }
	status := c.MakeRequest(fmt.Sprintf(endpointPlayer, id), nil, &player)
	fmt.Println(status)
	return player.People[0]
}

func GetPlayerStats(c *Client, params *playerParams) []AllSkaterStats {
	var playerStats struct{ Stats []AllSkaterStats `json:"stats"` }
	status := c.MakeRequest(fmt.Sprintf(endpointPlayerStats, params.id), parseParams(params), &playerStats)
	fmt.Println(status)
	return playerStats.Stats
}

func GetPlayerStatsTypes(c *Client) []string {
	var statTypes [] struct{ DisplayName string `json:"displayName"` }
	status := c.MakeRequest(endpointStatTypes, nil, &statTypes)
	fmt.Println(status)
	statTypesString := make([]string, len(statTypes))
	for index, value := range statTypes {
		statTypesString[index] = value.DisplayName
	}
	return statTypesString
}

func parseParams(params *playerParams) map[string]string {
	query := map[string]string{}
	query["id"] = string(params.id)                     // Player id
	query["season"] = createSeasonString(params.season) // Player stats for that season (use the year season started)
	query["stats"] = combineStringArray(params.stat)    // Obtains single season statistics for a player
	return query
}

type playerParams struct {
	id     int
	season int
	stat   []string
}

func NewPlayerParams() *playerParams {
	return &playerParams{}
}

func (pParams *playerParams) SetId(id int) *playerParams {
	pParams.id = id
	return pParams
}

func (pParams *playerParams) SetSeason(season int) *playerParams {
	pParams.season = season
	return pParams
}

func (pParams *playerParams) SetStat(stat ...string) *playerParams {
	pParams.stat = stat
	return pParams
}

type Player struct {
	ID                 int      `json:"id"`
	FullName           string   `json:"fullName"`
	Link               string   `json:"link"`
	FirstName          string   `json:"firstName"`
	LastName           string   `json:"lastName"`
	PrimaryNumber      string   `json:"primaryNumber"`
	BirthDate          string   `json:"birthDate"`
	CurrentAge         int      `json:"currentAge"`
	BirthCity          string   `json:"birthCity"`
	BirthStateProvince string   `json:"birthStateProvince"`
	BirthCountry       string   `json:"birthCountry"`
	Nationality        string   `json:"nationality"`
	Height             string   `json:"height"`
	Weight             int      `json:"weight"`
	Active             bool     `json:"active"`
	AlternateCaptain   bool     `json:"alternateCaptain"`
	Captain            bool     `json:"captain"`
	Rookie             bool     `json:"rookie"`
	ShootsCatches      string   `json:"shootsCatches"`
	RosterStatus       string   `json:"rosterStatus"`
	CurrentTeam        Team     `json:"currentTeam"`
	PrimaryPosition    Position `json:"primaryPosition"`
}

type Position struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}

type AllSkaterStats struct {
	Type struct {
		DisplayName string `json:"displayName"`
	} `json:"type"`
	Splits []struct {
		Season     string      `json:"season"`
		Stat       SkaterStats `json:"stat"`
		GoalieStat GoalieStats `json:"stat"`
	} `json:"splits"`
}

type SkaterStats struct {
	TimeOnIce                   string  `json:"timeOnIce"`
	Assists                     int     `json:"assists"`
	Goals                       int     `json:"goals"`
	Pim                         int     `json:"pim"`
	Shots                       int     `json:"shots"`
	Games                       int     `json:"games"`
	Hits                        int     `json:"hits"`
	PowerPlayGoals              int     `json:"powerPlayGoals"`
	PowerPlayPoints             int     `json:"powerPlayPoints"`
	PowerPlayTimeOnIce          string  `json:"powerPlayTimeOnIce"`
	EvenTimeOnIce               string  `json:"evenTimeOnIce"`
	PenaltyMinutes              string  `json:"penaltyMinutes"`
	FaceOffPct                  float64 `json:"faceOffPct"`
	ShotPct                     float64 `json:"shotPct"`
	GameWinningGoals            int     `json:"gameWinningGoals"`
	OverTimeGoals               int     `json:"overTimeGoals"`
	ShortHandedGoals            int     `json:"shortHandedGoals"`
	ShortHandedPoints           int     `json:"shortHandedPoints"`
	ShortHandedTimeOnIce        string  `json:"shortHandedTimeOnIce"`
	Blocked                     int     `json:"blocked"`
	PlusMinus                   int     `json:"plusMinus"`
	Points                      int     `json:"points"`
	Shifts                      int     `json:"shifts"`
	TimeOnIcePerGame            string  `json:"timeOnIcePerGame"`
	EvenTimeOnIcePerGame        string  `json:"evenTimeOnIcePerGame"`
	ShortHandedTimeOnIcePerGame string  `json:"shortHandedTimeOnIcePerGame"`
	PowerPlayTimeOnIcePerGame   string  `json:"powerPlayTimeOnIcePerGame"`
}

type GoalieStats struct {
	TimeOnIce                  string  `json:"timeOnIce"`
	Ot                         int     `json:"ot"`
	Shutouts                   int     `json:"shutouts"`
	Ties                       int     `json:"ties"`
	Wins                       int     `json:"wins"`
	Losses                     int     `json:"losses"`
	Saves                      int     `json:"saves"`
	PowerPlaySaves             int     `json:"powerPlaySaves"`
	ShortHandedSaves           int     `json:"shortHandedSaves"`
	EvenSaves                  int     `json:"evenSaves"`
	ShortHandedShots           int     `json:"shortHandedShots"`
	EvenShots                  int     `json:"evenShots"`
	PowerPlayShots             int     `json:"powerPlayShots"`
	SavePercentage             float64 `json:"savePercentage"`
	GoalAgainstAverage         float64 `json:"goalAgainstAverage"`
	Games                      int     `json:"games"`
	GamesStarted               int     `json:"gamesStarted"`
	ShotsAgainst               int     `json:"shotsAgainst"`
	GoalsAgainst               int     `json:"goalsAgainst"`
	TimeOnIcePerGame           string  `json:"timeOnIcePerGame"`
	PowerPlaySavePercentage    float64 `json:"powerPlaySavePercentage"`
	ShortHandedSavePercentage  float64 `json:"shortHandedSavePercentage"`
	EvenStrengthSavePercentage float64 `json:"evenStrengthSavePercentage"`
}
