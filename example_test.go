package gonhl_test

import (
	"fmt"
	"github.com/henrymxu/gonhl"
	"time"
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

// Example of retrieving Player information.
// Information about Connor McDavid.
func ExampleClient_GetPlayer() {
	client := gonhl.NewClient()
	player, _ := client.GetPlayer(8478402)
	fmt.Println(player.FullName)
	fmt.Println(player.Height.Format())
	fmt.Println(player.Weight)
	fmt.Println(player.BirthDate.Format("2006-01-02"))
	// Output:
	// Connor McDavid
	// 6' 1"
	// 193
	// 1997-01-13
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
	playerStat, _ := gonhl.ConvertPlayerStatsToSkaterStats(player[0].Splits[0].Stat)
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
	playerStat, _ := gonhl.ConvertPlayerStatsToGoalieStats(player[0].Splits[0].Stat)
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
	teams, _ := client.GetTeams(teamsParams)
	fmt.Println(teams[0].Name)
	fmt.Println(teams[0].TeamStats[0].Type.DisplayName)
	fmt.Println(teams[0].Roster.Link)
	teamStat, _ := gonhl.ConvertTeamStatsToTeamStats(teams[0].TeamStats[0].Splits[0].Stat)
	fmt.Println(teamStat.GamesPlayed > 0)
	rankedStat, _ := gonhl.ConvertTeamStatsToTeamRanks(teams[0].TeamStats[0].Splits[1].Stat)
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
	team, _ := client.GetTeam(teamsParams)
	fmt.Println(team.Name)
	// Output: New York Islanders
}

func ExampleClient_GetTeamStats() {
	client := gonhl.NewClient()
	teamStats, _ := client.GetTeamStats(1)
	fmt.Println(teamStats[0].Splits[0].Team.ID)
	fmt.Println(teamStats[0].Type.DisplayName)
	stats, _ := teamStats[0].Splits[0].Stat.(gonhl.TeamStats)
	fmt.Println(stats.Wins > 0)
	fmt.Println(teamStats[1].Type.DisplayName)
	ranks, _ := teamStats[1].Splits[0].Stat.(gonhl.TeamStatRanks)
	fmt.Println(ranks.Wins != "")
	// Output:
	// 1
	// statsSingleSeason
	// true
	// regularSeasonStatRankings
	// true
}

// Example of retrieving a prospect.
// Prospect information for Connor McDavid.
func ExampleClient_GetProspect() {
	client := gonhl.NewClient()
	prospect, _ := client.GetProspect(54223)
	fmt.Println(prospect.FullName)
	// Output: Connor Mcdavid
}

// Example of retrieving a schedule.
// Be careful of dates as they may not be the same timezone.
func ExampleClient_GetSchedule() {
	client := gonhl.NewClient()
	date, _ := time.Parse("2006-01-02", "2018-12-04")
	scheduleParams := gonhl.NewScheduleParams().
		SetDate(date).ShowTicketRetailers()
	schedule, _ := client.GetSchedule(scheduleParams)
	fmt.Println(schedule.Dates[0].Games[0].GameDate.Format("2006-01-02"))
	fmt.Println(schedule.Dates[0].Date.Format("2006-01-02"))
	fmt.Println(schedule.Dates[0].Games[0].Venue.Name)
	fmt.Println(schedule.Dates[0].Games[0].Teams.Home.Score)
	fmt.Println(schedule.Dates[0].Games[0].Teams.Away.Score)
	fmt.Println(len(schedule.Dates[0].Games[0].Tickets))
	fmt.Println(len(schedule.Dates[0].Games[0].Broadcasts))
	// Output:
	// 2018-12-05
	// 2018-12-04
	// BB&T Center
	// 5
	// Boston Bruins
	// 23
	// 0
}

// Example of retrieving only the live data of a game.
func ExampleClient_GetGameLiveData() {
	client := gonhl.NewClient()
	feed, _ := client.GetGameLiveData(2018020514)
	fmt.Println(feed.Boxscore.Teams.Home.Team.Name)
	fmt.Println(feed.Boxscore.Teams.Home.OnIcePlus[0].PlayerID)
	fmt.Println(feed.Plays.CurrentPlay.About.DateTime.String())
	// Output:
	// Montréal Canadiens
	// 8470642
	// 2018-12-18 03:00:57 +0000 UTC
}

func ExampleClient_GetGameLiveFeed() {
	client := gonhl.NewClient()
	feed, _ := client.GetGameLiveFeed(2018020514)
	fmt.Println(feed.LiveData.Boxscore.Teams.Home.Team.Name)
	fmt.Println(feed.LiveData.Boxscore.Teams.Home.OnIcePlus[0].PlayerID)
	// Output:
	// Montréal Canadiens
	// 8470642
}

func ExampleClient_GetGamePlaysAndPlayers() {
	client := gonhl.NewClient()
	location, _ := time.LoadLocation("America/Toronto")
	date := time.Date(2018, 10, 27, 16, 0, 0, 0, location)
	liveData, _ := client.GetGameLiveDataDiff(2018020150, date)
	fmt.Println(liveData.Plays.AllPlays[15].Result.Description)
	fmt.Println(liveData.Plays.AllPlays[15].Result.Event)
	fmt.Println(liveData.Plays.AllPlays[15].Result.EventCode)
	fmt.Println(liveData.Plays.AllPlays[15].Result.EventTypeID)
	fmt.Println(liveData.Boxscore.Teams.Home.OnIcePlus[0].PlayerID)
	fmt.Println(liveData.Boxscore.Teams.Home.Players["ID8476887"].Person.FullName)
	// Output:
	// Zack Kassian hit Anthony Bitetto
	// Hit
	// NSH55
	// HIT
	// 8474056
	// Filip Forsberg
}