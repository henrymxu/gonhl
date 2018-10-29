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
	conference := client.GetConference(6)
	fmt.Println(conference.Name)
	// Output: Eastern
}

func ExampleClient_GetDivision() {
	client := gonhl.NewClient()
	division := client.GetDivision(18)
	fmt.Println(division.Name)
	// Output: Metropolitan
}

func ExampleClient_GetDraft() {
	client := gonhl.NewClient()
	draft := client.GetDraft(2018)
	fmt.Println(draft[0].DraftYear)
	// Output: 2018
}

// Example of retrieving Skater Stats.
// Stats for Connor McDavid point total in the 2017 season.
func ExampleClient_GetPlayerStats() {
	client := gonhl.NewClient()
	playerParams := gonhl.NewPlayerParams().
		SetId(8478402).
		SetStat("statsSingleSeason").
		SetSeason(2017)
	player := client.GetPlayerStats(playerParams)
	fmt.Println(player[0].Type.DisplayName)

	playerStat, _ := player[0].Splits[0].Stat.(gonhl.PlayerStats)
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
	player := client.GetPlayerStats(playerParams)
	fmt.Println(player[0].Type.DisplayName)
	playerStat, _ := player[0].Splits[0].Stat.(gonhl.PlayerStats)
	fmt.Println(playerStat.SavePercentage != nil)
	fmt.Println(playerStat.Wins != nil)
	// Output:
	// vsConference
	// true
	// true
}

func ExampleClient_GetTeams() {
	client := gonhl.NewClient()
	teamsParams := gonhl.NewTeamsParams().
		ShowDetailedRoster().
		ShowTeamStats().
		SetTeamIds(1, 2, 3)
	teams := client.GetTeams(*teamsParams)
	fmt.Println(teams[0].Name)
	fmt.Println(teams[0].TeamStats[0].Type.DisplayName)
	fmt.Println(teams[0].Roster.Link)
	// Output:
	// New Jersey Devils
	// statsSingleSeason
	// /api/v1/teams/1/roster
}

func ExampleClient_GetTeam() {
	//TODO want to validate schedule params
	client := gonhl.NewClient()
	teamsParams := gonhl.NewTeamsParams().
		ShowDetailedRoster().
		ShowScheduleNext().
		ShowSchedulePrev().
		ShowTeamStats().
		SetTeamIds(2)
	team := client.GetTeam(*teamsParams)
	fmt.Println(team.Name)
	// Output: New York Islanders
}

func ExampleClient_GetProspect() {
	client := gonhl.NewClient()
	prospect := client.GetProspect(54223)
	fmt.Println(prospect.FullName)
	// Output: Connor McDavid
}