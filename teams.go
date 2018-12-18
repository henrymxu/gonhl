package gonhl

import (
	"encoding/json"
	"fmt"
)

const endpointTeams = "/teams"
const endpointTeam = endpointTeams + "/%d"
const endpointTeamRoster = endpointTeam + "/roster"
const endpointTeamStats = endpointTeam + "/stats"

// GetTeams retrieves information about NHL teams based on TeamsParams.
// If no parameters are passed, the current NHL teams with minimal information are retrieved.
// Stats must be casted to appropriate type.  Types can be determined using the DisplayName.
func (c *Client) GetTeams(params *TeamsParams) ([]Team, int) {
	var teams teams
	status := c.makeRequest(endpointTeams, parseTeamsParams(*params), &teams)
	return parseTeams(teams.Teams), status
}

// GetTeam retrieves information about a single NHL team based on TeamsParams.
// A team Id must be set.  If multiple Ids are set, only the first value is used.
// Stats must be casted to appropriate type.  Types can be determined using the DisplayName.
func (c *Client) GetTeam(params *TeamsParams) (Team, int) {
	var teams teams
	status := c.makeRequest(fmt.Sprintf(endpointTeam, params.ids[0]), parseTeamsParams(*params), &teams)
	return parseTeams(teams.Teams)[0], status
}

// parseTeams replaces the team struct from the api and with the proper Team struct.
func parseTeams(teams []team) []Team {
	parsedTeams := make([]Team, len(teams))
	for index, team := range teams {
		parsedTeams[index].TeamStats = parseStats(team.TeamStats)
		parsedTeams[index].internalTeam = team.internalTeam
	}
	return parsedTeams
}

// parseStats takes the retrieved team stats json from the api, unmarshals them, then reinserts into the parent struct.
// The parent struct holds an interface{} type and requires reflection to access the proper values of the stat.
// The proper types can be converted to using ConvertTeamStatsToTeamStats and ConvertTeamStatsToTeamRanks.
func parseStats(stats []teamStatsForType) []TeamStatsForType {
	parsedStats := make([]TeamStatsForType, len(stats))
	for statType, stat := range stats {
		parsedStats[statType].Splits = make([]TeamStatsSplit, len(stat.Splits))
		for splitType, split := range stat.Splits {
			var testmap map[string]string
			json.Unmarshal(*split.Stat, &testmap)
			if _, ok := testmap["wins"]; ok {
				var rankedStats TeamStatRanks
				json.Unmarshal(*split.Stat, &rankedStats)
				parsedStats[statType].Splits[splitType].Stat = rankedStats
			} else {
				var actualStats TeamStats
				json.Unmarshal(*split.Stat, &actualStats)
				parsedStats[statType].Splits[splitType].Stat = actualStats
			}
			parsedStats[statType].Splits[splitType].Team = split.Team
		}
		parsedStats[statType].Type = stat.Type
	}
	return parsedStats
}

// GetRoster retrieves the current roster of an NHL team using a team ID.
// The same roster can be retrieved with the GetTeam(s) endpoints by using ShowTeamRoster() when building teamParams.
func (c *Client) GetRoster(teamId int) (Roster, int) {
	var roster Roster
	status := c.makeRequest(fmt.Sprintf(endpointTeamRoster, teamId), nil, &roster)
	return roster, status
}

//TODO Fix this, missing ranking stats
// GetTeamStats retrieves the current stats of an NHL team using a team ID.
// The same stats can be retrieved with the GetTeam(s) endpoints by using ShowTeamStats() when building teamParams.
// Stats must be casted to appropriate type.  Types can be determined using the DisplayName.
func (c *Client) GetTeamStats(teamId int) ([]TeamStatsForType, int) {
	var teamStats struct {
		Stats []teamStatsForType `json:"stats"`
	}
	status := c.makeRequest(fmt.Sprintf(endpointTeamStats, teamId), nil, &teamStats)
	return parseStats(teamStats.Stats), status
}

// API Endpoint
type teams struct {
	Teams []team `json:"teams"`
}

// API Endpoint
type Roster struct {
	Roster []RosterPlayer `json:"roster"`
	Link   string         `json:"link"`
}

type team struct {
	TeamStats []teamStatsForType `json:"teamStats"`
	internalTeam
}

type Team struct {
	TeamStats []TeamStatsForType `json:"teamStats"`
	internalTeam
}

type internalTeam struct {
	ID                   int        `json:"id"`
	Name                 string     `json:"name"`
	Link                 string     `json:"link"`
	Venue                Venue      `json:"venue"`
	Abbreviation         string     `json:"abbreviation"`
	TriCode              string     `json:"triCode"`
	TeamName             string     `json:"teamName"`
	LocationName         string     `json:"locationName"`
	FirstYearOfPlay      string     `json:"firstYearOfPlay"`
	Division             Division   `json:"division"`
	Conference           Conference `json:"conference"`
	Franchise            Franchise  `json:"franchise"`
	Roster               Roster     `json:"roster"`
	NextGameSchedule     Schedule   `json:"nextGameSchedule"`
	PreviousGameSchedule Schedule   `json:"previousGameSchedule"`
	ShortName            string     `json:"shortName"`
	OfficialSiteURL      string     `json:"officialSiteUrl"`
	FranchiseID          int        `json:"franchiseId"`
	Active               bool       `json:"active"`
}

