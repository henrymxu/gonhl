package gonhl

import (
	"encoding/json"
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
// To determine if a skater or goalie is retrieved, use IsPlayerGoalie.
// Stats must be casted to appropriate type.  Types can be determined using the DisplayName.
func (c *Client) GetPlayerStats(params *PlayerParams) ([]PlayerStatsForType) {
	var playerStats struct{ Stats []playerStatsForType `json:"stats"` }
	status := c.makeRequest(fmt.Sprintf(endpointPlayerStats, params.id), parseParams(params), &playerStats)
	fmt.Println(status)
	parsedStats := make([]PlayerStatsForType, len(playerStats.Stats))
	for statType, stat := range playerStats.Stats {
		parsedStats[statType].Type = stat.Type
		parsedStats[statType].Splits = make([]StatSplits, len(stat.Splits))
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
		}
	}
	return parsedStats
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

type playerStatsForType struct {
	Type   StatType `json:"type"`
	Splits []struct {
		Season             string              `json:"season"`
		Stat               *json.RawMessage    `json:"stat"`
		IsHome             *bool               `json:"isHome"`
		IsWin              *bool               `json:"isWin"`
		IsOT               *bool               `json:"isOT"`
		Month              *int                `json:"month"`
		Opponent           StatSplitIdentifier `json:"opponent"`
		OpponentDivision   StatSplitIdentifier `json:"opponentDivision"`
		OpponentConference StatSplitIdentifier `json:"opponentConference"`
	} `json:"splits"`
}

type PlayerStatsForType struct {
	Type   StatType     `json:"type"`
	Splits []StatSplits `json:"splits"`
}

type StatSplits struct {
	Season string      `json:"season"`
	Stat   interface{} `json:"stat"`

	IsHome *bool `json:"isHome"`

	IsWin *bool `json:"isWin"`
	IsOT  *bool `json:"isOT"`

	Month *int `json:"month"`

	Opponent           StatSplitIdentifier `json:"opponent"`
	OpponentDivision   StatSplitIdentifier `json:"opponentDivision"`
	OpponentConference StatSplitIdentifier `json:"opponentConference"`
}

type StatSplitIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type StatType struct {
	DisplayName string `json:"displayName"`
}

type PlayerStats struct {
	TimeOnIce        string `json:"timeOnIce"`
	Games            int    `json:"games"`
	TimeOnIcePerGame string `json:"timeOnIcePerGame"`
}

// Position specific stats are pointers to differentiate between 0 value and nil value.
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
