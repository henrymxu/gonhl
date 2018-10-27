package main

import (
	"fmt"
	"github.com/henrymxu/gonhl"
	"time"
)

func main() {
	client := gonhl.NewClient()
	/*var roster gonhl.Teams
	status := client.MakeRequest("/teams", params, &roster)
	fmt.Println(status)
	fmt.Println(roster)*/
	//client.MakeRequest("/draft/prospects", params)

	//var schedule gonhl.Schedule
	//testTeams(client)
	//testPlayerStats(client)
	//testBoxscore(client)
	testLinescore(client)
}

func testSchedule(client *gonhl.Client) {
	date, _ := gonhl.CreateDateFromString("2018-10-25")
	scheduleParams := gonhl.NewScheduleParams()
	scheduleParams.
		SetDate(date).
		ShowBroadcasts().
		ShowLinescore().
		ShowTicketRetailers()
	schedule := gonhl.GetSchedule(client, scheduleParams)
	fmt.Println(schedule.Dates[0].Games[0].Linescore)
}

func testStandings(client *gonhl.Client) {
	date, _ := gonhl.CreateDateFromString("2019-01-28")
	standingsParams := gonhl.NewStandingsParams()
	standingsParams.
		SetDate(date).
		ShowDetailedRecords().
		SetSeason(2017).SetStandingsType("wildCard")
	standings := gonhl.GetStandings(client, standingsParams)
	fmt.Println(standings)
}

func testStandingsTypes(client *gonhl.Client) {
	standingsTypes := gonhl.GetStandingsTypes(client)
	fmt.Println(standingsTypes)
}

func testConferences(client *gonhl.Client) {
	conferences := gonhl.GetConferences(client)
	fmt.Println(conferences)
	conference := gonhl.GetConference(client, 6) //Eastern
	fmt.Println(conference)
}

func testDivisions(client *gonhl.Client) {
	divisions := gonhl.GetDivisions(client)
	fmt.Println(divisions)
	division := gonhl.GetDivision(client, 18) //Metropolitan
	fmt.Println(division)
}

func testTeams(client *gonhl.Client) {
	teamsParams := gonhl.NewTeamsParams()
	teamsParams.
		SetDetailedRoster().
		SetScheduleNext().
		SetSchedulePrev().
		SetTeamStats()
	teams := gonhl.GetTeams(client, *teamsParams)
	fmt.Println(teams)
}

func testRoster(client *gonhl.Client) {
	roster := gonhl.GetRoster(client, 1)
	fmt.Println(roster)
}

func testStats(client *gonhl.Client) {
	stats := gonhl.GetStats(client, 1)
	fmt.Println(stats)
}

func testPlayerStats(client *gonhl.Client) {
	statsParams := gonhl.NewPlayerParams()
	statsParams.SetId(8476368).
		SetSeason(2017).
		SetStat("byMonth", "vsTeam")
	stats := gonhl.GetPlayerStats(client, statsParams)
	fmt.Println(stats[0].Splits[0].Stat)

	statTypes := gonhl.GetPlayerStatsTypes(client)
	fmt.Println(statTypes)
}

func testDraft(client *gonhl.Client) {
	draft := gonhl.GetDraft(client, 2015)
	fmt.Println(draft)
}

func testProspect(client *gonhl.Client) {
	prospect := gonhl.GetProspect(client, 54223)
	fmt.Println(prospect)
}

func testGameContent(client *gonhl.Client) {
	content := gonhl.GetGameContent(client, 2018020134)
	fmt.Println(content.Editorial.Preview.Items[0].TokenData)
}

func testGameLive(client *gonhl.Client) {
	live := gonhl.GetGameLive(client, 2018020150)
	fmt.Println(live.GameData.Status.DetailedState)
}

func testLivePlays(client *gonhl.Client) {
	location, _ := time.LoadLocation("America/Toronto")
	date := time.Date(2018, 10, 27, 16, 0, 0, 0, location)
	live := gonhl.GetLiveData(client, 2018020150, date)
	fmt.Println(live.ScoringPlays)
}

func testBoxscore(client *gonhl.Client) {
	boxscore := gonhl.GetGameBoxscore(client, 2018020150)
	fmt.Println(boxscore)
}

func testLinescore(client *gonhl.Client) {
	linescore := gonhl.GetGameLinescore(client, 2018020150)
	fmt.Println(linescore)
}