type TeamStatsForType struct {
	Type   StatType         `json:"type"`
	Splits []TeamStatsSplit `json:"splits"`
}

type teamStatsForType struct {
	Type   StatType `json:"type"`
	Splits []struct {
		Stat *json.RawMessage `json:"stat"`
		Team Team             `json:"team"`
	} `json:"splits"`
}

type TeamStatsSplit struct {
	Stat interface{} `json:"stat"`
	Team Team        `json:"team"`
}

type TeamStats struct {
	GamesPlayed            int     `json:"gamesPlayed"`
	Wins                   int     `json:"wins"`
	Losses                 int     `json:"losses"`
	Ot                     int     `json:"ot"`
	Pts                    int     `json:"pts"`
	PtPctg                 float64 `json:"ptPctg,string"`
	GoalsPerGame           float64 `json:"goalsPerGame"`
	GoalsAgainstPerGame    float64 `json:"goalsAgainstPerGame"`
	EvGGARatio             float64 `json:"evGGARatio"`
	PowerPlayPercentage    float64 `json:"powerPlayPercentage,string"`
	PowerPlayGoals         float64 `json:"powerPlayGoals"`
	PowerPlayGoalsAgainst  float64 `json:"powerPlayGoalsAgainst"`
	PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
	PenaltyKillPercentage  float64 `json:"penaltyKillPercentage,string"`
	ShotsPerGame           float64 `json:"shotsPerGame"`
	ShotsAllowed           float64 `json:"shotsAllowed"`
	WinScoreFirst          float64 `json:"winScoreFirst"`
	WinOppScoreFirst       float64 `json:"winOppScoreFirst"`
	WinLeadFirstPer        float64 `json:"winLeadFirstPer"`
	WinLeadSecondPer       float64 `json:"winLeadSecondPer"`
	WinOutshootOpp         float64 `json:"winOutshootOpp"`
	WinOutshotByOpp        float64 `json:"winOutshotByOpp"`
	FaceOffsTaken          float64 `json:"faceOffsTaken"`
	FaceOffsWon            float64 `json:"faceOffsWon"`
	FaceOffsLost           float64 `json:"faceOffsLost"`
	FaceOffWinPercentage   float64 `json:"faceOffWinPercentage,string"`
	ShootingPctg           float64 `json:"shootingPctg"`
	SavePctg               float64 `json:"savePctg"`
}

type TeamStatRanks struct {
	Wins                     string `json:"wins"`
	Losses                   string `json:"losses"`
	Ot                       string `json:"ot"`
	Pts                      string `json:"pts"`
	PtPctg                   string `json:"ptPctg"`
	GoalsPerGame             string `json:"goalsPerGame"`
	GoalsAgainstPerGame      string `json:"goalsAgainstPerGame"`
	EvGGARatio               string `json:"evGGARatio"`
	PowerPlayPercentage      string `json:"powerPlayPercentage"`
	PowerPlayGoals           string `json:"powerPlayGoals"`
	PowerPlayGoalsAgainst    string `json:"powerPlayGoalsAgainst"`
	PowerPlayOpportunities   string `json:"powerPlayOpportunities"`
	PenaltyKillOpportunities string `json:"penaltyKillOpportunities"`
	PenaltyKillPercentage    string `json:"penaltyKillPercentage"`
	ShotsPerGame             string `json:"shotsPerGame"`
	ShotsAllowed             string `json:"shotsAllowed"`
	WinScoreFirst            string `json:"winScoreFirst"`
	WinOppScoreFirst         string `json:"winOppScoreFirst"`
	WinLeadFirstPer          string `json:"winLeadFirstPer"`
	WinLeadSecondPer         string `json:"winLeadSecondPer"`
	WinOutshootOpp           string `json:"winOutshootOpp"`
	WinOutshotByOpp          string `json:"winOutshotByOpp"`
	FaceOffsTaken            string `json:"faceOffsTaken"`
	FaceOffsWon              string `json:"faceOffsWon"`
	FaceOffsLost             string `json:"faceOffsLost"`
	FaceOffWinPercentage     string `json:"faceOffWinPercentage"`
	SavePctRank              string `json:"savePctRank"`
	ShootingPctRank          string `json:"shootingPctRank"`
}

type RosterPlayer struct {
	Person       Player   `json:"person"`
	JerseyNumber int      `json:"jerseyNumber,string"`
	Position     Position `json:"position"`
}

type Venue struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Link     string `json:"link"`
	City     string `json:"city"`
	TimeZone struct {
		ID     string `json:"id"`
		Offset int    `json:"offset"`
		Tz     string `json:"tz"`
	} `json:"timeZone"`
}

type Franchise struct {
	FranchiseID int    `json:"franchiseId"`
	TeamName    string `json:"teamName"`
	Link        string `json:"link"`
}
