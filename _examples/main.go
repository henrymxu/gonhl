package main

import (
	"fmt"
	"github.com/henrymxu/gonhl"
	"time"
)

func main() {
	client := gonhl.NewClient()
	testLivePlays(client)
}

func testSchedule(client *gonhl.Client) {
	date, _ := gonhl.CreateDateFromString("2018-10-25")
	scheduleParams := gonhl.NewScheduleParams()
	scheduleParams.
		SetDate(date).
		ShowBroadcasts().
		ShowLinescore().
		ShowTicketRetailers()
	schedule, _ := client.GetSchedule(scheduleParams)
	fmt.Println(schedule.Dates[0].Games[0].Linescore)
}

func testStandings(client *gonhl.Client) {
	date, _ := gonhl.CreateDateFromString("2019-01-28")
	standingsParams := gonhl.NewStandingsParams()
	standingsParams.
		SetDate(date).
		ShowDetailedRecords().
		SetSeason(2017).SetStandingsType("wildCard")
	standings, _ := client.GetStandings(standingsParams)
	fmt.Println(standings)
}

func testStandingsTypes(client *gonhl.Client) {
	standingsTypes, _ := client.GetStandingsTypes()
	fmt.Println(standingsTypes)
}

func testConferences(client *gonhl.Client) {
	conferences, _ := client.GetConferences()
	fmt.Println(conferences)
	conference, _ := client.GetConference(6) //Eastern
	fmt.Println(conference)
}

func testDivisions(client *gonhl.Client) {
	divisions, _ := client.GetDivisions()
	fmt.Println(divisions)
	division, _ := client.GetDivision(18) //Metropolitan
	fmt.Println(division)
}

func testTeams(client *gonhl.Client) {
	teamsParams := gonhl.NewTeamsParams()
	teamsParams.
		ShowDetailedRoster().
		ShowScheduleNext().
		ShowSchedulePrev().
		ShowTeamStats()
	teams, _ := client.GetTeams(teamsParams)
	fmt.Println(teams)
}

func testRoster(client *gonhl.Client) {
	roster, _ := client.GetRoster(1)
	fmt.Println(roster)
}

func testStats(client *gonhl.Client) {
	stats, _ := client.GetTeamStats(1)
	fmt.Println(stats)
}

func testPlayerStats(client *gonhl.Client) {
	statsParams := gonhl.NewPlayerParams()
	statsParams.SetId(8476368).
		SetSeason(2017).
		SetStat("byMonth", "vsTeam")
	stats, _ := client.GetPlayerStats(statsParams)
	fmt.Println(stats[0].Splits[0].Stat)

	statTypes, _ := client.GetPlayerStatsTypes()
	fmt.Println(statTypes)
}

func testDraft(client *gonhl.Client) {
	draft, _ := client.GetDraft(2015)
	fmt.Println(draft)
}

func testProspect(client *gonhl.Client) {
	prospect, _ := client.GetProspect(54223)
	fmt.Println(prospect)
}

func testGameContent(client *gonhl.Client) {
	content, _ := client.GetGameContent(2018020134)
	fmt.Println(content.Editorial.Preview.Items[0].TokenData)
}

func testGameLive(client *gonhl.Client) {
	live, _ := client.GetGameLiveFeed(2018020150)
	fmt.Println(live.GameData.Status.DetailedState)
}

func testLivePlays(client *gonhl.Client) {
	location, _ := time.LoadLocation("America/Toronto")
	date := time.Date(2018, 10, 27, 16, 0, 0, 0, location)
	live, _ := client.GetGamePlays(2018020150, date)
	fmt.Println(live.ScoringPlays)
}

func testBoxscore(client *gonhl.Client) {
	boxscore, _ := client.GetGameBoxscore(2018020150)
	fmt.Println(boxscore)
}

func testLinescore(client *gonhl.Client) {
	linescore, _ := client.GetGameLinescore(2018020150)
	fmt.Println(linescore)
}