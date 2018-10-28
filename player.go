package gonhl

import "fmt"

const endpointPlayer = "/people/%d"
const endpointPlayerStats = "/people/%d/stats/"

const endpointStatTypes = "/statTypes"

// GetPlayer retrieves information about a single NHL player using a player ID.
func GetPlayer(c *Client, id int) Player {
	var player struct{ People []Player `json:"people"` }
	status := c.makeRequest(fmt.Sprintf(endpointPlayer, id), nil, &player)
	fmt.Println(status)
	return player.People[0]
}

// GetPlayerStats retrieves stats about a single NHL player based on PlayerParams.
// The PlayerParams must not be nil and PlayerParams.Id must not be nil.
func GetPlayerStats(c *Client, params *PlayerParams) []AllSkaterStats {
	var playerStats struct{ Stats []AllSkaterStats `json:"stats"` }
	status := c.makeRequest(fmt.Sprintf(endpointPlayerStats, params.id), parseParams(params), &playerStats)
	fmt.Println(status)
	return playerStats.Stats
}

// GetPlayerStatsTypes retrieves information about the various enums that can be used when retrieving player stats.
// Pass values retrieved from here to SetStat for PlayerParams.
func GetPlayerStatsTypes(c *Client) []string {
	var statTypes [] struct{ DisplayName string `json:"displayName"` }
	status := c.makeRequest(endpointStatTypes, nil, &statTypes)
	fmt.Println(status)
	statTypesString := make([]string, len(statTypes))
	for index, value := range statTypes {
		statTypesString[index] = value.DisplayName
	}
	return statTypesString
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
