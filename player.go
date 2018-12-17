package gonhl

import (
	"encoding/json"
	"fmt"
)

const endpointPlayer = "/people/%d"
const endpointPlayerStats = "/people/%d/stats/"

const endpointStatTypes = "/statTypes"

// GetPlayer retrieves information about a single NHL player using a player ID.
func (c *Client) GetPlayer(id int) (Player, int) {
	var player struct{ People []Player `json:"people"` }
	status := c.makeRequest(fmt.Sprintf(endpointPlayer, id), nil, &player)
	return player.People[0], status
}

// GetPlayerStats retrieves stats about a single NHL player based on PlayerParams.
// The PlayerParams must not be nil and all fields must be set (id, season, statType).
// To determine if a skater or goalie is retrieved, use IsPlayerGoalie.
// Stats must be casted to appropriate type.  Types can be determined using the DisplayName.

// Internally this method takes the retrieved player stats json from the api, unmarshals them, then reinserts into the parent struct.
// The parent struct holds an interface{} type and requires reflection to access the proper values of the stat.
// The proper types can be converted to using ConvertPlayerStatsToSkaterStats and ConvertStatsToGoalieStats.
func (c *Client) GetPlayerStats(params *PlayerParams) ([]PlayerStatsForType, int) {
	var playerStats struct{ Stats []playerStatsForType `json:"stats"` }
	status := c.makeRequest(fmt.Sprintf(endpointPlayerStats, params.id), parseParams(params), &playerStats)
	parsedStats := make([]PlayerStatsForType, len(playerStats.Stats))
	for statType, stat := range playerStats.Stats {
		parsedStats[statType].ID = params.id
		parsedStats[statType].Type = stat.Type
		parsedStats[statType].Splits = make([]StatsSplit, len(stat.Splits))
		for splitType, split := range stat.Splits {
			switch stat.Type.DisplayName {
			case "regularSeasonStatRankings":
				var testmap map[string]string
				json.Unmarshal(*split.Stat, &testmap)
				if _, ok := testmap["rankGoals"]; ok {
					var skaterStat SkaterStatsByRank
					json.Unmarshal(*split.Stat, &skaterStat)
					parsedStats[statType].Splits[splitType].Stat = skaterStat
				} else {
					var goalieStat GoalieStatsByRank
					json.Unmarshal(*split.Stat, &goalieStat)
					parsedStats[statType].Splits[splitType].Stat = goalieStat
				}
			case "goalsByGameSituation":
				var testmap map[string]int
				json.Unmarshal(*split.Stat, &testmap)
				if _, ok := testmap["gameWinningGoals"]; ok {
					var skaterStat SkaterGoalsBySituation
					json.Unmarshal(*split.Stat, &skaterStat)
					parsedStats[statType].Splits[splitType].Stat = skaterStat
				} else {
					var goalieStat GoalsBySituation
					json.Unmarshal(*split.Stat, &goalieStat)
					parsedStats[statType].Splits[splitType].Stat = goalieStat
				}
			default:
				var testmap map[string]interface{}
				json.Unmarshal(*split.Stat, &testmap)
				if _, ok := testmap["faceOffPct"]; ok {
					var skaterStat SkaterStats
					json.Unmarshal(*split.Stat, &skaterStat)
					parsedStats[statType].Splits[splitType].Stat = skaterStat
				} else {
					var goalieStat GoalieStats
					json.Unmarshal(*split.Stat, &goalieStat)
					parsedStats[statType].Splits[splitType].Stat = goalieStat
				}
			}
			parsedStats[statType].Splits[splitType].internalStatsSplit = split.internalStatsSplit
		}
	}
	return parsedStats, status
}

// GetPlayerStatsTypes retrieves information about the various enums that can be used when retrieving player stats.
// Pass values retrieved from here to SetStat for PlayerParams.
func (c *Client) GetPlayerStatsTypes() ([]string, int) {
	var statTypes []StatType
	status := c.makeRequest(endpointStatTypes, nil, &statTypes)
	statTypesString := make([]string, len(statTypes))
	for index, value := range statTypes {
		statTypesString[index] = value.DisplayName
	}
	return statTypesString, status
}

type Player struct {
	ID                 int      `json:"id"`
	FullName           string   `json:"fullName"`
	Link               string   `json:"link"`
	FirstName          string   `json:"firstName"`
	LastName           string   `json:"lastName"`
	PrimaryNumber      int      `json:"primaryNumber,string"`
	BirthDate          JsonDate `json:"birthDate"`
	CurrentAge         int      `json:"currentAge"`
	BirthCity          string   `json:"birthCity"`
	BirthStateProvince string   `json:"birthStateProvince"`
	BirthCountry       string   `json:"birthCountry"`
	Nationality        string   `json:"nationality"`
	Height             Height   `json:"height,string"`
	Weight             int      `json:"weight"`
	Active             bool     `json:"active"`
	AlternateCaptain   bool     `json:"alternateCaptain"`
	Captain            bool          `json:"captain"`
	Rookie             bool          `json:"rookie"`
	ShootsCatches      string        `json:"shootsCatches"`
	RosterStatus       string        `json:"rosterStatus"`
	CurrentTeam        Team          `json:"currentTeam"`
	PrimaryPosition    Position      `json:"primaryPosition"`
}

type Position struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}

type PlayerStatsForType struct {
	ID     int
	Type   StatType     `json:"type"`
	Splits []StatsSplit `json:"splits"`
}

