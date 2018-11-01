package gonhl_test

import (
	"fmt"
	"github.com/henrymxu/gonhl"
)

func ExampleNewClient() {
	client := gonhl.NewClient()
	fmt.Println(client)
}

func ExampleClient_GetConference() {
	client := gonhl.NewClient()
	conference, _ := client.GetConference(6)
	fmt.Println(conference.Name)
	// Output: Eastern
}

func ExampleClient_GetDivision() {
	client := gonhl.NewClient()
	division, _ := client.GetDivision(18)
	fmt.Println(division.Name)
	// Output: Metropolitan
}

func ExampleClient_GetDraft() {
	client := gonhl.NewClient()
	draft, _ := client.GetDraft(2018)
	fmt.Println(draft[0].DraftYear)
	// Output: 2018
}

// Example of retrieving Skater Stats.
// Stats for Connor McDavid point total in the 2017 season.
func ExampleClient_GetPlayerStats_skater() {
	client := gonhl.NewClient()
	playerParams := gonhl.NewPlayerParams().
		SetId(8478402).
		SetStat("statsSingleSeason").
		SetSeason(2017)
	player, _ := client.GetPlayerStats(playerParams)
	fmt.Println(player[0].Type.DisplayName)

	playerStat, _ := player[0].Splits[0].Stat.(gonhl.SkaterStats)
	fmt.Println(player[0].Splits[0].Season)
	fmt.Println(playerStat.Points)
	// Output:
	// statsSingleSeason
	// 20172018
	// 108
}

// Example of retrieving Goalie Stats.
// Stats for Henrik Lundqvist save percentage and wins in the in progress 2018 season.
func ExampleClient_GetPlayerStats_goalie() {
	client := gonhl.NewClient()
	playerParams := gonhl.NewPlayerParams().
		SetId(8468685).
		SetStat("vsConference", "vsDivision").
		SetSeason(2018)
	player, _ := client.GetPlayerStats(playerParams)
	fmt.Println(player[0].Type.DisplayName)
	playerStat, _ := player[0].Splits[0].Stat.(gonhl.GoalieStats)
	fmt.Println(playerStat.SavePercentage > 0)
	fmt.Println(playerStat.Saves > 0)
	// Output:
	// vsConference
	// true
	// true
}

// Example of retrieving multiple teams.
// Gets detailed roster, and stats (values and rank) for 3 NHL teams.
func ExampleClient_GetTeams() {
	client := gonhl.NewClient()
	teamsParams := gonhl.NewTeamsParams().
		ShowDetailedRoster().
		ShowTeamStats().
		SetTeamIds(1, 2, 3)
	teams, _ := client.GetTeams(*teamsParams)
	fmt.Println(teams[0].Name)
	fmt.Println(teams[0].TeamStats[0].Type.DisplayName)
	fmt.Println(teams[0].Roster.Link)
	teamStat, _ := teams[0].TeamStats[0].Splits[0].Stat.(gonhl.TeamStats)
	fmt.Println(teamStat.GamesPlayed > 0)
	rankedStat, _ := teams[0].TeamStats[0].Splits[1].Stat.(gonhl.TeamStatsRank)
	fmt.Println(rankedStat.Wins != "")
	// Output:
	// New Jersey Devils
	// statsSingleSeason
	// /api/v1/teams/1/roster
	// true
	// true
}

// Example of retrieving a single team.
// Gets detailed roster, scheduling, and stats (values and rank) for the New York Islanders.
func ExampleClient_GetTeam() {
	//TODO want to validate schedule params
	client := gonhl.NewClient()
	teamsParams := gonhl.NewTeamsParams().
		ShowDetailedRoster().
		ShowScheduleNext().
		ShowSchedulePrev().
		ShowTeamStats().
		SetTeamIds(2)
	team, _ := client.GetTeam(*teamsParams)
	fmt.Println(team.Name)
	// Output: New York Islanders
}

// Example of retrieving a prospect.
// Prospect information for Connor McDavid.
func ExampleClient_GetProspect() {
	client := gonhl.NewClient()
	prospect, _ := client.GetProspect(54223)
	fmt.Println(prospect.FullName)
	// Output: Connor Mcdavid
}