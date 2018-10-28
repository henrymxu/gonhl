package gonhl

import (
	"fmt"
)

const endpointTeams = "/teams"
const endpointTeam = endpointTeams + "/%d"
const endpointTeamRoster = endpointTeam + "/roster"
const endpointTeamStats = endpointTeam + "/stats"

// GetTeams retrieves information about NHL teams based on TeamsParams.
// If no parameters are passed, the current NHL teams with minimal information are retrieved.
func (c *Client) GetTeams(params TeamsParams) []Team {
	var teams teams
	status := c.makeRequest(endpointTeams, parseTeamsParams(params), &teams)
	fmt.Println(status)
	return teams.Teams
}

//TODO review this
func (c *Client) GetTeam(params TeamsParams) Team {
	var teams teams
	status := c.makeRequest(fmt.Sprintf(endpointTeam, params.ids[0]), parseTeamsParams(params), &teams)
	fmt.Println(status)
	return teams.Teams[0]
}

// GetRoster retrieves the current roster of an NHL team using a team ID.
// The same roster can be retrieved with the GetTeam(s) endpoints by using ShowTeamRoster() when building teamParams.
func (c *Client) GetRoster(teamId int) Roster {
	var roster Roster
	status := c.makeRequest(fmt.Sprintf(endpointTeamRoster, teamId), nil, &roster)
	fmt.Println(status)
	return roster
}

// GetTeamStats retrieves the current stats of an NHL team using a team ID.
// The same stats can be retrieved with the GetTeam(s) endpoints by using ShowTeamStats() when building teamParams.
func (c *Client) GetTeamStats(teamId int) []AllTeamStats {
	var stats struct{ Stats []AllTeamStats `json:"stats"` }
	status := c.makeRequest(fmt.Sprintf(endpointTeamStats, teamId), nil, &stats)
	fmt.Println(status)
	return stats.Stats
}

// API Endpoint
type teams struct {
	Teams []Team `json:"teams,omitempty"`
}

// API Endpoint
type Roster struct {
	Roster []RosterPlayer `json:"roster"`
	Link   string         `json:"link"`
}

type Team struct {
	ID                   int            `json:"id"`
	Name                 string         `json:"name"`
	Link                 string         `json:"link"`
	Venue                Venue          `json:"venue"`
	Abbreviation         string         `json:"abbreviation"`
	TriCode              string         `json:"triCode"`
	TeamName             string         `json:"teamName"`
	LocationName         string         `json:"locationName"`
	FirstYearOfPlay      string         `json:"firstYearOfPlay,omitempty"`
	Division             Division       `json:"division"`
	Conference           Conference     `json:"conference"`
	Franchise            Franchise      `json:"franchise"`
	TeamStats            []AllTeamStats `json:"teamStats"`
	Roster               Roster         `json:"roster"`
	NextGameSchedule     Schedule       `json:"nextGameSchedule"`
	PreviousGameSchedule Schedule       `json:"previousGameSchedule"`
	ShortName            string         `json:"shortName"`
	OfficialSiteURL      string         `json:"officialSiteUrl"`
	FranchiseID          int            `json:"franchiseId"`
	Active               bool           `json:"active"`
}

type AllTeamStats struct {
	Type StatType `json:"type"`
	Splits []struct {
		Stat TeamStats `json:"stat"`
		Team Team      `json:"team"`
	} `json:"splits"`
}

type TeamStats struct {
	GamesPlayed            int     `json:"gamesPlayed"`
	Wins                   int     `json:"wins"`
	Losses                 int     `json:"losses"`
	Ot                     int     `json:"ot"`
	Pts                    int     `json:"pts"`
	PtPctg                 string  `json:"ptPctg"`
	GoalsPerGame           float64 `json:"goalsPerGame"`
	GoalsAgainstPerGame    float64 `json:"goalsAgainstPerGame"`
	EvGGARatio             float64 `json:"evGGARatio"`
	PowerPlayPercentage    string  `json:"powerPlayPercentage"`
	PowerPlayGoals         float64 `json:"powerPlayGoals"`
	PowerPlayGoalsAgainst  float64 `json:"powerPlayGoalsAgainst"`
	PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
	PenaltyKillPercentage  string  `json:"penaltyKillPercentage"`
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
	FaceOffWinPercentage   string  `json:"faceOffWinPercentage"`
	ShootingPctg           float64 `json:"shootingPctg"`
	SavePctg               float64 `json:"savePctg"`
}

type RosterPlayer struct {
	Person       Player   `json:"person"`
	JerseyNumber string   `json:"jerseyNumber"`
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
