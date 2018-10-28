package gonhl

import (
	"fmt"
)

const endpointPlayer = "/people/%d"
const endpointPlayerStats = "/people/%d/stats/"

const endpointStatTypes = "/statTypes"

// GetPlayer retrieves information about a single NHL player using a player ID.
func (c *Client) GetPlayer(id int) Player {
	var player struct{ People []Player `json:"people"` }
	status := c.makeRequest(fmt.Sprintf(endpointPlayer, id), nil, &player)
	fmt.Println(status)
	return player.People[0]
}

// GetPlayerStats retrieves stats about a single NHL player based on PlayerParams.
// The PlayerParams must not be nil and all fields must be set (id, season, statType).
// To determine if a skater or goalie is retrieved, position specific types can be checked for nil.
func (c *Client) GetPlayerStats(params *PlayerParams) ([]AllPlayerStats) {
	var playerStats struct{ Stats []AllPlayerStats `json:"stats"` }
	status := c.makeRequest(fmt.Sprintf(endpointPlayerStats, params.id), parseParams(params), &playerStats)
	fmt.Println(status)
	return playerStats.Stats
}

// GetPlayerStatsTypes retrieves information about the various enums that can be used when retrieving player stats.
// Pass values retrieved from here to SetStat for PlayerParams.
func (c *Client) GetPlayerStatsTypes() []string {
	var statTypes []StatType
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

type AllPlayerStats struct {
	Type   StatType `json:"type"`
	Splits []struct {
		Season string      `json:"season"`
		Stat   PlayerStats `json:"stat"`
	} `json:"splits"`
}

type StatType struct {
	DisplayName string `json:"displayName"`
}

// Position specific stats are pointers to differentiate between 0 value and nil value.
type PlayerStats struct {
	// Player Stats
	TimeOnIce        string `json:"timeOnIce"`
	Games            int    `json:"games"`
	TimeOnIcePerGame string `json:"timeOnIcePerGame"`

	// Skater Stats
	Assists                     *int     `json:"assists"`
	Goals                       *int     `json:"goals"`
	Pim                         *int     `json:"pim"`
	Shots                       *int     `json:"shots"`
	Hits                        *int     `json:"hits"`
	PowerPlayGoals              *int     `json:"powerPlayGoals"`
	PowerPlayPoints             *int     `json:"powerPlayPoints"`
	PowerPlayTimeOnIce          *string  `json:"powerPlayTimeOnIce"`
	EvenTimeOnIce               *string  `json:"evenTimeOnIce"`
	PenaltyMinutes              *string  `json:"penaltyMinutes"`
	FaceOffPct                  *float64 `json:"faceOffPct"`
	ShotPct                     *float64 `json:"shotPct"`
	GameWinningGoals            *int     `json:"gameWinningGoals"`
	OverTimeGoals               *int     `json:"overTimeGoals"`
	ShortHandedGoals            *int     `json:"shortHandedGoals"`
	ShortHandedPoints           *int     `json:"shortHandedPoints"`
	ShortHandedTimeOnIce        *string  `json:"shortHandedTimeOnIce"`
	Blocked                     *int     `json:"blocked"`
	PlusMinus                   *int     `json:"plusMinus"`
	Points                      *int     `json:"points"`
	Shifts                      *int     `json:"shifts"`
	EvenTimeOnIcePerGame        *string  `json:"evenTimeOnIcePerGame"`
	ShortHandedTimeOnIcePerGame *string  `json:"shortHandedTimeOnIcePerGame"`
	PowerPlayTimeOnIcePerGame   *string  `json:"powerPlayTimeOnIcePerGame"`

	// Goalie Stats
	Ot                         *int     `json:"ot"`
	Shutouts                   *int     `json:"shutouts"`
	Ties                       *int     `json:"ties"`
	Wins                       *int     `json:"wins"`
	Losses                     *int     `json:"losses"`
	Saves                      *int     `json:"saves"`
	PowerPlaySaves             *int     `json:"powerPlaySaves"`
	ShortHandedSaves           *int     `json:"shortHandedSaves"`
	EvenSaves                  *int     `json:"evenSaves"`
	ShortHandedShots           *int     `json:"shortHandedShots"`
	EvenShots                  *int     `json:"evenShots"`
	PowerPlayShots             *int     `json:"powerPlayShots"`
	SavePercentage             *float64 `json:"savePercentage"`
	GoalAgainstAverage         *float64 `json:"goalAgainstAverage"`
	GamesStarted               *int     `json:"gamesStarted"`
	ShotsAgainst               *int     `json:"shotsAgainst"`
	GoalsAgainst               *int     `json:"goalsAgainst"`
	PowerPlaySavePercentage    *float64 `json:"powerPlaySavePercentage"`
	ShortHandedSavePercentage  *float64 `json:"shortHandedSavePercentage"`
	EvenStrengthSavePercentage *float64 `json:"evenStrengthSavePercentage"`
}