type playerStatsForType struct {
	Type   StatType `json:"type"`
	Splits []struct {
		Stat *json.RawMessage `json:"stat"`
		internalStatsSplit
	} `json:"splits"`
}

type internalStatsSplit struct {
	Season             string              `json:"season"`
	IsHome             *bool               `json:"isHome"`
	IsWin              *bool               `json:"isWin"`
	IsOT               *bool               `json:"isOT"`
	Month              *int                `json:"month"`
	Opponent           StatSplitIdentifier `json:"opponent"`
	OpponentDivision   StatSplitIdentifier `json:"opponentDivision"`
	OpponentConference StatSplitIdentifier `json:"opponentConference"`
}

type StatsSplit struct {
	Stat interface{} `json:"stat"`
	internalStatsSplit
}

type StatSplitIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type StatType struct {
	DisplayName string `json:"displayName"`
}

// PlayerStats holds values that are used in both SkaterStats and GoalieStats.
// Only used as an anonymous struct.
type PlayerStats struct {
	TimeOnIce        string `json:"timeOnIce"`
	Games            int    `json:"games"`
	TimeOnIcePerGame string `json:"timeOnIcePerGame"`
}

type SkaterStats struct {
	PlayerStats
	Assists                     int     `json:"assists"`
	Goals                       int     `json:"goals"`
	Pim                         int     `json:"pim"`
	Shots                       int     `json:"shots"`
	Hits                        int     `json:"hits"`
	PowerPlayGoals              int     `json:"powerPlayGoals"`
	PowerPlayPoints             int     `json:"powerPlayPoints"`
	PowerPlayTimeOnIce          string  `json:"powerPlayTimeOnIce"`
	EvenTimeOnIce               string  `json:"evenTimeOnIce"`
	PenaltyMinutes              int     `json:"penaltyMinutes,string"`
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
	EvenTimeOnIcePerGame        string  `json:"evenTimeOnIcePerGame"`
	ShortHandedTimeOnIcePerGame string  `json:"shortHandedTimeOnIcePerGame"`
	PowerPlayTimeOnIcePerGame   string  `json:"powerPlayTimeOnIcePerGame"`
}

type GoalieStats struct {
	PlayerStats
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
	GamesStarted               int     `json:"gamesStarted"`
	ShotsAgainst               int     `json:"shotsAgainst"`
	GoalsAgainst               int     `json:"goalsAgainst"`
	PowerPlaySavePercentage    float64 `json:"powerPlaySavePercentage"`
	ShortHandedSavePercentage  float64 `json:"shortHandedSavePercentage"`
	EvenStrengthSavePercentage float64 `json:"evenStrengthSavePercentage"`
}

type SkaterStatsByRank struct {
	RankPowerPlayGoals   string `json:"rankPowerPlayGoals"`
	RankBlockedShots     string `json:"rankBlockedShots"`
	RankAssists          string `json:"rankAssists"`
	RankShotPct          string `json:"rankShotPct"`
	RankGoals            string `json:"rankGoals"`
	RankHits             string `json:"rankHits"`
	RankPenaltyMinutes   string `json:"rankPenaltyMinutes"`
	RankShortHandedGoals string `json:"rankShortHandedGoals"`
	RankPlusMinus        string `json:"rankPlusMinus"`
	RankShots            string `json:"rankShots"`
	RankPoints           string `json:"rankPoints"`
	RankOvertimeGoals    string `json:"rankOvertimeGoals"`
	RankGamesPlayed      string `json:"rankGamesPlayed"`
}

type GoalieStatsByRank struct {
	ShotsAgainst        string `json:"shotsAgainst"`
	Ot                  string `json:"ot"`
	PenaltyMinutes      string `json:"penaltyMinutes"`
	TimeOnIce           string `json:"timeOnIce"`
	ShutOuts            string `json:"shutOuts"`
	Saves               string `json:"saves"`
	Losses              string `json:"losses"`
	GoalsAgainst        string `json:"goalsAgainst"`
	SavePercentage      string `json:"savePercentage"`
	Games               string `json:"games"`
	GoalsAgainstAverage string `json:"goalsAgainstAverage"`
	Wins                string `json:"wins"`
}

// GoalsBySituation holds all the relevant information for goalies
// and most information for skaters.
type GoalsBySituation struct {
	GoalsInFirstPeriod       int `json:"goalsInFirstPeriod"`
	GoalsInSecondPeriod      int `json:"goalsInSecondPeriod"`
	GoalsInThirdPeriod       int `json:"goalsInThirdPeriod"`
	GoalsInOvertime          int `json:"goalsInOvertime"`
	ShootOutGoals            int `json:"shootOutGoals"`
	ShootOutShots            int `json:"shootOutShots"`
	GoalsTrailingByOne       int `json:"goalsTrailingByOne"`
	GoalsTrailingByTwo       int `json:"goalsTrailingByTwo"`
	GoalsTrailingByThreePlus int `json:"goalsTrailingByThreePlus"`
	GoalsWhenTied            int `json:"goalsWhenTied"`
	GoalsLeadingByOne        int `json:"goalsLeadingByOne"`
	GoalsLeadingByTwo        int `json:"goalsLeadingByTwo"`
	GoalsLeadingByThreePlus  int `json:"goalsLeadingByThreePlus"`
	PenaltyGoals             int `json:"penaltyGoals"`
	PenaltyShots             int `json:"penaltyShots"`
}

type SkaterGoalsBySituation struct {
	GameWinningGoals int `json:"gameWinningGoals"`
	EmptyNetGoals    int `json:"emptyNetGoals"`
	GoalsBySituation
}
